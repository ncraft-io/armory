
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


## 查询服务范围内的所有数据库及表单

### 请求路径
```http
GET /armory/unitable/v1/databases
```


### 请求参数

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
| `Array<armory.unitable.Database>` |  |


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `armory.unitable.Database`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `name` | `string` |  | N |  | 数据的名称 |
| `displayName` | `string` |  | N |  | 数据库的可显示的名称 |
| `description` | `string` |  | N |  | 数据库的描述 |
| `tables` | `Array<armory.unitable.Table>` |  | N |  | 数据库包含的表单信息 |


#### `armory.unitable.Table`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `columns` | `Array<armory.unitable.Column>` |  | N |  | 表单包含的列的元信息 |
| `createTime` | `string` | `Timestamp` | N |  | 表单创建时间 |
| `database` | `string` |  | N |  | 表单所在的数据库名 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时，作为sheet名称 |
| `id` | `string` |  | N |  | 表单ID |
| `jsonStyle` | `string` |  | N |  | 表格字段导出json的风格，默认与数据库一致为：snake，lower_camel |
| `name` | `string` |  | N |  | 表单名 |
| `tenant` | `string` |  | N |  | 租户名 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `armory.unitable.Table`
| field | type | format | required | default | description |
|---|---|---|---|---|---|


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `force` | `boolean` |  | 否 |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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


#### `armory.unitable.Column`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `createTime` | `string` | `Timestamp` | N |  | 表单列创建时间 |
| `database` | `string` |  | N |  | 所属的表单所在的数据库名称 |
| `dimensional` | `boolean` |  | N |  | 是否是维度相关的，即可枚举的值 |
| `displayName` | `string` |  | N |  | 可以是显示中文的名称 |
| `editable` | `boolean` |  | N |  | 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限 |
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `example` | `mojo.core.Value` |  | N |  | 示例的值 |
| `exportName` | `string` |  | N |  | 导出时使用的名称，比如Excel导出时 |
| `filterable` | `boolean` |  | N |  | 字段能够进行过滤操作 |
| `format` | `string` |  | N |  | 当列为String时，指定更详细的类型，比如时间、几何等 "time", "geometry" |
| `groupDisplayName` | `string` |  | N |  | 所属的列的组合名称 |
| `id` | `string` |  | N |  | 列的ID |
| `indexed` | `boolean` |  | N |  | 是否需要被索引 |
| `name` | `string` |  | N |  | 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格 |
| `originalName` | `string` |  | N |  | 如果该列名为使用函数后的组合名称时，其为原始字段的名称 |
| `referenced` | `string` |  | N |  | 是否为引用字段，可以设置是否自动join |
| `repeated` | `boolean` |  | N |  | is Array type |
| `show` | `boolean` |  | N |  | 是否需要显示 |
| `statistical` | `boolean` |  | N |  | 是否可以被统计 |
| `tableId` | `string` |  | N |  | 所属的表单ID |
| `temporal` | `boolean` |  | N |  | 是否是临时的 |
| `type` | `string` |  | N |  | 列的数据库类型 "bool", "integer", "float", "string" |
| `unique` | `boolean` |  | N |  | 是否需要在表内是唯一的 |
| `updateTime` | `string` | `Timestamp` | N |  | 表单列更新时间 |


#### `mojo.core.Value`
| type | format | description |
|---|---|---|
| `null` |  |  |
| `boolean` |  |  |
| `string` |  |  |
| `string` | `Bytes` | the format is: `b64.{base64 encoded bytes}` |
| `integer` | `Int64` |  |
| `number` | `Float64` |  |
| `mojo.core.Object` |  |  |
| `array` |  |  |


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
| `database` | `string` |  | specify the database name |
| `table` | `string` |  | specify the table name |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `query` | `string` |  | 否 |  | specify the query expression loads form config file |
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
| `filename` | `string` |  |  |


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


## 导出行(实现不同的权限控制)

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/rows:export/{filename}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  |  |
| `table` | `string` |  |  |
| `filename` | `string` |  |  |


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


## 查询行的相关字段的统计值如果未设置stats字段，则检查columns中是否配置了statistical，如果有可以自动进行统计对于文本类型，只统计 group 的 count对于数字类型，则统计 count,sum,avg,max,min对于时间类型，则统计时间范围，并可以按年、按月、按天、按小时进行count统计如果设置了stats字段，则只按照stats字段的表达式进行统计对于文本类型，支持 group text_field   ==>  count text_field group text_field对于数字类型，支持 count number_field, sum number_field对于时间类型，支持 range time_field, years time_field, months, days, hours输出基本输出{ "field_name": {"function_name": "value" }}group函数{ "field_name": {"group_by": {"field_value1": {"count": value}, "field_value2": {"count": value}}

### 请求路径
```http
GET /armory/unitable/v1/databases/{database}/tables/{table}/rows:stat
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `database` | `string` |  | specify the database name |
| `table` | `string` |  | specify the table name |


#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `filter` | `string` |  | 否 |  | additional filter expression |
| `stats` | `Array<string>` |  | 否 |  | specify the statistics field expression, like `sum field` |


### 返回值

#### 返回对象