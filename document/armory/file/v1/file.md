
# 文件服务
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


## 上传文件，采用http form的形式进行上传，返回文件元信息

### 请求路径
```http
POST /armory/file/v1/files
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 文件名 |
| `mineType` | `string` |  | N |  | 文件的mine-type |
| `url` | `string` | `Url` | N |  | 文件的url |
| `size` | `integer` | `Int64` | N |  | 文件大小，单位：Byte |
| `content` | `string` | `Bytes` | N |  | 文件内容 |
| `createTime` | `string` | `Timestamp` | N |  | 文件创建时间 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 文件名 |
| `mineType` | `string` |  | N |  | 文件的mine-type |
| `url` | `string` | `Url` | N |  | 文件的url |
| `size` | `integer` | `Int64` | N |  | 文件大小，单位：Byte |
| `content` | `string` | `Bytes` | N |  | 文件内容 |
| `createTime` | `string` | `Timestamp` | N |  | 文件创建时间 |


## 查询服务范围内的所有数据库及表单

### 请求路径
```http
GET /armory/file/v1/files/{name}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `name` | `string` |  | 文件名称 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 文件名 |
| `mineType` | `string` |  | N |  | 文件的mine-type |
| `url` | `string` | `Url` | N |  | 文件的url |
| `size` | `integer` | `Int64` | N |  | 文件大小，单位：Byte |
| `content` | `string` | `Bytes` | N |  | 文件内容 |
| `createTime` | `string` | `Timestamp` | N |  | 文件创建时间 |


## 批量上传文件，采用http form的形式进行上传，返回文件元信息

### 请求路径
```http
POST /armory/file/v1/files:batch
```


### 请求参数

#### Body 请求对象
| type | description |
|---|---|
| `Array<armory.file.BinaryFile>` |  |


#### `armory.file.BinaryFile`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 文件名 |
| `mineType` | `string` |  | N |  | 文件的mine-type |
| `url` | `string` | `Url` | N |  | 文件的url |
| `size` | `integer` | `Int64` | N |  | 文件大小，单位：Byte |
| `content` | `string` | `Bytes` | N |  | 文件内容 |
| `createTime` | `string` | `Timestamp` | N |  | 文件创建时间 |


### 返回值

#### 返回对象
| type | description |
|---|---|
| `Array<armory.file.BinaryFile>` |  |


#### `armory.file.BinaryFile`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 文件名 |
| `mineType` | `string` |  | N |  | 文件的mine-type |
| `url` | `string` | `Url` | N |  | 文件的url |
| `size` | `integer` | `Int64` | N |  | 文件大小，单位：Byte |
| `content` | `string` | `Bytes` | N |  | 文件内容 |
| `createTime` | `string` | `Timestamp` | N |  | 文件创建时间 |
