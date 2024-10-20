**sshaskgopass [![Go Reference](https://pkg.go.dev/badge/github.com/250meters/sshaskgopass.svg)](https://pkg.go.dev/github.com/250meters/sshaskgopass)**
================

`sshaskgopass` implements `ssh-askpass` and obtains passwords, passphrases, and OTP tokens directly from [gopass](https://github.com/gopasspw/gopass).

## Install
----------

To install `sshaskgopass`, run:
```shell
go install github.com/250meters/sshaskgopass@latest
```

## Setup
--------

Add the following exports to your shell's rc file to redirect SSH secret requests to `sshaskgopass`:

```shell
export GOPASS=$(go env GOPATH)/bin/gopass
export SSH_ASKPASS=$(go env GOPATH)/bin/sshaskgopass
export SSH_ASKPASS_REQUIRE=force
```

## Secret Naming Convention
-------------------------

When calling `gopass`, `sshaskgopass` relies on a secret naming convention. The format is:

```shell
ssh/<type>/[rest of the secret]
```

Where `<type>` is one of `otp`, `passphrase`, or `password`. For example:

### OTP

* SSH prompt: `Please provide (user@server.tld) otp code: `
* Gopass secret name: `ssh/otp/user@server.tld`
* Add or edit the secret: `gopass edit -c ssh/otp/user@server.tld`

### Passphrase

* SSH prompt: `Enter passphrase for key '/home/user/.ssh/id_rsa': `
* Gopass secret name: `ssh/passphrase/home/user/.ssh/id_rsa`
* Add or edit the secret: `gopass edit -c ssh/passphrase/home/user/.ssh/id_rsa`

### Password

* SSH prompt: `user@server.tld's password: `
* Gopass secret name: `ssh/password/user@server.tld`
* Add or edit the secret: `gopass edit -c ssh/password/user@server.tld`

## Know Issue
`sshaskgopass` is stateless.  When `gopass` has the requested secrect but the secret is incorrect ssh will reinvoke `sshaskgopass` for the same secret.  This will happen about three times and then ssh will prompt for a less secure secret if available.  That will present itself to you the user as being initially prompted for an unexpected less secure secret.  For instance a user's password instead of a key's passphrase.  When that happens double check your `gopass` secrets for correctness.  Also if you can disable ssh password based authentication you should.

## Contributing
-------------

To make pull requests as smooth as possible, include the provided `.gitconfig` file in your local project's Git settings:

```shell
git config --local include.path ../.gitconfig
```
