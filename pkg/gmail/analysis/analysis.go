package analysis

import "github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"

func Analysis() {
	sm := headertrack.NewSM()
	sm.GetR()
}
