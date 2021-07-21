package language

import "strings"

// parse_accept_language(
//   "en-US, fr-CA, fr-FR",  # the client's Accept-Language header, a string
//   ["fr-FR", "en-US"]      # the server's supported languages, a set of strings
// )
// returns: ["en-US", "fr-FR"]

// parse_accept_language("fr-CA, fr-FR", ["en-US", "fr-FR"])
// returns: ["fr-FR"]

// parse_accept_language("en-US", ["en-US", "fr-CA"])
// returns: ["en-US"]

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

func ParseAcceptLanguage(header string, accepted []string) []string {
	var r []string
	var regionTags []string
	if header == "" {
		return []string{}
	}

	if len(accepted) == 0 {
		return []string{}
	}

	h := strings.ReplaceAll(header, " ", "")
	l := strings.Split(h, ",")

	for _, language := range l {
		if !strings.Contains(language, "-") {
			regionTags = append(regionTags, language)
		}
	}

	for _, language := range l {
		for _, a := range accepted {
			s := strings.Split(a, "-")
			region := s[0]
			if strings.Compare(a, language) == 0 {
				if !contains(r, language) {
					r = append(r, language)
				}
			} else if contains(regionTags, region) {
				if !contains(r, a) {
					r = append(r, a)
				}
			}
		}
	}

	return r
}
