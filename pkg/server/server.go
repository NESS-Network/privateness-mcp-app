package server

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	wt "github.com/quic-go/webtransport-go"

	"github.com/jeff-bouchard/privateness-mcp-app/pkg/auth"
	"github.com/jeff-bouchard/privateness-mcp-app/pkg/meter"
	"github.com/jeff-bouchard/privateness-mcp-app/pkg/billing"
)

type Config struct {
	Listen   string
	BasePath string
	Rates    billing.Rates
}

// Minimal self-signed cert generator for dev use.
func devCertPaths() (certPath, keyPath string, err error) {
	base := "devcert"
	os.MkdirAll(base, 0o755)
	certPath = filepath.Join(base, "cert.pem")
	keyPath = filepath.Join(base, "key.pem")
	if _, err := os.Stat(certPath); err == nil { return certPath, keyPath, nil }
	// NOTE: For brevity, we do not generate certs here.
	// Provide your own certs or mount via env TLS_CERT/TLS_KEY.
	return certPath, keyPath, nil
}

func Start(cfg Config) error {
	listen := cfg.Listen
	if listen == "" { listen = ":443" }

	certFile := os.Getenv("TLS_CERT")
	keyFile := os.Getenv("TLS_KEY")

	if certFile == "" || keyFile == "" {
		// Try dev certs (must exist); else fail gracefully.
		cf, kf, _ := devCertPaths()
		certFile, keyFile = cf, kf
	}

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil { return err }

	srv := &http.Server{Addr: listen}

	wtServer := wt.Server{
		H3: wt.H3Server{
			Server: srv,
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{certificate}, MinVersion: tls.VersionTLS13},
		},
	}

	http.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade to WebTransport session
		sess, err := wtServer.Upgrade(w, r)
		if err != nil {
			log.Println("upgrade error:", err)
			return
		}
		go handleSession(sess, cfg)
	})

	log.Printf("listening on %s (HTTP/3 + WebTransport)\n", listen)
	return wtServer.ListenAndServe()
}

func handleSession(sess *wt.Session, cfg Config) {
	ctx := sess.Context()
	// First incoming bidirectional stream used for auth
	stream, err := sess.AcceptStream(ctx)
	if err != nil { log.Println("accept stream:", err); _ = sess.Close(); return }

	m := meter.New()
	// Wrap stream I/O for metering
	// Note: webtransport-go exposes Readable/Writeable via stream.(io.ReadWriteCloser)
	rw := meter.Wrap(stream, m)

	clientPubKey, err := auth.Authenticate(rw)
	if err != nil { log.Println("auth failed:", err); _ = sess.Close(); return }

	start := time.Now()
	log.Printf("session auth ok: %s\n", clientPubKey)

	// Handle more streams concurrently
	go func() {
		for {
			st, err := sess.AcceptStream(ctx)
			if err != nil { log.Println("stream end:", err); return }
			go handleGenericStream(st) // TODO route by initial JSON header
		}
	}()

	// Wait for session to end
	<-ctx.Done()

	// Billing on close
	cost := billing.Cost(m.BytesIn(), m.BytesOut(), time.Since(start), cfg.Rates)
	if cost > 0 {
		if err := billing.Charge(clientPubKey, cost); err != nil {
			log.Println("billing error:", err)
		}
	}
	log.Printf("session closed. in=%d out=%d dur=%s cost=%.6f\n", m.BytesIn(), m.BytesOut(), time.Since(start), cost)
}

func handleGenericStream(st wt.Stream) {
	// TODO: Inspect first JSON frame to decide handler (cmd/resp/notify/file/dht/transport)
	// For skeleton, just echo length
	buf := make([]byte, 4096)
	n, _ := st.Read(buf)
	st.Write([]byte("ok:"))
	st.Write([]byte(string(n)))
	st.Close()
}
