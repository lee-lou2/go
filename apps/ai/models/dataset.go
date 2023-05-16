package models

import (
	"github.com/lee-lou2/go/configs/db"
	"gorm.io/gorm"
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

// Create 데이터셋 저장
func (obj *Dataset) Create() (*Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	conn.Create(obj)
	return obj, nil
}

// Update 데이터셋 수정
func (obj *Dataset) Update(objId int) (*Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	// output 데이터가 있으면 status 를 2로 변경
	if obj.Output != "" {
		obj.Status = 2
	}
	// obj id 로 데이터 조회
	var dataset Dataset
	conn.First(&dataset, objId)
	// 데이터가 없으면 빈 값
	if dataset.ID == 0 {
		return &dataset, nil
	}
	conn.Model(&dataset).Updates(obj)
	return &dataset, nil
}

// GetDatasets 데이터셋 목록 조회
func GetDatasets(status int) ([]Dataset, error) {
	conn, err := db.GetDB()
	if err != nil {
		return nil, err
	}
	var datasets []Dataset
	conn.Where("status = ?", status).Find(&datasets)
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
