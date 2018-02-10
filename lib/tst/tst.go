// Package tst contains common testing helpers
package tst

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"testing"
)

// GetWR returns a new ResponseRecorder and a new Request to pass to a Handler
func GetWR(method, target string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, target, body)
	return w, r
}

// ParseW returns the whole Response and the body string from given ResponseRecorder
func ParseW(w *httptest.ResponseRecorder) (*http.Response, string) {
	r := w.Result()
	body, _ := ioutil.ReadAll(r.Body)
	return r, string(body)
}

// These come from https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
// but then improved a bit

// Assert fails the test if the condition is false.
func Assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	assert(tb, condition, msg, v...)
}

// Ok fails the test if given err is not nil.
func Ok(tb testing.TB, err error) {
	assert(tb, err == nil, "unexpected error: %s", err)
}

// Eq fails the test if exp is not equal to act.
func Eq(tb testing.TB, exp, act interface{}) {
	assert(tb, reflect.DeepEqual(exp, act), "\n\n\texp: %#v\n\n\tgot: %#v", exp, act)
}

// Regexp fails the test if given value did not match given regexp pattern
// or if the pattern failed to compile
func Regexp(tb testing.TB, pattern string, value string) {
	matched, err := regexp.MatchString(pattern, value)
	assert(tb, err == nil, "error: %s\n\n compiling regexp: %s", err, pattern)
	assert(tb, matched, "\n\n\texpected pattern: %s\n\n\tgot: %s", pattern, value)
}

// The actual assert function, always reports 'runtime.Caller(2)' and beyond
// so it's meant to be called by the public assert functions
func assert(tb testing.TB, condition bool, msg string, v ...interface{}) {
	if !condition {
		backtrace := getBacktrace()
		fmt.Printf(backtrace+" \033[31m"+msg+"\033[39m\n\n", v...)
		tb.FailNow()
	}
}

func getBacktrace() string {
	return "\033[1;91m" + strings.Join(getCallers(), "") + "\033[00m"
}

func getCallers() []string {
	callers := []string{}
	ok := true
	var file string
	var line int
	for i := 4; ok; i++ {
		_, file, line, ok = runtime.Caller(i)
		callers = append(callers, fmt.Sprintf("\n%s:%d", filepath.Base(file), line))
	}
	return callers[:len(callers)-3]
}
