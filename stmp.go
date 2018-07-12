package type2json

import (
	"database/sql/driver"
	"regexp"
	"strconv"
	"time"
)

// 转为时间戳
type Stmp time.Time

// sql driver exec
func (this *Stmp) Value() (driver.Value, error) {
	return this.GetValue(), nil
}

// sql driver scan
func (this *Stmp) Scan(v interface{}) error {
	switch buf := v.(type) {
	case time.Time:
		this.SetValue(buf)
	default:
		return ErrType
	}
	return nil
}

// 赋值
func (this *Stmp) SetValue(v time.Time) {
	*this = Stmp(v)
}

// 取值
func (this *Stmp) GetValue() time.Time {
	return time.Time(*this)
}

// 生成json
func (this *Stmp) MarshalJSON() ([]byte, error) {
	stamp := strconv.Quote(strconv.FormatInt(timetoint64(this.GetValue()), 10))
	return []byte(stamp), nil
}

// 解析json
func (this *Stmp) UnmarshalJSON(szData []byte) error {
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
	// 正则判断首尾是否为数字
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

