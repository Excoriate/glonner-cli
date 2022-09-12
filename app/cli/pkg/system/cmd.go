package system

import (
	logger "github.com/glonner/pkg/log"
	"os/exec"
	"strings"
)

func RunCMD(log logger.ILogger, cmd string, args ...string) (string, error) {
	cmdNormalised := strings.TrimSpace(cmd)

	out, err := exec.Command(cmdNormalised, args...).CombinedOutput()
	// Convert out to string
	var outStr string

	if out != nil {
		log.LogDebug(outStr)
		outStr = string(out)
	}

	log.LogDebug("Running command: " + cmdNormalised + " " + strings.Join(args, " "))

	if err != nil {
		log.LogDebugF(err.Error())
		return "", err
	}

	return outStr, nil
}
