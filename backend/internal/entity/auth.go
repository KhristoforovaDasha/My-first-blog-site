package entity

type AuthManager interface {
	MakeAuth(userId uint) (string, error)
	FetchAuth(tknString string) (*map[string]string, error)
}
