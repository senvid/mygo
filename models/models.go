package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"

	// "github.com/unknwon/com"
	// "os"
	// "path"
	"strconv"
	"time"
)

// const (
// 	_DB_NAME  = "data/beeblog.db"
// 	_DRIVER   = "sqlite3"
// )
const (
	_DB_NAME = "blog:blog@tcp(localhost:3306)/blog?charset=utf8"
	_DRIVER  = "mysql"
)

type Category struct {
	Id               int64
	Title            string
	Created          time.Time `orm:"index;type(timestamp);auto_now_add"`
	TopicTime        time.Time `orm:"index;type(timestamp);auto_now_add"`
	Views            int64     `orm:"index"`
	TopticCount      int64
	TopticLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:size(5000)`
	Attachment      string
	Created         time.Time `orm:"index;type(timestamp);auto_now_add"`
	Updated         time.Time `orm:"index;type(timestamp);auto_now_add"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegistrtDB() {
	// if !com.IsExist(_DB_NAME) {
	// 	os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
	// 	os.Create(_DB_NAME)
	// }

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_DRIVER, orm.DR_MySQL)
	// orm.RegisterDataBase("default", _DRIVER, _DB_NAME, 30)
	orm.RegisterDataBase("default", _DRIVER, _DB_NAME, 30)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	o.Using("default")
	cate := &Category{Title: name}
	qs := o.QueryTable("Category")
	err := qs.Filter("Title", name).One(cate)

	if err == nil {
		return nil
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	o.Using("default")
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {

	o := orm.NewOrm()
	o.Using("default")
	cates := make([]*Category, 0)
	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)
	return cates, err
}
