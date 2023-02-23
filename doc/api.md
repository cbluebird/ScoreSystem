---
title: 综测平台 v1.0.0
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

# 综测平台

> v1.0.0

Base URLs:

# 用户操作

## POST 修改密码

POST /api/user/bind/repassword

> Body 请求参数

```json
{
  "password": "string",
  "userid": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» password|body|string| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": "string",
  "msg": "string",
  "data": "string"
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
|» code|string|true|none||none|
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 登录

POST /api/user/login

body中加session，把时间切掉

> Body 请求参数

```json
{
  "account": "202203150203",
  "password": "1234567"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» account|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "admin": 1,
    "userid": 1
  },
  "msg": "ok"
}
```

```json
{
  "code": 404,
  "msg": "用户不存在",
  "data": null
}
```

```json
{
  "code": 405,
  "msg": "密码错误",
  "data": null
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
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» admin|integer|true|none||none|
|»» adminID|integer|true|none||none|
|»» userid|integer|true|none||none|

## POST 插入用户测试数据

POST /api/user/createuser

> Body 请求参数

```json
{
  "account": "202203150202",
  "name": "陈王彬",
  "password": "zjut123456",
  "age": 2022,
  "profession": "计算机",
  "admin": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» account|body|string| 是 |none|
|» name|body|string| 是 |none|
|» password|body|string| 是 |none|
|» age|body|integer| 是 |none|
|» profession|body|string| 是 |none|
|» admin|body|integer| 是 |判断是不是辅导员|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "OK"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 话题广场/言论库

## POST 发表言论

POST /api/square/pulish

> Body 请求参数

```json
{
  "word": "hello,world"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» word|body|string| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "OK",
  "data": null
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
|» msg|string|true|none||none|
|» data|object|true|none||none|

## GET 删除言论

GET /api/square/delete

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|squareId|query|string| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "OK"
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
|» msg|string|true|none||none|
|» data|object|true|none||none|

## GET 查看言论

GET /api/square/iquire

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page_num|query|integer| 否 |none|
|page_size|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "page": 1,
    "pagesize": 10,
    "square": [
      {
        "SquareId": 2,
        "UserId": 1,
        "Word": "hello",
        "Time": "2023-02-20T21:06:09.315+08:00"
      }
    ],
    "total": 1
  },
  "msg": "ok"
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
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» page|integer|true|none||none|
|»» pagesize|integer|true|none||none|
|»» total|integer|true|none||none|
|»» square|[object]|true|none||none|
|»»» Word|string|true|none||none|
|»»» SquareId|integer|false|none||none|
|»»» UserId|integer|false|none||none|
|»»» Time|string|false|none||none|
|» msg|string|true|none||none|

# 学生操作

## GET 获取往年所有分数

GET /api/student/get/grade

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|year|query|integer| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "Score": 27
  },
  "msg": "ok"
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
|» data|object|true|none||none|
|»» Score|integer|true|none||none|
|» msg|string|true|none||none|

## GET 获得具体的分数

GET /api/student/get/detailscore

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|year|query|integer| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "detaillist": [
      {
        "Description": "哈哈哈",
        "Class": 1,
        "Score": 10.1,
        "Module": 1
      },
      {
        "Description": "哈哈哈",
        "Class": 2,
        "Score": 24,
        "Module": 2
      }
    ]
  },
  "msg": "ok"
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
|» msg|string|true|none||none|
|» data|object|true|none||none|

## POST 提交综测申请

POST /api/student/add/application

module的说明：1：美术，2：智育，3：创新，4：劳动，5：道德，6：体育。

> Body 请求参数

```json
{
  "age": 2022,
  "module": 1,
  "description": "哈哈哈",
  "score": 10.1,
  "class": 1,
  "admin": 2,
  "userid": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» age|body|integer| 是 |none|
|» module|body|integer| 是 |none|
|» description|body|string| 是 |none|
|» score|body|number| 是 |none|
|» class|body|integer| 是 |none|
|» admin|body|integer| 是 |none|
|» userid|body|integer| 是 |none|

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

## GET 查询申报表

GET /api/student/inquire

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|page_size|query|integer| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "application": [
      {
        "ID": 1,
        "Age": 2022,
        "Module": 1,
        "Class": 1,
        "CreateTime": "2023-02-20T21:10:28.325+08:00",
        "Userid": 1,
        "Score": 10.1,
        "Description": "lalala",
        "Sta": 1,
        "AdminId": 1,
        "Reason": "ok"
      },
      {
        "ID": 2,
        "Age": 2022,
        "Module": 2,
        "Class": 2,
        "CreateTime": "2023-02-20T21:13:16.221+08:00",
        "Userid": 1,
        "Score": 8.2,
        "Description": "lalala",
        "Sta": 5,
        "AdminId": 1,
        "Reason": "ok"
      },
      {
        "ID": 3,
        "Age": 2022,
        "Module": 3,
        "Class": 1,
        "CreateTime": "2023-02-22T21:26:19.924+08:00",
        "Userid": 1,
        "Score": 10.2,
        "Description": "hahaha",
        "Sta": 5,
        "AdminId": 1,
        "Reason": "ok"
      }
    ],
    "page": 1,
    "pagesize": 10,
    "total": 3
  },
  "msg": "ok"
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
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» application|[object]|true|none||none|
|»»» ID|integer|true|none||none|
|»»» Age|integer|true|none||none|
|»»» Module|integer|true|none||none|
|»»» Class|integer|true|none||none|
|»»» CreateTime|string|true|none||none|
|»»» Userid|integer|true|none||none|
|»»» Score|number|true|none||none|
|»»» Description|string|true|none||none|
|»»» Sta|integer|true|none||none|
|»»» AdminId|integer|true|none||none|
|»»» Reason|string|true|none||none|
|»» page|integer|true|none||none|
|»» pagesize|integer|true|none||none|
|»» total|integer|true|none||none|

## POST 创建申诉

POST /api/student/creatappeal

> Body 请求参数

```json
{
  "applicationid": 0,
  "reason": "string",
  "userid": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» applicationid|body|integer| 是 |none|
|» reason|body|string| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "申诉发送成功"
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
|» data|null|true|none||none|
|» msg|string|true|none||none|

## GET 获取老师信息

GET /api/admin/inquireAdmin

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "teacherlist": [
      {
        "teacherName": "石团可",
        "teacherId": 47
      },
      {
        "teacherName": "深工变于",
        "teacherId": 4
      },
      {
        "teacherName": "向段斯名最准",
        "teacherId": 70
      },
      {
        "teacherName": "题花造技",
        "teacherId": 25
      }
    ]
  },
  "msg": "ok"
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
|» data|object|true|none||none|
|»» teacherlist|[object]|true|none||none|
|»»» TeacherName|string|true|none||none|
|»»» TeacherId|integer|true|none||none|
|» msg|string|true|none||none|

# 学生操作/建议库

## POST 发表建议

POST /api/student/pulishSuggestion

> Body 请求参数

```json
{
  "suggestion": "string",
  "adminName": "string",
  "anon": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|body|body|object| 否 |none|
|» suggestion|body|string| 是 |none|
|» adminName|body|string| 是 |none|
|» anon|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## GET 删除建议

GET /api/student/deleteSuggestion

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|suggestionId|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## POST 更新建议

POST /api/student/updateSuggestion

> Body 请求参数

```json
{
  "suggestionid": 0,
  "suggestion": "string",
  "adminName": "string",
  "anon": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|body|body|object| 否 |none|
|» suggestionid|body|integer| 是 |none|
|» suggestion|body|string| 是 |none|
|» adminName|body|string| 是 |none|
|» anon|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## GET 查看建议

GET /api/student/iquireSuggestion

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|page_num|query|integer| 否 |none|
|page_size|query|integer| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

# 辅导员操作

## POST 成绩录入与管理

POST /api/admin/putscore

成绩入

> Body 请求参数

```json
{
  "score": 0,
  "module": 0,
  "type": 0,
  "studentname": "string",
  "class": 0,
  "age": 0,
  "userid": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 是 |none|
|body|body|object| 否 |none|
|» score|body|number| 是 |none|
|» module|body|integer| 是 |none|
|» type|body|integer| 是 |none|
|» studentname|body|string| 是 |none|
|» class|body|integer| 是 |none|
|» age|body|integer| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "OK"
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
|» data|null|true|none||none|
|» msg|string|true|none||none|

## GET 获取学生成绩列表

GET /api/admin/instructor/scoresearch

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|query|query|string| 否 |查询参数|
|pagenum|query|number| 否 |当前页码|
|pagesize|query|number| 否 |每条显示条数|

> 返回示例

> 成功

```json
{
  "code": 83,
  "msg": "sed nulla Lorem",
  "data": {
    "totalpage": 34,
    "pagenum": 85,
    "students": [
      {
        "name": "谭艳",
        "score1": 64,
        "score2": 27,
        "score3": 15,
        "score4": 13,
        "score5": 82,
        "score6": 58
      },
      {
        "name": "郝敏",
        "score1": 46,
        "score2": 80,
        "score3": 49,
        "score4": 84,
        "score5": 43,
        "score6": 9
      },
      {
        "name": "熊娟",
        "score1": 78,
        "score2": 21,
        "score3": 23,
        "score4": 68,
        "score5": 38,
        "score6": 94
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
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» total|number|true|none|总记录数|none|
|»» pagenum|number|true|none|当前页码|none|
|»» students|[object]|true|none||none|
|»»» name|string|true|none||none|
|»»» studentid|number|true|none||none|
|»»» score1|number|true|none||none|
|»»» score2|number|true|none||none|
|»»» score3|number|true|none||none|
|»»» score4|number|true|none||none|
|»»» score5|number|true|none||none|
|»»» score6|number|true|none||none|

## GET 待处理的申请

GET /api/admin/getunprocessedapplication

缺少total

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|pagenum|query|integer| 否 |none|
|page_size|query|integer| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "application": [
      {
        "ID": 4,
        "Age": 2022,
        "Module": 3,
        "Class": 1,
        "CreateTime": "2023-02-09T21:27:47.522+08:00",
        "Userid": 2,
        "Score": 5.1,
        "Description": "一眼33",
        "Status": 5,
        "AdminId": 1,
        "Reason": ""
      },
      {
        "ID": 5,
        "Age": 2022,
        "Module": 3,
        "Class": 1,
        "CreateTime": "2023-02-09T21:27:54.826+08:00",
        "Userid": 2,
        "Score": 5.1,
        "Description": "一眼2333",
        "Status": 5,
        "AdminId": 1,
        "Reason": ""
      },
      {
        "ID": 6,
        "Age": 2022,
        "Module": 5,
        "Class": 1,
        "CreateTime": "2023-02-10T19:11:23.068+08:00",
        "Userid": 2,
        "Score": 7,
        "Description": "我的天啊",
        "Status": 5,
        "AdminId": 1,
        "Reason": ""
      },
      {
        "ID": 7,
        "Age": 2022,
        "Module": 4,
        "Class": 1,
        "CreateTime": "2023-02-10T19:11:37.001+08:00",
        "Userid": 2,
        "Score": 6,
        "Description": "我好爱你",
        "Status": 5,
        "AdminId": 1,
        "Reason": ""
      },
      {
        "ID": 8,
        "Age": 2022,
        "Module": 3,
        "Class": 1,
        "CreateTime": "2023-02-10T19:11:49.991+08:00",
        "Userid": 2,
        "Score": 8.99,
        "Description": "你算什么",
        "Status": 5,
        "AdminId": 1,
        "Reason": ""
      }
    ]
  },
  "msg": "ok"
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
|» data|object|true|none||none|
|»» application|[object]|true|none||none|
|»»» ID|integer|true|none||none|
|»»» Age|integer|true|none||none|
|»»» Module|integer|true|none||none|
|»»» Class|integer|true|none||none|
|»»» CreateTime|string|true|none||none|
|»»» Userid|integer|true|none||none|
|»»» Score|number|true|none||none|
|»»» Description|string|true|none||none|
|»»» Status|integer|true|none||none|
|»»» AdminId|integer|true|none||none|
|»»» Reason|string|true|none||none|
|» msg|string|true|none||none|

## GET 重新审批的申请

GET /api/admin/getreapprovedapplication

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|page_size|query|integer| 否 |none|
|userid|query|integer| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": {
    "application": [
      {
        "ID": 2,
        "Age": 2022,
        "Module": 2,
        "Class": 2,
        "CreateTime": "2023-02-20T21:13:16.221+08:00",
        "Userid": 1,
        "Score": 8.2,
        "Description": "lalala",
        "Sta": 3,
        "AdminId": 1,
        "Reason": "ok"
      }
    ],
    "page": 1,
    "pagesize": 10,
    "total": 2
  },
  "msg": "ok"
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
|» data|object|true|none||none|
|»» application|[object]|true|none||none|
|»»» ID|integer|false|none||none|
|»»» Age|integer|false|none||none|
|»»» Module|integer|false|none||none|
|»»» Class|integer|false|none||none|
|»»» CreateTime|string|false|none||none|
|»»» Userid|integer|false|none||none|
|»»» Score|number|false|none||none|
|»»» Description|string|false|none||none|
|»»» Sta|integer|false|none||none|
|»»» AdminId|integer|false|none||none|
|»»» Reason|string|false|none||none|
|»» page|integer|true|none||none|
|»» pagesize|integer|true|none||none|
|»» total|integer|true|none||none|
|» msg|string|true|none||none|

## POST 审批申请

POST /api/admin/judge

辅导员审批请求,status的数值代表了审批结果，1为通过，2为不通过，3为重新审批，4为最终不通过，5为未审批

> Body 请求参数

```json
{
  "status": 1,
  "id": 3,
  "reason": "ok",
  "userid": 1
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» status|body|integer| 是 |none|
|» id|body|integer| 是 |application|
|» reason|body|string| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "OK"
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
|» data|null|true|none||none|
|» msg|string|true|none||none|

## POST 删除学生成绩

POST /api/admin/deletestudent

> Body 请求参数

```json
{
  "studentid": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» studentid|body|number| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "msg": "OK",
  "data": null
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
|» msg|string|true|none||none|
|» data|string|true|none||none|

## POST 辅导员撤销审批

POST /api/admin/revokeapplication

> Body 请求参数

```json
{
  "applicationid": 0,
  "userid": 0
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 是 |none|
|body|body|object| 否 |none|
|» applicationid|body|integer| 是 |none|
|» userid|body|integer| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 200,
  "data": null,
  "msg": "撤销成功"
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
|» data|null|true|none||none|
|» msg|string|true|none||none|

# 辅导员操作/理由库

## GET 查询理由库

GET /api/admin/inquireReason

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## POST 添加理由库

POST /api/admin/pulishReason

> Body 请求参数

```json
{
  "reason": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|body|body|object| 否 |none|
|» reason|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## GET 删除理由库

GET /api/admin/deleteReason

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|reasonId|query|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

## POST 更新理由库

POST /api/admin/updateReason

> Body 请求参数

```json
{
  "ID": 0,
  "reason": "string"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|score-system-session|cookie|string| 否 |none|
|body|body|object| 否 |none|
|» ID|body|integer| 是 |none|
|» reason|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
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
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|

# 数据模型

<h2 id="tocS_Tag">Tag</h2>

<a id="schematag"></a>
<a id="schema_Tag"></a>
<a id="tocStag"></a>
<a id="tocstag"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||标签ID编号|
|name|string|false|none||标签名称|

<h2 id="tocS_Category">Category</h2>

<a id="schemacategory"></a>
<a id="schema_Category"></a>
<a id="tocScategory"></a>
<a id="tocscategory"></a>

```json
{
  "id": 1,
  "name": "string"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|false|none||分组ID编号|
|name|string|false|none||分组名称|

<h2 id="tocS_Pet">Pet</h2>

<a id="schemapet"></a>
<a id="schema_Pet"></a>
<a id="tocSpet"></a>
<a id="tocspet"></a>

```json
{
  "id": 1,
  "category": {
    "id": 1,
    "name": "string"
  },
  "name": "doggie",
  "photoUrls": [
    "string"
  ],
  "tags": [
    {
      "id": 1,
      "name": "string"
    }
  ],
  "status": "available"
}

```

### 属性

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|id|integer(int64)|true|none||宠物ID编号|
|category|[Category](#schemacategory)|true|none||分组|
|name|string|true|none||名称|
|photoUrls|[string]|true|none||照片URL|
|tags|[[Tag](#schematag)]|true|none||标签|
|status|string|true|none||宠物销售状态|

#### 枚举值

|属性|值|
|---|---|
|status|available|
|status|pending|
|status|sold|

