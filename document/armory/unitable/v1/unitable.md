
# 统一表单服务
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


## 查询表单

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |


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
| `Array<armory.unitable.Table>` |  |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `armory.unitable.Table`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `columns` | `Array<armory.unitable.Column>` |  | N |  | 表单包含的列的元信息 |
| `createTime` | `string` | `Timestamp` | N |  | 表单创建时间 |
| `database` | `string` |  | N |  | 表单所在的数据库名 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时，作为sheet名称 |
| `id` | `string` |  | N |  | 表单ID |
| `name` | `string` |  | N |  | 表单名 |
| `tenant` | `string` |  | N |  | 租户名 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单更新时间 |


## 创建表单

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 获取表单

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `columns` | `Array<armory.unitable.Column>` |  | N |  | 表单包含的列的元信息 |
| `createTime` | `string` | `Timestamp` | N |  | 表单创建时间 |
| `database` | `string` |  | N |  | 表单所在的数据库名 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时，作为sheet名称 |
| `id` | `string` |  | N |  | 表单ID |
| `name` | `string` |  | N |  | 表单名 |
| `tenant` | `string` |  | N |  | 租户名 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单更新时间 |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 更新表单

### 请求路径
```http
PUT /armory/unitable/v1/databases/{database}/tables/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `id` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
对象为空

## 删除表单

### 请求路径
```http
DELETE /armory/unitable/v1/databases/{database}/tables/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `id` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `force` | `boolean` |  | 否 |  |  |


### 返回值

#### 返回对象
对象为空

## 直接从数据库中同步指定的表

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables/{id}:sync
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `columns` | `Array<armory.unitable.Column>` |  | N |  | 表单包含的列的元信息 |
| `createTime` | `string` | `Timestamp` | N |  | 表单创建时间 |
| `database` | `string` |  | N |  | 表单所在的数据库名 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时，作为sheet名称 |
| `id` | `string` |  | N |  | 表单ID |
| `name` | `string` |  | N |  | 表单名 |
| `tenant` | `string` |  | N |  | 租户名 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单更新时间 |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 在指定的表内查找列

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/columns
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


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
| `Array<armory.unitable.Column>` |  |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 在指定的表内创建列

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables/{table}/columns
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 在指定的表内获取列

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


## 在指定的表内更新列

### 请求路径
```http
PUT /armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
对象为空

## 在指定的表内删除列

### 请求路径
```http
DELETE /armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 在指定的表内批量更新列

### 请求路径
```http
PUT /armory/unitable/v1/databases/{database}/tables/{table}/columns:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象
| type | description |
|---|---|
| `Array<armory.unitable.Column>` |  |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
对象为空

## 在指定的表内批量创建列

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables/{table}/columns:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象
| type | description |
|---|---|
| `Array<armory.unitable.Column>` |  |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `tableId` | `string` |  | N |  | 所属的表单名 |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "integer", "number", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


### 返回值

#### 返回对象
对象为空

## 在指定的表内批量删除列

### 请求路径
```http
DELETE /armory/unitable/v1/databases/{database}/tables/{table}/columns:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `ids` | `Array<string>` |  | 否 |  |  |


### 返回值

#### 返回对象
对象为空

## 查询行

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/rows
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


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
| `Array<mojo.core.Object>` |  |


## 在指定的表内新增数据

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables/{table}/rows
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象

### 返回值

#### 返回对象

## 在指定的表内获取某一行数据

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象

## 在指定的表内更新指定ID的数据

### 请求路径
```http
PUT /armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


#### Body 请求对象

### 返回值

#### 返回对象
对象为空

## 删除指定某一行

### 请求路径
```http
DELETE /armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空

## 批量更新行数据

### 请求路径
```http
PUT /armory/unitable/v1/databases/{database}/tables/{table}/rows:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象
| type | description |
|---|---|
| `Array<mojo.core.Object>` |  |


### 返回值

#### 返回对象
对象为空

## 批量创建行数据

### 请求路径
```http
POST /armory/unitable/v1/databases/{database}/tables/{table}/rows:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Body 请求对象
| type | description |
|---|---|
| `Array<mojo.core.Object>` |  |


### 返回值

#### 返回对象
对象为空

## 批量删除行

### 请求路径
```http
DELETE /armory/unitable/v1/databases/{database}/tables/{table}/rows:batch
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `ids` | `Array<string>` |  | 否 |  |  |


### 返回值

#### 返回对象
对象为空

## 导出行(实现不同的权限控制)

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/rows:export
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |


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
| `Array<mojo.core.Object>` |  |
