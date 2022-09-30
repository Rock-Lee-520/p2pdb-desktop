package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/pingcap/log"
)

type Configs struct {
	Id          uint64 `orm:"null;column(id)"`
	Topic       string `orm:"null;column(topic)"`
	PeerId      string `orm:"null;column(peer_id)"`
	Username    string `orm:"null;column(username)"`
	CreatedTime string `orm:"null;column(created_time);type(timestamp);null"`
	UpdatedTime string `orm:"null;column(updated_time);type(timestamp);null"`
}

func (t *Configs) TableName() string {
	return "configs"
}

func (t *Configs) InitTable() {
	log.Info("call method initTable=====")
	orm.RegisterModel(new(Configs))
	//orm.RegisterModel(new(Attributes))
}
