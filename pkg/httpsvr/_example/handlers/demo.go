package handlers

import (
	"github.com/dup2X/gomicrosvr/pkg/errcode"
	"github.com/dup2X/gomicrosvr/pkg/httpsvr/_example/common"
	"github.com/dup2X/gomicrosvr/pkg/httpsvr/_example/idl"
)

// DemoProc ...
func DemoProc(req *idl.RequestSt) (*idl.ResponseSt, *common.DemoErr) {
	return &idl.ResponseSt{1024, 1}, common.GenErr(errcode.Success, nil)
}
