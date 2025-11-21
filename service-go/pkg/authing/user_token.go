package authing

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ncraft-io/armory/go/pkg/armory/auth"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/auth/jwt"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/redis"

	_ "github.com/ncraft-io/ncraft/go/pkg/ncraft/redis/goredis"
)

type UserToken struct {
	JWT   *jwt.JWT
	Redis redis.Redis
}

func NewUserToken() *UserToken {
	token := &UserToken{
		JWT:   jwt.NewJWT(),
		Redis: redis.New(nil),
	}
	if token.JWT == nil {
		panic("should set jwt config and enable to true")
	}
	if token.Redis == nil {
		panic("should set redis config")
	}
	return token
}

func (t *UserToken) CreateToken(user *auth.User) string {
	if t != nil && user != nil {
		token := make(jwt.Token)

		token.SetIssuer("armory").SetSubject(user.Id).SetIssuedAt()

		if !user.Active {
			token.Set("active", false)
		}

		// Sign a JWT!
		signed, err := t.JWT.Generate(token)
		if err != nil {
			fmt.Printf("failed to sign token: %s\n", err)
		}

		return signed
	}

	return ""
}

func (t *UserToken) RefreshToken(token string) {
}

func NewSessionKey(subject string) string {
	return fmt.Sprintf("authing/subjects/%s", subject)
}

func (t *UserToken) Get(subject string) *auth.User {
	user, err := redis.Get(t.Redis, NewSessionKey(subject))
	if err != nil {
		return nil
	}

	usr := &auth.User{}
	err = jsoniter.UnmarshalFromString(user, usr)
	if err != nil {
		return nil
	}
	return usr
}

func (t *UserToken) ParseToken(tokenStr string) (jwt.Token, error) {
	return t.JWT.Parse(tokenStr)
}

func (t *UserToken) SetSession(user *auth.User) {
	if user != nil {
		us, _ := jsoniter.MarshalToString(&auth.User{
			Id:           user.Id,
			Name:         user.Name,
			Description:  user.Description,
			Domain:       user.Domain,
			PhoneNumber:  user.PhoneNumber,
			EmailAddress: user.EmailAddress,
			NickName:     user.NickName,
			Active:       user.Active,
			LoginTime:    user.LoginTime,
		})
		_, err := redis.Set(t.Redis, NewSessionKey(user.Id), us)
		if err != nil {
			logs.ErrLogw("failed to update the session to redis", "error", err)
		}
	}
}

func (t *UserToken) GetSession(ctx context.Context) (string, *auth.User) {
	token := jwt.GetContextToken(ctx)
	subject := jwt.GetContextSubject(ctx)
	return token, t.Get(subject)
}

func (t *UserToken) DeleteSession(user *auth.User) {
	if user != nil {
		key := NewSessionKey(user.Id)
		_, _ = redis.Del(t.Redis, key)
	}
}
