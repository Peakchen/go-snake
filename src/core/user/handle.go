package user


type RegHandler func(this *EntityUser, entity interface{})

var entityRegMap = map[int]RegHandler{
	M_ACCOUNT: 		SetAcc,
	M_SERVERINNER: 	SetInner,
	M_WXROLE: 		SetWxRole,
	M_ROLE: 		SetRole,
	M_Mail: 		SetMail,
	M_Chat: 		SetChat,
	M_Friend: 		SetFriend,
	M_Map: 			SetMap,
}

func SetAcc(this *EntityUser, entity interface{}){
	this.IAcc = entity.(IAcc)
}

func SetInner(this *EntityUser, entity interface{}){
	this.IInner = entity.(IInner)
}

func SetWxRole(this *EntityUser, entity interface{}){
	this.IWxRole = entity.(IWxRole)
}

func SetRole(this *EntityUser, entity interface{}){
	this.IRole = entity.(IRole)
}

func SetMail(this *EntityUser, entity interface{}){
	this.IMail = entity.(IMail)
}

func SetChat(this *EntityUser, entity interface{}){
	this.IChat = entity.(IChat)
}

func SetFriend(this *EntityUser, entity interface{}){
	this.IFriend = entity.(IFriend)
}

func SetMap(this *EntityUser, entity interface{}){
	this.IMap = entity.(IMap)
}
