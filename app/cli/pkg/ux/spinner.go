package ux

import (
	"github.com/briandowns/spinner"
	"github.com/pterm/pterm"
	"os"
	"time"
)

func GetLoader() spinner.Spinner {
	return *spinner.New([]string{
		"⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎",
		"⡜ ⡜", "⡜ ⡜", "⢎ ⡱", "⢎ ⡱", "⢣ ⢣", "⢣ ⢣",
		"⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎", "⡱ ⢎",
		"⣂⠔⠩", "⢄⡪⠑", "⠰⣉⠆", "⠊⢕⡠", "⠍⠢⣐",
	}, 30*time.Millisecond, spinner.WithWriter(os.Stderr))
}

func GetLoaderStandard(title string) *spinner.Spinner {
	mySpinner := GetLoader()
	mySpinner.Prefix = title
	return &mySpinner
}

func GetSpinner(msg string, seconds int64) *pterm.SpinnerPrinter {
	s, _ := pterm.DefaultSpinner.Start(msg)
	var duration time.Duration
	duration = time.Duration(seconds) * time.Second
	time.Sleep(duration)

	return s
}
