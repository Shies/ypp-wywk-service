package dao

import (
	"ypp-wywk-service/model"
	"log"
)

const (
	_userByOne = "SELECT id, create_time FROM t_user WHERE mobile = ? LIMIT 1"
)

func (d *Dao) UserExists(mobilephone string) (user *model.TUser, err error) {
	row := d.db.QueryRow(_userByOne, mobilephone)
	user = &model.TUser{}
	if err = row.Scan(&user.ID, &user.CreateTime); err != nil {
		if err == noRow {
			err = nil
			user = nil
		} else {
			log.Println("DB query failed:", err.Error())
		}
	}
	return
}