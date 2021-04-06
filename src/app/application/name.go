package application

var AppStr string

func SetAppName(node string) {
	AppStr = node
}

func GetAppName() string{
	return AppStr
}