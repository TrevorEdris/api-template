SHELL ?= /bin/bash
export IMAGEORG ?= tedris
export IMAGE ?= template-golang-kubernetes
export VERSION ?= $(shell printf "`./util/version`${VERSION_SUFFIX}")
export GIT_HASH =$(shell git rev-parse HEAD)
export PROJECT_SLUG := ${IMAGEORG}-${IMAGE}
export DEV_DOCKER_COMPOSE=deployments/local/docker-compose.dev.yaml

# Blackbox files that need to be decrypted.
clear_files=$(shell blackbox_list_files)
encrypt_files=$(patsubst %,%.gpg,${clear_files})

# =========================[ Common Targets ]========================
# These are targets that almost certainly will not need to be changed
# as they are common to nearly all repos.
# ===================================================================

.PHONY: all
all: build

.PHONY: help
help: ## List of available commands
	@awk 'BEGIN {FS = ":.*?##"} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1 $$2}' $(MAKEFILE_LIST)

# -------------------------[ General Tools ]-------------------------

.PHONY: clear
clear: ${clear_files}

${clear_files}: ${encrypt_files}
	@blackbox_decrypt_all_files

.PHONY: decrypt
decrypt: ${clear_files}

.PHONY: encrypt
encrypt: ${encrypt_files}
	blackbox_edit_end $^

.PHONY: submodules
	@git submodule update --init --recursive || printf "\nWarning: Could not pull submodules\n"

.PHONY: version
version: submodules util/version
	@echo ${VERSION}

# =========================[ Custom Targets ]========================
# These are targets that _may_ need to be customized to the specific
# project implemented in the repo.
# ===================================================================

# ---------------------------[ Local App ]---------------------------
.PHONY: up
up: ## Run the API locally and print logs to stdout
	docker-compose -f ${DEV_DOCKER_COMPOSE} up -d
	make -s logs

.PHONY: down
down: ## Stop all containers
	docker-compose -f ${DEV_DOCKER_COMPOSE} down

.PHONY: restart
restart: ## Restart all containers
	docker-compose -f ${DEV_DOCKER_COMPOSE} restart

.PHONY: logs
logs: ## Print logs in stdout
	docker-compose -f ${DEV_DOCKER_COMPOSE} logs -f api

# -----------------------------[ Build ]-----------------------------

.PHONY: build
build: decrypt submodules version
	@docker build -f deployments/container/Dockerfile -t ${IMAGEORG}/${IMAGE}:${VERSION} --target builder .
	@docker tag ${IMAGEORG}/${IMAGE}:${VERSION} ${IMAGEORG}/${IMAGE}:latest
	@docker tag ${IMAGEORG}/${IMAGE}:${VERSION} ${IMAGEORG}/${IMAGE}-build:latest

# -----------------------------[ Test ]------------------------------

.PHONY: test
test: build
	@test/test_unit

# -----------------------------[ Publish ]---------------------------

.PHONY: finalize
finalize: test
	@docker build -f container/Dockerfile -t ${IMAGEORG}/${IMAGE}:${VERSION} .
	@docker tag ${IMAGEORG}/${IMAGE}:${VERSION} ${IMAGEORG}/${IMAGE}:latest

.PHONY: publish_only
publish_only:
	@docker push ${IMAGEORG}/${IMAGE}:${VERSION}

.PHONY: publish
publish: finalize publish_only

# -----------------------------[ Deploy ]----------------------------

.PHONY: deploy_only
deploy_only: decrypt
	@kube/deploy

.PHONY: deploy
deploy: publish deploy_only

# ----------------------------[ Release ]----------------------------
# TODO
