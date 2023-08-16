package auth

// import "github.com/iammuho/natternet/internal/user/application/auth/dto"

// type SigninQueryHandler struct {
// 	userService services.Services
// 	getStream   *getStream.Client
// }

// func NewSigninQueryHandler(service services.Services, getStream *getStream.Client) *SigninQueryHandler {
// 	return &SigninQueryHandler{
// 		userService: service,
// 		getStream:   getStream,
// 	}
// }

// func (s *SigninQueryHandler) Handle(req *dto.SigninReqDTO) (*jwt.JWTResponse, *errorhandler.Response) {
// 	res, err := s.userService.Signin(req)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return jwt.CreatePair(map[string]interface{}{
// 		"ID": res.UserID,
// 	})
// }
