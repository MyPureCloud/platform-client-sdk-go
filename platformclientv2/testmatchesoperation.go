package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Testmatchesoperation - Information about the Trigger test mode processing step
type Testmatchesoperation struct { 
	// Name - The name of the processing step
	Name *string `json:"name,omitempty"`


	// Step - The number of the processing step
	Step *int `json:"step,omitempty"`


	// Matches - Whether or not the operation matches expectations
	Matches *bool `json:"matches,omitempty"`


	// Details - Details about why the operation did or did not succeed
	Details *[]Matchcriteriatestresult `json:"details,omitempty"`

}

func (o *Testmatchesoperation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Testmatchesoperation
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Step *int `json:"step,omitempty"`
		
		Matches *bool `json:"matches,omitempty"`
		
		Details *[]Matchcriteriatestresult `json:"details,omitempty"`
		*Alias
	}{ 
		Name: o.Name,
		
		Step: o.Step,
		
		Matches: o.Matches,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Testmatchesoperation) UnmarshalJSON(b []byte) error {
	var TestmatchesoperationMap map[string]interface{}
	err := json.Unmarshal(b, &TestmatchesoperationMap)
	if err != nil {
		return err
	}
	
	if Name, ok := TestmatchesoperationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Step, ok := TestmatchesoperationMap["step"].(float64); ok {
		StepInt := int(Step)
		o.Step = &StepInt
	}
	
	if Matches, ok := TestmatchesoperationMap["matches"].(bool); ok {
		o.Matches = &Matches
	}
    
	if Details, ok := TestmatchesoperationMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Testmatchesoperation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
