package models

import (
	"errors"
	"strings"
)

// User model
type User struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	EMail   string `json:"email,omitempty"`
	Pass    string `json:"password,omitempty"`
	RegDate string `json:"regdate,omitempty"`
}

func (user *User) CheckUser() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.trim()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("nome é obrigatório e deve ser informado")
	}
	if user.EMail == "" {
		return errors.New("E-Mail é obrigatório e deve ser informado")
	}
	if user.Pass == "" {
		return errors.New("senha é obrigatória e deve ser informada")
	}
	if user.RegDate == "" {
		return errors.New("data de cadastro é obrigatória e deve ser informada")
	}
	return nil
}

func (user *User) trim() {
	user.Name = strings.TrimSpace(user.Name)
	user.EMail = strings.TrimSpace(user.EMail)
	user.Pass = strings.TrimSpace(user.Pass)
	user.RegDate = strings.TrimSpace(user.RegDate)
}
