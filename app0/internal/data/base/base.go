package base

// 关于数据连接的一些考虑
// 一个应用在启动的时候，需要连接到哪些数据库肯定是已知的，如果是多库的情况，应该为每个库都维护一个自己相关的连接
// data层不应该由biz层直接传入，而且data层的操作也是一样，应该具体且明确的表示
import (
	"context"
	"github.com/aveyuan/vbasedata"
	"vmapp/app0/internal/biz/repo"
	"vmapp/app0/internal/conf"
	"vmapp/app0/internal/models"

	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	contextTxKey struct{}
)

type Data struct {
	Mysql   *gorm.DB
	Redis   redis.UniversalClient
	cleanup []func()
}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.Mysql.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx.WithContext(ctx)
	}
	return d.Mysql.WithContext(ctx)
}

// NewTransaction 实现接口TX
func NewTransaction(d *Data) repo.Transaction {
	return d
}

// NewData 初始化数据
func NewData(c *conf.AppConf, bc *conf.BootComponent) (*Data, func(), error) {
	data := &Data{}
	// 程序退出后，资源释放
	cleanup := func() {
		bc.Logger.Info("清理连接资源")
		for _, v := range data.cleanup {
			v()
		}
	}
	if c.Data == nil {
		return data, cleanup, nil
	}

	if c.Data.DB != nil {
		g, f, err := vbasedata.NewGorm(c.Data.DB, bc.Logger)
		if err != nil {
			panic(err)
		}
		data.cleanup = append(data.cleanup, f)
		data.Mysql = g
		g.AutoMigrate(new(models.User))
	}

	if c.Data.Redis != nil {
		r, f, err := vbasedata.NewRedis(c.Data.Redis, bc.Logger)
		if err != nil {
			bc.Logger.Errorf("Redis 连接失败,%v", err)
			panic(err)
		}
		data.cleanup = append(data.cleanup, f)
		data.Redis = r
	}

	return data, cleanup, nil
}
