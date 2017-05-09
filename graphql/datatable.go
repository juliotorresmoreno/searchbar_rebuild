package graphql

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/juliotorresmoreno/searchbar/api"
	"github.com/juliotorresmoreno/searchbar/db"
	"github.com/juliotorresmoreno/searchbar/models"
)

var cache = db.GetCache()

var tipos = map[string]graphql.Type{
	"id":                 graphql.Int,
	"score":              graphql.String,
	"store":              graphql.String,
	"fire_place":         graphql.String,
	"heating":            graphql.String,
	"kitchen_features":   graphql.String,
	"lot_features":       graphql.String,
	"structure_features": graphql.String,
	"constuction":        graphql.String,
	"roof":               graphql.String,
	"sewer":              graphql.String,
	"water":              graphql.String,
	"air_conditioning":   graphql.String,
	"pool_spa":           graphql.String,
	"location":           graphql.Int,
}

var locationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Fields: graphql.Fields{
		"score": &graphql.Field{
			Type: tipos["score"],
		},
		"stores": &graphql.Field{
			Type: graphql.NewList(tipos["store"]),
		},
		"fire_place": &graphql.Field{
			Type: tipos["fire_place"],
		},
		"heating": &graphql.Field{
			Type: tipos["heating"],
		},
		"kitchen_features": &graphql.Field{
			Type: tipos["kitchen_features"],
		},
		"lot_features": &graphql.Field{
			Type: tipos["lot_features"],
		},
		"structure_features": &graphql.Field{
			Type: tipos["structure_features"],
		},
		"constuction": &graphql.Field{
			Type: tipos["constuction"],
		},
		"roof": &graphql.Field{
			Type: tipos["roof"],
		},
		"sewer": &graphql.Field{
			Type: tipos["sewer"],
		},
		"water": &graphql.Field{
			Type: tipos["water"],
		},
		"air_conditioning": &graphql.Field{
			Type: tipos["air_conditioning"],
		},
		"pool_spa": &graphql.Field{
			Type: tipos["pool_spa"],
		},
		"location": &graphql.Field{
			Type: graphql.NewList(tipos["location"]),
		},
	},
})

//GetData Obtiene los datos
var GetData = graphql.Fields{
	"query": &graphql.Field{
		Type:        graphql.NewList(locationType),
		Description: "",
		Args: graphql.FieldConfigArgument{
			"term": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			term := params.Args["term"].(string)
			if term == "" {
				return nil, fmt.Errorf("Specify the search term")
			}
			data := api.LocationQuery(term)
			result := make([]map[string]interface{}, 0)
			for _, v := range data.Data {
				result = append(result, map[string]interface{}{
					"score":              v.Score,
					"stores":             v.Stores,
					"fire_place":         v.FirePlace,
					"heating":            v.Heating,
					"kitchen_features":   v.KitchenFeatures,
					"lot_features":       v.LotFeatures,
					"structure_features": v.StructureFeatures,
					"constuction":        v.Constuction,
					"roof":               v.Roof,
					"sewer":              v.Sewer,
					"water":              v.Water,
					"air_conditioning":   v.AirConditioning,
					"pool_spa":           v.PoolSpa,
					"location":           v.Location,
				})
			}
			return result, nil
		},
	},
}

//SetData Establece los datos
var SetData = graphql.Fields{
	"importData": &graphql.Field{
		Type: graphql.NewObject(graphql.ObjectConfig{
			Name: "result",
			Fields: graphql.Fields{
				"message": &graphql.Field{
					Type: graphql.String,
				},
			},
		}),
		Args: graphql.FieldConfigArgument{},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			result := map[string]string{"message": "success"}
			return result, nil
		},
	},
	"describeElement": &graphql.Field{
		Type: locationType,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			data := models.Datatable{}
			return data, nil
		},
	},
}
