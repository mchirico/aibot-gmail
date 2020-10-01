package analysis

import "github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"

func Analysis(count int, label string) ([]map[string]string, error) {
	sm := headertrack.NewSM()
	lt := headertrack.LabelCount{}
	lt.Count = count
	lt.Label = label
	return sm.GetR(lt)
}
