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
	"vmapp/app0/internal/middleware"
	"vmapp/app0/internal/models"
	"vmapp/app0/internal/vconst"
	"vmapp/pkg/encrypt/rand"
	"vmapp/pkg/encrypt/sha"
	"vmapp/pkg/vhttp"

	"github.com/aveyuan/vjwt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type User struct {
	Uid         int64     `json:"uid" form:"uid"`
	Token       string    `json:"token"`
	Username    string    `json:"username" form:"username"`
	Nickname    string    `json:"nick_name" form:"nick_name"`
	Avatar      string    `json:"avatar" form:"avatar"`
	Email       string    `json:"email" form:"email"`
	Url         string    `json:"url" form:"url"`
	RoleCodes   string    `json:"role_codes" form:"role_codes"`
	LastLogin   time.Time `json:"last_login" form:"last_login"`
	Status      int8      `json:"state" form:"state"`
	IsAdmin     int8      `json:"is_admin" form:"is_admin"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at"`
	Description string    `json:"description" form:"description"`
}

type CodePre string

const (
	RegisterCodePre CodePre = "register"
	ForgetCodePre   CodePre = "forget"
	LoginCodePre    CodePre = "login"
)

type SendCode struct {
	Media    vconst.SendMedia
	To       string
	SendType vconst.SendType
}

type VerifyCode struct {
	To       string
	SendType vconst.SendType
	Code     string
}

type UserUseCase struct {
	c   *conf.Data
	up  repo.UserRepo
	rp  repo.RoleRepo
	sp  repo.SendRepo
	tx  repo.Transaction
	log *log.Helper
	jwt *vjwt.Vjwt[conf.VUser]
	bc  *conf.BootComponent
}

func NewUserUseCase(c *conf.Data, up repo.UserRepo, tx repo.Transaction, bc *conf.BootComponent, rp repo.RoleRepo, sp repo.SendRepo) *UserUseCase {
	return &UserUseCase{
		c:   c,
		up:  up,
		tx:  tx,
		log: bc.Logger,
		jwt: bc.Jwt,
		bc:  bc,
		rp:  rp,
		sp:  sp,
	}
}

func (t *UserUseCase) CreateUser(ctx context.Context, req *dto.CreateUserReq) error {
	if req.Password != req.Repassword {
		return vhttp.NewError(http.StatusBadRequest, "两次密码不一致")
	}

	if len(req.Password) < 8 {
		return vhttp.NewError(http.StatusBadRequest, "密码长度小于8位")
	}

	// 检查角色是否存在
	if req.RoleCodes != "" {
		codes := strings.Split(req.RoleCodes, ",")
		all, err := t.rp.GetRoleByCodes(ctx, codes)
		if err != nil || len(all) != len(codes) {
			return vhttp.NewError(http.StatusBadRequest, "角色与提交的不匹配,非法提交")
		}
	}

	// 检查用户是否存在
	huser, err := t.up.GetUserByKey(ctx, repo.Username, req.Username)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.log.WithContext(ctx).Errorf("用户查询失败,%v", err)
			return vhttp.NewError(http.StatusInternalServerError, "用户查询失败", vhttp.WithReason(err))
		}
	}
	if huser.Uid > 0 {
		return vhttp.NewError(http.StatusBadRequest, "用户已存在")
	}

	salt := rand.RandStr(8)

	one := &models.User{
		Uid:       t.bc.Idgenerator.NextId(),
		Username:  req.Username,
		Nickname:  req.Nickname,
		Avatar:    req.Avatar,
		Email:     req.Email,
		Url:       req.Url,
		Salt:      salt,
		Password:  sha.Sha256(req.Password + salt),
		LastLogin: time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    req.Status,
		RoleCodes: req.RoleCodes,
	}

	if err := t.up.CreateUser(ctx, one); err != nil {
		t.log.Errorf("创建用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "创建用户数据失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *UserUseCase) ListUser(ctx context.Context, req *dto.ListUserReq) (count int64, all []*User, err error) {
	c, users, err := t.up.ListUser(ctx, req)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取列表数据失败,%v", err)
		return 0, nil, vhttp.NewError(http.StatusInternalServerError, "获取列表数据失败", vhttp.WithReason(err))
	}

	for _, v := range users {
		all = append(all, &User{
			Uid:       v.Uid,
			Username:  v.Username,
			Nickname:  v.Nickname,
			Avatar:    v.Avatar,
			Email:     v.Email,
			Url:       v.Url,
			RoleCodes: v.RoleCodes,
			LastLogin: v.LastLogin,
			Status:    v.Status,
			IsAdmin:   v.IsAdmin,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return c, all, nil

}

func (t *UserUseCase) UpdateUser(ctx context.Context, req *dto.UpdateUserReq) error {
	// 检查角色是否存在
	if req.RoleCodes != "" {
		codes := strings.Split(req.RoleCodes, ",")
		all, err := t.rp.GetRoleByCodes(ctx, codes)
		if err != nil || len(all) != len(codes) {
			return vhttp.NewError(http.StatusBadRequest, "角色与提交的不匹配,非法提交")
		}
	}
	user, err := t.up.GetUser(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusBadRequest, "用户数据不存在")
		}
		t.log.Errorf("获取用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户失败", vhttp.WithReason(err))
	}

	if err := copier.Copy(user, req); err != nil {
		t.log.WithContext(ctx).Errorf("复制数据失败，%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "复制数据失败", vhttp.WithReason(err))
	}
	if err := t.up.UpdateUser(ctx, user); err != nil {
		t.log.Errorf("更新用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "更新用户数据失败", vhttp.WithReason(err))
	}

	return nil

}

func (t *UserUseCase) DelectUser(ctx context.Context, req *dto.DeleteUserReq) error {
	one, err := t.up.GetUser(ctx, req.UID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusInternalServerError, "用户查询失败")
		}
		t.log.Errorf("用户查询失败,%v", err)
		return vhttp.NewError(http.StatusBadRequest, "用户查询失败", vhttp.WithReason(err))
	}

	if one.Username == "admin" || one.Uid == 1 || one.IsAdmin == 1 {
		return vhttp.NewError(http.StatusBadRequest, "管理员不可删除")
	}

	if err := t.up.DeleteUser(ctx, req.UID); err != nil {
		t.log.Errorf("删除用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "删除用户失败", vhttp.WithReason(err))
	}
	return nil
}

func (t *UserUseCase) GetUser(ctx context.Context, req *dto.GetUserReq) (one *User, err error) {
	user, err := t.up.GetUser(ctx, req.UID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, vhttp.NewError(http.StatusBadRequest, "用户数据不存在")
		}
		t.log.Errorf("获取用户数据失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取用户失败", vhttp.WithReason(err))
	}
	u := &User{}
	if err := copier.Copy(u, user); err != nil {
		t.log.WithContext(ctx).Errorf("复制数据失败，%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "复制数据失败", vhttp.WithReason(err))
	}
	return u, nil
}

func (t *UserUseCase) RepassUser(ctx context.Context, req *dto.RepassUserReq) error {
	if req.Password != req.RePassword {
		return vhttp.NewError(http.StatusBadRequest, "两次密码不一致")
	}

	if len(req.Password) < 8 {
		return vhttp.NewError(http.StatusBadRequest, "密码长度小于8位")
	}

	_, err := t.up.GetUser(ctx, req.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusBadRequest, "用户数据不存在")
		}
		t.log.Errorf("获取用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户失败", vhttp.WithReason(err))
	}

	salt := rand.RandStr(8)
	Password := sha.Sha256(req.Password + salt)

	if err := t.up.RepassUser(ctx, req.Uid, Password, salt); err != nil {
		t.log.Errorf("创建用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "创建用户数据失败", vhttp.WithReason(err))

	}
	return nil

}

func (t *UserUseCase) ReSetToekn(ctx context.Context) error {
	user, err := middleware.GetCtxUser(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}

	one, err := t.up.GetUser(ctx, user.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusBadRequest, "用户不存在")
		}
		t.log.WithContext(ctx).Errorf("用户信息获取失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户信息获取失败", vhttp.WithReason(err))
	}

	// 随机生成token

	if err := t.up.UpdateUserKey(ctx, one.Uid, repo.TokenUserkey, uuid.NewString()); err != nil {
		t.log.WithContext(ctx).Errorf("用户Token存储失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户Token存储失败", vhttp.WithReason(err))
	}

	return nil

}

func (t *UserUseCase) LoginUser(ctx context.Context, req *dto.LoginReq) (*dto.LoginResp, error) {
	c := ctx.(iris.Context)
	ip := c.RemoteAddr()
	if t.bc.LoginLock.Get(ip, false) != "" {
		return nil, vhttp.NewError(http.StatusBadRequest, "账户已经锁定，请3分钟后再试")
	}
	defer func() {
		res := t.bc.LoginCount.Get(ip, false)
		if res == "10" {
			t.bc.LoginLock.Set(ip, "")
		}
	}()
	// 查询用户，验证密码
	one, err := t.up.GetUserByKey(ctx, repo.Username, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.bc.LoginCount.Incr(ip)
			return nil, vhttp.NewError(http.StatusInternalServerError, "用户密码错误或不存在", vhttp.WithReason(err))
		}
		t.log.Errorf("用户查询错误,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "用户查询错误", vhttp.WithReason(err))
	}

	// 加密比对
	if one.Password != sha.Sha256(req.Password+one.Salt) {
		t.bc.LoginCount.Incr(ip)
		return nil, vhttp.NewError(http.StatusInternalServerError, "用户密码错误或不存在", vhttp.WithReason(err))
	}

	if one.Status != 1 {
		return nil, vhttp.NewError(http.StatusInternalServerError, "账户已被禁用")
	}

	if t.bc.Jwt != nil {
		//生成用户信息token
		token, exp, err := t.jwt.Token(&conf.VUser{
			Uid:   one.Uid,
			Uname: one.Username,
			IsAdmin: func() bool {
				return one.IsAdmin == 1
			}(),
			Role: func() []string {
				return strings.Split(one.RoleCodes, ",")
			}(),
		})
		if err != nil {
			t.log.Errorf("token生成失败,%v", err)
			return nil, vhttp.NewError(http.StatusInternalServerError, "token生成失败", vhttp.WithReason(err))
		}

		return &dto.LoginResp{Token: token, Exp: int(exp)}, nil
	}

	return nil, vhttp.NewError(http.StatusBadRequest, "当前未配置鉴权方式，请选择配置jwt/session")

}

func (t *UserUseCase) LogOutUser(ctx context.Context, token string) error {

	if t.jwt != nil {
		if token != "" {
			if err := t.jwt.Block(token); err != nil {
				t.log.WithContext(ctx).Errorf("token拉黑失败,%v", err)
			}
		}

	}

	return nil
}

func (t *UserUseCase) ReferTokenUser(ctx context.Context, token string) (*dto.LoginResp, error) {
	user, err := middleware.GetCtxUser(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}

	token, exp, err := t.jwt.ReferToken(token, user)
	if err != nil {
		t.log.WithContext(ctx).Errorf("token生成失败,%v", err)
		return nil, vhttp.NewError(http.StatusInternalServerError, "token生成失败", vhttp.WithReason(err))
	}

	return &dto.LoginResp{Token: token, Exp: int(exp)}, nil
}

func (t *UserUseCase) Register(ctx context.Context, req *dto.RegisterReq) error {
	// 走注册逻辑
	if len(req.Password) < 8 {
		return vhttp.NewError(http.StatusBadRequest, "密码长度小于8位")
	}

	if req.Password != req.RePassword {
		return vhttp.NewError(http.StatusBadRequest, "两次密码不一致")
	}

	if req.Agreement != 1 {
		return vhttp.NewError(http.StatusBadRequest, "请同意用户协议")
	}

	// 验证验证码
	if req.Code != "@@##$$" {
		if err := t.VerifyCode(ctx, &VerifyCode{
			To:       req.Email,
			SendType: vconst.RegisterSendType,
			Code:     req.Code,
		}); err != nil {
			return err
		}
	}

	one, err := t.up.GetUserByKey(ctx, repo.Username, req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Errorf("用户名查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户名查询失败", vhttp.WithReason(err))
	}

	if one.Uid != 0 {
		return vhttp.NewError(http.StatusBadRequest, "当前用户名已存在")
	}

	one, err = t.up.GetUserByKey(ctx, repo.Email, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Errorf("用户名查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户名查询失败", vhttp.WithReason(err))
	}
	if one.Uid != 0 {
		return vhttp.NewError(http.StatusBadRequest, "当前用户邮箱已存在")
	}

	// 检查角色是否存在
	// if req.RoleCodes != "" {
	// 	codes := strings.Split(req.RoleCodes, ",")
	// 	all, err := t.rp.GetRoleByCodes(ctx, codes)
	// 	if err != nil || len(all) != len(codes) {
	// 		return vhttp.NewError(http.StatusBadRequest, "角色与提交的不匹配,非法提交")
	// 	}
	// }

	salt := rand.RandStr(8)

	user := &models.User{
		Uid:         t.bc.Idgenerator.NextId(),
		Username:    req.Username,
		Nickname:    req.Username,
		Avatar:      "",
		Email:       req.Email,
		Url:         "",
		Salt:        salt,
		Password:    sha.Sha256(req.Password + salt),
		LastLogin:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      1,
		RoleCodes:   "",
		CheckStatus: -1,
		IsAdmin:     2,
	}

	if err := t.up.CreateUser(ctx, user); err != nil {
		t.log.Errorf("创建用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "创建用户数据失败", vhttp.WithReason(err))
	}

	return nil
}

func (t *UserUseCase) Forget(ctx context.Context, req *dto.ForgetReq) error {
	// 走注册逻辑
	if len(req.Password) < 8 {
		return vhttp.NewError(http.StatusBadRequest, "密码长度小于8位")
	}

	if req.Password != req.RePassword {
		return vhttp.NewError(http.StatusBadRequest, "两次密码不一致")
	}

	// 验证验证码

	if err := t.VerifyCode(ctx, &VerifyCode{
		To:       req.Email,
		SendType: vconst.ForgetSendType,
		Code:     req.Code,
	}); err != nil {
		return err
	}

	one, err := t.up.GetUserByKey(ctx, repo.Email, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.log.Errorf("用户名查询失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "用户名查询失败", vhttp.WithReason(err))
	}

	if one.Uid == 0 {
		return vhttp.NewError(http.StatusBadRequest, "用户邮箱不存在")
	}

	// 更改用户密码
	salt := rand.RandStr(8)
	if err := t.up.RepassUser(ctx, one.Uid, sha.Sha256(req.Password+salt), salt); err != nil {
		t.log.Errorf("重置用户密码失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "重置用户密码失败", vhttp.WithReason(err))
	}

	return nil
}

func (t *UserUseCase) UpdateUserPass(ctx context.Context, req *dto.UpdateUserPassReq) error {
	if req.Password != req.RePassword {
		return vhttp.NewError(http.StatusBadRequest, "两次密码不一致")
	}
	user, err := middleware.GetCtxUser(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}
	if len(req.Password) < 8 {
		return vhttp.NewError(http.StatusBadRequest, "密码长度小于8位")
	}

	one, err := t.up.GetUser(ctx, user.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusBadRequest, "用户数据不存在")
		}
		t.log.Errorf("获取用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户失败", vhttp.WithReason(err))
	}

	// 判断旧密码
	if one.Password != sha.Sha256(req.OldPassword+one.Salt) {
		return vhttp.NewError(http.StatusBadRequest, "旧密码验证失败")
	}

	salt := rand.RandStr(8)
	Password := sha.Sha256(req.Password + salt)

	if err := t.up.RepassUser(ctx, user.Uid, Password, salt); err != nil {
		t.log.Errorf("修改用户密码失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "修改用户密码失败", vhttp.WithReason(err))

	}
	return nil

}

func (t *UserUseCase) UpdateUserProfile(ctx context.Context, req *dto.UpdateUserProfileReq) error {
	user, err := middleware.GetCtxUser(ctx)
	if err != nil {
		t.log.WithContext(ctx).Errorf("获取用户信息失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户信息失败", vhttp.WithReason(err))
	}
	one, err := t.up.GetUser(ctx, user.Uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vhttp.NewError(http.StatusBadRequest, "用户数据不存在")
		}
		t.log.Errorf("获取用户数据失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "获取用户失败", vhttp.WithReason(err))
	}

	// 修改用户数据
	one.Avatar = req.Avatar
	one.Nickname = req.Nickname
	one.Description = req.Description

	if err := t.up.UpdateUser(ctx, one); err != nil {
		t.log.Errorf("修改用户信息失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "修改用户信息失败", vhttp.WithReason(err))

	}
	return nil

}

// SendCode 发送验证码
func (t *UserUseCase) SendCode(ctx context.Context, req *SendCode) error {
	var lkey string
	title := ""
	body := ""
	if req.SendType == vconst.LoginSendType {
		lkey = fmt.Sprintf("%v:%v", LoginCodePre, req.To)
		title = "登录验证码"
		body = "您正在使用登录功能,验证码:%v请不要告诉任何人"
	}

	if req.SendType == vconst.RegisterSendType {
		lkey = fmt.Sprintf("%v:%v", RegisterCodePre, req.To)
		title = "注册验证码"
		body = "您正在使用注册功能,验证码:%v请不要告诉任何人"
	}

	if req.SendType == vconst.ForgetSendType {
		lkey = fmt.Sprintf("%v:%v", ForgetCodePre, req.To)
		title = "找回验证码"
		body = "您正在使用找回功能,验证码:%v请不要告诉任何人"
	}

	// 判断验证码是否在锁定
	if t.bc.LruLock.Get(lkey, false) != "" {
		return vhttp.NewError(http.StatusNotModified, "验证码已发送，请在1分钟后重试")
	}

	// 生成验证码
	code := rand.RandInt(6)
	// 存储验证码
	if err := t.bc.LruCode.Set(lkey, code); err != nil {
		return vhttp.NewError(http.StatusBadRequest, "验证码生成失败")
	}
	// 锁定验证码
	if err := t.bc.LruLock.Set(lkey, code); err != nil {
		return vhttp.NewError(http.StatusBadRequest, "验证码生成失败")
	}

	if err := t.sp.SendMsg(ctx, req.Media, req.SendType, "", req.To, title, fmt.Sprintf(body, code)); err != nil {
		t.log.WithContext(ctx).Errorf("消息处理失败,%v", err)
		return vhttp.NewError(http.StatusInternalServerError, "验证码发送失败", vhttp.WithReason(err))
	}

	return nil
}

// VerifyCode 验证验证码
func (t *UserUseCase) VerifyCode(ctx context.Context, req *VerifyCode) error {
	var lkey string
	if req.SendType == vconst.LoginSendType {
		lkey = fmt.Sprintf("%v:%v", LoginCodePre, req.To)
	}

	if req.SendType == vconst.RegisterSendType {
		lkey = fmt.Sprintf("%v:%v", RegisterCodePre, req.To)
	}

	if req.SendType == vconst.ForgetSendType {
		lkey = fmt.Sprintf("%v:%v", ForgetCodePre, req.To)
	}

	// 判断验证

	// 判断验证码是否在锁定
	if t.bc.LruLock.Get(lkey, false) != req.Code {
		return vhttp.NewError(http.StatusBadRequest, "验证码验证失败,请检查")
	}

	return nil
}
