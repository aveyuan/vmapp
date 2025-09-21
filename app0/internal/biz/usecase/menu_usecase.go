package usecase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

type SelectMenu struct {
	Name     string        `json:"name"`
	Value    interface{}   `json:"value"`
	Selected bool          `json:"selected"`
	Disabled bool          `json:"disabled"`
	Children []*SelectMenu `json:"children"`
}

type ListTreeMenu struct {
	models.Menu
	IsParent bool            `json:"is_parent"`
	Children []*ListTreeMenu `json:"children"`
}

type SysMenu struct {
	HomeInfo struct {
		Title string `json:"title"`
		Href  string `json:"href"`
	} `json:"homeInfo"`
	LogoInfo struct {
		Title string `json:"title"`
		Image string `json:"image"`
		Href  string `json:"href"`
	} `json:"logoInfo"`
	MenuInfo []*SysMenuList `json:"menuInfo"`
}

// MenuTreeList 菜单结构体
type SysMenuList struct {
	ID     int32          `json:"id"`
	Pid    int32          `json:"pid"`
	Title  string         `json:"title"`
	Icon   string         `json:"icon"`
	Href   string         `json:"href"`
	Target string         `json:"target"`
	Remark string         `json:"remark"`
	Child  []*SysMenuList `json:"child"`
}

type MenuUseCase struct {
	mp  repo.MenuRepo
	tx  repo.Transaction
	log *log.Helper
	bcf  *conf.Data
	rp  repo.RoleRepo
	bc  *conf.BootComponent
}

func NewMenuUseCase(mp repo.MenuRepo, tx repo.Transaction, bc *conf.BootComponent, bcf *conf.Data, rp repo.RoleRepo) *MenuUseCase {
	return &MenuUseCase{
		mp:  mp,
		tx:  tx,
		log: bc.Logger,
		bcf:  bcf,
		rp:  rp,
		bc:  bc,
	}
}

func (t *MenuUseCase) CreateMenu(ctx context.Context, req *dto.CreateMenuReq) error {
	var one models.Menu
	if err := copier.Copy(&one, &req); err != nil {
		t.log.WithContext(ctx).Errorf("数据合并失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据合并失败", vhttp.WithReason(err))
	}
	one.CreatedAt = time.Now()
	if err := t.mp.CreateMenu(ctx, &one); err != nil {
		t.log.WithContext(ctx).Errorf("创建菜单失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "创建菜单失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *MenuUseCase) ListMenu(ctx context.Context) (all []*models.Menu, err error) {
	all, err = t.mp.ListMenu(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取菜单列表失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取菜单列表失败", vhttp.WithReason(err))
	}
	return
}

func (t *MenuUseCase) UpdateMenu(ctx context.Context, req *dto.UpdateMenuReq) error {
	if req.Id == req.Pid {
		return vhttp.NewError(http.StatusBadRequest, "上级ID不能是当前ID")
	}
	one, err := t.mp.GetMenu(ctx, req.Id)
	if err != nil {
		t.log.WithContext(ctx).Errorf("数据查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据查询失败", vhttp.WithReason(err))
	}

	if one.Id == req.Pid {
		return vhttp.NewError(http.StatusBadRequest, "上级不能是自己")
	}

	if one.Pid != req.Pid {
		// 查询是转移到了下级引发逻辑错误
		res, err := t.mp.ListMenu(ctx)
		if err != nil {
			t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
		}

		// 查询是否在逻辑错误
		if t.isParentInSubmenu(res, req.Pid, one.Id) {
			return vhttp.NewError(http.StatusInternalServerError, "新的上级不能是自己的下级")
		}
	}

	if err := copier.Copy(&one, &req); err != nil {
		t.log.WithContext(ctx).Errorf("数据合并失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "数据合并失败", vhttp.WithReason(err))
	}
	one.UpdatedAt = time.Now()
	if err := t.mp.UpdateMenu(ctx, one); err != nil {
		t.log.WithContext(ctx).Errorf("更新菜单失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "更新菜单失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *MenuUseCase) DelectMenu(ctx context.Context, req *dto.DeleteMenuReq) error {
	one, err := t.mp.GetMenu(ctx, req.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
	}

	if one.Id == 0 {
		return vhttp.NewError(http.StatusBadRequest, "菜单不存在")
	}

	// 查询是否有下级
	res, err := t.mp.GetMenuByPid(ctx, req.ID)
	if err != nil {
		t.log.WithContext(ctx).Errorf("菜单获取失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
	}

	if len(res) > 0 {
		return vhttp.NewError(http.StatusBadRequest, "当前菜单包含子级，不可删除!")
	}

	// 查询角色是否关联了菜单
	all, err := t.rp.GetAllMenuIDS(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("角色菜单查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "角色菜单查询失败", vhttp.WithReason(err))
	}

	// 查询角色是否关联了菜单
	if len(all) > 0 {
		var codes []string
		for _, v := range all {
			codes = append(codes, strings.Split(v, ",")...)
		}

		for _, v := range codes {
			if fmt.Sprintf("%v", one.Id) == v {
				return vhttp.NewError(http.StatusBadRequest, "当前菜单已有角色关联，不能删除！")
			}
		}
	}

	err = t.mp.DeleteMenu(ctx, req.ID)
	if err != nil {
		t.log.WithContext(ctx).Errorf("删除菜单失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "删除菜单失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *MenuUseCase) GetMenu(ctx context.Context, req *dto.GetMenuReq) (*models.Menu, error) {
	res, err := t.mp.GetMenu(ctx, req.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, vhttp.NewError(http.StatusBadRequest, "数据不存在", vhttp.WithReason(err))
		}
		t.log.WithContext(ctx).Errorf("获取菜单失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
	}
	return res, nil
}



func (t *MenuUseCase) GetSelectMenu(ctx context.Context, pid int32) (all []*SelectMenu, err error) {
	all = append(all, &SelectMenu{
		Name:  "顶级菜单",
		Value: 0,
		Selected: func() bool {
			return  pid == 0
		}(),
	})

	all = append(all, t.getSelectMenuList(ctx, pid)...)

	return all, nil
}

func (t *MenuUseCase) GetSelectRoleMenu(ctx context.Context, id int32) (all []*SelectMenu, err error) {

	all = append(all, t.getSelectMenuList(ctx, id)...)

	return all, nil
}

func (t *MenuUseCase) GetRoleMenu(ctx context.Context, rid int32) (all []*SelectMenu, err error) {
	menus, err := t.mp.ListMenu(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取菜单列表失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "菜单获取失败", vhttp.WithReason(err))
	}

	var rids []string
	// 获取role关联的菜单
	role, err := t.rp.GetRole(ctx, rid)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.WithContext(ctx).Errorf("获取角色失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取角色失败", vhttp.WithReason(err))
	}
	if role != nil {
		rids = strings.Split(role.MenuIds, ",")
	}

	return t.buildRoleMenu(0, rids, menus), nil
}

// getSelectMenuList 获取菜单列表
func (t *MenuUseCase) getSelectMenuList(ctx context.Context, pid int32) []*SelectMenu {
	all, err := t.mp.ListMenu(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取菜单列表失败,%v", err)
		return nil
	}
	return t.buildSelectMenu(0, pid, all)
}

// buildMenuChild 递归获取子菜单
func (t *MenuUseCase) buildSysMenuChild(pid int32, menuList []*models.Menu) []*SysMenuList {
	var treeList []*SysMenuList
	for _, v := range menuList {
		if pid == v.Pid && v.Type == 1 {
			node := &SysMenuList{
				ID:    v.Id,
				Title: v.Title,
				Icon:  v.Icon,
				Href: func() string {
					if v.Href == "" {
						return ""
					}
					if strings.Contains(v.Href, "http") {
						return v.Href
					}
					return v.Href
				}(),
				Target: v.Target,
				Pid:    v.Pid,
			}
			child := t.buildSysMenuChild(v.Id, menuList)
			if len(child) != 0 {
				node.Child = child
			}
			// todo 后续此处加上用户的权限判断
			treeList = append(treeList, node)
		}
	}
	return treeList
}

// buildSelectMenu 递归获取子菜单
func (t *MenuUseCase) buildSelectMenu(pid int32, selctid int32, menuList []*models.Menu) []*SelectMenu {
	var treeList []*SelectMenu
	for _, v := range menuList {
		if pid == v.Pid {
			node := &SelectMenu{
				Name:  v.Title,
				Value: v.Id,
				Selected: func() bool {
					return v.Id == selctid
				}(),
			}
			child := t.buildSelectMenu(v.Id, selctid, menuList)
			if len(child) != 0 {
				node.Children = child
			}
			treeList = append(treeList, node)
		}
	}
	return treeList
}

// buildSelectRoleMenu 递归获取子菜单
func (t *MenuUseCase) buildRoleMenu(pid int32, selctids []string, menuList []*models.Menu) []*SelectMenu {
	var treeList []*SelectMenu
	for _, v := range menuList {
		if pid == v.Pid {
			node := &SelectMenu{
				Name:  v.Title,
				Value: v.Id,
				Selected: func() bool {
					for _, sid := range selctids {
						if strconv.Itoa(int(v.Id)) == sid {
							return true
						}
					}
					return false
				}(),
			}
			child := t.buildRoleMenu(v.Id, selctids, menuList)
			if len(child) != 0 {
				node.Children = child
			}
			treeList = append(treeList, node)
		}
	}
	return treeList
}

func (t *MenuUseCase) ListMenuTree(ctx context.Context) (all []*ListTreeMenu, err error) {
	res, err := t.mp.ListMenu(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取菜单列表失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取菜单列表失败", vhttp.WithReason(err))
	}

	return t.buildListMenuTree(0, res), nil
}

// buildSelectMenu 递归获取子菜单
func (t *MenuUseCase) buildListMenuTree(pid int32, menuList []*models.Menu) []*ListTreeMenu {
	var treeList []*ListTreeMenu
	for _, v := range menuList {
		if pid == v.Pid {
			node := &ListTreeMenu{
				Menu: *v,
			}
			child := t.buildListMenuTree(v.Id, menuList)
			if len(child) != 0 {
				node.Children = child
				node.IsParent = true
			}
			treeList = append(treeList, node)
		}
	}
	return treeList
}

// 检查新设置的父级是否在当前菜单的子菜单中
func (t *MenuUseCase) isParentInSubmenu(menuList []*models.Menu, parentID int32, currentMenuID int32) bool {
	for _, menu := range menuList {
		if menu.Pid == currentMenuID {
			// 如果找到当前菜单的子菜单
			if menu.Id == parentID {
				// 如果新设置的父级在当前菜单的子菜单中，则返回 true
				return true
			} else {
				// 否则递归检查当前菜单的子菜单的子菜单
				if t.isParentInSubmenu(menuList, parentID, menu.Id) {
					return true
				}
			}
		}
	}
	return false
}
