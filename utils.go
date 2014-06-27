package main

import (
	"github.com/Toorop/govh"
)

// handleErrFromOvh handle error from OVH API
func handleErrFromOvh(err error) {
	if err == nil {
		return
	}
	if err.Error() == govh.ErrInvalidCredential.Error() || err.Error() == govh.ErrInvalidkey.Error() {
		dieInvalidConsumerKey()
	} else {
		dieError(err)
	}
}

// inslice return true if search is in slice, false otherwise
func inSliceStr(search string, in []string) bool {
	for _, v := range in {
		if v == search {
			return true
		}
	}
	return false
}
