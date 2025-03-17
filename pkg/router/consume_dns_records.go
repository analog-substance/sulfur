package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/pocketbase/pocketbase/core"
	"golang.org/x/net/publicsuffix"
	"io"
	"log"
	"net/http"
	"time"
)

type dnsRecord struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
	TTL   int    `json:"ttl"`
}

func consumeDNSRecord(e *core.RequestEvent) error {
	records := []dnsRecord{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &records)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	for _, record := range records {

		rootDomain, err := publicsuffix.EffectiveTLDPlusOne(record.Name)
		if err != nil {
			return err
		}

		rdr, err := model.FindRootDomain(rootDomain)
		if err != nil {
			log.Println(err)
		} else {

			log.Println(rdr)
		}

		dnsr, err := model.DNSRecordFirstOrCreate(record.Name, record.Value, record.Type)
		if err != nil {
			log.Println(err)
			continue
		}
		dnsr.SetTTL(time.Duration(record.TTL) * time.Second)
		dnsr.SetLastSeen(time.Now())
		dnsr.SetLastResolved(time.Now())

		err = dnsr.Save()
		if err != nil {
			log.Println(err)
		}

	}
	return e.String(http.StatusOK, "done")
}
