
.PHONY: all
all:
	go build -o bin/fynebind ./cmd/fynebind

.PHONY: macos-app
macos-app: all
	fyne package --executable bin/fynebind --name "FyneBinds" --appVersion 0.1.0 --icon assets/macos_icon.png
	mv "FyneBinds.app" bin/

.PHONY: clean
clean:
	rm -r bin/

.PHONY: demo
demo:
	go run fyne.io/fyne/v2/cmd/fyne_demo@latest
