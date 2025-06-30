package procops

import (
	"fmt"
	"os/exec"
)

func ExecuteIndicatorCmd(cmd string) error {
	err := exec.Command("calc.exe").Run()
	if err != nil {
		return fmt.Errorf("unable to execute the specified command: %s - %w", cmd, err)
	}
	return err
}
