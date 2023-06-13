package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// User model
type User struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	EMail   string `json:"email,omitempty"`
	Pass    string `json:"password,omitempty"`
	RegDate string `json:"regdate,omitempty"`
}

func (user *User) CheckUser(level string) error {
	//
	if err := user.validate(level); err != nil {
		return err
	}

	if err := user.trim(level); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(level string) error {
	//
	if user.Name == "" {
		return errors.New("nome é obrigatório e deve ser informado")
	}

	if user.EMail == "" {
		return errors.New("E-Mail é obrigatório e deve ser informado")
	}
	if err := checkmail.ValidateFormat(user.EMail); err != nil {
		return errors.New("E-Mail inválido")
	}

	if level == "ins" && user.Pass == "" {
		return errors.New("senha é obrigatória e deve ser informada")
	}

	if user.RegDate == "" {
		return errors.New("data de cadastro é obrigatória e deve ser informada")
	}

	return nil
}

func (user *User) trim(level string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.EMail = strings.TrimSpace(user.EMail)
	user.Pass = strings.TrimSpace(user.Pass)
	user.RegDate = strings.TrimSpace(user.RegDate)

	if level == "ins" {
		hash, err := security.Hash(user.Pass)
		if err != nil {
			return err
		}
		user.Pass = string(hash)
	}

	return nil
}
