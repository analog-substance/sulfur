package sulfur

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/analog-substance/sulfur/cmd/acid/pkg/model"
	"log"
	"net"
	"net/http"
	"strings"
)

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

	log.Println(resp.StatusCode)
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
//	// we may have a URL
//	parsedURL, err := url.Parse("a" + scopeItem)
//	if err == nil {
//		// no errors, we have a URL
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
