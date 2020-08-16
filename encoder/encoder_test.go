package encoder

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const (
	testDataFolder = "testdata"
)

var (
	update = flag.Bool("update", false, "update golden files")
)

func checkErrorFatalf(t *testing.T, message string, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s err %v\n", message, err)
	}
}

func openImageFile(t *testing.T, filename string) io.Reader {
	t.Helper()
	path := filepath.Join(testDataFolder, filename)
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	checkErrorFatalf(t, fmt.Sprintf("cant open image file %s", path), err)
	return file
}

func openGoldenFile(t *testing.T, filename, source string, update bool) string {
	t.Helper()
	path := filepath.Join(testDataFolder, filename)
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	checkErrorFatalf(t, fmt.Sprintf("cant open golden file %s", path), err)
	defer file.Close()

	if update {
		err := os.Truncate(path, 0)
		checkErrorFatalf(t, fmt.Sprintf("cant open golden file %s", path), err)
		_, err = file.WriteString(source)
		checkErrorFatalf(t, fmt.Sprintf("cant write golden file %s", path), err)
		return source
	}

	content, err := ioutil.ReadAll(file)
	checkErrorFatalf(t, fmt.Sprintf("cant read golden file %s", path), err)
	return string(content)
}

func TestShouldEncodeImage(t *testing.T) {
	type args struct {
		filename string
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{"Should encode tux image", args{"tux.jpg"}, "tux.golden"},
		{"Should encode julio image", args{"julio.jpg"}, "julio.golden"},
		{"Should encode gopher-science.jpg image", args{"gopher-science.png"}, "gopher-science.golden"},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			reader := openImageFile(t, tC.args.filename)
			data, err := Encode(reader, 70)
			checkErrorFatalf(t, fmt.Sprintf("cant encode file %s", tC.args.filename), err)

			got := openGoldenFile(t, tC.want, data, *update)

			if !reflect.DeepEqual(got, data) {
				t.Fatalf("Encode() = %v want %v", got, data)
			}
		})
	}
}

func TestShouldEncodeImageFail(t *testing.T) {
	reader, _ := os.OpenFile("fail.jpeg", os.O_RDONLY, 0644)
	_, err := Encode(reader, 70)

	if err == nil {
		t.Fatalf("Encode() = %v want %v", err, errors.New("Encode should fail"))
	}
}
