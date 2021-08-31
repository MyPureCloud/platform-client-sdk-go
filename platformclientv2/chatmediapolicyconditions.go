package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Chatmediapolicyconditions
type Chatmediapolicyconditions struct { 
	// ForUsers
	ForUsers *[]User `json:"forUsers,omitempty"`


	// DateRanges
	DateRanges *[]string `json:"dateRanges,omitempty"`


	// ForQueues
	ForQueues *[]Queue `json:"forQueues,omitempty"`


	// WrapupCodes
	WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`


	// Languages
	Languages *[]Language `json:"languages,omitempty"`


	// TimeAllowed
	TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`


	// Duration
	Duration *Durationcondition `json:"duration,omitempty"`

}

func (o *Chatmediapolicyconditions) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Chatmediapolicyconditions
	
	return json.Marshal(&struct { 
		ForUsers *[]User `json:"forUsers,omitempty"`
		
		DateRanges *[]string `json:"dateRanges,omitempty"`
		
		ForQueues *[]Queue `json:"forQueues,omitempty"`
		
		WrapupCodes *[]Wrapupcode `json:"wrapupCodes,omitempty"`
		
		Languages *[]Language `json:"languages,omitempty"`
		
		TimeAllowed *Timeallowed `json:"timeAllowed,omitempty"`
		
		Duration *Durationcondition `json:"duration,omitempty"`
		*Alias
	}{ 
		ForUsers: o.ForUsers,
		
		DateRanges: o.DateRanges,
		
		ForQueues: o.ForQueues,
		
		WrapupCodes: o.WrapupCodes,
		
		Languages: o.Languages,
		
		TimeAllowed: o.TimeAllowed,
		
		Duration: o.Duration,
		Alias:    (*Alias)(o),
	})
}

func (o *Chatmediapolicyconditions) UnmarshalJSON(b []byte) error {
	var ChatmediapolicyconditionsMap map[string]interface{}
	err := json.Unmarshal(b, &ChatmediapolicyconditionsMap)
	if err != nil {
		return err
	}
	
	if ForUsers, ok := ChatmediapolicyconditionsMap["forUsers"].([]interface{}); ok {
		ForUsersString, _ := json.Marshal(ForUsers)
		json.Unmarshal(ForUsersString, &o.ForUsers)
	}
	
	if DateRanges, ok := ChatmediapolicyconditionsMap["dateRanges"].([]interface{}); ok {
		DateRangesString, _ := json.Marshal(DateRanges)
		json.Unmarshal(DateRangesString, &o.DateRanges)
	}
	
	if ForQueues, ok := ChatmediapolicyconditionsMap["forQueues"].([]interface{}); ok {
		ForQueuesString, _ := json.Marshal(ForQueues)
		json.Unmarshal(ForQueuesString, &o.ForQueues)
	}
	
	if WrapupCodes, ok := ChatmediapolicyconditionsMap["wrapupCodes"].([]interface{}); ok {
		WrapupCodesString, _ := json.Marshal(WrapupCodes)
		json.Unmarshal(WrapupCodesString, &o.WrapupCodes)
	}
	
	if Languages, ok := ChatmediapolicyconditionsMap["languages"].([]interface{}); ok {
		LanguagesString, _ := json.Marshal(Languages)
		json.Unmarshal(LanguagesString, &o.Languages)
	}
	
	if TimeAllowed, ok := ChatmediapolicyconditionsMap["timeAllowed"].(map[string]interface{}); ok {
		TimeAllowedString, _ := json.Marshal(TimeAllowed)
		json.Unmarshal(TimeAllowedString, &o.TimeAllowed)
	}
	
	if Duration, ok := ChatmediapolicyconditionsMap["duration"].(map[string]interface{}); ok {
		DurationString, _ := json.Marshal(Duration)
		json.Unmarshal(DurationString, &o.Duration)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Chatmediapolicyconditions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
