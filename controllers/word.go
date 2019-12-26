package controllers

import (
	"github.com/astaxie/beego"
	"recitewords/models"
)

type WordController struct {
	beego.Controller
	worddao *models.WordDao
}

func (this *WordController) AddWord() {
	date := this.GetString("date")
	raw := this.GetString("raw")
	explaination := this.GetString("explaination")
	element := models.WordElement{raw, explaination, nil}
	this.worddao = models.NewWordDao()
	
	if success, err := this.worddao.AddWord(&element, date); success == false {
		this.Data["json"] = err.Error()
	}else {
		this.Data["json"] = map[string]string{"method": "AddWord"}
	}
	this.ServeJSON()
}

func (this *WordController) GetWords() {
	raw := this.GetString("raw")
	date := this.GetString("date")
	this.worddao = models.NewWordDao()
	
	if explaination, err := this.worddao.GetWord(raw, date); err == nil {
		this.Data["json"] = map[string]string{"result": explaination.Explanation}
	}else {
		this.Data["json"] = map[string]string{"result": err.Error()}
	}
	this.ServeJSON()
}

func (this *WordController) DeleteWord() {
	raw := this.GetString("raw")
	this.worddao = models.NewWordDao()
	
	result, err := this.worddao.DeleteWord(raw)
	this.Data["json"] = map[string]interface{}{"result": result, "error": err}
	this.ServeJSON()
}


func (this *WordController) Export() {
	
	this.Data["json"] = map[string]string{"method":"Export"}
	this.ServeJSON()
}


func (this *WordController) Index() {
	date := this.GetString("date")
	this.worddao = models.NewWordDao()
	
	if list, err := this.worddao.GetWordsByDay(date); err != nil {
		this.TplName = "index.tpl"
	}else{
		this.Data["list"] = list.List
		this.TplName = "wordsindex.tpl"
	}
	
}

func (this *WordController) Config() {
	this.worddao = models.NewWordDao()
	this.Data["json"] = map[string]string{"from_worddao_echoconfig": this.worddao.EchoConfig(), "from_beego_appconfig_string": beego.AppConfig.String("redishost")}
	this.ServeJSON()
}