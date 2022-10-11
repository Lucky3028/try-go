package app_errors

type ErrCode string

const (
	Unknown ErrCode = "U000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	DataNotFound     ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
)

func (code ErrCode) Wrap(err error, message string) error {
	return &ApplicationError{ErrCode: code, Message: message, Err: err}
}
