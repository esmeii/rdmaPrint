// Code generated by go-bindata. DO NOT EDIT.
// sources:
// vloadstorelocal.hsaco

package vloadstorelocal


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataVloadstorelocalhsaco = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5f\x6f\xdb\x54\x14\x3f\xbe\x71\xfe\xd4\x49\xff\x6c\x32\x25\x6c" +
	"\xd1\x30\x55\xa5\x54\x13\x89\x52\xaf\x8c\xa8\x2f\xf4\xdf\xd6\x4d\x4b\xbb\x6e\x65\x1b\x83\x4d\x9b\x17\x3b\xa9\x99" +
	"\x63\x07\xc7\xa9\x9a\x09\xba\x82\xa6\x69\x42\x95\x98\x10\x0f\x7b\x2b\x0f\xf0\xc6\x03\x9f\x60\x29\x5f\x01\x21\x94" +
	"\x87\x3d\xf0\xb2\x8f\x00\x4f\x88\xa0\x6b\xdf\x9b\xd8\x66\xee\x0a\x74\x1a\x12\xfe\x49\xc9\xf1\x3d\xbf\x7b\xee\x39" +
	"\xe7\xde\xd3\xa4\x37\xf7\xde\x3d\x55\x3a\x8d\x18\x66\x06\x08\x22\xf0\x0b\x30\xf8\x61\xc4\x69\x53\xe2\xc1\x1b\x8e" +
	"\x3c\x6e\xeb\x8a\x90\x80\x19\x48\x01\x07\x31\x00\x60\x5d\xfd\xfc\x72\x97\xf1\xca\x04\xd1\x33\xc4\x2e\x08\xdf\x26" +
	"\xbd\x92\xc6\x83\xed\xa2\xae\xb6\x5f\x76\x90\x57\xba\xed\x70\xac\x20\x10\xbd\x4f\xd6\xc1\x2b\xa9\x1d\xfa\x9b\x76" +
	"\x34\xbf\x8b\x4f\x2d\x99\xdd\x87\x9d\x3b\x3e\x8c\x0b\x4f\x2d\x39\xb6\xc7\xbc\x04\x81\xa5\xf3\x49\xf3\xf6\xc9\x9b" +
	"\x03\x5e\xc9\xba\xec\x12\xc4\xff\xec\xd2\x82\xdd\x9d\xae\xcd\x51\xbb\x1e\x1c\x3d\x0b\xf1\x5e\x6e\x54\x37\xbb\xb4" +
	"\xb0\xb8\x72\xc9\xe9\x8b\xcb\x23\x49\xf4\x52\x4d\xae\x96\xf5\x9c\x54\x93\xf1\x6b\xad\x21\x61\x51\x55\xef\x94\xb5" +
	"\x5c\xb5\xb2\x51\x2c\x9c\x20\xe3\xff\x9c\x00\xe0\x88\x4d\x2e\x97\xe3\x2e\x2b\x66\x43\x35\xf4\x69\x81\xe2\x03\x61" +
	"\xf2\x4d\xa1\x20\x5c\xe7\xce\x29\xa6\xae\x68\x8d\x3e\xc3\x09\x42\x4e\x58\x96\x6a\x4a\x5f\x25\x08\xc2\xba\xd6\xd0" +
	"\x38\xfc\xb0\xda\xaa\xdd\x32\x34\x17\x9f\xc5\xd4\xcc\x6d\x39\x6b\xd3\x25\x49\xaf\x36\xa5\x6a\xdf\xf8\x7c\x5d\xd1" +
	"\xe7\x4b\xc2\xbc\x87\xed\x85\x63\x87\x21\x0a\xd7\x6d\x76\xd6\xac\x36\x3c\x4e\x39\x47\x3c\x23\x1c\x89\xa3\x4f\xef" +
	"\xb6\xea\x8a\x87\xce\xaa\xba\x75\x3c\xdb\xe3\x57\xd5\x3b\x5e\xd3\x62\x8f\x9a\xd5\xd4\xaa\x3e\xfd\x4c\xea\xb2\xa4" +
	"\x35\x95\x73\xaa\x2e\x53\x7a\x51\x33\x6e\x49\xda\x5c\xb3\x52\x51\x4c\x6f\x2f\x1c\x00\xed\x75\xf6\x84\xd8\x1f\x5d" +
	"\x96\xcd\xd5\xba\x54\x56\x2e\x34\x25\x6d\xba\x37\x44\x9f\x2f\x97\x29\xe3\x60\x41\xa9\x48\x4d\xcd\x0a\x4e\xba\xfc" +
	"\x7f\x4c\x5a\xba\x61\xd5\xea\xff\x3c\xf1\xa9\xe0\xc4\xa7\x82\x13\x5f\x68\xe9\x52\x4d\x2d\xaf\xae\x49\xa6\x22\xaf" +
	"\x18\xaa\x6e\xed\x77\x02\x9c\xce\x4a\xdf\xd9\x54\xf0\xdc\x94\x8c\xf2\xbf\xab\x07\xa3\xa9\x5b\xc1\x53\xd3\x54\x5d" +
	"\xec\x01\x4d\xcc\x5c\xcb\x56\x05\xcf\xc5\x25\x77\x31\x3c\x2f\xa3\x03\x2a\xd3\x33\xaa\x2c\x2b\xba\x53\x69\xe7\x2b" +
	"\x95\x86\x62\xbd\xb7\xc7\x62\x9d\x9c\x7a\xf1\xfe\xaf\xbe\x64\xff\xef\xbf\x04\xff\xcb\x86\xbe\x47\x5d\x9c\x2d\xee" +
	"\xf3\x33\x22\x8c\x6a\xbf\x51\xcd\x1b\xb2\xb2\x62\x1a\xf5\xde\x77\x26\xb1\xc2\xdf\xe8\x92\x59\x5d\x55\xaa\x35\x45" +
	"\xb7\x9c\xb8\xdf\xa6\x7f\x95\x8b\xa6\xd1\xac\x13\xea\xb4\xba\xa1\xc8\x0e\x5f\x20\xf4\x8a\xa9\xae\x4b\x96\x12\xdc" +
	"\xc1\x3b\x38\xc9\x9c\xc6\x7b\x45\x5a\x57\x2a\xa6\x41\x9d\x0a\x42\xaf\xd6\x96\x9b\xb5\xd5\xc5\x95\x8b\xfd\xaf\xf7" +
	"\xc9\xc9\x3e\x73\xd9\xc3\x88\x94\x59\x92\x36\x4e\x6b\x92\x75\xc5\x30\x6f\x3b\x51\xdb\x83\x8a\x6f\x9d\xe4\xf2\xf9" +
	"\x3c\x17\xfc\xff\x1a\xe3\x7a\x1d\x09\xe0\x42\x84\x08\x11\xe2\x45\x82\x21\x5b\x44\xc6\xde\xdd\x45\x9e\xfb\xc1\x73" +
	"\x0a\xbe\x83\x87\xf6\x5e\xcf\xfb\xe1\x76\xc6\xf5\x9c\x84\x57\x3c\x1c\xcb\xb2\xb1\x6e\xb7\xdb\x3d\xc8\xb8\x0f\x0a" +
	"\x08\xd0\x2e\xde\x8b\x46\x98\xd8\x2e\xde\x8a\xcf\x03\xda\x4d\x03\xc0\x5d\xd8\x6e\x43\x17\xee\xe3\xa8\x13\x90\xf8" +
	"\x32\x01\x48\x64\x10\x12\xef\x21\x34\x1e\x45\x8f\x3e\xf9\x50\xd8\x7a\xac\xc0\x83\x76\x04\xb8\x5d\x3c\xce\x0e\x62" +
	"\xc7\x0e\xc3\x17\x3f\x7e\x86\x10\xdc\x07\x18\xff\x01\x7e\x7f\x8c\xc7\x60\x10\xbb\x09\xe3\xbc\x88\xb2\xa3\xc5\xef" +
	"\xf9\x94\xb8\x35\x3a\x54\xfc\x9a\x67\xc5\xad\xd1\x58\xf1\x21\x9f\x11\xb7\x46\x8f\x15\x01\xae\x3d\xe1\xec\xad\xf0" +
	"\xb5\x27\x43\x00\x43\x58\x22\x7b\x59\xae\x3d\x89\x01\xc4\x58\x98\x10\xef\x21\x98\x68\x44\xb6\xdb\x28\x92\xe9\xf0" +
	"\xa9\x41\x00\x26\xd3\xe1\xb9\x24\x3c\x40\xfc\xc4\x36\x4a\x4f\xec\x00\x1a\x33\xd9\xed\x76\x2c\x9e\xe9\xf0\x23\x87" +
	"\x80\x8d\x66\x3a\xfc\xd0\x30\x00\x8e\x09\x10\x44\xd0\xb1\x4d\x34\x9e\x11\x77\xf8\xd1\xb1\xe1\xec\xb1\x22\x02\x41" +
	"\x8c\xa0\xb1\x4d\x0e\xf3\x3c\x82\x9d\xf4\x91\xb1\x43\x91\xb1\x62\x0a\xb7\xd3\xb8\x3f\x6c\x22\x9e\x17\xe1\xd5\xd1" +
	"\x22\x4a\xa7\xc5\x8f\xa2\xdb\x6d\x2e\x99\xe9\xf0\x6c\x14\x12\x03\x99\x0e\x8f\x22\x50\x8f\x6d\xb7\x87\x86\x33\x1d" +
	"\x3e\x31\x00\xa9\xc1\x4c\x87\x8f\xc5\x9d\x79\x83\xcf\xdb\x70\xf4\x48\x11\x18\xbd\xc3\xe3\x39\x8e\xd8\x12\x39\xdc" +
	"\xc7\x38\x47\x60\xa3\x0e\x17\x8b\x7b\x39\xbc\x06\x89\x01\x87\xe3\x92\x5e\x0e\x57\x5c\x6a\xd0\xe1\x86\x86\xbd\x5c" +
	"\xca\xfe\x4d\x85\xf8\x86\x4f\xdb\x2f\xbb\xae\x42\x84\x08\x11\x22\x44\x88\x10\x21\x42\x84\x08\x11\xe2\xbf\x06\x7a" +
	"\xd6\x7c\x93\x9c\xb3\xd3\xe3\xf6\x34\x91\x51\x22\xbf\x21\x04\xdd\xf5\xd3\x9f\x09\x7e\xfd\xa3\x6b\x60\xb9\x41\x78" +
	"\x7a\xae\xfc\x55\xf2\xd9\xfe\x4a\xaa\x7e\x5b\x31\xa7\x85\x52\x69\x41\x98\xca\x17\xa0\xac\x49\x7a\x55\x58\x77\x8e" +
	"\x5a\x85\x62\xbe\x20\xec\x1d\x2f\x43\xa2\x9e\x39\xe4\xd5\xc7\x89\xfe\xa7\xc3\x5e\xfd\x20\x8e\x19\xc5\xfe\x72\x3f" +
	"\xc0\xee\x86\xe2\xfd\x73\x79\x8a\xbc\x6e\x58\x0a\xe4\xe5\x96\xde\x68\xd5\x20\x5f\xd5\x9b\xf9\x35\xa9\xb1\x06\xe4" +
	"\x1d\xeb\x2d\x13\xf2\x96\xb2\x61\xd9\x2d\xa9\xa6\x96\x21\x5f\x36\x6a\x35\x45\xb7\x20\xdf\x68\xd5\x2c\xe9\x16\xe4" +
	"\x1b\x6b\x0d\xcb\x74\x9e\x1c\x09\x73\x73\x85\x1b\x93\xf6\xbb\x08\xeb\x5a\x43\x83\x1b\x0b\x57\x97\x67\x97\xce\xce" +
	"\xef\x9d\xf0\x3e\xc1\x90\x39\xa0\xc7\xfe\x41\xe7\xff\x14\xfe\xbb\x17\x71\xb2\xf6\xd4\x8c\xd6\x03\x95\xee\x7a\x60" +
	"\x5c\xf7\x1c\xa8\x7e\x18\x00\x7e\xeb\x76\x0d\x6a\x4f\xeb\x81\xca\x8c\x2f\xac\x84\xcf\xff\x6b\x64\x6c\xca\xd3\xfa" +
	"\xa1\x72\xc4\x67\xcf\xfa\xe4\xeb\xe4\x5e\x02\xe5\x69\xbd\x52\xe9\xff\x59\xcb\xdf\xce\xba\xef\x88\x40\xf0\xbd\x92" +
	"\xa0\x01\x72\xc4\x36\x42\x15\x01\xf7\x3d\xa2\xbe\xfc\xa9\x9b\x93\x64\xc8\x82\xcf\x4d\x9d\xd8\x8f\x07\xb8\xa7\xf2" +
	"\x1d\xf7\xda\xbb\xf0\x88\xd8\x6f\x90\x76\x92\xc4\xe0\x5f\xbf\x45\x77\xec\x2e\x8c\x90\x7b\x3f\x57\x03\xfc\x53\x5c" +
	"\x08\xb0\x57\x89\xfd\xd1\xe7\xd8\xff\x19\x00\x00\xff\xff\x93\x1a\x83\x4c\x88\x24\x00\x00")

func bindataVloadstorelocalhsacoBytes() ([]byte, error) {
	return bindataRead(
		_bindataVloadstorelocalhsaco,
		"vloadstorelocal.hsaco",
	)
}



func bindataVloadstorelocalhsaco() (*asset, error) {
	bytes, err := bindataVloadstorelocalhsacoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "vloadstorelocal.hsaco",
		size: 9352,
		md5checksum: "",
		mode: os.FileMode(509),
		modTime: time.Unix(1551896327, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"vloadstorelocal.hsaco": bindataVloadstorelocalhsaco,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"vloadstorelocal.hsaco": {Func: bindataVloadstorelocalhsaco, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
