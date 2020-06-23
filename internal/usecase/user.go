package usecase

import (
	"errors"
	"log"
	"strings"

	cah "github.com/j4rv/cah/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type userController struct {
	store cah.UserStore
}

// NewUserUsecase returns a cah.UserUsecases
func NewUserUsecase(uc *cah.Usecases, store cah.UserStore) cah.UserUsecases {
	return &userController{store: store}
}

func (uc userController) Register(name, pass string) (cah.User, error) {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return cah.User{}, errors.New("Username cannot be empty")
	}
	if pass == "" {
		return cah.User{}, errors.New("Password cannot be empty")
	}
	_, err := uc.store.ByName(name)
	if err == nil {
		return cah.User{}, errors.New("That username already exists. Please try another")
	}
	passHash, err := userPassHash(pass)
	if err != nil {
		//Never log passwords! But this one caused an error and will not be stored, its an ok exception
		log.Println("ERROR while trying to hash password.", pass, err)
		return cah.User{}, errors.New("That password could not be protected correctly. Please try another")
	}
	return uc.store.Create(name, passHash)
}

func (uc userController) ByID(id uint) (cah.User, bool) {
	u, err := uc.store.ByID(id)
	return u, err == nil
}

func (uc userController) Login(name, pass string) (cah.User, bool) {
	trimmedName := strings.TrimSpace(name)
	u, err := uc.store.ByName(trimmedName)
	if err != nil {
		return cah.User{}, false
	}
	if !userCorrectPass(pass, u.Password) {
		return cah.User{}, false
	}
	return u, true
}

// internal

const userPassCost = 10

func userPassHash(p string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), userPassCost)
	return b, err
}

func userCorrectPass(pass string, storedhash []byte) bool {
	err := bcrypt.CompareHashAndPassword(storedhash, []byte(pass))
	return err == nil
}
