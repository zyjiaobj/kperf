// Code generated by go-bindata.
// sources:
// templates/single_chart.html
// DO NOT EDIT!

package utils

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

var _templatesSingle_chartHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x7b\xeb\x92\xdb\x36\x96\xf0\x7f\x3f\xc5\x19\x25\x19\xa9\xbf\x16\x20\x02\x04\x41\x50\x91\xba\x3e\xa7\xe3\xc4\xae\x92\x93\x54\xe2\xc9\xee\x4e\xaf\x2b\xc5\x26\xd1\x12\xdb\x14\xa9\x25\xd9\xb7\xa4\xfc\xee\x5b\x07\x20\x75\x25\xd5\xb2\x33\xf3\x63\x76\x97\x2e\x8b\x20\x70\x70\xee\x17\x00\x64\x4f\xfe\xf2\xed\x8f\x97\xef\xfe\xe3\xa7\x57\xb0\xa8\x96\xe9\xc5\x8b\x89\xbd\xbd\x98\x2c\x74\x18\x5f\xbc\x00\x00\x98\x2c\x75\x15\x42\xb4\x08\x8b\x52\x57\xd3\xde\x5d\x75\x43\x54\xaf\x1e\xaa\x92\x2a\xd5\x17\x3f\xe9\xe2\x06\xe2\xb0\x5c\x5c\xe7\x61\x11\x4f\x46\xb6\xd7\x42\x94\x51\x91\xac\x2a\x28\x8b\x68\xda\x5b\x54\xd5\xaa\x1c\x8f\x46\x51\x9c\xd1\xdb\x32\xd6\x69\x72\x5f\xd0\x4c\x57\xa3\x6c\xb5\x1c\xdd\xfe\xd7\x9d\x2e\x9e\xfe\xbf\x4b\x3d\xca\x46\x71\x52\x56\x75\x0f\x5d\x26\x08\xdd\xbb\x98\x8c\x2c\xae\x4f\x45\xac\x91\xf5\xaa\xb4\x38\xeb\x07\xa2\xb3\xe3\x78\xed\x03\x5e\x51\x9e\x95\x15\xa4\x7a\xae\xb3\xf8\x0d\x3e\xc0\x14\xae\xd6\xa3\x78\xf5\x57\x61\xb5\x18\x8f\x46\x6f\x3d\xc6\xc1\x63\x7c\x49\x84\x50\xe0\x84\xf8\x6b\x5a\xc0\xc0\x01\x15\x48\x70\x60\xa7\x8f\x98\xbe\xbf\xf7\x87\xed\xe8\x84\xe0\xd4\x07\x26\x3c\xaa\x66\x52\x52\x06\x7e\xa0\x28\x8f\x88\xeb\x50\x05\x9e\x4b\x5d\xf0\x71\x9c\x3b\x20\x03\xea\x62\x63\xe1\x63\x77\x24\x19\x45\x5a\xcc\x71\x28\x23\x52\x52\xdf\x00\x10\xc6\x9d\x99\xa7\x18\x82\x22\x4e\x8b\x88\xe0\x0c\xc2\x1c\x7f\xdd\x74\x15\xce\xfe\xbd\x8b\x2b\x5f\x3a\x10\x48\xe7\x35\x97\x22\x22\x8c\x39\xd4\x03\x87\x70\x07\x85\xa1\x9e\x69\x70\xc7\xf9\x15\x47\x9d\x7a\xb8\x19\x80\x7a\x70\x21\x02\x19\xd5\x33\xb1\xcf\x00\x40\x0d\x70\x8f\x83\x0e\x98\x61\xd2\x0c\x34\xb3\x3b\x99\x12\xd2\x83\x40\xd2\x60\x16\x48\xea\x81\x2b\x05\x95\x11\xe1\x0a\xb8\x43\x05\x71\x03\xd4\x97\x44\x26\x02\xca\x90\x9a\x48\x99\x70\xa8\x0f\xc2\x75\x29\x8b\x18\x36\x5d\x17\x04\xa3\x02\x3c\x0f\xf5\x8a\xda\xc6\xd6\x42\x78\x1e\x15\x91\x2b\xa8\x0f\x0e\x48\x8f\x0a\xc2\x79\x0d\x40\x10\x60\x16\x20\x62\x10\x9e\x30\x68\x88\xeb\x12\x46\x64\x40\x99\xa1\x85\x02\x88\x99\xe7\x05\x86\x39\xe4\x88\x18\x8e\xa4\xb4\xf7\x40\x1c\x51\xb4\xf4\x3d\x1a\x00\x5a\x86\xbf\x76\x85\xa2\x2c\x22\x82\xd3\x00\xdd\x86\xa3\x5a\x38\x0d\x08\x73\x04\x48\x87\xb2\x99\x72\x40\x78\x1c\x61\x38\x4a\xe1\xfa\xc8\x01\xb6\x14\x37\xac\x33\x16\x50\x95\x32\x29\x28\x03\xae\x04\x55\x51\x03\xc7\x41\x32\xca\x0c\x16\x68\xd0\x2d\x5c\xee\x53\x15\x59\x72\x88\xc2\x43\xb1\x91\x19\x41\x0c\xb9\x40\x08\xf0\x7c\x46\x03\x83\x86\x20\x39\x30\x2d\x4b\x8e\x18\x72\x33\xdf\x0f\x70\x92\xf4\xa9\x6b\xf9\x32\x80\x04\xe9\x19\x34\xa4\xc1\xd7\xa9\x03\xe6\x38\xc8\xb1\xc7\xf8\x4c\x06\x1c\x5c\xd7\x04\x18\x30\xc4\x8b\x0f\xf8\xdf\x3c\x60\x2f\x3e\xc8\x80\xa7\x4c\x39\xe0\x32\x4e\x99\x99\x23\x03\xde\x89\xde\x13\xcc\x48\xcd\xa9\x98\x49\xc9\xc1\x15\x3e\x75\x53\xee\x3b\x94\x83\x1b\x50\x37\xe2\x18\x3c\x2e\x12\xf3\xa9\x0b\xae\xa4\x12\x98\x42\xcf\x10\x33\xdf\x43\xa7\x90\x2e\xa7\x5e\x2a\x24\xe5\xc0\xd1\xec\x91\xa0\x12\xb8\xf1\x35\x0c\x50\x21\xa9\x22\x02\xb5\xe3\x0a\xea\xce\x90\x4b\xe5\x28\xaa\x52\xc2\x05\x33\xc1\xeb\x47\x08\xa9\x80\xa1\x92\x3d\x46\x7d\xe2\x53\xdf\x4c\x21\x38\xc5\xa0\x26\x06\xf5\xcc\x47\x38\x21\x78\x44\x98\x09\x67\x45\x15\x51\x54\xe2\x2c\x64\x8b\x18\xb6\xdc\xb5\x18\x8c\x63\x78\x73\x81\xc6\x66\x9c\x70\x41\x39\xf2\x23\x6c\xcb\x53\xd4\x25\xc7\x54\xef\x07\x12\xb5\x9a\xa2\x8d\x98\x72\x22\x26\x69\x80\x4e\xed\x12\x41\x39\x91\x1e\x0d\x88\xab\xcc\xfd\xb5\xf4\xd0\xc9\x1d\x50\x32\x22\x1b\x30\x2f\x40\xce\xb1\x65\xc2\xc9\x99\xb9\xbe\x03\x1c\x13\xd9\x6b\x26\x8d\xab\xba\x28\x90\x43\x3c\x81\x0a\x96\x54\x12\x57\x61\x8c\x05\x48\x13\x98\x72\x66\x8c\x0b\xb4\x9f\xc5\x0a\x88\x0b\x90\x77\x84\x81\x1a\x76\xc1\x1d\x35\x13\xbe\x80\xc0\x55\xd1\x06\x0c\x89\xdb\x96\x25\xde\x88\x41\xd9\x6b\x25\x79\x64\x29\x03\x52\x26\xc6\xaa\xb5\x28\xb3\x5a\xea\x4e\xad\x28\xa6\x30\x6f\xf8\x3e\xf5\x16\xc4\x08\xc1\xd0\xc4\xa8\x6b\xdf\xfc\x0a\xc2\x99\x4b\x7d\xc2\x31\xc5\xd6\x6d\x15\x98\xb0\x65\x12\x33\xab\xf0\x4d\x3f\xa7\xdc\x04\xa5\x20\x5c\x61\x24\x78\x18\x08\x1c\x7d\x41\x92\x40\x60\xb7\x4b\xe5\x25\xf3\x5c\x74\x40\x2f\xa0\x12\xa4\x00\x21\x50\x22\x89\x59\x2a\xa0\x01\xa6\x49\xc7\x07\x25\x80\x05\x82\x7a\xc0\x54\x40\x7d\x4c\x93\xa9\x27\x3d\x70\xa8\x1b\xf9\x0a\x65\x07\x26\x18\x75\x89\x74\x4d\xe6\xc7\xa6\xf9\x05\x87\xe0\xb8\xe9\x37\x3d\x82\x6c\x46\x45\xb7\x5f\x78\xa8\x3b\xf4\x0c\xdf\x23\xbe\x17\x49\x94\x1f\x7f\xc0\xb4\x98\xef\x21\x4d\xe3\xdf\xb6\x9f\x6c\xfa\x6d\xd3\x0c\x81\x93\x12\xdf\x03\x83\x03\xd1\x1c\x05\xdd\xc2\x6f\x5a\x35\x11\x30\xc3\x69\x83\x06\x1a\x34\x5d\x90\xb0\x1e\xac\xfb\x4d\xb3\x61\xa7\xc1\x01\xb5\x50\x5d\x80\x70\x4c\xe2\xd4\x8a\xf3\x7b\x7f\xad\xbc\xf7\x7b\x2b\x89\x52\xa7\x3a\xaa\x5e\xa6\x29\x2e\x26\x60\xba\xbb\x7a\x90\x8e\x0a\x03\x09\x66\xa9\xc0\xcc\x12\x81\x05\x1c\x9a\x1e\xbb\x90\x60\x01\x5f\x3a\x84\x7b\x32\x22\x4a\x51\xce\x85\x71\x2d\x07\x7c\x46\x7d\x5f\x9a\x26\x93\x4e\x69\x1f\xa1\x7e\x6c\xfe\x93\x4d\x37\xd9\x3c\x92\xe6\x11\xa1\xde\xda\x2c\xe5\x44\xc6\x4b\x1d\x89\xf8\x5d\x25\x90\xa2\xbd\x2b\x55\x32\xcc\xfd\xae\xc4\x36\xd4\x7d\x78\x07\xee\x49\x7b\x57\xa8\x17\x03\x63\xda\x75\xdf\xd2\x21\x52\x38\x97\x5c\x7a\x94\x0b\x65\x78\x32\x5e\xed\x52\xc7\x51\xc6\xaf\x19\x8f\x1c\x90\x8a\x06\x01\x07\xee\x30\x03\xe6\x7a\xdc\xac\x99\x5c\x8f\x97\x42\x28\xc2\x95\x85\xc7\xb6\xeb\xf1\xc8\x21\x76\x02\xa9\x27\x60\x27\xa9\x07\xfb\x7b\xda\xbf\xcb\x7e\x39\xa6\x7f\x23\x36\xa6\x28\xc6\x51\xab\x5c\x51\xd7\x64\x4e\x15\x48\xc2\x14\xa7\xd2\x93\x18\xa7\x4e\x20\xd3\x40\xd0\xc0\x37\xb1\x1a\xf8\xf2\x25\xf3\x24\x45\xe1\x9b\xbb\x63\xfe\x19\x93\xfa\x3c\x6a\xec\xd4\x62\x81\xb0\x75\x26\xe1\x3e\x75\x7c\x4e\x94\x4f\xa5\x4a\x99\xc3\xa8\xe7\x4a\x52\xdf\x2f\x15\x96\x6d\x05\x6e\xa0\xa8\xe4\xc2\xac\x2c\x45\xe0\x52\x57\xd8\xb6\x55\xa2\xcb\x5b\x2d\x10\x18\xa3\x85\x5b\x2e\x45\x4c\x3b\x22\x4c\x50\x5f\xa1\xb1\xb9\xa2\x52\x10\x97\x4a\x26\x89\x60\xd4\x51\x8a\x04\x54\x4a\x91\x32\xee\x53\xc1\x05\xa9\xef\x91\xa4\x8e\x30\x95\x0b\xad\x63\x40\xb0\xf4\xb9\x8e\xa8\x1f\xec\xe4\xb7\x8c\x2b\xcb\x13\xc1\x6a\xdd\xe2\x39\x8d\xc2\xc1\x2a\x1c\xac\xc2\xa1\x56\x38\x58\x85\xcf\x38\x26\x5d\x57\x82\x8f\x56\x66\xf2\x92\x29\x49\x7d\x0e\x92\x7b\xd4\x45\x7d\x22\x15\xd7\x31\x33\x6a\x8a\x4b\x29\x05\x3a\x30\xe1\xae\xa0\x8a\xc9\x34\x60\xd4\xe5\x8a\xd4\x37\x61\xbc\xb0\xb9\x05\x3e\x95\x2e\x07\x7b\xbb\x94\xbe\x4b\x85\x8f\x8b\x0b\x4e\x7d\x47\x80\x17\x78\x54\x78\x36\x9c\xcc\x0a\x44\x62\x2d\x3b\xee\xc6\xb8\x82\x08\x4c\xb5\xf2\xb9\x02\x26\x14\x15\x4c\x9a\xe5\x90\x31\x4a\xcd\x53\xcd\x0d\xd4\x37\xcb\x4d\x73\xb3\xdc\xd4\xbc\x5d\xba\x9e\x43\x3d\xae\x40\xb9\x8c\x72\x34\x3a\x57\xd4\x43\xa3\x5b\x92\x78\x8f\xb8\x90\xd4\xf7\xb8\xdd\x6a\xec\x87\x0a\x46\xb2\x61\x8a\x58\xa6\x48\xcd\x14\xa9\x99\x6a\x14\x65\x73\x00\xc6\xd6\xb1\x14\x83\x06\xf3\x0c\xd3\x9c\x2a\xa1\xb0\xdc\x2a\x34\x17\x55\x0a\x84\x4b\x95\xcf\x53\x4f\x51\x2f\xe0\xc4\xde\xc2\xc0\x43\x9f\x82\xfa\x66\xf3\x99\x1f\x50\xe5\x28\x62\x6f\xbb\xf0\x2f\x99\xe7\x53\xb4\xac\xbd\x6d\x42\xca\xc4\xf5\x5e\x60\xcf\x5e\x7d\xff\xea\x87\x6f\x7f\x9b\xbd\xf9\xe1\xd5\x6f\xaf\xdf\x7c\xff\xfa\x1d\x4c\x81\x8b\x76\xa0\x37\xef\x5e\xbd\xfd\xed\xbb\x37\xff\xfe\xea\x5b\x98\x82\x6a\x87\xb9\x7c\xfd\xf2\xe7\x35\x8c\x4f\xbd\x0d\xbd\x9b\xbb\x2c\xaa\x92\x3c\x83\x30\xbe\xbd\x2b\x2b\x1d\x7f\x5f\x24\xf1\x37\x79\x55\xe5\xcb\x81\xdd\x26\x7e\x1b\x56\xe1\x19\xfc\xb1\x53\x3f\xef\xc3\x02\xaa\xbc\x0a\xd3\x4b\xdc\x48\xc3\x14\x36\xa0\xb4\xd0\xf1\x5d\xa4\x07\x03\x33\x3e\x84\xa4\xd2\xcb\x21\x84\x67\x30\xbd\xd8\x43\x82\x57\xa1\xab\xbb\x22\xb3\xb8\xe0\xdc\x00\xd3\x54\x67\xf3\x6a\xb1\x03\xfa\x71\x08\xce\xd9\x01\x07\x96\xe8\xbf\x25\x71\xb5\x80\x29\x7c\x39\xe8\x7d\xd1\x83\x73\xb3\xb7\xaf\xbe\xcd\x97\x6f\xe2\x33\xfa\x80\x63\x83\x96\x99\x49\xa6\x2f\xf3\xbb\xac\x82\x29\x0c\xb6\x04\xf9\x7f\x2d\x0a\x3b\xdf\x16\xce\xf2\xb6\x81\xdb\x28\xff\x0c\x46\xdb\x0c\xbd\x68\x91\xf2\x6d\x58\x2d\x68\xa4\x93\x74\xb0\xa6\x7f\xb6\x41\xb5\x31\xf6\x7a\xee\xc7\x16\x3b\xcd\x75\xf5\xca\x88\xf8\xe3\x0a\x9f\xdf\xe9\xe5\x2a\x0d\x2b\x3d\x48\xe2\x21\x98\xa3\x8a\xe1\x16\xbf\x43\x48\xc3\x6b\x9d\x96\x43\x28\x75\x91\xe8\x72\x08\xe1\x63\x52\xce\xb0\xef\xe7\xbc\x0a\x2b\xdd\x66\x58\x3b\xfd\xfb\x22\xbf\x5b\xbd\x0d\x57\x30\x85\x3f\x3e\x76\xc0\xb4\x0e\x6f\x69\xeb\x26\x2f\x5e\x85\xd1\x62\x30\xc8\xc2\xa5\xee\xf0\x00\xc4\x36\x47\x5a\x3f\x84\x4b\x8d\x15\xac\xdf\x0a\xb2\x0a\x8b\x0a\x1d\x0d\x31\xd1\x72\x95\x26\xd5\xa0\xff\x5b\xff\xec\x00\x36\xb9\x81\x81\x81\x6d\x4c\x75\x01\xce\xbe\x90\xcd\xb5\x4d\xd6\xcc\xb9\x72\xde\x1f\x00\x7e\x04\x9d\x96\xfa\x04\x0c\xc8\xd9\xe1\xec\x56\x06\xff\xb2\xab\x62\xba\x08\xcb\x1f\x1f\xb2\x9f\x8a\x7c\xa5\x8b\xea\x69\xb0\xc6\x7a\xd6\xc5\xf9\xee\xfc\xab\xf5\x84\xf7\x30\x85\x1f\xaf\x6f\x75\x54\xd1\x0f\xfa\xa9\x1c\xec\xc2\x9d\x35\x3a\xf9\x6a\xfb\x04\xa8\x2d\xde\xda\x59\x5f\x1b\xfd\x2a\xab\x69\xad\x09\xef\x46\xeb\x61\xc0\x65\xfa\x61\xb6\xf6\x8b\xdd\x7c\xb1\x0c\x57\x47\x1d\xa4\x0e\x9e\x76\x45\xe0\xbc\x61\xeb\x48\x12\xe5\xd9\x78\x5b\xce\xab\x3d\xa5\xed\x49\xf3\xbe\xc5\xf4\xcf\x49\xb5\x7c\xb2\xeb\xaf\x77\x79\x9e\xa2\x73\x1e\xf2\xd8\x40\xbc\x4c\xd3\x71\x87\x08\xe5\x22\x7f\x18\x43\x55\xdc\x75\x08\x62\x82\x7a\x0c\x7d\xbb\xd4\x86\x30\x4d\xfb\xc7\x24\xde\x59\x91\xb7\x03\xe6\x59\x94\x26\xd1\x87\xf1\x26\xa9\x0c\xba\x1c\x6d\x23\x29\x26\x48\x4c\x97\xf5\x99\x23\x4d\xb2\xa4\x1a\xc4\x79\x74\xb7\xd4\x59\x45\x31\x2b\xa5\x1a\x9b\xdf\x3c\xbd\x89\x07\x49\x7c\x76\x18\x9a\xdb\xf8\x72\x93\xbc\x60\xda\x20\x46\x04\x36\xa1\x0d\x8e\x4f\xb4\xe2\xe9\xf8\x30\xef\xec\x48\x68\x50\x51\x6b\xe2\x2b\xe7\x3d\x8d\x77\xd2\x11\x96\x99\x0e\x6f\xdb\xbe\x1a\x5a\x57\xa6\x2c\x35\x3e\x8f\x96\xea\x9c\xf6\xb1\x9b\xfb\x03\x9e\xb6\x44\x69\x9a\x9d\x93\x1b\x35\x95\x6b\x35\x59\x74\xed\xe4\x0e\x15\xf3\xf1\xd0\x13\x96\x4f\x7f\xcb\xfe\x61\xce\x79\x97\x9d\xe8\x9e\x7b\x5b\x96\xff\x73\xd0\x7f\xbc\x83\xde\x84\x69\xf9\x2f\xe8\xa1\x2f\xba\x9f\x6e\xf2\x02\x06\xeb\x85\x02\x24\xd9\x5e\x01\x6c\x73\x8e\x9d\xdc\x7c\xd5\x6f\x1e\xcd\x94\x3e\x9c\x5b\x54\xef\x5b\x93\x36\x9c\xee\xf8\x06\xcb\x10\x46\xa3\x26\x3d\xaf\x51\x7f\x66\x55\xb2\x6c\xbd\x3f\x39\x2e\x56\x61\xb1\x0c\x87\xa0\x87\x60\x4b\xe8\xf1\x28\xb1\xfa\xab\x57\x51\x85\x5e\xa5\x61\xa4\x07\x7b\xba\x19\x42\xbf\x3f\x04\x76\xdc\xcb\xff\x57\x86\x1b\x2e\xdd\x36\xab\x86\xed\x98\xab\x57\x42\xc7\xd4\xdf\x5c\x9f\x51\x54\xe0\xb9\x65\xe8\xf3\xf8\x8f\xe7\x04\x68\x8d\xc8\xf5\xc8\xbf\x60\xc6\x40\xd7\x88\xf2\xe5\x32\xcf\x9e\x5d\xa1\x85\xf3\x24\x7a\xf7\xb4\xd2\x5d\x25\xb0\x32\x63\x57\x7d\xdc\xb3\x61\x70\x5c\x87\x05\xde\xca\x2a\x8c\x3e\x60\xa3\x4a\x52\x1d\xf7\x5b\x56\x8f\x87\x31\x8c\x5e\xf7\xf7\x3c\x5f\x8e\xe1\x8f\x96\xd1\x42\x97\x55\x5e\xe8\xf6\x41\x9c\xfa\x6b\xa2\x1f\xda\x47\xcb\xf0\x5e\xbf\x2c\xdf\x2c\xc3\xb9\x99\x7e\x44\x31\xf5\xee\x20\x2c\xcb\x64\x9e\x0d\x76\x92\xe4\xf0\x50\x63\x67\x5f\xb7\xed\x63\x0f\xf5\x54\x67\xc3\x0e\x05\xea\xc7\x6a\x5c\xef\x4d\x3b\x01\x7e\xa9\x9e\xba\x31\x80\xa9\x01\x59\xf5\x4b\xf2\xbb\x1e\x03\x93\x9f\xbf\xe8\xa9\xf2\x3c\xad\x92\x55\x27\xab\x45\x32\x9f\xeb\x62\x0c\x7d\xdc\x27\x77\x2c\x64\x70\xe8\xa7\x3c\xc9\x2a\x04\xec\x66\xd8\xfa\x4d\x3f\x2a\xf2\xb2\x0b\x13\x5e\x66\x83\x7e\x0c\x0f\x5e\xd7\x61\xf4\x01\x33\x4c\x16\x5f\xe6\x69\x8e\xfc\x7d\x21\x43\x3f\x50\xde\xe1\x66\xb9\xb9\xda\xe3\xb9\x45\x25\xf0\x19\x16\xe0\x9f\x82\x7c\x95\x97\x09\xc6\xf5\x4e\xd9\x42\xfd\x1d\x4b\x97\xb5\xab\x19\xb8\x4f\x21\x76\x93\x17\xcb\xb0\x32\xa6\xd9\x29\x92\xe1\xb2\x7c\xae\x38\x46\x79\x56\x69\x73\x24\x64\xe1\x31\x9f\xa1\xad\x7f\x0d\xd3\x3b\x0d\xe7\xd0\x9f\x5c\x17\xa3\x8b\x6e\x8d\xdb\x59\xeb\xaa\x82\xf9\xf7\xf9\x9a\x82\x84\x93\xfa\xec\x7e\x52\xde\xcf\xe1\x3e\xd1\x0f\xdf\xe4\x8f\xd3\x9e\x39\x51\x74\xb8\x30\x3f\x3d\xb8\xd7\x45\x99\xe4\xd9\xb4\xc7\x28\xeb\xc1\xe3\x32\xcd\x4a\xfb\xa9\xc8\x78\x34\x7a\x78\x78\xa0\x0f\x2e\xcd\x8b\xf9\x88\x3b\x8e\x33\x2a\xef\xe7\x3d\x30\x67\x5f\xd3\x1e\xe3\x3d\x58\xe8\x64\xbe\xa8\x4c\xfb\xa2\x0f\xe7\xcf\xd6\x90\xfe\x64\x15\x56\x0b\x28\xab\x22\xff\xa0\xa7\xbd\x7e\x73\x32\x17\xa1\xf3\xa1\x26\x7a\x70\x93\xa4\x69\xfb\x48\x6c\xbb\x4f\xda\x79\x9b\xb9\xf6\x68\xea\x07\xbb\x09\xdf\xac\x4a\xea\x37\x19\x66\x3d\x72\x66\x50\x8f\x2e\x26\xa3\xf9\xc5\x04\xc5\x3b\x62\x06\xb0\x87\x9f\xc6\x94\xe7\x53\xab\xdc\x73\xe8\xc3\x9a\xd7\x0d\x3d\xec\x1f\x6f\x06\xee\x4f\x34\xf4\xc7\xbd\xb4\xb8\x7d\xd5\x7e\x5b\x33\xd0\x0e\x77\x52\xa6\xb2\x4a\xea\x8a\xc7\x6b\x73\x42\x3b\x06\xa7\x3d\x0a\xb0\x5a\x8c\x77\x0f\x5d\xda\x01\x9b\x22\xfd\x36\x8f\x31\x59\x2d\xef\x30\x3b\xa6\xfa\x50\xf6\x8e\x5c\x7a\x9d\x3f\x76\xb1\x78\xa3\xc3\xea\x0e\xcb\xd9\x4e\x91\x39\x05\xf1\xbc\x48\x3a\x05\x4f\xf5\x4d\x35\x86\xbe\xfb\x55\x47\x4a\x2d\xd0\xd3\xc7\xd0\x17\x5d\x00\x8d\xe6\x8e\x9f\x77\xb7\xcf\x45\xab\x86\x49\x36\xb3\x29\xbb\x7b\x73\x70\x9d\x17\xb1\x2e\x9a\x54\x5d\xe8\xf8\x24\x7d\x3e\xbe\x7c\x4c\xca\xf1\xde\xe7\x58\xcd\xf5\x7c\x9d\x09\x2b\x3d\xcf\x8b\xa7\x23\xa5\xe6\x1a\x0b\x48\x58\x3c\x7d\x1f\xae\xc6\x76\x45\xd8\x0d\x6b\x3d\xa8\x3e\x3d\xee\x84\x32\x07\xb1\xb3\x24\x3b\x5a\x37\xe0\x84\xed\x54\x73\xe1\x22\xeb\xd9\x42\xd4\x5c\x91\x55\x70\xaf\x98\x5f\x87\x03\x2e\xd8\x10\xb8\xab\xea\x1f\x76\xd6\xfb\xdc\xe5\x6e\x37\x87\xeb\xc3\xf3\xe7\xb8\x2b\xcc\xe1\xfa\x78\xff\xb4\xfd\x53\x8b\xf5\x41\x6f\xcb\xbe\xf0\xe9\x4f\x79\x8d\x49\x79\x47\x5c\xe6\x7f\xa4\x81\x3f\x5f\xd9\xb6\x74\x74\x2c\xaf\x37\xad\xed\x97\x35\x66\x8f\xf3\x5d\x91\x2f\x2f\x7f\xf9\x75\xfb\x25\x4d\xdc\xf1\x5a\xcd\x86\x1c\x4c\xe1\xea\xfd\xc1\x98\x25\xdf\x3e\x96\x6e\x1f\xaf\xb7\x8c\x57\xe1\x75\xaa\x61\x0a\xbd\xde\xc1\x50\xa1\xd3\x97\x45\x01\x53\xc3\x53\xfd\x72\xa5\xf7\x9f\x59\x6f\x77\x23\x66\x5e\x5c\x7c\x49\x93\xf2\xd5\x72\x55\x3d\xd9\x7d\xc4\xc0\x4e\x3d\x83\xbf\xfe\xb5\xc6\xb2\x79\xed\xc2\xda\xd6\x5b\xeb\x93\x9c\x04\xa6\xe0\x7c\x0d\x09\x4c\x76\x27\x7e\x0d\xc9\xf9\x79\xd7\x4a\x0d\x27\x1a\x8f\x45\x1d\xd8\x69\x57\xc9\xfb\xf6\x22\x6b\xe5\x3d\x9f\x42\x6f\x52\x15\x17\xed\xbe\xd2\x2a\x92\x25\x40\xab\x22\x59\x0e\xce\x3a\xdf\xc0\x34\xec\xe4\xd7\xb7\x56\x77\x3b\xd3\x1a\x25\x0e\x7b\x47\x96\x0a\x6b\x5d\xdc\x5a\x5d\xdc\xc2\xa4\x46\xb7\xd6\xc5\x6d\xb7\x2e\xb6\x65\x48\x60\x3a\xed\x7e\xcd\xb5\x0f\x7d\x7b\x32\x34\xec\xbc\xce\x83\x69\xcd\xde\xb3\xf3\x4e\x3e\xac\x80\x75\x4c\x5d\xdd\x02\x01\xd6\x7d\x1e\xd7\x76\x65\xe1\x52\x8f\x6b\x9e\xae\x6e\x3b\xce\xcc\xda\x2e\x5b\xdc\xae\x3e\x61\x46\x9d\x2e\xed\x19\xc0\xc9\xb3\xca\xa7\xe5\x75\x9e\xda\x3d\x53\xfb\x96\x69\xff\xea\xce\x5a\xa7\x41\x9c\xac\xf9\xcf\x70\x04\x93\x98\xae\x92\xc6\x4e\x6b\xbd\xff\x13\xdd\xc1\x9c\xd3\x7d\x16\xc9\xcf\xac\x0d\xf0\xc9\xaa\xf9\xb4\xf8\x83\xbd\xcc\xb4\xb8\x48\xb2\x58\x3f\x4e\x46\xd5\xa2\x23\x47\xed\xf0\xfd\x29\x8a\xdc\x25\xd3\xc3\x4d\x0e\x9c\x43\xef\x54\x52\x7f\x52\x81\xa7\xab\xe4\x90\xcf\xb5\x99\x4f\xe4\xf7\x64\xb5\x6c\x93\x8a\x5b\x49\xc5\xcf\x91\xfa\xd3\xab\x8a\x3d\x3e\x46\xad\xc5\xe9\xd8\x09\xde\xd6\xab\xf3\x72\x91\xdc\x54\x2d\x9f\xb7\x44\x79\x76\x93\xcc\x61\xfa\x67\x3f\x16\x11\xde\xd9\x69\xe7\x7f\x96\x60\xcb\x16\x11\x05\x3d\x7d\xb9\x74\x5b\xfe\xac\xc3\xf8\xbb\x24\xd5\xe5\xe0\x06\x7f\xf7\xfd\x07\x3d\xcb\x0c\xd4\x15\xb2\xcd\xbf\x50\x01\x08\x03\x53\x73\x2b\xaf\x9c\x96\x15\x82\x5d\xf6\x84\xb1\xc6\xd2\x9d\xe9\x07\x40\xa2\x3f\x9b\x8e\x41\x4b\xc9\x46\xba\xa3\x4a\x3f\x56\xe7\x23\x5a\xe9\xb2\x32\x4c\x50\x2c\x0a\x9d\x0b\x04\x8b\x9d\xe6\x59\x9a\x87\x31\xf2\x72\xd2\xab\xcc\x2d\x63\x59\xf3\xed\x2c\x1f\x37\x9f\x36\x0d\x8d\x6c\xe6\x78\x7f\x08\xd5\x22\x29\x69\xa1\xcb\xbb\xb4\xea\x3e\xaa\x37\x73\x69\x94\xea\xb0\x38\xf2\x4a\x25\xda\x3b\x91\xdf\xe2\x87\x5a\x23\x77\xcf\xfd\x72\xd0\xfb\x22\x0a\xb3\xfb\xb0\x24\x2b\x5d\xdc\x90\x58\x57\x61\x92\x12\xe3\x03\xbd\x33\xba\xa8\x96\xe9\xa0\x37\xa9\xae\xf3\xf8\xe9\x62\xfd\xa1\x56\x8d\xbb\x8e\x08\x13\x10\x16\xe0\xd4\xb7\x00\x5b\xba\xc6\xdb\xcb\xf2\x9d\x7e\xb4\xf6\x69\xb1\xe3\xd1\x4c\x11\xa6\xba\xa8\x06\xfd\xbf\x65\xe5\xdd\x6a\x95\x17\x95\x8e\x8d\x92\x5b\x3e\x32\xea\x0a\x4f\xdb\xda\xff\x13\x2c\xdc\xed\x98\x05\xc4\xb4\x87\x2e\x34\x8a\xca\xb2\xb7\xf9\x8b\x2c\x94\x76\x8f\xa1\x9b\x3c\xab\x48\x69\x4f\x5a\xc5\xea\xf1\xeb\x2d\xfc\xeb\xa6\x55\xd8\xee\xbc\x65\x58\xcc\x93\x8c\x54\xf9\x6a\x0c\xdc\xd9\x9e\x08\xeb\x43\x09\x12\xe5\x69\x1a\xae\x4a\x3d\x86\xa6\x75\x04\x7f\x15\x0f\xf7\x7b\x16\x7b\x44\x2d\xda\x31\xb0\xd5\x23\x94\x79\x9a\xc4\xf0\x45\xac\x35\xd7\x72\x97\x3a\x4a\x4e\xc2\x34\x99\x67\x63\x88\x74\x56\xe9\x62\x77\x7c\x15\xc6\x71\x92\xcd\xc7\x40\x7d\xaf\xd0\xcb\x56\x9e\xa8\x71\x2b\x93\xb7\x8e\x68\x4c\xee\x0b\x6e\x06\x1f\xb4\x3d\x18\xba\xce\xd3\xf8\xeb\x16\xad\x8d\xc1\xdb\x9f\x57\xab\xd3\x9e\x38\x39\x47\x38\x32\xd9\x66\x97\x21\x73\xf8\x3a\x06\xe1\x7d\xd5\x4e\x8c\x39\xab\x47\x63\xa3\x4d\x6b\x17\x30\x4e\xca\x55\x1a\x3e\x8d\x21\xc9\x70\xd1\x49\xae\xd3\x3c\xfa\xf0\x1c\x0f\x49\xb6\xba\xab\xda\x39\x61\x8e\xf3\x55\xeb\xf4\x96\x90\x6d\xc7\x10\x38\x7b\xb2\x2c\x6a\x95\x7a\xce\x01\xf7\x8d\x98\x87\x72\x6d\xfb\xa8\xb3\xeb\xdb\x60\x63\x07\xc3\xe5\xe2\xc5\x64\x64\xff\x66\xf3\xc5\xc4\x04\x48\x94\x86\x65\x39\xed\x6d\xe7\x95\x55\x38\xd7\xcd\x5f\x6e\xc6\xc9\xfd\x0e\x08\x6a\x63\x2b\xc6\x0e\xc6\x8d\x0f\xf5\x2e\xcc\xae\xe6\x97\xfc\xae\x88\xf4\x78\x32\x8a\x93\xfb\xad\x29\x56\x97\x36\x6c\x0d\x3a\xc8\xb3\x68\x11\x66\x73\x3d\xed\x6d\xd7\x29\x93\x7c\x6d\xb1\xea\xc1\xa8\xe6\x67\x83\xeb\x80\x74\x7d\x86\xa8\x8b\x7d\xfe\x92\x78\xda\x3b\xb4\x45\xef\x62\x9f\x31\x1b\x87\xed\xd0\x75\xb2\xbd\x98\x8c\x4c\xe3\x90\x9b\xfd\xbf\x07\x35\xab\x85\xf2\xfe\x67\x53\x3c\x60\x0a\xbd\x3f\xfe\xa0\xa8\x95\x8f\x1f\x7b\xbb\x30\xeb\xca\x83\x40\x2d\x6c\x1e\x42\x9f\x54\xc5\x7a\xbd\xe1\x86\xfe\xd9\x21\x92\x53\x3f\x30\xd8\xfa\xe8\x77\x08\xbd\x14\x1d\x73\xab\x82\x7c\x4a\x55\xfb\x67\x54\xb1\x4d\x45\x98\x8c\x4c\xff\x0b\xf4\x6f\xf3\xb7\xc9\xff\x1d\x00\x00\xff\xff\xd3\xbd\x92\xf0\xb3\x3c\x00\x00")

func templatesSingle_chartHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSingle_chartHtml,
		"templates/single_chart.html",
	)
}

func templatesSingle_chartHtml() (*asset, error) {
	bytes, err := templatesSingle_chartHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/single_chart.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"templates/single_chart.html": templatesSingle_chartHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"single_chart.html": &bintree{templatesSingle_chartHtml, map[string]*bintree{}},
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

