package thirdparty

import (
	"fmt"
	"github.com/mchirico/aibot-gmail/pkg/gmail"
	"github.com/mchirico/aibot-gmail/pkg/gmail/headertrack"
	"testing"
)

func Test(t *testing.T) {
	ht := headertrack.NewSM()
	r, err := gmail.GetMessage(ht)
	fmt.Println(r,err)


}
