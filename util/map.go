package util

import "encoding/json"

type FileMap struct {
	SourceName string   `json:"source_name"`
	DstPath    []string `json:"dst_path"`
}

func Marshal(v FileMap) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte) (*FileMap, error) {
	v := &FileMap{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	return v, nil
}
