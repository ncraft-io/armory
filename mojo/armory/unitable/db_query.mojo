
/// Database Query Entity
type DbQuery {
    type Parameter {
        name: String @1
        type: String @2
        is_array: Bool @3
        pg_array: Bool @4
    }

    id: String @1
    name: String @2

    sql: String @3

    parameters: [Parameter] @4 @db.json
    // table: String @5
    database: String @6
    json_style: String @7 //< 表格字段导出json的风格，默认与数据库一致为：snake，lower_camel

    columns: [Column] @15 @db.json //< the meta info of the column in the query

    create_time: Timestamp @100
    update_time: Timestamp @101
}
