package image

import (
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	encimg "github.com/ikeikeikeike/gopkg/encoding/image"
	behavior "github.com/ikeikeikeike/gopkg/net/http"
	"github.com/ikeikeikeike/gopkg/str"
)

func ExtractImagePaths(str string) (paths []string, err error) {
	reader := strings.NewReader(str)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("src")
		paths = append(paths, val)
	})
	return
}

type FileInfo struct {
	Filename string
	Url      string
	Ext      string
	Mime     string
	Width    int
	Height   int
}

func NewFileInfo(url string) (*FileInfo, error) {
	c := behavior.NewUserBehavior()
	resp, err := c.Get(str.Clean(url))
	if err != nil {
		return nil, err
	}

	mime := resp.Header.Get("Content-Type")
	filename := path.Base(url)

	ext, err := encimg.ImageExt(filename, mime)
	if err != nil {
		return nil, err
	}
	img, err := encimg.Decord(resp.Body, ext)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Filename: filename,
		Url:      url,
		Ext:      ext,
		Mime:     mime,
		Width:    img.Bounds().Dx(),
		Height:   img.Bounds().Dy(),
	}, nil
}

type Info struct {
	*behavior.UserBehavior
	FileInfo *FileInfo
}

func NewInfo() *Info {
	return &Info{
		UserBehavior: behavior.NewUserBehavior(),
	}
}

func (c *Info) Fetch(url string) (*FileInfo, error) {
	resp, err := c.Get(str.Clean(url))
	if err != nil {
		return nil, err
	}

	mime := resp.Header.Get("Content-Type")
	filename := path.Base(url)

	ext, err := encimg.ImageExt(filename, mime)
	if err != nil {
		return nil, err
	}
	img, err := encimg.Decord(resp.Body, ext)
	if err != nil {
		return nil, err
	}

	c.FileInfo = &FileInfo{
		Filename: filename,
		Url:      url,
		Ext:      ext,
		Mime:     mime,
		Width:    img.Bounds().Dx(),
		Height:   img.Bounds().Dy(),
	}
	return c.FileInfo, nil
}
