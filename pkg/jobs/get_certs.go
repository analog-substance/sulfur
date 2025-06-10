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
	"log"
	"strings"
	"time"
)

type CheckCertStatus struct {
	IPPort iface.IPPort
	Domain string
	Certs  []*x509.Certificate
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

func CheckCerts() {
	records, err := model.GetCertsQueue()
	if err != nil {
		app_state.GetApp().Logger().Error("failed to get dns records to lookup: ", "err", err)
		log.Println("failed to get dns records to lookup: ", "err", err)
		return
	}

	log.Println("Total ports to check: ", len(records))

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
						log.Println("failed to create cert: ", err)
						continue
					}
					certRecord.SetSubject(cert.Subject.CommonName)
					certRecord.SetAlternativeNames(strings.Join(cert.DNSNames, ","))
					certRecord.SetIssuer(cert.Issuer.CommonName)
					certRecord.SetIssued(cert.NotBefore)
					certRecord.SetExpires(cert.NotAfter)

					if err := certRecord.Save(); err != nil {
						log.Println("failed to save cert: ", err, cert.Subject, cert.Issuer, cert.DNSNames, cert.EmailAddresses)
					}

					if portCert, err := model.IPPortCertificateFirstOrCreate(result[i].IPPort.ProxyRecord().Id, certRecord.Id()); err != nil {
						log.Println("failed to create port cert: ", err)
						continue
					} else if err := portCert.Save(); err != nil {
						log.Println("failed to save port cert: ", err)
						continue
					}
				}
			}
		}
	}
}
