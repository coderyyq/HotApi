package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type Weibo struct {
}

func (*Weibo) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://m.weibo.cn/api/container/getIndex?containerid=106003type%3D25%26t%3D3%26disable_hot%3D1%26filter_type%3Drealtimehot")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	updateTime := time.Unix(gjson.Get(jsonStr, "data.cardlistInfo.starttime").Int(), 0).Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "data.cards.0.card_group."+strconv.Itoa(i)+".itemid"); !index.Exists() {
			break
		}
		popularity := gjson.Get(jsonStr, "data.cards.0.card_group."+strconv.Itoa(i)+".desc_extr").Raw
		if strings.Contains(popularity, `\u`) {
			popularity, _ = strconv.Unquote(popularity)
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data.cards.0.card_group."+strconv.Itoa(i)+".desc").Str,
			Description: "",
			Picture:     "",
			Popularity:  popularity,
			URL:         gjson.Get(jsonStr, "data.cards.0.card_group."+strconv.Itoa(i)+".scheme").Str,
		})
	}

	return model.HotSearchData{Source: "微博热搜", UpdateTime: updateTime, HotList: hotList}, nil
}
