// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
	"database/sql/driver"
)

type CabinetBrand string

const (
	CabinetBrandUnknown        CabinetBrand = "UNKNOWN"
	CabinetBrandKaixin         CabinetBrand = "KAIXIN"
	CabinetBrandYundong        CabinetBrand = "YUNDONG"
	CabinetBrandTuobang        CabinetBrand = "TUOBANG"
	CabinetBrandXiliulouServer CabinetBrand = "XILIULOUSERV" // 西六楼服务器版
)

func (b CabinetBrand) Name() string {
	switch b {
	case CabinetBrandKaixin:
		return "凯信"
	case CabinetBrandYundong:
		return "云动"
	case CabinetBrandTuobang:
		return "拓邦"
	case CabinetBrandXiliulouServer:
		return "西六楼"
	}
	return "-"
}

func (b CabinetBrand) RpcName() string {
	switch b {
	default:
		return ""
	case CabinetBrandKaixin:
		return "kxcab"
	case CabinetBrandYundong:
		return "ydcab"
	case CabinetBrandTuobang:
		return "tbcab"
	case CabinetBrandXiliulouServer:
		return "xllscab"
	}
}

func (b CabinetBrand) String() string {
	return string(b)
}

func (b *CabinetBrand) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case string:
		*b = CabinetBrand(v)
	}
	return nil
}

func (b CabinetBrand) Value() (driver.Value, error) {
	return string(b), nil
}

type BatteryBrand string

const (
	BatteryBrandUnknown BatteryBrand = "UNKNOWN" // 未知
	BatteryBrandXC      BatteryBrand = "XC"      // 星创电池
	BatteryBrandTB      BatteryBrand = "TB"      // 拓邦电池
)

func (b BatteryBrand) RpcName() string {
	switch b {
	default:
		return ""
	case BatteryBrandXC:
		return "xcbms"
	case BatteryBrandTB:
		return "xcbms"
	}
}

func (b BatteryBrand) String() string {
	return string(b)
}

func (b *BatteryBrand) Scan(src interface{}) error {
	switch v := src.(type) {
	case nil:
		return nil
	case string:
		*b = BatteryBrand(v)
	}
	return nil
}

func (b BatteryBrand) Value() (driver.Value, error) {
	return string(b), nil
}
