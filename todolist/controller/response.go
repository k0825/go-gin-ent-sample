package controller

import (
	"time"

	"github.com/google/uuid"
)

type jsonTime struct {
	time.Time
}

// 出力形式はRFC3339で指定
func (j jsonTime) format() string {
	return j.Format(time.RFC3339)
}

func (j jsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}

type todoFindApiResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Tags        []string  `json:"tags"`
	StartsAt    jsonTime  `json:"starts_at"`
	EndsAt      jsonTime  `json:"ends_at"`
	CreatedAt   jsonTime  `json:"created_at"`
	UpdatedAt   jsonTime  `json:"updated_at"`
}

type todoCreateApiResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Tags        []string  `json:"tags"`
	StartsAt    jsonTime  `json:"starts_at"`
	EndsAt      jsonTime  `json:"ends_at"`
	CreatedAt   jsonTime  `json:"created_at"`
	UpdatedAt   jsonTime  `json:"updated_at"`
}
