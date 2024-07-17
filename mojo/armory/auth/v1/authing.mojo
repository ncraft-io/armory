/// 权限管理服务
interface Authing {
    /// 创建账号
    @entity("Account")
    @http.post("armory/auth/v1/accounts")
    create_account(database: String @1, table: String @2) -> Account
}