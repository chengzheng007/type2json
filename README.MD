this package provide simple function,to convert time or date type marshal to json
in golang, we know that, a common time.Time type, when marshal it to json, it has extra info we don't need:

```go
	type Temp struct {
		MyTime time.Time `json:"my_time"`
	}
	a := &Temp{MyTime:time.Now()}
	buf, _ := json.Marshal(a)
	fmt.Println(string(buf))
```

output:

```
{"ctime":"2018-07-12T20:07:11.949157595+08:00"}
```

this package can marshal the common time.Time to yyyy-mm-dd hh:mm:ss，no other redundant things, there is also Date and Stmp type，but notice that, unmarshal Date to time.Time will lose the time part!!

use example:

```go

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
```

output

```
marshal byt({"my_date":"2018-07-12","my_time":"2018-07-13 19:58:47","my_stmp":"1531396727"})
unmarshal MyDate(2018-07-12 00:00:00) MyTime(2018-07-13 19:58:47) MyStmp(2018-07-12 19:58:47)
```

Additional, if you define a struct like A, it also support sql Scan, you can Scan value from database timestamp type into MyTime directly！

Enjoy it!


