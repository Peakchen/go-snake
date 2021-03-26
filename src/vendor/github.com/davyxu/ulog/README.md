# ulog
Structured, easy use logger for golang


* 基本用法
```go
    // 基本用法
	Global().SetLevel(DebugLevel)
	Debugln("debug")
	Infof("info")
	Warnf("warning %d", 123)
	Errorf("error %d", 567)

	Global().SetReportCaller(true)
	// 全局颜色输出
	WithColorName("purple").Infoln("WithColorName ", Purple.String())
	WithColor(DarkGreen).Errorf("WithColor %s", DarkGreen.String())


    // 独立日志实例
	l := New()

	textFormat := &TextFormatter{
		EnableColor: true,
	}

	// 从文本解析
	err := textFormat.ParseColorRule(`
    {
        "Rule":[
            {"Text":"panic:","Color":"Red"},
            {"Text":"[DB]","Color":"Green"},
            {"Text":"#http.listen","Color":"Blue"},
            {"Text":"#http.recv","Color":"Blue"},
            {"Text":"#http.send","Color":"Purple"}
        ]
    }
    `)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	l.SetFormatter(textFormat)

	// 从文本确定颜色
	l.Infoln("panic: must be red")
	l.Infoln("[DB] write db")
	l.Infoln("#http.recv come data")
```

* Json输出
```go
	Global().SetFormatter(&JSONFormatter{})
	Global().SetLevel(DebugLevel)
	Global().SetReportCaller(true)
	// 单行kv
	WithField("key", "value").Infof("noraml json")

	Global().SetFormatter(&JSONFormatter{
		PrettyPrint: true,
	})

	// 多行kv
	Global().WithFields(Fields{
		"name": "monk",
		"age":  80,
	}).Errorf("error json with pretty print")
```