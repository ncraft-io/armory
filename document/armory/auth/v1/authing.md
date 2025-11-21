
# 权限管理服务
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## 查询用户

### 请求路径
```http
GET /armory/auth/v1/domains/{domain}/users
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `page_size` | `integer` | `Int32` | 否 |  | the page size for pagination request |
| `page_token` | `string` |  | 否 |  | the page token for pagination request, usually like "1", "2" ... |
| `skip` | `integer` | `Int32` | 否 |  | skip the first items count for the request |
| `filter` | `string` |  | 否 |  | the mojo expression for DB query |
| `order` | `mojo.core.Ordering` |  | 否 |  | setting the order field for result, like "name desc" |
| `field_mask` | `string` | `FieldMask` | 否 |  | control the fields which need to be retrieved |
| `unique` | `boolean` |  | 否 |  | make the fields which returns are unique, equals to "SELECT DISTINCT" in sql |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<armory.auth.User>` |  |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 创建用户

### 请求路径
```http
POST /armory/auth/v1/domains/{domain}/users
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 获取用户

### 请求路径
```http
GET /armory/auth/v1/domains/{domain}/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 更新用户

### 请求路径
```http
PUT /armory/auth/v1/domains/{domain}/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |
| `id` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


### 返回值

#### 返回对象
对象为空

## 删除用户

### 请求路径
```http
DELETE /armory/auth/v1/domains/{domain}/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 变更用户密码

### 请求路径
```http
PUT /armory/auth/v1/domains/{domain}/users/{id}/passwords
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |
| `id` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `old_password` | `string` |  | 否 |  |  |
| `new_password` | `string` |  | 否 |  |  |


### 返回值

#### 返回对象
对象为空

## 激活用户

### 请求路径
```http
POST /armory/auth/v1/domains/{domain}/users/{id}:active
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  | 用户域 |
| `id` | `string` |  | 用户ID |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `reset_password` | `string` |  | 否 |  | 用户重置的密码 |
| `passcode` | `string` |  | 否 |  | 绑定TOTP应用后获取的密码 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `user` | `armory.auth.User` |  | N |  |  |
| `token` | `string` |  | N |  |
| `totp` | `armory.auth.TOTP` |  | N |  | need user to set the totp app when the user login first time. |
| `loginTime` | `string` | `Timestamp` | N |  |  |


#### `armory.auth.TOTP`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `secret` | `string` |  | N |  |
| `qrCode` | `string` |  | N |  | base64 encoded png image file |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 用户登录

### 请求路径
```http
POST /armory/auth/v1/domains/{domain}/users:login
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `user` | `string` |  | 否 |  |  |
| `captcha` | `string` |  | 否 |  |  |
| `password` | `string` |  | 否 |  |  |
| `otp` | `string` |  | 否 |  | one time password |
| `type` | `string` |  | 否 |  | account type |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `user` | `armory.auth.User` |  | N |  |  |
| `token` | `string` |  | N |  |
| `totp` | `armory.auth.TOTP` |  | N |  | need user to set the totp app when the user login first time. |
| `loginTime` | `string` | `Timestamp` | N |  |  |


#### `armory.auth.TOTP`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `secret` | `string` |  | N |  |
| `qrCode` | `string` |  | N |  | base64 encoded png image file |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 用户登出

### 请求路径
```http
POST /armory/auth/v1/domains/{domain}/users:logout
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `user` | `string` |  | 否 |  | user可以是用户ID，name, mail等唯一性字段 |


### 返回值

#### 返回对象
对象为空

## 查询用户

### 请求路径
```http
GET /armory/auth/v1/users
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `page_size` | `integer` | `Int32` | 否 |  | the page size for pagination request |
| `page_token` | `string` |  | 否 |  | the page token for pagination request, usually like "1", "2" ... |
| `skip` | `integer` | `Int32` | 否 |  | skip the first items count for the request |
| `filter` | `string` |  | 否 |  | the mojo expression for DB query |
| `order` | `mojo.core.Ordering` |  | 否 |  | setting the order field for result, like "name desc" |
| `field_mask` | `string` | `FieldMask` | 否 |  | control the fields which need to be retrieved |
| `unique` | `boolean` |  | 否 |  | make the fields which returns are unique, equals to "SELECT DISTINCT" in sql |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<armory.auth.User>` |  |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 创建用户

### 请求路径
```http
POST /armory/auth/v1/users
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 获取用户

### 请求路径
```http
GET /armory/auth/v1/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 更新用户

### 请求路径
```http
PUT /armory/auth/v1/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |
| `id` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


### 返回值

#### 返回对象
对象为空

## 删除用户

### 请求路径
```http
DELETE /armory/auth/v1/users/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 变更用户密码

### 请求路径
```http
PUT /armory/auth/v1/users/{id}/passwords
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |
| `id` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `old_password` | `string` |  | 否 |  |  |
| `new_password` | `string` |  | 否 |  |  |


### 返回值

#### 返回对象
对象为空

## 激活用户

### 请求路径
```http
POST /armory/auth/v1/users/{id}:active
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  | 用户域 |
| `id` | `string` |  | 用户ID |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `reset_password` | `string` |  | 否 |  | 用户重置的密码 |
| `passcode` | `string` |  | 否 |  | 绑定TOTP应用后获取的密码 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `user` | `armory.auth.User` |  | N |  |  |
| `token` | `string` |  | N |  |
| `totp` | `armory.auth.TOTP` |  | N |  | need user to set the totp app when the user login first time. |
| `loginTime` | `string` | `Timestamp` | N |  |  |


#### `armory.auth.TOTP`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `secret` | `string` |  | N |  |
| `qrCode` | `string` |  | N |  | base64 encoded png image file |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 用户登录

### 请求路径
```http
POST /armory/auth/v1/users:login
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `domain` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `user` | `string` |  | 否 |  |  |
| `captcha` | `string` |  | 否 |  |  |
| `password` | `string` |  | 否 |  |  |
| `otp` | `string` |  | 否 |  | one time password |
| `type` | `string` |  | 否 |  | account type |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `user` | `armory.auth.User` |  | N |  |  |
| `token` | `string` |  | N |  |
| `totp` | `armory.auth.TOTP` |  | N |  | need user to set the totp app when the user login first time. |
| `loginTime` | `string` | `Timestamp` | N |  |  |


#### `armory.auth.TOTP`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `secret` | `string` |  | N |  |
| `qrCode` | `string` |  | N |  | base64 encoded png image file |


#### `armory.auth.User`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `name` | `string` |  | N |  |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `salt` | `string` |  | N |  | 密码盐 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `active` | `boolean` |  | N |  | 是否激活 |
| `otpSecret` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |


## 用户登出

### 请求路径
```http
POST /armory/auth/v1/users:logout
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `user` | `string` |  | 否 |  | user可以是用户ID，name, mail等唯一性字段 |


### 返回值

#### 返回对象
对象为空