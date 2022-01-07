HOSTNAME=hashicorp.com
NAMESPACE=Groupe-Hevea
NAME=phpipam
VERSION=1.3.0


OS_ARCH=$(shell go env GOHOSTOS)_$(shell go env GOHOSTARCH)

INSTALL_DIR=~/.terraform.d/plugins/$(HOSTNAME)/$(NAMESPACE)/$(NAME)/$(VERSION)/$(OS_ARCH)

test:
	go test -v $(shell go list ./... | grep -v /vendor/)

testacc:
	TF_ACC=1 go test -v ./plugin/providers/phpipam -run="TestAcc"

test-local: build deploy

build:
	goreleaser build --snapshot --single-target --rm-dist

deploy: build
	mkdir -p $(INSTALL_DIR)
	cp $(shell ls dist/terraform-provider-phpipam_linux_amd64/*) $(INSTALL_DIR)
	chmod +x $(INSTALL_DIR)/*

release: release_bump release_build

release_bump:
	scripts/release_bump.sh

release_build:
	scripts/release_build.sh

clean:
	rm -rf dist/
	cd examples && $(MAKE) clean

test_tf: deploy
	cd examples && $(MAKE) all
