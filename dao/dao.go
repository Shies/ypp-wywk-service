package dao

import (
	"database/sql"
	"time"

	"ypp-wywk-service/conf"
	"ypp-wywk-service/net"

	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

var (
	noRow = sql.ErrNoRows
)

type Dao struct {
	c 	     *conf.Config
	db       *sql.DB
	dbact    *sql.DB
	dbreport *sql.DB
	redis 	 redis.Conn
	http	 *net.Client
	err 	 error
}

func New(c *conf.Config) *Dao {
	dao := &Dao{
		c: c,
		err: nil,
		db: nil,
		redis: getRedisConn(c.Redis),
		http: net.NewClient(c.HttpClient),
	}
	dao.db, dao.err = getConn(c.Db)
	if dao.err != nil {
		panic(dao.err)
	}

	dao.dbact, dao.err = getConn(c.DbAct)
	if dao.err != nil {
		panic(dao.err)
	}

	dao.dbreport, dao.err = getConn(c.DbReport)
	if dao.err != nil {
		panic(dao.err)
	}

	return dao
}

func getRedisConn(r *conf.Redis) (redis.Conn) {
	conn, err := redis.Dial(r.Proto, r.Host+":"+r.Port)
	if err != nil {
		panic(err)
	}
	if _, err := conn.Do("AUTH", r.Auth); err != nil {
		conn.Close()
		panic(err)
	}
	if _, err := conn.Do("SELECT", 0); err != nil {
		conn.Close()
		panic(err)
	}

	return conn
}

func getConn(db *conf.Db) (*sql.DB, error) {
	connStr := createConnStr(db.Username, db.Password, db.Addr, db.Port, db.Db_name)
	DB, err := sql.Open("mysql", connStr)
	DB.SetConnMaxLifetime(time.Second)
	return DB, err
}

func createConnStr(username string, password string, addr string, port string, db_name string) string {
	return username + ":" + password + "@tcp(" + addr + ":" + port + ")/" + db_name + "?charset=utf8&parseTime=true&loc=Local"
}

func Now(t time.Time) int64 {
	return t.Unix()
}

