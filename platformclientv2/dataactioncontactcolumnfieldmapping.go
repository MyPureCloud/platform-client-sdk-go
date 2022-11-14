package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dataactioncontactcolumnfieldmapping
type Dataactioncontactcolumnfieldmapping struct { 
	// ContactColumnName - The name of a contact column whose data will be passed to the data action
	ContactColumnName *string `json:"contactColumnName,omitempty"`


	// DataActionField - The name of an input field from the data action that the contact column data will be passed to
	DataActionField *string `json:"dataActionField,omitempty"`

}

func (o *Dataactioncontactcolumnfieldmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dataactioncontactcolumnfieldmapping
	
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

func (o *Dataactioncontactcolumnfieldmapping) UnmarshalJSON(b []byte) error {
	var DataactioncontactcolumnfieldmappingMap map[string]interface{}
	err := json.Unmarshal(b, &DataactioncontactcolumnfieldmappingMap)
	if err != nil {
		return err
	}
	
	if ContactColumnName, ok := DataactioncontactcolumnfieldmappingMap["contactColumnName"].(string); ok {
		o.ContactColumnName = &ContactColumnName
	}
    
	if DataActionField, ok := DataactioncontactcolumnfieldmappingMap["dataActionField"].(string); ok {
		o.DataActionField = &DataActionField
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Dataactioncontactcolumnfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
