package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Createintegrationrequest - Details for an Integration
type Createintegrationrequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - The name of the integration, used to distinguish this integration from others of the same type.
	Name *string `json:"name,omitempty"`


	// IntegrationType - Type of the integration to create.
	IntegrationType *Integrationtype `json:"integrationType,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Createintegrationrequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Createintegrationrequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		IntegrationType *Integrationtype `json:"integrationType,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		IntegrationType: o.IntegrationType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Createintegrationrequest) UnmarshalJSON(b []byte) error {
	var CreateintegrationrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateintegrationrequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CreateintegrationrequestMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := CreateintegrationrequestMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if IntegrationType, ok := CreateintegrationrequestMap["integrationType"].(map[string]interface{}); ok {
		IntegrationTypeString, _ := json.Marshal(IntegrationType)
		json.Unmarshal(IntegrationTypeString, &o.IntegrationType)
	}
	
	if SelfUri, ok := CreateintegrationrequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createintegrationrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
