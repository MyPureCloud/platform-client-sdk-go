package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Validateworkplanresponse
type Validateworkplanresponse struct { 
	// WorkPlan - The work plan reference associated with this response
	WorkPlan *Workplanreference `json:"workPlan,omitempty"`


	// Valid - Whether the work plan is valid or not
	Valid *bool `json:"valid,omitempty"`


	// Messages - Validation messages for this work plan
	Messages *Validateworkplanmessages `json:"messages,omitempty"`

}

func (o *Validateworkplanresponse) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Validateworkplanresponse
	
	return json.Marshal(&struct { 
		WorkPlan *Workplanreference `json:"workPlan,omitempty"`
		
		Valid *bool `json:"valid,omitempty"`
		
		Messages *Validateworkplanmessages `json:"messages,omitempty"`
		*Alias
	}{ 
		WorkPlan: o.WorkPlan,
		
		Valid: o.Valid,
		
		Messages: o.Messages,
		Alias:    (*Alias)(o),
	})
}

func (o *Validateworkplanresponse) UnmarshalJSON(b []byte) error {
	var ValidateworkplanresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ValidateworkplanresponseMap)
	if err != nil {
		return err
	}
	
	if WorkPlan, ok := ValidateworkplanresponseMap["workPlan"].(map[string]interface{}); ok {
		WorkPlanString, _ := json.Marshal(WorkPlan)
		json.Unmarshal(WorkPlanString, &o.WorkPlan)
	}
	
	if Valid, ok := ValidateworkplanresponseMap["valid"].(bool); ok {
		o.Valid = &Valid
	}
    
	if Messages, ok := ValidateworkplanresponseMap["messages"].(map[string]interface{}); ok {
		MessagesString, _ := json.Marshal(Messages)
		json.Unmarshal(MessagesString, &o.Messages)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Validateworkplanresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
