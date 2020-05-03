package dao

import (
	"go_pay/model"
	"go_pay/util"
)

func SaveFlow(flow *model.Flow) bool {
	util.DB.Table("flow").Save(flow)
	return true
}
