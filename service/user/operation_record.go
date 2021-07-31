package user

import (
	"go-temp/global"
	"go-temp/model/user"
)

type OperationRecordService struct {
}

func (ors *OperationRecordService) CreateOperationRecord(o user.OperationRecord) (err error) {
	err = global.DB.Create(&o).Error
	return err
}
