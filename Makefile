PLATFORMS := linux/amd64 darwin/amd64 linux/386 darwin/386 windows/amd64/.exe windows/386/.exe

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))
ext = $(word 3, $(temp))
package := cscurl
base = $(package)$(ext)
out_dir = dist/$(os)/$(arch)
out_file = $(out_dir)/$(base)
main := cscurl.go
current_dir := $(shell pwd)

build:
	mkdir -p dist
	go build -o dist/$(base) $(main)

package: $(PLATFORMS)

$(PLATFORMS):
	mkdir -p dist
	GOOS=$(os) GOARCH=$(arch) go build -o $(out_file) $(main)
	cd $(out_dir); tar czf $(current_dir)/dist/$(package)_$(arch)-$(os).tar.gz $(base)

upload: package
	github-release cloudshare/go-sdk $(TAG) master '' 'dist/*.gz'

clean:
	rm -rf dist

.PHONY: package $(PLATFORMS) build clean
