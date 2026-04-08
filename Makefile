BIN=./build/bin
BUILD=./build
APP_NAME="basic_api"
IMAGE_TAG=$(shell git rev-parse --short HEAD)
FULL_IMAGE=$(DOCKER_USER)/$(APP_NAME):$(IMAGE_TAG)

.PHONY: build
build:
	@echo "building $(APP_NAME)..."
	@mkdir -p $(BIN)
	@go build -o $(BIN)/$(APP_NAME) .

.PHONY: clean
clean:
	@echo "cleaning service..."
	@rm $(BIN)/$(APP_NAME)

.PHONY: test
test:
	@echo "beginning testing..."
	@go test -race ./...

PORT=8080
.PHONY: run
run: build
	@echo "running $(APP_NAME)"
	@./$(BIN)/$(APP_NAME) \
		-port=$(PORT)

.PHONY: image
image:
	@echo "building image $(FULL_IMAGE)..."
	@podman build -t $(FULL_IMAGE) -f $(BUILD)/Dockerfile .

.PHONY: start-container
start-container: image
	@echo "Running Docker container $(FULL_IMAGE)..."
	@podman run -it --rm -p $(PORT):8080 $(FULL_IMAGE)

.PHONY: kustomize-dev
kustomize-dev:
	@echo "Generating complete YAML for DEV"
	@kubectl kustomize deploy/k8s/overlays/dev

.PHONY: kustomize-uat
kustomize-uat:
	@echo "Generating complete YAML for UAT"
	@kubectl kustomize deploy/k8s/overlays/uat


.PHONY: kustomize-%
kustomize-prod:
	@echo "Generating complete YAML for $*..."
	@kubectl kustomize deploy/k8s/overlays/$*

.PHONY: deploy-%
deploy-%:
	@echo "Deploying to $*..."
	@kubectl apply -k deploy/k8s/overlays/$*

