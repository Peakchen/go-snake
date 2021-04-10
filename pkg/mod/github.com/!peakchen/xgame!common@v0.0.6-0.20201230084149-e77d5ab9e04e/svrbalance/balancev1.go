package svrbalance

// balance v1: loop servers person, if some one has location to connect, then will push client into it.
// add stefan 20190702 17:30
/*
	{

		s1
		s2     -------->  ctrl svr balance
		s3

								server	person
						---------s1  	sx
		client	select	---------s2	 	sy				select second max person server to client.
						---------s3	 	sz
		if s1 person sx has beyond max person limit, then begin loop find s2, which sy has not arive person limit,
		server will distribute s2 for client connection. firstly, full server persons , then find next server.
	}
*/
import ()

type TSvrBalanceV1 struct {
	sb map[string]*TExternal
}

func (this *TSvrBalanceV1) NewBalance() {

}

func (this *TSvrBalanceV1) AddSvr(svr string) {
	_, ok := this.sb[svr]
	if ok {
		return
	}

	this.sb[svr] = &TExternal{
		Persons: 0,
	}
}

// some one connect gateway to balance route push one server.
func (this *TSvrBalanceV1) Push(svr string) {
	ex, ok := this.sb[svr]
	if ok {
		return
	}

	ex.Persons++
}

// get second max server persons
func (this *TSvrBalanceV1) GetSvr() (s string) {
	var (
		secmin int32 = 0
		max    int32 = 0
		loop   int   = 0
		sblen  int   = len(this.sb)
	)
	for svr, ex := range this.sb {
		loop++
		if ex.Persons > max {
			max = ex.Persons
		}
		if secmin == 0 {
			secmin = ex.Persons
		} else if secmin < max && secmin < ex.Persons {
			secmin = ex.Persons
		}
		if secmin > 0 && (loop+1 == sblen) {
			s = svr
			break
		}
	}
	return
}
