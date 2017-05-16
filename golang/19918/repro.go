package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

func func3(id uint64) (bool, error) {
	foundSomething := false
	cmd := exec.Command("echo", "foobar")
	cmd.Stderr = ioutil.Discard
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return foundSomething, err
	}
	if err := cmd.Start(); err != nil { // main.go:186
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		return foundSomething, err
	}
	go func() {
		time.Sleep(5 * time.Second)
		err := cmd.Process.Kill()
		log.Printf("killing it: %d, err: %v", id, err)
	}()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Index(line, ":")
		if i < 0 {
			continue
		}
		if strings.Contains(line, "foo") {
			foundSomething = true
		}
	}
	err = cmd.Wait()
	return foundSomething, err
}

func main() {
	i := uint64(0)
	ticker := time.NewTicker(5 * time.Minute)
	for {
		i += 1
		go func(id uint64) {
			fs, err := func3(i)
			log.Printf("#%d: fs: %v err: %v\n", id, fs, err)
		}(i)

		<-ticker.C
	}
}
