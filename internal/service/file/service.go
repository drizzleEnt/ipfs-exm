package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/drizzleent/ipffs-exm/internal/model"
	shell "github.com/ipfs/go-ipfs-api"
)

type fileService struct {
}

func NewFileService() *fileService {
	return &fileService{}
}

func (fs *fileService) AddJSON(info *model.User) (string, error) {
	res, err := json.Marshal(info)

	if err != nil {
		return "", fmt.Errorf("failed to marshal struct into json: %s", err.Error())
	}

	reader := bytes.NewReader(res)

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(reader)
	if err != nil {
		return "", fmt.Errorf("failed to add json to ipfs: %s", err.Error())
	}
	return cid, nil
}

func (fs *fileService) CatJSON(cid string, info *model.User) error {
	sh := shell.NewShell("localhost:5001")
	data, err := sh.Cat(cid)
	if err != nil {
		return fmt.Errorf("failed cat %s %s", cid, err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	err = json.Unmarshal(buf.Bytes(), info)
	if err != nil {
		return fmt.Errorf("failed unmarshal %s %s", buf.String(), err.Error())
	}
	return nil
}

func (fs *fileService) AddFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %s", filePath, err.Error())
	}
	defer f.Close()

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(f)
	if err != nil {
		return "", fmt.Errorf("failed to add file to ipfs: %s", err.Error())
	}
	return cid, nil
}

func (fs *fileService) CatFile(cid string) (string, error) {
	sh := shell.NewLocalShell()
	data, err := sh.Cat(cid)
	if err != nil {
		return "", fmt.Errorf("failed cat %s %s", cid, err.Error())
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)

	return buf.String(), nil
}

func (fs *fileService) GetFile(cid string) error {
	sh := shell.NewLocalShell()
	err := sh.Get(cid, "./files/get/")
	if err != nil {
		return fmt.Errorf("failed get %s %s", cid, err.Error())
	}

	return nil
}
