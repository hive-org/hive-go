package hive_go

import "github.com/satori/go.uuid"

type IAuthenticationBackendUser interface {
	GetIsAdmin() bool
	GetRoles() []string
	GetUserID() uuid.UUID
}

type Secret struct {
	Id      uuid.UUID
	Created int64
	Value   uuid.UUID
}
