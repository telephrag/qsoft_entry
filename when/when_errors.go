package when

import "errors"

var (
	ErrYearNotAnInteger = errors.New("integer expected as a request parameter")
)
