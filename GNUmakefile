TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
WEBSITE_REPO=github.com/hashicorp/terraform-website
PKG_NAME=hdns
export CGO_ENABLED:=0

default: build

build: fmtcheck
	go install

clean:
	@rm -rf bin

clean-release:
	@rm -rf _output

release: \
	clean \
	clean-release \
	_output/terraform-provider-hdns_linux_amd64.zip \
	_output/terraform-provider-hdns_darwin_amd64.zip \
	_output/terraform-provider-hdns_freebsd_amd64.zip \
	_output/terraform-provider-hdns_freebsd_386.zip \
	_output/terraform-provider-hdns_freebsd_arm.zip \
	_output/terraform-provider-hdns_linux_amd64.zip \
	_output/terraform-provider-hdns_linux_386.zip \
	_output/terraform-provider-hdns_linux_arm.zip \
	_output/terraform-provider-hdns_openbsd_amd64.zip \
	_output/terraform-provider-hdns_openbsd_386.zip \
	_output/terraform-provider-hdns_solaris_amd64.zip \
	_output/terraform-provider-hdns_windows_amd64.zip \
	_output/terraform-provider-hdns_windows_386.zip

bin/darwin_amd64/terraform-provider-hdns:  GOARGS = GOOS=darwin GOARCH=amd64
bin/freebsd_amd64/terraform-provider-hdns:  GOARGS = GOOS=freebsd GOARCH=amd64
bin/freebsd_386/terraform-provider-hdns:  GOARGS = GOOS=freebsd GOARCH=386
bin/freebsd_arm/terraform-provider-hdns:  GOARGS = GOOS=freebsd GOARCH=arm
bin/linux_amd64/terraform-provider-hdns:  GOARGS = GOOS=linux GOARCH=amd64
bin/linux_386/terraform-provider-hdns:  GOARGS = GOOS=linux GOARCH=386
bin/linux_arm/terraform-provider-hdns:  GOARGS = GOOS=linux GOARCH=arm
bin/openbsd_amd64/terraform-provider-hdns:  GOARGS = GOOS=openbsd GOARCH=amd64
bin/openbsd_386/terraform-provider-hdns:  GOARGS = GOOS=openbsd GOARCH=386
bin/solaris_amd64/terraform-provider-hdns:  GOARGS = GOOS=solaris GOARCH=amd64
bin/windows_amd64/terraform-provider-hdns:  GOARGS = GOOS=windows GOARCH=amd64
bin/windows_386/terraform-provider-hdns:  GOARGS = GOOS=windows GOARCH=386

bin/%/terraform-provider-hdns: clean
	$(GOARGS) go build -o $@ -a .

_output/terraform-provider-hdns_%.zip: NAME=terraform-provider-hdns_$(VERSION)_$*
_output/terraform-provider-hdns_%.zip: DEST=_output/$(NAME)
_output/terraform-provider-hdns_%.zip: bin/%/terraform-provider-hdns
	mkdir -p $(DEST)
	cp bin/$*/terraform-provider-hdns README.md CHANGELOG.md LICENSE $(DEST)
	cd $(DEST) && zip -r ../$(NAME).zip .

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 30m

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck errcheck test-compile release clean clean-release website website-test

