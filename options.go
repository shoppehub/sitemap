package sitemap

import (
	"os"
	"path"
)

const (
	// MaxSitemapLinks defines max links per Sitemap
	MaxSitemapLinks = 50000
)

type options struct {
	defaultHost string
	publicPath  string
	filename    string
	compress    bool
	pretty      bool
	maxLinks    int
}

func NewOptions() *options {
	pwd, _ := os.Getwd()
	return &options{
		defaultHost: "http://www.example.com",
		publicPath:  pwd,
		filename:    "Sitemap.xml",
		compress:    false,
		pretty:      false,
		maxLinks:    MaxSitemapLinks,
	}
}

func (o *options) SetDefaultHost(host string) {
	o.defaultHost = host
}

func (o *options) SetPublicPath(path string) {
	o.publicPath = path
}

func (o *options) SetFilename(filename string) {
	if path.Ext(filename) != ".xml" {
		filename = filename + ".xml"
	}
	o.filename = filename
}

func (o *options) SetCompress(compress bool) {
	o.compress = compress
}

func (o *options) SetPretty(pretty bool) {
	o.pretty = pretty
}

func (o *options) SetMaxLinks(max int) {
	if max < MaxSitemapLinks && max > 0 {
		o.maxLinks = max
	}
}
