package router

import (
	"encoding/json"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/jobs"
	"github.com/analog-substance/sulfur/pkg/model"
	"github.com/analog-substance/sulfur/pkg/sulfur"
	"github.com/miekg/dns"
	"github.com/pocketbase/pocketbase/core"
	"io"
	"net/http"
	"time"
)

func importDNSRecords(e *core.RequestEvent) error {
	records := []sulfur.DNSRecord{}
	jsonBytes, err := io.ReadAll(e.Request.Body)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	err = json.Unmarshal(jsonBytes, &records)
	if err != nil {
		return e.String(http.StatusBadRequest, "invalid request body")
	}

	go ImportDNSRecords(records)
	return e.String(http.StatusOK, "done")
}

func ImportDNSRecords(records []sulfur.DNSRecord) {
	logger := app_state.GetApp().Logger().WithGroup("importDNSRecords")
	logger.Info("import dns started", "count", len(records))

	extra := []sulfur.DNSRecord{}
	for _, record := range records {

		// lets see if we can lookup missing A records
		// this occurs when r53 has an alias A record
		if record.Value == "" {
			if record.Type == "A" {
				dnsData, err := jobs.DNSQueryMultiple(record.Name, []uint16{dns.TypeA})
				if err != nil {
					logger.Error("JIT DNS lookup failed", "record", record, "err", err)
					continue
				}

				recordCount := len(dnsData.A)
				logger.Info("JIT DNS lookup completed", "record", record, "recordCount", recordCount, "dnsData", dnsData)
				for i, aRecord := range dnsData.A {
					logger.Debug("processing JIT dns record", "aRecord", aRecord, "record", record)

					if i == 0 {
						record.Value = aRecord
					} else {
						extra = append(extra, sulfur.DNSRecord{Name: record.Name, Value: aRecord, Type: record.Type, TTL: int(dnsData.TTL)})
					}
				}
			}
		}
	}

	logger.Info("appending extra records from JIT resolution", "count", len(extra))
	records = append(records, extra...)

	for _, record := range records {
		if record.Value == "" {
			logger.Error("skipping DNS record without a value", "record", record)
			continue
		}

		dnsr, err := model.DNSRecordFirstOrCreate(record.Name, record.Value, record.Type)
		if err != nil {
			logger.Error("unable to find or create dns record", "err", err)
			continue
		}

		ttl := time.Duration(record.TTL) * time.Second

		dnsr.SetTTL(ttl / time.Second)
		dnsr.SetLastSeen(time.Now())
		dnsr.SetLastResolved(time.Now())
		if record.ExternalReference != "" {
			extRef, err := model.ExternalReferenceFirstOrCreate(record.ExternalReference)
			if err != nil {
				logger.Error("unable to find or create external reference", "extRef", record.ExternalReference, "err", err)
			} else {
				if err := extRef.Save(); err != nil {
					logger.Error("unable to save external reference", "extRef", record.ExternalReference, "err", err)
				} else {
					dnsr.SetExternalReference(extRef)
				}
			}
		}

		err = dnsr.Save()
		if err != nil {
			logger.Error("err saving dns record", "record", "err", err)
		}
	}
	logger.Info("import dns complete", "count", len(records))

}
