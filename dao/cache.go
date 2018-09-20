package dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

const (
	_wywkUserKey = "WYWK_USER_%s"
	_wywkCheckYppSend = "WYWK_CHECK_YPP_SEND"
	_wywkCheckReserve = "RESERVE_SEAT_HASH_TABLE"
)

func (d *Dao) UserCacheKey(mobilephone string) string {
	return fmt.Sprintf(_wywkUserKey, mobilephone)
}

func (d *Dao) CheckSendCacheKey() string {
	return _wywkCheckYppSend
}

func (d *Dao) CheckReserveCacheKey() string {
	return _wywkCheckReserve
}

func (d *Dao) SetUserCache(mobile string, val string) (reply bool) {
	var (
		conn = d.redis
		key = d.UserCacheKey(mobile)
	)

	defer conn.Close()
	v, err := redis.Bool(conn.Do("SET", key, val))
	if err != nil {
		fmt.Println(err)
		return
	}

	return v
}

func (d *Dao) CheckSendCache(mobile string) (reply bool) {
	var (
		conn = d.redis
		key = d.CheckSendCacheKey()
	)

	defer conn.Close()
	isExist, err := redis.Bool(conn.Do("hexists", key, mobile))
	if err != nil {
		fmt.Println(err)
		return
	}

	if !isExist {
		_, err := redis.Bool(conn.Do("hset", key, mobile, 1))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	return isExist
}

func (d *Dao) CheckReserveCache(cardno string) (reply bool) {
	var (
		conn = d.redis
		key = d.CheckReserveCacheKey()
	)

	defer conn.Close()
	v, err := redis.Bool(conn.Do("hexists", key, cardno))
	if err != nil {
		fmt.Println(err)
		return
	}

	return v
}