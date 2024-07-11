/// unitable的列的元信息
@entity
type Column {
    id: String @1 //< 列的ID
    database: String @2 @db.index //< 所属的表单所在的数据库名称
    table_id: String @3 @db.index //< 所属的表单名

    name: String @4 @db.index //< 表单的列名，符合数据库的列名规格，采用 `[a-z][a-z_0-9]*` 规格
    display_name: String @5 //< 可以是显示中文的名称
    export_name: String @6 //< 导出时使用的名称，比如Excel导出时

    group_display_name: String @7 //< 所属的列的组合名称

    type: String @8 //< 列的数据库类型 "integer", "number", "string"
    format: String @9 //< 当列为String时，指定更详细的类型，比如时间、几何等

    indexed: Bool @10  //< 是否需要被索引
    unique: Bool @11   //< 是否需要在表内是唯一的
    show: Bool @12     //< 是否需要显示
    editable: Bool @13 //< 字段是否可编辑，控制前端显示时，允许用户编辑，实际可否编辑还得检查相应权限
    filterable: Bool @14 //< 字段能够进行过滤操作
    temporal: Bool @15 //< 是否是临时的
    dimensional: Bool @16 //< 是否是维度相关的，即可枚举的值

    referenced: String @20 //< 是否为引用字段，可以设置是否自动join

    create_time: Timestamp @100 //< 表单列创建时间
    update_time: Timestamp @101 //< 表单列更新时间
    delete_time: Timestamp @102 //< 表单列删除时间
}
