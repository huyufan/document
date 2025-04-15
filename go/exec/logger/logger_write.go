package logger

// The csyslog function is necessary here because cgo does not appear
// to be able to call a variadic function directly and syslog has the
// same signature as printf.

// #include <stdlib.h>
// #include <syslog.h>
// void csyslog(int p, const char *m) {
//     syslog(p, "%s", m);
// }
import "C"
import (
	"bytes"
	"errors"
)

const (
	NumMessages = 10 * 1024
)

type logMessage struct {
	bytes.Buffer
	level C.int
}

var (
	ErrLogFullBuf = errors.New("Log message queue is full")
)
