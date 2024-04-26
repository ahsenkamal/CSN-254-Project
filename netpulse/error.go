package netpulse

import "fmt"

const (
	// Warnings
	warnInterface = 801
	warnClose     = 802
	warnDuplicate = 803
	// Errors
	errNew              = 901
	errDispatch         = 902
	errService          = 903
	errInitialize       = 904
	errStart            = 905
	errJoin             = 906
	errInterface        = 907
	errPort             = 908
	errNodeHeader       = 909
	errServiceHeader    = 910
	errVersionHeader    = 911
	errGroupHeader      = 912
	errVerbose          = 913
	errDispatchHeader   = 914
	errDispatchAction   = 915
	errScheme           = 916
	errResUnmarshal     = 917
	errResUnmarshalJSON = 918
	errUnknownService   = 919
	errTimeout          = 920
	errRECV             = 921
	errREPL             = 922
	errLogLevel         = 923
	errAdd              = 924
	errReqMarshal       = 925
	errReqUnmarshal     = 926
	errReqUnmarshalJSON = 927
	errReqUnmarshalHTTP = 928
	errReqWhisper       = 929
	errResWhisper       = 930
	errLeave            = 931
	errUnzip            = 932
	errUnzipRead        = 933
	errDo               = 934
	errClosed           = 935
	errWait             = 936
)

type Error struct {
	Codes   []int
	message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("netpulse: %s %v", e.message, e.Codes)
}

func (e *Error) escalate(code int) *Error {
	e.Codes = append(e.Codes, code)
	return e
}

func newError(code int, format string, v ...interface{}) *Error {
	return &Error{Codes: []int{code}, message: fmt.Sprintf(format, v...)}
}
