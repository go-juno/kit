package sql

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/xerrors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlClient interface {
	// Updates 指定单个字段或者多个字段
	Updates(model interface{}, values map[string]interface{}) (bool, error)
	// Save 全更新
	Save(values interface{}) (bool, error)
	// UpDateBySelect 更新选定字段
	UpDateBySelect(model interface{}, queries []string, values interface{}) (bool, error)
	Create(value interface{}) *gorm.DB
	Del(whereQuery string, value interface{}) (bool, error)
}

// Config 配置
type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
	Port     int    `yaml:"port"`
}

// NewMysqlDBClient  连接数据库
func NewMysqlDBClient(config Config) (MysqlClient, error) {
	sqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
	dia := mysql.Open(sqlURI)

	logConfig := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logConfig,
	)
	db, err := gorm.Open(dia, &gorm.Config{Logger: newLogger, DisableAutomaticPing: true})
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return nil, err
	}

	return &mysqlClient{
		db: db,
	}, nil
}

type mysqlClient struct {
	db *gorm.DB
}

func (m mysqlClient) Updates(model interface{}, values map[string]interface{}) (bool, error) {
	err := m.db.Model(model).Updates(values).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return true, nil
}

func (m mysqlClient) UpDateBySelect(model interface{}, queries []string, values interface{}) (bool, error) {
	query := ""
	queryLen := len(queries) - 1
	for index, quer := range queries {
		if index < queryLen {
			query = query + quer + ","
		} else {
			query = query + quer
		}
	}
	err := m.db.Model(model).Select(query).Updates(values).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return true, nil
}

func (m mysqlClient) Save(values interface{}) (bool, error) {
	err := m.db.Save(&values).Error
	if err != nil {
		err = xerrors.Errorf("%w", err)
		return false, err
	}
	return true, nil
}

func (m mysqlClient) Create(value interface{}) *gorm.DB {
	return m.db.Create(value)

}

func (m mysqlClient) Del(whereQuery string, value interface{}) (bool, error) {
	if whereQuery != "" {
		err := m.db.Where(whereQuery).Delete(&value).Error
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return false, err
		}
	} else {
		err := m.db.Delete(&value).Error
		if err != nil {
			err = xerrors.Errorf("%w", err)
			return false, err
		}
	}
	return true, nil
}
