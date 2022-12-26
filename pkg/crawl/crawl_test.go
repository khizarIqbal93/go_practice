package crawl

import "testing"

type urlTest struct {
	urlString         string
	isRoot            bool
	expectedUrlString string
}

var urlTests = []urlTest{
	urlTest{"https://go.dev/doc/code#Testing", true, "https://go.dev/doc/code#Testing"},
	urlTest{"https://www.youtube.com/watch?v=IbFGG4T3_Yo", true, "https://www.youtube.com/watch?v=IbFGG4T3_Yo"},
}

func TestSetPageUrl(t *testing.T) {
	for _, test := range urlTests {
		page := Page{}
		page.SetPageUrl(test.urlString, test.isRoot)
		if page.Page != test.expectedUrlString {
			t.Errorf("Output not equal to expected %q", test.expectedUrlString)
		}
	}
}
