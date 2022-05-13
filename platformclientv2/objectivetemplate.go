package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Objectivetemplate
type Objectivetemplate struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Zones
	Zones *[]Objectivezone `json:"zones,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Objectivetemplate) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Objectivetemplate
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Zones *[]Objectivezone `json:"zones,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Zones: o.Zones,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Objectivetemplate) UnmarshalJSON(b []byte) error {
	var ObjectivetemplateMap map[string]interface{}
	err := json.Unmarshal(b, &ObjectivetemplateMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ObjectivetemplateMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ObjectivetemplateMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Zones, ok := ObjectivetemplateMap["zones"].([]interface{}); ok {
		ZonesString, _ := json.Marshal(Zones)
		json.Unmarshal(ZonesString, &o.Zones)
	}
	
	if SelfUri, ok := ObjectivetemplateMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Objectivetemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
