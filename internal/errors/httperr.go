package errors

import (
	"errors"
	"log"
	"maps"
	"slices"

	"github.com/danielgtaylor/huma/v2"
)

func fieldErrors(fields map[string]string) []error {
	errs := make([]error, 0, len(fields))
	for _, name := range slices.Sorted(maps.Keys(fields)) {
		errs = append(errs, &huma.ErrorDetail{Message: fields[name], Location: "body." + name})
	}
	return errs
}

func From(err error) error {

	var appErr *AppError

	if !errors.As(err, &appErr) {
		log.Printf("Unexpected Error : %v", err)
		return huma.Error500InternalServerError("Something went wrong")
	}

	switch appErr.Kind {
	case KindInvalid:
		return huma.Error422UnprocessableEntity(appErr.Message, fieldErrors(appErr.Fields)...)
	case KindUnauthenticated:
		return huma.Error401Unauthorized(appErr.Message)
	case KindForbidden:
		return huma.Error403Forbidden(appErr.Message)
	case KindNotFound:
		return huma.Error404NotFound(appErr.Message)
	case KindConflict:
		return huma.Error409Conflict(appErr.Message)
	default:
		log.Printf("error with unmapped kind %d: %v", appErr.Kind, err)
		return huma.Error500InternalServerError("something went wrong")
	}

}
