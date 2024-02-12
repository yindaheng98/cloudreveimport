package util

import (
	"fmt"
	"testing"
)

func Test_Marshal(t *testing.T) {
	data, _ := Marshal(FileMap{
		SourceName: "/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg",
		DstPath:    []string{"www", "ggg", "ggg", "iii.png"},
	})
	t.Log(string(data))
}

func Test_Unmarshal(t *testing.T) {
	v, _ := Unmarshal(
		[]byte("{\"source_name\":\"/gallery/data/twitter/eumi_114/1426380913575743488_1.jpg\",\"dst_path\":[\"www\",\"ggg\",\"ggg\",\"iii.png\"]}"),
	)
	t.Log(fmt.Sprintf("%+v", v))
}
