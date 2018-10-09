package main

import (
	"flag"
	"regexp"
	"math"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/pkg/errors"
	"log"
)

var (
	// Command line flags config.
	ignorePattern        string
	matchPattern         string
	searchLimit          int64
	matchDirectoriesOnly bool

	ignoreRegex             *regexp.Regexp
	matchRegex              *regexp.Regexp
	searchLimitCount        int64 = math.MaxInt64
	searchLimitReachedError       = errors.New("maximum file limit reached")
)

const defaultBasePath = "."
const defaultIgnorePattern = `^(\.git|\.hg|\.svn)$`

func init() {
	flag.StringVar(&ignorePattern, "i", "", "Ignore pattern")
	flag.StringVar(&matchPattern, "m", "", "Match file/folder pattern")
	flag.Int64Var(&searchLimit, "l", -1, "Max files limit")
	flag.BoolVar(&matchDirectoriesOnly, "d", false, "Directory search only")
}

func main() {
	flag.Parse()

	if matchPattern != "" {
		matchRegex = regexp.MustCompile(matchPattern) // will panic
	}

	if ignorePattern != "" {
		ignoreRegex = regexp.MustCompile(ignorePattern)
	} else {
		ignoreRegex = regexp.MustCompile(defaultIgnorePattern)
	}

	if searchLimit > 0 {
		searchLimitCount = searchLimit
	}

	basePath := defaultBasePath
	if flag.NArg() > 0 {
		basePath = filepath.FromSlash(flag.Arg(0))
	}

	results := performSearch(basePath)

	for result := range results {
		if _, err := os.Stdout.Write([]byte(result + "\n")); err != nil {
			os.Exit(2)
		}
	}

}

func performSearch(basePath string) chan string {
	q := make(chan string, 32)

	go func() {
		n := int64(0)
		basePath = appendPathSep(basePath)

		performMatch := func(path string, info os.FileInfo) error {
			if matchRegex != nil && !matchRegex.MatchString(info.Name()) {
				return nil
			}

			n++
			if n > searchLimitCount {
				return searchLimitReachedError
			}

			q <- filepath.Clean(path)
			return nil
		}

		var err error
		err = filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
			if info == nil {
				return err
			}
			name := info.Name()
			if matchDirectoriesOnly {
				if info.IsDir() && name != defaultBasePath {
					if ignoreRegex.MatchString(name) {
						return filepath.SkipDir
					}
					return performMatch(path, info)
				}
				return nil
			} else {
				if !info.IsDir() {
					if ignoreRegex.MatchString(name) {
						return nil
					}
					return performMatch(path, info)
				} else {
					if ignoreRegex.MatchString(name) {
						return filepath.SkipDir
					}
				}
			}

			return nil
		})

		if err != nil && err != searchLimitReachedError {
			log.Fatal(err)
		}

		close(q)
	}()

	return q
}

func appendPathSep(basePath string) string {
	sep := string(os.PathSeparator)

	if !strings.HasSuffix(basePath, sep) {
		return fmt.Sprintf("%s%s", basePath, sep)
	}

	return basePath
}
