package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

type Toutiao struct {
}

func (*Toutiao) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://www.toutiao.com/hot-event/hot-board/?origin=toutiao_pc")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	time := gjson.Get(jsonStr, "impr_id").Str[:14]
	updateTime := time[:4] + "-" + time[4:6] + "-" + time[6:8] + " " + time[8:10] + ":" + time[10:12] + ":" + time[12:14]

	var hotList []model.HotItem

	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".ClusterId"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".Title").Str,
			Description: "",
			Picture:     gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".Image.url").Str,
			Popularity:  gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".HotValue").Str,
			URL:         gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".Url").Str,
		})
	}

	return model.HotSearchData{Source: "头条热榜", UpdateTime: updateTime, HotList: hotList}, nil
}
