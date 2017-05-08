package etl

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"encoding/json"

	"time"

	"github.com/juliotorresmoreno/searchbar/db"
	"github.com/juliotorresmoreno/searchbar/models"
	"gopkg.in/mgo.v2/bson"
	redis "gopkg.in/redis.v5"
)

//Etl define un objeto encargado de la transformacion de los datos
//@param Source: Origen de los datos
//@param Database: Base de datos de la cual se va a importar los datos
type Etl struct {
	temp  map[string]models.Datatable
	words map[string][]string
	cache *db.Cache
}

//NewEtl Constructor del nuevo Etl
func NewEtl() *Etl {
	result := Etl{}
	return &result
}

//Open Establece la conexion con la base de datos cache redis
func (etl *Etl) Open() {
	if etl.cache == nil {
		etl.cache = db.GetCache()
	}
}

//Close Cierra la conexion con la base de datos cache redis
func (etl *Etl) Close() {
	if etl.cache != nil {
		etl.cache.Close()
	}
}

//ImportData Encargado de hacer la importacion propiamente
func (etl *Etl) ImportData(source, database, collection string) error {
	session, db, err := db.GetMongoConnection(source, database)
	defer session.Close()
	if err != nil {
		return err
	}
	data := make([]models.Datatable, 0)
	err = db.C(collection).Find(bson.M{}).All(&data)
	if err != nil {
		return err
	}
	zero := 0 * time.Second
	for _, v := range data {
		id := key(source, database, collection, v.ID.Hex())
		data, _ := json.Marshal(v)
		etl.cache.Set(id, string(data), zero)
		etl.addTerm(id, &v)
	}
	return nil
}

func key(source, database, collection, id string) string {
	source = strings.Replace(source, ":", "_", 1)
	source = strings.Replace(source, ".", "_", 3)
	h := md5.New()
	io.WriteString(h, source+"_"+database+"_"+collection+"_"+id)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//addTerm Agrega un indice de palabras en la base de datos
func (etl *Etl) addTerm(id string, row *models.Datatable) {
	etl.addWord(id, row.AirConditioning)
	etl.addWord(id, row.Constuction)
	etl.addWord(id, row.FirePlace)
	etl.addWord(id, row.Heating)
	etl.addWord(id, row.KitchenFeatures)
	etl.addWord(id, row.LotFeatures)
	etl.addWord(id, row.PoolSpa)
	etl.addWord(id, row.Roof)
	etl.addWord(id, row.Sewer)
	etl.addWord(id, row.Water)
	etl.addWord(id, row.StructureFeatures)
}

func (etl *Etl) addWord(id, term string) {
	_term := strings.Split(term, " ")
	for _, v := range _term {
		if v != "" {
			key := "word_" + v
			data := etl.cache.ZRange(key, 0, -1)
			words, _ := data.Result()
			for _, t := range words {
				if t == key {
					return
				}
			}
			etl.cache.ZAdd(key, redis.Z{
				Score:  0,
				Member: id,
			})
		}
	}
}
