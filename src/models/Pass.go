package models

type Pass struct {
	NewPass string `json:"newpass"`
	OldPass string `json:"oldpass"`
}
