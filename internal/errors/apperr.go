package errors

type Kind int

const (
	KindInternal Kind = iota
	KindInvalid
	KindUnauthenticated
	KindForbidden
	KindNotFound
	KindConflict
)

type AppError struct {
	Kind    Kind
	Message string
	Fields  map[string]string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(kind Kind, message string) *AppError {
	return &AppError{Kind: kind, Message: message}
}

func Invalid(fields map[string]string) *AppError {
	return &AppError{Kind: KindInvalid, Message: "Validation Failed", Fields: fields}
}
