package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"postgres/internal/model"
)

// ToAuthFromDomain конвертирует domain модель в gRPC модель
// Импортируйте вашу gRPC модель, например:
// authGrpc "week_1/grpc/pkg/auth_v1"
//
// func ToAuthFromDomain(auth *model.Auth) *authGrpc.User {
// 	if auth == nil {
// 		return nil
// 	}
//
// 	return &authGrpc.User{
// 		Id:        auth.ID,
// 		Name:      auth.Name,
// 		Password:  auth.Password,
// 		CreatedAt: timestamppb.New(auth.CreatedAt),
// 	}
// }

// ToDomainFromAuth конвертирует gRPC модель в domain модель
// func ToDomainFromAuth(user *authGrpc.User) *model.Auth {
// 	if user == nil {
// 		return nil
// 	}
//
// 	return &model.Auth{
// 		ID:        user.Id,
// 		Name:      user.Name,
// 		Password:  user.Password,
// 		CreatedAt: user.CreatedAt.AsTime(),
// 	}
// }

// Временная заглушка, чтобы избежать ошибок компиляции
var _ = timestamppb.Now()
var _ = model.Auth{}
