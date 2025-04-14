package mongo

import "go.mongodb.org/mongo-driver/v2/bson"

// type utils interface {
// 	StructToBson(any) *bson.M
// }

func StructToBson(s any) *bson.M {
	// convert struct bytes
	bytes, err := bson.Marshal(s)
	if err != nil {
		return nil
	}
	var result bson.M
	err = bson.Unmarshal(bytes, &result)
	if err != nil {
		return nil
	}
	return &result
}
