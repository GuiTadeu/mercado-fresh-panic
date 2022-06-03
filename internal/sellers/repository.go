package sellers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Repository interface {
	GetAll() ([]Seller, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Seller, error) {
	jsonFile, err := os.Open("sellers.json")

	if err != nil {
		return []Seller{}, errors.New("error on opening JSON file")
	}

	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return []Seller{}, errors.New("error on loading JSON data")
	}

	var seller []Seller

	json.Unmarshal(jsonData, &seller)

	defer jsonFile.Close()

	return seller, nil
}

func NewRepository() Repository {
	return &repository{}
}
