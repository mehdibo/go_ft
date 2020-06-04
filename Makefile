GOCMD=go
NAME=go_ft

all: $(NAME)

$(NAME): deps
	echo "Building binary..."
	$(GOCMD) build -o $(NAME) -v

deps:
	echo "Installing dependencies..."
	$(GOCMD) mod vendor

all_platforms: deps
	for GOOS in darwin linux; do \
        for GOARCH in 386 amd64; do \
			export GOOS $GOOS ; \
			export GOARCH $GOARCH; \
			echo "Building binary for $$GOOS - $$GOARCH" ; \
			go build -o $(NAME)-$$GOOS-$$GOARCH; \
		done \
    done