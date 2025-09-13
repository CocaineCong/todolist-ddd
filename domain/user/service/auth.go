package service

//
// type AuthenticationService interface {
// 	AuthenticateByToken(ctx context.Context, token string) (*entity.User, error)
// 	AuthenticateByCredentials(ctx context.Context, username, password string) (*entity.User, error)
// }
//
// type AuthenticationServiceImpl struct {
// 	userRepo     repository.User
// 	tokenService auth.TokenService // 依赖抽象接口
// }
//
// func (a *AuthenticationServiceImpl) AuthenticateByToken(ctx context.Context, token string) (*entity.User, error) {
// 	// 1. 验证Token
// 	claims, err := a.tokenService.ParseToken(token)
// 	if err != nil {
// 		return nil, errors.New(consts.UserTokenParseFailed)
// 	}
//
// 	// 2. 业务逻辑：根据用户ID获取用户
// 	user, err := a.userRepo.GetUserByID(ctx, claims.Id)
// 	if err != nil {
// 		return nil, errors.New(consts.UserNotFound)
// 	}
//
// 	// 3. 业务验证：检查用户状态
// 	if !user.IsActive() {
// 		return nil, errors.New(consts.UserNotActive)
// 	}
//
// 	return user, nil
// }
