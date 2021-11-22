build:
	@ CGO_ENABLED=0 CC=i686-w64-mingw32-gcc CXX=i686-w64-mingw32-g++ GOOS=windows GOARCH=386 go build -ldflags="-H windowsgui"