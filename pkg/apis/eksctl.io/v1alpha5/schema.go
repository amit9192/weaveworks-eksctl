// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/schema.json (43.659kB)

package v1alpha5

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

var _schemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5d\x7b\x8f\xdb\xb6\xb2\xff\x3f\x9f\x82\x70\x0f\xee\x69\x00\x7b\x7d\x5a\xe0\x3e\x90\xdb\x1b\xc0\xd9\x4d\x93\xbd\xc9\x3e\xb0\xce\x69\x71\x9b\x2d\x10\x5a\x1a\xdb\xec\x4a\xa4\x4a\x52\xbb\x71\xda\x7c\xf7\x0b\x52\x0f\xeb\x41\x4a\xa4\x2c\xa7\x9b\xf3\x57\xb2\xb2\x38\x1c\xce\xfc\x38\x9c\x19\x0e\xa9\x3f\x9e\x20\x34\xf9\x1b\x87\xf5\xe4\x19\x9a\x7c\x33\x0f\x61\x4d\x28\x91\x84\x51\x31\x3f\x8d\x52\x21\x81\x9f\x32\xba\x26\x9b\xc9\x54\xbd\x28\x77\x09\xa8\x17\xd9\xea\x37\x08\x64\xf6\xec\x6f\x22\xd8\x42\x8c\xd5\xe3\xad\x94\xc9\xb3\xf9\xfc\x37\xc1\xe8\x2c\x7b\x3a\x63\x7c\x33\x0f\x39\x5e\xcb\xd9\x3f\xfe\x73\x9e\x3d\xfb\x26\x6b\x57\xe9\x6a\xf2\x0c\x29\x3e\x10\x9a\x14\x7d\x46\x2c\x0d\x7f\xc6\x32\xd8\x96\x3f\x21\x34\x49\x38\x4b\x80\x4b\x02\xa2\xf2\x14\xa1\x49\x90\x35\x7a\xcb\x36\x1b\x42\x37\xb5\xdf\x7a\x07\x57\x76\x54\xb4\x2e\x9b\x7e\xce\xff\xf7\x79\xba\xef\x1f\xd6\xc0\x39\x84\x57\x3c\x04\x3e\x79\x86\xde\x5b\x79\xc8\x7f\xf8\xb5\x6c\x8b\xc3\x50\xf7\x8c\xa3\xeb\xea\x28\xd6\x38\x12\x50\xbe\x14\x82\x08\x38\x49\xd4\x7b\x8a\xe3\x80\x51\x89\x09\x15\x28\xd0\x2a\x40\x09\xe6\x38\x06\x09\x5c\x20\x0e\x11\x96\x10\x22\xc9\x50\x45\x56\x25\xa1\x8f\x33\x42\x25\x44\x11\xf9\x6d\xb6\x95\x71\x34\x3b\x94\xf0\x93\x8a\x20\xda\x3a\x6a\x0b\xde\xaa\x2a\xa0\x78\x15\xc1\xbb\x5d\xd2\xf8\x01\xa1\x09\x91\x10\x37\x1f\x56\x20\x27\x24\x57\x7d\x4c\xeb\xbf\x86\xb0\xc6\x69\x24\xd5\x0b\x93\xca\x2f\x9f\xab\xaf\x95\x24\x30\xe7\x78\x37\x54\xc1\x55\xce\xc7\xd4\x2e\xf0\x6e\x2d\xe4\xc0\x42\x51\x2e\x64\x4f\x1d\xfb\x92\x37\x6a\x3a\xb3\x00\x15\xf5\x72\xf8\x3d\x25\x1c\xc2\xba\x88\x62\x90\x38\xc4\x12\xb7\xe5\x63\x83\x03\x4e\xc8\x4f\xc0\x45\xc6\xf2\x1f\x26\x9d\x19\xd4\x5e\x53\x7a\xed\x07\xf8\x88\xe3\x24\xd2\x9d\xbc\xaf\xe3\x04\xee\x44\x20\xa3\x13\xc2\xe6\xf7\xdf\xe1\x28\xd9\xe2\x7f\xaf\xc2\xe5\xd7\x27\x06\xe0\x4c\xf0\x3d\x26\x11\x5e\x91\x88\xc8\xdd\x2f\x8c\x3e\x06\xc4\x4e\xab\x06\xc7\x60\x22\x91\x87\xc1\x33\x93\x5d\x63\xbe\xc1\x12\xae\x39\x5b\x93\xc8\x79\xc8\xe6\x2e\x7f\xac\xd1\x3a\x68\xb4\x1b\x22\xdd\x86\xf9\x8a\x48\x33\x05\x82\x63\x2f\x41\x9d\x2f\x2e\xcc\x84\xee\x08\x0d\x8f\x8c\xd6\xfa\xbc\xeb\x05\x6a\x8c\x29\xde\x40\x78\xc9\x42\x78\xc5\x59\x9a\x1c\xa6\xb5\x8b\x06\x35\x57\xbd\x35\x06\x5d\x33\x46\x4b\x00\xf4\x5e\x51\xdc\x68\xfe\x50\x2a\xf0\x06\x7e\xfd\x76\xae\xff\x9d\x6b\xfe\x09\xdd\xcc\x68\xf9\xc6\x53\x84\x69\x88\xde\xe7\x23\x43\xfb\x1f\xca\x46\x70\x27\x66\xf9\xcf\xba\x9d\x98\x3f\xad\xb3\xd0\x63\x1f\x15\x4b\x3f\x60\xb4\xe5\xb0\xfe\x9f\xdb\x89\x9d\x93\xdb\xc9\xf3\x26\xe3\x3f\xcc\xf1\x73\xcd\x5f\xab\x7d\x9b\xa9\xdb\xc9\xf3\xf6\x20\x14\x01\x33\xb8\x4a\x13\xea\x03\xd5\x0b\xd8\xdb\xdc\x3a\x39\x3a\x0e\x24\x46\xc5\xc2\x8f\x8c\x23\x42\xd7\x8c\xc7\x58\x3d\xd2\x82\x2c\xa6\x02\x12\x0a\x28\x06\x6d\x9b\x20\xe2\xa5\xee\xde\x5e\x1d\xb1\xe0\xa2\x44\x01\x01\x07\x29\x5e\xd2\x80\xef\x0a\x06\x1c\xb4\xb9\x6c\x35\x33\x53\x97\x58\xa6\x2d\x7d\x76\x02\x64\x99\x35\x31\x92\xbb\x4f\x02\x2f\x5a\x3f\x5d\x9f\x0e\x75\xa1\xb4\xe5\x9c\x1a\x57\x7f\xd3\x14\x68\x18\xef\x06\xcf\x66\x94\x77\x5a\xc5\x8e\x65\xae\x73\xd9\x37\xaf\xba\x9d\x0a\x6f\xeb\xab\xb1\x98\x8d\xe2\x42\x62\x24\x88\x42\x70\xee\xe2\x4d\x95\x43\xb7\x02\xc4\x21\x89\x70\x00\x21\x7a\x20\x72\x8b\x72\xbd\xa1\xc5\xf5\xb9\xb3\xf3\xe8\x4d\xd8\xe4\x36\xbe\xa4\x61\xc2\x08\x95\xc2\x25\x30\x48\x38\xb9\xc7\x12\x16\x41\x00\xa2\x05\xee\xc2\xb8\xac\x18\x8b\x00\x5b\xe6\x45\x92\xae\x22\x12\xf8\x12\xf0\x02\x70\x9d\x49\x5b\xdf\xa3\xa8\x76\xcb\xa2\x50\x94\xde\x39\x4e\x08\x12\xc0\xef\x81\x23\xc8\xa5\x8a\xb0\xee\xad\x6a\xd4\x9c\xd5\x3b\x88\xb8\x49\xc5\xca\x49\x72\x50\x6e\x31\xdb\x58\xf8\xf2\x23\x04\xa9\x22\x77\xc3\x22\x58\xdc\x5c\xf6\x38\x52\x9d\x2e\x6a\x83\xda\x35\xf0\x98\x08\x65\x4d\xc4\x0b\x96\xd2\x10\xf3\xdd\x10\xea\x4a\x12\x24\x50\x3a\x66\x69\x1d\xbb\xc8\x7b\xdd\xdc\x4b\x69\x59\xa3\x7a\x90\x2b\x9c\x33\x78\x80\x00\x2b\x14\x46\x12\x9a\xb2\x07\x57\xe7\x67\xa7\x47\x9a\x77\x8d\x21\xbb\x0f\xa5\x1f\x35\x0d\x7a\x1e\xd8\x32\x0d\xbf\x03\x47\x23\x5a\x05\x1c\x45\xe8\x7c\x71\x81\xb0\x94\x9c\xac\x52\x09\x02\xb1\x35\xc2\xc5\x84\xf6\x34\x03\x7d\xd4\x2c\xf3\xbe\x81\x68\x07\x2b\x80\x29\x65\x12\xd7\x13\x7e\xdd\xb2\x38\x5e\x60\x5d\x49\x60\x9a\x08\xfc\xf1\xd9\x8c\x73\x2c\x25\x0e\xb6\xd7\x2c\x22\x41\x6b\x9e\x98\x4d\xc0\x39\x8d\x08\x85\x33\x16\xa4\x31\xd0\x56\x87\x26\x75\xa0\x44\x93\x47\x61\xde\x46\xad\xbd\x59\xbf\xea\x7f\x72\x4b\x04\xca\xb1\xa5\xac\xb4\x16\xbe\x8f\x23\x3c\xbc\x97\x5e\x89\x2c\x6e\x2e\x1f\x57\x8e\x24\xc2\x2b\x88\xbe\x5a\xb0\x51\x1c\xc3\xd0\x44\x83\x95\xa0\x48\x70\x30\x2e\xd5\xc4\x7b\x09\xf1\xa3\x3f\x20\xde\x69\x99\xa7\xae\x00\x48\xe2\xcd\xd7\x05\x11\xaf\xa5\x53\x83\xc8\x88\x81\xf6\x3c\x99\x9a\x6d\x75\xd7\x6c\xb7\xd9\xc6\x1e\x7c\x74\x86\x47\x5a\x21\x63\x2e\x97\x14\x11\x1c\xe7\xd6\x2c\x37\x66\xa8\x88\x32\x75\x22\x20\x8b\x72\x52\x3e\xc4\x8b\xf6\xa5\xee\xb4\x9c\x2e\x9b\x98\xb7\x2e\xaa\xdc\xd3\x0d\xf4\x02\x4f\x41\x7c\x44\x65\x64\xea\x56\x1e\x46\x4b\x6a\x9e\x82\xef\xa2\x64\x12\xb2\xce\x92\xf5\xee\x61\x34\xe7\x0b\x87\x4d\x25\x13\xd3\xbf\xa3\x31\xba\xd1\xce\x19\x18\x93\xe4\x57\x67\xf1\xaa\xa9\x9f\xc3\xb6\x8b\x46\xb2\xa2\xb9\x52\x0c\x7c\x1d\xc5\x8e\x95\xfb\xb5\x1b\xa0\xc0\x71\x54\x86\xee\x43\x82\x7f\x27\x62\xa6\x09\x74\xb9\x78\xe7\x62\x91\x54\xfc\xf4\x80\xdd\xc3\x48\x2f\x45\x14\xc4\x47\xb4\x48\x97\x8b\x77\x28\x27\x5b\x37\xd5\x88\x25\xf5\x05\xd0\xcd\x2e\xf5\xd3\x33\x09\xd7\xdd\xe2\x63\x3e\xae\x35\x08\x54\x27\x6b\x12\x60\x09\x8b\x54\x6e\x19\x27\x72\x77\x66\xd8\x82\x70\xf3\xe4\x0f\x71\xd7\x8b\xec\xd3\xd8\xee\x63\x70\x77\x39\x8a\x55\xf6\x42\x6a\x39\x18\x17\x51\x4f\xeb\xea\x35\xb2\x3f\x1a\xe4\x11\x07\x1c\xce\x18\x8d\x76\xa3\x64\x10\x1c\xc8\x19\x01\x9f\xae\x28\x78\xe5\x86\x87\x2d\x59\x96\x1d\x2d\x90\x0f\x8c\xdf\x1d\x6d\x99\xca\x12\xc2\x8f\x9e\x63\x2f\x44\x17\x6a\x68\x0f\x73\x44\x6b\x9c\x77\xa2\x1d\xe7\x8c\x3a\x12\x39\x52\xfc\xcc\x70\x07\x21\x13\x1c\x7f\xba\x3e\x75\x32\xbe\xa9\x64\x8b\x28\x62\x6a\x0a\x9f\x5f\xdf\xff\xc7\xa0\x9d\x8a\x80\x84\xdc\x2d\x9e\xdd\x10\xb9\x4d\x57\x27\x01\x8b\xff\x7c\x00\x7c\x0f\x0a\x01\xe2\xcf\xac\x98\xe5\xcf\xe4\x6e\xf3\x67\x2a\x49\x24\xfe\x24\x09\x05\x79\x72\x7e\x7d\x09\x96\x2c\x4d\x60\xdf\x91\xe9\xe8\xbd\xb5\x8f\x63\xb6\xda\x1f\x25\xc7\xa7\xe7\x67\x37\x87\xe5\xca\x0f\x19\xea\xf0\xcd\xe8\x35\xe3\x68\x0f\x56\xa4\x86\x81\xb0\x10\x2c\x20\x59\xf0\x3b\x45\x70\xb2\x39\x41\x92\xa1\x54\x40\xb6\xed\x25\x20\xc1\x5c\x21\x4b\xbf\xac\x08\x14\x50\xcb\xf1\x85\x14\x4d\xba\x43\x38\x9c\x6d\x59\x1b\xbe\x2e\x10\xfe\x82\x6c\x19\x75\x4a\x06\x57\xb8\x18\xc9\x51\xec\x58\xbd\x53\xf1\x34\x3b\xcc\x6a\xb6\xcf\xe6\x03\xb9\x2f\x93\x6e\x14\x10\xa4\x6a\x49\xcf\x4a\x25\x46\x75\x62\xb6\x98\x67\x5b\xd9\xcb\xe1\x7d\xb4\x10\x96\x70\x98\x69\xe9\x43\x88\xb2\x1e\x74\x8d\x0a\x5a\xbe\xf2\x06\xab\x2b\xa9\xfe\x91\xb6\xdc\x82\x7e\xb8\x2c\x4d\x33\xac\xc1\x64\x31\x09\x30\x07\x04\x44\x6e\x81\x17\xab\x42\x65\xa6\xa8\x91\xb4\x27\xd4\xbe\xe4\x63\x8a\xe4\x16\x04\x68\x22\x77\xb0\x83\x10\xad\x76\x68\xf1\x8b\x6e\x17\x30\x7a\x0f\x94\x00\xad\xa5\xd6\xfa\xa5\xf7\x45\x19\x1b\xb8\xf2\x93\x5a\xd9\x86\x5e\xbe\xac\xb0\x37\xe8\xd2\xbc\x58\x38\xc0\x7b\xda\xb1\xf0\x36\xcc\x4b\xd7\x62\xd7\x69\x40\x46\xf4\x5d\x36\x11\x5b\xe1\x28\xb7\xac\xda\xf1\xc0\x51\x84\x82\x2d\x89\x0a\x17\x64\x5e\xb7\xc9\x9e\x2e\x8d\x3f\xfd\x9a\xa7\xd3\x28\xc7\x74\x4b\x85\xb5\xc4\x33\x5e\xe2\xab\x36\x42\xb6\x56\x10\x46\x39\x8f\x28\xc9\x98\x3c\xf1\x9a\x4a\x4e\x34\xfa\xf7\x33\xbc\xab\x14\xba\xc6\x75\xbe\xb8\x40\x9c\x45\xf0\x77\x81\x16\x37\x97\xc5\x8a\x2d\x19\xe2\x29\x45\x09\x0b\x05\x62\x54\xb2\x82\x67\xbf\xf1\x1e\x44\xbb\xdf\x12\x43\x04\x81\x64\x7c\xcc\x1a\xe0\x65\x4e\x73\x0c\xd7\x2d\x5b\x6e\xb4\xc6\x79\x1a\x81\x50\x03\xcf\x78\x46\xca\x77\x8c\x18\xd6\x05\xee\x22\xd8\x42\x98\x46\x70\x80\x9c\x0f\xeb\xc9\x67\x99\xfb\x52\xee\x4b\x97\x5c\x1f\xb6\x24\xd8\x96\x93\x48\x6c\x59\x1a\x85\x05\xb0\x42\x86\x68\x16\x87\x22\x5d\x09\xa6\x77\x8e\xf3\x69\x97\x49\x04\xc2\x52\x26\x27\xe8\x7c\x8d\x28\xa3\x7a\x26\xde\x93\x10\xc2\xa9\x36\x58\xc5\x8a\xa7\x16\x27\xd5\xb0\xc8\x3f\x3e\x90\x28\x42\x2b\x50\x7d\x85\x7e\x0a\x7a\x24\x2c\x1b\x35\xfd\xd5\x25\xdb\xbd\x3c\x83\x66\x62\xdc\x64\x3c\x8d\x36\xa5\xdb\x4d\x18\x2f\x75\x9e\xcd\x5d\x91\xab\x5b\x4a\x42\x37\x42\xab\xab\x36\x63\xcb\x69\x6c\x36\x12\x6e\x06\x62\x60\x27\x1d\x6b\x74\x69\x2e\x9d\xd6\xea\x6c\x5b\xd7\x79\xc1\x7e\xd4\xc5\x11\x35\xf1\xbe\x49\x57\xc0\x29\x48\x10\x48\x33\x8d\x4a\x18\x55\xd6\xbc\x86\x41\xf6\x33\x20\x23\xf4\xe0\x58\xd0\x31\xa0\xfe\xc2\xc6\x69\x49\x0e\xad\x39\x8b\x51\x66\x04\x47\x94\xc4\x30\xfa\x23\xed\xb6\xd9\x6a\x14\x46\x35\x0a\x07\xf8\x0d\xae\x26\x61\xa8\xc3\x50\x18\x84\x57\xc4\xa9\xb6\x6e\xc5\x98\x14\x92\xe3\xa4\xed\xdd\x23\xbb\x73\xd6\x3a\x99\x55\x45\xab\xea\x08\xd7\x8d\x4f\x07\xa9\xab\xe2\x6d\x23\x2d\x0e\x09\x73\xa3\x73\xa3\xde\x1c\x88\x1b\xdd\x8b\x69\x04\xd3\x0e\x41\x8d\x02\xa8\xfc\x74\x90\x0e\xc7\x4c\x5b\x7e\xd5\x13\x97\xfa\x20\x29\xa1\x1b\xf4\x8a\xc8\xab\x44\xf9\xe8\xfb\x9d\x12\x1d\xd4\x45\x84\xde\xa9\xdf\x49\x56\x94\xa7\xde\x43\x6a\x68\x82\x48\xc6\x77\xee\x10\xfc\xa2\x4c\xd5\x40\xdb\x28\x7a\xac\xe0\xd7\xb6\xbc\xec\x15\x6c\x59\x17\x6c\xa5\x3d\x3b\x84\xf9\x8a\x48\x8e\xf9\x0e\xfd\xef\xf2\xea\x72\xfe\x7f\x8b\x8b\xb7\x65\x55\xa3\x98\x22\x91\x06\x5b\x84\x05\xd2\x99\x07\xc3\x59\x58\xc6\x75\xf5\xab\x2e\x87\x24\xe0\xbb\xcb\x70\x4c\x06\x0c\x2b\x4a\x21\xe0\xd6\x19\xbd\x91\x83\x79\x1c\x93\x1f\x71\x4c\xa2\x71\x2b\xf9\x1e\xf7\x81\xda\x10\x84\x12\xdb\x29\x4e\x70\x40\xa4\x75\xe4\x0a\x14\x1b\xb0\xd8\x39\xe7\x43\xa6\xa5\xe6\xac\xc7\x4c\x09\x15\x12\xd3\x40\x9f\x38\x1f\x55\x0b\x8f\xda\xef\xeb\xf5\xa1\x62\xfc\x71\x49\x3e\x59\x25\xd2\xa9\x9d\x98\xd0\xc1\x6d\x47\xaf\xeb\xca\x33\x75\xf9\xb6\xae\xe1\xde\x8a\xe6\x56\xa2\x8d\xbc\x5e\x98\x2c\x59\x06\xe1\x78\x38\xbc\x84\xe3\x72\xf9\xfa\x5f\x24\x8c\xad\xf0\x7e\xcf\xa2\x34\x06\x17\xd5\x77\x79\xde\x1b\xb2\xc1\xab\x9d\x04\xbf\x9d\x3c\x4b\xab\x0a\xd7\xff\xf5\x8f\xb1\x82\xef\xbd\xd5\xb6\xd9\x91\x0e\x73\x67\x98\x27\x86\x69\x67\x96\x6a\xa7\x8d\x6f\xe0\xb1\x6d\x89\x3a\x27\x45\x13\x82\x0d\x63\x3b\x6a\x38\x80\x29\x7a\xf9\x66\x39\x6b\x9d\xd7\x46\xef\xae\xce\xae\xd0\x4f\x38\x22\x61\xb9\x19\x43\x63\x9c\x24\x10\xa2\x35\x81\xcc\x0f\x08\x91\xdc\x72\xf6\xa0\x88\x00\xe7\xcc\xbd\x86\xe6\x38\xbd\xd7\xdd\x05\x90\x9c\x04\xe2\x94\x45\x2a\x06\xa9\x97\x4f\x5a\xfc\x85\x0d\xc7\x34\x8d\x30\x57\xd0\x70\x76\x1b\xaa\x8d\xc6\xb4\x95\x71\xc6\xff\x5f\xef\x2e\x78\x4d\xcf\xaa\x34\x0c\x83\x19\x05\xba\x3a\xc9\xb4\xda\x65\x99\xa7\x00\x6b\x17\xbe\x38\xb3\xab\xcf\xd2\xeb\x03\xcb\xfb\x63\xef\x21\x0b\xc4\xaf\xdf\x6e\xa5\x4c\xc4\xb3\xf9\x5c\xfd\x75\x82\x1f\xc4\x09\x8e\xf1\x27\x46\x4f\x02\x16\xcf\x17\x3f\x2f\xf5\x1d\x21\x3f\x16\x6d\xe6\x2a\x48\x10\x72\xfe\x4f\x01\xfc\x55\x4a\x42\x98\xe3\x07\x31\xdb\x43\x60\x86\xc5\x2c\x1f\x53\x50\x02\xec\x44\x21\xfd\xa9\xf3\x1c\xe8\x1b\xc6\xfe\x70\xfe\x17\x62\xfd\x76\xf2\xdc\x20\xb9\xfd\x91\xff\x62\x66\x15\x55\x59\x0e\x21\xfa\x97\xaf\xfa\x19\xa3\x8a\xc3\x0b\xf1\x86\x7d\xe2\x51\x50\x9e\xc5\x5a\xe7\x67\xda\xd0\x9d\x9e\x9f\xdd\x78\x46\x69\xd5\x96\x75\xf5\x1d\x31\x80\x3a\x20\xc9\xb7\x4c\x20\x20\xeb\x1d\x7a\x1f\xa4\x42\xb2\x18\x2d\x2e\xce\xf7\x57\x61\x64\xcf\x66\x38\x26\x33\x91\x26\x09\xe3\x72\xfe\x74\x8a\x6e\xf5\x0e\xf9\x4c\x88\xf8\x76\x52\xfc\xa5\xfe\xc7\x38\xba\xd5\x67\x6c\x48\x70\x3b\xf1\xf2\x5c\x0a\x26\x5a\x17\x63\x18\x18\x50\xd3\x65\xcf\xaa\x9a\x26\x53\xf4\x6f\xbf\xa7\x4c\xfe\x77\xc1\x55\xf6\x57\xf5\x69\xf1\x84\xf1\xfc\x61\xc6\x65\xf6\x7f\xdf\xc8\xf2\x38\xf1\xaa\xd8\x74\xad\x9c\xa8\x63\x0d\xb2\x5c\xad\xd3\xa2\xe6\xbd\x02\x3d\xc6\x70\xba\x0b\xca\x6f\x49\x4c\x64\x76\xa5\x4b\x96\x06\xd5\xa8\x22\x01\x5a\xfc\xb2\x87\xb4\x82\x43\x6e\xf6\xe7\xdf\x7c\x62\x14\x66\xf8\x01\x73\x98\x65\xe0\xc9\x7e\xf0\xbb\xee\x25\xeb\xb6\x05\x5d\x97\x8e\xf2\x4b\x5e\x5a\xdc\xda\xef\x7b\x59\x31\x29\x23\xe0\x2c\xb8\x03\xc7\x3a\xb7\xd2\xee\xbc\xa8\x36\x35\x12\x0f\x22\x2c\x04\x09\xde\x32\x1c\xbe\xc0\x91\xf2\xe4\xf9\x25\x8e\x1f\xa7\xb2\x17\x79\x89\x22\x20\x9d\xe2\x5e\xe5\xfc\x8a\xac\x6e\x49\x09\xb9\x5c\xdd\x37\x8d\x02\x9f\x7e\x95\x7a\x13\xb7\x88\x53\xe7\x34\xcf\x2e\x97\x07\xd8\xe7\xf7\xa7\x99\xb1\xc3\x61\xc8\x41\xec\x71\x7c\x9f\x04\x33\x5a\xc6\x2e\xf3\x6f\x72\x4b\x99\xf7\x39\x0b\xa9\x98\xe5\x4d\x9e\x66\xdb\x83\xca\x99\x3f\xbb\x5c\xa2\x88\xb1\xbb\xfa\x0d\x35\xfd\xf2\x68\x81\xdb\xbd\xf7\xdb\xc9\xf3\xfa\x08\xf4\x8d\x56\xfd\x1c\xf5\x5a\xcc\x31\x72\x68\xb0\x12\x57\x89\x24\x31\xf9\x04\x56\xff\xc5\x92\x13\xa9\xc9\x27\xbb\xb0\x51\xa0\xf7\x2f\x5f\x2c\x75\xca\x3b\x26\x9f\xb4\x2b\xd7\xeb\xff\xbe\x3c\xfd\xbe\xed\x39\xc2\x4a\xcc\x58\xc1\x57\xc3\xbd\x75\x51\x57\xc1\x8e\xb3\x2b\xeb\xc8\xc5\xed\xe4\x79\x73\x80\x76\x4b\x75\x84\xfc\xe4\x38\x47\x6b\x0c\x84\xaf\x39\xac\xc9\xc7\xa3\x90\x1e\x3d\xa7\x5a\x10\x16\x67\x44\x64\x47\x60\x9c\x6f\xfd\xda\x4b\xda\x48\xc3\xd8\xdd\x5d\xba\x82\x08\xe4\x4b\x5d\x4c\xd9\xbc\x9c\xb3\xa3\x2f\x8f\xcb\x28\x72\x13\x47\x3e\x01\xfa\x90\x77\xf7\x21\x0f\xc9\x1a\x9e\x28\xf9\x44\xe8\x66\x26\xb7\x30\xcb\xdf\xf3\xbc\x88\xcf\xe2\x5f\xb6\xc9\x96\x56\x4b\x31\xf5\x43\xc0\x42\x78\x9e\xff\xf4\xc3\x5c\xff\x95\xf3\x67\x87\xff\x57\x9f\xfa\xbe\x66\xa1\xb8\x06\xae\x30\x33\x2c\x03\xfe\xaf\x92\x3d\x67\xf7\xc0\x39\x09\xe1\x45\xb1\x9d\x7b\xca\xe2\x18\xfb\x5e\xcc\x59\xc3\xe1\x55\x4e\x12\x7d\xc8\x22\xed\x0f\x7f\x17\xa8\xdc\x2d\x4e\x94\x5b\x91\xbd\xee\x05\xee\x92\x68\x86\xd7\x8c\x72\x0e\x57\x1b\x7d\xe3\x80\x13\xde\x1a\xeb\xa3\x74\x01\x41\x57\x7f\x41\x88\x56\xb0\x66\x1c\x1a\x23\x2c\xed\x64\x76\xc9\x0c\xb4\x0e\x1c\xba\xc8\x74\x60\x17\x16\xb1\x1e\xb6\x0b\x53\x63\xec\xa5\x5e\xe2\xd1\xfb\xa2\x08\x7b\xef\x8c\x59\x3d\xc4\x54\xc0\x2c\x7f\x7d\x96\xd7\xc1\xcd\xd6\x8c\xcf\xb4\xc9\xc6\xd1\xfe\xee\xc9\xa7\xda\x33\x2b\xff\xf4\x12\x58\xce\x57\xaf\xc3\xe8\xcc\xcc\xed\xe4\x79\x7b\x8c\xda\x87\xec\x60\xd2\x71\xcb\xaa\x7a\x00\xc0\xf1\x18\xc8\x7e\xf7\xea\x95\xe5\x3c\x93\x48\x98\xbc\x0a\x94\xfe\x3c\x09\x96\xed\xcc\x64\x07\xed\xaf\x75\x41\xa8\x28\xd4\x00\x81\x84\xd8\x16\xd7\xf0\x65\x65\xa7\x44\x0c\xd4\xbf\x33\x51\xe3\x20\xbf\xe6\x9d\x3f\x89\x4d\xa7\x1d\xbf\x1e\xee\xf9\x06\xa4\xc6\xcd\x5f\x79\xb5\x96\x5b\xc4\x9f\x31\x9b\x85\xde\x23\xc7\xfb\x4e\xa4\x8d\x12\xcc\xf6\x28\xf3\x7b\x59\xfb\xc3\xc9\x0e\x1a\xe7\x57\xd7\xd6\x8c\x41\xa7\xeb\x93\x35\x7f\x13\x8b\x37\xb0\x3b\x3f\x73\xbe\x15\xa3\x45\xc1\x21\xce\xea\x68\xfd\x55\x6c\x7e\xb7\xb8\xf6\x8f\xd3\x6a\xdd\xeb\x3d\x53\x74\x8f\x39\xc1\x34\x3b\x61\xf7\x0c\x7d\xb8\x9d\x6c\x92\xef\x6f\x27\x1f\x10\x11\xe8\x55\x7e\x03\xca\x75\xca\x13\x26\x00\x2d\x97\x67\xe8\xdb\x9c\xbb\xa7\x53\xf5\x2e\x61\xdf\xe5\xef\x5e\x73\x76\x4f\x04\x61\x14\x42\xa4\xc0\xa0\x5e\xd6\xaf\x88\xa0\x78\xe5\xdd\x96\xb3\x74\xb3\x4d\x52\x89\xca\x0c\x06\x7a\x7d\x96\xbf\x26\x8b\xd7\x4e\x59\xa4\x1f\xfb\x15\xe6\x9a\x06\x93\x39\x95\x59\xc6\x7c\x93\x7c\x9f\xfd\xa7\x08\x86\xfa\xc7\x57\x6d\x4e\xd8\x77\xad\xe6\xe6\x21\x57\x5b\x89\xa0\xdd\xca\x2e\x85\x5a\x4b\xd9\x6e\x69\x11\x4c\x05\x2e\x9b\xe4\xfb\xfa\x6f\x40\xd3\xb8\x7d\x7f\x7e\xf3\x35\x65\x2c\xd9\x77\xcd\x47\x22\x68\x3f\x92\xdf\x55\xad\x62\xe5\xba\xfd\x27\x0d\x8c\xfa\x16\x66\x0c\xae\xd3\x30\xa7\x16\xec\xf9\x12\x5b\x8a\xa6\x73\xe3\xa0\xa3\xca\xa2\xb3\x24\xa3\xe1\xad\x75\x64\x02\x0d\x81\xa4\x21\x2e\xed\xdb\x7c\xb1\xa5\x09\xcd\xf6\xce\x6c\x4f\xcc\x96\xb5\x63\xd1\xb0\x1b\x73\xf3\x2a\x61\x0f\xd9\xdb\xa9\x88\xb6\xaf\xe2\x92\xfa\xef\xf0\x11\x1a\xde\x69\x23\xf7\x67\xdb\xb4\xe8\x0b\x34\x5d\x22\x6f\x73\x92\xbd\x3b\x6b\x65\x74\xd1\xf3\x67\xa3\xdc\x9f\x5d\xab\x7b\xae\x5c\x4f\x23\xb7\x58\xea\xd3\xd2\xe5\x0e\x8f\xae\x6a\x6e\xfb\xd7\x8e\x57\x69\x0f\xee\x47\x77\xd3\xda\x8e\x7e\x61\xde\x51\xea\xf9\xac\xd2\x22\x8c\x09\x3d\x2d\xbe\xfb\x33\xc8\xd9\x29\x4e\x2f\x8d\x9e\x4f\x2c\x6f\x1d\xc3\x74\x87\xde\x57\xc1\x57\x9e\x98\xda\xe7\xe5\xf7\x05\x10\xf3\xea\x9b\x33\x26\x6a\x7f\xcf\xbf\xa9\x74\x32\x63\xeb\x59\x41\xc9\x2f\x01\x59\x63\xad\x9d\x9e\x3f\x94\x99\xdb\xc9\x73\xe3\x70\x1b\x69\x4a\xaf\x05\xc6\xa8\x6f\x93\x1a\x47\x9c\x4b\x3a\x97\x52\xc3\xb9\x8a\x20\xab\x48\x45\x2b\x2c\x20\x44\xfb\x2f\x2e\xb8\x9f\x55\x38\xa0\x0b\xf3\x0c\x72\xbc\x99\xfe\x51\x5f\x5f\xbc\x5f\xd9\xf5\x51\x15\xef\x83\xe8\x8e\x9b\x16\x83\x0e\xb9\x7b\xd0\x3e\xda\x0e\xd1\xb0\x0b\xec\xfd\xfa\x52\xe1\xe6\x22\x0c\x19\xbd\x2e\xce\x67\x78\xef\x9f\xd5\x9b\x0f\x9c\xf1\x5d\x37\xef\x1a\x70\xd2\xa1\xe6\x2e\x2d\x79\x08\xb9\x53\x46\x23\x9a\x1d\xdb\xf5\xf4\xfb\xd2\x2d\x3f\x1b\xd3\x4f\xcf\x6a\x50\x6c\x38\xb0\x5b\x97\x68\x75\x4e\x37\x7c\xe8\x17\x4d\x70\x92\x5c\x40\x3b\xc5\xe8\xd3\xf6\x9a\xc3\x3d\x81\x87\x61\x24\x52\xc9\x96\x01\x8e\x06\xba\x12\x01\x70\x99\x9d\x53\x1a\xd8\xde\xfa\x81\x3b\xa7\xe6\xb0\x1a\x26\x74\x58\x0f\x6c\xf7\x51\x02\xa7\x38\xea\xa8\x2c\xe9\x6c\xbf\x16\xd6\xbd\xee\xce\x76\x24\xc6\x1b\x78\x91\x92\x28\x1c\x28\xe7\x8f\x37\xf6\xcb\x5a\x0f\xfc\x6e\x47\x8d\x37\x33\xb2\x2c\x12\xb4\xe0\xc8\x30\x39\xec\x98\x6f\x80\xa1\x21\xeb\x86\xca\xa7\xc6\x59\xdb\x14\x93\x19\x9e\xc7\xb0\x76\xca\xd4\x0c\x3e\x94\x68\x26\x62\xb1\x6b\x3d\xa5\x0b\x96\x72\xd8\x6a\xba\xc2\x60\xef\x6d\x16\xb1\xde\xec\x31\x39\x5b\x2a\x54\xe7\xc4\x7e\x36\x9f\xa6\xf1\xca\x96\xe3\x65\xf4\x0c\x54\x0c\xfc\x02\x0b\x38\xa8\xf6\xa9\x20\x74\x0d\x3c\x00\x2a\xf1\x06\x16\x2b\x76\x0f\x07\xd3\x55\xc1\x75\x7e\x7b\x16\x61\x74\x29\x39\x96\xb0\x19\xf6\x79\xa4\x84\xc9\x02\x32\xd7\x8c\xb5\xab\x29\xec\xfc\xf8\x99\x8e\x1a\x50\x4c\x7a\xea\x93\xbf\xa7\x58\x3b\xc7\xd8\x2f\xca\x11\x6d\x80\x39\x0a\x7a\xaf\x3a\xde\xef\x6e\x97\x3b\xca\xea\xf1\xac\x7c\x3c\x77\x3f\x73\xd1\xd5\x59\x6b\xab\xb8\xd1\xcb\xed\xe4\x79\x9d\x1d\xc3\xc9\x88\xea\xa6\xac\x73\x24\x76\x7e\xf6\x28\xb7\xb9\x32\xe6\x40\x54\x2f\xe5\x2c\x72\x9f\x28\x3f\x67\x9f\x97\x1b\x0c\xdb\xa2\x1d\xd4\x81\x35\x60\x79\xcb\x02\x1c\x1d\x52\xc9\x90\x7f\x09\x08\x37\x78\x40\x0a\xf6\x51\xf9\x81\xa0\x43\x86\x3a\x90\x76\x45\xaf\x92\xa7\x96\x1a\x02\x25\x82\xa5\xbe\x52\x70\x04\x19\x64\x77\xfa\xd4\x38\xcd\x2f\xb8\xc4\x31\xa3\x1b\xbd\xd8\xee\x2f\x62\x44\x84\x0e\xae\x6b\x19\xbf\x43\xbb\xb4\xbc\x6c\xf1\x7e\x6a\x9a\x85\x6c\x44\xdf\x28\x06\x31\x60\x54\x72\x16\x89\xd6\x5c\xe8\xa8\x88\x70\x49\xf7\xb9\xd2\xb4\x58\xb4\xe5\x6b\xb7\xe8\x2f\x62\xc3\x22\xaf\xec\x32\xc7\x37\x30\x68\x85\x2e\x1b\x0f\xdd\x31\x2e\x09\x5c\x63\x69\x8d\xbd\x3a\x7d\x04\x96\xf2\xa0\x7e\x95\xe7\xf9\x5f\x57\xad\x36\x14\xf4\x5a\x7b\x56\xb1\x18\xb5\x65\xd5\x42\xbf\x70\x46\x8e\x21\xb4\x11\xd9\xd7\xfa\xd4\x57\x78\xbd\x05\x71\x48\xfe\xc4\x87\xba\x65\x0a\x19\x0a\xb2\x3a\x2f\x5d\x37\x27\x22\x5c\xcb\xb7\x16\x7b\x0a\x46\xc4\x06\x2c\x4e\xd2\xf6\x17\x06\x5c\xc9\x9f\xe6\xcd\xcd\x81\x85\xd7\x27\xbd\xdb\xc4\x2f\x8a\xf6\xe6\xb9\x96\xdd\x21\x65\xa8\x5c\x74\xa5\xbf\xdc\x53\x30\xf7\x60\x89\x16\x9c\xe9\x37\x5c\x64\xdf\x89\x68\xfa\x1e\x74\xc9\x93\x41\x87\x66\xe1\x98\xb3\x0f\x47\xf5\xda\x8b\x43\xcd\x4a\x12\x48\x8b\xe2\xf0\xfd\x10\x13\xcd\x9e\x19\xb6\x30\xcd\x9e\xce\xb9\x66\x3e\x52\xe0\x96\xb4\x63\x2c\x0a\xd9\x83\xf5\x0b\x35\xdd\x27\x7f\xf4\xae\xd6\xb0\x2a\xad\x2d\xe0\x90\x33\xd6\x5a\x51\x06\x7e\xe8\xde\x20\xbd\xd7\x79\x0f\x47\x5e\x77\x0a\x21\x4c\x8d\x3a\x31\x89\xda\x28\x85\x91\x91\x5d\x16\xdc\x01\x3f\x0a\xca\xfd\xe8\x3b\x23\xbe\xd4\x99\xcb\xf9\xfc\x24\xbd\x06\xfe\x4f\x4a\xac\xdf\x1f\xea\x04\xef\xe6\xb0\xe6\x31\xc4\x8c\xef\x0e\xa1\x40\xd3\xf8\x6a\xad\x5a\x1f\x2b\x25\x53\x11\x90\x79\xdc\xd6\xe1\x98\xb9\x1c\x19\xa2\x05\xf6\x8f\x82\x4f\x0f\xe2\x3d\xe0\x3c\x6d\xb9\x1a\xc3\x92\xa5\xae\x06\xec\xdc\x94\xa8\x3d\x2c\x1b\x37\xfa\xaa\xa9\x25\x72\xa4\xd5\xd3\x95\x76\x8f\xda\xce\x2d\xba\xb0\xdf\xc8\x19\xe1\xe0\x2e\x22\xa2\x35\x99\xff\xd2\x2c\xf7\xc3\x96\x48\x78\x1c\x5c\x79\x01\x70\xcf\xf7\xd4\x24\xe2\x91\x11\x59\xc0\x1d\x29\xde\x9b\xf1\xcd\x38\xc0\xf4\xee\xa2\x07\x9f\x17\xed\x28\xc3\xbe\xd2\xe5\x5f\x9d\xed\x3b\x2d\x67\x44\x50\xae\xf3\xb7\x38\xa5\xc1\x76\x99\x40\xeb\x23\x67\x6e\x19\x0e\xf3\x35\xb5\x56\x06\xbc\xc0\x52\x10\x9f\x9a\x86\xdc\x39\x94\x91\x71\x54\x7e\x06\x78\x1c\xcc\x74\x92\xeb\xc1\xc7\xd2\x14\x27\x5a\x11\x22\xb6\xa9\x54\xbe\xe5\x6b\x96\xb6\xbf\xff\xe0\x1f\x58\x2e\x6b\xe4\x8c\x80\x90\x58\xdc\x8d\xed\xbb\xef\x19\x78\x87\x85\xc3\x97\xf3\x0e\xb2\x50\x75\x91\xb5\x46\x36\x32\xb2\xf6\x91\xed\x51\xac\x93\x17\x79\x67\xe4\x2d\x2d\xa8\xb2\x3b\x40\xe2\xe5\x01\x21\xa1\x24\x31\xfc\x4c\x68\xc8\x1e\xbe\xb6\xdc\xe3\x7e\xdc\x96\xe1\x8c\x8d\xa6\x5c\x31\x68\xab\x34\x73\x1c\x44\xf9\x76\xe1\x8c\x2a\x3d\xb7\x5d\x56\x3d\xce\xe8\xcb\x8f\x09\x07\xe1\xf0\x01\x69\x23\xa0\x0e\x84\x23\x16\x77\x0e\x07\x9a\x46\x85\x4c\xd1\xe5\xd4\x2a\x86\x91\x91\xa4\x7a\x3c\x0a\x7e\x1c\x09\xf7\xa1\xa6\x9d\xcb\xb4\xe2\x65\x8d\xa3\x68\x85\x83\xbb\x77\xec\x2a\x2f\x1f\x18\x56\x51\x9f\x30\xb9\xaf\x3b\x18\x14\xd9\xa7\x92\x44\xe4\x13\xdc\x80\x00\x7e\x0f\x61\x59\xb1\x73\xa4\xaa\xad\x06\xc7\xd3\x4e\x91\xb8\xb0\x39\xb6\xb5\xca\x95\x78\x1c\x3b\xe5\x4e\xbc\x86\xb5\xab\xf6\xd7\x12\xec\x96\x88\xc5\x31\x91\x45\x8b\x0b\x4c\xc9\x1a\x84\x3d\x6d\xd3\x09\x2f\x7d\x8e\x68\xd4\x8a\xe7\x81\x5f\x29\xe9\xa1\xca\x01\x87\x57\xd4\x7e\x41\xdf\x01\xb7\x5c\x3f\x10\xb9\x7d\x0d\x51\xeb\x66\xa3\x91\xa6\x83\x4d\x5b\x2d\x25\x18\x65\x68\xe2\xd3\x24\x96\x51\xe6\x88\xcf\x17\x1f\x14\xd8\x8b\x0f\x64\x94\xdf\x2a\xba\x03\x48\x6a\x1f\xb8\xca\x6e\x16\x86\xc6\xe7\x1e\x10\xa1\x48\xec\x68\x70\x9c\x6f\x51\x8c\xc3\x59\x6d\x76\x1a\xbe\x79\x68\x9b\x9c\x2c\x95\x49\x2a\x1d\xb6\xbe\x7d\xf1\x9f\x1d\xa5\x1d\x95\x68\xb6\x97\x7c\x38\x49\xbf\xe5\x21\xeb\xd4\x34\xb6\xa9\x51\x8c\x63\x43\x3b\x04\x89\x49\x94\x7f\xa3\xe4\xf7\x94\x04\x77\x42\x62\x2e\x8b\xef\x3d\x96\x9f\x33\x01\xf5\x46\x13\x32\x38\x2c\x31\x36\xc2\x77\x55\x8e\xc5\x4a\x0d\xbc\x37\xf5\x8f\xe7\xd8\x33\x8e\x1c\x53\x7b\xb1\xfc\x20\x80\x41\x8c\xc9\xb8\xab\xcb\x3a\x4a\x3f\x8e\x3e\xb9\x12\x2c\xb7\x8f\x20\xc6\xab\xe5\xb7\xf4\xa9\xe9\xe5\xf2\xb5\x5b\x1d\x8d\xdf\x78\x53\x3e\xae\x52\x52\x61\x3f\x3b\x70\x24\x2b\xa2\x86\x30\x6d\x61\xb7\xa5\x53\x13\x72\x9a\x8c\xb7\xf0\xda\xa9\x87\x2f\xbe\xd0\x36\xbf\x95\xb4\xbf\x67\x33\xfb\xd6\xd2\x71\x16\xd1\xfe\x5e\x6b\x36\x66\x09\x01\x07\x29\xf2\x43\xef\x4e\x27\x00\xee\x60\xb7\xb8\xb9\x74\x2f\xfd\xcf\xdf\x3f\x4a\xf4\x6b\xe3\xe5\xd0\x2f\x46\xb6\x2b\xa3\xdf\x5c\x2c\x11\x94\x52\x2a\x3e\x13\xea\xff\xb5\x07\x3f\xea\x35\x5d\x0d\xb8\x1e\xbe\xa2\x4c\xcb\xbc\x6e\x55\xbf\x52\x74\x7e\x5d\x5c\x0f\xab\xbc\xaa\xd3\xf3\xb3\x1b\x44\x99\xd4\xcc\x3a\x0f\xb7\x87\xcc\x93\x42\xd5\x9f\x9f\x7c\x7e\xf2\xff\x01\x00\x00\xff\xff\xf4\xb3\x69\x0a\x8b\xaa\x00\x00")

func schemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_schemaJson,
		"schema.json",
	)
}

func schemaJson() (*asset, error) {
	bytes, err := schemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x9e, 0xd8, 0x31, 0xa0, 0x72, 0x45, 0x19, 0x76, 0x71, 0x72, 0x10, 0xf0, 0x96, 0xa9, 0x78, 0x19, 0x93, 0x51, 0x7e, 0xcb, 0x19, 0x39, 0x7, 0x1c, 0xb0, 0x8f, 0xeb, 0x38, 0x1b, 0xa1, 0xa}}
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
	"schema.json": schemaJson,
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
	"schema.json": &bintree{schemaJson, map[string]*bintree{}},
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
