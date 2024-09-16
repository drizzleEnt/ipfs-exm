package main

import (
	"fmt"
	"log"

	"github.com/drizzleent/ipffs-exm/internal/model"
	"github.com/drizzleent/ipffs-exm/internal/service"
	"github.com/drizzleent/ipffs-exm/internal/service/file"
)

var (
	filePath = "./files/sm.txt"
	cidHW    = "QmRCFs5ygExCGqsP1iKoKSxSe8kceof2CyDFwTSYtcvktV"
)

func main() {
	var srv service.IPFSService
	srv = file.NewFileService()

	//ADD FILE TO IPFS
	// cid, err := srv.AddFile(filePath)
	// if err != nil {
	// 	log.Fatalf("error srv add file %s", err.Error())
	// }
	//fmt.Println(cid)

	//CAT FILE FROM IPFS
	// res, err := srv.CatFile(cidHW)
	// if err != nil {
	// 	log.Fatalf("error srv cat file %s", err.Error())
	// }
	// fmt.Println(res)

	//GET FILE FROM IPFS TO DIR
	// err := srv.GetFile(cidHW)
	// if err != nil {
	// 	log.Fatalf("error srv get file %s", err.Error())
	// }

	//ADD JSON TO IPFS
	// u := &model.User{
	// 	ID:    1,
	// 	Name:  "Kiwi",
	// 	Price: 5000,
	// }
	// cid, err := srv.AddJSON(u)
	// if err != nil {
	// 	log.Fatalf("error srv add json %s", err.Error())
	// }
	// fmt.Println(cid)

	//CAT JSON FROM IPFS
	var u model.User
	fmt.Printf("u: %v\n", u)
	err := srv.CatJSON("QmYMF2kVjafVc5NYqMhQpJn7MXTNraAEPdtf6cS3k4ZoUQ", &u)
	if err != nil {
		log.Fatalf("error srv cat json %s", err.Error())
	}
	fmt.Println(u)
}
