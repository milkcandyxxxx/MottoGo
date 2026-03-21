# [MottoGo](https://github.com/milkcandyxxxx/MottoGo)一言api使用文档

## 统一认证方式：

请求头：`X-API-Key`

权限分为admin（所以权限）与user（仅查看）

## 接口：

### /hitokoto

#### 方法：GET

#### 参数说明

| 参数名称 | 参数说明 | 请求类型 | 是否必须 | 数据类型 | 备注       |
| -------- | -------- | -------- | -------- | :------- | ---------- |
| id       | 图片id   | query    | 否       | uint     | 留空则随机 |
| t/type   | 图片类型 | query    | 否       | string   | 暂未添加   |
|          |          |          |          |          |            |



### /hitokoto/AddHit

#### 方法：POST

#### 数据类型：application/json

#### 格式说明：

```json
{
    "id": 2, // 填不填随意，id会自动追加
    "hitokoto": "树叶飞舞之处，火亦生生不息。",
    "from": "火影忍者",
    "from_who": "猿飞日斩",
    "type": "动漫"
}
```

### /hitokoto/DelHit

#### 方法：GET

#### 参数说明

| 参数名称 | 参数说明 | 请求类型 | 是否必须 | 数据类型 | 备注         |
| -------- | -------- | -------- | -------- | :------- | ------------ |
| id       | 图片id   | query    | 是       | uint     | 需要删除的id |

注：此方法删除后原数据id会乱，如对id有要求请谨慎删除