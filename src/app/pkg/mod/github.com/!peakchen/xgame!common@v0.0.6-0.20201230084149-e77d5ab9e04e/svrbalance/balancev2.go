package svrbalance

// balance v1: loop servers person, if some one has location to connect, then will push client into it.
// add stefan 20190704 16:25
/*
	{

		s1
		s2     -------->  ctrl svr balance
		s3

								server	person
						---------s1  	sx
		client	select	---------s2	 	sy				rand server for client to connect.
						---------s3	 	sz
		if s1 person sx has beyond max person limit, then begin loop find s2, which sy has not arive person limit,
		server will distribute s2 for client connection. firstly, rand server to connect.
	}
*/
import (
	"math/rand"
	"time"
)

type TSvrBalanceV2 struct {
	sb map[string]*TExternal
}

func (this *TSvrBalanceV2) NewBalance() {

}

func (this *TSvrBalanceV2) AddSvr(svr string) {
	_, ok := this.sb[svr]
	if ok {
		return
	}

	this.sb[svr] = &TExternal{
		Persons: 0,
	}
}

// some one connect gateway to balance route push one server.
func (this *TSvrBalanceV2) Push(svr string) {
	ex, ok := this.sb[svr]
	if ok {
		return
	}

	ex.Persons++
}

func (this *TSvrBalanceV2) getsvr() (svrs []string) {
	svrs = []string{}
	for svr, _ := range this.sb {
		svrs = append(svrs, svr)
	}
	return
}

// get second max server persons
func (this *TSvrBalanceV2) GetSvr() (s string) {
	var (
		svrs    []string = this.getsvr()
		svrslen int      = len(svrs)
	)

	randidx := rand.Intn(svrslen)
	s = svrs[randidx]
	return
}
