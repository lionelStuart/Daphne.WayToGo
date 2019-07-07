package DB

import (
	"database/sql"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"strconv"
	"sync"
)

type DBConf struct {
	Title string
	Mysql MySqlInfo
}

type MySqlInfo struct {
	Host     string
	Port     uint
	User     string
	Password string
	DB       string
}

var (
	dbConf *DBConf
	DB     *sql.DB
	err    error
)

func GetDBConf(path string) *DBConf {
	dbConf := &DBConf{}
	var once sync.Once
	once.Do(func() {
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(path, &dbConf); err != nil {
			panic(err)
		}
	})
	return dbConf
}

func init() {
	dbConf = GetDBConf("db.toml")
	DB, err = sql.Open("mysql", dbConf.Mysql.User+
		":"+dbConf.Mysql.Password+"@tcp("+
		dbConf.Mysql.Host+":"+string(dbConf.Mysql.Port)+")/"+
		dbConf.Mysql.DB+"?charset=utf8")
}

func NewConn() (*sql.DB, error) {
	DB, err := sql.Open("mysql", dbConf.Mysql.User+
		":"+dbConf.Mysql.Password+"@tcp("+
		dbConf.Mysql.Host+":"+strconv.FormatUint(uint64(dbConf.Mysql.Port), 10)+")/"+
		dbConf.Mysql.DB+"?charset=utf8")
	if err != nil {
		panic(err)
		return nil, err
	}
	return DB, nil
}
