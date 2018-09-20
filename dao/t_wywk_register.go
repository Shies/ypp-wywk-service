package dao

import (
	"log"
	"fmt"
	"ypp-wywk-service/model"
)

const (
	_wyByOne = "SELECT id, type_id, id_type, id_card, truename, create_time, status, active_time, mobile, card_no, card_create_time, mobile_update_time, send_count, last_send_date, common_code, update_timestamp FROM t_wy_register WHERE %s = ? LIMIT 1"
	_hmByOne = "SELECT id, type_id, id_type, id_card, truename, create_time, status, active_time, mobile, card_no, card_create_time, mobile_update_time, send_count, last_send_date, common_code, update_timestamp FROM t_hm_register WHERE %s = ? LIMIT 1"
	_registerInsert = "INSERT INTO t_wywk_register(data, create_time) VALUES(?, ?)"
	_registerWYInsert = "INSERT INTO t_wy_register(id, type_id, id_type, id_card, truename, create_time, status, active_time, mobile, card_no, card_create_time, mobile_update_time, send_count, last_send_date, common_code) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_registerHMInsert = "INSERT INTO t_hm_register(id, type_id, id_type, id_card, truename, create_time, status, active_time, mobile, card_no, card_create_time, mobile_update_time, send_count, last_send_date, common_code) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_registerWYUpdate = "UPDATE t_wy_register SET type_id = ?, id_type = ?, id_card = ?, truename = ?, create_time = ?, status = ?, active_time = ?, mobile = ?, card_no = ?, card_create_time = ?, mobile_update_time = ?, send_count = ?, last_send_date = ?, common_code = ? WHERE id = ?"
	_registerHMUpdate = "UPDATE t_hm_register SET type_id = ?, id_type = ?, id_card = ?, truename = ?, create_time = ?, status = ?, active_time = ?, mobile = ?, card_no = ?, card_create_time = ?, mobile_update_time = ?, send_count = ?, last_send_date = ?, common_code = ? WHERE id = ?"
)

func (d *Dao) WYExists(value string, key string) (reg *model.TWyRegister, err error) {
	sql := fmt.Sprintf(_wyByOne, key)
	row := d.db.QueryRow(sql, value)
	reg = &model.TWyRegister{}
	if err = row.Scan(&reg.ID, &reg.TypeID, &reg.IDType, &reg.IDCard, &reg.Truename, &reg.CreateTime, &reg.Status, &reg.ActiveTime, &reg.Mobile, &reg.CardNO, &reg.CardCreateTime, &reg.MobileUpdateTime, &reg.SendCount, &reg.LastSendDate, &reg.CommonCode, &reg.UpdateTimeStamp); err != nil {
		if err == noRow {
			err = nil
			reg = nil
		} else {
			log.Println("DB query failed:", err.Error())
		}
	}
	return
}

func (d *Dao) HMExists(value string, key string) (reg *model.TWyRegister, err error) {
	sql := fmt.Sprintf(_hmByOne, key)
	row := d.db.QueryRow(sql, value)
	reg = &model.TWyRegister{}
	if err = row.Scan(&reg.ID, &reg.TypeID, &reg.IDType, &reg.IDCard, &reg.Truename, &reg.CreateTime, &reg.Status, &reg.ActiveTime, &reg.Mobile, &reg.CardNO, &reg.CardCreateTime, &reg.MobileUpdateTime, &reg.SendCount, &reg.LastSendDate, &reg.CommonCode, &reg.UpdateTimeStamp); err != nil {
		if err == noRow {
			err = nil
			reg = nil
		} else {
			log.Println("DB query failed:", err.Error())
		}
	}
	return
}

func (d *Dao) AddWYWKRegister(r *model.TRegister) (err error) {
	sql, err := d.dbact.Prepare(_registerInsert)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.Data, r.CreateTime)
	return
}

func (d *Dao) AddWYRegister(r *model.TWyRegister) (err error) {
	sql, err := d.db.Prepare(_registerWYInsert)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.ID, r.TypeID, r.IDType, r.IDCard, r.Truename, r.CreateTime, r.Status, r.ActiveTime, r.Mobile, r.CardNO, r.CardCreateTime, r.MobileUpdateTime, r.SendCount, r.LastSendDate, r.CommonCode)
	return
}

func (d *Dao) AddHMRegister(r *model.TWyRegister) (err error) {
	sql, err := d.db.Prepare(_registerHMInsert)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.ID, r.TypeID, r.IDType, r.IDCard, r.Truename, r.CreateTime, r.Status, r.ActiveTime, r.Mobile, r.CardNO, r.CardCreateTime, r.MobileUpdateTime, r.SendCount, r.LastSendDate, r.CommonCode)
	return
}

func (d *Dao) UpdateHMRegister(r *model.TWyRegister) (err error) {
	sql, err := d.db.Prepare(_registerHMUpdate)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.TypeID, r.IDType, r.IDCard, r.Truename, r.CreateTime, r.Status, r.ActiveTime, r.Mobile, r.CardNO, r.CardCreateTime, r.MobileUpdateTime, r.SendCount, r.LastSendDate, r.CommonCode, r.ID)
	return
}

func (d *Dao) UpdateWYRegister(r *model.TWyRegister) (err error) {
	sql, err := d.db.Prepare(_registerWYUpdate)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.TypeID, r.IDType, r.IDCard, r.Truename, r.CreateTime, r.Status, r.ActiveTime, r.Mobile, r.CardNO, r.CardCreateTime, r.MobileUpdateTime, r.SendCount, r.LastSendDate, r.CommonCode, r.ID)
	return
}