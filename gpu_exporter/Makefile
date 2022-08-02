DOCKER_IMG=aibox03.bj.intel.com:5000/dgpu-exporter:latest

.PHONY: test
test:
	@echo "\n🛠️  Running unit tests..."
	go test ./... -v

.PHONY: build
build:
	@echo "\n🔧  Building Go binaries..."
	GOOS=linux GOARCH=amd64 go build -o bin/gpu-exporter .

.PHONY: image
image:
	@echo "\n📦 Building ycm-webhook Docker image..."
	DOCKER_BUILDKIT=1 docker build --build-arg http_proxy=${http_proxy} --build-arg https_proxy=${https_proxy} -t ${DOCKER_IMG} .

.PHONY: push
push: image
	@echo "\n📦 Pushing ycm-webhook image..."
	docker push ${DOCKER_IMG}
