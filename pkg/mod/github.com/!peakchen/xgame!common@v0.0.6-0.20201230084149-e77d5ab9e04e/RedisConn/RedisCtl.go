package RedisConn

/*
	ctrl some redis server change channl.
	redis can sentil n main and sub redis service.
*/

type TRedisCtrl struct {
	redisCns    []*TAokoRedis
	curRedisIdx int
}

func LazyInit() (redctl *TRedisCtrl) {
	redctl = &TRedisCtrl{}
	redctl.redisCns = []*TAokoRedis{}
	return
}

func (this *TRedisCtrl) AddRedisConn(c *TAokoRedis) {
	this.redisCns = append(this.redisCns, c)
}

func (this *TRedisCtrl) ChangeNextChannl() (c *TAokoRedis) {
	c = nil
	if this.curRedisIdx == len(this.redisCns) {
		return
	}
	mod := (this.curRedisIdx + 1) % len(this.redisCns)
	if mod == 0 {
		this.curRedisIdx = len(this.redisCns) - 1
	} else {
		this.curRedisIdx++
	}
	return this.redisCns[this.curRedisIdx]
}
