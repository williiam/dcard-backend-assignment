## DCARD 2022 backend-intern assignment
#### DEMO:

https://dacrd-backend-assignment.herokuapp.com/

## Installation and Setup Instructions

Clone down this repository. You will need `node` and `npm` installed globally on your machine.  

Installation:

`go install`  

 安裝並啟動postgres，並到/DB/db.go裡設定連線參數

 安裝並啟動redis，並到/service/redis.go裡設定連線參數

To Start Server:

`cd server`

`go run server`  


## Routes
### `localhost:8080/`:
Send a GET request and get a warm welcome :)
```JSON
{
  "message": "Welcome to URL shortener."
}
```

### `localhost:8080/shorten`:
POST a JSON object like below and in return, get the generated short link:
```JSON
{
  "LongURL": "https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716",
  "ExpTime": "2d"
}
```
Exp date valid format: `2d` for 2 days, `2h` for 2 hours, `2m` for 2 minutes and `2s` for 2 seconds.

#### Result:
```JSON
{
  "message": "Short url created successfully",
  "ShortURL": "http://localhost:8080/ZLgJHJB2"
}
```

## 架構設計說明

採用postgres作為RDBMS:


採用redis做為快取：

使用redis還有以下好處：
減輕資料庫負擔
加快server回應速度，因為不知要連線資料庫



RDBMS - postgres:
  
  只有一張table
  結構如下：

    url_mapping:
      id: PK
      short_url: 短網址
      original_url: 長網址
      count: 被使用計數


redis快取機制如下:

剛新增的短網址會存在快取裡6小時

若有url_mapping被使用次數>10 並且不在快取 則會將之加入快取


