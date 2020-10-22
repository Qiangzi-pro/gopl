package gather

import (
	"io"
	selenium_demo "learngo/gopl.io/crawler/zhenai/gather/selenium_demo"
)

type CityListGather struct {}

func (c CityListGather) Fetch(url string) (io.Reader, error) {
	return selenium_demo.GetPageContent(), nil
}
