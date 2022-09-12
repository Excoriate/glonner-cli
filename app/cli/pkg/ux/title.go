package ux

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
)

var (
	typesAllowed = []string{"info", "success", "warning", "error"}
)

func checkTitleType(titleType string) error {
	for _, t := range typesAllowed {
		if t == titleType {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Invalid title type. Please use one of the following: %s", typesAllowed))
}

func ShowHeader(title string, titleType string) error {
	err := checkTitleType(titleType)

	if err != nil {
		return err
	}

	switch titleType {
	case "info":
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.
			BgBlue)).WithTextStyle(pterm.NewStyle(pterm.FgDarkGray)).WithMargin(0).Println(title)

	case "warning":
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.
			BgYellow)).WithTextStyle(pterm.NewStyle(pterm.FgDarkGray)).WithMargin(0).Println(title)

	case "error":
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.
			BgRed)).WithTextStyle(pterm.NewStyle(pterm.FgDarkGray)).WithMargin(0).Println(title)

	case "success":
		pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.
			BgGreen)).WithTextStyle(pterm.NewStyle(pterm.FgDarkGray)).WithMargin(0).Println(title)
	}

	return nil
}
