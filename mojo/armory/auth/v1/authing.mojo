/// 权限管理服务
interface Authing {
    /// 创建用户
    @entity("User")
    @http.post("/armory/auth/v1/users")
    @http.post("/armory/auth/v1/domains/{domain}/users")
    create_user(domain: String @1
                user: User @2 @http.body) -> User

    @entity("User")
    @http.post("/armory/auth/v1/users:batch")
    @http.post("/armory/auth/v1/domains/{domain}/users:batch")
    batch_create_users(domain: String @1
                       users: [User] @2 @http.body) -> [User]

    /// 更新用户
    @entity("User")
    @http.put("/armory/auth/v1/users/{id}")
    @http.put("/armory/auth/v1/domains/{domain}/users/{id}")
    update_user(domain: String @1
                id: String @2
                user: User @3 @http.body)

    /// 激活用户
    @entity("User")
    @http.post("/armory/auth/v1/users/{id}:active")
    @http.post("/armory/auth/v1/domains/{domain}/users/{id}:active")
    active_user(domain: String @1 //< 用户域
                id: String @2 //< 用户ID
                reset_password: String @3 //< 用户重置的密码
                passcode: String @4) //< 绑定TOTP应用后获取的密码
                -> LogonUser

    /// 获取用户
    @entity("User")
    @http.get("/armory/auth/v1/users/{id}")
    get_user(id: String @1) -> User

    /// 查询用户
    @entity("User")
    @http.get("/armory/auth/v1/users")
    @http.get("/armory/auth/v1/domains/{domain}/users")
    list_user(domain: String @1) -> [User]

    /// 删除用户
    @entity("User")
    @http.delete("/armory/auth/v1/users/{id}")
    delete_user(id: String @1)

    /// 变更用户密码
    @entity("User")
    @http.put("/armory/auth/v1/users/{id}/passwords")
    @http.put("/armory/auth/v1/domains/{domain}/users/{id}/passwords")
    update_password(domain: String @1, id: String @2, old_password: String @3, new_password: String @4)

    //admin_reset_password()
    //signup(user: User @1) -> User
    //signoff()

    /// 用户登录
    @entity("User")
    @http.post("/armory/auth/v1/users:login")
    @http.post("/armory/auth/v1/domains/{domain}/users:login")
    login(user: String @1
          captcha: String @4
          password: String @5
          otp: String @6 //< one time password
          type: String @7 //< account type
          domain: String @8) -> LogonUser

    /// 用户登出
    @entity("User")
    @http.post("/armory/auth/v1/users:logout")
    @http.post("/armory/auth/v1/domains/{domain}/users:logout")
    logout(user: String @1) //< user可以是用户ID，name, mail等唯一性字段
}