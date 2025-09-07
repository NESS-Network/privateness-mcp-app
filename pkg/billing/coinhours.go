package billing

import "time"

type Rates struct {
	PerByteIn  float64 // coin hours per byte in
	PerByteOut float64 // coin hours per byte out
	PerSecond  float64 // coin hours per second
}

// Compute coin hours cost given usage and rates.
func Cost(bytesIn, bytesOut uint64, dur time.Duration, r Rates) float64 {
	return float64(bytesIn)*r.PerByteIn + float64(bytesOut)*r.PerByteOut + dur.Seconds()*r.PerSecond
}

// TODO: Implement real debit via Privateness wallet/MCP
func Charge(pubKey string, amount float64) error {
	// Placeholder: no-op
	_ = pubKey; _ = amount
	return nil
}
