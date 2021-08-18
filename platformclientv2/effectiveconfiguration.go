package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Effectiveconfiguration - Effective Configuration for an ClientApp. This is comprised of the integration specific configuration along with overrides specified in the integration type.
type Effectiveconfiguration struct { 
	// Properties - Key-value configuration settings described by the schema in the propertiesSchemaUri field.
	Properties *map[string]interface{} `json:"properties,omitempty"`


	// Advanced - Advanced configuration described by the schema in the advancedSchemaUri field.
	Advanced *map[string]interface{} `json:"advanced,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// Notes - Notes about the integration.
	Notes *string `json:"notes,omitempty"`


	// Credentials - Credentials required by the integration. The required keys are indicated in the credentials property of the Integration Type
	Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`

}

func (u *Effectiveconfiguration) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Effectiveconfiguration

	

	return json.Marshal(&struct { 
		Properties *map[string]interface{} `json:"properties,omitempty"`
		
		Advanced *map[string]interface{} `json:"advanced,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Credentials *map[string]Credentialinfo `json:"credentials,omitempty"`
		*Alias
	}{ 
		Properties: u.Properties,
		
		Advanced: u.Advanced,
		
		Name: u.Name,
		
		Notes: u.Notes,
		
		Credentials: u.Credentials,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Effectiveconfiguration) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
