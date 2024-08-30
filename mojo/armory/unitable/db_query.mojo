
type DbQuery {
    id: String @1
    name: String @2

    sql: String @3

    parameters: [String] @4
    table: String @5

    columns: [Column] @15 @db.json //< Query 包含的列的元信息
}
