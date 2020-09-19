package headertrack

import (
	"fmt"
)

type Details struct {
	Subject    map[string]int
	Snippet    map[string]int
	MessageID  map[string]int
	From       map[string]int
	ReturnPath map[string]int
	Id         map[string]int
	HistoryId  map[int]int
	AIMsgField map[string]int
	Count      int
	Matches    int
	State      string
	Index      map[string]string
}

func CreateIndex(headers map[string]string) (*Details, string) {
	details := &Details{
		Subject:    map[string]int{},
		Snippet:    map[string]int{},
		MessageID:  map[string]int{},
		From:       map[string]int{},
		ReturnPath: map[string]int{},
		Id:         map[string]int{},
		HistoryId:  map[int]int{},
		AIMsgField: map[string]int{},
		Count:      0,
		Matches:    0,
		State:      "",
		Index:      map[string]string{},
	}
	s := "Init:"
	if v, ok := headers["Snippet"]; ok {
		details.Snippet[v] += 1
		fmt.Printf("\nSnippet:\n\n->%s<-\n\n", v)
		s += v
	}
	if v, ok := headers["Message-ID"]; ok {
		details.MessageID[v] += 1
		s += v
	}
	if v, ok := headers["Return-Path"]; ok {
		details.ReturnPath[v] += 1
		s += v
	}
	if v, ok := headers["AI-Msg-Field"]; ok {
		details.AIMsgField[v] += 1
		s += v
	}
	details.Index[s] = "init"
	fmt.Printf("\n\n\ns == ->%s<-\n", s)
	return details, s
}

type SM struct {
	SentMail map[string]Details
}

func NewSM() *SM {
	return &SM{map[string]Details{}}
}

func (s *SM) Found(headers map[string]string) bool {
	details, idx := CreateIndex(headers)
	if _, ok := s.SentMail[idx]; ok {
		return true
	}
	s.SentMail[idx] = *details
	return false

}
