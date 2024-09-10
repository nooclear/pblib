package pblib

import "fmt"

func (p *PocketBase) GetFile(collection, recordId, file string) ([]byte, error) {
	query := fmt.Sprintf(`%s/api/files/%s/%s/%s`, p.Addr, collection, recordId, file)
	res, err := request("GET", query, nil)
	return res, err
}
