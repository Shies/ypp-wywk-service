package model

import (
	"time"
)

type TReport struct {
	CreateDate time.Time
	Source 	   string
	SendCount  int
	OpenCount  int
}

type TRegister struct {
	ID 		   int64
	Data 	   string
	CreateTime time.Time
}

type TUser struct {
	ID 				string
	YppNO			string
	Mobile			string
	Password 		string
	Status			int
	Token			string
	IsBind			int
	CardNO			string
	CreateTime  	time.Time
	BindTime		time.Time
	IsOnline		int
	IsSignHx		int
	IsGod			int
	Source 			string
	LastLoginTime 	time.Time
	ViewType		int
	IsAuth			int
	IsVUser			int
	ZhimaOpenID 	string
	WechatUnionID 	string
	QQOpenID 		string
	Device 			string
	IsSignYx		int8
	MarketChannel 	string
	UpdateTimeStamp time.Time
	TmpFlag		  	int8
	IsFace		  	int8
	BundleID 	  	string
}

type TWyRegister struct {
	ID 			  	 string
	TypeID 		  	 int
	IDType			 string
	IDCard			 string
	Truename		 string
	CreateTime		 time.Time
	Status 		     int8
	ActiveTime		 time.Time
	Mobile 		     string
	CardNO			 string
	CardCreateTime   time.Time
	MobileUpdateTime time.Time
	SendCount		 int
	LastSendDate	 time.Time
	CommonCode 	     string
	UpdateTimeStamp  time.Time
}
