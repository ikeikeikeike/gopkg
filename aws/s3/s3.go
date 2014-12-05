package s3

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/s3"
)

func Md5HexDigest(body io.Reader) string {
	hash := md5.New()
	c, _ := ioutil.ReadAll(body)
	io.WriteString(hash, string(c))
	return hex.EncodeToString(hash.Sum(nil))
}

type BaseClient struct {
	Auth   aws.Auth
	Bucket string
}

type NamedClient struct {
	*BaseClient
}

func NewNamedClient() *NamedClient {
	auth, _ := aws.EnvAuth()
	return &NamedClient{
		&BaseClient{Auth: auth},
	}
}

func (nc *NamedClient) SetAuth(key, secret string) *NamedClient {
	nc.Auth = aws.Auth{
		AccessKey: key,
		SecretKey: secret,
	}
	return nc
}

func (nc *NamedClient) SetBucket(name string) *NamedClient {
	nc.Bucket = name
	return nc
}

func (nc *NamedClient) PutByString(src, ctype string) (string, error) {
	return nc.put(strings.NewReader(src), ctype)
}

func (nc *NamedClient) PutByBytes(src []byte, ctype string) (string, error) {
	return nc.put(bytes.NewReader(src), ctype)
}

func (nc *NamedClient) PutByUrl(src string) (dst string, err error) {
	resp, err := http.Get(src)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return nc.put(resp.Body, resp.Header.Get("Content-Type"))
}

func (nc *NamedClient) put(body io.Reader, ctype string) (dst string, err error) {
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	b3 := new(bytes.Buffer)
	io.Copy(io.MultiWriter(b1, b2, b3), body)

	bucket := s3.New(nc.Auth, aws.APNortheast).Bucket(nc.Bucket)

	clength, _ := ioutil.ReadAll(b1)
	dst = fmt.Sprintf("%s.ico", Md5HexDigest(b2))

	err = bucket.PutReader(
		dst, b3, int64(len(clength)), ctype, s3.PublicRead, s3.Options{})
	return
}
