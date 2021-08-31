package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Pingidentity
type Pingidentity struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Disabled
	Disabled *bool `json:"disabled,omitempty"`


	// IssuerURI
	IssuerURI *string `json:"issuerURI,omitempty"`


	// SsoTargetURI
	SsoTargetURI *string `json:"ssoTargetURI,omitempty"`


	// SloURI
	SloURI *string `json:"sloURI,omitempty"`


	// SloBinding
	SloBinding *string `json:"sloBinding,omitempty"`


	// Certificate
	Certificate *string `json:"certificate,omitempty"`


	// Certificates
	Certificates *[]string `json:"certificates,omitempty"`


	// RelyingPartyIdentifier
	RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Pingidentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Pingidentity
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Disabled *bool `json:"disabled,omitempty"`
		
		IssuerURI *string `json:"issuerURI,omitempty"`
		
		SsoTargetURI *string `json:"ssoTargetURI,omitempty"`
		
		SloURI *string `json:"sloURI,omitempty"`
		
		SloBinding *string `json:"sloBinding,omitempty"`
		
		Certificate *string `json:"certificate,omitempty"`
		
		Certificates *[]string `json:"certificates,omitempty"`
		
		RelyingPartyIdentifier *string `json:"relyingPartyIdentifier,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Disabled: o.Disabled,
		
		IssuerURI: o.IssuerURI,
		
		SsoTargetURI: o.SsoTargetURI,
		
		SloURI: o.SloURI,
		
		SloBinding: o.SloBinding,
		
		Certificate: o.Certificate,
		
		Certificates: o.Certificates,
		
		RelyingPartyIdentifier: o.RelyingPartyIdentifier,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Pingidentity) UnmarshalJSON(b []byte) error {
	var PingidentityMap map[string]interface{}
	err := json.Unmarshal(b, &PingidentityMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PingidentityMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := PingidentityMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Disabled, ok := PingidentityMap["disabled"].(bool); ok {
		o.Disabled = &Disabled
	}
	
	if IssuerURI, ok := PingidentityMap["issuerURI"].(string); ok {
		o.IssuerURI = &IssuerURI
	}
	
	if SsoTargetURI, ok := PingidentityMap["ssoTargetURI"].(string); ok {
		o.SsoTargetURI = &SsoTargetURI
	}
	
	if SloURI, ok := PingidentityMap["sloURI"].(string); ok {
		o.SloURI = &SloURI
	}
	
	if SloBinding, ok := PingidentityMap["sloBinding"].(string); ok {
		o.SloBinding = &SloBinding
	}
	
	if Certificate, ok := PingidentityMap["certificate"].(string); ok {
		o.Certificate = &Certificate
	}
	
	if Certificates, ok := PingidentityMap["certificates"].([]interface{}); ok {
		CertificatesString, _ := json.Marshal(Certificates)
		json.Unmarshal(CertificatesString, &o.Certificates)
	}
	
	if RelyingPartyIdentifier, ok := PingidentityMap["relyingPartyIdentifier"].(string); ok {
		o.RelyingPartyIdentifier = &RelyingPartyIdentifier
	}
	
	if SelfUri, ok := PingidentityMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Pingidentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
