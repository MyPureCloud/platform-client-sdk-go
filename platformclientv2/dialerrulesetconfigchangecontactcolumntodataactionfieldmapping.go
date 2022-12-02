package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping
type Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping struct { 
	// ContactColumnName - The name of a contact column whose data will be passed to the data action
	ContactColumnName *string `json:"contactColumnName,omitempty"`


	// DataActionField - The name of an output field from the data action that the contact column data will be passed to
	DataActionField *string `json:"dataActionField,omitempty"`


	// AdditionalProperties
	AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`

}

func (o *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping
	
	return json.Marshal(&struct { 
		ContactColumnName *string `json:"contactColumnName,omitempty"`
		
		DataActionField *string `json:"dataActionField,omitempty"`
		
		AdditionalProperties *map[string]interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		ContactColumnName: o.ContactColumnName,
		
		DataActionField: o.DataActionField,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) UnmarshalJSON(b []byte) error {
	var DialerrulesetconfigchangecontactcolumntodataactionfieldmappingMap map[string]interface{}
	err := json.Unmarshal(b, &DialerrulesetconfigchangecontactcolumntodataactionfieldmappingMap)
	if err != nil {
		return err
	}
	
	if ContactColumnName, ok := DialerrulesetconfigchangecontactcolumntodataactionfieldmappingMap["contactColumnName"].(string); ok {
		o.ContactColumnName = &ContactColumnName
	}
    
	if DataActionField, ok := DialerrulesetconfigchangecontactcolumntodataactionfieldmappingMap["dataActionField"].(string); ok {
		o.DataActionField = &DataActionField
	}
    
	if AdditionalProperties, ok := DialerrulesetconfigchangecontactcolumntodataactionfieldmappingMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerrulesetconfigchangecontactcolumntodataactionfieldmapping) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
