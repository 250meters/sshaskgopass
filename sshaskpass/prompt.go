package sshaskpass

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type Prompt string

func (p Prompt) IsOTP() bool {
	return strings.Contains(string(p), "code:")
}

func (p Prompt) RequestedOtp() string {
	return p.between("(", ")")
}

// returns true if the promt matches "Enter passphrase for key '/path/to/key/file': "
func (p Prompt) IsPassphraseForKey() bool {
	return strings.Contains(string(p), "passphrase for key")
}

func (p Prompt) RequestedPassphraseForKey() string {
	return p.between("'", "'")
}

func (p Prompt) IsPassword() bool {
	return strings.Contains(string(p), "assword:")
}

func (p Prompt) RequestedPassword() string {
	return p.before("'")
}

func (p Prompt) AskPassword() (password string, err error) {
	defer p.clearPrompt()
	fmt.Fprint(os.Stderr, p)
	bp, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("could not read password: %w", err)
	}
	return string(bp), nil
}

func (p Prompt) AskOther() (resp string, err error) {
	defer p.clearPrompt()
	fmt.Fprint(os.Stderr, p)
	resp, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read response: %w", err)
	}

	// Remove trailing newlines from the input
	resp = resp[:len(resp)-1]

	return resp, nil
}

func (p Prompt) clearPrompt() {
	s := string(p)

	fmt.Fprint(os.Stderr, "\033[G")  // cursor at begining of line
	fmt.Fprint(os.Stderr, "\033[2K") // clear the line

	firstLine := true
	for _, c := range s {
		if c == '\n' {
			if firstLine {
				firstLine = false
			} else {
				fmt.Fprint(os.Stderr, "\033[1A") // Move cursor up one line
			}

			fmt.Fprint(os.Stderr, "\033[2K") // clear the line
		}
	}
}

func (p Prompt) between(first, last string) string {
	ps := string(p)
	i := strings.Index(ps, first)
	if i >= 0 {
		j := strings.Index(ps[i+1:], last)
		if j >= 0 {
			return ps[i+1 : j+i+1]
		}
	}
	return ""
}

func (p Prompt) before(end string) string {
	ps := string(p)
	i := strings.Index(ps, end)
	if i >= 0 {
		return ps[:i]
	}
	return ""
}
