package upload

import (
	"errors"
	"github.com/dulumao/Guten-utils/rand"
	"github.com/h2non/filetype"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"
)

const PathSeparator = string(os.PathSeparator)

type Upload struct {
	Path          string
	DirectoryName string
	FileName      string
	Ext           string
	DateTimeName  bool
	Mimes         []string
}

func New() *Upload {
	return &Upload{
		Path:         "web/assets/uploads",
		DateTimeName: true,
		Mimes:        make([]string, 0),
	}
}

func (self *Upload) SetPath(path string) *Upload {
	self.Path = path

	return self
}

func (self *Upload) SetAutoSubDirectory(used bool) *Upload {
	if used {
		self.DirectoryName = time.Now().Format("20060102") + PathSeparator
	}

	return self
}

func (self *Upload) SetFileName(name string) *Upload {
	self.FileName = name

	return self
}

func (self *Upload) SetAutoFileName(used bool) *Upload {
	self.DateTimeName = used

	return self
}

func (self *Upload) SetOnlyMIMESupported(mimes ...string) *Upload {
	self.Mimes = append(self.Mimes, mimes...)

	return self
}

func (self *Upload) GetMIME(file *multipart.FileHeader) []byte {
	src, err := file.Open()

	if err != nil {
		return nil
	}

	defer src.Close()

	head := make([]byte, 261)
	src.Read(head)
	src.Seek(0, 0)

	return head
}

func (self *Upload) SetSaveFile(file *multipart.FileHeader) (*Upload, error) {
	src, err := file.Open()

	if err != nil {
		return self, err
	}

	defer src.Close()

	head := make([]byte, 261)
	src.Read(head)
	src.Seek(0, 0)

	if len(self.Mimes) > 0 {
		for _, mime := range self.Mimes {
			if filetype.IsMIME(head, mime) {
				goto UPLOADFILE
			}
		}

		return self, errors.New("类型错误")
	}

UPLOADFILE:

	self.Ext = path.Ext(file.Filename)

	if self.Ext == "" {
		self.Ext = ".bin"
	}

	if self.DateTimeName {
		self.FileName = time.Now().Format("150405") + rand.Str(4) + self.Ext
	} else {
		self.FileName = file.Filename
	}

	if _, err := os.Stat(self.Path + PathSeparator + self.DirectoryName); err != nil {
		os.MkdirAll(self.Path+PathSeparator+self.DirectoryName, os.ModePerm)
	}

	// Destination
	dst, err := os.Create(self.Path + PathSeparator + self.DirectoryName + self.FileName)

	if err != nil {
		return self, err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return self, err
	}

	return self, nil
}
