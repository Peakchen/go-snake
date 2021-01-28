// add by stefan
package define

import (
	"net/http"
)

type ERouteId int32

const (
	ERouteId_ER_Invalid    ERouteId = 0
	ERouteId_ER_ESG        ERouteId = 1
	ERouteId_ER_ISG        ERouteId = 2
	ERouteId_ER_DB         ERouteId = 3
	ERouteId_ER_BigWorld   ERouteId = 4
	ERouteId_ER_Login      ERouteId = 5
	ERouteId_ER_SmallWorld ERouteId = 6
	ERouteId_ER_DBProxy    ERouteId = 7
	ERouteId_ER_Game       ERouteId = 8
	ERouteId_ER_Client     ERouteId = 9
	ERouteId_ER_CenterGate ERouteId = 10
	ERouteId_ER_MMS        ERouteId = 11
	ERouteId_ER_ISG_SERVER ERouteId = 12
	ERouteId_ER_ISG_CLIENT ERouteId = 13
	ERouteId_ER_SG         ERouteId = 14
	ERouteId_ER_Max        ERouteId = (ERouteId_ER_SG + 1)
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

type HandlerServMux struct {
	muxhandler *http.ServeMux
}

//message format
/*
	Session string/int64
	MainID	int32
	SubID	int32
	message proto.Message
*/
