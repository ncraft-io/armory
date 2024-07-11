
/// 统一表单服务
interface Unitable {
    /// 创建表单
    @entity("Table")
    @http.post("/armory/unitable/v1/databases/{database}/tables")
    create_table(database: String @1, table: Table @2 @http.body) -> Table

    /// 更新表单
    @entity("Table")
    @http.put("/armory/unitable/v1/databases/{database}/tables/{id}")
    update_table(database: String @1, table: Table @2 @http.body, id: String @3)

    /// 获取表单
    @entity("Table")
    @http.get("/armory/unitable/v1/databases/{database}/tables/{id}")
    get_table(database: String @1, id: String @2) -> Table

    /// 查询表单
    @entity("Table")
    @http.get("/armory/unitable/v1/databases/{database}/tables")
    list_tables(database: String @1) -> [Table]

    /// 删除表单
    @entity("Table")
    @http.delete("/armory/unitable/v1/databases/{database}/tables/{id}")
    delete_table(database: String @1, id: String @2, force: Bool @3)

    /// 直接从数据库中同步指定的表
    @entity("Table")
    @http.post("/armory/unitable/v1/databases/{database}/tables/{id}:sync")
    sync_table(database: String @1, id: String @2) -> Table

    /// 在指定的表内创建列
    @entity("Column")
    @http.post("/armory/unitable/v1/databases/{database}/tables/{table}/columns")
    create_column(database: String @1, table: String @2, column: Column @3 @http.body) -> Column

    /// 在指定的表内更新列
    @entity("Column")
    @http.put("/armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}")
    update_column(database: String @1, table: String @2, column: Column @3 @http.body, id: String @4)

    /// 在指定的表内获取列
    @entity("Column")
    @http.get("/armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}")
    get_column(database: String @1, table: String @2, id: String @3) -> Column

    /// 在指定的表内删除列
    @entity("Column")
    @http.delete("/armory/unitable/v1/databases/{database}/tables/{table}/columns/{id}")
    delete_column(database: String @1, table: String @2, id: String @3)

    /// 在指定的表内查找列
    @entity("Column")
    @http.get("/armory/unitable/v1/databases/{database}/tables/{table}/columns")
    list_columns(database: String @1, table: String @2) -> [Column]

    /// 在指定的表内批量创建列
    @entity("Column")
    @http.post("/armory/unitable/v1/databases/{database}/tables/{table}/columns:batch")
    batch_create_columns(database: String @1, table: String @2, columns: [Column] @3 @http.body)

    /// 在指定的表内批量更新列
    @entity("Column")
    @http.put("/armory/unitable/v1/databases/{database}/tables/{table}/columns:batch")
    batch_update_column(database: String @1, table: String @2, columns: [Column] @3 @http.body)

    /// 在指定的表内批量删除列
    @entity("Column")
    @http.delete("/armory/unitable/v1/databases/{database}/tables/{table}/columns:batch")
    batch_delete_column(database: String @1, table: String @2, ids: [String] @3)

    /// 在指定的表内新增数据
    @http.post("/armory/unitable/v1/databases/{database}/tables/{table}/rows")
    create_row(database: String @1, table: String @2, row: Object @3 @http.body) -> Object

    /// 在指定的表内更新指定ID的数据
    @http.put("/armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}")
    update_row(database: String @1, table: String @2, id: String @3, row: Object @4 @http.body)

    /// 在指定的表内获取某一行数据
    @http.get("/armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}")
    get_row(database: String @1, table: String @2, id: String @3) -> Object

    /// 删除指定某一行
    @http.delete("/armory/unitable/v1/databases/{database}/tables/{table}/rows/{id}")
    delete_row(database: String @1, table: String @2, id: String @3)

    /// 查询行
    @http.get("/armory/unitable/v1/databases/{database}/tables/{table}/rows")
    list_row(database: String @1, table: String @2) -> [Object]

    /// 导出行(实现不同的权限控制)
    @http.get("/armory/unitable/v1/databases/{database}/tables/{table}/rows:export")
    export_row(database: String @1, table: String @2) -> [Object]

    /// 批量创建行数据
    @http.post("/armory/unitable/v1/databases/{database}/tables/{table}/rows:batch")
    batch_create_rows(database: String @1,table: String @2, rows: [Object] @3 @http.body)

    /// 批量更新行数据
    @http.put("/armory/unitable/v1/databases/{database}/tables/{table}/rows:batch")
    batch_update_rows(database: String @1,table: String @2, rows: [Object] @4 @http.body)

    /// 批量删除行
    @http.delete("/armory/unitable/v1/databases/{database}/tables/{table}/rows:batch")
    batch_delete_rows(database: String @1, table: String @2, ids: [String] @3)
}
