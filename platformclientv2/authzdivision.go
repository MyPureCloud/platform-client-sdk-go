package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Authzdivision
type Authzdivision struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Description - A helpful description for the division.
	Description *string `json:"description,omitempty"`


	// HomeDivision - A flag indicating whether this division is the \"Home\" (default) division. Cannot be modified and any supplied value will be ignored on create or update.
	HomeDivision *bool `json:"homeDivision,omitempty"`


	// ObjectCounts - A count of objects in this division, grouped by type.
	ObjectCounts *map[string]int `json:"objectCounts,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Authzdivision) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Authzdivision
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		HomeDivision *bool `json:"homeDivision,omitempty"`
		
		ObjectCounts *map[string]int `json:"objectCounts,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Description: o.Description,
		
		HomeDivision: o.HomeDivision,
		
		ObjectCounts: o.ObjectCounts,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Authzdivision) UnmarshalJSON(b []byte) error {
	var AuthzdivisionMap map[string]interface{}
	err := json.Unmarshal(b, &AuthzdivisionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AuthzdivisionMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := AuthzdivisionMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Description, ok := AuthzdivisionMap["description"].(string); ok {
		o.Description = &Description
	}
	
	if HomeDivision, ok := AuthzdivisionMap["homeDivision"].(bool); ok {
		o.HomeDivision = &HomeDivision
	}
	
	if ObjectCounts, ok := AuthzdivisionMap["objectCounts"].(map[string]interface{}); ok {
		ObjectCountsString, _ := json.Marshal(ObjectCounts)
		json.Unmarshal(ObjectCountsString, &o.ObjectCounts)
	}
	
	if SelfUri, ok := AuthzdivisionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Authzdivision) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
