package singleton

//Singleton singleton object's interface
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

//GetInstance get singleton instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
