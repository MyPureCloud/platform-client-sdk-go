package platformclientv2
import (
	"encoding/json"
)

// Inbounddomainpatchrequest
type Inbounddomainpatchrequest struct { 
	// MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
	MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`


	// CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
	CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`

}

// String returns a JSON representation of the model
func (o *Inbounddomainpatchrequest) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
