package dashboard

import (
	"fms/common"
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/util"
	"github.com/gin-gonic/gin"
)

type InterfaceDashboard interface {
	FaultsSum()gin.HandlerFunc
}

type StructDashboard struct {
	S interface{}
}

func InitDashboardAPI()InterfaceDashboard {
	s := &StructDashboard{}
	return s
}

func (sd *StructDashboard)FaultsSum()gin.HandlerFunc  {
	return func(c *gin.Context) {
		faults := new([]moddle.FaultList)
		db := database.MysqlInit()
		faultTime := make([]moddle.Time,0)
		faultType := make([]string,0)
		faultClass := make([]string,0)
		faultStatus := make([]string,0)
		faultDuration := make(map[moddle.Time]int)
		db.Model(moddle.FaultList{}).Find(&faults)
		_=db.Close()
		for _,v := range *faults{
			faultTime = append(faultTime,v.CreatedAt)
			faultType = append(faultType,v.FaultType)
			faultClass = append(faultClass,v.FaultLevel)
			faultStatus = append(faultStatus,v.Status)
			faultDuration[v.CreatedAt] = v.Duration
		}
		b := make([]common.JSKeyValueCommon,0)
		d := common.JSKeyValueCommon{}
		response.Success(c,gin.H{"sum":util.HandleTimeData(faultTime),"type":util.HandleType(faultType,b,d),"class":util.HandleType(faultClass,b,d),"status":util.HandleType(faultStatus,b,d),"duration":util.HandleDuration(faultDuration)},"查询成功!")
	}
}