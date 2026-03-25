# [MottoGo](https://github.com/milkcandyxxxx/MottoGo)一言api使用文档

## 部署方法：

下载对应平台，将MottoGo，config.yaml放置于同一文件下即可。

https://github.com/milkcandyxxxx/MottoGo/releases

## 配置文件：

密码均为MD5值目前需自己加密填入

```yaml
server:
  port: 8080 # 端口
security:
  key:
    admin: # 管理员
      - 202cb962ac59075b964b07152d234b70 
    user: # 用户
      - a2550eeab0724a691192ca13982e6ebd
      - c4ca4238a0b923820dcc509a6f75849b
      - 76419c58730d9f35de7ac538c2fd6737
```

## 限流器：

每个ip一秒内最多访问10次

## 统一认证方式：

请求头：`Authorization` `Bearer 123`

权限分为admin（所有权限）与user（仅查看）

## 接口：

### /hitokoto

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

### /hitokoto/AddHit

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

### /hitokoto/DelHit/:uuid

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
