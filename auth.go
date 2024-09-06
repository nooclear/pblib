package pocketbase

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Bearer string

func (p *PocketBase) AuthWithPass(email, pass string) string {
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
		panic(err)
	}

	data := request("POST", query, bytes.NewBuffer(idBytes))
	res := struct {
		Admin struct {
			ID      string `json:"id"`
			Created string `json:"created"`
			Updated string `json:"updated"`
			Avatar  int    `json:"avatar"`
			Email   string `json:"email"`
		} `json:"admin"`
		Token string `json:"token"`
	}{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}

	return res.Token
}

func (p *PocketBase) AuthRefresh() bool {
	query := fmt.Sprintf(`%s/api/admins/auth-refresh`, p.Addr)
	data := request("POST", query, nil)
	res := struct {
		Admin struct {
			ID      string `json:"id"`
			Created string `json:"created"`
			Updated string `json:"updated"`
			Avatar  int    `json:"avatar"`
			Email   string `json:"email"`
		} `json:"admin"`
		Token string `json:"token"`
	}{}

	if err := json.Unmarshal(data, &res); err != nil {
		panic(err)
	}
	// Time format time.Now().UTC().Format("2006-01-02 15:04:05.000Z")
	Bearer = res.Token
	return true
}
