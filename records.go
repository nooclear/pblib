package pblib

import (
	"bytes"
    "fmt"
)

type RecordArgs struct {
	Page int
	PerPage int
	Sort string
	Filter string
}

func (p *PocketBase) GetRecord(collection string, args RecordArgs) ([]byte, error) {
	query := fmt.Sprintf("%s/api/collections/%s/records/?page=%d&perPage=%d", p.Addr, collection, args.Page, args.PerPage)
	if args.Sort != "" {
		query += "&sort="+args.Sort
	}
	if args.Filter != "" {
		query += fmt.Sprintf("&filter=(%s)", args.Filter)
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
