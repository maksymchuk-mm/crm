package service

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type SignInInput struct {
	login    int64
	Password string
}
