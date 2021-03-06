all: build

deps:
	go get -d -v lexs.com/...

vet:
	@ for f in $$(find . -type f -iregex '.*\.go' | grep -v "vendor"); do \
		go fmt $$f; \
	done 

	@ for f in $$(find . -type f -iregex '.*\.go'  | grep -v "vendor"); do \
		go tool vet $$f; \
	done 

updatedeps:
	go get -d -v -u -f lexs.com/...

testdeps:
	go get -d -v -t lexs.com/...

updatetestdeps:
	go get -d -v -t -u -f lexs.com/...

build:
	bash/build.sh

gen:
	@ if ! which thrift > /dev/null; then \
		echo "error: thrift not installed" >&2; \
		exit 1; \
	fi

	@ for f in $$(find proto -type f -iregex '.*\.thrift'); do \
		thrift -out proto -r --gen go $$f; \
	done

	@ for d in $$(find proto/* -type 'd' | grep "_sevice-remote"); do \
		rm -rf $$d; \
	done 
	
	@ for f in $$(find . -type f -iregex '.*\.go'  | grep 'consts.go\|GoUnusedProtection__.go'); do \
		rm -f $$f; \
	done

test: testdeps
	go test -v -cpu 1,4 lexs.com/...

testrace: testdeps
	go test -v -race -cpu 1,4 lexs.com/...

clean:
	@ #go clean -i lexs.com/...
	@ rm -rf _output
	
	@ for d in $$(find proto/* -type 'd' | grep -v "thrift"); do \
		rm -rf $$d; \
	done 

coverage: testdeps
	./coverage.sh --coveralls

.PHONY: \
	all \
	deps \
	vet	\
	updatedeps \
	testdeps \
	updatetestdeps \
	build \
	gen \
	test \
	testrace \
	clean \
	coverage

