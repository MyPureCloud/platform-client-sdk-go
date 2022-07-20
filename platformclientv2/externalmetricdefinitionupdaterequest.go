package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Externalmetricdefinitionupdaterequest
type Externalmetricdefinitionupdaterequest struct { 
	// Name - The name of the External Metric Definition
	Name *string `json:"name,omitempty"`


	// Precision - The decimal precision of the External Metric Definition. Must be at least 0 and at most 5
	Precision *int `json:"precision,omitempty"`


	// DefaultObjectiveType - The default objective type of the External Metric Definition
	DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`


	// Enabled - True if the External Metric Definition is enabled
	Enabled *bool `json:"enabled,omitempty"`

}

func (o *Externalmetricdefinitionupdaterequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Externalmetricdefinitionupdaterequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Precision *int `json:"precision,omitempty"`
		
		DefaultObjectiveType *string `json:"defaultObjectiveType,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Precision: o.Precision,
		
		DefaultObjectiveType: o.DefaultObjectiveType,
		
		Enabled: o.Enabled,
		Alias:    (*Alias)(o),
	})
}

func (o *Externalmetricdefinitionupdaterequest) UnmarshalJSON(b []byte) error {
	var ExternalmetricdefinitionupdaterequestMap map[string]interface{}
	err := json.Unmarshal(b, &ExternalmetricdefinitionupdaterequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := ExternalmetricdefinitionupdaterequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Precision, ok := ExternalmetricdefinitionupdaterequestMap["precision"].(float64); ok {
		PrecisionInt := int(Precision)
		o.Precision = &PrecisionInt
	}
	
	if DefaultObjectiveType, ok := ExternalmetricdefinitionupdaterequestMap["defaultObjectiveType"].(string); ok {
		o.DefaultObjectiveType = &DefaultObjectiveType
	}
    
	if Enabled, ok := ExternalmetricdefinitionupdaterequestMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Externalmetricdefinitionupdaterequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
