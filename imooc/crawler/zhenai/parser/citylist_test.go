package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	const ressultSize = 470
	var expectedUrls = []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	result := ParserCityList(contents)
	if len(result.Requests) != ressultSize {
		t.Errorf("test failed, %d", len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Items[i] != url {
			t.Errorf("expected url #%d, %s", i, url)
		}
	}
}
