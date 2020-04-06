install_linux:
	go build cmd/baila/baila.go && sudo mv baila /usr/local/bin
test:
	go test ./api