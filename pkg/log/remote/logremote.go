package remote

import (
	"github.com/mchirico/aibot-gmail/pkg/httputils"
	"log"
	"os"
)

func Log(email, value string) (string, error) {

	url := "https://tasks.cwxstat.io/gmail"

	key := "Authorization"
	KeyValue := os.Getenv("TASK_KEY")

	h := httputils.NewHTTP()
	h.Header(key, KeyValue)

	h.Header("Email", email)
	h.Header("Value", value)

	r, err := h.Get(url)
	if err != nil {
		log.Printf("log.remote.err: %s\n", err)
		return "", err
	}

	return string(r), nil

}
