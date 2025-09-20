package rbac_repo

import (
	"context"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/data/base"

	"github.com/go-kratos/kratos/v2/log"
)

type rbacRepo struct {
	data *base.Data
	log  *log.Helper
}

func NewCasRepo(data *base.Data, component *conf.BootComponent) repo.RbacRepo {
	return &rbacRepo{
		data: data,
		log:  component.Logger,
	}
}

func (t *rbacRepo) AutoAddRbac(ctx context.Context, role string, policies []*repo.Policy) error {
	if policies == nil {
		return t.data.Rbac.AutoAddRbacs(role, nil)
	}
	var rules [][]string
	for _, v := range policies {
		rules = append(rules, []string{v.Source, v.Action})
	}

	return t.data.Rbac.AutoAddRbacs(role, rules)

}
