package models

import (
	"errors"
	"filetool/models/system"
)

var NULL = ""

var ErrFielExisted = errors.New("select field Existed")
var ErrUnpermission = errors.New("without permission")
var ErrNotFound = errors.New("sql: no rows in result set")
var ErrNoMoreFound = errors.New("sql: no more found in db")
var ErrCountNum = errors.New("sql: count failed in db")
var ErrNotAsset = errors.New("wrong: device not in asset")
var StationNoExisted = errors.New("stationNo existed")

type (
	UserInfo struct {
		system.TUserInfo
	}

	ChatRecord struct {
		system.TChatRecordInfo
	}
)
