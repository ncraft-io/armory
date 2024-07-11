/// 统一表单服务
interface Authing {
    /// 创建表单
    @entity("Account")
    @http.post("armory/auth/v1/accounts")
    create_account(database: String @1, table: String @2)
}