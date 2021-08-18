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


	// Action - The action used to activate and then confirm a WhatsApp Integration.
	Action *string `json:"action,omitempty"`


	// AuthenticationMethod - The authentication method used to confirm a WhatsApp Integration activation. If action is set to Activate, then authenticationMethod is a required field. 
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`


	// ConfirmationCode - The confirmation code sent by Whatsapp to you during the activation step. If action is set to Confirm, then confirmationCode is a required field.
	ConfirmationCode *string `json:"confirmationCode,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Whatsappintegrationupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Whatsappintegrationupdaterequest

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
		
		ConfirmationCode *string `json:"confirmationCode,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Action: u.Action,
		
		AuthenticationMethod: u.AuthenticationMethod,
		
		ConfirmationCode: u.ConfirmationCode,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Whatsappintegrationupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
