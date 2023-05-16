package models

import (
	"github.com/lee-lou2/go/configs/db"
	"gorm.io/gorm"
	"strings"
)

// Dataset 데이터셋
type Dataset struct {
	gorm.Model
	Instruction string `json:"instruction"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	Status      int    `json:"status"`
	Category    string `json:"category"`
}

// SetCategory 쉼표로 구분된 문자열을 배열로 변환
func (obj *Dataset) SetCategory(categories []string) {
	obj.Category = strings.Join(categories, ",")
}

// GetCategory 쉼표로 구분된 문자열을 배열로 변환
func (obj *Dataset) GetCategory() []string {
	return strings.Split(obj.Category, ",")
}

// Save 데이터셋 저장
func (obj *Dataset) Save() (*Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	conn.Create(obj)
	return obj, nil
}

// Update 데이터셋 수정
func (obj *Dataset) Update() (*Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	// output 데이터가 있으면 status 를 2로 변경
	if obj.Output != "" {
		obj.Status = 2
	}
	conn.Save(obj)
	return obj, nil
}

// GetDatasets 데이터셋 목록 조회
func GetDatasets() ([]Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	var datasets []Dataset
	conn.Where("status = ?", 2).Find(&datasets)
	return datasets, nil
}

// GetInstruction 1개의 데이터셋 조회
func GetInstruction() (*Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	var dataset Dataset
	conn.Where("status = ?", 0).First(&dataset)
	// 데이터가 없으면 빈 값
	if dataset.ID == 0 {
		return &dataset, nil
	}
	conn.Model(&dataset).Update("status", 1)
	return &dataset, nil
}
