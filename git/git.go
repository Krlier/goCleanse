package git

import (
	"fmt"
	"os"
	"os/exec"
)

// ChangeDir changes the working directory to the one provided
func ChangeDir(path string) error {
	if err := os.Chdir(path); err != nil {
		return err
	}

	return nil
}

// Clone clones a git repository
func Clone(destination string, repoURL string) error {
	cmd := exec.Command("git", "clone", repoURL, destination)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// CreateBranch creates a new branch to the git repository in the current directory
func CreateBranch(branchName string) error {
	cmd := exec.Command("git", "checkout", "-b", branchName)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// RemoveSensitiveCode removes a sensitive string from a git repository and it's history
func RemoveSensitiveCode(sensitive string, substitute string, fileExtension string, repoPath string) error {
	gitCommand := fmt.Sprintf("git filter-branch -f --tree-filter 'git ls-files -z \"*.%s\" |xargs -0 perl -p -i -e \"s#(%s)#%s#g\"' -- --all", fileExtension, sensitive, substitute)

	cmd := exec.Command("bash", "-c", gitCommand)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// SignCommits signs all commits from a given user from his commiter e-mail
func SignCommits(commiterEmail string) error {
	gitCommand := fmt.Sprintf("git filter-branch -f --commit-filter 'if [ \"$GIT_COMMITTER_EMAIL\" = \"%s\" ]; then git commit-tree -S \"$@\"; else git commit-tree \"$@\"; fi' HEAD", commiterEmail)
	fmt.Println(gitCommand)
	cmd := exec.Command("bash", "-c", gitCommand)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Push pushes the current repository
func Push(branchName string, remote string) error {
	cmd := exec.Command("git", "push", remote, branchName, "--force")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
