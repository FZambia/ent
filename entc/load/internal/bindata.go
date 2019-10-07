// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/main.tmpl
// schema.go
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x51\x5d\x6b\xdb\x30\x14\x7d\xb6\x7e\xc5\x99\xe9\xa8\x5d\x52\xa5\xed\xdb\x06\x79\x28\x6d\x06\x19\x5b\x3b\x48\x61\x0f\x5d\x29\x8a\x7d\x9d\x88\x3a\x92\x77\xa5\x94\x05\xa1\xff\x3e\x24\x27\x61\x7b\xb2\xa5\x73\xee\xf9\xd0\x0d\x61\x7a\x21\xee\xec\xb0\x67\xbd\xde\x78\xdc\x5c\x5d\x7f\xba\x1c\x98\x1c\x19\x8f\x2f\xaa\xa1\x95\xb5\x6f\x58\x98\x46\xe2\xb6\xef\x91\x49\x0e\x09\xe7\x77\x6a\xa5\x78\xda\x68\x07\x67\x77\xdc\x10\x1a\xdb\x12\xb4\x43\xaf\x1b\x32\x8e\x5a\xec\x4c\x4b\x0c\xbf\x21\xdc\x0e\xaa\xd9\x10\x6e\xe4\xd5\x11\x45\x67\x77\xa6\x15\xda\x64\xfc\xdb\xe2\x6e\xfe\xb0\x9c\xa3\xd3\x3d\xe1\x70\xc7\xd6\x7a\xb4\x9a\xa9\xf1\x96\xf7\xb0\x1d\xfc\x3f\x66\x9e\x89\xa4\xb8\x98\xc6\x28\x44\x08\x68\xa9\xd3\x86\x50\x6e\x95\x36\x25\x62\x14\xd3\x29\xee\x52\x9e\x35\x19\x62\xe5\xa9\xc5\x6a\x8f\x73\x32\xbe\x39\x5d\x9d\x4b\xdc\x3f\xe2\xe1\xf1\x09\xf3\xfb\xc5\x93\x14\x83\x6a\xde\xd4\x9a\x90\x34\x84\xd0\xdb\xc1\xb2\x47\x25\x8a\xd2\xba\x52\x14\xe5\x6a\xef\x29\xfd\x84\x00\x4f\xdb\xa1\x57\x9e\x50\x8e\x2c\x97\x2d\x33\x34\xb0\x36\xbe\x43\xf9\xf1\x77\x09\xf9\xe3\xa0\x18\xa3\xa8\x73\xcc\xb3\x95\x72\x84\xcf\x33\xe4\xef\x11\x4f\xb3\xef\x8a\xe1\x9a\x0d\x6d\x95\xc3\x0c\xcf\x2f\x64\xbc\x5c\x18\x4f\xdc\xa9\x86\x42\x96\x66\x65\xd6\x84\xb3\xd7\x09\xce\x8c\xda\x66\x19\xf9\xa0\xb6\xe4\x92\x7e\x51\x84\x70\x79\xd0\x8f\x51\xa6\xc3\x29\x8a\x0b\xb1\x3c\xcc\xc4\x38\xc9\x5a\x64\x5a\x5c\xc6\x28\xa2\x10\xdd\xce\x34\xb9\x73\x55\x23\x88\x22\x05\xe9\xb5\x21\x87\xe7\x97\xe7\x97\x54\x5a\x14\x9d\x65\xbc\x4e\x0e\xf9\x92\xef\x18\xe5\x98\x37\x88\xa2\x58\x4d\x40\xcc\x09\xfb\xae\xd8\x6d\x54\xbf\xcc\x60\x35\x72\x6a\x51\x14\xba\xcb\x8c\x0f\x33\x18\xdd\xe7\x99\xa2\x53\xba\xaf\x88\x39\xc1\xa9\xc2\xe8\x3b\x83\x1a\x06\x32\x6d\x95\x8f\x13\xac\x6a\x91\x50\xeb\xe4\xd2\xb7\x76\xe7\xe5\x4f\xd6\x9e\xaa\xbc\x0f\xf9\xd5\x6a\x73\x24\x8e\x71\xab\xf2\x97\x29\xeb\xba\x3e\x75\x3b\xba\x24\x7b\xcb\xb9\xe4\xa8\x45\xcc\xa3\xd6\xd2\xb3\x36\xeb\xc4\x91\xf3\xc4\xa9\xea\x3a\x73\xe6\x7f\xb4\xaf\xae\xb3\xd2\x7f\x5b\x1f\x4b\x8d\x4b\x3f\x3c\x66\x8c\xe2\x6f\x00\x00\x00\xff\xff\xe4\x6e\x0c\x4d\x4b\x03\x00\x00")

func templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateMainTmpl,
		"template/main.tmpl",
	)
}

func templateMainTmpl() (*asset, error) {
	bytes, err := templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/main.tmpl", size: 843, mode: os.FileMode(420), modTime: time.Unix(1567406598, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\xdd\x6f\xdb\x38\x12\x7f\xb6\xfe\x8a\xa9\x81\x06\x56\xe0\x95\x7b\x8b\xc3\x01\xa7\xc2\x0f\x8b\x6e\x16\xc8\xed\x35\x2d\x36\xb9\x7b\x09\x82\xac\x2c\x0d\x6d\x36\x12\xa5\x92\x74\x9a\x6c\x90\xff\x7d\x31\x43\x52\x1f\x96\x9d\x7e\x26\x2f\xb1\x86\x33\xc3\x99\xdf\x7c\x4a\x8b\x05\xbc\xa9\x9b\x7b\x2d\xd7\x1b\x0b\x3f\xbf\xfa\xc7\xbf\x7f\x6a\x34\x1a\x54\x16\x7e\xcb\x72\x5c\xd5\xf5\x0d\x9c\xaa\x3c\x81\x5f\xca\x12\x98\xc9\x00\x9d\xeb\x5b\x2c\x92\x68\xb1\x80\x8b\x8d\x34\x60\xea\xad\xce\x11\xf2\xba\x40\x90\x06\x4a\x99\xa3\x32\x58\xc0\x56\x15\xa8\xc1\x6e\x10\x7e\x69\xb2\x7c\x83\xf0\x73\xf2\x2a\x9c\x82\xa8\xb7\xaa\x20\x15\x52\x31\xcb\x7f\x4f\xdf\x9c\x9c\x9d\x9f\x80\x90\x25\x06\x9a\xae\x6b\x0b\x85\xd4\x98\xdb\x5a\xdf\x43\x2d\xc0\xf6\xee\xb3\x1a\x31\x89\xa2\x26\xcb\x6f\xb2\x35\x42\x59\x67\x45\x14\xc9\xaa\xa9\xb5\x85\x59\x34\x99\xa2\xca\xeb\x42\xaa\xf5\xe2\x83\xa9\xd5\x34\x9a\x4c\x45\x65\xe9\x9f\x46\x51\x62\x6e\xa7\x51\x34\x99\xae\xa5\xdd\x6c\x57\x49\x5e\x57\x0b\xe1\x1d\x96\x2a\xdf\xae\x32\x5b\xeb\x05\x2a\xe6\xff\x1c\xcf\xc2\xe4\x1b\xac\xb2\x05\x16\x6b\xfc\x1a\x7e\x21\xb1\x2c\xa6\x51\x1c\x11\x0a\xe7\x4c\x03\x8d\x1e\x7f\x03\x99\x02\x54\x36\xf1\x07\x76\x93\x59\xf8\x94\x19\x76\x13\x0b\x10\xba\xae\x20\x83\xbc\xae\x9a\x52\x12\xd6\x06\x35\x78\x28\x92\xc8\xde\x37\x18\x54\x1a\xab\xb7\xb9\x85\x87\x68\x72\x96\x55\x08\xe1\xcf\x58\x2d\xd5\xba\x7d\xfc\x93\x40\x4a\xa7\x2a\xab\x70\x5e\x57\xd2\x62\xd5\xd8\xfb\xe9\x9f\xd1\xe4\x4d\xad\x84\x0c\x7c\x64\x50\x8f\xe0\x85\x72\xa6\x0c\xc5\x4e\x8a\x35\x9a\xa0\xfc\xf2\xea\x98\x9e\x77\xee\x22\xbc\xcc\x50\xea\x37\x82\xc4\x74\x52\xfc\x3c\x94\x62\xd4\x76\xc4\x4e\x55\x81\x77\xe1\xba\xcb\xab\x63\x7e\x1e\x8a\x49\xc7\x32\x94\x3b\x67\x68\xfc\xa5\x97\x57\xc7\xbd\xe7\x20\xe7\xd0\xbb\xde\x73\xeb\x23\xc7\xed\x7d\x6d\xa4\x95\xb5\x82\x02\x4d\xae\xe5\x0a\x0d\x64\xc0\xdc\xd0\x84\x23\x9f\xcd\x2e\xec\x3e\x38\xad\x5c\x17\x9e\x9e\xd5\x52\x59\x80\xc5\xc2\x2b\x62\xdb\x83\x16\x47\x2a\xa5\xb1\x49\x34\x79\x2b\xef\xb0\x38\x55\x24\xb2\xaa\xeb\x12\xb8\x9c\x0a\x99\x67\x16\x0d\x48\xd1\x13\xa0\xd4\xa9\x88\xfb\x27\xa9\x9c\xa0\x54\xa7\x5e\xaf\xbb\xab\x22\xd2\xf0\x2e\x47\x72\x77\x39\x77\x1d\x36\xe3\x2c\x75\xf4\x6f\x48\x52\x27\x78\x20\x47\x77\x93\xf4\x70\x96\x9e\x2a\x51\x77\x6c\xc7\xec\x73\x72\x71\xdf\x20\x1f\x78\x31\xba\x70\x28\x76\x91\xf5\x94\x1f\xba\xcd\x66\x3b\xb9\x7d\x2e\xff\xea\xd9\x78\x2c\x95\xfd\xd7\x3f\x47\x52\x46\xfe\xb5\x73\xd9\x89\xda\x56\x6d\x49\xc0\xe5\xd5\xf0\xba\x50\x14\xc4\x34\x94\xfb\x9f\x92\x1f\xb7\xed\x85\x1c\x67\x18\x5d\xb7\x65\xa6\xa1\xe0\x99\x2c\xcb\x6c\x55\xe2\x93\x82\xca\x33\x0d\x45\xdf\x35\x94\x9c\x59\xf9\xa4\x68\xed\x99\x86\xa2\xbf\xa2\xc8\xb6\xa5\x7d\xda\xdc\xc2\x31\xed\x38\xda\x14\x99\xc5\x20\x7f\xc8\x51\x66\xba\xde\xab\xe0\xb4\xaa\xb6\xb6\xf5\xf8\x80\x02\x19\x98\x86\xb2\xff\xcf\x4a\x59\x50\x8b\xe6\x10\x71\x51\x8c\x65\x6f\x5b\xa6\xdd\x46\x52\xeb\x6c\x8d\xbf\xe3\xfd\x13\x79\x64\x1c\xd3\xf5\x0d\xde\x0f\xa5\xdb\x5e\xe0\xf2\x69\xf8\x18\xa4\x43\x37\xd9\xb9\x18\x15\x91\x6f\x9f\xf4\xd8\x04\xa6\x3d\xfd\xab\xdf\xf2\x76\xca\xfa\xce\xa2\xa6\x14\xf0\xc5\xe9\xfa\x48\x81\x42\x2a\x2c\xf6\xf6\xb4\xbe\xae\xae\xa2\xdb\x1a\xf3\xb0\x1c\xaa\xaa\xb6\xf2\x87\x7c\xe3\x5a\xa7\xb2\xde\xa7\x70\x54\xdd\x6f\xea\xaa\xa2\x55\x66\x87\x31\x77\xe4\x9d\x18\xdc\xac\xdf\x67\x76\xb3\xcb\xdb\xdc\xac\xaf\x9b\xcc\x6e\x76\x2a\xb9\x5a\x61\x41\x0d\xce\x03\x1e\x6a\xd7\x93\xf7\xc0\xcc\xe3\x6f\xdc\x36\x99\xfc\x0d\x5d\x93\xe5\xf6\x34\xcd\x1f\x06\xdd\x97\x06\xed\x0f\x14\xee\xf2\x21\x9f\x46\x71\x3d\xbe\xfd\x0f\x14\x3e\xc5\xdd\x36\xd0\x31\x1f\x68\x78\x43\x78\xf7\xb5\xb8\x53\x75\x8b\xda\xe0\x2e\xab\x74\xe4\xdd\xeb\x3f\x6e\xa5\x1e\x45\x4d\x7b\xf2\x9e\xa8\xb9\x01\x39\x0e\x9b\xa3\x7f\x43\xdc\x9c\x60\x17\x38\xef\x69\x5b\xb7\x4f\x78\xea\x17\xaa\x76\x6c\x7c\x76\x89\xda\xe5\x3c\xb8\xc2\x9c\xe1\x27\x8e\x47\xae\x91\xf7\x86\x4c\x05\x8f\x48\xb9\x73\x8b\x7f\xb9\x15\xa7\xb1\xb5\x4e\x22\xb1\x55\x79\x90\x9c\x61\x01\xc7\xc4\x91\xfc\xda\x72\xc4\x3e\xc8\x0f\xd1\x44\x21\xa4\x4b\x38\xa2\xc7\x87\x68\x42\xa9\x95\xba\x34\xc0\x22\xb9\xc8\xd6\x73\xa2\xdd\x37\x98\xb6\x34\xca\xc6\x68\xc2\x59\xdd\x12\xe9\x81\x88\x0e\xb1\xd4\x11\xdd\x03\x91\x7d\x1e\xa4\x4c\xf6\x0f\x44\x0f\x31\x4f\x89\x1e\x1e\xdc\x81\xf0\xfa\xf9\x40\x78\xfd\x8f\xd1\x44\x0a\xd0\x28\xc8\x64\x77\xf2\x9a\x1f\x5f\x2c\x41\xc9\x92\xdc\x99\x28\x24\x32\x2c\x5b\xf7\x35\x8a\x98\x45\x35\xda\xad\x56\xa0\xb0\x43\xd6\x75\xc3\x31\xb4\xae\x9b\x3e\x8d\x2d\xcb\xce\x44\x11\xf6\x99\x3e\xba\x33\xb7\x1b\xcf\x01\xb5\xa6\xe7\x87\x68\x62\xd8\xe8\x23\xa6\x3f\x0c\xf0\xe3\x3f\xd1\x81\x48\x4b\xd1\xf0\x84\x28\xf3\x41\x70\xc2\x89\x8f\x10\x2f\x2f\x69\xff\x80\x29\xc3\x90\x84\xa3\x2e\x2e\x61\x05\x49\x3b\x1b\xc2\xbe\x11\x4d\xda\x2d\xa3\x3b\x0d\x14\xb6\x32\x0c\xea\xb4\xb5\xb2\x1d\xdd\xd1\xa4\x37\x73\x53\x7f\xdc\x51\xe8\xbc\x1b\xe8\x7c\x5e\xa2\x9a\x89\x22\xe9\xa8\x31\x31\xf9\x65\x23\xed\x6c\x0f\xeb\x87\x0b\x38\xfb\xd7\x5f\x4b\x52\xf6\x6f\xb0\xa8\x74\x9c\xed\x34\x6e\x2d\x6e\x29\x6d\x6e\x19\xc1\x60\xc3\xb2\x4b\xa8\x90\x36\xb2\x9c\x83\xa8\x6c\x72\x42\x21\x15\xb3\x69\x25\x8d\xa1\x02\xe6\xd6\x21\x49\x48\xd4\xda\x67\xce\xcb\x8f\xd3\x39\xe9\xa2\x90\xc6\xad\x6e\x5a\x4c\xd3\x25\xf0\x46\x4a\xde\xd2\xa6\x1a\xbf\x76\xf4\x17\x4b\x78\xc5\xd7\x19\xc1\x74\x58\xc2\x11\x1d\xf4\x33\xd7\x88\x39\x99\xe1\xd3\xf7\x6d\xa6\xcd\x26\x2b\xfd\x7b\x24\xbf\x4e\x23\x4f\xfd\xde\x7b\xa9\x54\x16\x35\xbd\xe6\xd2\xaf\x1a\x32\xf8\xcf\xf9\xbb\x33\x12\xe6\xee\x98\x67\x0a\x56\x94\xdc\x24\x5a\x38\x16\x52\xe0\x85\xeb\xd5\x07\xcc\xad\xff\xe7\xf3\x7e\x70\xe9\xcc\x84\xbb\xa9\xe9\xfa\x9b\x62\x98\xad\xe0\xf2\x6a\x75\x6f\x91\xd3\xbf\x5f\x02\x5c\x01\x4e\x96\x5c\x75\xef\xaa\x69\x58\x53\xdc\xe3\x2c\xee\x77\x17\x7a\x5f\xd2\x98\xdb\x99\xff\x2c\xc0\xed\xe7\x9d\xf0\x37\xc7\x31\x23\xcc\x22\x0e\x63\xba\x30\x5d\x82\x49\xa8\x90\x5d\xab\x0d\xbc\xaf\xf9\xf0\xc5\xfe\xc0\xa2\xd6\xac\xc2\x35\xec\x56\x4d\x26\x90\x3b\x7b\xd0\xd1\xde\x71\x40\x4d\x3f\x3f\x3c\x38\x2f\x3f\xa6\xf0\xf2\x96\xd2\xc1\x15\x38\x89\xbb\x94\xa0\x74\xb9\x9e\x03\xe7\x84\xce\xd4\x1a\xb9\xd7\x18\x97\x05\x89\x9b\x28\x4b\xc8\x9a\x06\x55\x31\xf3\x84\x79\xd7\xd5\x7b\x0d\x67\x16\xc7\x3e\xcb\xfc\x7b\x74\xdf\x01\xff\xfa\xfd\x9c\x2e\xc8\xe2\xae\x73\xc2\xdb\xc0\x8a\xfd\x81\x2c\xee\x06\xd6\xb2\x83\xe1\xb3\x40\xcf\xc5\xd3\x60\xfe\x11\xff\x22\x0d\x6e\xb2\xa6\xc0\x3a\x1c\x04\x44\x75\xa1\x4d\x99\xea\x7e\x33\x39\x34\x3b\x22\x77\x6d\xee\x71\xd0\xff\x69\xde\x26\x3e\x8f\x67\x26\xf6\xd5\xd4\xe5\x0b\xff\x34\xbe\x90\x6d\xed\xb3\xd3\x0f\x83\x7e\xa6\xfb\x92\x98\x19\x38\x76\x39\x1d\xc3\x28\xeb\x76\x6b\x83\x8b\x81\xa0\xe1\x97\xf7\x41\x9c\xf8\xa5\xff\x0b\xa2\xf4\xd5\x01\x92\x73\xa8\x7a\xf1\x71\x9f\x0d\x48\xa1\xdf\x39\xfa\x46\x78\xe3\xab\x3b\x8a\xd1\xd8\x84\xaf\xb7\x81\x8c\x60\x2b\x3e\xcc\x41\x74\x46\xb8\xab\x9d\x4e\x6a\x6b\xde\x84\x6e\xac\x0e\xb3\x9b\xd8\xf6\x58\xf3\x0d\xe6\xb0\x3d\xd4\x64\xdb\xf7\xb8\x25\x1c\x85\xdf\x4e\x29\xe7\x9e\x9f\x39\x1f\x38\xad\xc2\x97\x1c\x26\x5a\xed\xb2\x6a\xd2\xfb\x4c\x93\x82\x9c\x77\xca\x7d\x46\xf6\x33\xdb\xe7\x28\x18\xe1\x31\xa1\xd8\x1c\x84\xff\x79\x92\x60\x3f\xfc\x5f\x86\xfe\x0f\x4b\x85\x43\xc8\x07\x18\x99\xe7\x73\x00\xf6\x96\xb9\x76\x1c\x76\xf0\xc1\x27\x9d\x35\xa6\xff\x02\xec\xe9\x99\x2a\x5c\xf6\x07\x42\x85\x76\x53\x17\xf0\x49\xda\x0d\x68\xcc\xeb\x5b\xd4\x54\xf1\xa8\xcc\x56\x23\xa8\x1a\x9a\x4c\xc9\xdc\xd0\xeb\x74\xe5\x1a\x86\x54\x6b\x5f\xf6\xbd\x70\x89\xa2\x1b\xb4\x0f\xe0\x89\x31\x5c\x5e\x75\xdf\xde\x1e\x63\x98\x89\xb0\xf9\xb7\xe4\xdd\x01\x59\xa0\x40\x0d\xa4\x7e\x16\xbb\xfe\x29\xe0\x96\xa3\xe6\x8c\x9b\xc5\xaf\xe1\x76\x10\x04\x92\x5f\x0e\x62\xf0\xf2\x22\x78\xe7\x8c\xf7\xa1\x10\xc5\x1c\x6e\xb9\x00\x44\xc0\x96\xb0\x73\xb9\x48\x1d\x39\x84\xb3\x48\x82\x03\xf3\x1d\x74\xdd\x44\x1a\x81\xeb\xc8\xdf\x0b\x65\x7f\xcc\x8e\x16\x0a\x37\x17\x1d\x70\xc4\xf8\x1c\xb8\x0d\xbc\x19\x40\xe7\x60\x43\x3f\x8f\xf7\xa2\xd6\x17\x1e\x03\x17\x26\xdd\x08\xba\x70\xf0\xbd\xe0\x0d\x47\xfc\x08\xbe\x30\x91\x1d\x80\xcc\xfc\x8c\x08\x06\xa7\xf6\x60\x28\xdb\x91\xff\x14\x8a\xc1\x9b\x11\x8e\xdc\x6f\xc7\x28\x3a\xf2\xf7\x62\xd8\x1f\xbf\x23\x04\xdd\xcc\x74\xf8\xbd\xed\x26\xf7\xb3\xe0\xe7\xdc\xd9\x83\x9e\x33\xe2\x69\xec\x9c\x17\x1d\x72\xec\x5e\xbb\x44\x5b\xe8\xaf\xd1\xf1\xe0\x89\xac\xa2\x41\x61\x93\xdf\xa5\x2a\x66\x31\xbd\x02\x85\xf3\xf7\x96\x77\x96\x89\x85\x25\xd8\xe4\xa4\xc4\x6a\x36\xe8\xc2\x36\x7a\x8c\xfe\x0e\x00\x00\xff\xff\xc1\xa3\x85\x8a\xac\x1c\x00\x00")

func schemaGoBytes() ([]byte, error) {
	return bindataRead(
		_schemaGo,
		"schema.go",
	)
}

func schemaGo() (*asset, error) {
	bytes, err := schemaGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.go", size: 7340, mode: os.FileMode(420), modTime: time.Unix(1570458335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/main.tmpl": templateMainTmpl,
	"schema.go":          schemaGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"schema.go": &bintree{schemaGo, map[string]*bintree{}},
	"template": &bintree{nil, map[string]*bintree{
		"main.tmpl": &bintree{templateMainTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
