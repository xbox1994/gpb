package creater

import (
	"fmt"
	"github.com/xbox1994/wps-gpb/repository/loginer"
	"github.com/xbox1994/wps-gpb/repository/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Gitlab630Ldap struct {
}

func (g *Gitlab630Ldap) CreateRemoteRepo(answer model.Answer, repoCreatePreInfo loginer.RepoCreatePreInfo) {
	// 不会自动重定向的Client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	form := url.Values{}
	form.Add("project[name]", answer.RepoName)
	form.Add("project[namespace_id]", repoCreatePreInfo.RepoNamespaceId)
	req, _ := http.NewRequest("POST",
		answer.GitHostAddress+"projects",
		strings.NewReader(form.Encode()))
	req.Header.Add("Cookie", repoCreatePreInfo.Cookie)
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(body), answer.RepoName) {
		fmt.Printf("create repo %s success\n", answer.RepoName)
	}else{
		fmt.Printf("create repo %s failed\n", answer.RepoName)
	}
}
