package exceptions

type (
	NotFoundError struct {
		Message string
	}
	UnAuthorizedError struct {
		Message string
	}
	ValidationError struct {
		Message string
	}
)

func (err NotFoundError) Error() string {
	return err.Message
}

func (err UnAuthorizedError) Error() string {
	return err.Message
}

func (err ValidationError) Error() string {
	return err.Message
}
