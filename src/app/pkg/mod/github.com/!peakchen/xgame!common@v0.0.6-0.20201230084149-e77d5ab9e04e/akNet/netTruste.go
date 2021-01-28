package akNet

// add by stefan

var (
	GTrustedList   = map[string]bool{}
	GUnTrustedList = map[string]bool{}
)

func InitTrusted() {
	// for _, item := range serverConfig.GNetFilterConfig.Get() {
	// 	GTrustedList[item.White] = true
	// 	GUnTrustedList[item.Black] = true
	// }
}

func IsTrusted(ip string) bool {
	if GTrustedList[ip] {
		return true
	}

	//...
	return false
}

func IsUnTrusted(ip string) bool {
	return GUnTrustedList[ip]
}

func init() {
	InitTrusted()
}
