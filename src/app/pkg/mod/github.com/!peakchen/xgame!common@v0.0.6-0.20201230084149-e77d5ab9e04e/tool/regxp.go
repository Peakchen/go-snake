package tool

// add by stefan

import (
	"regexp"
)

func IsChineseName(str string) bool {
	ok, err := regexp.MatchString(`^([\u4e00-\u9fa5][·\u4e00-\u9fa5]{0,30}[\u4e00-\u9fa5])$`, str)
	return ok && err == nil
}

func IsEmail(str string) bool {
	ok, err := regexp.MatchString(`[\w+\.]+[\w+]+@+[0-9A-Za-z]+\.+[A-Za-z]+$`, str)
	return ok && err == nil
}

func IsMobilePhone(str string) bool {
	ok, err := regexp.MatchString(`^((\+86)|(86))?(-|\s)?1\d{10}$`, str)
	return ok && err == nil
}

func IsOnlyChinese(str string) bool {
	regular := `^[\u4e00-\u9fa5]+$`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsIP(str string) bool {
	regular := `((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsYMD(str string) bool {
	//year-month-day
	regular := `^([0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)|([0-9]{2}(0[48]|[2468][048]|[13579][26])|(0[48]|[2468][048]|[13579][26])00)-02-29)$`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsHMS_APM(str string) bool {
	//hour:min:second xxx
	regular := `(0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9] ([AP]M)`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsHMS(str string) bool {
	//hour:min:second
	regular := `(0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9]`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsYMDHMS(str string) bool {
	//year-month-day hour:min:second
	regular := `^([0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)|([0-9]{2}(0[48]|[2468][048]|[13579][26])|(0[48]|[2468][048]|[13579][26])00)-02-29) (0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9]$`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsNumber(str string) bool {
	regular := `^[0-9]*$`
	return regexp.MustCompile(regular).MatchString(str)
}

func IsFloat(str string) bool {
	regular := `^[0-9]+([.]{0,1}[0-9]+){0,1}$`
	return regexp.MustCompile(regular).MatchString(str)
}
