package routers

import (
	"recitewords/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    // 根据日期获取某一天的单词记录表
    beego.Router("/words/index", &controllers.WordController{}, "*:Index")
    // 存储当天的单词记录
    beego.Router("/words/add", &controllers.WordController{}, "*:AddWord")
    // 根据日期查看某天的单词记录
    beego.Router("/words/watch", &controllers.WordController{}, "*:GetWords")
    // 导出单词记录 可选维度 日、周、月、年
    beego.Router("/words/export", &controllers.WordController{}, "*:Export")
    // 删除当日、指定日期某个单词
    beego.Router("/words/delete", &controllers.WordController{}, "*:DeleteWord")
    
    // TODO 待删除的一波测试
    beego.Router("/words/config", &controllers.WordController{}, "*:Config")
}
