package tagreplace

import "testing"

func TestNewDefaultTagReplacer(t *testing.T) {
	src := `{{{{fee}}}}元{{{{day}}}}天`
	expect := `{{200}}元{{10}}天`
	rep := make(map[string]string)
	rep["fee"] = "200"
	rep["day"] = "10"
	tr := NewDefaultTagReplacer()
	dst := tr.ReplaceTag(src, rep)
	if dst != expect {
		t.Errorf("got:%s||expected:%s", dst, expect)
	}
}

func TestNewTagReplacer(t *testing.T) {
	src := `\\\\fee////元 \\day//天`
	expect := `\\200//元 10天`
	rep := make(map[string]string)
	rep["fee"] = "200"
	rep["day"] = "10"
	tr := NewTagReplacer(`\\`, `//`)
	dst := tr.ReplaceTag(src, rep)
	if dst != expect {
		t.Errorf("got:%s||expected:%s", dst, expect)
	}
}
