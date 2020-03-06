package ginApp

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Application struct {
	Addr   string
	Mode   string
	DB     *gorm.DB
	Engine *gin.Engine
	Cache  *redis.Client
}

func (a *Application) Init(cfg *Config) error {
	a.Mode = cfg.ServerConfig.Mode
	a.Addr = cfg.ServerConfig.ParseUrl()

	a.Engine = gin.Default()
	gin.SetMode(a.Mode)
	
	// 初始化数据库
	err := a.initDB(cfg)
	if err != nil{
		return err
	}

	// 初始化缓存
	err = a.initCache(cfg)
	if err != nil{
		return nil
	}
	return nil
}

func (a *Application) initDB(cfg *Config) error {
	var err error

	a.DB, err = gorm.Open(
		cfg.DBConfig.Dialect,
		cfg.DBConfig.ParseUri(),
		)

	if err != nil{
		return err
	}

	return nil
}

func (a *Application) initCache(cfg *Config)  error {
	a.Cache = redis.NewClient(&redis.Options{
		Addr:               cfg.RedisConfig.ParseUrl(),
		Password:           cfg.RedisConfig.Password,
		DB:                 cfg.RedisConfig.DB,
		MaxRetries:         0,
	})

	_, err := a.Cache.Ping().Result()
	if err != nil{
		return err
	}

	return nil
}

func (a *Application) Run() error {
	return a.Engine.Run(a.Addr)
}