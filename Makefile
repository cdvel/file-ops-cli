VERSION = "0.4.0"
.DEFAULT_GOAL := build
BIN_FILE=fops

release:
	docker build . -t fops-cli:latest
	docker run --rm --privileged \
		-v $$PWD:/go/src/github.com/cdvel/fops-cli \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-w /go/src/github.com/cdvel/fops-cli \
		-e GITHUB_TOKEN=${GITHUB_TOKEN} \
		fops-cli:latest --rm-dist

build:
	go build -o "${BIN_FILE}" --ldflags "-X fops-cli/cmd.Version=${VERSION}"

build_all:
	xgo -o build/"${BIN_FILE}" --targets "darwin/amd64,windows/amd64,linux/amd64" --ldflags "-X cmd.Version=${VERSION}" ./

run:
	"./${BIN_FILE}"