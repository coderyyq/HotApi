# CowApi

### 简述

基于 `Go` 语言实现的 `API` 接口，实时抓取各大新闻网站的热搜内容，并返回相关的详细信息，提供统一的 `API` 服务。

### 请求

```
GET http://127.0.0.1:9000/api/hot/{source}
```

### 参数

| 参数   | 类型   | 说明                                                                                                                                                 | 是否必填 |
| ------ | ------ | ---------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
| source | string | 来源平台标识，支持以下值：<br>`baidu`、`douyin`、`weibo`、`bilibili`、`kuaishou`、`pengpai`、`qqnews`、`sina`、`sougou`、`sspai`、`tieba`、`toutiao` | 是       |

### 响应

#### 成功：

```json
{
  "code": 1000,
  "data": {
    "source": "百度热搜",
    "update_time": "2025-06-16 09:09:50",
    "hot_list": [
      {
        "id": 1,
        "title": "总书记关心的世界文化遗产",
        "description": "近日，新华社推出系列微纪录片《总书记关心的世界文化遗产》。让我们一起行走在辽阔的神州大地上，回望古老的文化遗产，感受这幅古今辉映、气势恢宏的新时代大美画卷。",
        "picture": "https://fyb-2.cdn.bcebos.com/hotboard_image/0ca944f25e4ca42fc2ab6dce461c3762",
        "popularity": "7904246",
        "url": "https://www.baidu.com/s?wd=%E6%80%BB%E4%B9%A6%E8%AE%B0%E5%85%B3%E5%BF%83%E7%9A%84%E4%B8%96%E7%95%8C%E6%96%87%E5%8C%96%E9%81%97%E4%BA%A7"
      }
      // ...
    ]
  },
  "message": "success"
}
```

**注意**：

- 不是所有平台的返回数据都包含 `picture`、`description` 和 `popularity` 字段。

#### 失败：

```json
{
  "code": 1001,
  "data": {
    "source": "",
    "update_time": "",
    "hot_list": null
  },
  "message": "无法获取此源数据~"
}
```

### 支持来源平台

- `baidu`：百度
- `douyin`：抖音
- `weibo`: 微博
- `sina`：新浪
- `tieba`：贴吧
- `bilibili`：哔哩哔哩
- `kuaishou`：快手
- `pengpai`：澎湃
- `qqnews`：腾讯
- `sougou`：搜狗
- `sspai`：少数派
- `toutiao`：头条
