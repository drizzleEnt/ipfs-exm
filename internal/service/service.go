package service

import "github.com/drizzleent/ipffs-exm/internal/model"

type IPFSService interface {
	AddFile(string) (string, error)
	CatFile(string) (string, error)
	GetFile(string) error
	AddJSON(*model.User) (string, error)
	CatJSON(string, *model.User) error
}
