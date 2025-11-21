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
