package gmail

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mchirico/aibot-gmail/pkg/gmail/canned/replies"
	"github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"
	gfb "github.com/mchirico/go-firebase/pkg/gofirebase"
	"github.com/mchirico/go-gmail/mail/messages"
	"github.com/mchirico/go-pubsub/pubsub"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func SetupFB(ctx context.Context) (*gfb.FB, error) {
	credentials := "../../credentials/septapig-firebase-adminsdk.json"
	StorageBucket := "septapig.appspot.com"

	//defer cancel() // cancel when we are finished
	fb := &gfb.FB{Credentials: credentials, StorageBucket: StorageBucket}
	_, err := fb.CreateApp(ctx)
	return fb, err
}

func Lables() {
	messages.Labels()
}

func Domains(number_to_check int, doc string) (map[string]interface{}, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fb, err := SetupFB(ctx)

	var COL = "aibot"
	dsnap, _ := fb.ReadMap(ctx, COL, "domains")
	mi := dsnap.Data()

	if err != nil {
		fmt.Errorf("Can't setup Firebase")
	}

	r, err := messages.GetNewMessages("TRASH", number_to_check)
	if err != nil {
		return map[string]interface{}{}, err
	}
	domainsT := messages.Domains(r)

	r, err = messages.GetNewMessages("SPAM", number_to_check)
	if err != nil {
		return map[string]interface{}{}, err
	}
	domainsS := messages.Domains(r)

	for k, v := range domainsT {
		mi[k] = v
	}
	for k, v := range domainsS {
		mi[k] = v
	}

	fb.WriteMap(ctx, mi, COL, doc)
	return mi, nil

}

func Fb() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fb, err := SetupFB(ctx)
	if err != nil {
		return
	}

	number := 9
	doc := make(map[string]interface{})
	doc["application"] = "FirebaseGo"
	doc["function"] = "TestAuthenticate"
	doc["test"] = "This is example text..."
	doc["random"] = number

	var COL = "aibot"

	fb.WriteMap(ctx, doc, COL, "go-gofirebase-v4")
	fb.WriteMapCol2Doc2(ctx, doc, COL, "go-gofirebase-v4", "updates", "doc")

	resultFind, _ := fb.Find(ctx, COL, "function", "==", "TestAuthenticate")

	if resultFind["test"] != "This is example text..." {
		fmt.Errorf("Find not working")
	}

	dsnap, _ := fb.ReadMap(ctx, COL, "go-gofirebase-v4")
	result := dsnap.Data()

	fmt.Printf("Document data: %v %v\n", result["random"].(int64), number)

}

func RejectEmails(from string) error {
	return errors.New("error")
}

func RejectImmediate(from, snippet string) error {
	f := strings.ToLower(from)
	s := strings.ToLower(snippet)

	rejectFrom := []string{"postmaster", "alert", "reply", "human"}
	rejectText := []string{"w2", "n.j."}

	for _, v := range rejectFrom {
		if strings.Contains(f, v) {
			return errors.New("Postmaster")
		}
	}

	for _, v := range rejectText {
		if strings.Contains(s, v) {
			return errors.New("Bot")
		}
	}
	return nil
}

func EmailCount(from, snippet string) (int64, error) {

	err := RejectImmediate(from, snippet)
	if err != nil {
		return 0, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fb, err := SetupFB(ctx)
	if err != nil {
		return 0, err
	}
	var COL = "aibotEmailCount"

	dsnap, _ := fb.ReadMap(ctx, COL, from)
	result := dsnap.Data()
	var count int64
	if v, ok := result[from]; ok {
		count = v.(int64)
	}

	sarray := []interface{}{}
	if v, ok := result["snippet"]; ok {
		sarray = v.([]interface{})
	}

	doc := make(map[string]interface{})
	doc[from] = 1 + count
	doc["timeStamp"] = time.Now()
	sarray = append(sarray, snippet)
	doc["snippet"] = sarray

	fb.WriteMap(ctx, doc, COL, from)
	return count, nil

}

func StartWatch() (time.Time, error) {

	// FIXME: This won't work in docker..
	//    Use file find
	b, err := ioutil.ReadFile("../../credentials/topic_name.json")
	if err != nil {
		return time.Time{}, err
	}
	return messages.StartWatch("me", string(b))

}

var HT = headertrack.NewSM()

func GetR() ([]map[string]string, error) {
	r, err := messages.GetNewMessages("TRASH", 1)
	if err != nil {
		return []map[string]string{}, err
	}
	id := 0
	fmt.Println("Subject:", r[id]["Subject"])
	fmt.Println("MessageID:", r[id]["Message-ID"])
	fmt.Println(r[id]["Return-Path"])
	fmt.Println(r[id]["From"])
	fmt.Println(r[id]["Snippet"])
	fmt.Println("--->", r[id]["Id"])

	if HT.Found(r[id]) {
		fmt.Printf("\n\nAborted: Found\n\n")
		return r, errors.New("Id previously used")
	}

	return r, nil
}

func EmailEnough(r []map[string]string) bool {
	count, err := EmailCount(r[0]["From"], r[0]["Snippet"])
	if err != nil {
		return true
	}
	if count >= 2 {
		return true
	}
	return false
}

// Maybe add analytics here
func PostEmailEnough(r []map[string]string) {
	log.Println("PostEmailEnough")
}

func SendReply() {
	r, err := GetR()
	if err != nil {
		return
	}

	if EmailEnough(r) {
		PostEmailEnough(r)
		return
	}

	id := 0

	rMsg := &replies.Replies{}
	msg := rMsg.Get(0)

	subject := "Contract?...  Re: " + r[id]["Subject"]
	if v, ok := r[id]["Subject"]; ok {
		if strings.Index(v, "Re: Contract?...") == 0 {
			msg = rMsg.Get(1)
			subject = "Final Question"
		}
		if strings.Contains(v, "Final") {
			msg = rMsg.Get(2)
			subject = "Future Reference"
			fmt.Printf("\nContains Final\n")
		}
		if strings.Contains(v, "Future Reference") {
			fmt.Printf("DONE. Last msg sent")
			return
		}
	}

	fmt.Printf("\n\nSENDING!!!!\n\n")
	msgID := r[id]["Message-ID"]
	_, err = messages.ReplyAI(r[id]["Id"], msgID, "mc@cwxstat.com",
		r[id]["From"], subject, msg, "contract")
	if err != nil {
		log.Printf("messages.ReplyAI err: %v\n"+
			"r[id]{Id}: %v\n"+
			"msgID: %v\n"+
			"Subject: %v\n"+
			"From: %v\n", err, r[id]["Id"], msgID, subject, r[id]["From"])

		messages.Send2(r[id]["From"], "Contract C2C, 100% Remote", rMsg.Get(2))
	}

}

func StopWatch() error {
	return messages.StopWatch("me")
}

type Pub struct {
	Email     string `json:"emailAddress"`
	HistoryId int    `json:"historyId"`
}

var COUNT int64 = 0

func RunEmail() {

	g := pubsub.NewG()
	var buf bytes.Buffer
	m := map[int]bool{}
	pub := &Pub{}
	for i := 0; i < 4; i++ {

		COUNT += 1
		msg, n, err := g.PullMsgsTimeOut(&buf, "gmail-sub", 10)
		if err != nil {
			log.Printf("No message")
		}
		if n == 0 {
			log.Printf("Looping. n: %v, COUNT: %v\n", n, COUNT)
			continue
		}
		fmt.Printf("msg: %s\n", msg)

		json.Unmarshal([]byte(msg), &pub)

		if _, ok := m[pub.HistoryId]; !ok {
			m[pub.HistoryId] = true
			fmt.Printf("Digest: %d\n", pub.HistoryId)
			SendReply()
		} else {
			fmt.Printf("Skipped: %d\n", pub.HistoryId)
		}

	}
}
