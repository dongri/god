package GoBeer

import (
	"strings"
	"time"
)

// Format ...
func Format(t time.Time, layout string) string {
	pattern := map[string]string{
		"%a": "Mon",
		"%A": "Monday",
		"%b": "Jan",
		"%B": "January",
		"%d": "02",
		"%e": "2",
		"%m": "01",
		"%y": "06",
		"%Y": "2006",
		"%H": "15",
		"%I": "03",
		"%l": "3",
		"%M": "04",
		"%P": "pm",
		"%p": "PM",
		"%S": "05",
		"%Z": "MST",
		"%%": "%",
	}
	for k, v := range pattern {
		layout = strings.Replace(layout, k, v, -1)
	}
	return t.Format(layout)
}
