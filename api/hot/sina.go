package hot

import (
	"HotApi/model"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Sina struct {
}

func (*Sina) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://sinanews.sina.cn/h5/top_news_list.d.html")
	if err != nil {
		return model.HotSearchData{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.HotSearchData{}, err
	}

	var jsonStr string
	reg := regexp.MustCompile(`SM = ({.*?});`)
	result := reg.FindAllStringSubmatch(string(body), -1)
	if len(result) > 0 && len(result[0]) > 1 {
		jsonStr = result[0][1]
	} else {
		return model.HotSearchData{}, errors.New("获取数据失败~")
	}

	updateTime := time.Unix(gjson.Get(jsonStr, "data.data.date").Int(), 0).Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		if index := gjson.Get(jsonStr, "data.data.hotList."+strconv.Itoa(i)+".@type"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data.data.hotList."+strconv.Itoa(i)+".info.title").Str,
			Description: "",
			Picture:     "",
			Popularity:  gjson.Get(jsonStr, "data.data.hotList."+strconv.Itoa(i)+".info.hotValue").Str,
			URL:         gjson.Get(jsonStr, "data.data.hotList."+strconv.Itoa(i)+".base.base.url").Str,
		})
	}

	return model.HotSearchData{Source: "新浪热榜", UpdateTime: updateTime, HotList: hotList}, nil
}
