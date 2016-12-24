package issue_test

import (
	"fmt"
	"os"
	"os/exec"
)

func ExampleTest() {
	fmt.Println("foo")

	cmd := exec.Command("/bin/sh", "-c", "sleep 999d")
	cmd.Stdout = os.Stdout
	go cmd.Run()

	// Output:
	// foo
}
