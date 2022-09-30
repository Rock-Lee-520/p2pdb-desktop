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
	Id             uint64 `orm:"column(id)"`
	EmailId        string `orm:"column(email_id)"`
	Title          string `orm:"null;column(title)"`
	Topic          string `orm:"null;column(topic)"`
	Content        string `orm:"null;column(content)"`
	SentPeerId     string `orm:"null;column(sent_peer_id)"`
	ReceivedPeerId string `orm:"null;column(received_peer_id)"`
	Status         string `orm:"null;column(status)"`
	Username       string `orm:"null;column(username)"`
	CreatedTime    string `orm:"null;column(created_time);type(timestamp);null"`
	UpdatedTime    string `orm:"null;column(updated_time);type(timestamp);null"`
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
	log.Debug(t)
	id, err := Ormer.Insert(t)
	return id, err
}
