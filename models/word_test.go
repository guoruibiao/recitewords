package models

import (
	"testing"
	"github.com/astaxie/beego"
)

var dao *WordDao
func init() {
	dao = NewWordDao()
}


func TestWordDao_AddWord(t *testing.T) {
	wordElement := WordElement{
		Raw:"test",
		Explanation: "测试",
	}
	t.Log(beego.AppConfig.String("redishost"))
	if success, err := dao.AddWord(&wordElement); success == false {
		t.Error(err)
	}else{
		t.Log("add word success", success)
	}
}

func TestWordDao_EchoConfig(t *testing.T) {
	t.Log("获取结果为 [" + dao.EchoConfig() +"]")
}