package proto

type StreamType string

const (
	StreamCmd      StreamType = "cmd"
	StreamResp     StreamType = "resp"
	StreamNotify   StreamType = "notify"
	StreamFile     StreamType = "file"
	StreamDHT      StreamType = "dht"
	StreamSkywire  StreamType = "skywire"
	StreamYgg      StreamType = "ygg"
	StreamI2P      StreamType = "i2p"
	StreamTor      StreamType = "tor"
	StreamAmzWG    StreamType = "amneziawg"
	StreamAmzXRAY  StreamType = "amneziaxray"
)

type StreamHeader struct {
	StreamType StreamType `json:"stream_type"`
	StreamID   uint32     `json:"stream_id"`
	Flags      uint16     `json:"flags"`
	Length     uint32     `json:"length"`
}

type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      uint64      `json:"id"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	ID      uint64      `json:"id"`
}
