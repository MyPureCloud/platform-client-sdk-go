package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Routingconversationattributesrequest
type Routingconversationattributesrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Priority - Priority for the conversation.  Each point of priority is equivalent to one minute of time in queue.  Range:[-25000000, 25000000].  To reset, specify 0.
	Priority *int `json:"priority,omitempty"`

	// SkillIds - Skill requirements for the conversation.  To remove all skill requirements, specify an empty list, i.e. [].
	SkillIds *[]string `json:"skillIds,omitempty"`

	// LanguageId - Language requirement for the conversation.  To remove the language requirement, specify an empty string, i.e., \"\".
	LanguageId *string `json:"languageId,omitempty"`

	// RequestScoredAgents
	RequestScoredAgents *[]Requestscoredagent `json:"requestScoredAgents,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Routingconversationattributesrequest) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Routingconversationattributesrequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Routingconversationattributesrequest
	
	return json.Marshal(&struct { 
		Priority *int `json:"priority,omitempty"`
		
		SkillIds *[]string `json:"skillIds,omitempty"`
		
		LanguageId *string `json:"languageId,omitempty"`
		
		RequestScoredAgents *[]Requestscoredagent `json:"requestScoredAgents,omitempty"`
		Alias
	}{ 
		Priority: o.Priority,
		
		SkillIds: o.SkillIds,
		
		LanguageId: o.LanguageId,
		
		RequestScoredAgents: o.RequestScoredAgents,
		Alias:    (Alias)(o),
	})
}

func (o *Routingconversationattributesrequest) UnmarshalJSON(b []byte) error {
	var RoutingconversationattributesrequestMap map[string]interface{}
	err := json.Unmarshal(b, &RoutingconversationattributesrequestMap)
	if err != nil {
		return err
	}
	
	if Priority, ok := RoutingconversationattributesrequestMap["priority"].(float64); ok {
		PriorityInt := int(Priority)
		o.Priority = &PriorityInt
	}
	
	if SkillIds, ok := RoutingconversationattributesrequestMap["skillIds"].([]interface{}); ok {
		SkillIdsString, _ := json.Marshal(SkillIds)
		json.Unmarshal(SkillIdsString, &o.SkillIds)
	}
	
	if LanguageId, ok := RoutingconversationattributesrequestMap["languageId"].(string); ok {
		o.LanguageId = &LanguageId
	}
    
	if RequestScoredAgents, ok := RoutingconversationattributesrequestMap["requestScoredAgents"].([]interface{}); ok {
		RequestScoredAgentsString, _ := json.Marshal(RequestScoredAgents)
		json.Unmarshal(RequestScoredAgentsString, &o.RequestScoredAgents)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Routingconversationattributesrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
