package akWebNet

import (
	"context"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_HeartBeat"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

//广播消息
/*
	@param 1: 自身会话
	@param 2：是否不给自己广播
	@param 3：消息ID
	@param 4：消息参数
*/
func BroadCastMsgExceptSession(sess *WebSession, bMsg2Me bool, mainId, subId uint16, data proto.Message) {
	msg, err := PackMsgOp(mainId, subId, data, PACK_PROTO)
	if msg == nil || err != nil {
		akLog.Error("pack msg fail: ", mainId, subId)
		return
	}
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func(k, v interface{}) bool {
		if v != nil {
			sess := v.(*WebSession)
			if !bMsg2Me && sess.RemoteAddr == sess.RemoteAddr {
				return true
			}
			sess.Write(websocket.BinaryMessage, msg)
		}

		return true
	})
}

func BroadCastMsgExceptID(mainId, subId uint16, data proto.Message) {
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func(k, v interface{}) bool {
		if v != nil {
			sess := v.(*WebSession)
			msg, err := PackMsgOp(mainId, subId, data, PACK_PROTO)
			if msg == nil {
				akLog.Error("pack msg fail: ", mainId, subId, err)
				return false
			}
			sess.Write(websocket.BinaryMessage, msg)
		}

		return true
	})
}

func loopSignalCheck(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
	}()

	chsignal := make(chan os.Signal, 1)
	signal.Notify(chsignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-chsignal:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				akLog.FmtPrintln("check exit signal:", s)
				return
			default:
				akLog.FmtPrintln("other signal:", s)
			}
		}
	}
}

func SendMsg(sess *WebSession, mainId, subId uint16, data proto.Message) (succ bool, err error) {
	var msg []byte
	msg, err = PackMsgOp(mainId, subId, data, PACK_PROTO)
	if err != nil {
		akLog.Error("pack msg fail: ", mainId, subId, err)
		return
	}
	sess.Write(websocket.BinaryMessage, msg)
	return true, nil
}

func IsGateWayActor(actor define.ERouteId) bool {
	return actor == define.ERouteId_ER_ESG || actor == define.ERouteId_ER_ISG_SERVER || actor == define.ERouteId_ER_ISG_CLIENT
}

func MsgProc(sess *WebSession, data []byte, pt PACK_TYPE) {
	info, err := GetUnPackMsgInfo(data, pt)
	if err != nil || info == nil {
		akLog.Error("msg un pack info fail: ", pt, err)
		return
	}
	msgCallBack := func(dstses *WebSession, pt PACK_TYPE) {
		msg, cb, err := UnPackMsgOp(pt)
		if err != nil {
			akLog.Error("msg proc fail: ", pt, err)
			return
		}
		//callback define: func (sess *WebSession, proto Message)(bool,error)
		params := []reflect.Value{
			reflect.ValueOf(dstses),
			reflect.ValueOf(msg),
		}
		ret := cb.Call(params)
		succ := ret[0].Interface().(bool)
		reterr := ret[1].Interface()
		if reterr != nil || !succ {
			akLog.Error("message proc return err: ", reterr.(error).Error())
		}
	}
	actor := sess.GetActor().GetActorType()
	if IsGateWayActor(actor) {
		switch info.Actor {
		case 0:
			msgCallBack(sess, pt)
		case uint16(define.ERouteId_ER_SG):
			msgCallBack(sess, pt)
		case uint16(define.ERouteId_ER_Login),
			uint16(define.ERouteId_ER_Game):
			var dstsess *WebSession
			if actor == define.ERouteId_ER_ESG { // route msg to inner client
				dstsess = GwebSessionMgr.GetSessionByActor(define.ERouteId_ER_ISG)
			} else if actor == define.ERouteId_ER_ISG_SERVER { // route msg to external gateway
				dstsess = GwebSessionMgr.GetSessionByActor(define.ERouteId_ER_ESG)
			} else if actor == define.ERouteId_ER_ISG_CLIENT { // route msg to gate way innner client
				dstsess = GwebSessionMgr.GetSessionByActor(define.ERouteId(info.Actor))
			} else {
				akLog.Error("invalid server actor.")
				return
			}
			if dstsess != nil {
				dstsess.Write(websocket.BinaryMessage, data)
			}
		default:
			akLog.Error("invalid route actor")
		}
	} else {
		msgCallBack(sess, pt)
	}
}

func sendHeartBeat(sess *WebSession) (succ bool, err error) {
	rsq := &MSG_HeartBeat.CS_HeartBeat_Req{}
	return SendMsg(sess, uint16(MSG_MainModule.MAINMSG_HEARTBEAT), uint16(MSG_HeartBeat.SUBMSG_CS_HeartBeat), rsq)
}
