package main

import (
	"errors"
	dockerdistributionerrcode "github.com/docker/distribution/registry/api/errcode"
	dockerdistributionapi "github.com/docker/distribution/registry/api/v2"
)


// isNotFoundImageError heuristically attempts to determine whether an error
// is saying the remote source couldn't find the image (as opposed to an
// authentication error, an I/O error etc.)
// TODO drive this into containers/image properly
func isNotFoundImageError(err error) bool {
	return isDockerManifestUnknownError(err)
}

// isDockerManifestUnknownError is a copy of code from containers/image,
// please update there first.
func isDockerManifestUnknownError(err error) bool {
	var ec dockerdistributionerrcode.ErrorCoder
	if !errors.As(err, &ec) {
		return false
	}
	return ec.ErrorCode() == dockerdistributionapi.ErrorCodeManifestUnknown
}
