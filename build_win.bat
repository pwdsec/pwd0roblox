go env -w GOOS=windows GOARCH=386
go build -o bin/win/x86/pwd0roblox_x86.exe
go env -w GOOS=windows GOARCH=amd64
go build -o bin/win/x64/pwd0roblox_x64.exe