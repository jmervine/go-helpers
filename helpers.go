package helpers

// credit: https://stackoverflow.com/a/24809646

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// WARNING: If the app directory name changes, this will need to change as well.
var AppName string

// Any place that includes this package, should be able to set `helpers.DebugEnabled = true`
// to turn on debug messaging.
var DebugEnabled bool

func Error(err interface{}) error {
	e := assertError(err)
	return fmt.Errorf("%s", format(e.Error()))
}

func Errorf(t string, args ...any) error {
	return fmt.Errorf("%s", formatf(t, args...))
}

func PrintError(err interface{}) {
	e := assertError(err)
	log.Println(formatf("[error] in %v", e))
}

func PrintErrorf(t string, args ...any) {
	e := assertError(fmt.Sprintf(t, args...))
	log.Printf("[error] in %v\n", format(e.Error()))
}

func FatalError(err interface{}) {
	e := assertError(err)
	log.Fatalf("[error] in %v\n", format(e.Error()))
}

func FatalErrorf(t string, args ...any) {
	e := assertError(fmt.Sprintf(t, args...))
	log.Fatalf("[error] in %v\n", format(e.Error()))
}

func PrintDebugf(t string, args ...any) {
	if DebugEnabled {
		log.Printf("[debug] %s\n", fmt.Sprintf(t, args...))
	}
}

func PrintDebug(s string) {
	if DebugEnabled {
		log.Printf("[debug] %s\n", s)
	}
}

func PrintWarnf(t string, args ...any) {
	log.Printf("[warn] %s\n", fmt.Sprintf(t, args...))
}

func PrintWarn(s string) {
	log.Printf("[warn] %s\n", s)
}

func trunc(path string) string {
	parts := strings.Split(path, fmt.Sprintf("/%s/", AppName))
	if len(parts) == 2 {
		return parts[1]
	}

	// WARNING: This might be fagile, using some checking to be safe and return
	// the unmodified path, if it's not exactly what's expected.
	return path
}

func format(s string) string {
	pc, filename, line, _ := runtime.Caller(2)
	filename = trunc(filename)
	pcName := trunc(runtime.FuncForPC(pc).Name())
	return fmt.Sprintf("%s[%s:%d]: %s", pcName, filename, line, s)
}

func formatf(t string, args ...any) string {
	pc, filename, line, _ := runtime.Caller(2)
	filename = trunc(filename)
	pcName := trunc(runtime.FuncForPC(pc).Name())
	return fmt.Sprintf("%s[%s:%d]: %s", pcName, filename, line, fmt.Sprintf(t, args...))
}

func assertError(err interface{}) (e error) {
	switch v := err.(type) {
	case error:
		e = v
	case string:
		e = fmt.Errorf("%s", v)
	default:
		panic(fmt.Sprintf("'%v' is neither a string or an error", err))
	}
	return e
}
