package account

import (
	"go-snake/akmessage"
	"go-snake/common/akOrm"
	"go-snake/core/usermgr"
	"go-snake/core/user"
	"go-snake/dbmodel/acc_model"
	"go-snake/core/msg"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_ACC_REGISTER), (*akmessage.CS_AccRegister)(nil), func(actor user.IEntityUser, pb proto.Message) {
		actor.HandlerRegister(pb)
	})
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_LOGIN), (*akmessage.CS_Login)(nil), func(actor user.IEntityUser, pb proto.Message) {
		actor.HandlerLogin(pb)
	})
	msg.RegisterActorMessageProc(uint32(akmessage.MSG_CS_LOGOUT), (*akmessage.CS_Logout)(nil), func(actor user.IEntityUser, pb proto.Message) {
		actor.HandlerLogout(pb)
	})
}

func (this *Acc) HandlerRegister(pb proto.Message) {
	reg := pb.(*akmessage.CS_AccRegister)
	//akLog.FmtPrintln("AccRegister...", reg.Acc, reg.Pwd)

	rsp := func(ret akmessage.ErrorCode) {
		res := &akmessage.SC_AccRegister{
			Ret: ret,
		}
		this.SendMsg(akmessage.MSG_SC_ACC_REGISTER, res)
	}

	accM := &accdb.Acc{}
	exist, err := akOrm.HasExistAcc(accM, reg.Acc, reg.Pwd)
	if err != nil {
		rsp(akmessage.ErrorCode_Invaild)
		return
	}
	if exist {
		rsp(akmessage.ErrorCode_AccountExisted)
		return
	}

	accM = accdb.NewAcc(reg.Acc, reg.Pwd)
	if accM != nil {
		akLog.FmtPrintln("AccRegister success...")

		this.user = accM
		this.SetSessionID(this.GetSessionID())
		this.SetID(accM.GetDBID())
		usermgr.GetEntityMgr().AddEnity(accM.GetDBID(), this)
		rsp(akmessage.ErrorCode_Success)
	} else {
		rsp(akmessage.ErrorCode_Invaild)
	}
}

func (this *Acc) HandlerLogin(pb proto.Message) {
	login := pb.(*akmessage.CS_Login)
	//akLog.FmtPrintln("login...", login.Acc, login.Pwd)

	rsp := func(ret akmessage.ErrorCode) {
		res := &akmessage.SC_Login{
			Ret: ret,
		}
		this.SendMsg(akmessage.MSG_SC_LOGIN, res)
	}

	accM := &accdb.Acc{}
	exist, err := akOrm.HasExistAcc(accM, login.Acc, login.Pwd)
	if err != nil {
		rsp(akmessage.ErrorCode_Invaild)
		return
	}
	if !exist {
		rsp(akmessage.ErrorCode_AccountNotExisted)
		return
	}

	akLog.FmtPrintln("login success...")
	this.user = accM
	this.SetSessionID(this.GetSessionID())
	this.SetID(accM.GetDBID())
	usermgr.GetEntityMgr().AddEnity(accM.GetDBID(), this)
	rsp(akmessage.ErrorCode_Success)
	
}

func (this *Acc) HandlerLogout(pb proto.Message) {
	logout := pb.(*akmessage.CS_Logout)
	akLog.FmtPrintln("logout...", logout)
}
