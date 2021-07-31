package initialize

import (
	"fmt"
	"go-temp/config"
	"go-temp/global"
	"go-temp/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for _, detail := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				_, err := global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				})
				if err != nil {
					return
				}
			}(detail)
		}
	}
}
