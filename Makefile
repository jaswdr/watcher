DOCKER=docker
TOKEN=$(GITHUB_TOKEN)
DOCKER_RUN=$(DOCKER) run \
		   -it \
		   --rm \
		   -e GITHUB_TOKEN=$(TOKEN) \
		   -v $(HOME)/.ssh:/root/.ssh \
		   -v $(HOME)/.gnupg:/root/.gnupg \
		   -v $(PWD):/go/src/github.com/jaswdr/watch \
		   -w /go/src/github.com/jaswdr/watch \
		   jaschweder/golang:alpine

default : test

clean :
	rm -rf ./dist

shell :
	$(DOCKER_RUN) sh

dep :
	$(DOCKER_RUN) dep ensure -v

test :
	$(DOCKER_RUN) go run main.go "clear && date" .tmp

test-error :
	$(DOCKER_RUN) go run main.go

release :
	goreleaser

snapshot :
	goreleaser --snapshot
