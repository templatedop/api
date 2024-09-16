package common

import (
	"net/http"

	"github.com/templatedop/api/util/wrapper"
)

type (
	RoundTripperWrapper = wrapper.Wrapper[http.RoundTripper]
)
