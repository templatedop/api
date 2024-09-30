package handler

import "github.com/templatedop/api/examples/server/core/domain"

type ListUsersResponse struct {
	Users []*domain.User `json:"users"`
	Meta  meta           `json:"meta"`
}