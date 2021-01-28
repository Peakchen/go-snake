package akNet

import "sync"

// add by stefan

// client session pool

var (
	GClientSessionPool *TClientSession
	GServerSessionPool *TServerSession
)

func init() {
	GClientSessionPool = &TClientSession{
		sessionpool: &sync.Pool{
			New: func() interface{} {
				return new(TConnSession)
			},
		},
	}

	GServerSessionPool = &TServerSession{
		sessionpool: &sync.Pool{
			New: func() interface{} {
				return new(TConnSession)
			},
		},
	}
}

type TClientSession struct {
	sessionpool *sync.Pool
}

func (this *TClientSession) Push(s *TConnSession) {
	this.sessionpool.Put(s)
}

func (this *TClientSession) Get() (s *TConnSession) {
	s = this.sessionpool.Get().(*TConnSession)
	return
}

// server session pool

type TServerSession struct {
	sessionpool *sync.Pool
}

func (this *TServerSession) Push(s *TConnSession) {
	this.sessionpool.Put(s)
}

func (this *TServerSession) Get() (s *TConnSession) {
	s = this.sessionpool.Get().(*TConnSession)
	return
}
