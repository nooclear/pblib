package pblib

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Bearer string

type AuthRes struct {
	Admin struct {
		ID      string `json:"id"`
		Created string `json:"created"`
		Updated string `json:"updated"`
		Avatar  int    `json:"avatar"`
		Email   string `json:"email"`
	} `json:"admin"`
	Token string `json:"token"`
}

func (p *PocketBase) AuthWithPass(email, pass string) ([]byte, error) {
	query := fmt.Sprintf("%s/api/admins/auth-with-password", p.Addr)
	id := struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}{
		Identity: email,
		Password: pass,
	}
	idBytes, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	data, err := request("POST", query, bytes.NewBuffer(idBytes))
	if err != nil {
		return nil, err
	}
	authRes := AuthRes{}

	err = json.Unmarshal(data, &authRes)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *PocketBase) AuthRefresh() ([]byte, error) {
	query := fmt.Sprintf(`%s/api/admins/auth-refresh`, p.Addr)
	data, err := request("POST", query, nil)
	if err != nil {
		return nil, err
	}
	authRes := AuthRes{}

	if err := json.Unmarshal(data, &authRes); err != nil {
		return nil, err
	}

	Bearer = authRes.Token
	return data, nil
}
