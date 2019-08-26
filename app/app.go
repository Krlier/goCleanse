package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Krlier/goCleanse/git"
)

func main() {
	// Loads needed environment variables
	repURL := os.Getenv("GCLEANSE_REPURL")
	clonePath := os.Getenv("GCLEANSE_CLONEPATH")
	newBranchName := os.Getenv("GCLEANSE_NEWBRANCHNAME")
	password := os.Getenv("GCLEANSE_OLDSTRING")
	newString := os.Getenv("GCLEANSE_NEWSTRING")
	fileExtension := os.Getenv("GCLEANSE_FILEEXTENSION")
	remote := os.Getenv("GCLEANSE_REMOTE")

	// Verifies if an e-mail has been provided as argument
	if len(os.Args) > 1 {
		if err := git.ChangeDir(clonePath); err != nil {
			fmt.Println(err)
		}

		if err := git.SignCommits(os.Args[1]); err != nil {
			panic(err)
		}

		if err := git.Push(newBranchName, remote); err != nil {
			panic(err)
		}

		return
	}

	// Clones the repository
	if err := git.Clone(clonePath, repURL); err != nil {
		log.Fatal(err)
	}

	// Changes directory to the just cloned repository
	if err := git.ChangeDir(clonePath); err != nil {
		log.Fatal(err)
	}

	// Creates a new branch in which the changes will be pushed to
	if err := git.CreateBranch(newBranchName); err != nil {
		log.Fatal(err)
	}

	// Searches and removes the sensitive content from the repository
	if err := git.RemoveSensitiveCode(password, newString, fileExtension, clonePath); err != nil {
		log.Fatal(err)
	}

	// Push the repository
	if err := git.Push(newBranchName, remote); err != nil {
		log.Fatal(err)
	}
}
