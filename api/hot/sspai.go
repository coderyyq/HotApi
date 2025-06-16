package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Sspai struct {
}

func (*Sspai) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	resp, err := http.Get("https://sspai.com/api/v1/article/tag/page/get?limit=40&offset=0&tag=%E7%83%AD%E9%97%A8%E6%96%87%E7%AB%A0&released=false")
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
		if index := gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".id"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".title").Str,
			Description: gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".summary").Str,
			Picture:     "",
			Popularity:  "",
			URL:         "https://sspai.com/post/" + strconv.FormatInt(gjson.Get(jsonStr, "data."+strconv.Itoa(i)+".id").Int(), 10),
		})
	}

	return model.HotSearchData{Source: "少数派最热", UpdateTime: updateTime, HotList: hotList}, nil
}
