package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert 문서 등록
func (c *NewCollection) Insert(document interface{}) (primitive.ObjectID, error) {
	insertResult, err := c.InsertOne(context.Background(), document)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return insertResult.InsertedID.(primitive.ObjectID), nil
}

// Get 문서 조회
func (c *NewCollection) Get(id primitive.ObjectID, document interface{}) (interface{}, error) {
	filter := bson.M{"_id": id}

	// 문서 조회
	err := c.FindOne(context.Background(), filter).Decode(document)
	if err != nil {
		return nil, err
	}
	// 데이터 반환
	return document, nil
}

// Find 문서 조회
func (c *NewCollection) Find(key string) (interface{}, error) {
	results := map[string]interface{}{}

	// 문서 조회
	err := c.FindOne(context.Background(), bson.M{"key": key}).Decode(&results)
	if err != nil {
		return nil, err
	}
	// 데이터 반환
	return results["value"], nil
}

// Delete 문서 삭제
func (c *NewCollection) Delete(id primitive.ObjectID) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 문서 삭제
	deleteResult, err := c.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	if deleteResult.DeletedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

// Update 문서 업데이트
func (c *NewCollection) Update(id primitive.ObjectID, update interface{}) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 문서 업데이트
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

// Push 문서 배열 필드에 값 추가
func (c *NewCollection) Push(id primitive.ObjectID, field string, values ...interface{}) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 필드에 값을 추가
	update := bson.M{"$push": bson.M{field: bson.M{"$each": values}}}
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}

// Increment 문서 필드 값 증가
func (c *NewCollection) Increment(id primitive.ObjectID, field string, value int) error {
	// 쿼리 필터 지정
	filter := bson.M{"_id": id}

	// 필드의 값을 증가
	update := bson.M{"$inc": bson.M{field: value}}
	updateResult, err := c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	if updateResult.ModifiedCount == 0 {
		return fmt.Errorf("document not found")
	}

	return nil
}
