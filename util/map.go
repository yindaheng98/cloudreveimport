package util

import (
	"encoding/json"
	"time"
)

const (
	ImportFileCommand       = "ImportFile"
	DeleteFileCommand       = "DeleteFile"
	UpdateFileStatCommand   = "UpdateFileStat"
	UpdateFolderTimeCommand = "UpdateFolderTime"
)

type Command struct {
	Command            string    `json:"command"`
	SourceName         string    `json:"source_name"`
	DstPath            []string  `json:"dst_path"`
	CreatedAtTimestamp int64     `json:"created_at"`
	UpdatedAtTimestamp int64     `json:"updated_at"`
	Size               uint64    `json:"size"`
	CreatedAt          time.Time `json:"-"`
	UpdatedAt          time.Time `json:"-"`
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func Marshal(v Command) ([]byte, error) {
	v.CreatedAtTimestamp = max(v.CreatedAt.Unix(), 0)
	v.UpdatedAtTimestamp = max(v.UpdatedAt.Unix(), 0)
	return json.Marshal(v)
}

func Unmarshal(data []byte) (*Command, error) {
	v := &Command{}
	if err := json.Unmarshal(data, v); err != nil {
		return nil, err
	}
	v.CreatedAt = time.Unix(v.CreatedAtTimestamp, 0)
	v.UpdatedAt = time.Unix(v.UpdatedAtTimestamp, 0)
	return v, nil
}
