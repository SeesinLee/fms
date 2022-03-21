package prometheusAPI

import (
	"fms/database"
	"fms/moddle"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var (
	NumCounts sync.WaitGroup
	WG sync.WaitGroup
	Rate float64
	RLock sync.RWMutex
)
var (
	instance string
	url string
)

func InitGetCode() {
	go getError()
	go callingRate()
	go func() {
		for {
			WG.Add(2)
			go StartCurlApp()
			go F3CallingRating()
			WG.Wait()
			time.Sleep(time.Second * 20 )
		}
	}()

}

func getError () { //应用实例宕机时等待应用恢复，并把错误信息写入数据库，判断是否为发版引起的应用实例宕机
	faultMsg := new(moddle.FaultList)
	var (
		instance AppStatusMg
		ok bool
		a GetCodeInter
		db *gorm.DB
		b bool
	)
	for {
		instance,ok = <- InstanceInfo.Instance
		if ok && Rate < 0.5 {
			RLock.Lock()
			a =<- InstanceInfo.InstanceClass
			RLock.Unlock()
			faultMsg.Status = "待确认"
			faultMsg.FaultSketch = fmt.Sprintf("应用实例：%s状态不正常，疑似宕机", instance.Instance)
			faultMsg.StartAt = time.Now()
			faultMsg.EndedAt = time.Now()
			faultMsg.FaultType = "业务程序故障"
			fmt.Printf("%s is down \n", instance.Instance)
			db = database.MysqlInit()
			db.Model(moddle.FaultList{}).Create(&faultMsg)
			_ = db.Close()
			b = a.WhenError()
			if b == true {
				instance.Cancel()
			}
		}
	}
}

func callingRate(){
	for {
		Rate = <- Rating
	}
}

func StartCurlApp () {
	NumCounts.Add(len(viper.GetStringSlice("instances")))
	for _,v := range viper.GetStringSlice("instances"){
		instance = v
		url = fmt.Sprintf("http://prometheus.corp.taojinkf.com/api/v1/query?query=probe_http_status_code{instance=\"%s\"}", instance)
		go InitStatusChan().CurlApp(url)
	}
	WG.Done()
}