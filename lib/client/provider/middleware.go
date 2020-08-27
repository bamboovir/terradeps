package provider


import (
	"fmt"
	resty "github.com/go-resty/resty/v2"
)

// MiddleWareSet defination
type MiddleWareSet struct {
	Map map[string]resty.RequestMiddleware
}

// GetMiddleware defination
func (m *MiddleWareSet) GetMiddleware(atom string) (resty.RequestMiddleware, bool){
	v, ok := m.Map[atom]
	return v, ok
}

// MustGetMiddleware defination
func (m *MiddleWareSet) MustGetMiddleware(atom string) (resty.RequestMiddleware) {
	v, ok := m.Map[atom]
	if !ok {
		panic(fmt.Sprintf("middleware [%s] is not exist", atom))
	}
	return v
}