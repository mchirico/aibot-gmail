package analysis

import "github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"

func Analysis() ([]map[string]string, error) {
	sm := headertrack.NewSM()
	lt := headertrack.LabelCount{}
	lt.Count = 10
	lt.Label = "SPAM"
	return sm.GetR(lt)
}
