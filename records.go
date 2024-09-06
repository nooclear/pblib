package pocketbase

import (
	"bytes"
	"fmt"
)

// filter = (email="user@example.com" && other="")
func (p *PocketBase) GetRecord(collection, filter string) []byte {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	if filter != "" {
		query += `/?filter=` + filter
	}
	data := request("GET", query, nil)
	return data
}

func (p *PocketBase) InsertRecord(collection string, data []byte) []byte {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	res := request("POST", query, bytes.NewBuffer(data))
	return res
}

func (p *PocketBase) UpdateRecord(collection, recordId string, data []byte) []byte {
	query := fmt.Sprintf("%s/api/collections/%s/records/%s", p.Addr, collection, recordId)
	res := request("PATCH", query, bytes.NewBuffer(data))
	return res
}

func (p *PocketBase) DeleteRecord(collection, recordId string) []byte {
	query := fmt.Sprintf("%s/api/collections/%s/records/%s", p.Addr, collection, recordId)
	res := request("DELETE", query, nil)
	return res
} //TODO
