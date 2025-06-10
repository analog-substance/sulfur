package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"log"
	"net/http"
	"time"
)

func consumeDNSRecord(e *core.RequestEvent) error {
	records := []sulfur.DNSRecord{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &records)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	for _, record := range records {
		dnsr, err := model.DNSRecordFirstOrCreate(record.Name, record.Value, record.Type)
		if err != nil {
			log.Println("unable to find or create dns record", err)
			continue
		}
		dnsr.SetTTL(time.Duration(record.TTL) * time.Second)
		dnsr.SetLastSeen(time.Now())
		dnsr.SetLastResolved(time.Now())

		rootDomain, err := model.RootDomainFirstOrCreate(record.Name)
		if err != nil {
			log.Println("unable to find or create root domain record", err)
			continue
		}

		if rootDomain.Id() == "" {
			err = rootDomain.Save()
			if err != nil {
				log.Println("unable to save root domain", err)
				continue
			}
		}

		log.Println("Found root domain", rootDomain.DomainName())
		dnsr.SetRootDomain(rootDomain)

		err = dnsr.Save()
		if err != nil {
			log.Println("err saving dns record", err)
		}
	}
	return e.String(http.StatusOK, "done")
}
