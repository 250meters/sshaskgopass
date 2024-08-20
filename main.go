package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/250meters/sshaskgopass/gopass"
	"github.com/250meters/sshaskgopass/sshaskpass"
)

func main() {
	msg := sshaskpass.Prompt(strings.Join(os.Args[1:], " "))
	ctx := context.Background()

	if msg.IsOTP() {
		code, err := gopass.New("").Otp(ctx, msg.RequestedOtp())
		if err != nil {
			log.Println("Error: ", err)
			code, err = msg.AskOther()
			if err != nil {
				log.Println("Error: ", err)
				os.Exit(-1)
				return
			}
		}

		fmt.Print(code)
		return
	}

	if msg.IsPassword() {
		pass, err := gopass.New("").Password(ctx, msg.RequestedPassword())
		if err != nil {
			log.Println("Error: ", err)
			pass, err = msg.AskPassword()
			if err != nil {
				log.Println("Error: ", err)
				os.Exit(-1)
				return
			}
		}

		fmt.Print(pass)
		return
	}

	if msg.IsPassphraseForKey() {
		pass, err := gopass.New("").Passphrase(ctx, msg.RequestedPassphraseForKey())
		if err != nil {
			log.Println("Error: ", err)
			pass, err = msg.AskPassword()
			if err != nil {
				log.Println("Error: ", err)
				os.Exit(-1)
				return
			}
		}

		fmt.Print(pass)
		return
	}

	resp, err := msg.AskOther()
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(-1)
		return
	}

	fmt.Println(resp)
}
