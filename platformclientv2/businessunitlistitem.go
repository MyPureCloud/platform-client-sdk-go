package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Businessunitlistitem
type Businessunitlistitem struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Authorized - Whether the user has authorization to interact with this business unit
	Authorized *bool `json:"authorized,omitempty"`


	// Division - The division to which this entity belongs.
	Division *Divisionreference `json:"division,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Businessunitlistitem) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Businessunitlistitem
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Authorized *bool `json:"authorized,omitempty"`
		
		Division *Divisionreference `json:"division,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Authorized: o.Authorized,
		
		Division: o.Division,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Businessunitlistitem) UnmarshalJSON(b []byte) error {
	var BusinessunitlistitemMap map[string]interface{}
	err := json.Unmarshal(b, &BusinessunitlistitemMap)
	if err != nil {
		return err
	}
	
	if Id, ok := BusinessunitlistitemMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := BusinessunitlistitemMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Authorized, ok := BusinessunitlistitemMap["authorized"].(bool); ok {
		o.Authorized = &Authorized
	}
	
	if Division, ok := BusinessunitlistitemMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if SelfUri, ok := BusinessunitlistitemMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Businessunitlistitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
