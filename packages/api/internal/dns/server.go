package dns

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	resolver "github.com/miekg/dns"

	"github.com/e2b-dev/infra/packages/shared/pkg/smap"
)

const ttl = 0

const defaultRoutingIP = "127.0.0.1"

type dnsRecord struct {
	IP         string
	InternalID string
}
type DNS struct {
	mu      sync.Mutex
	records *smap.Map[dnsRecord]
}

func New() *DNS {
	return &DNS{
		records: smap.New[dnsRecord](),
	}
}

func (d *DNS) Add(sandboxID, internalID, ip string) {
	d.records.Insert(d.hostname(sandboxID), dnsRecord{IP: ip, InternalID: internalID})
}

func (d *DNS) Remove(sandboxID, internalID string) {
	d.records.RemoveCb(d.hostname(sandboxID), func(key string, v dnsRecord, exists bool) bool {
		return v.InternalID == internalID
	})
}

func (d *DNS) get(hostname string) (string, bool) {
	v, ok := d.records.Get(hostname)
	if !ok {
		return "", false
	}

	return v.IP, true
}

func (*DNS) hostname(sandboxID string) string {
	return fmt.Sprintf("%s.", sandboxID)
}

func (d *DNS) handleDNSRequest(w resolver.ResponseWriter, r *resolver.Msg) {
	m := new(resolver.Msg)
	m.SetReply(r)
	m.Compress = false
	m.Authoritative = true

	for _, q := range m.Question {
		if q.Qtype == resolver.TypeA {
			a := &resolver.A{
				Hdr: resolver.RR_Header{
					Name:   q.Name,
					Rrtype: resolver.TypeA,
					Class:  resolver.ClassINET,
					Ttl:    ttl,
				},
			}

			sandboxID := strings.Split(q.Name, "-")[0]
			ip, found := d.get(sandboxID)
			if found {
				a.A = net.ParseIP(ip).To4()
			} else {
				a.A = net.ParseIP(defaultRoutingIP).To4()
			}

			m.Answer = append(m.Answer, a)
		}
	}

	err := w.WriteMsg(m)
	if err != nil {
		log.Printf("Failed to write message: %s\n", err.Error())
	}
}

func (d *DNS) Start(address string, port int) error {
	mux := resolver.NewServeMux()

	mux.HandleFunc(".", d.handleDNSRequest)

	server := resolver.Server{Addr: fmt.Sprintf("%s:%d", address, port), Net: "udp", Handler: mux}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start DNS server: %w", err)
	}

	return nil
}
