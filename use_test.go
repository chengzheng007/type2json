package type2json

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

/*
  this test case show you how to use packae
*/

func TestStmp(t *testing.T) {
	type A struct {
		MyDate Date `json:"my_date"`
		MyTime Time `json:"my_time"`
		MyStmp Stmp `json:"my_stmp"`
	}

	a := &A{}
	a.MyDate.SetValue(time.Now())
	a.MyTime.SetValue(time.Now().AddDate(0, 0, 1))
	a.MyStmp.SetValue(time.Now())

	byt, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("json.Marshal(%v) error(%v)\n", a, err)
		return
	}
	fmt.Printf("marshal byt(%s)\n", byt)

	b := &A{}
	if err := json.Unmarshal(byt, b); err != nil {
		fmt.Printf("json.Unmarshal(%s, b) error(%v)\n", byt, err)
		return
	}
	fmt.Printf("unmarshal MyDate(%v) MyTime(%v) MyStmp(%v)\n",
		b.MyDate.GetValue().Format("2006-01-02 15:04:05"),
		b.MyTime.GetValue().Format("2006-01-02 15:04:05"),
		b.MyStmp.GetValue().Format("2006-01-02 15:04:05"))
}
