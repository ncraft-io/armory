package handlers

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"github.com/mojo-lang/mojo/go/pkg/mojo/core"
	"github.com/ncraft-io/armory/service-go/pkg/authing"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/auth/jwt"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
	"github.com/pquerna/otp/totp"
	"github.com/segmentio/ksuid"
	"image/png"

	"github.com/ncraft-io/armory/go/pkg/armory/auth"
	pb "github.com/ncraft-io/armory/go/pkg/armory/auth/v1"
	"github.com/ncraft-io/armory/service-go/pkg/model"
)

type authingServer struct {
	pb.UnimplementedAuthingServer
}

// NewService returns a naive, stateless implementation of Interface.
func NewService() pb.AuthingServer {
	return authingServer{}
}

// CreateUser implements Interface.
func (s authingServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*auth.User, error) {
	if in.User == nil {
		return nil, core.NewInvalidArgumentError("not set the user")
	}
	if len(in.User.Name) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user name")
	}
	if len(in.User.Password) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user password")
	}

	if len(in.User.Id) == 0 {
		in.User.Id = ksuid.New().String()
	}

	if len(in.User.Domain) == 0 {
		in.User.Domain = in.Domain
	}

	in.User.Active = false
	in.User.Salt = rand.Text()

	sum := sha256.New().Sum([]byte(in.User.Password + in.User.Salt))
	//sum := sha256.Sum256([]byte(in.User.Password + in.User.Salt))
	in.User.Password = base64.StdEncoding.EncodeToString(sum[:])

	if in.User.CreateTime == nil {
		in.User.CreateTime = core.Now()
	}
	in.User.UpdateTime = core.Now()

	if _, err := model.GetUserModel().Create(ctx, in.User); err != nil {
	}

	resp := &auth.User{
		Id:   in.User.Id,
		Name: in.User.Name,
	}

	return resp, nil
}

// BatchCreateUsers implements Interface.
func (s authingServer) BatchCreateUsers(ctx context.Context, in *pb.BatchCreateUsersRequest) (*pb.BatchCreateUsersResponse, error) {
	if len(in.Users) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user")
	}

	resp := &pb.BatchCreateUsersResponse{}
	for _, user := range in.Users {
		request := &pb.CreateUserRequest{
			Domain: in.Domain,
			User:   user,
		}
		if response, err := s.CreateUser(ctx, request); err != nil {
			logs.ErrLogw("failed to create the user", "user", user.Name)
		} else {
			resp.Users = append(resp.Users, response)
		}
	}

	return resp, nil
}

// UpdateUser implements Interface.
func (s authingServer) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*core.Null, error) {
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user id")
	}
	if in.User == nil {
		return nil, core.NewInvalidArgumentError("not set the user")
	}

	if user, err := model.GetUserModel().Get(ctx, in.Id); err != nil {
		_ = user
	}

	in.User.Password = ""
	in.User.Salt = ""
	in.User.UpdateTime = core.Now()

	if _, err := model.GetUserModel().Update(ctx, in.User); err != nil {
	}

	resp := &core.Null{}
	return resp, nil
}

// ActiveUser implements Interface.
func (s authingServer) ActiveUser(ctx context.Context, in *pb.ActiveUserRequest) (*auth.LogonUser, error) {
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user id")
	}
	if len(in.ResetPassword) == 0 {
		return nil, core.NewInvalidArgumentError("active user should reset new password")
	}
	disableOTP := authing.GetAuthing().Config.DisableOTP
	if !disableOTP {
		if len(in.Passcode) == 0 {
			return nil, core.NewInvalidArgumentError("active user should set the totp passcode")
		}
	}

	// test the auth token
	// token := jwt.GetContextToken(ctx)

	if user, err := model.GetUserModel().Get(ctx, in.Id); err != nil || user == nil {
		return nil, core.NewUnauthenticatedError("user or password is not valid")
	} else {
		if !disableOTP {
			secret, _ := base32.StdEncoding.DecodeString(user.OtpSecret)
			key, _ := totp.Generate(totp.GenerateOpts{
				Issuer:      "xd",
				AccountName: user.Name,
				Secret:      secret,
			})
			if valid := totp.Validate(in.Passcode, key.Secret()); !valid {
				authing.GetUserToken().DeleteSession(user)
				return nil, core.NewInvalidArgumentError("the passcode is not valid")
			}
		}

		sum := sha256.New().Sum([]byte(in.ResetPassword + user.Salt))
		password := base64.StdEncoding.EncodeToString(sum[:])
		updated := &auth.User{
			Id:         in.Id,
			Password:   password,
			Active:     true,
			UpdateTime: core.Now(),
		}
		if _, err = model.GetUserModel().Update(ctx, updated); err != nil {
			authing.GetUserToken().DeleteSession(user)
			return nil, core.NewInternalError("can't update the user")
		}

		authing.GetUserToken().SetSession(user)
		user.Active = true
		return &auth.LogonUser{
			User: &auth.User{
				Id: user.Id,
			},
			Token:     authing.GetUserToken().CreateToken(user),
			LoginTime: user.LoginTime,
		}, nil
	}
}

// DeleteUser implements Interface.
func (s authingServer) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*core.Null, error) {
	resp := &core.Null{}
	return resp, nil
}

// UpdatePassword implements Interface.
func (s authingServer) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordRequest) (*core.Null, error) {
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user id")
	}
	if len(in.NewPassword) == 0 {
		return nil, core.NewInvalidArgumentError("the new password is empty")
	}

	selfUpdate := false
	if subject, ok := ctx.Value("jwt:subject").(string); ok {
		if subject == in.Id {
			if len(in.OldPassword) == 0 {
				return nil, core.NewInvalidArgumentError("should set the old password")
			}
			selfUpdate = true
		}
	}

	if user, err := model.GetUserModel().Get(ctx, in.Id); err != nil || user == nil {
		return nil, core.NewNotFoundError("user not found")
	} else {
		// check the old password
		if selfUpdate {
			sum := sha256.New().Sum([]byte(in.OldPassword + user.Salt))
			if user.Password != base64.StdEncoding.EncodeToString(sum[:]) {
				return nil, core.NewInvalidArgumentError("the old password is invalid")
			}
		}

		// check the user, should be admin

		sum := sha256.New().Sum([]byte(in.NewPassword + user.Salt))
		password := base64.StdEncoding.EncodeToString(sum[:])
		updated := &auth.User{
			Id:         in.Id,
			Password:   password,
			UpdateTime: core.Now(),
		}
		if _, err = model.GetUserModel().Update(ctx, updated); err != nil {
		}
	}

	return &core.Null{}, nil
}

// Login implements Interface.
func (s authingServer) Login(ctx context.Context, in *pb.LoginRequest) (*auth.LogonUser, error) {
	if len(in.User) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user name or email or phone number")
	}
	if len(in.Password) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user password")
	}

	if user, err := model.GetUserModel().Get(ctx, in.User); err != nil || user == nil {
		return nil, core.NewUnauthenticatedError("user or password is not valid")
	} else {
		if !user.CheckPassword(in.Password) {
			return nil, core.NewUnauthenticatedError("user or password is not valid")
		}

		logon := &auth.LogonUser{
			User: &auth.User{
				Id:          user.Id,
				Name:        user.Name,
				Description: user.Description,
				Domain:      user.Domain,
				NickName:    user.NickName,
				LoginTime:   user.LoginTime,
			},
			Token:     authing.GetUserToken().CreateToken(user),
			Totp:      nil,
			LoginTime: user.LoginTime,
		}

		if authing.GetAuthing().Config.DisableOTP {
			if user.Active {
				authing.GetUserToken().SetSession(user)
			} else {
				logon.Totp = &auth.TOTP{
					// Secret: key.Secret(),
					QrCode: "empty",
				}
				// return nil, core.NewUnauthenticatedError("user is not active")
			}
		} else {
			if user.Active {
				if len(in.Otp) == 0 {
					return nil, core.NewInvalidArgumentError("active user should set the totp passcode")
				}

				secret, _ := base32.StdEncoding.DecodeString(user.OtpSecret)
				key, _ := totp.Generate(totp.GenerateOpts{
					Issuer:      "xd",
					AccountName: user.Name,
					Secret:      secret,
				})
				if valid := totp.Validate(in.Otp, key.Secret()); !valid {
					return nil, core.NewUnauthenticatedError("user or password is not valid")
				}

				authing.GetUserToken().SetSession(user)
			} else {
				key, _ := totp.Generate(totp.GenerateOpts{
					Issuer:      "xd", // user.Domain ?? "ARMORY"
					AccountName: user.Name,
				})

				// Convert TOTP key into a PNG
				var buf bytes.Buffer
				if img, err := key.Image(200, 200); err != nil {
					return nil, core.NewInternalError("failed to generate the QR code")
				} else {
					if err = png.Encode(&buf, img); err != nil {
						return nil, core.NewInternalError("failed to generate the QR code")
					}

					if _, err = model.GetUserModel().Update(ctx, &auth.User{
						Id:        user.Id,
						OtpSecret: key.Secret(),
						LoginTime: core.Now(),
					}); err != nil {
						return nil, core.NewInternalError("failed to update the user info")
					}

					logon.Totp = &auth.TOTP{
						// Secret: key.Secret(),
						QrCode: base64.StdEncoding.EncodeToString(buf.Bytes()),
					}
				}
			}
		}

		return logon, nil
	}
}

// Logout implements Interface.
func (s authingServer) Logout(ctx context.Context, in *pb.LogoutRequest) (*core.Null, error) {
	if len(in.User) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user id")
	}
	if subject := jwt.GetContextSubject(ctx); subject != in.User {
		return nil, core.NewInvalidArgumentError("invalid user id")
	}

	_, session := authing.GetUserToken().GetSession(ctx)
	if session == nil {
		return nil, core.NewUnauthenticatedError("invalid authenticated.")
	}

	authing.GetUserToken().DeleteSession(session)
	return &core.Null{}, nil
}

// GetUser implements Interface.
func (s authingServer) GetUser(ctx context.Context, in *pb.GetUserRequest) (*auth.User, error) {
	if len(in.Id) == 0 {
		return nil, core.NewInvalidArgumentError("not set the user name or email or phone number")
	}

	if user, err := model.GetUserModel().Get(ctx, in.Id); err != nil {
		return nil, core.NewNotFoundError("not found the user")
	} else {
		user.Password = ""
		user.Salt = ""
		user.PasswordAlgorithm = ""
		return user, nil
	}
}

// ListUser implements Interface.
func (s authingServer) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	resp := &pb.ListUserResponse{
		// Users:
		// TotalCount:
		// NextPageToken:
	}
	return resp, nil
}
