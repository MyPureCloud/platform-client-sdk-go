package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegrationrequest
type Whatsappintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the WhatsApp Integration
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// PhoneNumber - The phone number associated to the whatsApp integration
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// WabaCertificate - The waba(WhatsApp Business Manager) certificate associated to the WhatsApp integration phone number
	WabaCertificate *string `json:"wabaCertificate,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Whatsappintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		WabaCertificate *string `json:"wabaCertificate,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		PhoneNumber: o.PhoneNumber,
		
		WabaCertificate: o.WabaCertificate,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappintegrationrequest) UnmarshalJSON(b []byte) error {
	var WhatsappintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WhatsappintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WhatsappintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SupportedContent, ok := WhatsappintegrationrequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if PhoneNumber, ok := WhatsappintegrationrequestMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
	
	if WabaCertificate, ok := WhatsappintegrationrequestMap["wabaCertificate"].(string); ok {
		o.WabaCertificate = &WabaCertificate
	}
	
	if SelfUri, ok := WhatsappintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
