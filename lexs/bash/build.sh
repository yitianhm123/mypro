#!/bin/bash -e

# set some environment variables
readonly LEXS_ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
readonly LEXS_OUTPUT="${LEXS_ROOT}/_output/local"
readonly LEXS_OUTPUT_SRCPATH="${LEXS_OUTPUT}/src"
readonly LEXS_OUTPUT_BINPATH="${LEXS_OUTPUT}/bin"

readonly LEXS_TARGETS=(
	cmd/account
	cmd/property
  )

eval $(go env)

# enable/disable failpoints
toggle_failpoints() {
	FAILPKGS="lexsserver/ lexsserver/auth/"

	mode="disable"
	if [ ! -z "$FAILPOINTS" ]; then mode="enable"; fi
	if [ ! -z "$1" ]; then mode="$1"; fi

	if which gofail >/dev/null 2>&1; then
		gofail "$mode" $FAILPKGS
	elif [ "$mode" != "disable" ]; then
		echo "FAILPOINTS set but gofail not found"
		exit 1
	fi
}

lexs_setup_gopath() {
	# preserve old gopath to support building with unvendored tooling deps (e.g., gofail)
	if [ -n "$GOPATH" ]; then
		GOPATH=":$GOPATH"
	fi
	export GOPATH=${LEXS_OUTPUT}

	rm -rf ${LEXS_OUTPUT_SRCPATH}
	mkdir -p ${LEXS_OUTPUT_SRCPATH}

	ln -s ${LEXS_ROOT} ${LEXS_OUTPUT_SRCPATH}/lexs
}

lexs_build_target() {
	toggle_failpoints
	
	for arg; do
		# echo "target: ${arg}, ${arg##*/}"
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $GO_BUILD_FLAGS \
		-installsuffix cgo -ldflags "$GO_LDFLAGS" \
		-o ${LEXS_OUTPUT_BINPATH}/${arg##*/}.x lexs/${arg} || return
	done
}

lexs_make_ldflag() {
  local key=${1}
  local val=${2}

  echo "-X lexs/cmd/version.${key}=${val}"
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the project on the git tree status.
lexs_version_ldflags() {
	local -a ldflags=($(lexs_make_ldflag "buildDate" "$(date -u +'%Y-%m-%dT%H:%M:%SZ')"))

	local git_sha=`git rev-parse --short HEAD || echo "GitNotFound"`
	if [ ! -z "$FAILPOINTS" ]; then
		git_sha="$git_sha"-FAILPOINTS
	fi

	ldflags+=($(lexs_make_ldflag "gitSHA" "${git_sha}"))

	echo "${ldflags[*]-}"
}

toggle_failpoints

# only build when called directly, not sourced
if echo "$0" | grep "build.sh$" >/dev/null; then
	# force new gopath so builds outside of gopath work
	lexs_setup_gopath
	lexs_version_ldflags
	#lexs_build
	lexs_build_target "${LEXS_TARGETS[@]}"
fi
