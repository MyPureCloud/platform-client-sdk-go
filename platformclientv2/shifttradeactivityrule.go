package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Shifttradeactivityrule
type Shifttradeactivityrule struct { 
	// ActivityCategory - The activity category to which to apply this rule
	ActivityCategory *string `json:"activityCategory,omitempty"`


	// Action - The action this rule invokes
	Action *string `json:"action,omitempty"`


	// ActivityCodeIdReplacement - The activity code ID with which to replace activities belonging to the original category if applicable (required if action == Replace, must be a default activity code ID)
	ActivityCodeIdReplacement *string `json:"activityCodeIdReplacement,omitempty"`

}

func (o *Shifttradeactivityrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Shifttradeactivityrule
	
	return json.Marshal(&struct { 
		ActivityCategory *string `json:"activityCategory,omitempty"`
		
		Action *string `json:"action,omitempty"`
		
		ActivityCodeIdReplacement *string `json:"activityCodeIdReplacement,omitempty"`
		*Alias
	}{ 
		ActivityCategory: o.ActivityCategory,
		
		Action: o.Action,
		
		ActivityCodeIdReplacement: o.ActivityCodeIdReplacement,
		Alias:    (*Alias)(o),
	})
}

func (o *Shifttradeactivityrule) UnmarshalJSON(b []byte) error {
	var ShifttradeactivityruleMap map[string]interface{}
	err := json.Unmarshal(b, &ShifttradeactivityruleMap)
	if err != nil {
		return err
	}
	
	if ActivityCategory, ok := ShifttradeactivityruleMap["activityCategory"].(string); ok {
		o.ActivityCategory = &ActivityCategory
	}
    
	if Action, ok := ShifttradeactivityruleMap["action"].(string); ok {
		o.Action = &Action
	}
    
	if ActivityCodeIdReplacement, ok := ShifttradeactivityruleMap["activityCodeIdReplacement"].(string); ok {
		o.ActivityCodeIdReplacement = &ActivityCodeIdReplacement
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Shifttradeactivityrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
