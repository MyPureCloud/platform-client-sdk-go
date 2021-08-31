package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Inbounddomain
type Inbounddomain struct { 
	// Id - Unique Id of the domain such as: example.com
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// MxRecordStatus - Mx Record Status
	MxRecordStatus *string `json:"mxRecordStatus,omitempty"`


	// SubDomain - Indicates if this a PureCloud sub-domain.  If true, then the appropriate DNS records are created for sending/receiving email.
	SubDomain *bool `json:"subDomain,omitempty"`


	// MailFromSettings - The DNS settings if the inbound domain is using a custom Mail From. These settings can only be used on InboundDomains where subDomain is false.
	MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`


	// CustomSMTPServer - The custom SMTP server integration to use when sending outbound emails from this domain.
	CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Inbounddomain) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Inbounddomain
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		MxRecordStatus *string `json:"mxRecordStatus,omitempty"`
		
		SubDomain *bool `json:"subDomain,omitempty"`
		
		MailFromSettings *Mailfromresult `json:"mailFromSettings,omitempty"`
		
		CustomSMTPServer *Domainentityref `json:"customSMTPServer,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		MxRecordStatus: o.MxRecordStatus,
		
		SubDomain: o.SubDomain,
		
		MailFromSettings: o.MailFromSettings,
		
		CustomSMTPServer: o.CustomSMTPServer,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Inbounddomain) UnmarshalJSON(b []byte) error {
	var InbounddomainMap map[string]interface{}
	err := json.Unmarshal(b, &InbounddomainMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InbounddomainMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := InbounddomainMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if MxRecordStatus, ok := InbounddomainMap["mxRecordStatus"].(string); ok {
		o.MxRecordStatus = &MxRecordStatus
	}
	
	if SubDomain, ok := InbounddomainMap["subDomain"].(bool); ok {
		o.SubDomain = &SubDomain
	}
	
	if MailFromSettings, ok := InbounddomainMap["mailFromSettings"].(map[string]interface{}); ok {
		MailFromSettingsString, _ := json.Marshal(MailFromSettings)
		json.Unmarshal(MailFromSettingsString, &o.MailFromSettings)
	}
	
	if CustomSMTPServer, ok := InbounddomainMap["customSMTPServer"].(map[string]interface{}); ok {
		CustomSMTPServerString, _ := json.Marshal(CustomSMTPServer)
		json.Unmarshal(CustomSMTPServerString, &o.CustomSMTPServer)
	}
	
	if SelfUri, ok := InbounddomainMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Inbounddomain) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
