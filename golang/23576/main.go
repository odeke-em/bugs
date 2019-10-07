package main

import (
	"os"

	git "gopkg.in/libgit2/git2go.v26"
)

func main() {
	pwd, err := os.Getwd()
	check(err)

	repo, err := git.OpenRepositoryExtended(pwd, 0, "")
	check(err)

	rw, err := repo.Walk()
	check(err)

	rw.Push(nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
