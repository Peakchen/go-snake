package svrbalance

// balance v1: loop servers person, if some one has location to connect, then will push client into it.
// add stefan 20190704 16:00
/*
	{

		s1
		s2     -------->  ctrl svr balance
		s3

								server	person
						---------s1  	sx
		client	select	---------s2	 	sy              select min person server to client.
						---------s3	 	sz
		if s1 person sx has beyond max person limit, then begin loop find s2, which sy has not arive person limit,
		server will distribute s2 for client connection, firstly, select min server persons , then find next min server.
	}
*/
import ()

type TSvrBalanceV3 struct {
	sb map[string]*TExternal
}

func (this *TSvrBalanceV3) NewBalance() {

}

func (this *TSvrBalanceV3) AddSvr(svr string) {
	_, ok := this.sb[svr]
	if ok {
		return
	}

	this.sb[svr] = &TExternal{
		Persons: 0,
	}
}

// some one connect gateway to balance route push one server.
func (this *TSvrBalanceV3) Push(svr string) {
	ex, ok := this.sb[svr]
	if ok {
		return
	}

	ex.Persons++
}

// get min server persons
func (this *TSvrBalanceV3) GetSvr() (s string) {
	var (
		min   int32 = 0
		loop  int   = 0
		sblen int   = len(this.sb)
	)
	for svr, ex := range this.sb {
		loop++
		if min < ex.Persons {
			min = ex.Persons
		}
		if min > 0 && (loop+1 == sblen) {
			s = svr
			break
		}
	}
	return
}
