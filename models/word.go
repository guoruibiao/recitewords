package models

import (
	"recitewords/utils"
			"recitewords/callbacks"
	"github.com/pkg/errors"
		"github.com/astaxie/beego"
		"github.com/garyburd/redigo/redis"
)

const (
	WORDSKEY_PREXIX = "wordshash";
)


type WordElement struct {
	Raw string
	Explanation string
	Extra map[string]interface{}
}

type WordList struct {
	List []WordElement
}

type WordDao struct {
	redisutil *utils.RedisUtil
}

func NewWordDao() *WordDao {
	redisconfig := utils.GetRedisConfig(
		"tcp",
		beego.AppConfig.String("redishost") + ":" + beego.AppConfig.String("redisport"),
		)
	return &WordDao{
		utils.NewRedisUtil(*redisconfig),
	}
}

func (this *WordDao) AddWord(wordElement *WordElement, date string) (result bool, err error) {
	params := []interface{}{WORDSKEY_PREXIX + date, wordElement.Raw, wordElement.Explanation}
	_, err = this.redisutil.Run("hset", callbacks.DoCommandImpl, params...)
	if err != nil {
		return false, errors.New(err.Error())
	}
	return true, nil
}

func (this *WordDao) GetWord(raw, date string) (wordElement *WordElement, err error) {
	params := []interface{}{WORDSKEY_PREXIX + date, raw}
	result, err := redis.String(this.redisutil.Run("hget", callbacks.DoCommandImpl, params...))
	if err != nil {
		return nil, errors.New("未找到【" + raw +"】的翻译结果")
	}
	wordElement = &WordElement{}
	wordElement.Raw = raw
	wordElement.Explanation = result
	return
}

func (this *WordDao) GetWordsByDay(date string) (wordList *WordList, err error) {
	// TODO date 格式可能不尽然会按照方法要求的那样
	params := []interface{}{WORDSKEY_PREXIX + date}
	maps, err := redis.StringMap(this.redisutil.Run("hgetall", callbacks.DoCommandImpl, params...))
	if err != nil {
		return nil, err
	}
	list := &WordList{}
	for raw, explination := range maps {
		word := &WordElement{
			Raw:raw,
			Explanation: explination,
		}
		list.List = append(list.List, *word)
	}
	return list, nil
}


// 删除某一天的没啥必要 仅支持当天的吧
func (this *WordDao) DeleteWord(raw string) (success bool, err error) {
	params := []interface{}{WORDSKEY_PREXIX + utils.DateString("YYYY-mm-dd", 0), raw}
	if _, err = this.redisutil.Run("hdel", callbacks.DoCommandImpl, params...); err == nil {
		return true, nil
	}
	return false, err
}

func (this *WordDao) EchoConfig() string{
	host := beego.AppConfig.String("redishost")
	return host
}