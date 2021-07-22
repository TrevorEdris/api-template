SHELL ?= /bin/bash
export IMAGEORG ?= tedris
export IMAGE ?= template-golang-kubernetes
export VERSION ?= $(shell printf "`./util/version`${VERSION_SUFFIX}")
export GIT_HASH =$(shell git rev-parse HEAD)
export PROJECT_SLUG := ${IMAGEORG}-${IMAGE}

# Blackbox files that need to be decrypted.
clear_files=$(shell blackbox_list_files)
encrypt_files=$(patsubst %,%.gpg,${clear_files})

# =========================[ Common Targets ]========================

.PHONY: all
all: build

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

# -----------------------------[ Build ]-----------------------------

.PHONY: build
build: decrypt submodules version
	@docker build -f container/Dockerfile -t ${IMAGEORG}/${IMAGE}:${VERSION} --target builder .
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
