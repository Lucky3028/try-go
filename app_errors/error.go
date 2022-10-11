package app_errors

type ApplicationError struct {
	ErrCode
	Message string
	Err     error `json:"-"`
}

func (appErr ApplicationError) Error() string {
	return appErr.Err.Error()
}

func (appErr ApplicationError) Unwrap() error {
	return appErr.Err
}
