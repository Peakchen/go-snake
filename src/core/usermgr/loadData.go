package usermgr

func (self *EntityManager) Add(fs ...func(em *EntityManager)){
	
	for _, f := range fs {
		self.loads = append(self.loads, f)
	}

}

func (this *EntityManager) LoadAll() {

	for _, f := range this.loads{
		f(this)	
	}

}

