package hot

import (
	"HotApi/model"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

type Pengpai struct {
}

func (*Pengpai) GetHotSearchData(maxNum int) (model.HotSearchData, error) {
	resp, err := http.Get("https://cache.thepaper.cn/contentapi/wwwIndex/rightSidebar")
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
		if index := gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".contId"); !index.Exists() {
			break
		}
		hotList = append(hotList, model.HotItem{
			Id:          i + 1,
			Title:       gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".name").Str,
			Description: gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".nodeInfo.summarize").Str,
			Picture:     gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".smallPic").Str,
			Popularity:  gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".pubTimeNew").Str,
			URL:         "https://www.thepaper.cn/newsDetail_forward_" + gjson.Get(jsonStr, "data.hotNews."+strconv.Itoa(i)+".contId").Str,
		})
	}

	return model.HotSearchData{Source: "澎湃热榜", UpdateTime: updateTime, HotList: hotList}, nil
}
