package pocketbase

import "fmt"

func (p *PocketBase) GetFile(collection, recordId, file string) []byte {
	query := fmt.Sprintf(`%s/api/files/%s/%s/%s`, p.Addr, collection, recordId, file)
	data := request("GET", query, nil)
	return data
}
