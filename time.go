package type2json


import (
	"database/sql/driver"
	"regexp"
	"strconv"
	"time"
)

type Time time.Time

// sql driver exec
func (this *Time) Value() (driver.Value, error) {
	return this.GetValue(), nil
}

// sql driver scan
func (this *Time) Scan(v interface{}) error {
	switch buf := v.(type) {
	case time.Time:
		this.SetValue(buf)
	default:
		return ErrType
	}
	return nil
}

// 赋值
func (this *Time) SetValue(v time.Time) {
	*this = Time(v)
}

// 取值
func (this *Time) GetValue() time.Time {
	return time.Time(*this)
}

// 生成json
func (this *Time) MarshalJSON() ([]byte, error) {
	stamp := strconv.Quote(timetoa(this.GetValue()))
	return []byte(stamp), nil
}

// 解析json
func (this *Time) UnmarshalJSON(szData []byte) error {
	strData := string(szData)
	// 判断是否为null
	if strData == "null" {
		return nil
	}
	// 去掉双引号
	t, err := strconv.Unquote(strData)
	if err != nil {
		t = strData
	}
	// 正则判断是否为数字
	bNumber, err := regexp.MatchString("^\\d+$", t)
	if err != nil {
		return err
	}
	if !bNumber {
		v, err := atotime(t)
		this.SetValue(v)
		return err
	}
	i64, err := atoi64(t)
	if err != nil {
		return err
	}
	v, err := i64totime(i64)
	this.SetValue(v)
	return err
}
