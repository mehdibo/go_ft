GOCMD=go
NAME=go_ft
INSTALL_PATH=/usr/local/bin
CONFIG_FILE_EXAMPLE=config.example.yaml
CONFIG_FILE_TARGET=$$HOME/.go_ft.yaml

all: $(NAME)

$(NAME): deps
	@echo "Building binary..."
	$(GOCMD) build -o $(NAME) -v

deps:
	@echo "Installing dependencies..."
	$(GOCMD) mod vendor

all_platforms: deps
	for GOOS in darwin linux; do \
        for GOARCH in 386 amd64; do \
			export GOOS $GOOS ; \
			export GOARCH $GOARCH; \
			@echo "Building binary for $$GOOS - $$GOARCH" ; \
			go build -o $(NAME)-$$GOOS-$$GOARCH; \
		done \
    done

install: $(NAME)
	sudo mv $(NAME) $(INSTALL_PATH)
	cp $(CONFIG_FILE_EXAMPLE) $(CONFIG_FILE_TARGET)
	@echo "You need to edit $(CONFIG_FILE_TARGET) with your credentials"

uninstall:
	sudo rm $(INSTALL_PATH)/$(NAME)
	rm $(CONFIG_FILE_TARGET)