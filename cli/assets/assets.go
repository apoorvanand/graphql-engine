// Code generated by go-bindata.
// sources:
// assets/unversioned/console.html
// assets/v1.0-alpha/console.html
// assets/v1.0/console.html
// DO NOT EDIT!

package assets

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsUnversionedConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5f\x6b\xfb\x36\x14\x7d\xff\x7d\x8a\x8b\xc6\xe8\xcb\x2c\xa5\x5d\x61\xc5\xb1\x03\x63\x30\x06\xdb\xa0\x0f\xeb\x5e\x8b\x2a\x5f\xdb\x37\x95\x25\x4f\x57\x49\xea\x95\x7c\xf7\xe1\x7f\x89\x9b\x0c\xc6\xca\x46\x5e\x74\xff\x48\xe7\x9e\x23\xe5\x38\xab\x63\x63\xc1\x6a\x57\xe5\x02\x5d\xb2\x63\xb1\xf9\x02\x90\xd5\xa8\x8b\x7e\x01\x90\x59\x72\xaf\x10\xd0\xe6\x82\x8c\x77\x02\x62\xd7\x62\x2e\xa8\xd1\x15\xaa\xd6\x55\x02\xea\x80\x65\x2e\xea\x18\x5b\x4e\x95\xe2\xe8\x83\xae\x50\x56\xde\x57\x16\x75\x4b\x2c\x8d\x6f\x54\xad\x79\x17\x74\x52\x05\xdd\xd6\x7f\xd8\x04\x5d\x45\x0e\x95\xf1\x8e\xbd\x45\xa5\x99\x31\xb2\x2a\xf5\xbe\xc7\x90\xc3\xb1\x6a\xc2\x67\x13\xa8\x8d\x63\x00\x07\x72\x85\x3f\xc8\xe7\x67\x74\x7b\xc8\xe1\x7d\xcc\x02\xe8\x96\x7e\xf2\x1c\x53\x78\x7f\x97\xd3\xfa\x78\xfc\x66\x51\x7d\xf4\x21\xa6\x20\xc6\x72\x1f\x1c\x8f\xe2\x54\x37\x96\x7e\xc7\xc0\xe4\xdd\x70\xc0\x39\x5c\x9c\x51\xe8\xa8\xbf\x6f\xe9\x29\xd8\xa1\xe7\x1c\x5e\xf7\x2c\xcf\xfa\x98\x5a\xce\x64\x0c\x32\xff\x8c\xdd\x38\xf3\x1c\x2d\x3a\x76\xc1\x3e\x06\x2c\xe9\x2d\x05\xa1\x16\xc3\x8e\x9a\xfd\xea\x0b\x4c\x41\x18\x4b\x62\xac\x1c\xd7\xa3\x5e\xea\x2c\x58\xa6\xe6\x7b\xcc\x5e\x7c\xd1\xcd\x82\xc6\xce\xe2\xb8\x96\x8d\x26\xf7\x83\x77\x11\x5d\x3c\x89\x59\x10\xb7\x56\x77\x29\xdc\x38\xef\xf0\x66\x3d\xa5\x7d\xab\x0d\xc5\x2e\x85\xd5\x9c\x89\x41\x3b\xa6\x38\x50\x9d\xaa\x20\xef\x56\x0c\x96\x1c\xea\x30\xb6\x1d\xaf\x80\x24\xd7\xfe\xf0\x37\x68\x2f\xd6\x9b\xd7\x6b\xb8\xdb\x4f\xc0\x65\x6a\x22\x39\x46\x05\xed\x81\x8a\x5c\x58\xaf\x0b\x72\x95\x98\x1e\xd3\x58\x30\x56\x33\xe7\xa2\xd5\x15\x26\x73\x03\x0c\xdb\xf3\x49\x59\x68\xc8\x25\x35\x52\x55\xc7\x14\x6e\x57\xab\x7d\x3d\x8f\x74\xa0\x22\xd6\x43\xee\xeb\xf5\x25\x9f\xd2\xe2\xdb\x9c\xd4\x96\x2a\x97\x50\xc4\x86\x53\x30\xe8\x22\x86\xb9\x54\x7a\x17\x93\x52\x37\x64\xbb\x14\x58\x3b\x4e\x18\x03\x95\x73\x79\xbb\xe3\x48\x65\x97\x98\x51\xbb\xcb\xdd\x27\x2a\xfd\xbd\xb6\xda\xcd\x6c\x2e\x19\x4c\x38\x4c\x7f\x62\x0a\x77\xd8\xac\x4f\xf9\x46\x87\x8a\x5c\x12\x7d\x9b\x42\xf2\xed\xb2\x62\xbc\xf5\x21\x85\xaf\x1e\xee\xfb\xdf\x39\xbf\xc0\xfc\x65\xd4\x4b\x4a\x39\x2b\xaa\xfa\x29\x4e\xfa\xaa\x82\xf6\xd3\xab\x5b\x2c\xe7\xeb\x98\x38\x89\x79\xe8\xc5\x1b\x11\x9b\xe5\x86\xb3\x05\x0d\xac\xb8\x46\x8c\x97\xbe\x63\x0a\xb7\x65\x69\xac\xdf\x15\xa5\xd5\x01\x07\xd7\xd1\x5b\xfd\xa6\x2c\xbd\xb0\x1a\xe8\xeb\x03\xb2\x6f\x50\xdd\xcb\xef\xe4\x4a\x19\xfe\x98\x96\x0d\x39\x69\x98\x85\xfa\x17\xb0\x9f\xb2\xbb\xfe\x0f\x3f\x38\xde\xc9\x16\x54\xcf\x7d\x00\x07\x53\xeb\xc0\x18\x73\xf1\xf4\xdb\x8f\xc9\x83\xf8\x68\x83\xc0\xc1\xfc\xf7\xe0\x7b\x74\x85\x0f\x72\x7b\x8d\xbe\x59\xda\xc9\xff\x3c\xc5\x20\xc1\x3f\xcd\x90\xa9\xd1\xc9\x32\xd5\x7f\xb8\x36\x5f\xfe\x0a\x00\x00\xff\xff\xdf\x70\xc8\xd7\xc0\x06\x00\x00")

func assetsUnversionedConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsUnversionedConsoleHtml,
		"assets/unversioned/console.html",
	)
}

func assetsUnversionedConsoleHtml() (*asset, error) {
	bytes, err := assetsUnversionedConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/unversioned/console.html", size: 1728, mode: os.FileMode(436), modTime: time.Unix(1530690087, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10AlphaConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5f\x6b\xfb\x36\x14\x7d\xff\x7d\x8a\x8b\xc6\xe8\xcb\x2c\xa5\x5d\x61\xc5\xb1\x03\x63\x30\x06\xdb\xa0\x0f\xeb\x5e\x8b\x2a\x5f\xdb\x37\x95\x25\x4f\x57\x49\xea\x95\x7c\xf7\xe1\x7f\x89\x9b\x0c\xc6\xca\x46\x5e\x74\xff\x48\xe7\x9e\x23\xe5\x38\xab\x63\x63\xc1\x6a\x57\xe5\x02\x5d\xb2\x63\xb1\xf9\x02\x90\xd5\xa8\x8b\x7e\x01\x90\x59\x72\xaf\x10\xd0\xe6\x82\x8c\x77\x02\x62\xd7\x62\x2e\xa8\xd1\x15\xaa\xd6\x55\x02\xea\x80\x65\x2e\xea\x18\x5b\x4e\x95\xe2\xe8\x83\xae\x50\x56\xde\x57\x16\x75\x4b\x2c\x8d\x6f\x54\xad\x79\x17\x74\x52\x05\xdd\xd6\x7f\xd8\x04\x5d\x45\x0e\x95\xf1\x8e\xbd\x45\xa5\x99\x31\xb2\x2a\xf5\xbe\xc7\x90\xc3\xb1\x6a\xc2\x67\x13\xa8\x8d\x63\x00\x07\x72\x85\x3f\xc8\xe7\x67\x74\x7b\xc8\xe1\x7d\xcc\x02\xe8\x96\x7e\xf2\x1c\x53\x78\x7f\x97\xd3\xfa\x78\xfc\x66\x51\x7d\xf4\x21\xa6\x20\xc6\x72\x1f\x1c\x8f\xe2\x54\x37\x96\x7e\xc7\xc0\xe4\xdd\x70\xc0\x39\x5c\x9c\x51\xe8\xa8\xbf\x6f\xe9\x29\xd8\xa1\xe7\x1c\x5e\xf7\x2c\xcf\xfa\x98\x5a\xce\x64\x0c\x32\xff\x8c\xdd\x38\xf3\x1c\x2d\x3a\x76\xc1\x3e\x06\x2c\xe9\x2d\x05\xa1\x16\xc3\x8e\x9a\xfd\xea\x0b\x4c\x41\x18\x4b\x62\xac\x1c\xd7\xa3\x5e\xea\x2c\x58\xa6\xe6\x7b\xcc\x5e\x7c\xd1\xcd\x82\xc6\xce\xe2\xb8\x96\x8d\x26\xf7\x83\x77\x11\x5d\x3c\x89\x59\x10\xb7\x56\x77\x29\xdc\x38\xef\xf0\x66\x3d\xa5\x7d\xab\x0d\xc5\x2e\x85\xd5\x9c\x89\x41\x3b\xa6\x38\x50\x9d\xaa\x20\xef\x56\x0c\x96\x1c\xea\x30\xb6\x1d\xaf\x80\x24\xd7\xfe\xf0\x37\x68\x2f\xd6\x9b\xd7\x6b\xb8\xdb\x4f\xc0\x65\x6a\x22\x39\x46\x05\xed\x81\x8a\x5c\x58\xaf\x0b\x72\x95\x98\x1e\xd3\x58\x30\x56\x33\xe7\xa2\xd5\x15\x26\x73\x03\x0c\xdb\xf3\x49\x59\x68\xc8\x25\x35\x52\x55\xc7\x14\x6e\x57\xab\x7d\x3d\x8f\x74\xa0\x22\xd6\x43\xee\xeb\xf5\x25\x9f\xd2\xe2\xdb\x9c\xd4\x96\x2a\x97\x50\xc4\x86\x53\x30\xe8\x22\x86\xb9\x54\x7a\x17\x93\x52\x37\x64\xbb\x14\x58\x3b\x4e\x18\x03\x95\x73\x79\xbb\xe3\x48\x65\x97\x98\x51\xbb\xcb\xdd\x27\x2a\xfd\xbd\xb6\xda\xcd\x6c\x2e\x19\x4c\x38\x4c\x7f\x62\x0a\x77\xd8\xac\x4f\xf9\x46\x87\x8a\x5c\x12\x7d\x9b\x42\xf2\xed\xb2\x62\xbc\xf5\x21\x85\xaf\x1e\xee\xfb\xdf\x39\xbf\xc0\xfc\x65\xd4\x4b\x4a\x39\x2b\xaa\xfa\x29\x4e\xfa\xaa\x82\xf6\xd3\xab\x5b\x2c\xe7\xeb\x98\x38\x89\x79\xe8\xc5\x1b\x11\x9b\xe5\x86\xb3\x05\x0d\xac\xb8\x46\x8c\x97\xbe\x63\x0a\xb7\x65\x69\xac\xdf\x15\xa5\xd5\x01\x07\xd7\xd1\x5b\xfd\xa6\x2c\xbd\xb0\x1a\xe8\xeb\x03\xb2\x6f\x50\xdd\xcb\xef\xe4\x4a\x19\xfe\x98\x96\x0d\x39\x69\x98\x85\xfa\x17\xb0\x9f\xb2\xbb\xfe\x0f\x3f\x38\xde\xc9\x16\x54\xcf\x7d\x00\x07\x53\xeb\xc0\x18\x73\xf1\xf4\xdb\x8f\xc9\x83\xf8\x68\x83\xc0\xc1\xfc\xf7\xe0\x7b\x74\x85\x0f\x72\x7b\x8d\xbe\x59\xda\xc9\xff\x3c\xc5\x20\xc1\x3f\xcd\x90\xa9\xd1\xc9\x32\xd5\x7f\xb8\x36\x5f\xfe\x0a\x00\x00\xff\xff\xdf\x70\xc8\xd7\xc0\x06\x00\x00")

func assetsV10AlphaConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10AlphaConsoleHtml,
		"assets/v1.0-alpha/console.html",
	)
}

func assetsV10AlphaConsoleHtml() (*asset, error) {
	bytes, err := assetsV10AlphaConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0-alpha/console.html", size: 1728, mode: os.FileMode(436), modTime: time.Unix(1530689993, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsV10ConsoleHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x5f\x6b\xfb\x36\x14\x7d\xff\x7d\x8a\x8b\xc6\xe8\xcb\x2c\xa5\x5d\x61\xc5\xb1\x03\x63\x30\x06\xdb\xa0\x0f\xeb\x5e\x8b\x2a\x5f\xdb\x37\x95\x25\x4f\x57\x49\xea\x95\x7c\xf7\xe1\x7f\x89\x9b\x0c\xc6\xca\x46\x5e\x74\xff\x48\xe7\x9e\x23\xe5\x38\xab\x63\x63\xc1\x6a\x57\xe5\x02\x5d\xb2\x63\xb1\xf9\x02\x90\xd5\xa8\x8b\x7e\x01\x90\x59\x72\xaf\x10\xd0\xe6\x82\x8c\x77\x02\x62\xd7\x62\x2e\xa8\xd1\x15\xaa\xd6\x55\x02\xea\x80\x65\x2e\xea\x18\x5b\x4e\x95\xe2\xe8\x83\xae\x50\x56\xde\x57\x16\x75\x4b\x2c\x8d\x6f\x54\xad\x79\x17\x74\x52\x05\xdd\xd6\x7f\xd8\x04\x5d\x45\x0e\x95\xf1\x8e\xbd\x45\xa5\x99\x31\xb2\x2a\xf5\xbe\xc7\x90\xc3\xb1\x6a\xc2\x67\x13\xa8\x8d\x63\x00\x07\x72\x85\x3f\xc8\xe7\x67\x74\x7b\xc8\xe1\x7d\xcc\x02\xe8\x96\x7e\xf2\x1c\x53\x78\x7f\x97\xd3\xfa\x78\xfc\x66\x51\x7d\xf4\x21\xa6\x20\xc6\x72\x1f\x1c\x8f\xe2\x54\x37\x96\x7e\xc7\xc0\xe4\xdd\x70\xc0\x39\x5c\x9c\x51\xe8\xa8\xbf\x6f\xe9\x29\xd8\xa1\xe7\x1c\x5e\xf7\x2c\xcf\xfa\x98\x5a\xce\x64\x0c\x32\xff\x8c\xdd\x38\xf3\x1c\x2d\x3a\x76\xc1\x3e\x06\x2c\xe9\x2d\x05\xa1\x16\xc3\x8e\x9a\xfd\xea\x0b\x4c\x41\x18\x4b\x62\xac\x1c\xd7\xa3\x5e\xea\x2c\x58\xa6\xe6\x7b\xcc\x5e\x7c\xd1\xcd\x82\xc6\xce\xe2\xb8\x96\x8d\x26\xf7\x83\x77\x11\x5d\x3c\x89\x59\x10\xb7\x56\x77\x29\xdc\x38\xef\xf0\x66\x3d\xa5\x7d\xab\x0d\xc5\x2e\x85\xd5\x9c\x89\x41\x3b\xa6\x38\x50\x9d\xaa\x20\xef\x56\x0c\x96\x1c\xea\x30\xb6\x1d\xaf\x80\x24\xd7\xfe\xf0\x37\x68\x2f\xd6\x9b\xd7\x6b\xb8\xdb\x4f\xc0\x65\x6a\x22\x39\x46\x05\xed\x81\x8a\x5c\x58\xaf\x0b\x72\x95\x98\x1e\xd3\x58\x30\x56\x33\xe7\xa2\xd5\x15\x26\x73\x03\x0c\xdb\xf3\x49\x59\x68\xc8\x25\x35\x52\x55\xc7\x14\x6e\x57\xab\x7d\x3d\x8f\x74\xa0\x22\xd6\x43\xee\xeb\xf5\x25\x9f\xd2\xe2\xdb\x9c\xd4\x96\x2a\x97\x50\xc4\x86\x53\x30\xe8\x22\x86\xb9\x54\x7a\x17\x93\x52\x37\x64\xbb\x14\x58\x3b\x4e\x18\x03\x95\x73\x79\xbb\xe3\x48\x65\x97\x98\x51\xbb\xcb\xdd\x27\x2a\xfd\xbd\xb6\xda\xcd\x6c\x2e\x19\x4c\x38\x4c\x7f\x62\x0a\x77\xd8\xac\x4f\xf9\x46\x87\x8a\x5c\x12\x7d\x9b\x42\xf2\xed\xb2\x62\xbc\xf5\x21\x85\xaf\x1e\xee\xfb\xdf\x39\xbf\xc0\xfc\x65\xd4\x4b\x4a\x39\x2b\xaa\xfa\x29\x4e\xfa\xaa\x82\xf6\xd3\xab\x5b\x2c\xe7\xeb\x98\x38\x89\x79\xe8\xc5\x1b\x11\x9b\xe5\x86\xb3\x05\x0d\xac\xb8\x46\x8c\x97\xbe\x63\x0a\xb7\x65\x69\xac\xdf\x15\xa5\xd5\x01\x07\xd7\xd1\x5b\xfd\xa6\x2c\xbd\xb0\x1a\xe8\xeb\x03\xb2\x6f\x50\xdd\xcb\xef\xe4\x4a\x19\xfe\x98\x96\x0d\x39\x69\x98\x85\xfa\x17\xb0\x9f\xb2\xbb\xfe\x0f\x3f\x38\xde\xc9\x16\x54\xcf\x7d\x00\x07\x53\xeb\xc0\x18\x73\xf1\xf4\xdb\x8f\xc9\x83\xf8\x68\x83\xc0\xc1\xfc\xf7\xe0\x7b\x74\x85\x0f\x72\x7b\x8d\xbe\x59\xda\xc9\xff\x3c\xc5\x20\xc1\x3f\xcd\x90\xa9\xd1\xc9\x32\xd5\x7f\xb8\x36\x5f\xfe\x0a\x00\x00\xff\xff\xdf\x70\xc8\xd7\xc0\x06\x00\x00")

func assetsV10ConsoleHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsV10ConsoleHtml,
		"assets/v1.0/console.html",
	)
}

func assetsV10ConsoleHtml() (*asset, error) {
	bytes, err := assetsV10ConsoleHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/v1.0/console.html", size: 1728, mode: os.FileMode(436), modTime: time.Unix(1530690011, 0)}
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
	"assets/unversioned/console.html": assetsUnversionedConsoleHtml,
	"assets/v1.0-alpha/console.html": assetsV10AlphaConsoleHtml,
	"assets/v1.0/console.html": assetsV10ConsoleHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"unversioned": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsUnversionedConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10ConsoleHtml, map[string]*bintree{}},
		}},
		"v1.0-alpha": &bintree{nil, map[string]*bintree{
			"console.html": &bintree{assetsV10AlphaConsoleHtml, map[string]*bintree{}},
		}},
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

