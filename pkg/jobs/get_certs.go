package jobs

import (
	"context"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/analog-substance/sulfur/pkg/app_state"
	"github.com/analog-substance/sulfur/pkg/iface"
	"github.com/analog-substance/sulfur/pkg/model"
	"strings"
	"time"
)

type CheckCertStatus struct {
	IPPort iface.IPPort
	Domain string
	Certs  []*x509.Certificate
}

func CheckCerts() {
	logger := app_state.GetApp().Logger().WithGroup("CheckCerts")
	records, err := model.GetCertsQueue()
	if err != nil {
		logger.Error("failed to get cert queue: ", "err", err)
		return
	}

	logger.Info("Total ports to check", "count", len(records))

	total := 0

	c := make(chan CheckCertStatus)

	for _, record := range records {
		for _, domain := range record.GetIP().GetDomains() {
			total += 1
			go CheckCert(domain, record, c)
		}
	}

	result := make([]CheckCertStatus, total)
	for i, _ := range result {
		result[i] = <-c
		if len(result[i].Certs) > 0 {
			for _, cert := range result[i].Certs {
				if !cert.IsCA {
					fp := fmt.Sprintf("%x", sha1.Sum(cert.Raw))
					certRecord, err := model.CertificateFirstOrCreate(fp)
					if err != nil {
						logger.Error("failed tp create cert", "error", err)
						continue
					}
					certRecord.SetSubject(strings.ToLower(cert.Subject.CommonName))
					certRecord.SetAlternativeNames(strings.ToLower(strings.Join(cert.DNSNames, ",")))
					certRecord.SetIssuer(cert.Issuer.CommonName)
					certRecord.SetIssued(cert.NotBefore)
					certRecord.SetExpires(cert.NotAfter)

					if err := certRecord.Save(); err != nil {
						logger.Error("failed to save cert", "error", err, "subject", cert.Subject, "issuer", cert.Issuer)
						continue
					}

					if portCert, err := model.IPPortCertificateFirstOrCreate(result[i].IPPort.Id(), certRecord.Id()); err != nil {
						logger.Error("failed to create cert port", "error", err)
						continue
					} else {
						portCert.SetLastSeen(time.Now())

						if err := portCert.Save(); err != nil {
							logger.Error("failed to save cert", "error", err)
							continue
						}
					}
				}
			}
		}
	}
}

func CheckCert(domain string, ipPort iface.IPPort, c chan CheckCertStatus) {

	if strings.HasPrefix(domain, "*.") {
		domain = fmt.Sprintf("s%d.%s", time.Now().UnixMilli(), domain[2:])
	}

	var conf = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         domain,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	d := tls.Dialer{
		Config: conf,
	}
	conn, err := d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%d", ipPort.GetIP().Address(), ipPort.Port()))
	cancel() // Ensure cancel is always called

	if err != nil {
		c <- CheckCertStatus{
			IPPort: ipPort,
			Certs:  []*x509.Certificate{},
		}
		return
	}
	defer conn.Close()

	tlsConn := conn.(*tls.Conn)

	c <- CheckCertStatus{
		IPPort: ipPort,
		Domain: domain,
		Certs:  tlsConn.ConnectionState().PeerCertificates,
	}

}
