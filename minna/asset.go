package minna

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets66006826044fa749a8098d688840a3d8dbeeebca = "Hello World\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"test.txt"}}, map[string]*assets.File{
	"/test.txt": &assets.File{
		Path:     "/test.txt",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1507658456, 1507658456000000000),
		Data:     []byte(_Assets66006826044fa749a8098d688840a3d8dbeeebca),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1521919554, 1521919554168840907),
		Data:     nil,
	}}, "")
