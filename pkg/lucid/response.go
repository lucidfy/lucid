package lucid

type Response struct {
	Type  string
	Value interface{}
}

type ResponseError struct {
	Message interface{}
	Error   error
	Code    interface{}
}

type ResponseValidationError struct {
	ValidationError interface{}
}
