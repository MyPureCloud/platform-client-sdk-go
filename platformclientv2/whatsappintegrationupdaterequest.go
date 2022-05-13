package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Whatsappintegrationupdaterequest
type Whatsappintegrationupdaterequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - WhatsApp Integration name
	Name *string `json:"name,omitempty"`


	// SupportedContent - Defines the SupportedContent profile configured for an integration
	SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`


	// MessagingSetting
	MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`


	// Action - The action used to activate and then confirm a WhatsApp Integration.
	Action *string `json:"action,omitempty"`


	// AuthenticationMethod - The authentication method used to confirm a WhatsApp Integration activation. If action is set to Activate, then authenticationMethod is a required field. 
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`


	// ConfirmationCode - The confirmation code sent by Whatsapp to you during the activation step. If action is set to Confirm, then confirmationCode is a required field.
	ConfirmationCode *string `json:"confirmationCode,omitempty"`


	// PhoneNumber - Phone number to associate with the WhatsApp integration
	PhoneNumber *string `json:"phoneNumber,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Whatsappintegrationupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappintegrationupdaterequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SupportedContent *Supportedcontentreference `json:"supportedContent,omitempty"`
		
		MessagingSetting *Messagingsettingreference `json:"messagingSetting,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
		
		ConfirmationCode *string `json:"confirmationCode,omitempty"`
		
		PhoneNumber *string `json:"phoneNumber,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SupportedContent: o.SupportedContent,
		
		MessagingSetting: o.MessagingSetting,
		
		Action: o.Action,
		
		AuthenticationMethod: o.AuthenticationMethod,
		
		ConfirmationCode: o.ConfirmationCode,
		
		PhoneNumber: o.PhoneNumber,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Whatsappintegrationupdaterequest) UnmarshalJSON(b []byte) error {
	var WhatsappintegrationupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &WhatsappintegrationupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WhatsappintegrationupdaterequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WhatsappintegrationupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if SupportedContent, ok := WhatsappintegrationupdaterequestMap["supportedContent"].(map[string]interface{}); ok {
		SupportedContentString, _ := json.Marshal(SupportedContent)
		json.Unmarshal(SupportedContentString, &o.SupportedContent)
	}
	
	if MessagingSetting, ok := WhatsappintegrationupdaterequestMap["messagingSetting"].(map[string]interface{}); ok {
		MessagingSettingString, _ := json.Marshal(MessagingSetting)
		json.Unmarshal(MessagingSettingString, &o.MessagingSetting)
	}
	
	if Action, ok := WhatsappintegrationupdaterequestMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if AuthenticationMethod, ok := WhatsappintegrationupdaterequestMap["authenticationMethod"].(string); ok {
		o.AuthenticationMethod = &AuthenticationMethod
	}
    
	if ConfirmationCode, ok := WhatsappintegrationupdaterequestMap["confirmationCode"].(string); ok {
		o.ConfirmationCode = &ConfirmationCode
	}
    
	if PhoneNumber, ok := WhatsappintegrationupdaterequestMap["phoneNumber"].(string); ok {
		o.PhoneNumber = &PhoneNumber
	}
    
	if SelfUri, ok := WhatsappintegrationupdaterequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Whatsappintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
