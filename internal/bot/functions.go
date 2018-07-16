package bot

import (
	"strings"

	"github.com/anaskhan96/soup"
)

// Returns number of a release
func getNumber(s string) string {
	array := strings.Split(s, " ")
	return array[len(array)-1]
}

func parseNews(s string) (topics string, err error) {
	root := soup.HTMLParse(s)
	if root.Error != nil {
		return "", root.Error
	}

	var news []string
	for _, topic := range root.FindAll("li") {
		if topic.Error != nil {
			continue
		}
		desc := topic.Find("a")
		if desc.Error == nil {
			news = append(news, desc.Text())
		}
	}

	news[0] = "* " + news[0]
	return strings.Join(news, "\n* "), nil
}
