package dao

import (
	"log"
	"ypp-wywk-service/model"
)

const (
	_smsCount = "SELECT COUNT(*) FROM sms_count WHERE source = ? AND create_date = ?"
	_smsByOne = "SELECT send_count FROM sms_count WHERE source = ? AND create_date = ? LIMIT 1"
	_smsInsert = "INSERT INTO sms_count(create_date, source, send_count) VALUES(?, ?, ?)"
	_upInsert = "UPDATE sms_count SET send_count = ? WHERE source = ? AND create_date = ?"
)

func (d *Dao) SmsCount(source string, now string) (count int64, err error) {
	row := d.dbreport.QueryRow(_smsCount, source, now)
	if err = row.Scan(&count); err != nil {
		if err == noRow {
			err = nil
			count = 0
		} else {
			log.Println("DB query failed:", err.Error())
		}
	}
	return
}

func (d *Dao) SmsByOne(source string, now string) (sendCount int64, err error) {
	row := d.dbreport.QueryRow(_smsByOne, source, now)
	if err = row.Scan(&sendCount); err != nil {
		if err == noRow {
			err = nil
			sendCount = 0
		} else {
			log.Println("DB query failed:", err.Error())
		}
	}
	return
}

func (d *Dao) AddReport(r *model.TReport) (err error) {
	sql, err := d.dbreport.Prepare(_smsInsert)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(r.CreateDate, r.Source, r.SendCount)
	return
}

func (d *Dao) UpdateReport(sendCount int, source string, now string) (err error) {
	sql, err := d.dbreport.Prepare(_upInsert)
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
	_, err = sql.Exec(sendCount, source, now)
	return
}

