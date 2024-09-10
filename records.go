package pblib

import (
	"bytes"
    "encoding/json"
    "fmt"
)

type RecordsArgs struct {
	Page int
	PerPage int
	Sort string
}

// GetRecord filter = email="user@example.com"
func (p *PocketBase) GetRecord(collection, filter string, args RecordsArgs) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	if filter != "" {
		query += fmt.Sprintf(`/?filter=(%s)`, filter)
	}
	arg, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	res, err := request("GET", query, bytes.NewBuffer(arg))
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
