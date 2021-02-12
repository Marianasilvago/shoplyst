package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"log"
)


type products struct {

	ProductList []struct{

		Title string `json:"title"`
		Merchant string `json:"merchant"`
	}`json:"productList"`

}



func main() {


	url := "https://api-in-dev.shortlyst.com/shopalyst-service/v1/products"


	Client := http.Client {

		Timeout: time.Second * 2, 
	}

	req,err := http.NewRequest(http.MethodGet, url,nil)
	if err != nil {

		log.Fatal(err)
	}


	res , getErr := Client.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {

		defer res.Body.Close()
	}

	body,readErr := ioutil.ReadAll(res.Body)

	

	if readErr != nil {

		log.Fatal(readErr)
	}

	var productList products

	jsonErr := json.Unmarshal([]byte(body), &productList)

	if jsonErr != nil {

		log.Fatalf("unable to parse value : %q,error:%s",string(body),jsonErr.Error())
	}

	fmt.Printf("HTTP:%s\n",res.Status)

	for i:=0 ;i < len(productList.ProductList) ;i++{

		fmt.Println("Title   :" + productList.ProductList[i].Title)
		fmt.Println("Merchant    :" + productList.ProductList[i].Merchant)
		fmt.Println("-------------------------------------------------")
	}

}