package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pingcap/log"
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
	Id          uint64 `orm:"column(id)"`
	Title       string `orm:"column(title)"`
	Topic       string `orm:"column(parent_id)"`
	Content     string `orm:"column(user_id)"`
	Status      string `orm:"column(status)"`
	Username    string `orm:"column(username)"`
	CreatedTime string `orm:"column(created_time);type(timestamp);null"`
	UpdatedTime string `orm:"column(updated_time);type(timestamp);null"`
}

func (t *Emails) TableName() string {
	return "emails"
}

func (t *Emails) InitTable() {
	log.Info("call method initTable=====")
	orm.RegisterModel(new(Emails))
	//orm.RegisterModel(new(Attributes))
}
