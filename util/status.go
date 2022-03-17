package util

var (
	StatusChan  Status
	a int
)

type Status struct {
	ID int
	UserStatus chan int
}

func NewChan(){
	StatusChan.UserStatus = make(chan int,1)
}

func (s Status)Set(i int){
	s.UserStatus <- i
}

func (s Status)Watch() int {
	a = <- s.UserStatus
	return a
}

func (s Status)Clean(){
	s = Status{}
}

