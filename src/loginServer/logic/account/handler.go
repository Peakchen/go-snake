package account

import (
	"go-snake/akmessage"
	"go-snake/common/akOrm"
	"go-snake/loginServer/base"
	"go-snake/loginServer/entityMgr"
	"go-snake/loginServer/logic/account/acc_model"
	"go-snake/loginServer/msg"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_ACC_REGISTER), (*akmessage.CS_AccRegister)(nil), func(actor entityMgr.IEntityUser, pb proto.Message) {
		actor.HandlerRegister(pb)
	})
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_LOGIN), (*akmessage.CS_Login)(nil), func(actor entityMgr.IEntityUser, pb proto.Message) {
		actor.HandlerLogin(pb)
	})
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_LOGOUT), (*akmessage.CS_Logout)(nil), func(actor entityMgr.IEntityUser, pb proto.Message) {
		actor.HandlerLogout(pb)
	})
}

func (this *Acc) HandlerRegister(pb proto.Message) {
	reg := pb.(*akmessage.CS_AccRegister)
	akLog.FmtPrintln("AccRegister...", reg.Acc, reg.Pwd)

	rsp := func(acc *Acc, ret akmessage.ErrorCode) {
		res := &akmessage.SC_AccRegister{
			Ret: ret,
		}
		acc.SendMsg(akmessage.MSG_SC_ACC_REGISTER, res)
	}

	accM := &acc_model.Acc{}
	if akOrm.HasExistAcc(accM, reg.Acc, reg.Pwd) {
		rsp(this, akmessage.ErrorCode_AccountExisted)
		return
	}

	accM = acc_model.NewAcc(reg.Acc, reg.Pwd)
	if accM != nil {
		this.user = accM
		this.SetSessionID(this.GetSessionID())
		this.SetID(accM.GetUserID())
		base.GetEntityMgr().SetEntityByID(accM.GetUserID(), this)
		rsp(this, akmessage.ErrorCode_Success)
	} else {
		rsp(this, akmessage.ErrorCode_Invaild)
	}
}

func (this *Acc) HandlerLogin(pb proto.Message) {
	login := pb.(*akmessage.CS_Login)
	akLog.FmtPrintln("login...", login.Acc, login.Pwd)

	rsp := func(acc *Acc, ret akmessage.ErrorCode) {
		res := &akmessage.SC_Login{
			Ret: ret,
		}
		acc.SendMsg(akmessage.MSG_SC_LOGIN, res)
	}

	accM := &acc_model.Acc{}
	if !akOrm.HasExistAcc(accM, login.Acc, login.Pwd) {
		rsp(this, akmessage.ErrorCode_Invaild)
		return
	}
	this.user = accM
	this.SetSessionID(this.GetSessionID())
	this.SetID(accM.GetUserID())
	base.GetEntityMgr().SetEntityByID(accM.GetUserID(), this)
	rsp(this, akmessage.ErrorCode_Success)
}

func (this *Acc) HandlerLogout(pb proto.Message) {
	logout := pb.(*akmessage.CS_Logout)
	akLog.FmtPrintln("logout...", logout)
}
