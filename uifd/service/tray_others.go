//go:build !windows
// +build !windows

package main

import (
	"uifd/uif"
)

func TrayInit() {
	uif.WaitQuitSnignal()
}

func SetQuicLink() error {
	return nil
}
