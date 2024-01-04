package link

import (
	"fmt"
	"net/url"
	"strings"
)

type Link struct {
	URL      *url.URL
	Hostname string
}

func IsUrl(str string) bool {
	url, err := url.Parse(str)
	if err != nil {
		return false
	}
	if url.Scheme == "" || url.Host == "" {
		return false
	}
	return true
}

func getHostname(hostname string) string {
	// Check if the hostname has www. in it and if so remove it
	substring := strings.Split(hostname, ".")

	if substring[0] == "www" {
		hostname = strings.Join(substring[1:], ".")
	}

	return hostname
}

func (l *Link) IsFixableUrl() bool {
	l.Hostname = getHostname(l.URL.Hostname())
	switch l.Hostname {
	case "instagram.com", "twitter.com", "x.com", "reddit.com", "tiktok.com":
		return true
	}
	return false
}

func (l *Link) Fix() string {
	return fmt.Sprintf(fixURL(l))
}

func fixURL(link *Link) string {
	switch link.Hostname {
	case "instagram.com":
		return fixInstagram(link)
	case "twitter.com":
		return fixTwitter(link)
	case "x.com":
		return fixX(link)
	case "reddit.com":
		return fixReddit(link)
	case "tiktok.com":
		return fixTikTok(link)

	}
	return ""
}

func fixInstagram(link *Link) string {
	newURL := link.URL.Scheme + "://dd" + link.Hostname + link.URL.Path
	return newURL
}

func fixTwitter(link *Link) string {
	newURL := link.URL.Scheme + "://fx" + link.Hostname + link.URL.Path
	return newURL
}

func fixX(link *Link) string {
	newURL := link.URL.Scheme + "://fixv" + link.Hostname + link.URL.Path
	return newURL
}

func fixReddit(link *Link) string {
	newURL := link.URL.Scheme + "://vx" + link.Hostname + link.URL.Path
	return newURL
}

func fixTikTok(link *Link) string {
	newURL := link.URL.Scheme + "://vx" + link.Hostname + link.URL.Path
	return newURL
}
