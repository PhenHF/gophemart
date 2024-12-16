package jwtauth

type InvalidTokenError struct {
	Err error
}

func (te *InvalidTokenError) Error() string {
	return "invalid token error"
}

func newInvalidTokenError(err error) error {
	return &InvalidTokenError{
		Err: err,
	}
}
