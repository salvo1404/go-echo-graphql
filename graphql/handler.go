package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
)

func NewGraphHandler(db *gorm.DB) (*handler.Handler, error) {
	schemaConfig := graphql.SchemaConfig{
		Query: newQuery(db),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
