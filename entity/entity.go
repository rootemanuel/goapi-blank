package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type RootEntity struct {
	ObjectID      bson.ObjectId `json:"_id,omitempty"bson:"_id,omitempty"`
	Id            int           `bson:"id" json:"id"`
	CampanhaId    int           `bson:"campanhaId" json:"campanhaId"`
	CupomEle      int           `bson:"cupomEle" json:"cupomEle"`
	Valor         int           `bson:"valor" json:"valor"`
	Cenario       string        `bson:"cenario" json:"cenario"`
	Data          string        `bson:"data" json:"data"`
	Ec            string        `bson:"ec" json:"ec"`
	TipoTransacao string        `bson:"tipoTransacao"S json:"tipoTransacao"`
	Qtde          int           `bson:"qtde" json:"qtde"`
	Umi           string        `bson:"umi" json:"umi"`
	QtdeResgate   int           `bson:"qtdeResgate" json:"qtdeResgate"`
}
