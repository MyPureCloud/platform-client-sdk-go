package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdefinitioncreaterequest
type Externalmetricdefinitioncreaterequest struct { 
	// Name - The name of the External Metric Definition
	Name *string `json:"name,omitempty"`


	// Unit - The unit of the External Metric Definition
	Unit *string `json:"unit,omitempty"`


	// UnitDefinition - The unit definition of the External Metric Definition
	UnitDefinition *string `json:"unitDefinition,omitempty"`


	// Precision - The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
	Precision *int `json:"precision,omitempty"`


	// DefaultObjectiveType - The default objective type of the External Metric Definition
	DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`


	// Enabled - True if the External Metric Definition is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Externalmetricdefinitioncreaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdefinitioncreaterequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Unit *string `json:"unit,omitempty"`
		
		UnitDefinition *string `json:"unitDefinition,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Unit: o.Unit,
		
		UnitDefinition: o.UnitDefinition,
		
		Precision: o.Precision,
		
		DefaultObjectiveType: o.DefaultObjectiveType,
		
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdefinitioncreaterequest) UnmarshalJSON(b []byte) error {
	var ExternalmetricdefinitioncreaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdefinitioncreaterequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ExternalmetricdefinitioncreaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Unit, ok := ExternalmetricdefinitioncreaterequestMap["unit"].(string); ok {
		o.Unit = &Unit
	}
    
	if UnitDefinition, ok := ExternalmetricdefinitioncreaterequestMap["unitDefinition"].(string); ok {
		o.UnitDefinition = &UnitDefinition
	}
    
	if Precision, ok := ExternalmetricdefinitioncreaterequestMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if DefaultObjectiveType, ok := ExternalmetricdefinitioncreaterequestMap["defaultObjectiveType"].(string); ok {
		o.DefaultObjectiveType = &DefaultObjectiveType
	}
    
	if Enabled, ok := ExternalmetricdefinitioncreaterequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdefinitioncreaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
