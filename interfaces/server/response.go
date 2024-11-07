package server

import (
	"github.com/ruohuaii/goddd/infrastructure/faults"
)

type Response struct {
	Errors []*faults.Faults `json:"errors"`
	Data   any              `json:"data"`
}
