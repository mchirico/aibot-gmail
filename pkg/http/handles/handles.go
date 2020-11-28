package handles

import (
	"encoding/json"
	"fmt"
	//"github.com/mchirico/aibot-gmail/pkg/etcd"
	"github.com/mchirico/aibot-gmail/pkg/gmail"
	"github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"
	"github.com/mchirico/go-gmail/mail/messages"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var Count = 0
var CountStatus = 0

func BaseRoot(w http.ResponseWriter, r *http.Request) {

	version := "v0.0.1"

	switch r.Method {
	case "GET":
		Count += 1

		msg := fmt.Sprintf("\nversion: %v\naibot: %v\n", version, Count)
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}

func Status(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		//etcd.D()
		CountStatus += 1

		//msg := fmt.Sprintf("\nstatus: %v\naibot: %v\n", Count, CountStatus)
		msg := fmt.Sprintf("COUNT: %v\n", gmail.COUNT)
		w.Write([]byte(msg))
	case "POST":
		// msg := fmt.Sprintf("Hello world: POST: %v", r.FormValue("user"))
		w.Write([]byte("post"))
	default:
		w.Write([]byte(`"Sorry, only GET and POST methods are supported."`))
	}

}

var PUBSUBTOKEN = os.Getenv("PUBSUBTOKEN")

type PubSubMessage struct {
	Message struct {
		Data []byte `json:"data,omitempty"`
		ID   string `json:"id"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

/*
Subscript gets called from PubSub. The PUBSUBTOKEN must
match.
*/
func Subscript(w http.ResponseWriter, r *http.Request) {

	key := fmt.Sprintf("%s", r.FormValue("key"))
	if key != PUBSUBTOKEN {
		w.Write([]byte("bad key"))
		return
	}
	w.Write([]byte("checking length"))
	if len(key) > 7 {

		var m PubSubMessage
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("ioutil.ReadAll: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &m); err != nil {
			log.Printf("json.Unmarshal: %v", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		data := string(m.Message.Data)
		log.Printf("Running %s!", data)
		sm := headertrack.NewSM()
		r, err := gmail.GetMessage(sm)
		if err != nil {
			lpmsg := gmail.LOOPMSG{}
			lpmsg.Send1 = messages.ReplyAI
			lpmsg.Send2 = messages.Send2
			lpmsg.LoopMsg(r)

		}

	}

}
