

type LogonUser {
    user: User @1

    token: String @2

    totp: TOTP @5 //< need user to set the totp app when the user login first time.

    login_time: Timestamp @1000
}

type TOTP {
    secret: String @5
    qr_code: String @6 //< base64 encoded png image file
}