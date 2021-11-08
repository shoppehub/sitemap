package sitemap

import (
	"encoding/xml"
	"fmt"
	"time"
)

type ChangeFreq string

type InvalidPriorityError struct {
	msg string
}

func (e *InvalidPriorityError) Error() string {
	return e.msg
}

const (
	Always  ChangeFreq = "always"
	Hourly  ChangeFreq = "hourly"
	Daily   ChangeFreq = "daily"
	Weekly  ChangeFreq = "weekly"
	Monthly ChangeFreq = "monthly"
	Yearly  ChangeFreq = "yearly"
	Never   ChangeFreq = "never"
)

type po float64

func (f po) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := fmt.Sprintf("%.6f", f)
	return xml.Attr{Name: name, Value: s}, nil
}

type url struct {
	*base
	XMLName    xml.Name   `xml:"url"`
	Loc        string     `xml:"loc"`
	LastMod    string     `xml:"lastmod,omitempty"`
	ChangeFreq ChangeFreq `xml:"changefreq,omitempty"`
	Priority   po         `xml:"priority,omitempty"`
	Token      []xml.Token
}

func NewUrl() *url {
	return &url{
		base:       &base{},
		Loc:        "",
		LastMod:    "",
		ChangeFreq: "",
		Priority:   0,
	}
}

// 网址
func (u *url) SetLoc(loc string) *url {
	u.Loc = loc
	return u
}

// 最后一次修改时间
func (u *url) SetLastMod(lastMod time.Time, layout string) *url {
	if layout == "" {
		layout = "2006-01-02"
	}
	u.LastMod = lastMod.Format(layout)
	return u
}

// 更新频率
func (u *url) SetChangeFreq(freq ChangeFreq) *url {
	u.ChangeFreq = freq
	return u
}

// 网页优先级
func (u *url) SetPriority(priority float64) *url {
	if priority < 0 || priority > 1 {
		panic(InvalidPriorityError{"Valid values range from 0.0 to 1.0"})
	}
	u.Priority = po(priority)
	return u
}

// 对于单个网页上的多个视频，为该网页创建一个 <loc> 标记，并为该网页上的每个视频创建一个子级 <video> 元素。
func (u *url) AppendVideo(video *video) {
	u.setNs(VideoXmlNS)
	u.Token = append(u.Token, video)
}

// 对于单个网页上的多个图片，每个 <url> 标记最多可包含 1000 个 <image:image> 标记。
func (u *url) AppendImage(image *image) {
	u.setNs(ImageXmlNS)
	u.Token = append(u.Token, image)
}

func (u *url) AppendNews(news *news) {
	u.setNs(NewsXmlNS)
	u.Token = append(u.Token, news)
}
