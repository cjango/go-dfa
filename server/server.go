package server

import (
	"dfa/core"
	"encoding/json"
	"net/http"
)

func Start(port string, types []string) {
	http.HandleFunc("/check", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			return
		}
		_ = request.ParseForm()
		content := request.FormValue("content")

		var msg core.Message

		if result, ok := core.Check(content, types); !ok {
			msg = core.Message{
				Status:  500,
				Message: result,
			}
		} else {
			msg = core.Message{
				Status: 200,
			}
		}
		ret, _ := json.Marshal(msg)
		_, _ = writer.Write(ret)
	})
	_ = http.ListenAndServe(":"+port, nil)
}
