package tagreplace

import (
	"bytes"
	"regexp"
	"strings"
)

// TagReplacer ...
type TagReplacer interface {
	ReplaceTag(src string, rep map[string]string) string
}

type tagReplace struct {
	preFix     string
	posFix     string
	compileExp string
	reg        *regexp.Regexp
}

var (
	defaultPrefix  = "{{"
	defaultPostfix = "}}"
	tagExpr        = `([\w\s\d]+)`
	escapeBytes    = []byte{'\\', '|', '^', '$', '.', '*', '+', '?', '{', '}', '(', ')', '[', ']'}
)

// NewDefaultTagReplacer ...
func NewDefaultTagReplacer() TagReplacer {
	return newTagReplace(defaultPrefix, defaultPostfix)
}

// NewTagReplacer ...
func NewTagReplacer(preFix, posFix string) TagReplacer {
	return newTagReplace(preFix, posFix)
}

func needEscape(b byte) bool {
	for _, esByte := range escapeBytes {
		if esByte == b {
			return true
		}
	}
	return false
}

func genCompileExp(preFix, posFix string) string {
	buf := bytes.NewBuffer(nil)
	for i := 0; i < len(preFix); i++ {
		if needEscape(preFix[i]) {
			buf.WriteByte('\\')
		}
		buf.WriteByte(preFix[i])
	}
	buf.WriteString(tagExpr)
	for i := 0; i < len(posFix); i++ {
		if needEscape(posFix[i]) {
			buf.WriteByte('\\')
		}
		buf.WriteByte(posFix[i])
	}
	return buf.String()
}

func newTagReplace(preFix, posFix string) TagReplacer {
	tr := &tagReplace{
		preFix:     preFix,
		posFix:     posFix,
		compileExp: genCompileExp(preFix, posFix),
	}
	tr.reg = regexp.MustCompile(tr.compileExp)
	return tr
}

func (t *tagReplace) replace(src string, rep map[string]string) string {
	src = t.reg.ReplaceAllStringFunc(src, func(i string) string {
		key := strings.TrimLeft(i, t.preFix)
		key = strings.TrimRight(key, t.posFix)
		if value, ok := rep[key]; ok {
			return value
		}
		return i
	})
	return src
}

func (t *tagReplace) ReplaceTag(src string, rep map[string]string) string {
	return t.replace(src, rep)
}
