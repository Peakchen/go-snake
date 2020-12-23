package messageBase

import "go-snake/akmessage"

func GetActorRegisterReq(sid string, st akmessage.ServerType) []byte {
	return SSPackMsg_pb(sid, 0, akmessage.MSG_SS_REGISTER_REQ, &akmessage.SS_Register_Req{St: st})
}

func GetActorRegisterRsp(sid string, st akmessage.ServerType) []byte {
	return SSPackMsg_pb(sid, 0, akmessage.MSG_SS_REGISTER_RSP, &akmessage.SS_Register_Resp{})
}

func SS_HeatBeatMsg(sid string) []byte {
	return SSPackMsg_pb(sid, 0, akmessage.MSG_SS_HEARTBEAT_REQ, &akmessage.SS_HeartBeat_Req{})
}

func CS_HeatBeatMsg(sid string) []byte {
	return nil
}
