package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// says it cannot find list.bat in the path.
	runScript("subfolder", "list.bat")
	// correctly runs .\subfolder\list2.bat from .\subfolder
	runScript("subfolder", "list2.bat")
	// runs .\subfolder\list.bat from .\
	runScript("", "subfolder\\list.bat")
	// runs .\subfolder\list2.bat from .\
	runScript("", "subfolder\\list2.bat")
}

func runScript(folder, filename string) bool {
	fmt.Println("Running", folder, filename)
	cmd := exec.Command(filename)
	cmd.Dir = folder
	// Stream to std out
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		fmt.Println("Error running script: " + err.Error())
		return false
	}
	return true
}