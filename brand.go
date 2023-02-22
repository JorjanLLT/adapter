// Copyright (C) liasica. 2023-present.
//
// Created at 2023-01-14
// Based on adapter by liasica, magicrolan@qq.com.

package adapter

import (
    "database/sql/driver"
    "strings"
)

type CabinetBrand string

const (
    CabinetBrandUnknown CabinetBrand = "UNKNOWN"
    CabinetBrandKaixin  CabinetBrand = "KAIXIN"
    CabinetYundong      CabinetBrand = "YUNDONG"
)

func (b CabinetBrand) LoggerName(env Environment) string {
    var prefix string
    if env != Production {
        prefix = strings.ToLower(env.String()) + "-"
    }
    return prefix + "cabinet-" + strings.ToLower(b.String())
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
    return b, nil
}

type BatteryBrand string

const (
    BatteryBrandUnknown BatteryBrand = "UNKNOWN"
    BatteryBrandXC      BatteryBrand = "XC"
)

func (b BatteryBrand) LoggerName(env Environment) string {
    var prefix string
    if env != Production {
        prefix = strings.ToLower(env.String()) + "-"
    }
    return prefix + "battery-" + strings.ToLower(b.String())
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
    return b, nil
}
