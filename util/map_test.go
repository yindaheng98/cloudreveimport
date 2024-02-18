package util

import (
	"fmt"
	"testing"
	"time"
)

func Test_Marshal(t *testing.T) {
	data, _ := Marshal(Command{
		Command:    UpdateFileStatCommand,
		SourceName: "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg",
		DstPath:    []string{"www", "ggg", "ggg", "iii.png"},
		CreatedAt:  time.Now(),
	})
	t.Log(string(data))
}

func Test_Unmarshal(t *testing.T) {
	v, _ := Unmarshal(
		[]byte(`{"command":"UpdateFileStat","source_name":"/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg","dst_path":["www","ggg","ggg","iii.png"],"created_at":1707712939,"updated_at":0}`),
	)
	t.Log(fmt.Sprintf("%+v", v))
}
