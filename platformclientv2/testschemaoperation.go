package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testschemaoperation - Information about the Trigger test mode schema validation step
type Testschemaoperation struct { 
	// Name - The name of the processing step
	Name *string `json:"name,omitempty"`


	// Step - The number of the processing step
	Step *int `json:"step,omitempty"`


	// Matches - Whether or not the operation matches expectations
	Matches *bool `json:"matches,omitempty"`


	// Details - Details about why the operation did or did not succeed
	Details *[]string `json:"details,omitempty"`

}

func (o *Testschemaoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testschemaoperation
	
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

func (o *Testschemaoperation) UnmarshalJSON(b []byte) error {
	var TestschemaoperationMap map[string]interface{}
	err := json.Unmarshal(b, &TestschemaoperationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TestschemaoperationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Step, ok := TestschemaoperationMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if Matches, ok := TestschemaoperationMap["matches"].(bool); ok {
		o.Matches = &Matches
	}
    
	if Details, ok := TestschemaoperationMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testschemaoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
