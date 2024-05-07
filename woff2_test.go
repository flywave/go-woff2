package woff2

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	data, err := ioutil.ReadFile("NotoSans-Regular.ttf")
	if err != nil {
		t.FailNow()
	}

	si := MaxCompressedSize(data)
	if si == 0 {
		t.FailNow()
	}

	out := ConvertTTFToWOFF2(data)

	if out == nil {
		t.FailNow()
	}

	ioutil.WriteFile("./NotoSans-Regular.wof", out, fs.ModePerm)

	os.Remove("./NotoSans-Regular.wof")
}
