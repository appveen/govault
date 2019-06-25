#!/bin/bash

cd $WORKSPACE

set -e
if [ -f $WORKSPACE/../TOGGLE ]; then
	echo "****************************************************"
    echo "odp:govault :: Toggle mode is on, terminating build"
    echo "odp:govault :: BUILD CANCLED"
    echo "****************************************************"
    exit 0
fi

echo "****************************************************"
echo "govault :: Building excutables"
echo "****************************************************"

# echo "env GOOS=android GOARCH=arm go build -o exec/vault-android-arm cmd/main.go"
# env GOOS=android GOARCH=arm go build -o exec/vault-android-arm cmd/main.go || true

# echo "env GOOS=darwin GOARCH=386 go build -o exec/vault-darwin-386 cmd/main.go"
# env GOOS=darwin GOARCH=386 go build -o exec/vault-darwin-386 cmd/main.go

echo "env GOOS=darwin GOARCH=amd64 go build -o exec/vault-darwin-amd64 cmd/main.go"
env GOOS=darwin GOARCH=amd64 go build -o exec/vault-darwin-amd64 cmd/main.go || true

# echo "env GOOS=darwin GOARCH=arm go build -o exec/vault-darwin-arm cmd/main.go"
# env GOOS=darwin GOARCH=arm go build -o exec/vault-darwin-arm cmd/main.go || true

# echo "env GOOS=darwin GOARCH=arm64 go build -o exec/vault-darwin-arm64 cmd/main.go"
# env GOOS=darwin GOARCH=arm64 go build -o exec/vault-darwin-arm64 cmd/main.go || true

# echo "env GOOS=dragonfly GOARCH=amd64 go build -o exec/vault-dragonfly-amd64 cmd/main.go"
# env GOOS=dragonfly GOARCH=amd64 go build -o exec/vault-dragonfly-amd64 cmd/main.go || true

# echo "env GOOS=freebsd GOARCH=386 go build -o exec/vault-freebsd-386 cmd/main.go"
# env GOOS=freebsd GOARCH=386 go build -o exec/vault-freebsd-386 cmd/main.go || true

# echo "env GOOS=freebsd GOARCH=amd64 go build -o exec/vault-freebsd-amd64 cmd/main.go"
# env GOOS=freebsd GOARCH=amd64 go build -o exec/vault-freebsd-amd64 cmd/main.go || true

# echo "env GOOS=freebsd GOARCH=arm go build -o exec/vault-freebsd-arm cmd/main.go"
# env GOOS=freebsd GOARCH=arm go build -o exec/vault-freebsd-arm cmd/main.go || true

echo "env GOOS=linux GOARCH=386 go build -o exec/vault-linux-386 cmd/main.go"
env GOOS=linux GOARCH=386 go build -o exec/vault-linux-386 cmd/main.go

echo "env GOOS=linux GOARCH=amd64 go build -o exec/vault-linux-amd64 cmd/main.go"
env GOOS=linux GOARCH=amd64 go build -o exec/vault-linux-amd64 cmd/main.go || true

# echo "env GOOS=linux GOARCH=arm go build -o exec/vault-linux-arm cmd/main.go"
# env GOOS=linux GOARCH=arm go build -o exec/vault-linux-arm cmd/main.go || true

# echo "env GOOS=linux GOARCH=arm64 go build -o exec/vault-linux-arm64 cmd/main.go"
# env GOOS=linux GOARCH=arm64 go build -o exec/vault-linux-arm64 cmd/main.go || true

# echo "env GOOS=linux GOARCH=ppc64 go build -o exec/vault-linux-ppc64 cmd/main.go"
# env GOOS=linux GOARCH=ppc64 go build -o exec/vault-linux-ppc64 cmd/main.go || true

# echo "env GOOS=linux GOARCH=ppc64le go build -o exec/vault-linux-ppc64le cmd/main.go"
# env GOOS=linux GOARCH=ppc64le go build -o exec/vault-linux-ppc64le cmd/main.go || true

# echo "env GOOS=linux GOARCH=mips go build -o exec/vault-linux-mips cmd/main.go"
# env GOOS=linux GOARCH=mips go build -o exec/vault-linux-mips cmd/main.go || true

# echo "env GOOS=linux GOARCH=mipsle go build -o exec/vault-linux-mipsle cmd/main.go"
# env GOOS=linux GOARCH=mipsle go build -o exec/vault-linux-mipsle cmd/main.go || true

# echo "env GOOS=linux GOARCH=mips64 go build -o exec/vault-linux-mips64 cmd/main.go"
# env GOOS=linux GOARCH=mips64 go build -o exec/vault-linux-mips64 cmd/main.go || true

# echo "env GOOS=linux GOARCH=mips64le go build -o exec/vault-linux-mips64le cmd/main.go"
# env GOOS=linux GOARCH=mips64le go build -o exec/vault-linux-mips64le cmd/main.go || true

# echo "env GOOS=netbsd GOARCH=386 go build -o exec/vault-netbsd-386 cmd/main.go"
# env GOOS=netbsd GOARCH=386 go build -o exec/vault-netbsd-386 cmd/main.go || true

# echo "env GOOS=netbsd GOARCH=amd64 go build -o exec/vault-netbsd-amd64 cmd/main.go"
# env GOOS=netbsd GOARCH=amd64 go build -o exec/vault-netbsd-amd64 cmd/main.go || true

# echo "env GOOS=netbsd GOARCH=arm go build -o exec/vault-netbsd-arm cmd/main.go"
# env GOOS=netbsd GOARCH=arm go build -o exec/vault-netbsd-arm cmd/main.go || true

# echo "env GOOS=openbsd GOARCH=386 go build -o exec/vault-openbsd-386 cmd/main.go"
# env GOOS=openbsd GOARCH=386 go build -o exec/vault-openbsd-386 cmd/main.go || true

# echo "env GOOS=openbsd GOARCH=amd64 go build -o exec/vault-openbsd-amd64 cmd/main.go"
# env GOOS=openbsd GOARCH=amd64 go build -o exec/vault-openbsd-amd64 cmd/main.go || true

# echo "env GOOS=openbsd GOARCH=arm go build -o exec/vault-openbsd-arm cmd/main.go"
# env GOOS=openbsd GOARCH=arm go build -o exec/vault-openbsd-arm cmd/main.go || true

# echo "env GOOS=plan9 GOARCH=386 go build -o exec/vault-plan9-386 cmd/main.go"
# env GOOS=plan9 GOARCH=386 go build -o exec/vault-plan9-386 cmd/main.go || true

# echo "env GOOS=plan9 GOARCH=amd64 go build -o exec/vault-plan9-amd64 cmd/main.go"
# env GOOS=plan9 GOARCH=amd64 go build -o exec/vault-plan9-amd64 cmd/main.go || true

# echo "env GOOS=solaris GOARCH=amd64 go build -o exec/vault-solaris-amd64 cmd/main.go"
# env GOOS=solaris GOARCH=amd64 go build -o exec/vault-solaris-amd64 cmd/main.go || true

echo "env GOOS=windows GOARCH=386 go build -o exec/vault-windows-386.exe cmd/main.go"
env GOOS=windows GOARCH=386 go build -o exec/vault-windows-386.exe cmd/main.go

echo "env GOOS=windows GOARCH=amd64 go build -o exec/vault-windows-amd64.exe cmd/main.go"
env GOOS=windows GOARCH=amd64 go build -o exec/vault-windows-amd64.exe cmd/main.go

