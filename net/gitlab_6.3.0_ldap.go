package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client :=
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	form := url.Values{}
	form.Add("username", "wangtianyi1")
	form.Add("password", "TWwty@0032")
	req, _ := http.NewRequest("POST",
		"http://wpsgit.kingsoft.net/users/auth/ldap/callback",
		strings.NewReader(form.Encode()))
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	req, _ = http.NewRequest("POST",
		"http://wpsgit.kingsoft.net/",
		strings.NewReader(form.Encode()))
	req.Header.Add("Cookie", res.Cookies()[2].Name+"="+res.Cookies()[2].Value)
	res, _ = client.Do(req)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
