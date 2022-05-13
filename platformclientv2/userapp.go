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

func (o *Userapp) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationType: o.IntegrationType,
		
		Config: o.Config,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Userapp) UnmarshalJSON(b []byte) error {
	var UserappMap map[string]interface{}
	err := json.Unmarshal(b, &UserappMap)
	if err != nil {
		return err
	}
	
	if Id, ok := UserappMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := UserappMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if IntegrationType, ok := UserappMap["integrationType"].(map[string]interface{}); ok {
		IntegrationTypeString, _ := json.Marshal(IntegrationType)
		json.Unmarshal(IntegrationTypeString, &o.IntegrationType)
	}
	
	if Config, ok := UserappMap["config"].(map[string]interface{}); ok {
		ConfigString, _ := json.Marshal(Config)
		json.Unmarshal(ConfigString, &o.Config)
	}
	
	if SelfUri, ok := UserappMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Userapp) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
