package filters

import "strings"

func IgnoreEmail(s string) bool {
	found := false
	ignore_words := []string{"job-search", "alert", "indeedapply", "accumatchconsulting.com",
		"enterprisesolutioninc.com", "huxley.com", "datarobot.com", "jobs", "facebook", "security",
		"candidates", "Mohammad", "singh", "thetechhires.com", "monster.com", "themuse.com"}
	for _, v := range ignore_words {
		if strings.Contains(strings.ToLower(s), strings.ToLower(v)) {
			return true
		}
	}
	return found
}

func IsEmail(s string) bool {
	count := 0
	emailValid := []string{"@", ".com", ".io"}
	for _, v := range emailValid {
		if strings.Contains(strings.ToLower(s), strings.ToLower(v)) {
			count += 1
		}
	}
	if count > 1 {
		return true
	}
	return false
}
