
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

    parameters: [Parameter] @4
    table: String @5

    columns: [Column] @15 @db.json //< Query 包含的列的元信息
}
