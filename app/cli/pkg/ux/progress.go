package ux

import (
	"fmt"
	"github.com/pterm/pterm"
	"time"
)

type ProgressBar struct {
	Title                   string
	SleepTime               time.Duration
	ProgressMessage         string
	StdPrefixSuccessMessage string
	StdPrefixErrorMessage   string
	StdPrefixWarningMessage string
}

type IProgressBar interface {
	OnSuccess(items string, customMsg string)
	OnWarning(item string, customMsg string)
	OnFail(item string, customMsg string)
	Start(items []string) *pterm.ProgressbarPrinter
	ForceThreeDotsOnProgressMessage(msg string) string
}

func NewProgressBar(title string, sleepTime int, progressMsg string, stdPrefixSuccessMsg string,
	stdPrefixErrorMsg string, stdPrefixWarningMsg string) *ProgressBar {
	timeIn := time.Duration(sleepTime) * time.Second
	return &ProgressBar{
		Title:                   title,
		SleepTime:               timeIn,
		ProgressMessage:         progressMsg,
		StdPrefixSuccessMessage: stdPrefixSuccessMsg,
		StdPrefixErrorMessage:   stdPrefixErrorMsg,
		StdPrefixWarningMessage: stdPrefixWarningMsg,
	}
}

func (pb *ProgressBar) OnSuccess(item, customMsg string) {
	if customMsg != "" {
		pterm.Success.Println(customMsg)
	} else {
		if item != "" {
			pterm.Success.Println(fmt.Sprintf("%s %s", pb.StdPrefixSuccessMessage, item))
		} else {
			pterm.Success.Println(pb.StdPrefixSuccessMessage)
		}
	}
}

func (pb *ProgressBar) OnWarning(item, customMsg string) {
	if customMsg != "" {
		pterm.Warning.Println(customMsg)
	} else {
		if item != "" {
			pterm.Warning.Println(fmt.Sprintf("%s %s", pb.StdPrefixWarningMessage, item))
		} else {
			pterm.Warning.Println(pb.StdPrefixWarningMessage)
		}
	}
}

func (pb *ProgressBar) OnFail(item, customMsg string) {
	if customMsg != "" {
		pterm.Error.Println(customMsg)
	} else {
		if item != "" {
			pterm.Error.Println(fmt.Sprintf("%s %s", pb.StdPrefixErrorMessage, item))
		} else {
			pterm.Error.Println(pb.StdPrefixErrorMessage)
		}
	}
}

func (pb *ProgressBar) Start(items []string) *pterm.ProgressbarPrinter {
	p, _ := pterm.DefaultProgressbar.WithTotal(len(items)).WithTitle(pb.Title).Start()

	return p
}

func (pb *ProgressBar) ForceThreeDotsOnProgressMessage(msg string) string {
	if msg[len(msg)-3:] != "..." {
		return fmt.Sprintf("%s...", msg)
	} else {
		return msg
	}
}
