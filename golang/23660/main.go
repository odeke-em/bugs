package main

import git "gopkg.in/libgit2/git2go.v26"

func f(r *git.Repository, x int) {
	r.CreateCommit(
		"",
		nil,
		nil,
		"",
		nil,
		x,
	)
}

func main() {
}
