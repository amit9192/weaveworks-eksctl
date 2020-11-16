// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/10-eksclt.al2.conf (985B)
// assets/bootstrap.al2.sh (1.246kB)
// assets/bootstrap.ubuntu.sh (2.235kB)
// assets/install-ssm.al2.sh (159B)
// assets/kubelet.yaml (464B)

package nodebootstrap

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var __10EkscltAl2Conf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xdf\x6b\xdb\x30\x10\x7e\xf7\x5f\x21\x68\x1f\x36\x88\x6c\xda\xbd\x15\xfc\xe0\x25\x6e\x31\x73\x9d\x12\xa7\xdb\x60\x1b\x46\x91\x2e\xe9\x11\xf9\x64\x24\x39\x6d\x57\xf2\xbf\x0f\xc5\xf1\x96\xb2\x32\xf6\x26\xdd\x77\xf7\x7d\xdf\xfd\x38\x63\xb0\x75\xd2\x6b\xee\x3a\x90\xb8\x46\xc9\xdc\xb3\xf3\xd0\x2a\xa6\xac\xe9\x38\x12\xeb\x09\x3d\x5b\x1b\xcb\xb6\xfd\x0a\x34\xf8\xc9\xe1\x93\xb5\xe2\xa7\x21\x56\x22\xf5\x4f\xec\x92\xbd\xcb\xca\xcb\xf7\x51\xf4\xad\x06\xbb\x43\x09\x3f\xa2\x33\x56\x1a\x29\x34\x6b\xc1\x0b\x25\xbc\x60\x9d\xb0\xa2\x05\x0f\xd6\x5d\xb1\x45\x7e\x53\xcc\xab\x09\xcb\xbe\xd4\xcd\x2c\xbf\xce\xee\xcb\x65\x33\xc4\xa2\x9c\x76\x68\x0d\xb5\x40\xfe\x1a\x35\xa4\x09\x78\x99\x0c\x16\x93\x91\x2b\x06\xda\x45\x67\xec\x46\x9b\x95\xd0\x4c\x90\x62\xce\x0b\x8f\xf2\x95\xc6\xb4\xbc\xaf\x97\xf9\xa2\x99\x55\xf5\x84\x55\xf3\x59\xde\x94\xd9\xc7\xbc\x1c\x3f\xcb\xac\xa8\x96\xf5\x3f\xe5\x8e\xfd\x1e\xd5\x86\x76\xc8\x10\x7f\x43\xec\x40\x59\xdc\x4d\x58\x51\xd5\xcb\xac\x9a\xe6\x4d\x31\xfb\x2f\x6e\x1d\x58\x0f\x0a\x51\xfe\x04\xb2\xf6\xc2\xfa\xf4\xe4\x99\xf4\xce\x26\x2b\xa4\xb1\x80\x7d\x8f\x18\xe3\x9c\x8c\x02\x8e\x5d\x7a\xfe\x72\x54\xde\x9f\x02\x5a\xac\x40\xbb\x11\x1c\xda\xde\x4f\x84\xee\x1e\x44\x3c\xe8\xc7\x68\x12\x24\xe7\x05\x49\xe0\xa8\xd2\xf3\x97\x13\xe3\x23\x57\x2b\x9e\x78\x67\x54\x20\xba\xcd\xbe\x36\x77\xf3\x59\x3d\x42\x16\x36\xe8\x3c\xd8\x83\x5e\xea\x6d\x0f\xa7\xc1\x47\xf4\x0f\xdc\x0b\x24\xff\xdb\xc4\x30\xee\xb1\x5c\x6a\xd3\x2b\xde\x59\xb3\x43\x05\x36\x15\x8f\x6e\x04\x0c\x85\x3a\xb0\xdc\xf6\xe4\xb1\x85\x54\x19\xb9\x05\x3b\x76\x07\xfe\xd1\xd8\x2d\xef\x74\xbf\x41\x4a\x25\xe1\x58\x47\xc8\x57\x48\x5c\xa1\x4d\x13\xd3\xf9\x44\x12\x86\xb1\x9d\xc0\xd2\xd0\x7a\xc0\xc3\x1a\x02\x4e\xe0\x63\x75\xcc\xe8\x8c\xe2\x48\x6b\x2b\x4e\x2c\x60\x2b\x36\x90\x9e\xbf\x84\x2b\xcd\x3f\xd5\x4d\x3e\x5d\x34\xd9\x74\x3a\xbf\xaf\x96\xfb\x58\x6d\x6d\x0c\xd2\xc6\x03\xfc\xfa\x88\xf7\xc7\x68\x9d\x2f\x3e\x17\xd3\xbc\x6e\x66\xf3\xdb\xac\xa8\xf6\x61\xf9\x49\x27\x7a\x07\x57\x1f\xe2\x0b\x0e\x5b\xb7\xea\x51\xab\xf8\xe2\x68\x22\xec\x38\xd8\xc4\xcd\x5f\xb7\x32\x84\xe3\x67\xd1\xea\x3f\xa3\x7a\x2b\x31\x1c\x55\xc8\x8a\x7e\x05\x00\x00\xff\xff\x30\x3f\x63\x4d\xd9\x03\x00\x00")

func _10EkscltAl2ConfBytes() ([]byte, error) {
	return bindataRead(
		__10EkscltAl2Conf,
		"10-eksclt.al2.conf",
	)
}

func _10EkscltAl2Conf() (*asset, error) {
	bytes, err := _10EkscltAl2ConfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10-eksclt.al2.conf", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x4a, 0xbf, 0xa0, 0x4b, 0x49, 0x6d, 0x88, 0xe7, 0x32, 0x9e, 0xe6, 0xa2, 0xaa, 0x6d, 0xbd, 0xa8, 0x2, 0xb8, 0xe7, 0x4a, 0x8e, 0xcb, 0xa8, 0xa, 0x46, 0xea, 0xe0, 0x70, 0x6f, 0x65, 0x2d, 0xc8}}
	return a, nil
}

var _bootstrapAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x61\x4f\x1a\x4d\x10\xc7\xdf\xef\xa7\x98\xe7\xbc\x18\xc9\xd3\xe5\xd4\x58\x13\x51\x9a\x50\x38\xd3\x4b\x11\x48\xc1\x56\x63\xec\x65\xd9\x1b\xca\xc6\x65\xf7\x72\x3b\xa0\x86\x5c\x3f\x7b\xb3\x78\x58\x50\xeb\x2b\x98\x99\xff\xec\xfc\xf7\x37\xb7\x3b\xff\x45\x63\x65\xa2\xb1\x70\x53\xc6\x1c\x12\x70\x0b\x58\x14\xf8\xa0\x68\x1d\xe6\x2a\xc7\x89\x50\x7a\x1d\x1b\x3b\x37\x0e\x89\xb1\xc9\xdc\x48\x52\xd6\xc0\x2f\xa4\x74\x26\x1e\xd2\xdc\x66\x6e\xaf\x06\x4b\x06\x70\x3f\x55\x1a\xa1\x40\x91\x81\x32\x8e\x84\x91\x98\xd2\x63\x8e\xe0\x35\xa7\x90\x59\x06\x00\xa0\x26\x00\x37\x37\x10\x84\xcb\x2d\x51\x19\x40\xb3\xe9\xb3\x07\x65\x00\xb7\xb7\xb0\xbb\x5b\xa9\x7c\xb3\x2f\xfe\x86\x9f\x37\xfb\xfc\xe4\xf6\xff\xd0\x97\x4f\x81\xa6\x68\x56\x07\x02\xa0\x9c\x5a\xa8\x94\x55\xaa\x40\x9a\x17\x4f\xf5\x89\x62\x00\x99\x35\x08\x67\x10\x21\xc9\x08\xef\x9c\x24\x1d\xad\xdd\xd7\x67\x22\x67\x25\x63\x3b\x70\xe9\x10\x92\x8b\xce\x70\x71\x08\x64\xfd\x0d\x61\x86\x24\x32\x41\x82\x8d\xfa\x5f\xe3\x5e\x33\x08\xf7\xe4\xbc\xd0\xc0\xb9\x53\x1a\x0d\x01\xbf\x82\xc1\xe5\x08\xf8\x17\x08\xae\xb8\xb8\x77\x1c\xe5\x21\x5f\x37\x71\xb2\x77\x68\x38\x91\xe6\x0e\xa5\x35\x99\x6b\xc0\xf1\xfe\x7e\x00\x53\xa2\xbc\x11\x45\x07\xc7\x27\xf5\xc3\x8f\x47\xf5\xea\x37\xd2\x82\xd0\x51\x24\x72\x15\xad\x3a\x6b\xc1\x0b\xdc\xd5\xb9\x15\xee\x17\x4e\xde\xb1\xd0\x80\x70\xe5\x3f\x80\xe0\xfd\xd1\xbe\x8d\xfb\xbe\x28\x3c\x08\x3c\x93\x5e\xbf\x13\xa7\xc9\xc0\x5f\x7c\xd3\x01\x68\x2b\x85\xe6\x2a\x5f\x1c\xd5\x02\x96\xf4\x86\xa3\x56\xaf\x1d\xa7\x49\xe7\x95\x70\xbd\x63\xae\xb2\x4d\xe5\xe8\x7a\x10\xff\x5b\xeb\xbf\x87\x5a\xc0\x5a\x3f\x86\xe9\x30\xfe\xf6\x3d\x69\xc7\xc3\xb4\xd3\xbf\x68\x25\xbd\x57\x3d\x0e\x8b\x85\x92\xe8\xa2\xcc\xce\x84\xf2\xc8\x18\x73\x76\x5e\x48\xdc\xda\xf5\xdd\x7c\x8c\x1a\xa9\x8e\x66\x01\x3b\x40\x53\xe5\x40\x0a\x03\x76\x81\x45\xa1\x32\x84\x8b\xd6\x55\x3a\xe8\x77\x86\xec\xaf\xc5\x6e\x72\x1e\xb7\xaf\xdb\xdd\x77\x7c\x6a\x35\x41\x2e\x1f\xa5\xf6\x6e\x57\xa8\xba\xad\xcf\x71\x77\xd8\x0c\xc2\xe5\x46\x58\x7e\x30\x36\x7b\x52\xaf\xc4\xcd\x70\xf9\x7a\x4a\xe9\x9d\x4b\x41\xf0\xe9\x4d\xe3\x2b\xe0\x2b\xfb\x67\x67\x71\xff\xfc\x79\x31\xd5\xa0\x64\x50\x6e\xad\x61\x63\x42\xd2\x29\x5f\x70\xdf\x28\xfa\xb8\x7c\x13\x74\xb8\x7c\x23\x5b\xb2\x35\xa8\x66\xb8\x5c\xff\x6d\xf0\x8a\x4f\xf5\x9e\xfc\xa3\xdd\x1e\x10\xd4\xca\x2d\x3c\xdb\x74\x98\xbf\x0f\x73\x8f\x8e\x70\x26\x49\x43\x26\x70\x66\x0d\x2f\x50\x5b\x91\x6d\xe4\xd1\x88\xb1\x46\xa8\x88\x6c\x14\x1c\x89\x82\x9e\xf3\x7f\x02\x00\x00\xff\xff\x6d\x49\xb4\xd1\xde\x04\x00\x00")

func bootstrapAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapAl2Sh,
		"bootstrap.al2.sh",
	)
}

func bootstrapAl2Sh() (*asset, error) {
	bytes, err := bootstrapAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcf, 0xb, 0x12, 0xec, 0x9f, 0xc4, 0x34, 0x62, 0x3, 0xdd, 0x85, 0x9f, 0xc6, 0x71, 0xab, 0x8b, 0xfb, 0x4c, 0x46, 0x31, 0x3f, 0x8c, 0x45, 0x1b, 0xbd, 0xdc, 0x6d, 0x38, 0xee, 0xaa, 0xe5, 0xf1}}
	return a, nil
}

var _bootstrapUbuntuSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x55\x71\x6f\x1a\xb9\x13\xfd\xdf\x9f\x62\x7e\x1b\xd4\x5f\xd0\xc5\xbb\x4d\xae\x57\xa9\x69\x39\x1d\x97\xd0\x3b\xd4\x14\xa2\x42\xee\x5a\x45\x39\x64\xec\x81\xb5\x30\xf6\xca\x9e\x85\x46\x11\xf7\xd9\x4f\x5e\x76\x09\xa1\x6d\xfe\x62\xed\xf7\x66\xfc\x3c\x7e\x33\x1c\xfd\x2f\x9b\x6a\x9b\x4d\x45\xc8\x19\x0b\x48\xc0\x1d\xa0\xf7\xf8\x55\x53\xb3\x2c\x74\x81\x33\xa1\x4d\xb3\xb6\xae\xb4\x01\x89\xb1\x59\x69\x25\x69\x67\x61\x8e\x34\x59\x8a\xaf\x93\xc2\xa9\x70\xdc\x86\x07\x06\xb0\xce\xb5\x41\xf0\x28\x14\x68\x1b\x48\x58\x89\x13\xba\x2f\x10\x22\xe7\x2d\x28\xc7\x00\x00\xf4\x0c\xe0\xf6\x16\x92\xd6\xc3\x13\xd2\x26\x81\x4e\x27\xee\x9e\x6e\x12\xb8\xbb\x83\x17\x2f\x6a\x56\x0c\x8e\xe0\xbf\xf0\xcf\xed\x4b\xfe\xe6\xee\xa7\x56\x84\xdf\x02\xe5\x68\xab\x84\x00\x28\x73\x07\x35\xf3\x6d\xbd\xe7\x91\x4a\xbf\x25\xcc\x34\x03\x50\xce\x22\xbc\x83\x0c\x49\x66\xb8\x08\x92\x4c\xd6\xc8\x4f\x97\xa2\x60\x1b\xc6\x8e\xe0\x26\x20\xf4\x3f\x5e\x8e\x56\x67\x40\x2e\x5e\x11\x96\x48\x42\x09\x12\x6c\x3c\xfc\xd0\x1b\x74\x92\xd6\xb1\x2c\xbd\x01\xce\x83\x36\x68\x09\xf8\x67\xb8\xbe\x19\x03\xff\x13\x92\xcf\x5c\xac\x03\x47\x79\xc6\x9b\x20\x4e\x6e\x81\x96\x13\x19\x1e\x50\x3a\xab\xc2\x39\xbc\x7e\xf9\x32\x81\x9c\xa8\x38\xcf\xb2\xd3\xd7\x6f\xd2\xb3\x5f\x5e\xa5\xf5\x6f\x66\x04\x61\xa0\x4c\x14\x3a\xab\x22\xdb\xc9\x41\xbd\xeb\xbc\x75\xbd\x0f\x94\x3c\x23\xe1\x1c\x5a\x95\xfe\x04\x92\xe7\x8f\x8e\x61\x3c\xc6\x65\xad\xd3\x24\xd6\x64\x30\xbc\xec\x4d\xfa\xd7\xf1\xe2\xfb\x0a\xc0\x38\x29\x0c\xd7\xc5\xea\x55\x3b\x61\xfd\xc1\x68\xdc\x1d\x5c\xf4\x26\xfd\xcb\x6f\x88\xcd\x23\x73\xad\xf6\x99\xe3\x2f\xd7\xbd\x1f\x73\xa3\x21\xda\x09\xeb\xfe\x3d\x9a\x8c\x7a\x9f\xfe\xea\x5f\xf4\x46\x93\xcb\xe1\xc7\x6e\x7f\xf0\x4d\x4c\x40\xbf\xd2\x12\x43\xa6\xdc\x52\xe8\x58\x32\x16\x5c\xe9\x25\x3e\x79\xea\x45\x39\x45\x83\x94\xa2\x5d\xc1\x11\x50\xae\x03\x48\x61\xc1\xad\xd0\x7b\xad\x10\x3e\x76\x3f\x4f\xae\x87\x97\x23\xc6\x1e\x25\x5e\xf5\xdf\xf7\x2e\xbe\x5c\x5c\x3d\xa3\xd3\xe8\x19\x72\x79\x2f\x4d\x54\xcb\xa4\x20\xf8\xf5\xbb\xc7\x56\xd5\xaa\x0e\x7f\xf7\xae\x37\x7c\xbf\xab\x6a\xeb\xa1\xfe\xda\x3c\xa9\x61\xeb\x61\x6f\xb5\x39\x28\xda\x1e\x18\xd7\x1b\xd6\x68\xef\xb4\x1e\x9a\xcf\x73\x5e\x2b\xae\x1d\x1e\xfb\xe8\x69\x54\xd2\xde\xb0\xa8\x84\x05\x2b\x0a\x10\x46\x8b\x00\xb5\x5a\x8e\x8b\x90\xd6\xdf\xcd\xde\x21\x4d\x92\xd9\xd1\x24\x99\x66\x6f\x4b\x0b\xe4\x8a\xfd\x64\x2c\xdc\x07\xc2\x65\xe4\x79\x0c\x48\x3c\x4e\x16\x54\x8c\x1d\x33\x80\x23\x18\x0f\x2f\x87\xe7\xb1\x9d\x03\x42\xc8\x5d\x69\x14\x4c\x11\x8c\x73\x0b\x54\x20\x08\x70\x85\xfe\x1e\x48\x2f\xb1\x49\x0a\x81\x84\xa7\x00\x65\x71\x52\x65\x58\xe7\x5a\xe6\xa0\x03\xac\x73\x41\xb0\x46\x50\x0e\xb4\x85\xee\xd5\x19\x1c\xef\xb0\xa9\x08\xa8\xc0\x59\x28\x8c\xd0\x16\xb6\x9a\xd4\x36\x81\xb0\x0a\x96\x28\x2c\xc5\xb6\x9f\xc6\x81\xe5\x49\x4c\x0d\xc6\xe5\xd2\x05\x6a\xd8\xa0\x74\x20\xef\x42\xfb\x04\xa6\x25\x81\xa6\xff\x87\x2a\xde\x3a\x02\x69\x50\x78\xc8\xdd\x3a\x06\x19\x27\x54\x7d\xa5\x99\x77\xcb\x47\xe1\xb1\x3e\x6b\x4d\xb9\x2b\x09\x72\xb1\xd2\x76\x5e\x25\x20\x07\xb2\x0c\xe4\x96\x3a\x60\x8c\xdb\x12\x35\x05\x34\x33\x06\xf0\x8c\xa3\x77\xd6\x7a\x9e\xf6\x43\x42\x63\xea\x8a\xc1\x00\x66\x46\xcc\x43\xe7\xb8\x1a\x9c\x89\x75\x0a\xb9\x2e\xf6\x7c\x9a\x6c\x81\xa5\xf8\xca\xa3\xb1\xf6\x3c\xd7\x40\x55\x8c\x11\x53\x34\xa1\x89\xbb\xea\xfe\xde\xbb\x1a\x6d\x4e\x84\x29\x72\x91\x6e\x0f\x4e\xb5\xcb\xf6\x66\xc3\x81\xe7\x4f\xb6\x59\xf4\x0c\xab\xee\xda\x47\x77\x6d\xd9\x1c\x58\x38\xc5\xb5\x9d\x79\xc1\xa5\xb3\x24\xb4\x45\xcf\xf5\x52\xcc\x63\x54\x9c\x20\xbd\x0f\xa3\x49\xef\xe2\xd3\xa4\x7b\x71\x31\xbc\x19\x8c\x37\xa9\x5a\xf8\x14\xa5\x4f\xb7\xf0\x65\xef\x7d\xf7\xe6\x6a\x3c\xf9\xd4\xfb\xa3\x3f\x1c\x6c\xea\xdd\x83\xb1\xb3\x89\xe5\xca\x0a\x51\x06\x3c\xff\x39\x3d\x8d\xae\x9e\x96\xda\xa8\xf4\xb4\x16\x21\x8d\x2b\x15\x2f\xbc\x5b\x69\x85\xbe\x23\xd6\xa1\x01\xac\xe6\x53\x6d\xb9\xd2\xbe\x93\xb9\x82\x32\x69\x75\xfc\xfb\xdd\x83\xa5\xb3\xb3\x2d\x1e\xdf\x25\xe2\x16\x29\x55\x0d\x63\x77\x29\x5f\xda\xd8\x05\x1d\xe5\xe4\x02\x7d\x53\x6e\xa4\xb5\xf3\x0b\x5e\x98\x72\xae\x6d\x47\x5a\x5d\x03\x1e\xe7\x3a\x10\x7a\x1e\x4b\xd9\x21\x5f\xe2\x21\x10\x7d\xc8\x63\x6e\xda\xbd\xd4\xb8\xdb\x1f\x8c\x77\x4f\x59\x75\xb5\xb3\x33\x3d\xef\x1c\x7a\x6a\xbb\x9d\xde\x8b\xa5\x79\xd4\xf9\x3d\x62\x34\x5f\xc3\x6a\x47\x83\x6d\x47\xc4\xe3\x68\x89\xb5\x8c\xf3\xa9\x32\xde\xed\x6f\x77\x9b\x84\xb5\x59\x33\x48\x84\x7f\xc2\x63\xff\x05\x00\x00\xff\xff\xdc\xe3\x6a\xe7\xbb\x08\x00\x00")

func bootstrapUbuntuShBytes() ([]byte, error) {
	return bindataRead(
		_bootstrapUbuntuSh,
		"bootstrap.ubuntu.sh",
	)
}

func bootstrapUbuntuSh() (*asset, error) {
	bytes, err := bootstrapUbuntuShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bootstrap.ubuntu.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2e, 0x6d, 0xef, 0x2a, 0x1f, 0x75, 0xdf, 0x7e, 0x9b, 0x71, 0x2e, 0xaf, 0x4d, 0xe2, 0xe8, 0x1d, 0xea, 0x9d, 0x97, 0x2a, 0xf6, 0xdc, 0xbb, 0x3d, 0x76, 0xeb, 0xd, 0x4a, 0xe6, 0x50, 0x26, 0x3d}}
	return a, nil
}

var _installSsmAl2Sh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xca\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4f\x11\x71\x5d\xe6\x4c\xa9\x44\x0d\xa4\xe9\xd0\xff\x07\xac\xa7\x77\x35\x2b\x61\x96\x0f\xde\xfd\xb6\x35\xcf\xad\x29\xde\xa5\xc0\x28\x75\x88\xcd\x69\x1f\xe7\xc9\xdd\x77\x7b\xaa\xc7\xe9\x1c\x47\xc2\x58\xca\x3a\xba\x78\x82\x1a\x21\x75\x89\x76\xfd\x8e\xac\x40\xaf\xfa\xb2\x64\xc1\x02\xad\x3f\x18\x62\xa9\x2d\xec\x6a\x80\x3a\xf9\x1f\x7e\x01\x00\x00\xff\xff\x93\x2c\xf6\x43\x9f\x00\x00\x00")

func installSsmAl2ShBytes() ([]byte, error) {
	return bindataRead(
		_installSsmAl2Sh,
		"install-ssm.al2.sh",
	)
}

func installSsmAl2Sh() (*asset, error) {
	bytes, err := installSsmAl2ShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "install-ssm.al2.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0xf9, 0xf8, 0x5e, 0xb9, 0xcf, 0xfb, 0x94, 0xb4, 0x85, 0xa3, 0x62, 0xf0, 0x3b, 0x88, 0x44, 0xe3, 0x84, 0xfa, 0x85, 0x39, 0xdb, 0xed, 0xa2, 0x6a, 0x1, 0x7b, 0xe5, 0x49, 0x21, 0xef, 0x3d}}
	return a, nil
}

var _kubeletYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x4f\x4f\xc3\x30\x0c\xc5\xef\xf9\x14\xfe\x04\x6d\x07\x9a\x04\xb9\x8d\x4d\x70\x60\x27\x36\xe0\xec\xa6\xee\x16\x35\x8d\x27\xc7\x19\x7f\x3e\x3d\x5a\x5a\x90\x26\xa1\x9c\x9e\xde\xb3\xdf\x4f\xce\xe0\x63\x67\xe1\x39\xb7\x14\x48\xd7\x1c\x7b\x7f\xc8\x82\xea\x39\x1a\x3c\xf9\x37\x92\xe4\x39\x5a\x18\xa6\x40\xe5\x4a\xa2\x1a\xee\x52\xe5\xb9\x3e\x2f\x5a\x52\x5c\x18\x83\x5d\x27\x94\x92\x85\xa6\x2a\xcf\xb8\x90\x93\x92\x6c\x78\x44\x1f\x2d\xcc\xb2\x0a\xec\x30\x18\x83\x59\x8f\x14\xd5\xbb\x52\x64\x0d\x00\x46\x8e\x5f\x23\xe7\x74\x11\x00\x14\xb1\x0d\xd4\x59\xe8\x31\x24\x32\x00\x1f\xd4\x1e\x99\x87\xc9\x75\xe8\x8e\xb4\xdf\x6f\x2d\xdc\x8c\x4d\xba\x1e\x50\xc9\x97\xfc\xe7\xb2\xb9\x9f\xc3\xc1\x53\xd4\xf5\xea\xd1\x07\xb2\x50\x93\xba\x9a\x86\xe4\x34\xd4\x0e\x2b\x27\x3a\xd1\xb0\xf8\xef\x3f\x98\x91\x3b\xb2\xf0\x3e\x55\xfe\x5b\xbe\x9a\x47\xa8\x2b\x18\xcb\x5f\x8c\x62\xbe\x46\xbc\xb6\x6f\x9b\x64\x4c\x22\x39\x93\xec\xb7\xbb\x07\x66\x4d\x2a\x78\x9a\x61\x8d\x3b\x08\xe7\xd3\x46\xfc\x99\xc4\xc2\xa4\xfa\x64\x4c\x4f\xa8\x59\xe8\x09\x95\xca\x59\x5e\x58\x51\x69\xfe\xaa\x5d\x59\xb7\x26\x51\xdf\x5f\xee\x48\xf3\xb6\x9f\x00\x00\x00\xff\xff\x1f\x2f\xa9\x0f\xd0\x01\x00\x00")

func kubeletYamlBytes() ([]byte, error) {
	return bindataRead(
		_kubeletYaml,
		"kubelet.yaml",
	)
}

func kubeletYaml() (*asset, error) {
	bytes, err := kubeletYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubelet.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe1, 0x70, 0xd5, 0xbb, 0x81, 0xa2, 0xa6, 0x76, 0x99, 0x80, 0xe7, 0xe2, 0x47, 0xc5, 0xa0, 0xe0, 0xb4, 0xe1, 0x42, 0x2c, 0xb0, 0x60, 0xa0, 0xb0, 0x97, 0x53, 0xa7, 0x1a, 0x9, 0xc3, 0x1, 0x6d}}
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
	"10-eksclt.al2.conf":  _10EkscltAl2Conf,
	"bootstrap.al2.sh":    bootstrapAl2Sh,
	"bootstrap.ubuntu.sh": bootstrapUbuntuSh,
	"install-ssm.al2.sh":  installSsmAl2Sh,
	"kubelet.yaml":        kubeletYaml,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"10-eksclt.al2.conf": {_10EkscltAl2Conf, map[string]*bintree{}},
	"bootstrap.al2.sh": {bootstrapAl2Sh, map[string]*bintree{}},
	"bootstrap.ubuntu.sh": {bootstrapUbuntuSh, map[string]*bintree{}},
	"install-ssm.al2.sh": {installSsmAl2Sh, map[string]*bintree{}},
	"kubelet.yaml": {kubeletYaml, map[string]*bintree{}},
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
