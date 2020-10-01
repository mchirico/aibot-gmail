package analysis

import (
	"context"
	"errors"
	"fmt"
	"github.com/mchirico/aibot-gmail/pkg/gmail/analysis/filters"
	"github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"
	gfb "github.com/mchirico/go-firebase/pkg/gofirebase"
	"strings"
	"time"
)

func Analysis(count int, label string) ([]map[string]string, error) {
	sm := headertrack.NewSM()
	lt := headertrack.LabelCount{}
	lt.Count = count
	lt.Label = label
	return sm.GetR(lt)
}

// TODO: Directory sensitive... ../../
func SetupFB(ctx context.Context) (*gfb.FB, error) {
	credentials := "../../../credentials/septapig-firebase-adminsdk.json"
	StorageBucket := "septapig.appspot.com"

	//defer cancel() // cancel when we are finished
	fb := &gfb.FB{Credentials: credentials, StorageBucket: StorageBucket}
	_, err := fb.CreateApp(ctx)
	return fb, err
}

func UpdateSummary() {

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
	doc["timeStamp"] = time.Now()
	doc["random"] = number

	var COL = "aibot"

	fb.WriteMap(ctx, doc, COL, "go-gofirebase-v4")
	fb.WriteMapCol2Doc2(ctx, doc, COL, "go-gofirebase-v4", "updates", "doc")

}

func FilterJunkString(s []string) ([]string, error) {
	valid := []string{}
	err := errors.New("not found")
	for _, e := range s {
		if filters.IgnoreEmail(e) {
			continue
		}
		if !filters.IsEmail(e) {
			continue
		}
		valid = append(valid, e)
		err = nil
	}
	return valid, err
}

func Filter2orMore(m map[string]interface{}) ([]string, error) {
	email := []string{}
	for k, v := range m {

		switch t := v.(type) {
		case string:
			continue
		case int:
			if t < 2 {
				continue
			}
		case int64:
			if t < 2 {
				continue
			}
		default:
			continue
		}

		if strings.Contains(k, "@") {
			fmt.Println(k, v)
			email = append(email, k)
		}

	}
	return email, nil
}

func AnalysisFirebase(t time.Time) (map[string]interface{}, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fb, err := SetupFB(ctx)
	if err != nil {
		return map[string]interface{}{}, err
	}
	var COL2 = "aibotEmailCount"

	resultFind, err := fb.Find(ctx, COL2, "timeStamp", ">=", t)
	return resultFind, err

}
