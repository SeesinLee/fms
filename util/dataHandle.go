package util

import (
	"fms/common"
	"fms/moddle"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HandleTimeData(date []moddle.Time)[]int{
	a := make([]int,0)
	e := strings.Split(strings.Split(time.Now().String()," ")[0],"-")
	for _,v := range date{
		b := strings.Split(strings.Split(v.String()," ")[0],"-")
		if b[0] != e[0]{
			continue
		}
		d,_ :=strconv.Atoi(b[1])
		a = append(a,d)
	}
	c := make(map[int]common.MonthSum)
	return HandleData(HandleCount(a,c))
}

func HandleCount(a []int,b map[int]common.MonthSum)map[int]common.MonthSum{
	if len(a) < 1 {
		return b
	}
	var left,right []int
	for _,v := range a[1:] {
		if v == a[0]{
			left =append(left,v)
		} else {
			right = append(right,v)
		}
	}
	left = append(left,a[0])
	b[left[0]] = common.MonthSum{
		Value: len(left),
		Month: left[0],
	}
	return HandleCount(right,b)
}

func HandleData(a map[int]common.MonthSum)[]int{
	d := common.MonthSum{}
	var b = []int{1,2,3,4,5,6,7,8,9,10,11,12}
	c := make([]int,0)
	for _,v := range b{
		if a[v] == d {
			continue
		}
		c = append(c,a[v].Value)
	}
	return c
}

func HandleType(a []string,b []common.JSKeyValueCommon,c common.JSKeyValueCommon)[]common.JSKeyValueCommon{
	if len(a) < 1 {
		return b
	}
	var left,right []string
	for _,v := range a[1:] {
		if v == a[0]{
			left =append(left,v)
		} else {
			right = append(right,v)
		}
	}
	left = append(left,a[0])
	c.Name = left[0]
	c.Value = len(left)
	b = append(b,c)
	return HandleType(right,b,c)
}

func HandleDuration(a map[moddle.Time]int)[]float64{
	b := map[int]common.AverageCompute{}
	e := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	f := make([]float64,0)
	g := common.AverageCompute{}
	h := strings.Split(strings.Split(time.Now().String(), " ")[0], "-")
	for n:= 1 ;n < 13 ;n ++ {
		c := common.AverageCompute{}
		for i, v := range a {
			p := strings.Split(strings.Split(i.String(), " ")[0], "-")
			if p[0] != h[0] {
				continue
			}
			d,_ := strconv.Atoi(p[1])
			if d != n{
				continue
			}
			c.Value = c.Value + v
			c.Denominator += 1
			b[n] = c
		}
	}
	for _,v := range e {
		if b[v] == g {
			continue
		}
		vl,_:=strconv.ParseFloat(fmt.Sprintf("%.2f",float64(b[v].Value/b[v].Denominator/60)),64)
		f = append(f,vl)
	}
	return f
}