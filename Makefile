goBuildStatic:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a ./...

goInstallStatic:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -ldflags='-w -s -extldflags "-static"' -a ./...

goBuildWindows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/oa2server.exe

goBuildOsx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/oa2server-osx

precommit: securitycheck lint testci

securitycheck: govulncheck gosec

test:
	go test ./...

testRace:
	go test --race ./...

testci:
	go test -vet=all -failfast -short ./...

updateDepsInteractive:
ifeq ($(shell which go-mod-upgrade 2>/dev/null),)
	go install github.com/oligot/go-mod-upgrade@latest
endif
	$(MAKE) updateGo; go-mod-upgrade; go mod tidy

updateDeps:
	$(MAKE) updateGo; go get -u ./...; go mod tidy


goInstallLint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
ifeq ($(shell which golangci-lint 2>/dev/null),)
	$(MAKE) goInstallLint
endif
	golangci-lint run


goInstallPkgsite:
	go install golang.org/x/pkgsite/cmd/pkgsite@latest

doc:
ifeq ($(shell which pkgsite 2>/dev/null),)
	$(MAKE) goInstallPkgsite
endif
	pkgsite

goInstallVulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest

govulncheck:
ifeq ($(shell which govulncheck 2>/dev/null),)
	$(MAKE) goInstallVulncheck
endif
	govulncheck ./...


goInstallGosec:
	go install github.com/securego/gosec/v2/cmd/gosec@latest

gosec:
ifeq ($(shell which gosec 2>/dev/null),)
	$(MAKE) goInstallGosec
endif
	gosec --quiet ./...


goInstallGomajor:
	go install github.com/icholy/gomajor@latest

depsListMajor:
ifeq ($(shell which gomajor 2>/dev/null),)
	$(MAKE) goInstallGomajor
endif
	gomajor list

updateGo:
	go env -w "GOTOOLCHAIN=$(shell curl -s https://go.dev/VERSION\?m=text | head -n 1)+auto"

GOVER := $(shell curl -s https://go.dev/VERSION?m=text | head -n 1)
goLinuxReinstall:
	sudo rm -rf /usr/local/go && curl "https://dl.google.com/go/${GOVER}.linux-amd64.tar.gz" | sudo tar -C /usr/local -xz

goInstallAll: goInstallPkgsite goInstallVulncheck goInstallGosec goInstallGomajor
