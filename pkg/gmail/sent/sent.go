package sent

import (
	"errors"
	"fmt"
	"github.com/mchirico/date/parse"
	"github.com/mchirico/go-gmail/mail/messages"
	"time"
)

func GetDateFromMap(m map[string]string) (time.Time, time.Duration, string, error) {
	if date, ok := m["Date"]; ok {

		tdate, err := parse.DateTimeParse(date).GetTime()
		if err != nil {
			fmt.Println(date)
			fmt.Printf("err: %v\n", err)
			return time.Time{}, time.Now().Sub(time.Time{}), "", err
		}

		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			return time.Time{}, time.Now().Sub(time.Time{}), "", err
		}
		return tdate, time.Now().Sub(tdate), tdate.In(loc).Format("2006-01-02 15:04:05"), nil

	}
	return time.Time{}, time.Now().Sub(time.Time{}), "", errors.New("Can't extract time.Time")
}

func GetSnippetFromMap(m map[string]string) (string, error) {
	if snippet, ok := m["Snippet"]; ok {
		return snippet,nil
	}
	return "",errors.New("Snippet not found")
}

func GetTo(m map[string]string) (string, error) {
	if to, ok := m["To"]; ok {
		return to,nil
	}
	if to, ok := m["TO"]; ok {
		return to,nil
	}
	return "",errors.New("To not found")
}

func GetMsgId(m map[string]string) (string, error) {
	if v, ok := m["Message-Id"]; ok {
		return v,nil
	}
	if v, ok := m["Message-ID"]; ok {
		return v,nil
	}
	return "",errors.New("Message-id not found")
}

func GetSubject(m map[string]string) (string, error) {
	if v, ok := m["Subject"]; ok {
		return v,nil
	}
	return "",errors.New("Subject not found")
}

func GetAImsgField(m map[string]string) (string, error) {
	if v, ok := m["AI-Msg-Field"]; ok {
		return v,nil
	}
	return "",errors.New("AI-Msg-Field")
}

// SENT
func Msg(label string) {
	r, _ := messages.GetNewMessages(label, 40)
	for _, m := range r {
		_, _, sdate, err := GetDateFromMap(m)
		if err != nil {
			continue
		}

		snippet,err := GetSnippetFromMap(m)
		if err != nil {
			continue
		}
		subject,err := GetSubject(m)
		if err != nil {
			continue
		}
		to,err := GetTo(m)
		if err != nil {
			continue
		}
		msgId,err := GetMsgId(m)
		if err != nil {
			continue
		}
		msgAI,_ := GetAImsgField(m)

		if len(snippet) > 50 {
			fmt.Println(sdate,",",to,msgAI,subject, snippet[0:50],msgId[0:10])
		} else {
			fmt.Println(sdate, ",", to, msgAI, subject, snippet, msgId[0:10])
		}


	}
}
