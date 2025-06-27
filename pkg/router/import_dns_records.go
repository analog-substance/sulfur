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
					logger.Error("failed to lookup missing dns record", "name", record.Name, "type", record.Type, "err", err)
				}
				if len(record.Value) > 0 {
					record.Value = dnsData.A[0]
					if len(record.Value) > 1 {
						extra = append(extra, sulfur.DNSRecord{Name: record.Name, Value: record.Value, Type: record.Type, TTL: int(dnsData.TTL)})
					}
				}
			}
		}
	}

	records = append(records, extra...)

	for _, record := range records {
		if record.Value == "" {
			logger.Error("skipping DNS record without a value", "name", record.Name, "type", record.Type)
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

		err = dnsr.Save()
		if err != nil {
			logger.Error("err saving dns record", "recordName", record.Name, "recordVal", record.Value, "recordType", record.Type, "err", err)
		}
	}
	logger.Info("import dns complete", "count", len(records))

}
