package file

import (
	"encoding/json"
	"fmt"
	"github.com/elissonalvesilva/interview-hash/my-ecommerce/domain/entity"
	"io/ioutil"
	"os"
)

func ReadDBFile(pathname string) []entity.Product {
	jsonFile, err := os.Open(pathname)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Opened: ", pathname)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products []entity.Product

	json.Unmarshal(byteValue, &products)

	return products
}