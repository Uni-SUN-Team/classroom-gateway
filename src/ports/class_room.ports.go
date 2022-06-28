package ports

import "unisun/api/classroom-gateway/src/model/original"

type ClassRoomService interface {
	FindAll() (*original.ResponseSuccess, error)
	FindById(id string) (*original.ResponseSingleSuccess, error)
	FindBySlug(slug string) (*original.ResponseSuccess, error)
}
