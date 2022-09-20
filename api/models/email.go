package models

import (
	"github.com/Rock-liyi/p2pdb/infrastructure/util/log"
	"github.com/beego/beego/v2/client/orm"
)

const (
	INBOX     = "inbox"
	SENT      = "sent"
	IMPORTANT = "important"
	DRAFTS    = "drafts"
	SPAM      = "spam"
	TRASH     = "trash"
)

type Emails struct {
	//Id             string `orm:"pk;column(id)"`
	EmailId        string `orm:"pk;column(email_id)"`
	Title          string `orm:"column(title)"`
	Topic          string `orm:"column(topic)"`
	Content        string `orm:"column(content)"`
	SentPeerId     string `orm:"column(sent_peer_id)"`
	ReceivedPeerId string `orm:"column(received_peer_id)"`
	Status         string `orm:"column(status)"`
	Username       string `orm:"column(username)"`
	CreatedTime    string `orm:"column(created_time);type(timestamp);null"`
	UpdatedTime    string `orm:"column(updated_time);type(timestamp);null"`
}

func (t *Emails) TableName() string {
	return "emails"
}

func (t *Emails) InitTable() {
	log.Info("call method initTable=====")
	orm.RegisterModel(new(Emails))
	//orm.RegisterModel(new(Attributes))
}

func (t *Emails) Insert() (int64, error) {
	var Ormer orm.Ormer
	Ormer = orm.NewOrm()
	email := new(Emails)
	id, err := Ormer.Insert(email)
	return id, err
}
