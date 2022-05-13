package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Workplanreference - Work plan information
type Workplanreference struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// ManagementUnit - The management unit to which this work plan belongs.  Nullable in some routes
	ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Workplanreference) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Workplanreference
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		ManagementUnit *Managementunitreference `json:"managementUnit,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		ManagementUnit: o.ManagementUnit,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Workplanreference) UnmarshalJSON(b []byte) error {
	var WorkplanreferenceMap map[string]interface{}
	err := json.Unmarshal(b, &WorkplanreferenceMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WorkplanreferenceMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if ManagementUnit, ok := WorkplanreferenceMap["managementUnit"].(map[string]interface{}); ok {
		ManagementUnitString, _ := json.Marshal(ManagementUnit)
		json.Unmarshal(ManagementUnitString, &o.ManagementUnit)
	}
	
	if SelfUri, ok := WorkplanreferenceMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Workplanreference) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
