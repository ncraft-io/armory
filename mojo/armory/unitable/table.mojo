
/// 动态table的元信息
@entity
type Table {
    id: String @1 //< 表单ID
    name: String @2 @db.index //< 表单名
    display_name: String @3 //< 可以是显示中文的名称
    export_name: String @4 //< 导出时使用的名称，比如Excel导出时，作为sheet名称
    
    tenant: String @5 @db.index //< 租户名

    database: String @10 @db.index //< 表单所在的数据库名

    columns: [Column] @15 //< 表单包含的列的元信息

    create_time: Timestamp @100 //< 表单创建时间
    update_time: Timestamp @101 //< 表单更新时间
    delete_time: db.DeleteTime @102 //< 表单删除时间
}
