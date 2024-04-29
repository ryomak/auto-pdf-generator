build-cli:
	go build -o ./bin/pdfgenerator-cli cmd/cli/main.go
build-desktop:
	cd cmd/desktop && wails build -clean
	cp ./cmd/desktop/build/bin/pdfgenerator-desktop.app/Contents/MacOS/pdfgenerator-desktop ./bin/
dev-desktop:
	cd cmd/desktop && wails dev
install:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest