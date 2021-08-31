package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Inbounddomainpatchrequest
type Inbounddomainpatchrequest struct { 
	// MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
	MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`


	// CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
	CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`

}

func (o *Inbounddomainpatchrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Inbounddomainpatchrequest
	
	return json.Marshal(&struct { 
		MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`
		
		CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`
		*Alias
	}{ 
		MailFromSettings: o.MailFromSettings,
		
		CustomSMTPServer: o.CustomSMTPServer,
		Alias:    (*Alias)(o),
	})
}

func (o *Inbounddomainpatchrequest) UnmarshalJSON(b []byte) error {
	var InbounddomainpatchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &InbounddomainpatchrequestMap)
	if err != nil {
		return err
	}
	
	if MailFromSettings, ok := InbounddomainpatchrequestMap["mailFromSettings"].(map[string]interface{}); ok {
		MailFromSettingsString, _ := json.Marshal(MailFromSettings)
		json.Unmarshal(MailFromSettingsString, &o.MailFromSettings)
	}
	
	if CustomSMTPServer, ok := InbounddomainpatchrequestMap["customSMTPServer"].(map[string]interface{}); ok {
		CustomSMTPServerString, _ := json.Marshal(CustomSMTPServer)
		json.Unmarshal(CustomSMTPServerString, &o.CustomSMTPServer)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Inbounddomainpatchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
