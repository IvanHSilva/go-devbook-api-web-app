package models

import (
	"errors"
	"strings"
)

type Post struct {
	Id         uint64 `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Content    string `json:"content,omitempty"`
	AuthorId   uint64 `json:"authorid,omitempty"`
	AuthorName string `json:"authorname,omitempty"`
	Likes      uint64 `json:"likes"`
	RegDate    string `json:"regdate,omitempty"`
}

func (post *Post) CheckPost() error {
	//
	if err := post.validate(); err != nil {
		return err
	}

	if err := post.trim(); err != nil {
		return err
	}

	return nil
}

func (post *Post) validate() error {
	//
	if post.Title == "" {
		return errors.New("título é obrigatório e deve ser informado")
	}

	if post.Content == "" {
		return errors.New("conteúdo é obrigatório e deve ser informado")
	}

	if post.RegDate == "" {
		return errors.New("data de cadastro é obrigatória e deve ser informada")
	}

	return nil
}

func (post *Post) trim() error {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
	post.RegDate = strings.TrimSpace(post.RegDate)

	return nil
}
