| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `active` | `boolean` |  | N |  | 是否激活 |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  | 是否删除 |
| `description` | `string` |  | N |  |
| `domain` | `string` |  | N |  |
| `emailAddress` | `string` |  | N |  |
| `id` | `string` |  | N |  |
| `loginTime` | `string` | `Timestamp` | N |  | 用户最近的登录时间 |
| `name` | `string` |  | N |  |
| `nickName` | `string` |  | N |  |
| `otpSecret` | `string` |  | N |  |
| `password` | `string` |  | N |  | 存储密码 sh256(md5(raw_password) + salt)，客户端传输的密码必须采用md5加密 |
| `passwordAlgorithm` | `string` |  | N |  | default is md5 |
| `phoneNumber` | `string` |  | N |  |
| `salt` | `string` |  | N |  | 密码盐 |
| `updateTime` | `string` | `Timestamp` | N |  |  |
