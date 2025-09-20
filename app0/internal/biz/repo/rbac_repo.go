package repo

import "context"

type Policy struct {
	Source string
	Action string
}
type RbacRepo interface {
	AutoAddRbac(ctx context.Context, roleID string, policies []*Policy) error
}
