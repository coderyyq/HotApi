package hot

import (
	"HotApi/model"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

type Kuaishou struct {
}

func (*Kuaishou) GetHotSearchData(maxNum int) (HotSearchData model.HotSearchData, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.kuaishou.com/?isHome=1", nil)
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

	var jsonStr string
	reg := regexp.MustCompile(`window.__APOLLO_STATE__=({.*?});`)
	result := reg.FindAllStringSubmatch(string(body), -1)
	if len(result) > 0 && len(result[0]) > 1 {
		jsonStr = result[0][1]
	} else {
		return model.HotSearchData{}, errors.New("获取数据失败~")
	}

	updateTime := time.Now().Format("2006-01-02 15:04:05")

	var hotList []model.HotItem
	for i := 0; i < maxNum; i++ {
		index := gjson.Get(jsonStr, `defaultClient.$ROOT_QUERY\.visionHotRank({\"page\":\"home\"}).items.`+strconv.Itoa(i)+".id")
		if !index.Exists() {
			break
		}
		result := escapeSpecialCharacters(index.Str)
		hotList = append(hotList, model.HotItem{
			Id:          int(gjson.Get(jsonStr, "defaultClient."+result+".rank").Int() + 1),
			Title:       gjson.Get(jsonStr, "defaultClient."+result+".name").Str,
			Description: "",
			Picture:     gjson.Get(jsonStr, "defaultClient."+result+".poster").Str,
			Popularity:  gjson.Get(jsonStr, "defaultClient."+result+".hotValue").Str,
			URL:         "https://www.kuaishou.com/short-video/" + gjson.Get(jsonStr, "defaultClient."+result+".photoIds.json.0").Str,
		})
	}

	return model.HotSearchData{Source: "快手热榜", UpdateTime: updateTime, HotList: hotList}, nil
}

func escapeSpecialCharacters(str string) string {
	var result strings.Builder

	for _, char := range str {
		if char == '.' {
			result.WriteRune('\\')
		}
		result.WriteRune(char)
	}

	return result.String()
}
