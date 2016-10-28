package main

func main() {
	var a int
	switch a {
	case 0, 3:
	case -1, 1, 0, 2, 3:
	case -1:
	}

	switch "go" {
	case "go", "gopher":
	case "go":
	}

	var s interface{} = "hello"
	switch s.(type) {
	case string:
	case int, string:
	}

	var goarch = "darwin"
	var goarm = ""
	switch {
	case goarch == "arm" && goarm == "5":
	case goarch == "arm":
	case "come" == "as you are":
	}
}
