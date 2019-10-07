package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	tmpDir, err := ioutil.TempDir("", "30336")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	deeplyNestedDir := filepath.Join(tmpDir, "a", "b", "c", "d", "issueX")
	if err := os.MkdirAll(deeplyNestedDir, 0755); err != nil {
		log.Fatalf("MkdirAll failed: %v", err)
	}

	if false {
		mainContent := `
package main
func main() {}`

		mainPath := filepath.Join(deeplyNestedDir, "main.go")
		if err := ioutil.WriteFile(mainPath, []byte(mainContent), 0755); err != nil {
			log.Fatalf("Failed to write file contents to main.go: %v", err)
		}
	}

	symlinkPath := filepath.Join(tmpDir, "issueX")
	if err := os.Symlink(deeplyNestedDir, symlinkPath); err != nil {
		log.Fatalf("Symlink failed: %v", err)
	}

	originalCwd, _ := os.Getwd()
        fmt.Printf("DeeplyNestedDir: %q\nSymlinkPath:     %q\n\n", deeplyNestedDir, symlinkPath)
	if err := os.Chdir(tmpDir); err != nil {
		log.Fatalf("Failed to change to deeplyNestedDir: %v", err)
	}
	defer os.Chdir(originalCwd)
        _ = os.Chdir("issueX")

	paths := []string{
		filepath.Join(tmpDir, "issueX", "main.go"),
	}

	rps := RelPaths(paths)

	fmt.Printf("RelPaths:\n%s\n", rps)
}

func RelPaths(paths []string) []string {
	var out []string
	// TODO(rsc): Can this use Cwd from above?
	pwd, _ := os.Getwd()
	for _, p := range paths {
		rel, err := filepath.Rel(pwd, p)
                fmt.Printf("Rel: %s err: %v\np:   %q\npwd: %q\n", rel, err, p, pwd)
		if err == nil && len(rel) < len(p) {
			p = rel
		}
		out = append(out, p)
	}
	return out
}
