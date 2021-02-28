package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gojektech/heimdall/httpclient"
)

//Item ...
type Item struct {
	ID                     int64  `json:"id"`
	ElasticID              *int64 `json:"elastic_id"`
	PriceMaxBeforeDisxount int64  `json:"price_max_before_discount"`
	Image                  string `json:"image"`
	ShopID                 int64  `json:"shop_id"`
	RefereceitemID         string `json:"reference_item_id"`
	Currency               string `json:"currency"`
	RawDiscount            int64  `json:"raw_discount"`
	PriceBeforeDiscount    int64  `json:"price_before_discount"`
	CmtCount               int64  `json:"cmt_count"`
	ViewCount              int64  `json:"view_count"`
	CatID                  int64  `json:"catid"`
	IsOfficalShop          int    `json:"is_offical_shop"`
	Brand                  string `json:"brand"`
	PriceMin               int64  `json:"price_min"`
	PricwMinBeforeDiscount int64  `json:"price_min_before_discount"`
	Sold                   int64  `json:"sold"`
	Stock                  int64  `json:"stock"`
	PriceMax               int64  `json:"price_max"`
	Price                  int64  `json:"price"`
	Name                   string `json:"name"`
	Ctime                  int64  `json:"ctime"`
	ItemStatus             string `json:"item_status"`
	HistoricalSold         int64  `json:"historical_sold"`
}

func main() {

	db, err := sql.Open("mysql", "root:@tcp(10.8.13.61:3306)/shopee")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("select * from item")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var ItemArr []Item

	for results.Next() {
		var item Item
		// for each row, scan the result into our tag composite object
		err = results.Scan(
			&item.ID,
			&item.ElasticID,
			&item.PriceBeforeDiscount,
			&item.Image,
			&item.ShopID,
			&item.RefereceitemID,
			&item.Currency,
			&item.RawDiscount,
			&item.PriceBeforeDiscount,
			&item.CmtCount,
			&item.ViewCount,
			&item.CatID,
			&item.IsOfficalShop,
			&item.Brand,
			&item.PriceMin,
			&item.PricwMinBeforeDiscount,
			&item.Sold,
			&item.Stock,
			&item.PriceMax,
			&item.Price,
			&item.Name,
			&item.Ctime,
			&item.ItemStatus,
			&item.HistoricalSold,
		)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		ItemArr = append(ItemArr, item)
		// and then print out the tag's Name attribute

	}
	var url string
	var elasticHostname string = "10.8.13.61:9200"
	var elasticsearchIndexName string = "products"
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	for i := 0; i < len(ItemArr); i++ {
		payloadByte, _ := json.Marshal(ItemArr[0])
		payload := bytes.NewReader(payloadByte)

		ID := ItemArr[i].ID
		IDstring := strconv.FormatInt(ID, 10)
		url = "http://" + elasticHostname + "/" + elasticsearchIndexName + "/_doc/" + IDstring

		res, _ := client.Put(url, payload, headers)
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}

}
