GOCMD=go
NAME=go_ft

all: $(NAME)

$(NAME): deps
	$(GOCMD) build -o $(NAME) -v

deps:
	$(GOCMD) mod vendor

all_platforms: deps
	for GOOS in darwin linux; do \
        for GOARCH in 386 amd64; do \
			export GOOS $GOOS ; \
			export GOARCH $GOARCH; \
			echo "Echo building for $$GOOS - $$GOARCH" ; \
			go build -o $(NAME)-$$GOOS-$$GOARCH; \
		done \
    done