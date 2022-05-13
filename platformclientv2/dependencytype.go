package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dependencytype
type Dependencytype struct { 
	// Id - The dependency type identifier
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Versioned
	Versioned *bool `json:"versioned,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Dependencytype) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dependencytype
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Versioned *bool `json:"versioned,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Versioned: o.Versioned,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Dependencytype) UnmarshalJSON(b []byte) error {
	var DependencytypeMap map[string]interface{}
	err := json.Unmarshal(b, &DependencytypeMap)
	if err != nil {
		return err
	}
	
	if Id, ok := DependencytypeMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := DependencytypeMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Versioned, ok := DependencytypeMap["versioned"].(bool); ok {
		o.Versioned = &Versioned
	}
    
	if SelfUri, ok := DependencytypeMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dependencytype) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
