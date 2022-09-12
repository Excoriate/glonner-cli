package system

import (
	logger "github.com/glonner/pkg/log"
	"os/exec"
	"strings"
)

func RunCMD(logger logger.ILogger, cmd string, args ...string) (string, error) {
	cmdNormalised := strings.TrimSpace(cmd)

	out, err := exec.Command(cmdNormalised, args...).CombinedOutput()
	// Convert out to string
	var outStr string

	if out != nil {
		logger.LogDebug(outStr)
		outStr = string(out)
	}

	logger.LogDebug("Running command: " + cmdNormalised + " " + strings.Join(args, " "))

	if err != nil {
		logger.LogDebugF(err.Error())
		return "", err
	}

	return outStr, nil
}
