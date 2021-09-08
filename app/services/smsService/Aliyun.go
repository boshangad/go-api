package smsService

type AliyunGateway struct {
	SmsService
}

//func (that AliyunGateway) Send() (err error) {
//	var (
//		client = db.DefaultClient()
//		ctx = context.Background()
//		smsLogModel *ent.SmsLog
//		typeId uint64
//		ok bool
//		dataStr []byte
//	)
//	if typeId, ok = TypeCorrespondId[""]; !ok {
//		err = errors.New("sending failed, undeclared type")
//		return
//	}
//	dataStr, err = json.Marshal(data)
//	if err != nil {
//		return
//	}
//	smsLogModel, err = client.SmsLog.Create().
//		SetAppID(appId).
//		SetDialCode(dialCode).
//		SetMobile(mobile).
//		SetScope(that.Scope).
//		SetGateway(that.GateWay).
//		SetIP(that.Ip).
//		SetContent(string(dataStr)).
//		SetTemplateID(that.TemplateId).
//		SetTemplateText(that.TemplateText).
//		SetStatus(StatusDraft).
//		Save(ctx)
//	if err != nil {
//		return
//	}
//
//	return
//}

func (that AliyunGateway) CheckCode(dialCode, mobile, code string, typeId uint64, appId uint64) (err error)  {
	return
}