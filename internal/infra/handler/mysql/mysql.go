package mysql

import (
	"fmt"
	"github.com/jinfeijie/pangu/internal/constant"
	"github.com/jinfeijie/pangu/internal/infra/handler"
	"github.com/jinfeijie/pangu/internal/infra/handler/configcenter"
	m "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

const configName = "mysql"

func init() {
	handler.Register(configName, newMysql(), constant.GroupBefore)
}

type mysql struct{}

func newMysql() *mysql {
	return &mysql{}
}

var dbmap sync.Map

func init() {
	dbmap = sync.Map{}
}

func (h *mysql) Serve() {
	for key, conf := range configcenter.GetConfig().Mysql {
		l := logger.New(lll{}, logger.Config{
			SlowThreshold:             1 * time.Millisecond,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Silent,
		})
		db, err := gorm.Open(m.New(m.Config{
			DSN: fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
				conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset,
			), // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  false, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			Logger: l,
		})

		if err != nil {
			panic(err)
		}

		d, _ := db.DB()
		err = d.Ping()
		if err != nil {
			panic(err)
		}
		dbmap.Store(key, db.Debug())
	}

}

func (h *mysql) Name() string {
	return configName
}

func GetDB(key ...string) *gorm.DB {
	if len(key) == 0 {
		db, exist := dbmap.Load("default")
		if exist {
			return db.(*gorm.DB)
		}
		return nil
	}

	if len(key) == 1 {
		db, exist := dbmap.Load(key[0])
		if exist {
			return db.(*gorm.DB)
		}
		return nil
	}
	return nil
}

type lll struct {
}

func (lll) Printf(s string, p ...interface{}) {
	log.Printf(s, p...)
}
