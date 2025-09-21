package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/dto"
	"vmapp/app0/internal/models"
	"vmapp/pkg/vhttp"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type RoleUseCase struct {
	rp   repo.RoleRepo
	up   repo.UserRepo
	mp   repo.MenuRepo
	rbac repo.RbacRepo
	tx   repo.Transaction
	log  *log.Helper
}

func NewRoleUseCase(rp repo.RoleRepo, tx repo.Transaction, bc *conf.BootComponent, up repo.UserRepo, mp repo.MenuRepo, rbac repo.RbacRepo) *RoleUseCase {
	return &RoleUseCase{
		rp:   rp,
		tx:   tx,
		log:  bc.Logger,
		up:   up,
		mp:   mp,
		rbac: rbac,
	}
}

func (t *RoleUseCase) CreateRole(ctx context.Context, req *dto.CreateRoleReq) error {
	if req.MenuIds != "" {
		ids := strings.Split(req.MenuIds, ",")
		all, err := t.mp.GetMenuByIds(ctx, ids)
		if err != nil || len(all) != len(ids) {
			return vhttp.NewError(http.StatusBadRequest, "角色与提交的不匹配,非法提交")
		}
		if req.IndexMenu != "" {
			for _, v := range all {
				if fmt.Sprintf("%v", v.Id) == req.IndexMenu {
					if v.Type != 1 {
						return vhttp.NewError(http.StatusBadRequest, "角色首页类型应该为菜单")
					}
				}
			}
		}
	}
	var one models.Role
	if err := copier.Copy(&one, &req); err != nil {
		t.log.WithContext(ctx).Errorf("数据合并失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据合并失败", vhttp.WithReason(err))
	}
	one.CreatedAt = time.Now()
	if err := t.rp.CreateRole(ctx, &one); err != nil {
		t.log.WithContext(ctx).Errorf("创建角色失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "创建角色失败", vhttp.WithReason(err))
	}
	// 添加相关权限
	if one.MenuIds != "" {
		all, err := t.mp.GetMenuByIds(ctx, strings.Split(one.MenuIds, ","))
		if err != nil {
			t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
		}
		// 组装数据，然后写入
		var Policy []*repo.Policy = make([]*repo.Policy, 0, len(all))
		for _, v := range all {
			if v.ApiUrl != "" && v.ApiMethod != "" {
				Policy = append(Policy, &repo.Policy{
					Source: v.ApiUrl,
					Action: v.ApiMethod,
				})
			}
			if v.Href != "" {
				Policy = append(Policy, &repo.Policy{
					Source: v.Href,
					Action: http.MethodGet,
				})
			}
		}
		if err := t.rbac.AutoAddRbac(ctx, one.RoleCode, Policy); err != nil {
			t.log.WithContext(ctx).Errorf("cas添加失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "cas添加失败", vhttp.WithReason(err))
		}

	}
	return nil
}

func (t *RoleUseCase) GetUserRole(ctx context.Context, uid int64) (all []*SelectMenu, err error) {
	roles, err := t.rp.ListRole(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取角色列表失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取角色列表失败", vhttp.WithReason(err))
	}
	// 查询出用户的权限列表
	user, err := t.up.GetUser(ctx, uid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}
	// 查看用户的权限
	ucodes := strings.Split(user.RoleCodes, ",")
	// 列表展现
	for _, v := range roles {
		all = append(all, &SelectMenu{
			Name:  v.RoleName,
			Value: v.RoleCode,
			Selected: func() bool {
				for _, v2 := range ucodes {
					if v.RoleCode == v2 {
						return true
					}
				}
				return false
			}(),
		})
	}

	return all, nil
}

func (t *RoleUseCase) ListRole(ctx context.Context) (all []*models.Role, err error) {
	all, err = t.rp.ListRole(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取角色列表失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取角色列表失败", vhttp.WithReason(err))
	}
	return
}

func (t *RoleUseCase) UpdateRole(ctx context.Context, req *dto.UpdateRoleReq) error {
	if req.MenuIds != "" {
		ids := strings.Split(req.MenuIds, ",")
		all, err := t.mp.GetMenuByIds(ctx, ids)
		if err != nil || len(all) != len(ids) {
			return vhttp.NewError(http.StatusBadRequest, "角色与提交的不匹配,非法提交")
		}
		if req.IndexMenu != "" {
			for _, v := range all {
				if fmt.Sprintf("%v", v.Id) == req.IndexMenu {
					if v.Type != 1 {
						return vhttp.NewError(http.StatusBadRequest, "角色首页类型应该为菜单")
					}
				}
			}
		}
	}
	one, err := t.rp.GetRole(ctx, req.Id)
	if err != nil {
		t.log.WithContext(ctx).Errorf("数据查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据查询失败", vhttp.WithReason(err))
	}
	if err := copier.Copy(&one, &req); err != nil {
		t.log.WithContext(ctx).Errorf("数据合并失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据合并失败", vhttp.WithReason(err))
	}
	one.UpdatedAt = time.Now()
	if err := t.rp.UpdateRole(ctx, one); err != nil {
		t.log.WithContext(ctx).Errorf("更新角色失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "更新角色失败", vhttp.WithReason(err))
	}
	if err := t.rbac.AutoAddRbac(ctx, one.RoleCode, nil); err != nil {
		t.log.WithContext(ctx).Errorf("cas移除失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "cas移除失败", vhttp.WithReason(err))
	}

	// 添加/更新相关权限
	if req.MenuIds != "" {
		all, err := t.mp.GetMenuByIds(ctx, strings.Split(req.MenuIds, ","))
		if err != nil {
			t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
		}
		// 组装数据，然后写入
		var Policy []*repo.Policy
		for _, v := range all {
			if v.ApiUrl != "" && v.ApiMethod != "" {
				Policy = append(Policy, &repo.Policy{
					Source: v.ApiUrl,
					Action: v.ApiMethod,
				})
			}
			if v.Href != "" {
				Policy = append(Policy, &repo.Policy{
					Source: v.Href,
					Action: http.MethodGet,
				})
			}
		}
		if err := t.rbac.AutoAddRbac(ctx, one.RoleCode, Policy); err != nil {
			t.log.WithContext(ctx).Errorf("cas添加失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "cas添加失败", vhttp.WithReason(err))
		}

	}
	return nil
}

func (t *RoleUseCase) DelectRole(ctx context.Context, req *dto.DeleteRoleReq) error {
	one, err := t.rp.GetRole(ctx, req.ID)
	if err != nil {
		t.log.WithContext(ctx).Errorf("数据查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据查询失败", vhttp.WithReason(err))
	}
	// 查询是否有用户关联了角色
	all, err := t.up.GetAllRoleCodes(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("用户角色查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户角色查询失败", vhttp.WithReason(err))
	}

	if len(all) > 0 {
		var codes []string
		for _, v := range all {
			codes = append(codes, strings.Split(v, ",")...)
		}

		for _, v := range codes {
			if one.RoleCode == v {
				return vhttp.NewError(http.StatusBadRequest, "当前角色已有用户关联，不能删除！")
			}
		}
	}

	if err := t.rbac.AutoAddRbac(ctx, one.RoleCode, nil); err != nil {
		t.log.WithContext(ctx).Errorf("cas移除失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "cas移除失败", vhttp.WithReason(err))
	}
	err = t.rp.DeleteRole(ctx, req.ID)
	if err != nil {
		t.log.WithContext(ctx).Errorf("删除角色失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "删除角色失败", vhttp.WithReason(err))
	}

	return nil
}

func (t *RoleUseCase) GetRole(ctx context.Context, req *dto.GetRoleReq) (*models.Role, error) {
	res, err := t.rp.GetRole(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, vhttp.NewError(http.StatusBadRequest, "数据不存在", vhttp.WithReason(err))
		}
		t.log.WithContext(ctx).Errorf("获取角色失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "角色获取失败", vhttp.WithReason(err))
	}
	return res, nil
}

func (t *RoleUseCase) GetMenuFromRoles(ctx context.Context, Roles []string) (all []*models.Menu, err error) {
	roles, err := t.rp.GetRoleByCodes(ctx, Roles)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取角色数据失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "角色获取失败", vhttp.WithReason(err))
	}
	var ids []string
	for _, v := range roles {
		if v.MenuIds != "" {
			ids = append(ids, strings.Split(v.MenuIds, ",")...)
		}
	}
	// 查询出菜单
	all, err = t.mp.GetMenuByIds(ctx, ids)
	if err != nil {
		t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
	}
	return
}
