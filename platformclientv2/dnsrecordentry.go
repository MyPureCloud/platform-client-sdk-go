package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dnsrecordentry
type Dnsrecordentry struct { 
	// Host - the hostname of the DNS entry
	Host *string `json:"host,omitempty"`


	// RecordContents - the payload of the DNS entry
	RecordContents *string `json:"recordContents,omitempty"`


	// VerificationStatus - the current status of the related verification process
	VerificationStatus *string `json:"verificationStatus,omitempty"`

}

func (o *Dnsrecordentry) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dnsrecordentry
	
	return json.Marshal(&struct { 
		Host *string `json:"host,omitempty"`
		
		RecordContents *string `json:"recordContents,omitempty"`
		
		VerificationStatus *string `json:"verificationStatus,omitempty"`
		*Alias
	}{ 
		Host: o.Host,
		
		RecordContents: o.RecordContents,
		
		VerificationStatus: o.VerificationStatus,
		Alias:    (*Alias)(o),
	})
}

func (o *Dnsrecordentry) UnmarshalJSON(b []byte) error {
	var DnsrecordentryMap map[string]interface{}
	err := json.Unmarshal(b, &DnsrecordentryMap)
	if err != nil {
		return err
	}
	
	if Host, ok := DnsrecordentryMap["host"].(string); ok {
		o.Host = &Host
	}
    
	if RecordContents, ok := DnsrecordentryMap["recordContents"].(string); ok {
		o.RecordContents = &RecordContents
	}
    
	if VerificationStatus, ok := DnsrecordentryMap["verificationStatus"].(string); ok {
		o.VerificationStatus = &VerificationStatus
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dnsrecordentry) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
