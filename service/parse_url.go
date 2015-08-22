package pezdispenser

import (
	"net/url"
	"strings"

	"labix.org/v2/mgo"
)

//ParseURL - copy from latest version of mgo, but there was a version conflict,
//need to figure that out
func ParseURL(url string) (*mgo.DialInfo, error) {
	uinfo, _ := extractURL(url)
	direct := false
	mechanism := ""
	service := ""
	source := ""

	info := mgo.DialInfo{
		Addrs:     uinfo.addrs,
		Direct:    direct,
		Database:  uinfo.db,
		Username:  uinfo.user,
		Password:  uinfo.pass,
		Mechanism: mechanism,
		Service:   service,
		Source:    source,
	}
	return &info, nil
}

type urlInfo struct {
	addrs   []string
	user    string
	pass    string
	db      string
	options map[string]string
}

func extractURL(s string) (*urlInfo, error) {
	if strings.HasPrefix(s, "mongodb://") {
		s = s[10:]
	}
	info := &urlInfo{options: make(map[string]string)}

	if c := strings.Index(s, "@"); c != -1 {
		pair := strings.SplitN(s[:c], ":", 2)

		info.user, _ = url.QueryUnescape(pair[0])
		if len(pair) > 1 {
			info.pass, _ = url.QueryUnescape(pair[1])

		}
		s = s[c+1:]
	}
	if c := strings.Index(s, "/"); c != -1 {
		info.db = s[c+1:]
		s = s[:c]
	}
	info.addrs = strings.Split(s, ",")
	return info, nil
}
