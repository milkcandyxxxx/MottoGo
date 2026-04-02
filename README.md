# [MottoGo](https://github.com/milkcandyxxxx/MottoGo)一言api使用文档

## 部署方法：

下载对应平台

https://github.com/milkcandyxxxx/MottoGo/releases

## 配置文件：

无配置文件时自动生成
基础模板

```yaml
server:
  port: 8080 # 运行端口
  allow_cors: true    # 是否开启跨域访问 (CORS)
security:
  require_userkey: false # 是否开启普通用户密钥，管理员为必须
  key:
    admin:
      - 415290769594460e2e485922904f345d  # 管理员密钥
    user:
      -   # 用户密钥
limit:
  rate: 10   # 每秒允许生成的令牌数（QPS）
  burst: 10  # 桶的最大容量（允许瞬间爆发的请求数）
```

## 限流器：

令牌桶算法

## 统一认证方式：

请求头：`Authorization` `Bearer 123`

权限分为admin（所有权限）与user（仅查看）

## 接口：

### /hitokoto（根据要求回去句子）

#### 方法：GET

#### 参数说明

| 参数名称 | 参数说明 | 请求类型 | 是否必须 | 数据类型 | 备注     |
| -------- | -------- | -------- | -------- | :------- | -------- |
| c        | 类型     | query    | 否       | string   | 留空则随 |
| a        | 作者     | query    | 否       | string   | 留空则随 |
| s        | 来源     | query    | 否       | string   |          |

注：三个选择可同时匹配

```
"动漫": "a",
"游戏": "b",
"文学": "c",
"原创": "e",
"来自网络": "f",
"诗句": "g",
"歌曲": "h",
"其他": "i",
```

#### 响应格式

```json
{
  "code": 200,
  "data": {
    "uuid": "bc3f077e-2d64-430a-82ac-44d6cb5c4e49",
    "hitokoto": "未知",
    "source": "未知",
    "author": "佚名",
    "category": "a",
    "created_at": "2026-03-25T15:59:34.4846072+08:00"
  },
  "msg": "查询成功"
}
```

### /hitokoto/all（获取所有句子）

#### 方法：POST

#### 数据类型：application/json

#### 响应格式

```json
{
  "code": 200,
  "data": [
    {
      "uuid": "e0c613d3-6d2c-41e2-9bcd-96ea68636c8b",
      "hitokoto": "年少不可得之物，终将困其一生。",
      "source": "网络感悟",
      "author": "佚名",
      "category": "f",
      "created_at": "2026-04-01T20:00:04.810443286+08:00"
    },
    {
      "uuid": "3e33365f-3510-4358-8530-d7e41d1e7fe0",
      "hitokoto": "最是人间留不住，朱颜辞镜花辞树。",
      "source": "蝶恋花·阅尽天涯离别苦",
      "author": "王国维",
      "category": "g",
      "created_at": "2026-04-01T20:01:25.941270554+08:00"
    },
    {
      "uuid": "aeb20645-1a59-4027-b52b-06ea9ef62f5c",
      "hitokoto": "我深知生命如蜉蝣，深知死亡总是如影随形。但此时哪怕再多一年，再多一日，再多一时也好，我辈仍愿人生得续。",
      "source": "铃芽之旅",
      "author": "岩户铃芽",
      "category": "a",
      "created_at": "2026-04-01T20:11:52.265668153+08:00"
    },
    {
      "uuid": "c2503810-79e7-4c26-a7e5-8bba55aebbad",
      "hitokoto": "世界上有两样东西不可直视，一是太阳，二是人心。",
      "source": "白夜行",
      "author": "东野圭吾",
      "category": "c",
      "created_at": "2026-04-01T20:12:24.284305422+08:00"
    },
    {
      "uuid": "2be66aa6-5079-49e8-b4e5-33e2863a4a24",
      "hitokoto": "我见青山多妩媚，料青山见我应如是。",
      "source": "贺新郎·甚矣吾衰矣",
      "author": "辛弃疾",
      "category": "g",
      "created_at": "2026-04-01T20:13:07.074925724+08:00"
    }
  ],
  "msg": "查询成功"
}
```



### /hitokoto/AddHit（添加句子）

#### 方法：POST

#### 数据类型：application/json

#### 格式说明：

```json
{
  "hitokoto": "最是人间留不住，朱颜辞镜花辞树。",
  "source": "阅微草堂笔记 / 蝶恋花",
  "author": "王国维",
  "category": "诗词"
}
```

#### 响应格式

```json
{
  "code": 201,
  "data": {
    "uuid": "e85770cc-ebb7-4ebb-b739-81bc603754c4",
    "hitokoto": "最是人间留不住，朱颜辞镜花辞树。",
    "source": "阅微草堂笔记 / 蝶恋花",
    "author": "王国维",
    "category": "诗词",
    "created_at": "2026-03-25T17:31:39.5615739+08:00"
  },
  "msg": "添加成功"
}
```

### /hitokoto/DelHit/:uuid（删除句子）

#### 方法：DELETE

#### 参数说明

| 参数名称 | 参数说明 | 请求类型 | 是否必须 | 数据类型 | 备注           |
| -------- | -------- |------| -------- | :------- | -------------- |
| uuid     | 图片uuid | path | 是       | string   | 需要删除的uuid |

#### 响应格式

```json
{
  "code": 200,
  "data": {
    "delete_count": 0,
    "delete_uuid": "25e8caed-3e62-40fc-9cf0-6993e50597f2"
  },
  "msg": "删除成功"
}
```

## 数据库格式

| **字段名**    | **类型**    | **说明**   | **约束/特性**              |
| ------------- | ----------- | ---------- | -------------------------- |
| **ID**        | `uint`      | 主键 ID    | 自增，JSON 序列化时隐藏    |
| **Uuid**      | `string`    | 唯一标识符 | 长度 36，非空，唯一索引    |
| **Hitokoto**  | `string`    | 语句正文   | 数据库类型为 `TEXT`，非空  |
| **Source**    | `string`    | 来源       | 默认值：'未知'，已建立索引 |
| **Author**    | `string`    | 作者       | 默认值：'佚名'，已建立索引 |
| **Category**  | `string`    | 分类标签   | 长度 20，非空，已建立索引  |
| **CreatedAt** | `time.Time` | 创建时间   | 自动填充创建时间           |
