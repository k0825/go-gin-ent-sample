package models

import (
	"fmt"
	"time"

	"github.com/k0825/go-gin-ent-sample/domainerrors"
)

type Todo struct {
	id          TodoId
	title       TodoTitle
	description TodoDescription
	image       TodoImage
	tags        []TodoTag
	startsAt    time.Time
	endsAt      time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTodo(
	id TodoId,
	title TodoTitle,
	description TodoDescription,
	image TodoImage,
	tags []TodoTag,
	startsAt time.Time,
	endsAt time.Time,
	createdAt time.Time,
	updatedAt time.Time) (*Todo, error) {

	if !uuidRegexp.MatchString(id.String()) {
		return nil, domainerrors.NewDomainValueError(fmt.Sprintf("id is not in uuid format: %s", id.String()))
	}

	if startsAt.After(endsAt) {
		return nil, domainerrors.NewDomainValueError(fmt.Sprintf("endsAt is not later than startsAt: %s, %s", startsAt, endsAt))
	}

	if createdAt.After(updatedAt) {
		return nil, domainerrors.NewDomainValueError(fmt.Sprintf("updatedAt is not later than createdAt: %s, %s", createdAt, updatedAt))
	}

	return &Todo{
		id:          id,
		title:       title,
		description: description,
		image:       image,
		tags:        tags,
		startsAt:    startsAt,
		endsAt:      endsAt,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}, nil
}

type PaginationMeta struct {
	Start int
	Take  int
	Total int
}

func NewPaginationMeta(start int, take int, total int) *PaginationMeta {
	return &PaginationMeta{
		Start: start,
		Take:  take,
		Total: total,
	}
}

func (todo Todo) GetId() TodoId {
	return todo.id
}

func (todo Todo) GetTitle() TodoTitle {
	return todo.title
}

func (todo Todo) GetDescription() TodoDescription {
	return todo.description
}

func (todo Todo) GetImage() TodoImage {
	return todo.image
}

func (todo Todo) GetTags() []TodoTag {
	return todo.tags
}

func (todo Todo) GetStartsAt() time.Time {
	return todo.startsAt
}

func (todo Todo) GetEndsAt() time.Time {
	return todo.endsAt
}

func (todo Todo) GetCreatedAt() time.Time {
	return todo.createdAt
}

func (todo Todo) GetUpdatedAt() time.Time {
	return todo.updatedAt
}
