// Code generated by go-bindata. DO NOT EDIT.
// sources:
// basic-float64.md (359B)
// calc-range-condition.md (211B)
// dry.md (133B)
// math-pow.md (106B)
// square-sum-loop.md (443B)
// sum-square-loop.md (195B)

package tpl

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _basicFloat64Md = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x31\x4f\xc3\x40\x0c\x85\x77\x7e\xc5\x1b\x41\x2a\x11\x8a\x10\x2b\x1b\x88\x1d\x89\xf5\xcc\xd5\x49\x2c\xae\x76\x74\xe7\xa4\xcd\xbf\x47\x4e\xe9\xc2\x6a\xf9\xfb\xde\xb3\x1f\xf1\x69\xc8\x95\xc9\x19\x84\xa1\x18\xb9\xe8\x88\xd9\x44\x1d\x2b\x95\x85\xb1\xd9\x82\x4c\x8a\xa5\x31\x7c\x62\x14\x71\xae\x54\x90\xfa\xee\x29\x1d\x20\xda\x9c\xe9\x08\x1b\x40\xf0\x6d\x66\x64\xd3\x95\x6b\x13\x53\xa4\xdd\xf8\xf2\x7c\xdf\x3f\xa4\x0e\x1f\x03\xc4\x21\x0d\xb9\x30\x55\xb8\xe1\xdd\xe0\x13\xf9\x9e\x41\x95\xe1\x75\x8b\x78\xff\xdf\x29\x82\x15\xa9\x4f\x41\xb3\xda\x32\x4e\x1d\xbe\x62\x76\x34\x6e\xa1\xf9\x51\x3b\xef\x9a\x33\xa9\xdf\xb0\x57\xbc\x59\x05\x5f\xe8\x34\x17\x3e\x40\x86\x7d\x23\x0e\x11\x07\x35\x10\x66\xaa\x74\x62\xe7\x0a\xd1\xa0\x16\xcd\x1e\xc5\x33\x95\x72\xad\xc6\x97\x99\xb3\xc7\xee\xf5\x98\x04\xab\x37\x13\xb5\x26\xa3\x86\xcc\x0d\x84\x95\xaa\xd0\x77\xe1\xf8\xc5\xfe\x89\x3f\xa2\xbb\xfb\x0d\x00\x00\xff\xff\xc3\xfb\x4e\x14\x67\x01\x00\x00")

func basicFloat64MdBytes() ([]byte, error) {
	return bindataRead(
		_basicFloat64Md,
		"basic-float64.md",
	)
}

func basicFloat64Md() (*asset, error) {
	bytes, err := basicFloat64MdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "basic-float64.md", size: 359, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x78, 0x2f, 0x31, 0x10, 0x9b, 0x27, 0x94, 0x16, 0x20, 0x96, 0xa, 0xf5, 0x81, 0xf9, 0x85, 0xc7, 0x29, 0xa9, 0xfd, 0x49, 0x4e, 0x42, 0xe2, 0xdf, 0xf8, 0x78, 0x71, 0xe, 0xed, 0x79, 0x18, 0xd5}}
	return a, nil
}

var _calcRangeConditionMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\xc1\x4d\xc4\x40\x0c\x85\xe1\x3b\x55\xbc\x0b\xda\x0b\x6c\x01\x54\x00\x5d\xc4\x8c\x1f\x8c\x45\xd6\x8e\x62\x6f\x50\xba\x47\x33\xec\xf1\x59\xfa\x7e\xbf\xe2\x9d\x3b\x2f\x89\x70\xe2\x57\x4e\x54\xc0\x6e\xdb\x1e\x07\x51\x9d\xc8\x8d\x54\xc4\xd7\x1c\xcb\x2e\xfe\xcd\x05\x6b\xc4\xf6\x36\x2f\x2d\x5c\xad\x2c\x1c\x4e\x6a\x0e\xfd\x49\xb4\xce\xf6\x33\x98\x83\x07\xf7\x73\x82\x17\x88\x2b\xac\x90\x3d\xee\xab\xfa\xa5\x86\x2e\x31\x47\x93\xb5\xdd\x57\x19\x9d\x44\x75\x29\x34\xf1\x11\x92\x23\x4c\xa9\x57\x7c\x38\xaa\x5b\xa2\x49\x12\xcb\x73\x2e\x98\xe3\x9f\x51\x1f\x6f\xca\x6e\xbc\x3e\xfd\x05\x00\x00\xff\xff\x7e\xca\x2e\x24\xd3\x00\x00\x00")

func calcRangeConditionMdBytes() ([]byte, error) {
	return bindataRead(
		_calcRangeConditionMd,
		"calc-range-condition.md",
	)
}

func calcRangeConditionMd() (*asset, error) {
	bytes, err := calcRangeConditionMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "calc-range-condition.md", size: 211, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7, 0x99, 0x99, 0x7a, 0xa8, 0x77, 0xc5, 0x2a, 0x33, 0x4e, 0x27, 0x7c, 0x91, 0x9, 0x59, 0x1e, 0x0, 0xad, 0x8b, 0x70, 0x2e, 0x1, 0xe7, 0xa1, 0x3b, 0x68, 0xcd, 0xa2, 0xfa, 0x82, 0x9e, 0x97}}
	return a, nil
}

var _dryMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcb\x31\xae\xc2\x30\x0c\x06\xe0\xfd\x9d\xe2\x3f\xc0\x83\x53\xb0\x77\xe8\xc4\x16\x93\xd8\xd4\x52\x6a\x97\xc4\x46\xea\xed\x11\x62\xfc\x86\xef\x82\xbb\x27\xaa\x67\x6f\x98\xba\x1f\x5d\xe5\xc4\xe9\x39\x50\x6e\x2a\xc2\x83\xad\x72\x81\xa4\xd5\x50\x37\x3c\x4e\x6c\xf4\x56\x7b\x42\x03\x95\x7a\x47\x59\x73\x5f\x64\x7d\x25\x0d\x9e\x05\x64\x0d\xe5\xa7\x45\xd6\xdc\x67\xf9\x87\xda\x0c\xa6\x06\x17\x0c\x3e\x98\xe2\xfb\x63\xa3\x40\xf5\xc6\xd7\xbf\x4f\x00\x00\x00\xff\xff\xec\x10\x8f\x5d\x85\x00\x00\x00")

func dryMdBytes() ([]byte, error) {
	return bindataRead(
		_dryMd,
		"dry.md",
	)
}

func dryMd() (*asset, error) {
	bytes, err := dryMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dry.md", size: 133, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfb, 0x74, 0x95, 0xb7, 0x27, 0x7f, 0xa3, 0xca, 0x29, 0x10, 0xb8, 0xb, 0xcf, 0xef, 0x98, 0x5a, 0x89, 0xed, 0xe, 0x71, 0x3c, 0xb5, 0xbf, 0x74, 0x6f, 0xd8, 0x29, 0x68, 0x44, 0x85, 0xbd, 0xbb}}
	return a, nil
}

var _mathPowMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\xca\xbb\x0d\xc2\x40\x0c\x80\xe1\x9e\x29\xfe\x1a\x41\xf6\xa0\x63\x84\x33\x91\xe1\x82\x12\x1f\xf8\xa1\xc0\xf6\x48\xf4\xdf\x99\xcb\x9d\xef\x28\x9e\x15\xc9\x2e\x96\xe4\x20\xde\x25\xae\x08\x56\xdb\x4d\xfd\xf4\x17\xb3\x18\x15\x4a\x76\x45\x3f\x2f\xd7\x88\x65\x18\xcd\x38\x62\x0d\x97\xec\xea\x64\x17\x63\x96\x75\x5d\xec\x41\xdb\x24\xfb\x74\x1d\x7b\x9b\x0e\xbf\x00\x00\x00\xff\xff\xc1\xc4\x7a\x68\x6a\x00\x00\x00")

func mathPowMdBytes() ([]byte, error) {
	return bindataRead(
		_mathPowMd,
		"math-pow.md",
	)
}

func mathPowMd() (*asset, error) {
	bytes, err := mathPowMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "math-pow.md", size: 106, mode: os.FileMode(420), modTime: time.Unix(1542537851, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd2, 0x1e, 0x7f, 0x7a, 0xce, 0x7, 0x5a, 0xc1, 0x5b, 0xe1, 0xe3, 0x86, 0x1d, 0x56, 0x79, 0xc9, 0x78, 0xea, 0x60, 0x73, 0xe0, 0x38, 0x9f, 0xe0, 0x3d, 0xe6, 0x78, 0x3d, 0x92, 0x48, 0xb1, 0x42}}
	return a, nil
}

var _squareSumLoopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x90\x31\x8f\xd5\x30\x10\x84\x7b\x7e\xc5\x50\x01\xd2\xdd\x29\xaf\x42\xa2\x41\xe8\x24\x10\x15\xc5\x41\xff\x36\xce\xe6\xd9\x7a\xb6\x37\x6f\x77\x4d\x94\x7f\x7f\x72\x92\xce\x92\x67\xbe\x99\xd9\x67\xfc\x40\x16\x59\x40\xcb\xa2\x42\x21\xc2\x05\x81\x72\x68\x99\x3c\xd5\x1b\xae\x6f\x8f\x46\xca\x7f\xe6\xb7\x56\xae\x58\x45\xef\xf6\x84\xb1\x39\x3c\xb2\xf2\x27\x03\xc1\xa2\xa8\x3f\x87\xe6\x2f\xf8\x1b\x19\x37\x65\x72\x14\xf2\xc8\x85\x3c\x85\x44\x15\xaf\xa4\x19\x3f\x35\xf1\xa4\x29\x44\xfc\xa2\x66\x86\x99\x02\x4f\x9d\x03\xa3\xc2\x58\x54\xc6\xcc\x05\x64\xd8\xa4\x61\x8d\x5c\x11\x19\x2b\x19\xc8\x61\x21\x8a\x64\xa4\xba\x1b\x2e\x5f\x87\xc1\xbe\xf5\xae\x73\xaa\x27\xa3\x15\xc8\xbc\x3f\x6b\x2b\x23\xab\x61\x56\x29\xb8\x74\xd5\x65\x18\x8e\x72\x8f\xc6\xe6\x49\xea\x81\x35\x4b\xb7\xca\x53\x8f\x1c\x9b\x6d\x7d\x1d\xc6\x6d\x67\x38\x53\x88\xac\xc7\xd6\xb3\xaf\xb4\x33\x8b\xaa\xad\xac\xd0\xbe\x51\xf1\x68\x29\xdc\xf3\xd6\x9d\x53\xb2\x20\xff\x59\xfb\xe5\x08\x0b\xb9\xb3\xd6\x17\xbc\x52\xdd\x37\xed\x65\x93\xc3\x45\xbe\xe3\xf3\xbf\x9a\xd3\x9d\x0f\xf8\xd3\xfe\x1f\xa8\xa2\x19\xef\x19\xbf\x6b\xf7\x72\x17\xc3\x98\x34\x44\xcc\xa2\x67\xb4\x7d\xfc\xf2\xe1\x3d\x00\x00\xff\xff\xbe\x3b\xc7\xb9\xbb\x01\x00\x00")

func squareSumLoopMdBytes() ([]byte, error) {
	return bindataRead(
		_squareSumLoopMd,
		"square-sum-loop.md",
	)
}

func squareSumLoopMd() (*asset, error) {
	bytes, err := squareSumLoopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "square-sum-loop.md", size: 443, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4a, 0x17, 0xf6, 0xbb, 0x73, 0x6, 0x8d, 0xba, 0x6a, 0x1a, 0xd8, 0x2e, 0xd6, 0x51, 0x59, 0xa5, 0x8c, 0xe0, 0x7e, 0xb, 0x17, 0x91, 0x86, 0x2b, 0xb6, 0xb7, 0x43, 0xe3, 0x2c, 0x5, 0x3c, 0x22}}
	return a, nil
}

var _sumSquareLoopMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xcd\x3d\x6e\xc3\x30\x10\x44\xe1\x3e\xa7\x98\x74\x09\x90\xe4\x0a\x29\x52\xa4\x74\xa1\x0b\x68\x65\xf1\x0f\x26\x39\xf2\x72\x09\x42\x3e\xbd\x61\xab\x74\xfd\xf0\xf0\x7d\xe3\x9f\x5c\xb1\x29\x83\x4a\x29\x4e\x1b\x44\x1d\x16\x5a\x44\x2b\xa2\x06\xa9\x2b\xb2\xdc\xf6\x1f\xfc\x49\xc5\xce\x0e\x9f\xea\x0a\x39\xf2\xd7\xb3\x61\xc8\x0e\x7a\x9c\x59\xb6\x6e\xa9\x06\xcc\x53\x2f\x27\x3f\x5d\xbb\xa8\x6b\x33\x46\xb2\xc8\x6e\x10\x64\x72\xfb\xc5\xc7\xab\x9a\x1b\x71\xa9\x1c\x88\x1c\x30\x1e\x8a\x45\x07\x4d\x21\x1a\x24\x07\x6a\xb2\x58\xe0\xa9\x90\xc7\xbb\x64\x57\xde\x3f\xdf\xee\x01\x00\x00\xff\xff\x90\x56\xe0\xb8\xc3\x00\x00\x00")

func sumSquareLoopMdBytes() ([]byte, error) {
	return bindataRead(
		_sumSquareLoopMd,
		"sum-square-loop.md",
	)
}

func sumSquareLoopMd() (*asset, error) {
	bytes, err := sumSquareLoopMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sum-square-loop.md", size: 195, mode: os.FileMode(420), modTime: time.Unix(1542628451, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbb, 0x6e, 0xea, 0x18, 0x19, 0xf0, 0x18, 0xd5, 0x1, 0xef, 0x50, 0x4c, 0xc3, 0x14, 0x6f, 0xa1, 0x67, 0x3f, 0xcb, 0x87, 0xaf, 0x9d, 0xde, 0xf3, 0xe6, 0x23, 0x5b, 0x5, 0x69, 0x98, 0x7a, 0xf1}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"basic-float64.md": basicFloat64Md,

	"calc-range-condition.md": calcRangeConditionMd,

	"dry.md": dryMd,

	"math-pow.md": mathPowMd,

	"square-sum-loop.md": squareSumLoopMd,

	"sum-square-loop.md": sumSquareLoopMd,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"basic-float64.md":        &bintree{basicFloat64Md, map[string]*bintree{}},
	"calc-range-condition.md": &bintree{calcRangeConditionMd, map[string]*bintree{}},
	"dry.md":                  &bintree{dryMd, map[string]*bintree{}},
	"math-pow.md":             &bintree{mathPowMd, map[string]*bintree{}},
	"square-sum-loop.md":      &bintree{squareSumLoopMd, map[string]*bintree{}},
	"sum-square-loop.md":      &bintree{sumSquareLoopMd, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
