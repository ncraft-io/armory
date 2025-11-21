| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `loginTime` | `string` | `Timestamp` | N |  |  |
| `token` | `string` |  | N |  |
| `totp` | `armory.auth.TOTP` |  | N |  | need user to set the totp app when the user login first time. |
| `user` | `armory.auth.User` |  | N |  |  |
