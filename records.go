package pblib

import (
	"bytes"
	"fmt"
)

// GetRecord filter = (email="user@example.com" && other="")
func (p *PocketBase) GetRecord(collection, filter string) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	if filter != "" {
		query += fmt.Sprintf(`/?filter=(%s)`, filter)
	}
	res, err := request("GET", query, nil)
	return res, err
}

func (p *PocketBase) ViewRecord(collection, recordID string) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records/%s", p.Addr, collection, recordID)
	res, err := request("GET", query, nil)
	return res, err
}

func (p *PocketBase) InsertRecord(collection string, data []byte) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	res, err := request("POST", query, bytes.NewBuffer(data))
	return res, err
}

func (p *PocketBase) UpdateRecord(collection, recordId string, data []byte) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records/%s", p.Addr, collection, recordId)
	res, err := request("PATCH", query, bytes.NewBuffer(data))
	return res, err
}

func (p *PocketBase) DeleteRecord(collection, recordId string) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records/%s", p.Addr, collection, recordId)
	res, err := request("DELETE", query, nil)
	return res, err
}
