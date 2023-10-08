package models

import (
	"regexp"

	"github.com/google/uuid"
)

type TodoId struct {
	value uuid.UUID
}

type TodoTitle struct {
	value string
}

type TodoDescription struct {
	value string
}

type TodoImage struct {
	value string
}

type TodoTag struct {
	value string
}

var uuidRegexp = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

func NewTodoId(value uuid.UUID) *TodoId {
	return &TodoId{value: value}
}

func (todoId TodoId) Value() uuid.UUID {
	return todoId.value
}

func (todoId TodoId) String() string {
	return todoId.value.String()
}

func NewTodoTitle(value string) (*TodoTitle, error) {
	return &TodoTitle{value: value}, nil
}

func (todoTitle TodoTitle) Value() string {
	return todoTitle.value
}

func (todoTitle TodoTitle) String() string {
	return todoTitle.value
}

func NewTodoDescription(value string) (*TodoDescription, error) {
	return &TodoDescription{value: value}, nil
}

func (todoDescription TodoDescription) Value() string {
	return todoDescription.value
}

func (todoDescription TodoDescription) String() string {
	return todoDescription.value
}

func NewTodoImage(value string) (*TodoImage, error) {
	return &TodoImage{value: value}, nil
}

func (todoImage TodoImage) Value() string {
	return todoImage.value
}

func (todoImage TodoImage) String() string {
	return todoImage.value
}

func NewTodoTag(value string) (*TodoTag, error) {
	return &TodoTag{value: value}, nil
}

func (todoTag TodoTag) Value() string {
	return todoTag.value
}

func (todoTag TodoTag) String() string {
	return todoTag.value
}
