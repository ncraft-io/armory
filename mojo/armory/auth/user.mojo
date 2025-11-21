
type User {
    id: String @1
    name: String @2
    description: String @3
    domain: String @4

    password: String @5 //< 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密
    salt: String @6 //< 密码盐
    password_algorithm: String @7 //< default is md5

    phone_number: String @10
    email_address: String @11

    nick_name: String @13

    active: Bool @20 //< 是否激活
    otp_secret: String @21

    is_admin: Bool @25 //< 是否是管理员

    login_time: Timestamp @30 //< 用户最近的登录时间

    create_time: Timestamp @1000
    update_time: Timestamp @1001
    delete_time: db.DeleteTime @1002 //< 是否删除
}
