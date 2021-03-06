// +build !windows

package main

import (
	"log"
	"log/syslog"
	"runtime"
)

func setupLogger() {
	if isDeamon && runtime.GOOS != "windows" {
		if logwriter, err := syslog.New(syslog.LOG_NOTICE, name); err == nil {
			log.SetOutput(logwriter)
		}
	}
}
