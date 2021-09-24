package smsService

type AliyunGateway struct {
	GatewayInterface
	Data
}

//func (that *AliyunGateway) Send() (err error) {
//	var (
//		client = db.DefaultClient()
//		ctx = context.Background()
//		smsLogModel *ent.SmsLog
//		typeId uint64
//		ok bool
//		dataStr []byte
//	)
//	dataStr, err = json.Marshal(that.Content)
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
//
//func (that AliyunGateway) Check(code string) (err error)  {
//	return
//}