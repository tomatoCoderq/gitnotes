package commands

import (
	"os/exec"
)



func ResolveGitRef (ref string) (string, error) {
	out, err := exec.Command("git", "rev-parse", ref).Output()
	if err != nil {
		return "", err
	}

	shortenSHA := out[:6]
	return string(shortenSHA), nil
}