package simplex

import (
	"regexp"
	"strings"
	url2 "net/url"
)

type Router struct {
	routes []*RoutePath
}

type RoutePath struct {
	regex  *regexp.Regexp
	method string
	h     HandlerFunc
}

// NewRouter returns new router instance.
func NewRouter() *Router {
	r := Router{
		routes: make([]*RoutePath, 0),
	}
	return &r
}

func newRoute() *RoutePath {
	route := RoutePath{}

	return &route
}

func (rt *Router) MatchRoute(url string, method string) (*RoutePath, error) {
	url = url2.QueryEscape(url)
	if !strings.HasSuffix(url, "%2F") {
		url += "%2F"
	}
	url = strings.Replace(url, "%2F", "/", -1)
	for _, route := range rt.routes {
		matched, _ := regexp.MatchString(route.regex.String(), url)
		if matched {
			return route ,nil
		}
	}

	return nil, nil
}


func (rt *Router) Route(pattern string, method string, handler HandlerFunc) (*RoutePath) {
	route := newRoute()
	route.regex = parsePattern(pattern)
	route.method = method
	route.h = handler

	return route
}

func parsePattern(pattern string) (*regexp.Regexp)  {
	pattern = url2.QueryEscape(pattern)
	if !strings.HasSuffix(pattern, "%2F") {
		pattern += "%2F"
	}
	pattern = strings.Replace(pattern, "%2F", "/", -1)
	regex := regexp.MustCompile("^" + pattern + "$")
	return regex
}
