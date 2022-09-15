package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/pingcap/log"
)

type Configs struct {
	Id          uint64 `orm:"column(id)"`
	Topic       string `orm:"column(topic)"`
	PeerId      string `orm:"column(peer_id)"`
	Username    string `orm:"column(username)"`
	CreatedTime string `orm:"column(created_time);type(timestamp);null"`
	UpdatedTime string `orm:"column(updated_time);type(timestamp);null"`
}

func (t *Configs) TableName() string {
	return "configs"
}

func (t *Configs) InitTable() {
	log.Info("call method initTable=====")
	orm.RegisterModel(new(Configs))
	//orm.RegisterModel(new(Attributes))
}
