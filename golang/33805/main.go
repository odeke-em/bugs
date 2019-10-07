package main

import (
	"regexp"
)

func main() {
	re := regexp.MustCompile("^" +
		"C:\\\\Users\\\\gopher\\\\AppData\\\\Local\\\\Temp\\\\1\\[0-9]+xyz$")

                dir := "C:\\Users\\gopher\\AppData\\Local\\Temp\\1"
                sep := "\\"
                want := "^" + regexp.QuoteMeta(dir) + sep + "[0-9]+xyz$"
	in := "" +
		"C:\\Users\\gopher\\AppData\\Local\\Temp\\1\\631670953xyz"
	println(re.MatchString(in))
	println(regexp.QuoteMeta(in))
        rex := regexp.MustCompile(want)
        println(rex.MatchString(in))
        println(want)

        // runtime, internal/poll: ensure that no thread is consumed, nor a syscall.Read if FD isn't yet ready for I/O
        // runtime, internal/poll: ensure that no thread is consumed, nor a syscall.Read if FD isn't yet ready for I/O
        // runtime, internal/poll: darwin: ensure that no thread is consumed, nor a syscall.Read if FD isn't yet ready for I/O
}
