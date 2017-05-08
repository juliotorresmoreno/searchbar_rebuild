package models

import "gopkg.in/mgo.v2/bson"

//Datatable Modelo de base de datos que define la estructura de la
//tabla que debe tener el origen de datos
type Datatable struct {
	ID                bson.ObjectId `bson:"_id" json:"id"`
	FirePlace         string        `bson:"FirePlace" json:"fire_place"`
	Heating           string        `bson:"Heating" json:"heating"`
	KitchenFeatures   string        `bson:"KitchenFeatures" json:"kitchen_features"`
	LotFeatures       string        `bson:"LotFeatures" json:"lot_features"`
	StructureFeatures string        `bson:"StructureFeatures" json:"structure_features"`
	Constuction       string        `bson:"Constuction" json:"constuction"`
	Roof              string        `bson:"Roof" json:"roof"`
	Sewer             string        `bson:"Sewer" json:"sewer"`
	Water             string        `bson:"Water" json:"water"`
	AirConditioning   string        `bson:"AirConditioning" json:"air_conditioning"`
	PoolSpa           string        `bson:"PoolSpa" json:"pool_spa"`
}
