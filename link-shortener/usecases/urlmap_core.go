package usecases

import (
	"fmt"
	"link-shortener/utils"
	"strconv"
)

type UrlMappingCore interface {
	CreateUrlMapping(originalURL string) (string, error)
}

type urlMappingCore struct {
	urlMappingrepo UrlMappingRepo
}

func NewUrlMappingCore(urlMappingrepo UrlMappingRepo) UrlMappingCore {
	return &urlMappingCore{urlMappingrepo: urlMappingrepo}
}

// CreateUrlMapping implements UrlMappingCore.
func (u *urlMappingCore) CreateUrlMapping(originalURL string) (string, error) {
	nodeID := int64(1) // change this to your node ID
	sf := utils.NewSnowflake(nodeID)
	id, err := sf.GenerateID()
	if err != nil {
		fmt.Println("Error:", err)
	}
	str := strconv.Itoa(int(id)) // Convert id to int before passing it to strconv.Itoa()

	return str, nil
}
