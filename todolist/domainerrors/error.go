package domainerrors

import "fmt"

type DomainValueError struct {
	message string
}

func (e *DomainValueError) Error() string {
	return e.message
}

func NewDomainValueError(message string) *DomainValueError {
	return &DomainValueError{message: message}
}

type InvalidValueError struct {
	key   string
	value string
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("%s is invalid value: %s", e.key, e.value)
}

func NewInvalidValueError(key string, value string) *InvalidValueError {
	return &InvalidValueError{key: key, value: value}
}

type NotFoundError struct {
	FieldName string
	Id        string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s %s is does not exist", e.FieldName, e.Id)
}

func NewNotFoundError(fieldName string, id string) *NotFoundError {
	return &NotFoundError{FieldName: fieldName, Id: id}
}

type InternalServerError struct {
	message string
}

func (e *InternalServerError) Error() string {
	return e.message
}

func NewInternalServerError(message string) *InternalServerError {
	return &InternalServerError{message: message}
}
