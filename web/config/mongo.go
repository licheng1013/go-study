/*
 * 作者：Lc
 */

package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"web/common"
)

var Mongo *mongo.Client

func init() {
	//mongodbInit()
}

func mongodbInit() {
	dsn := common.AppConfig.MongoUrl
	log.Println("Mongo:初始化！")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	v, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Panic(err)
	}
	Mongo = v

	database := "testing"
	numbers := "numbers"
	db := NewMongoDB(database, numbers)
	db.Context = ctx

	//插入
	//d := bson.D{{"name", "ok"}, {"value", 3.14159}}
	//one := db.Insert(d)

	// 查询条件
	//list := db.List(bson.D{{"name", bson.D{{"$eq", "Hello"}}}})
	//log.Println(list[0][0].Value)

	//修改
	one := db.GetById("62f11d0e8b0875f27f1234a4")
	m := one.Map()
	m["name"] = 13 //如果 m 与mongo中的数据一致则不进行修改返回0
	num := db.UpdateById(one.Map()["_id"], m)
	log.Println(num) //返回修改次数
	log.Println(db.GetById("62f11d0e8b0875f27f1234a4"))
}

type MongoDB struct {
	Context    context.Context
	Database   string
	Collection string
}

// GetCollection 获取集合
func (m MongoDB) GetCollection() *mongo.Collection {
	return Mongo.Database(m.Database).Collection(m.Collection)
}

// Insert 插入数据
func (m MongoDB) Insert(bson bson.D) string {
	result, err := m.GetCollection().InsertOne(m.Context, bson)
	if err != nil {
		log.Panic(err)
	}
	return fmt.Sprintf("%v", result.InsertedID)
}

// List 返回列表,传如空值则查询所有
func (m MongoDB) List(where bson.D) []bson.D {
	find, err := m.GetCollection().Find(m.Context, where)
	if err != nil {
		log.Panic(err)
	}
	var list []bson.D
	_ = find.All(m.Context, &list)
	return list
}

func (m MongoDB) GetById(id interface{}) bson.D {
	hex, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", id))
	if err != nil {
		log.Panic(err)
	}
	var one bson.D
	err = m.GetCollection().FindOne(m.Context, bson.D{{"_id", hex}}).Decode(&one)
	if err != nil {
		log.Panic(err)
	}
	return one
}

// Update 更新数据
func (m MongoDB) Update(filter, update bson.D) {
	one, err := m.GetCollection().UpdateMany(m.Context, filter, bson.D{{"$set", update}})
	if err != nil {
		log.Panic(err)
	}
	if one.UpsertedCount < 1 {
		panic("修改数据失败！")
	}
}

// UpdateById 根据Id更新插入数据 返回修改次数,如果修改的数据与mongodb中的一样则更新数量为0
func (m MongoDB) UpdateById(id interface{}, update interface{}) int64 {
	one, err := m.GetCollection().UpdateByID(m.Context, id, bson.D{{"$set", update}})
	//log.Println(one)
	if err != nil {
		log.Panic(err)
	}
	//log.Println(one.MatchedCount)
	//返回更新的数量
	return one.ModifiedCount
}

func NewMongoDB(database string, collection string) *MongoDB {
	return &MongoDB{Database: database, Collection: collection}
}
