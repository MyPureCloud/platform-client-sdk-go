package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Action
type Action struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// IntegrationId - The ID of the integration for which this action is associated
	IntegrationId *string `json:"integrationId,omitempty"`


	// Category - Category of Action
	Category *string `json:"category,omitempty"`


	// Contract - Action contract
	Contract *Actioncontract `json:"contract,omitempty"`


	// Version - Version of this action
	Version *int `json:"version,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Action) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Action

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationId *string `json:"integrationId,omitempty"`
		
		Category *string `json:"category,omitempty"`
		
		Contract *Actioncontract `json:"contract,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		IntegrationId: u.IntegrationId,
		
		Category: u.Category,
		
		Contract: u.Contract,
		
		Version: u.Version,
		
		Secure: u.Secure,
		
		Config: u.Config,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Action) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
