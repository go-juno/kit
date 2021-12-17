package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// LocalTime 封装时间
type LocalTime struct {
	time.Time
}

// MarshalJSON json序列化
func (t LocalTime) MarshalJSON() ([]byte, error) {
	if !t.IsZero() && t.UnixNano() > 0 {
		dateString := t.Format("2006-01-02 15:04:05")
		return json.Marshal(dateString)
	}
	return json.Marshal(nil)

}

// UnmarshalJSON json反列化
func (t *LocalTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	return
}

// Value ORM 序列化
func (t LocalTime) Value() (driver.Value, error) {

	return t.Format("2006-01-02 15:04:05"), nil
}

// Scan ORM 反序列化
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// NullString 数据库字符串空值
type NullString struct {
	sql.NullString
}

// MarshalJSON 数据库字符串空值json序列化
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON 数据库字符串空值json反序列化
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil && ns.String != "")
	return err
}

// Value 数据库字符串空值orm序列化
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}

// NullUint 数据库uint空值
type NullUint struct {
	Uint  uint
	Valid bool
}

// MarshalJSON 数据库uint空值json序列化
func (ns NullUint) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.Uint)
}

// UnmarshalJSON 数据库uint空值json反序列化
func (ns *NullUint) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.Uint)
	ns.Valid = (err == nil || ns.Uint == 0)
	return err
}

// Value 数据库uint空值orm序列化
func (ns NullUint) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Uint, nil
}

// LocalDate 封装日期
type LocalDate struct {
	time.Time
}

// MarshalJSON json序列化
func (t LocalDate) MarshalJSON() ([]byte, error) {
	if !t.IsZero() && t.UnixNano() > 0 {
		dateString := t.Format("2006-01-02")
		return json.Marshal(dateString)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON json反序列化
func (t *LocalDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "" {
		t.Time = time.Time{}
		return
	}
	t.Time, err = time.Parse("2006-01-02", s)
	return
}

// Value orm 序列化
func (t LocalDate) Value() (driver.Value, error) {
	return t.Format("2006-01-02"), nil
}

// Scan orm 反序列化
func (t *LocalDate) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalDate{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Model 基础表字段
type Model struct {
	Id        uint           `json:"id" gorm:"primary_key"`
	CreatedAt LocalTime      `json:"created_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6);comment:创建时间"`
	UpdatedAt LocalTime      `json:"updated_at" gorm:"type:datetime(6);not null;default:CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6);comment:更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
