package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

type Douyin struct {
}

func (*Douyin) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	resp, err := http.Get("https://www.iesdouyin.com/web/api/v2/hotsearch/billboard/word/")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	jsonStr := string(body)

	updateTime := gjson.Get(jsonStr, "active_time").Str

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "word_list."+strconv.Itoa(i)+".word"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "word_list."+strconv.Itoa(i)+".word").Str,
			Description: "",
			Picture:     "",
			Popularity:  strconv.FormatInt(gjson.Get(jsonStr, "word_list."+strconv.Itoa(i)+".hot_value").Int(), 10),
			URL:         "https://www.douyin.com/search/" + gjson.Get(jsonStr, "word_list."+strconv.Itoa(i)+".word").Str,
		})
	}

	return model.HotSearchData{Source: "抖音热榜", UpdateTime: updateTime, HotList: hotList}, nil
}
