package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicalertruleproperties
type V2mobiusalertstopicalertruleproperties struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`

}

func (o *V2mobiusalertstopicalertruleproperties) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias V2mobiusalertstopicalertruleproperties
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		Alias:    (*Alias)(o),
	})
}

func (o *V2mobiusalertstopicalertruleproperties) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicalertrulepropertiesMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicalertrulepropertiesMap)
	if err != nil {
		return err
	}
	
	if Id, ok := V2mobiusalertstopicalertrulepropertiesMap["id"].(map[string]interface{}); ok {
		IdString, _ := json.Marshal(Id)
		json.Unmarshal(IdString, &o.Id)
	}
	
	if Name, ok := V2mobiusalertstopicalertrulepropertiesMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := V2mobiusalertstopicalertrulepropertiesMap["type"].(string); ok {
		o.VarType = &VarType
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicalertruleproperties) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
