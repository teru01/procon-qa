# procon-qa
https://procon-qa.herokuapp.com/

## フロントエンド
- Vue
- TypeScript
- Vuetify
## バックエンド
- Go
- echo
- GORM


### 検証環境 (docker-compose)
docker-composeを使います。
```
git clone https://github.com/monkukui/procon-qa.git
cd procon-qa
docker-compose build
docker-compose up
```
としてから `localhost:8080` にアクセスしてください。
検証が終わったら `docker-compose down` してください。

### DB 設計
```
Questions {
  id: int ( primary_key )
  uid: int
  tid: int
  
  title: string
  body: string
  date: string
  state: string
  url: string
  completed: boolean
}

Answers {
  id: int ( primary_key )
  uid: int
  qid: int

  body: string
  date: string
}

Users {
  uid: int (primary_key)
  name: string
  pass: string (仮)
  rate: int (仮)
}

QuestionGood {
  uid: number
  qid: number
}

AnswerGood {
  uid: number
  aid: number
}
```

### API

## GET
質問を全取得する (GET)
```
/api/questions

[
  {
    "uid":1,
    "tid":2,
    "id":3,
    "title":"タイトル1",
    "body":"本文1",
    "url":"https://atcoder/abc123/proba",
    "state":"",
    "date":"1996/05/26",
    "completed":true
  },
  {
    "uid":1,
    "tid":0,
    "id":4,
    "title":"aa",
    "body":"bb",
    "url":"",
    "state":"",
    "date":"",
    "completed":false
  },
]
```

質問をページ取得する (GET)
```
/api/questions/<page(1-indexed)>

[
  {
    "uid":1,
    "tid":2,
    "id":3,
    "title":"タイトル1",
    "body":"本文1",
    "url":"https://atcoder/abc123/proba",
    "state":"",
    "date":"1996/05/26",
    "completed":true
  },
  {
    "uid":1,
    "tid":0,
    "id":4,
    "title":"aa",
    "body":"bb",
    "url":"",
    "state":"",
    "date":"",
    "completed":false
  },
]
```

ユーザーに紐づいた質問をページ取得する (GET)
```
/api/user-questions/<page(1-indexed)>

[
  {
    "uid":1,
    "tid":2,
    "id":3,
    "title":"タイトル1",
    "body":"本文1",
    "url":"https://atcoder/abc123/proba",
    "state":"",
    "date":"1996/05/26",
    "completed":true
  },
  {
    "uid":1,
    "tid":0,
    "id":4,
    "title":"aa",
    "body":"bb",
    "url":"",
    "state":"",
    "date":"",
    "completed":false
  },
]
```

質問を 1 つ取得する (GET)
```
/api/question/<question_id>
{
  "uid":1,
  "tid":2,
  "id":3,
  "title":"タイトル1",
  "body":"本文1",
  "url":"https://atcoder/abc123/proba",
  "state":"",
  "date":"1996/05/26",
  "completed":true
},
```
回答を 1 つ取得する (GET)
```
/api/answer/<answer_id>
{
  "uid":4,
  "qid":4,
  "id":12,
  "body":"自明では",
  "date":"yyyy/mm/dd-hh/mm/ss",
  "favo":0
}

```


ユーザー 1 つを取得する (GET)
```
/api/user/<user_id>
{
  "id":3,
  "name":"monkukui",
  "password":"mon"
}
```

特定の質問に対する, 回答の全取得 (GET)
```
/api/answers/:qid
[ 
  { 
    "uid": 3, 
    "qid": 4, 
    "id": 3, 
    "body": "自明だと思う", 
    "date": "yyyy/mm/dd-hh/mm/ss", 
    "favo": 0 
  }, 
  { 
    "uid": 3, 
    "qid": 4, 
    "id": 4, 
    "body": "それは自明では？", 
    "date": "yyyy/mm/dd-hh/mm/ss",
    "favo": 0 
  }
]

```


## POST
質問を投稿する (POST)
```
/api/questions
{
  title: "タイトル"
  body: "本文"
  date: "yyyy/mm/dd-hh:mm:ss"
  state: "WA"
  url: "https://atcoder.jp/contests/jsc2019-qual/submissions/7120555"
}
```
回答を投稿する (POST)
```
/api/answers
{
  body: "本文"
  date: "yyyy/mm/dd-hh:mm:ss"
  qid: 3
  favo: 0
}
```

## DELETE
質問を削除する (DELETE)
自分の投稿じゃない場合, 404 がかえる

```
/api/question/<question_id>
```

## PUT
質問の解決ヅミフラグを切り替え
自分の質問じゃない場合, 404 が帰る
```
/api/question/<question_id>/completed
```


### TODO

#### 優先度高

- [x] (GET) 質問の全取得
- [x] (GET) 質問の固定長取得 (10 個ほど)
- [x] (GET) 質問を 1 つ取得
- [x] (GET) 回答を 1 つ取得
- [ ] (GET) 未回答の質問を全取得
- [x] (GET) 特定のユーザが投稿した質問の全取得
- [x] (GET) 特定の質問に対する, 回答の全取得
- [x] (GET) ユーザを 1 つ取得
- [x] (POST) 質問の投稿
- [x] (POST) 回答の投稿
- [ ] (PUT) 質問の編集
- [ ] (PUT) 回答の編集
- [X] (DELETE) 質問の削除
- [ ] (DELETE) 回答の削除

####  優先度中

- [ ] (GET) タグ検索

#### 優先度低
- [ ] (POST) ファボ機能
- [ ] (PUT) レート変動
