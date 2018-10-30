package repo_creator

import (
	"fmt"
	"grb/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Gitlab630Ldap struct {
}

func (g *Gitlab630Ldap) CreateRepo(answer model.Answer) {
	client :=
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	form := url.Values{}
	form.Add("username", answer.Username)
	form.Add("password", answer.Password)
	req, _ := http.NewRequest("POST",
		answer.GitHostAddress+"users/auth/ldap/callback",
		strings.NewReader(form.Encode()))
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	form = url.Values{}
	form.Add("project[name]", answer.RepoName)
	form.Add("project[namespace_id]", "1102")
	req, _ = http.NewRequest("POST",
		answer.GitHostAddress+"projects",
		strings.NewReader(form.Encode()))
	req.Header.Add("Cookie", res.Cookies()[2].Name+"="+res.Cookies()[2].Value)
	res, _ = client.Do(req)
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
