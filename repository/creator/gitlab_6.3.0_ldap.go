package creator

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"grb/repository/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Gitlab630Ldap struct {
	cookie          string
	repoNamespaceId string
	client          http.Client
}

func (g *Gitlab630Ldap) Login(loginInfo model.LoginInfo) {
	fmt.Println("Login in...")
	g.client =
		http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	setCookie(loginInfo, g)
	setRepoNamespaceId(loginInfo, g)

	if g.repoNamespaceId == "" || g.cookie == "" {
		fmt.Printf("login info invalid: %s, %s\n", g.repoNamespaceId, g.cookie)
	}
}

func setCookie(loginInfo model.LoginInfo, g *Gitlab630Ldap) {
	form := url.Values{}
	form.Add("username", loginInfo.Username)
	form.Add("password", loginInfo.Password)
	req, _ := http.NewRequest("POST",
		loginInfo.GitHostAddress+"users/auth/ldap/callback",
		strings.NewReader(form.Encode()))
	res, _ := g.client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(body), loginInfo.GitHostAddress) {
		fmt.Println("Get Cookie success")
		g.cookie = res.Cookies()[2].Name + "=" + res.Cookies()[2].Value
	}
}

func setRepoNamespaceId(loginInfo model.LoginInfo, g *Gitlab630Ldap) {
	req, _ := http.NewRequest("GET", loginInfo.GitHostAddress+"projects/new", nil)
	req.Header.Add("Cookie", g.cookie)
	res, _ := g.client.Do(req)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("#project_namespace_id optgroup").Each(func(i int, s *goquery.Selection) {
		element := s.Find("option")
		if element.Text() == loginInfo.RepoNamespace {
			repoNamespaceId, _ := element.Attr("value")
			g.repoNamespaceId = repoNamespaceId
			fmt.Println("Get repo namespace id: " + g.repoNamespaceId)
			return
		}
	})
}

func (g *Gitlab630Ldap) CreateRepo(answer model.Answer) {
	form := url.Values{}
	form.Add("project[name]", answer.RepoName)
	form.Add("project[namespace_id]", g.repoNamespaceId)
	req, _ := http.NewRequest("POST",
		answer.GitHostAddress+"projects",
		strings.NewReader(form.Encode()))
	req.Header.Add("Cookie", g.cookie)
	res, _ := g.client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if strings.Contains(string(body), answer.RepoName) {
		fmt.Printf("create repo %s success\n", answer.RepoName)
	}else{
		fmt.Printf("create repo %s failed\n", answer.RepoName)
	}
}
