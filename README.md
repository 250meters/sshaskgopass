# Go example projects

[![Go Reference](https://pkg.go.dev/badge/github.com/250meters/sshaskgopass.svg)](https://pkg.go.dev/github.com/250meters/sshaskgopass)

Sshaskgopass implements sshaskpass and obtains passwords, passphrases and otp tokens directoy from [gopass](https://github.com/gopasspw/gopass).

## Install
```shell
go install https://github.com/250meters/sshaskgopass/cmd/sshaskgopass@latest
```

## Setup

Add these exports to your shell's rc to redirect ssh secret requests to sshaskgopass.

```shell
export SSH_ASKPASS=$(go env GOPATH)/bin/sshaskgopass
export SSH_ASKPASS_REQUIRE=force
```

## Secret Naming Convention
sshaskgopass relyies on a secret naming convention when calling gopass.  
Read ahead to discover the secret. The sshaskgopass secrets are first prefixed with 
"ssh/<type>/" where "<type>" is one of otp, passphrase or password.  The rest of the secret 
matches the ssh prompt.

## Secret Nameing Examples
### OTP
Ssh prompt: Please provide (user@server.tld) otp code:
gopass secret name: ssh/otp/user@server.tld

### Passphrase
Ssh prompt: Enter passphrase for key '/home/user/.ssh/id_rsa':
gopass secret name: ssh/passphrase/home/user/.ssh/id_rsa

### Password
Ssh prompt: user@server.tld's password:
gopass secret name: ssh/password/user@server.tld

## Contributing
### Post Clone Setup
To make pull requests as smooth as possiable please include the provided ".gitconfig" file in your local project git settings.

```shell
git config --local include.path ../.gitconfig
```
