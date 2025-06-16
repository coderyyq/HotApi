package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Tieba struct {
}

func (*Tieba) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	resp, err := http.Get("https://tieba.baidu.com/hottopic/browse/topicList")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	updateTime := time.UnixMilli(gjson.Get(jsonStr, "data.timestamp").Int()).Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		index := gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".idx_num")
		if !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          int(index.Int()),
			Title:       gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".topic_name").Str,
			Description: gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".topic_desc").Str,
			Picture:     gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".topic_pic").Str,
			Popularity:  strconv.FormatInt(gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".discuss_num").Int(), 10),
			URL:         gjson.Get(jsonStr, "data.bang_topic.topic_list."+strconv.Itoa(i)+".topic_url").Str,
		})
	}

	return model.HotSearchData{Source: "贴吧热议榜", UpdateTime: updateTime, HotList: hotList}, nil
}
