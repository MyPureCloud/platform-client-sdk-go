package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Updatedraftinput - Definition of an Action Draft to be created or updated.
type Updatedraftinput struct { 
	// Category - Category of action, Can be up to 256 characters long
	Category *string `json:"category,omitempty"`


	// Name - Name of action, Can be up to 256 characters long
	Name *string `json:"name,omitempty"`


	// Config - Configuration to support request and response processing
	Config *Actionconfig `json:"config,omitempty"`


	// Contract - Action contract
	Contract *Actioncontractinput `json:"contract,omitempty"`


	// Secure - Indication of whether or not the action is designed to accept sensitive data
	Secure *bool `json:"secure,omitempty"`


	// Version - Version of current Draft
	Version *int `json:"version,omitempty"`

}

func (u *Updatedraftinput) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Updatedraftinput

	

	return json.Marshal(&struct { 
		Category *string `json:"category,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Config *Actionconfig `json:"config,omitempty"`
		
		Contract *Actioncontractinput `json:"contract,omitempty"`
		
		Secure *bool `json:"secure,omitempty"`
		
		Version *int `json:"version,omitempty"`
		*Alias
	}{ 
		Category: u.Category,
		
		Name: u.Name,
		
		Config: u.Config,
		
		Contract: u.Contract,
		
		Secure: u.Secure,
		
		Version: u.Version,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Updatedraftinput) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
