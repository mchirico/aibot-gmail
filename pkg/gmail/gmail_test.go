package gmail

import (
	"fmt"
	"github.com/mchirico/date/parse"
	_ "github.com/mchirico/date/parse"
	"github.com/mchirico/go-gmail/mail/messages"
	"testing"
	"time"
)

func TestStartWatch(t *testing.T) {
	expire, err := StartWatch()
	if err != nil {
		t.Fatalf("Can't start watch")
	}
	t.Logf("Expire: %v", expire)
	//StopWatch()
}

func TestStopWatch(t *testing.T) {
	StopWatch()
}

func TestReject(t *testing.T) {
	err := RejectImmediate("postmaster","a")
	if err == nil {
		t.Fatal()
	}

	err = RejectImmediate("mchirico@gmail.com","a")
	if err != nil {
		t.Fatal()
	}

}

func TestFb(t *testing.T) {
	Fb()
}

func TestEmailCountFB(t *testing.T) {
	count, _ := EmailCount("garbo3", "snippet")
	if count <= 0 {
		t.Fatalf("Fail")
	}
}

func TestDomains(t *testing.T) {
	Domains(14, "domains")
}

func TestDomainsReject(t *testing.T) {
	//Domains(1400,"domainsReject")
}

func TestRunEmail(t *testing.T) {
	RunEmail()
}

func TestSendReply(t *testing.T) {
	SendReply()
}

func TestMessages(t *testing.T) {

	r, _ := messages.GetNewMessages("SENT", 4)
	for _, m := range r {
		if date, ok := m["Date"]; ok {
			s := string(date)
			//layout := "Mon, _2 Jan 2006 15:04:05 -0700"
			tt, err := parse.DateTimeParse(s).GetTime()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			loc, err := time.LoadLocation("America/New_York")
			if err != nil {
				t.Fatal()
			}
			fmt.Printf("tt>: %v\n", tt.In(loc).Format("2006-01-02 15:04:05 pm"))
			fmt.Printf("tt: %v\n", time.Now().In(loc).Format("2006-01-02 15:04:05 pm"))

			fmt.Println(time.Now().Unix())
			fmt.Println(date)
		}

	}
}
