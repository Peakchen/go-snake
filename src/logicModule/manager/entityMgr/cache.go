package entityMgr

var emgr *BaseEntityMgr

func GetEntityMgr()*BaseEntityMgr{
	
	if emgr == nil {
		emgr = newEntityMgr()
	}

	return emgr
}