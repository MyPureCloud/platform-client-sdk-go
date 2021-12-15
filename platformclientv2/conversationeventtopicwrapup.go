package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Conversationeventtopicwrapup - Call wrap up or disposition data.
type Conversationeventtopicwrapup struct { 
	// Code - The user configured wrap up code name.
	Code *string `json:"code,omitempty"`


	// Notes - Text entered by the agent to describe the call or disposition.
	Notes *string `json:"notes,omitempty"`


	// Tags - List of tags selected by the agent to describe the call or disposition.
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds - The length of time in seconds that the agent spent doing after call work., Note, the format of utc-millisec should be ignored, our code generator needs it to generate a Long for us internally
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime - The timestamp when the wrapup was finished.
	EndTime *time.Time `json:"endTime,omitempty"`

}

func (o *Conversationeventtopicwrapup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Conversationeventtopicwrapup
	
	EndTime := new(string)
	if o.EndTime != nil {
		
		*EndTime = timeutil.Strftime(o.EndTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndTime = nil
	}
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Notes *string `json:"notes,omitempty"`
		
		Tags *[]string `json:"tags,omitempty"`
		
		DurationSeconds *int `json:"durationSeconds,omitempty"`
		
		EndTime *string `json:"endTime,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Notes: o.Notes,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		EndTime: EndTime,
		Alias:    (*Alias)(o),
	})
}

func (o *Conversationeventtopicwrapup) UnmarshalJSON(b []byte) error {
	var ConversationeventtopicwrapupMap map[string]interface{}
	err := json.Unmarshal(b, &ConversationeventtopicwrapupMap)
	if err != nil {
		return err
	}
	
	if Code, ok := ConversationeventtopicwrapupMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Notes, ok := ConversationeventtopicwrapupMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if Tags, ok := ConversationeventtopicwrapupMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := ConversationeventtopicwrapupMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if endTimeString, ok := ConversationeventtopicwrapupMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Conversationeventtopicwrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
