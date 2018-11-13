package loginer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/xbox1994/gpb/repository/model"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Gitlab630Ldap struct {
}

func (g *Gitlab630Ldap) Login(loginInfo model.LoginInfo) (repoCreatePreInfo RepoCreatePreInfo) {
	// 不会自动重定向的Client
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// 登录获取Cookie
	fmt.Println("Login in...")
	form := url.Values{}
	form.Add("username", loginInfo.Username)
	form.Add("password", loginInfo.Password)
	req, _ := http.NewRequest("POST",
		loginInfo.GitHostAddress+"users/auth/ldap/callback",
		strings.NewReader(form.Encode()))
	res, _ := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("login failed, please check your login info")
		os.Exit(1)
	}
	if strings.Contains(string(body), loginInfo.GitHostAddress) {
		fmt.Println("Get Cookie success")
		repoCreatePreInfo.Cookie = res.Cookies()[2].Name + "=" + res.Cookies()[2].Value
	}
	res.Body.Close()

	// 登录后得到想要创建的Repo所在组的Id
	fmt.Println("Get repo group id by name...")
	req, _ = http.NewRequest("GET", loginInfo.GitHostAddress+"projects/new", nil)
	req.Header.Add("Cookie", repoCreatePreInfo.Cookie)
	res, _ = client.Do(req)
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	namespaceFound := false
	doc.Find("#project_namespace_id optgroup").Each(func(i int, s *goquery.Selection) {
		element := s.Find("option")
		if element.Text() == loginInfo.RepoGroupName {
			RepoGroupNameId, _ := element.Attr("value")
			repoCreatePreInfo.RepoGroupNameId = RepoGroupNameId
			fmt.Println("Get repo namespace id: " + RepoGroupNameId)
			namespaceFound = true
		}
	})
	res.Body.Close()

	if !namespaceFound {
		panic("namespace (groups or users) not found: " + loginInfo.RepoGroupName)
	}
	if repoCreatePreInfo.RepoGroupNameId == "" || repoCreatePreInfo.Cookie == "" {
		panic(fmt.Sprintf("login info invalid: %s, %s\n", repoCreatePreInfo.RepoGroupNameId, repoCreatePreInfo.Cookie))
	}

	return repoCreatePreInfo
}
