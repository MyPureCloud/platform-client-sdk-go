package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testtargetoperation - Information about the Trigger test mode target validation step
type Testtargetoperation struct { 
	// Name - The name of the processing step
	Name *string `json:"name,omitempty"`


	// Step - The number of the processing step
	Step *int `json:"step,omitempty"`


	// Matches - Whether or not the operation matches expectations
	Matches *bool `json:"matches,omitempty"`


	// Details - Details about why the operation did or did not succeed
	Details *[]string `json:"details,omitempty"`

}

func (o *Testtargetoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testtargetoperation
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Step *int `json:"step,omitempty"`
		
		Matches *bool `json:"matches,omitempty"`
		
		Details *[]string `json:"details,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Step: o.Step,
		
		Matches: o.Matches,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Testtargetoperation) UnmarshalJSON(b []byte) error {
	var TesttargetoperationMap map[string]interface{}
	err := json.Unmarshal(b, &TesttargetoperationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TesttargetoperationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Step, ok := TesttargetoperationMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if Matches, ok := TesttargetoperationMap["matches"].(bool); ok {
		o.Matches = &Matches
	}
    
	if Details, ok := TesttargetoperationMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testtargetoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
