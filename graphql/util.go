package graphql

import "github.com/graphql-go/graphql"

var result = graphql.NewObject(graphql.ObjectConfig{
	Fields: graphql.Fields{
		"result": &graphql.Field{
			Type: graphql.String,
		},
	},
})
