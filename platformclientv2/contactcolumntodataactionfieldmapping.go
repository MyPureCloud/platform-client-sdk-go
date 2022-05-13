package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactcolumntodataactionfieldmapping
type Contactcolumntodataactionfieldmapping struct { 
	// ContactColumnName - The name of a contact column whose data will be passed to the data action
	ContactColumnName *string `json:"contactColumnName,omitempty"`


	// DataActionField - The name of an input field from the data action that the contact column data will be passed to
	DataActionField *string `json:"dataActionField,omitempty"`

}

func (o *Contactcolumntodataactionfieldmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Contactcolumntodataactionfieldmapping
	
	return json.Marshal(&struct { 
		ContactColumnName *string `json:"contactColumnName,omitempty"`
		
		DataActionField *string `json:"dataActionField,omitempty"`
		*Alias
	}{ 
		ContactColumnName: o.ContactColumnName,
		
		DataActionField: o.DataActionField,
		Alias:    (*Alias)(o),
	})
}

func (o *Contactcolumntodataactionfieldmapping) UnmarshalJSON(b []byte) error {
	var ContactcolumntodataactionfieldmappingMap map[string]interface{}
	err := json.Unmarshal(b, &ContactcolumntodataactionfieldmappingMap)
	if err != nil {
		return err
	}
	
	if ContactColumnName, ok := ContactcolumntodataactionfieldmappingMap["contactColumnName"].(string); ok {
		o.ContactColumnName = &ContactColumnName
	}
    
	if DataActionField, ok := ContactcolumntodataactionfieldmappingMap["dataActionField"].(string); ok {
		o.DataActionField = &DataActionField
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
