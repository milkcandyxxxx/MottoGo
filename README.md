# [MottoGo](https://github.com/milkcandyxxxx/MottoGo)一言api使用文档

## 部署方法：

下载对应平台，将MottoGo，config.yaml，cartoon.jsonl放置于同一文件下即可。

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

请求头：`X-API-key`

权限分为admin（所有权限）与user（仅查看）

## 接口：

### /hitokoto

#### 方法：GET

#### 参数说明

| 参数名称 | 参数说明     | 请求类型 | 是否必须 | 数据类型 | 备注       |
| -------- | ------------ | -------- | -------- | :------- | ---------- |
| c        | 类型         | query    | 否       | uint     | 留空则随机 |
| all      | 是否输出全部 | query    | 否       | bool     |            |

```
"动漫": "a",
"游戏": "b",
"文学": "c",
"原创": "e",
"佚名": "f",
"其他": "g",
"诗句": "h",
```

#### 响应格式

```json
{
  "uuid": "4f5a5052-76c0-41eb-abef-fc903cc8978b",
  "hitokoto": "最是人间留不住，朱颜辞镜花辞树",
  "from": "《蝶恋花·阅尽天涯离别苦》",
  "from_who": "王国维",
  "type": "h"
}
```

### /hitokoto/AddHit

#### 方法：POST

#### 数据类型：application/json

#### 格式说明：

```json
{
    "hitokoto": "树叶飞舞之处，火亦生生不息。",
    "from": "火影忍者",
    "from_who": "猿飞日斩",
    "type": "a"
}
```

#### 响应格式

```json
{
  "data": {
    "uuid": "257e1915-3716-4fe4-bc37-fafa8ab66d98",
    "hitokoto": "最是人间留不住，朱颜辞镜花辞树",
    "from": "《蝶恋花·阅尽天涯离别苦》",
    "from_who": "王国维",
    "type": "h"
  },
  "ok": "Write success"
}
```

### /hitokoto/DelHit

#### 方法：DELETE

#### 参数说明

| 参数名称 | 参数说明 | 请求类型 | 是否必须 | 数据类型 | 备注           |
| -------- | -------- |------| -------- | :------- | -------------- |
| uuid     | 图片uuid | path | 是       | string   | 需要删除的uuid |

#### 响应格式

```json
{
  "hitokoto": "最是人间留不住，朱颜辞镜花辞树",
  "ok": "Delete successful"
}
```
