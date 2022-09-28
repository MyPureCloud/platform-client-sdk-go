package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Predictormodelfeature
type Predictormodelfeature struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// VarType - The type of feature.
	VarType *string `json:"type,omitempty"`


	// PercentageImportance - The percentage of how important a feature is in the model.
	PercentageImportance *float64 `json:"percentageImportance,omitempty"`

}

func (o *Predictormodelfeature) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Predictormodelfeature
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		PercentageImportance *float64 `json:"percentageImportance,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		VarType: o.VarType,
		
		PercentageImportance: o.PercentageImportance,
		Alias:    (*Alias)(o),
	})
}

func (o *Predictormodelfeature) UnmarshalJSON(b []byte) error {
	var PredictormodelfeatureMap map[string]interface{}
	err := json.Unmarshal(b, &PredictormodelfeatureMap)
	if err != nil {
		return err
	}
	
	if Id, ok := PredictormodelfeatureMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if VarType, ok := PredictormodelfeatureMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if PercentageImportance, ok := PredictormodelfeatureMap["percentageImportance"].(float64); ok {
		o.PercentageImportance = &PercentageImportance
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Predictormodelfeature) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
