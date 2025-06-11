package sulfur

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/analog-substance/sulfur/cmd/acid/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

func New(apiEndpoint, user, pass string) *APIClient {
	return &APIClient{
		url:      apiEndpoint,
		username: user,
		password: pass,
	}
}

type APIClient struct {
	url      string
	username string
	password string
	token    string
}

func (a *APIClient) getToken() (string, error) {
	if a.token != "" {
		return a.token, nil
	}

	if a.username == "" {
		return "", errors.New("username not set")
	}

	if a.password == "" {
		return "", errors.New("password not set")
	}

	auth, err := a.Authenticate(sulfur.AuthRequest{
		Identity: a.username,
		Password: a.password,
	})

	if err != nil {
		return "", err
	}

	a.token = auth.Token
	return a.token, nil
}

func (a *APIClient) Do(req *http.Request) (*http.Response, error) {

	token := ""
	if req.URL.Path != sulfur.AuthUserPath && req.URL.Path != sulfur.AuthSuperUserPath {
		var err error
		token, err = a.getToken()
		if err != nil {
			return nil, err
		}
	}

	// Set standard headers
	req.Header.Set("Authorization", token)
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

func (a *APIClient) PostJSON(path string, body []byte) (*http.Response, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", a.url, path), bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("error creating POST request: %v", err)
	}

	return a.Do(req)

}

func (a *APIClient) GetJSON(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", a.url, path), nil)

	if err != nil {
		return nil, fmt.Errorf("error creating GET request: %v", err)
	}

	return a.Do(req)
}

func (a *APIClient) PostStruct(path string, structData any, resStruc any) error {
	body, err := json.Marshal(structData)
	if err != nil {
		//return nil, err
		return err
	}

	//return a.PostJSON(path, body)
	resp, err := a.PostJSON(path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		//return nil, err
		return err
	}
	err = json.Unmarshal(responseBody, resStruc)
	if err != nil {
		//return nil, err
		return err
	}
	return nil
}

func (a *APIClient) GetStruct(path string, resStruc any) error {
	resp, err := a.GetJSON(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		//return nil, err
		return err
	}
	err = json.Unmarshal(responseBody, resStruc)
	if err != nil {
		//return nil, err
		return err
	}
	return nil
}

func (a *APIClient) Authenticate(authReq sulfur.AuthRequest) (*sulfur.AuthResponse, error) {
	authRes := sulfur.AuthResponse{}

	err := a.PostStruct(sulfur.AuthSuperUserPath, authReq, &authRes)
	if err != nil {
		return nil, err
	}

	return &authRes, nil
}

func (a *APIClient) ListRootDomains() (*sulfur.RootDomainsListResponse, error) {
	resStruct := sulfur.RootDomainsListResponse{}

	err := a.GetStruct(sulfur.RootDomainSPath, &resStruct)
	if err != nil {
		return nil, err
	}

	return &resStruct, nil
}

func (a *APIClient) ImportDNSRecords(domainsToImport []sulfur.DNSRecord) {
	body, err := json.Marshal(domainsToImport)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := a.PostJSON(sulfur.ImportDNSRecordsPath, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func (a *APIClient) ImportOrgRootDomains(orgId string, domainsToImport []sulfur.OrgRootDomain) {
	body, err := json.Marshal(domainsToImport)
	if err != nil {
		fmt.Println(err)
		return
	}

	apiPath := strings.Replace(sulfur.ImportOrgRootDomainsPath, "{org_id}", orgId, -1)

	resp, err := a.PostJSON(apiPath, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func AddCIDR(org string, cidr *model.CIDR) {

	cidr.Organization = org

	body, err := json.Marshal(cidr)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.Post("http://localhost:8090/api/collections/asset_cidrs/records", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
}

func Add(org string, scopeItems ...string) {

	for _, scopeItem := range scopeItems {
		if strings.Contains(scopeItem, "/") {
			// could be CIDR
			_, ipNet, err := net.ParseCIDR(scopeItem)
			if err == nil {
				// no err

				c := model.CIDRFromNetCIDR(ipNet)
				AddCIDR(org, c)
			}
		} else {
			// no /, maybe ip
			ip := net.ParseIP(scopeItem)
			if ip != nil {
				// no err

				c := model.CIDRFromNetIP(&ip)
				AddCIDR(org, c)
			}
		}

	}

}

//
//
//func normalizedScope(scopeItem string) string {
//	scopeItem = strings.TrimSpace(scopeItem)
//	if len(scopeItem) == 0 {
//		return ""
//	}
//
//	containsProto := strings.Contains(scopeItem, "://")
//	if !containsProto && strings.Contains(scopeItem, "/") {
//		// perhaps we have a CIDR
//		_, ipNet, err := net.ParseCIDR(scopeItem)
//		if err == nil {
//			return ipNet.String()
//		}
//	}
//
//	possibleIPv6 := ipv6Regexp.MatchString(scopeItem)
//	if !containsProto {
//		if possibleIPv6 {
//			scopeItem = fmt.Sprintf("[%s]", scopeItem)
//		}
//		scopeItem = fmt.Sprintf("https://%s", scopeItem)
//	}
//
//	// we may have a url
//	parsedURL, err := url.Parse("a" + scopeItem)
//	if err == nil {
//		// no errors, we have a url
//		if len(parsedURL.Host) > 0 {
//			hostname := strings.TrimSuffix(parsedURL.Hostname(), ".")
//			return hostname
//			//_, err := publicsuffix.EffectiveTLDPlusOne(hostname)
//			//if err == nil {
//			//	if state.Debug {
//			//		log.Println("hostname", hostname)
//			//	}
//			//	return hostname
//			//} else {
//			//	log.Println("root domain err", err)
//			//}
//		}
//	}
//
//	// must be invalid
//	return ""
//}
