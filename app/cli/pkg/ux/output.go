package ux

import (
	"github.com/pterm/pterm"
	"os"
)

func OutInfo(message string, header string) {
	if header == "" {
		pterm.Info.Println(message)
	} else {
		pterm.Info.Prefix = pterm.Prefix{
			Text: header,
		}
		pterm.Info.Println(message)
	}
}

func OutError(message string, header string, killOnError bool) {
	if header == "" {
		pterm.Error.Println(message)
	} else {
		pterm.Error.Prefix = pterm.Prefix{
			Text: header,
		}
		pterm.Error.Println(message)
	}

	if killOnError {
		os.Exit(1)
	}
}

func OutSuccess(message string, header string) {
	if header == "" {
		pterm.Success.Println(message)
	} else {
		pterm.Success.Prefix = pterm.Prefix{
			Text: header,
		}
		pterm.Success.Println(message)
	}
}

func OutFatal(message string, header string, killOnError bool) {
	if header == "" {
		pterm.Fatal.Println(message)
	} else {
		pterm.Fatal.Prefix = pterm.Prefix{
			Text: header,
		}
		pterm.Fatal.Println(message)
	}

	if killOnError {
		os.Exit(1)
	}
}

func OutWarn(message string, header string) {
	if header == "" {
		pterm.Warning.Println(message)
	} else {
		pterm.Warning.Prefix = pterm.Prefix{
			Text: header,
		}
		pterm.Warning.Println(message)
	}
}

func ShowAppTitle(appName string, topMsg string, bottomMsg string) {
	pterm.DefaultCenter.Println(topMsg)

	s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(appName)).Srender()
	pterm.DefaultCenter.Println(s)

	pterm.DefaultCenter.WithCenterEachLineSeparately().Println(bottomMsg)
}
