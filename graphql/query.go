package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/salvo1404/go-echo-graphql/graphql/field"
)

func newQuery(db *gorm.DB) *graphql.Object {
	fields := graphql.Fields{
		"inventories": field.GetInventoryField(db),
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "Query",
		Fields: fields,
	}

	return graphql.NewObject(rootQuery)
}
