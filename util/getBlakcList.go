package util

import (
	"fms/database"
)

func GetExist(token string)interface{}{
	re:= database.InitRedis()
	b,_ := re.Get().Do("exists",token)
	_=re.Get().Close()
	return b
}
