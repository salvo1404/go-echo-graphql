package field

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

var graphqlUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Inventory",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.ID},
			"make":      &graphql.Field{Type: graphql.String},
			"model":     &graphql.Field{Type: graphql.String},
			"variant":   &graphql.Field{Type: graphql.String},
			"series":    &graphql.Field{Type: graphql.String},
			"url":       &graphql.Field{Type: graphql.String},
			"url_cg":    &graphql.Field{Type: graphql.String},
			"status":    &graphql.Field{Type: graphql.String},
			"nvic":      &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.String},
			"updatedAt": &graphql.Field{Type: graphql.String},
			"deletedAt": &graphql.Field{Type: graphql.String},
		},
		Description: "Inventory data",
	},
)

type Inventory struct {
	// gorm.Model
	ID        uint       `gorm:"primary_key" json:"id"`
	Make      string     `json:"make"`
	Nvic      string     `json:"nvic"`
	Model     string     `json:"model"`
	Variant   string     `json:"variant"`
	Series    string     `json:"series"`
	Url       string     `json:"url"`
	UrlCg     string     `json:"url_cg"`
	status    string     `json:"nvic"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func GetInventoryField(db *gorm.DB) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(graphqlUser),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			var inv []*Inventory

			// var singleUser *user.User

			fmt.Println("Api call to http://stormtrooper.app/api/v1/inventory/1...")
			// response, err := http.Get("http://stormtrooper.app/api/v1/inventory/1")
			response, err := http.Get("http://stormtrooper.app/api/v1/inventory?dealer_id&id&manu_year&source_ref_id&rego&over_cap&enhanced&eif_generate_xml&source_id&perf_rating&stock_no&sortBy&orderBy&with&status=Live&user_id&dateFrom&dateTo&dateType&paginate=10")
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			}

			data, _ := ioutil.ReadAll(response.Body)

			// Working with 1
			// var result map[string]interface{}
			// json.Unmarshal([]byte(string(data)), &result)

			var objmap map[string]*json.RawMessage
			json.Unmarshal([]byte(string(data)), &objmap)

			json.Unmarshal(*objmap["data"], &inv)

			// Working with 1
			// payload := result["data"].(map[string]interface{})
			// jsonValue, _ := json.Marshal(payload)

			// errore := json.Unmarshal(jsonValue, &singleUser)
			// if errore != nil {
			// 	log.Println(errore)
			// }
			// fmt.Println(singleUser.ID, singleUser.Name, singleUser.Email)

			// u = append(u, singleUser)

			return inv, nil
		},
		Description: "inventory",
	}
}
