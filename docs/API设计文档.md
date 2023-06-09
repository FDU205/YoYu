---
title: 幽语App v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# 幽语App

> v1.0.0

Base URLs:

# User Module

<a id="opIduserLogin"></a>

## POST 用户登录

POST /api/user/login

用户登录

> Body 请求参数

```json
{
  "username": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 1,
  "error_msg": "string",
  "data": {
    "token": "string",
    "id": 0,
    "username": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|A status code of 0 means the request is successful; others mean errors, if there is an error, please attach the error messages.|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|
|»» token|string|true|none||none|
|»» id|integer|true|none||none|
|»» username|string|true|none||none|

## POST 用户注册

POST /api/user/register

> Body 请求参数

```json
{
  "username": "string",
  "password": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "username": "string",
    "token": "string",
    "id": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|If register successfully, return the user information in response data.|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» data|object|true|none||none|
|»» username|string|true|none||none|
|»» token|string|true|none||none|
|»» id|integer|true|none||none|

## POST 关注用户

POST /api/user/follow

> Body 请求参数

```json
{
  "follow_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» follow_id|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|

## DELETE 取消关注

DELETE /api/user/unfollow

> Body 请求参数

```json
{
  "follow_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» follow_id|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|

## GET 获取关注列表

GET /api/user/followlist

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "follows": [
      {
        "user_id": 0,
        "username": "string"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» data|object|true|none||none|
|»» follows|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» username|string|true|none||none|

## GET 获取粉丝列表

GET /api/user/fanslist

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "fans": [
      {
        "user_id": 0,
        "username": "string"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» data|object|true|none||none|
|»» fans|[object]|true|none||none|
|»»» user_id|integer|true|none||none|
|»»» username|string|true|none||none|

## GET 获取关注数

GET /api/user/followcount

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» count|integer|true|none||none|

## GET 获取粉丝数

GET /api/user/fanscount

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "count": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» count|integer|true|none||none|

## GET 是否关注

GET /api/user/isfollow

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|follow_id|query|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "yes": true
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» error_msg|string|true|none||none|
|» yes|boolean|true|none||none|

# MessageBox Module

## POST 创建提问箱

POST /api/messageBox

> Body 请求参数

```json
{
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "owner_id": 0,
    "title": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|[MessageBox](#schemamessagebox)|false|none||none|
|»» id|integer|false|none||none|
|»» owner_id|integer|false|none||none|
|»» title|string|false|none||none|

## GET 获取提问箱信息

GET /api/messageBox/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "owner_id": 0,
    "title": "string",
    "owner_name": "string",
    "posts": [
      0
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» owner_id|integer|true|none||none|
|»» title|string|true|none||none|
|»» owner_name|string|true|none||none|
|»» posts|[integer]|true|none||none|

## PUT 修改提问箱信息

PUT /api/messageBox/{id}

> Body 请求参数

```json
{
  "title": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "owner_id": 0,
    "title": "string"
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|[MessageBox](#schemamessagebox)|false|none||none|
|»» id|integer|false|none||none|
|»» owner_id|integer|false|none||none|
|»» title|string|false|none||none|

## DELETE 删除提问箱

DELETE /api/messageBox/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |Only the owner of library can delete it.|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|

## GET 查询提问箱

GET /api/messageBoxes

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|title|query|string| 否 |none|
|owner|query|integer| 否 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "messageBoxes": [
      {
        "id": 0,
        "owner_id": 0,
        "title": "string",
        "owner_name": "string"
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|true|none||none|
|»» messageBoxes|[object]|true|none||none|
|»»» id|integer|false|none||none|
|»»» owner_id|integer|false|none||none|
|»»» title|string|false|none||none|
|»»» owner_name|string|true|none||none|

# Post Module

## POST 创建帖子

POST /api/post

> Body 请求参数

```json
{
  "message_box_id": 0,
  "content": "string",
  "visibility": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» message_box_id|body|integer| 是 |none|
|» content|body|string| 是 |none|
|» visibility|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "poster_id": 0,
    "message_box_id": 0,
    "content": "string",
    "visibility": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|[Post](#schemapost)|false|none||none|
|»» id|integer|false|none||none|
|»» poster_id|integer|false|none||none|
|»» message_box_id|integer|false|none||none|
|»» content|string|false|none||none|
|»» visibility|integer|false|none||none|

## GET 查询所有帖子

GET /api/posts

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|message_box_id|query|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "posts": [
      {
        "id": 0,
        "poster_id": 0,
        "message_box_id": 0,
        "content": "string",
        "visibility": 0
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|
|»» posts|[[Post](#schemapost)]|true|none||none|
|»»» id|integer|false|none||none|
|»»» poster_id|integer|false|none||none|
|»»» message_box_id|integer|false|none||none|
|»»» content|string|false|none||none|
|»»» visibility|integer|false|none||none|

## GET 获取帖子信息

GET /api/post/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "poster_id": 0,
    "poster_name": "string",
    "content": "string",
    "visibility": 0,
    "message_box_id": 0,
    "threads": [
      {
        "id": 0,
        "post_id": 0,
        "content": "string",
        "type": 0
      }
    ],
    "channels": [
      "string"
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|
|»» id|integer|false|none||none|
|»» poster_id|integer|false|none||none|
|»» poster_name|string|true|none||none|
|»» content|string|false|none||none|
|»» visibility|integer|false|none||none|
|»» message_box_id|integer|true|none||none|
|»» threads|[[Channel](#schemachannel)]|true|none||none|
|»»» id|integer|true|none||none|
|»»» post_id|integer|true|none||none|
|»»» content|string|true|none||none|
|»»» type|integer|true|none||none|
|»» channels|[string]|false|none||none|

## DELETE 删除帖子

DELETE /api/post/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |Only the owner of library can delete it.|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {}
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|

## POST 创建channel

POST /api/post/channel

> Body 请求参数

```json
{
  "post_id": 0,
  "content": "string",
  "type": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» post_id|body|integer| 是 |none|
|» content|body|string| 是 |none|
|» type|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "post_id": 0,
    "content": "string",
    "type": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|[Channel](#schemachannel)|false|none||none|
|»» id|integer|true|none||none|
|»» post_id|integer|true|none||none|
|»» content|string|true|none||none|
|»» type|integer|true|none||none|

## GET 我的贴子

GET /api/mypost

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "id": 0,
    "poster_id": 0,
    "message_box_id": 0,
    "content": "string",
    "visibility": 0
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|[Post](#schemapost)|false|none||none|
|»» id|integer|false|none||none|
|»» poster_id|integer|false|none||none|
|»» message_box_id|integer|false|none||none|
|»» content|string|false|none||none|
|»» visibility|integer|false|none||none|

# Wall Module

## GET 获取今日表白墙

GET /api/wall

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 是 |none|
|page_size|query|integer| 是 |none|
|Authorization|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string",
  "data": {
    "posts": [
      {
        "id": 0,
        "poster_id": 0,
        "poster_name": "string",
        "content": "string",
        "visibility": 0
      }
    ]
  }
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|
|» data|object|false|none||none|
|»» posts|[object]|false|none||none|
|»»» id|integer|false|none||none|
|»»» poster_id|integer|false|none||none|
|»»» poster_name|string|false|none||none|
|»»» content|string|false|none||none|
|»»» visibility|integer|false|none||none|

## POST 创建表白信息

POST /api/wall/create

> Body 请求参数

```json
{
  "content": "string",
  "visibility": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|
|body|body|object| 否 |none|
|» content|body|string| 是 |none|
|» visibility|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "error_msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|false|none||none|
|» error_msg|string|false|none||none|

## GET 我的表白墙

GET /api/wall/mywall

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_size|query|integer| 是 |none|
|page_num|query|integer| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

<h2 id="tocS_Follow">Follow</h2>

<a id="schemafollow"></a>
<a id="schema_Follow"></a>
<a id="tocSfollow"></a>
<a id="tocsfollow"></a>

```json
{
  "follow_id": 0,
  "follow_by_id": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|follow_id|integer|true|none||none|
|follow_by_id|integer|true|none||none|

<h2 id="tocS_Channel">Channel</h2>

<a id="schemachannel"></a>
<a id="schema_Channel"></a>
<a id="tocSchannel"></a>
<a id="tocschannel"></a>

```json
{
  "id": 0,
  "post_id": 0,
  "content": "string",
  "type": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|post_id|integer|true|none||none|
|content|string|true|none||none|
|type|integer|true|none||none|

<h2 id="tocS_Wall">Wall</h2>

<a id="schemawall"></a>
<a id="schema_Wall"></a>
<a id="tocSwall"></a>
<a id="tocswall"></a>

```json
{
  "id": 0,
  "poster_id": 0,
  "content": "string",
  "visibility": 1,
  "date": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|true|none||none|
|poster_id|integer|true|none||none|
|content|string|true|none||none|
|visibility|integer|true|none||2：匿名|
|date|string|true|none||none|

<h2 id="tocS_MessageBox">MessageBox</h2>

<a id="schemamessagebox"></a>
<a id="schema_MessageBox"></a>
<a id="tocSmessagebox"></a>
<a id="tocsmessagebox"></a>

```json
{
  "id": 0,
  "owner_id": 0,
  "title": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||none|
|owner_id|integer|false|none||none|
|title|string|false|none||none|

<h2 id="tocS_Post">Post</h2>

<a id="schemapost"></a>
<a id="schema_Post"></a>
<a id="tocSpost"></a>
<a id="tocspost"></a>

```json
{
  "id": 0,
  "poster_id": 0,
  "message_box_id": 0,
  "content": "string",
  "visibility": 0
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||none|
|poster_id|integer|false|none||none|
|message_box_id|integer|false|none||none|
|content|string|false|none||none|
|visibility|integer|false|none||none|

<h2 id="tocS_User">User</h2>

<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": 0,
  "username": "string",
  "password": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer|false|none||none|
|username|string|false|none||none|
|password|string|false|none||none|

