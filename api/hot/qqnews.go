package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type QQnews struct {
}

func (*QQnews) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://i.news.qq.com/gw/event/pc_hot_ranking_list?offset=0&page_size=51")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	updateTime := time.Now().Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 1; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i)+".id"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i,
			Title:       gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i)+".title").Str,
			Description: gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i)+".abstract").Str,
			Picture:     "",
			Popularity:  "",
			URL:         gjson.Get(jsonStr, "idlist.0.newslist."+strconv.Itoa(i)+".url").Str,
		})
	}

	return model.HotSearchData{Source: "腾讯热点榜", UpdateTime: updateTime, HotList: hotList}, nil
}
