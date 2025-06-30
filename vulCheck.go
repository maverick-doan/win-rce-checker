package main

import (
	"fmt"
	"os"
	fileOps "win-rce-checker/pkg/file-ops"
	procOps "win-rce-checker/pkg/proc-ops"
)

var (
	path    string = "C:\\Users\\Public\\exploit_success.txt"
	content string = "Exploited"
	cmd            = "calc.exe"
)

func main() {
	// Ignore any commandline arguments
	_ = os.Args

	// Evidence file
	file, err := fileOps.WriteIndicatorFile(path, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	fmt.Println("Evidence file written.")

	// Evidence command execution
	err = procOps.ExecuteIndicatorCmd(cmd)
	if err != nil {
		fmt.Println("Calc failed (maybe SYSTEM?):", err)
	} else {
		fmt.Println("Calc launched.")
	}
}
