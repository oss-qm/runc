//go:build linux && no_systemd

package main

import (
	"errors"
	"os"
)

func sdGetListenFDs() []*os.File {
	return nil
}

func sdDetectUID() (int, error) {
	return -1, errors.New("no systemd support")
}
