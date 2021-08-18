package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Postactioninput - Definition of an Action to be created or updated.
type Postactioninput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// IntegrationId - The ID of the integration this action is associated to
	IntegrationId *string `json:"integrationId,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Contract - Action contract
	Contract *Actioncontractinput `json:"contract,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`

}

func (u *Postactioninput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Postactioninput

	

	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Contract *Actioncontractinput `json:"contract,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		*Alias
	}{ 
		Category: u.Category,
		
		Name: u.Name,
		
		IntegrationId: u.IntegrationId,
		
		Config: u.Config,
		
		Contract: u.Contract,
		
		Secure: u.Secure,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Postactioninput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
