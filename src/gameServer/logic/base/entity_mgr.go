package base

type EntityManager struct {
	Users map[string]*IEntityUser
}

var (
	entitys = &EntityManager{
		Users: make(map[string]*IEntityUser),
	}
)

func GetEntityMgr() *EntityManager {
	return entitys
}

func GetUserByID(id string) *IEntityUser {
	return entitys.Users[id]
}

func (this *EntityManager) GetEntityByID(rid string) *IEntityUser {
	return this.Users[rid]
}

func (this *EntityManager) SetEntityByID(rid string, entity *IEntityUser) {
	this.Users[rid] = entity
}
