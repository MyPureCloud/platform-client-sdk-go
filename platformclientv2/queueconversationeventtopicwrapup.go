package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Queueconversationeventtopicwrapup
type Queueconversationeventtopicwrapup struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Notes
	Notes *string `json:"notes,omitempty"`


	// Tags
	Tags *[]string `json:"tags,omitempty"`


	// DurationSeconds
	DurationSeconds *int `json:"durationSeconds,omitempty"`


	// EndTime
	EndTime *time.Time `json:"endTime,omitempty"`


	// AdditionalProperties
	AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`

}

func (o *Queueconversationeventtopicwrapup) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Queueconversationeventtopicwrapup
	
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
		
		AdditionalProperties *interface{} `json:"additionalProperties,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Notes: o.Notes,
		
		Tags: o.Tags,
		
		DurationSeconds: o.DurationSeconds,
		
		EndTime: EndTime,
		
		AdditionalProperties: o.AdditionalProperties,
		Alias:    (*Alias)(o),
	})
}

func (o *Queueconversationeventtopicwrapup) UnmarshalJSON(b []byte) error {
	var QueueconversationeventtopicwrapupMap map[string]interface{}
	err := json.Unmarshal(b, &QueueconversationeventtopicwrapupMap)
	if err != nil {
		return err
	}
	
	if Code, ok := QueueconversationeventtopicwrapupMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Notes, ok := QueueconversationeventtopicwrapupMap["notes"].(string); ok {
		o.Notes = &Notes
	}
	
	if Tags, ok := QueueconversationeventtopicwrapupMap["tags"].([]interface{}); ok {
		TagsString, _ := json.Marshal(Tags)
		json.Unmarshal(TagsString, &o.Tags)
	}
	
	if DurationSeconds, ok := QueueconversationeventtopicwrapupMap["durationSeconds"].(float64); ok {
		DurationSecondsInt := int(DurationSeconds)
		o.DurationSeconds = &DurationSecondsInt
	}
	
	if endTimeString, ok := QueueconversationeventtopicwrapupMap["endTime"].(string); ok {
		EndTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endTimeString)
		o.EndTime = &EndTime
	}
	
	if AdditionalProperties, ok := QueueconversationeventtopicwrapupMap["additionalProperties"].(map[string]interface{}); ok {
		AdditionalPropertiesString, _ := json.Marshal(AdditionalProperties)
		json.Unmarshal(AdditionalPropertiesString, &o.AdditionalProperties)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Queueconversationeventtopicwrapup) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
