/// 数据库元信息
type Database {
    name: String @1 //< 数据的名称
    display_name: String @2 //< 数据库的可显示的名称
    description: String @3 //< 数据库的描述

    tables: [Table] @5 //< 数据库包含的表单信息 
}
