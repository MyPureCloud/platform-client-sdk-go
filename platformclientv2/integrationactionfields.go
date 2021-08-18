package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Integrationactionfields
type Integrationactionfields struct { 
	// IntegrationAction - Reference to the Integration Action to be used when integrationAction type is qualified
	IntegrationAction *Integrationaction `json:"integrationAction,omitempty"`


	// RequestMappings - Collection of Request Mappings to use
	RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`

}

func (u *Integrationactionfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Integrationactionfields

	

	return json.Marshal(&struct { 
		IntegrationAction *Integrationaction `json:"integrationAction,omitempty"`
		
		RequestMappings *[]Requestmapping `json:"requestMappings,omitempty"`
		*Alias
	}{ 
		IntegrationAction: u.IntegrationAction,
		
		RequestMappings: u.RequestMappings,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Integrationactionfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
