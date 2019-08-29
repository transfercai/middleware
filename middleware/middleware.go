package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"transfercai/middleware/error"
)

func MakeHandleFunc(handle func(r *http.Request) (interface{}, *error.SError), ch chan int) func(w http.ResponseWriter, r *http.Request) {
	retJson := map[string]string{
		"return_code":    "500",
		"return_message": "connection is over limit",
	}
	retStr, err := json.Marshal(retJson)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal ret str:%s", err.Error()))
	}
	ret := func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ch:
			MiddleWare(w, r, handle)
		default:
			w.Write([]byte(retStr))
		}
	}
	return ret
}

func MiddleWare(w http.ResponseWriter, r *http.Request, handle func(r *http.Request) (interface{}, *error.SError)) {
	str := []byte("")
	ret, err := handle(r)
	if err != nil {
		str = FormatRetError(err)
		w.Write(str)
	} else {
		str = FormatRetJson(ret)
		w.Write(str)
	}
}

func FormatRetError(sError *error.SError) []byte {
	mp := map[string]interface{}{
		"return_code":    sError.ErrorCode,
		"return_message": sError.Message,
	}
	ret, err := json.Marshal(mp)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal ret str:%s", err.Error()))
	}
	return ret
}

func FormatRetJson(ret interface{}) []byte {
	retStr, err := json.Marshal(ret)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal ret str:%s", err.Error()))
	}
	return retStr
}

func LimitReq(limit int64, ch chan int) {
	tc := time.NewTicker(time.Millisecond*time.Duration(limit))
	for {
		<-tc.C
		ch <- 1
	}
}
