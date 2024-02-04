package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"hristoforovada-project/backend/internal/entity"
)

type UserService struct {
	userRepo entity.UserRepository
}

func NewUserService(userRepo entity.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) Get(id uint) (*entity.User, error) {
	userDB, err := s.userRepo.Get(id)
	if err != nil {
		return nil, err
	} else {
		return userDB, nil
	}
}

func (s *UserService) Update(user *entity.User) error {
	userDB, err := s.userRepo.Get(user.Id)
	if err != nil {
		return err
	}

	var newUserPasswordHash string
	if comparePasswordWithHash(user.Password, userDB.Password) != nil {
		newUserPasswordHash, err = generatePasswordHash(user.Password)
	}

	user.Password = newUserPasswordHash
	err = s.userRepo.Update(user)
	return err
}

func comparePasswordWithHash(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func generatePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash), err
}

func (s *UserService) Delete(user *entity.User) error {
	err := s.userRepo.Delete(user.Id)
	return err
}

func (s *UserService) Register(userReg *entity.UserRegister) error {
	if userReg.Login == "" {
		return entity.ErrInvalidLogin
	}
	if len(userReg.Password) < 8 {
		return entity.ErrInvalidPassword
	}

	passwordHash, err := generatePasswordHash(userReg.Password)
	if err != nil {
		return err
	}

	userReg.Password = passwordHash
	user := entity.User{UserRegister: *userReg}
	err = s.userRepo.Create(&user)
	return err
}

func (s *UserService) Login(userLog *entity.UserRegister) (uint, error) {
	userDB, err := s.userRepo.GetByLogin(userLog.Login)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%+v", userDB)
	fmt.Println()
	err = comparePasswordWithHash(userLog.Password, userDB.Password)
	if err != nil {
		return 0, err
	}

	return userDB.Id, nil
}
