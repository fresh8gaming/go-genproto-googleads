PKG := github.com/fresh8gaming/go-genproto-googleads
GOOGLE_PROTO := google.golang.org/genproto/googleapis/ads/googleads

TARGETS := v11
SRC := googleapis/bazel-bin/google/ads/googleads/$(VERSION)/gapi-ads-googleads-$(VERSION)-go.tar.gz

BUILD_PATH := googleapis/google/ads/googleads/v11
BUILD_FILE := BUILD.bazel

all: $(TARGETS)

$(TARGETS): googleapis
	make add-protos
	-rm -rf build
	make gen VERSION=$@

gen: _build
	-rm -rf $(VERSION) pb/$(VERSION)
	cp -a build/$(PKG)/$(VERSION) $(VERSION)
	cp -a build/$(GOOGLE_PROTO)/$(VERSION) pb/$(VERSION)

_build: $(SRC)
	mkdir build
	tar zxf $(SRC) -C build
	grep -l -R '$(GOOGLE_PROTO)' build/ | \
		xargs sed -i.bak -e 's,$(GOOGLE_PROTO),$(PKG)/pb,'
	find build -name '*.bak' -exec rm {} \;

$(SRC): googleapis
	cd googleapis; bazel build //google/ads/googleads/$(VERSION):gapi-ads-googleads-$(VERSION)-go

googleapis:
	git clone --depth=1 --branch master https://github.com/googleapis/googleapis

add-protos:
	cat go-bazel/v11.bazel >> $(BUILD_PATH)/$(BUILD_FILE)
	cat go-bazel/common.bazel >> $(BUILD_PATH)/common/$(BUILD_FILE)
	cat go-bazel/enums.bazel >> $(BUILD_PATH)/enums/$(BUILD_FILE)
	cat go-bazel/errors.bazel >> $(BUILD_PATH)/errors/$(BUILD_FILE)
	cat go-bazel/resources.bazel >> $(BUILD_PATH)/resources/$(BUILD_FILE)
	cat go-bazel/services.bazel >> $(BUILD_PATH)/services/$(BUILD_FILE)

clean:
	-rm -rf build
	-rm -rf googleapis
