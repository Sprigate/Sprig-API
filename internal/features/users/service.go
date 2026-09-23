package users

import (
	"errors"
	//"log"
	"sprig/internal/common"
	"sprig/internal/model"
	"strings"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(req CreateUserRequest) (*model.User, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("Nama user tidak boleh kosong")
	}

	if strings.TrimSpace(req.Email) == "" {
		return nil, errors.New("Email tidak boleh kosong")
	}

	if strings.TrimSpace(req.Password) == "" {
		return nil, errors.New("Password tidak boleh kosong")
	}

	encryptedPass, err := common.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: encryptedPass,
	}

	err = s.repository.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAllUser() ([]model.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(id uint64, req CreateUserUpdateRequest) (*model.User, error) {

	// Ambil entity lama, agar PATCH bisa tetap mempertahankan
	// data jika client ingin mengosongkannya (""/nil)
	existingUser, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	// jika field tidak kosong / nil, maka timpa saja
	if req.Name != nil && strings.TrimSpace(*req.Name) != "" {
		existingUser.Name = *req.Name
	}

	if req.Email != nil && strings.TrimSpace(*req.Email) != "" {
		existingUser.Email = *req.Email
	}

	if req.Password != nil && strings.TrimSpace(*req.Password) != "" {

		// bcrypt hashing (common.)
		encryptedPass, err := common.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}

		existingUser.Password = string(encryptedPass)

		// // DEBUG
		// match := common.CheckPassword(*req.Password, encryptedPass)
		// log.Printf("cocok : %t", match)
	}

	if err = s.repository.Save(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	err := s.repository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
