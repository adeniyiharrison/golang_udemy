package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adeniyiharrison/golang_udemy/15_nutritional_details_test/models"
)

var (
	bulkEndpoint      = "https://dsom-universal-products.dsom.us-east-1.staging.shipt.com/v1/universal-products"
	universalProducts models.UniversalProducts
)

func checkErrors(e error) {
	if e != nil {
		log.Panicln(e)
	}
}

// ReturnUniversalProducts is used to return bulk GET responses from UP API
func ReturnUniversalProducts(limit int) models.UniversalProducts {
	endpoint := fmt.Sprintf(bulkEndpoint+"?limit=%d", limit)
	resp, err := http.Get(endpoint)
	checkErrors(err)

	if resp.StatusCode != 200 {
		log.Panicf("resp.StatusCode: %d\n", resp.StatusCode)
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	checkErrors(err)

	err = json.Unmarshal(respBytes, &universalProducts)
	checkErrors(err)

	return universalProducts
}
