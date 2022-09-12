package ux

import (
	"github.com/pterm/pterm"
	"time"
)

func GetSpinner(msg string, seconds int64) *pterm.SpinnerPrinter {
	s, _ := pterm.DefaultSpinner.Start(msg)
	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)

	return s
}
