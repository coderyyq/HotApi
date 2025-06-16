package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Bilibili struct {
}

func (*Bilibili) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.bilibili.com/x/web-interface/ranking/v2?type=all", nil)
	if err != nil {
		return model.HotSearchData{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
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
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".aid"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".title").Str,
			Description: gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".desc").Str,
			Picture:     gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".pic").Str,
			Popularity:  strconv.FormatInt(gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".stat.view").Int(), 10),
			URL:         gjson.Get(jsonStr, "data.list."+strconv.Itoa(i)+".short_link_v2").Str,
		})
	}

	return model.HotSearchData{Source: "bilibili排行榜", UpdateTime: updateTime, HotList: hotList}, nil
}
