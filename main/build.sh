#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/linux-amd64 main.go
#
#CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o build/linux-x86 main.go
#
#CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/linux-arm64 main.go

#CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o build/mac-arm64 main.go

#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o build/mac-amd64 main.go

#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/win-amd64.exe main.go
#
#CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o build/win-x86.exe main.go
#
#CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o build/win-arm64.exe main.go


CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o /Users/lzp/soft/vino/ss main.go