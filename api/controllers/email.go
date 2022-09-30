package controllers

import (
	"api/models"

	"github.com/Rock-liyi/p2pdb/infrastructure/util/function"
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
)

func (this *MainController) Email() {

	this.TplName = "email/email.html"
}

type ReturnData struct {
	code    int         `json:"code"`
	message string      `json:"message"`
	data    interface{} `json:"data"`
}

// @Title SendEmail
// @Description Logs out current logged in user session
// @Success 200 {interface} SendEmail success
// @router /send-email [post]
func (this *MainController) SendEmail() {
	Topic := this.GetString("Topic")

	Title := this.GetString("Title")
	Content := this.GetString("Content")
	Username := "rock"
	Peerid := this.GetString("peerid")
	log.Debug(Topic)
	log.Debug(Title)
	log.Debug(Content)

	email := new(models.Emails)
	email.EmailId = function.GetUUID()
	email.Title = Title
	email.Topic = Topic
	email.Content = Content
	email.Username = Username
	email.Status = models.SENT
	email.SentPeerId = function.GetLocalPeerId()
	email.ReceivedPeerId = Peerid
	email.Id = uint64(function.GetSnowflakeId())
	id, err := email.Insert()
	log.Debug(id)
	if err == nil {
		this.Data["json"] = map[string]string{"code": "2000", "message": "Sending email is success", "data": string(function.GetSnowflakeId())}
	} else {
		this.Data["json"] = map[string]string{"code": "2000", "message": err.Error(), "data": ""}
	}

	//&ReturnData{code: 1000, message: "Sending email is success", data: ""}
	this.ServeJSON()

	//this.TplName = "email/compose.html"
	//return &ReturnData{code: 1000, message: "Sending email is success", data: ""}
}

func (this *MainController) Compose() {

	this.TplName = "email/compose.html"
}
