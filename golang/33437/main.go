package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var code = fmt.Sprintf(`
package main

type function struct {
	name   string
	params []string
}

var functions map[uint32]function

func main() {
     functions = make(map[uint32]function)
     %s
}`, makeFunctions(140000))

func makeFunctions(n int) string {
	sb := new(strings.Builder)
	for i := 0; i < n; i++ {
		name := fmt.Sprintf("fn-%d", i)
		fmt.Fprintf(sb, "\tfunctions[%d] = function{name: %q, params: []string{%q}}\n", i, name, name)
	}
	return sb.String()
}

func main() {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "issue-33437")
	if err != nil {
		log.Fatalf("Failed to create tempdir: %v", err)
	}
	defer os.Remove(tmpDir)

	generatedMainGoFile := filepath.Join(tmpDir, "main.go")
	if err := ioutil.WriteFile(generatedMainGoFile, []byte(code), 0655); err != nil {
		log.Printf("Failed to write generated main.go file: %v", err)
		return
	}

	ctx := context.Background() // Perhaps set it to a timer.
	cmd := exec.CommandContext(ctx, "go", "run", generatedMainGoFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to run the command: %v\n%s", err, output)
	}
}
