package bot

import (
	"regexp"
	"strings"
)

// Returns number of a release
func getNumber(s string) string {
	array := strings.Split(s, " ")
	return array[len(array)-1]
}

func parseNews(s string) (themes string) {
	r := regexp.MustCompile(`<li>(?:<a href=".*">)?(?P<news>.*?)(?:</a>\.)?</li>`)
	matches := r.FindAllStringSubmatch(s, -1)
	var news []string

	for _, m := range matches {
		news = append(news, m[1])
	}
	news[0] = "* " + news[0]
	return strings.Join(news, "\n* ")
}

