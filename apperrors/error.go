package apperrors

type MyAppError struct {
	ErrCode
	Message string
	Err     error
}

func (myErr *MyAppError) Error() string {
	//詳細なエラーロギングのための出力
	return myErr.Err.Error()
}

func (myErr *MyAppError) Unwrap() error {
	return myErr.Err
}
