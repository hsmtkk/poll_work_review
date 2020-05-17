package login

import (
	"io/ioutil"
	"log"
	"net/http"
)

type AuthCookieGetter interface {
	GetAuthCookie(userName, password string) (string, error)
}

const LoginURL = "https://techacademy.jp/mentor/login"
const SessionURL = "https://techacademy.jp/mentor/sessions"

func New() AuthCookieGetter {
	return &impl{client: http.DefaultClient, url: LoginURL}
}

func NewClient(client *http.Client) AuthCookieGetter {
	return &impl{client: client, url: LoginURL}
}

func NewClientURL(client *http.Client, url string) AuthCookieGetter {
	return &impl{client: client, url: url}
}

type impl struct {
	client *http.Client
	url    string
}

func (i *impl) GetAuthCookie(userName, password string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, i.url, nil)
	if err != nil {
		return "", err
	}
	resp, err := i.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	log.Println(string(bs))
	return "", nil
}
