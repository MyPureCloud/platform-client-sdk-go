package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Userapp - Details for a UserApp
type Userapp struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the userApp, used to distinguish this userApp from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Integration Type for the userApp
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// Config
	Config *Userappconfigurationinfo `json:"config,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Userapp) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Userapp

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationType *Integrationtype `json:"integrationType,omitempty"`
		
		Config *Userappconfigurationinfo `json:"config,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		IntegrationType: u.IntegrationType,
		
		Config: u.Config,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Userapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
