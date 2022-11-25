package template

var GlobalTmp = `package global

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	DefaultConfigDir  = "configs"
	DefaultConfigType = "yaml"
)

func NewDefaultConfig() *ViperConfig {
	vp := viper.New()
	vp.SetConfigType(DefaultConfigType)
	vp.AddConfigPath(DefaultConfigDir)
	vp.AddConfigPath(".")

	return &ViperConfig{
		vp,
	}
}

type ViperConfig struct {
	*viper.Viper
}

func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.Replace(s, ".", "_", -1))
}

func (c *ViperConfig) ReadData() error {
	return c.ReadInConfig()
}

func (c *ViperConfig) GetWithEnv(key string) interface{} {
	v := os.Getenv(FormatEnvKey(key))
	if len(v) > 0 {
		return v
	} else {
		return c.Get(key)
	}
}`

var GlobalConfigTmp = `package global

/*
	全局变量存放
*/

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var Config *ViperConfig
var RedisClient *redis.Client
var Log *logrus.Logger
var mysqlMap sync.Map
var MongoDBClient *mongo.Client

//进行项目的初始
func init() {
	//初始化日志对象
	if Log == nil {
		Log = logrus.New()
		Log.SetFormatter(&MyFormatter{})
	}

	//初始化配置文件
	if Config == nil {
		Config = NewDefaultConfig()
		if err := Config.ReadData(); err != nil {
			panic(err)
		}
	}

	//初始化基础库连接
	//if _, err := initMysqlDb(Config.GetString("Mysql.default.link")); err != nil {
	//	panic(err)
	//}
}`

var GlobalMysqlTmp = `package global

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//GetMysqlDB 获取默认数据库连接
func GetMysqlDB() (db *gorm.DB, err error) {
	dsn := Config.GetString("mysql.default.link")
	return GetMysqlDBByDSN(dsn)
}

//GetMysqlDBByDSN 获取数据库连接通过dsn
func GetMysqlDBByDSN(dsn string) (db *gorm.DB, err error) {
	value, ok := mysqlMap.Load(dsn)

	if !ok {
		if db, err = initMysqlDb(dsn); err != nil {
			return
		}
		return
	}

	db, _ = value.(*gorm.DB)

	var sqlDb *sql.DB
	if sqlDb, err = db.DB(); err != nil {
		if db, err = initMysqlDb(dsn); err != nil {
			return
		}
	}

	if err = sqlDb.Ping(); err != nil { //检查连接状态
		if db, err = initMysqlDb(dsn); err != nil {
			return
		}
	}

	return
}

func initMysqlDb(dsn string) (db *gorm.DB, err error) {
	if db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		//DefaultStringSize:       256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 change 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{}); err != nil {
		return
	}

	mysqlMap.Store(dsn, db)
	return
}`

var GlobalLoggerTmp = `package global

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
)

type MyFormatter struct{}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s] [%s]\t%s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}`

var GlobalMongoTmp = `package global

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDBClient() (c *mongo.Client, err error) {
	if MongoDBClient == nil {
		if err = initMongoDBClient(); err != nil {
			return
		}
	}

	if err = MongoDBClient.Ping(context.TODO(), nil); err != nil {
		if err = initMongoDBClient(); err != nil {
			return
		}
	}

	return MongoDBClient, nil
}

func initMongoDBClient() (err error) {
	mongodbUri := Config.Get("MongoDB.link").(string)
	clientOptions := options.Client().ApplyURI(mongodbUri)

	//连接
	var client *mongo.Client
	if client, err = mongo.Connect(context.TODO(), clientOptions); err != nil {
		return
	}

	MongoDBClient = client
	return
}

func GetMongoDB() (db *mongo.Database, err error) {
	var client *mongo.Client
	if client, err = GetMongoDBClient(); err != nil {
		return
	}
	database := Config.GetString("MongoDB.database")
	db = client.Database(database)
	return
}`

var GlobalRedisTmp = `package global

import "github.com/go-redis/redis"

func GetRedisClient() (c *redis.Client, err error) {
	if RedisClient == nil {
		RedisClient = redis.NewClient(&redis.Options{
			Addr: Config.GetString("Redis.link"),
			DB:   Config.GetInt("Redis.database"),
		})
	}

	if err = RedisClient.Ping().Err(); err != nil {
		return
	}
	c = RedisClient

	return
}`
