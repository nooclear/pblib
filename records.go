package pblib

import (
	"bytes"
    "fmt"
)


func (p *PocketBase) GetRecord(collection, args string) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records", p.Addr, collection)
	if args != "" {
		query += fmt.Sprintf("/?" + args)
	}
	res, err := request("GET", query, nil)
	return res, err
} // working on args/filter

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
