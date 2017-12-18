package service

import (
	"encoding/json"
	"fmt"
)

// Code used to describe the error code in the rpc service.
type Code uint32

const (
	// Separator separation of the RPCError field information.
	Separator string = "\n\n"
)

// RPCError defines the status from an RPC.
type RPCError struct {
	Code  Code
	Desc  string
	Extra map[string]interface{}
}

// Error returns the error information.
func (e RPCError) Error() string {
	errMsg := fmt.Sprintf("%d%s%s", e.Code, Separator, e.Desc)
	if e.Extra != nil {
		buf, err := json.Marshal(e.Extra)
		if err == nil {
			errMsg = fmt.Sprintf("%s%s%s", errMsg, Separator, string(buf))
		}
	}
	return errMsg
}

// NewRPCError returns system pre-defined error
func NewRPCError(code Code, desc string) *RPCError {
	return &RPCError{
		Code: code,
		Desc: desc,
	}
}
