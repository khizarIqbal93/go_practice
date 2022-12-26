package crawl

import (
	"net/url"
	"testing"
)

type urlTest struct {
	urlString         string
	isRoot            bool
	expectedUrlString string
}

var urlTests = []urlTest{
	{"https://go.dev/doc/code#Testing", true, "https://go.dev/doc/code#Testing"},
	{"https://www.youtube.com/watch?v=IbFGG4T3_Yo", true, "https://www.youtube.com/watch?v=IbFGG4T3_Yo"},
	{"/mmcgrana/gobyexample", false, "https://github.com/mmcgrana/gobyexample"},
	{"/login?return_to=%2Fmmcgrana%2Fgobyexample", false, "https://github.com/login?return_to=%2Fmmcgrana%2Fgobyexample"},
}

func TestSetPageUrl(t *testing.T) {
	for _, test := range urlTests {
		page := Page{}
		if !test.isRoot {
			page.ParentUrl, _ = url.Parse("https://github.com")
		}

		page.SetPageUrl(test.urlString, test.isRoot)
		if page.Page != test.expectedUrlString {
			t.Errorf("Expected %q but got %q", test.expectedUrlString, page.Page)
		}
	}
}
