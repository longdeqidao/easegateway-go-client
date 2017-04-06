package pdu

import "time"

type UpTime struct {
	UpTime time.Duration `json:"up_nanosec"`
}
