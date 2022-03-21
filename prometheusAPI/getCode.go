package prometheusAPI

import (
	"context"
	"encoding/json"
	"fms/f3CallingRequest"
	"fms/moddle"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)
type GetCodeInter interface {
	CurlApp(url string)
	WhenError()bool
}

type InstanceIsDown struct {		//应用宕机的时候，保留curl实例的信息
	InstanceClass chan GetCodeInter
	Instance chan AppStatusMg
}

var	InstanceInfo = InstanceIsDown{
	InstanceClass: make(chan GetCodeInter),
	Instance: make(chan AppStatusMg),
}

type AppStatusMg struct {
	Instance string
	Url string
	Ctx context.Context
	Cancel context.CancelFunc
}

var Nums = make(chan int,1)
var Rating = make(chan float64)
var Lock sync.Mutex
var pm = new(moddle.PrometheusMetrics)
var a  = new(f3CallingRequest.F3Calling)
var repF3 = new(http.Response)

var (
	f3Calling int
	errF3     error
	bodyF3    []byte
	vb     f3CallingRequest.Result
	b    int
	c float64
)

func InitStatusChan() GetCodeInter{
	ctx,cancel :=context.WithCancel(context.Background())
	Si := &AppStatusMg{
		Ctx: ctx,
		Cancel: cancel,
	}
	return Si
}

func (si *AppStatusMg)CurlApp(url string){  //curlPrometheus的应用状态接口
	si.Url = url
			rep,err := http.Get(url)
			if err != nil {
				logrus.Error(err)
				NumCounts.Done()
				return
			}
			defer rep.Body.Close()
			body, err := ioutil.ReadAll(rep.Body)
			if len(body) == 0 {
				NumCounts.Done()
				return
			}
			if err != nil {
				logrus.Error(err)
				NumCounts.Done()
				return
			}
			err = json.Unmarshal(body, &pm)
			if err != nil {
				logrus.Error(err)
				NumCounts.Done()
				return
			}
			if pm.Status != "success"{
				NumCounts.Done()
				return
			}
			for _, v := range *pm.Data.Result {
				if v.Value[1] != "200" || len(v.Value) != 2{
					si.Instance = v.Metric.Instance
					Lock.Lock()
					InstanceInfo.Instance <- *si
					InstanceInfo.InstanceClass <- si
					Lock.Unlock()
					ONE:
					for {
						select {
						case <-si.Ctx.Done():
							break ONE
						}
					}
				}
			}
	NumCounts.Done()
}

func (si *AppStatusMg)WhenError()bool { //出现应用宕机情况后等待应用恢复
	Pm := new(moddle.PrometheusMetrics)
	rep := new(http.Response)
	var (
		err  error
		body []byte
		v moddle.Result
		i int
	)
	ONE:
	for i = 0; i < 5; i++ {
		time.Sleep(time.Second * 5)
		rep,err = http.Get(si.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer rep.Body.Close()
		body, err = ioutil.ReadAll(rep.Body)
		if len(body) == 0 {
			continue
		}
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &Pm)
		if err != nil {
			continue
		}
		for _, v = range *Pm.Data.Result {
			if v.Value[1] != "200" {
				break ONE
			}
			return true
		}
	}
	return si.WhenError()
}

func F3CallingRating(){
		Nums <- f3Calling
		repF3, errF3 = http.Get("http://prometheus.corp.taojinkf.com/api/v1/query?query=fs_call_num{name=\"fs3\"}")
		if errF3 != nil {
			fmt.Println(errF3)
			WG.Done()
			return
		}
		defer repF3.Body.Close()
		bodyF3, errF3 = ioutil.ReadAll(repF3.Body)
		if len(bodyF3) == 0 {
			WG.Done()
			return
		}
		if errF3 != nil {
			panic(errF3)
		}
		errF3 = json.Unmarshal(bodyF3,&a)
		if errF3 != nil {
			WG.Done()
			return
		}
		if a.Status != "success" {
			WG.Done()
			return
		}
		for _, vb = range *a.Data.Result {
			f3Calling,_ =strconv.Atoi(vb.Value[1].(string))
		}
		b = <- Nums
		c = float64(b) / float64(f3Calling)
		Rating <- c
		WG.Done()
}