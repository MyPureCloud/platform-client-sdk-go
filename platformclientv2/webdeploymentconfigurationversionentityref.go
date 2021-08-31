package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webdeploymentconfigurationversionentityref
type Webdeploymentconfigurationversionentityref struct { 
	// Id - The configuration version ID
	Id *string `json:"id,omitempty"`


	// Name - The configuration version name
	Name *string `json:"name,omitempty"`


	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`


	// Version - The version of the configuration
	Version *string `json:"version,omitempty"`

}

func (o *Webdeploymentconfigurationversionentityref) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webdeploymentconfigurationversionentityref
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		Version *string `json:"version,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		SelfUri: o.SelfUri,
		
		Version: o.Version,
		Alias:    (*Alias)(o),
	})
}

func (o *Webdeploymentconfigurationversionentityref) UnmarshalJSON(b []byte) error {
	var WebdeploymentconfigurationversionentityrefMap map[string]interface{}
	err := json.Unmarshal(b, &WebdeploymentconfigurationversionentityrefMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebdeploymentconfigurationversionentityrefMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WebdeploymentconfigurationversionentityrefMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if SelfUri, ok := WebdeploymentconfigurationversionentityrefMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	
	if Version, ok := WebdeploymentconfigurationversionentityrefMap["version"].(string); ok {
		o.Version = &Version
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webdeploymentconfigurationversionentityref) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
