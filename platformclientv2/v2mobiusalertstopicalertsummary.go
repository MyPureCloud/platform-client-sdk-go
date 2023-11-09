package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2mobiusalertstopicalertsummary
type V2mobiusalertstopicalertsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entities
	Entities *[]V2mobiusalertstopicalertsummaryentity `json:"entities,omitempty"`

	// Conversation
	Conversation *V2mobiusalertstopicaddressableentityref `json:"conversation,omitempty"`

	// MetricType
	MetricType *string `json:"metricType,omitempty"`

	// IsTeamMemberRuleSummary
	IsTeamMemberRuleSummary *bool `json:"isTeamMemberRuleSummary,omitempty"`

	// TeamMemberRuleSummary
	TeamMemberRuleSummary *bool `json:"teamMemberRuleSummary,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2mobiusalertstopicalertsummary) SetField(field string, fieldValue interface{}) {
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

func (o V2mobiusalertstopicalertsummary) MarshalJSON() ([]byte, error) {
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
	type Alias V2mobiusalertstopicalertsummary
	
	return json.Marshal(&struct { 
		Entities *[]V2mobiusalertstopicalertsummaryentity `json:"entities,omitempty"`
		
		Conversation *V2mobiusalertstopicaddressableentityref `json:"conversation,omitempty"`
		
		MetricType *string `json:"metricType,omitempty"`
		
		IsTeamMemberRuleSummary *bool `json:"isTeamMemberRuleSummary,omitempty"`
		
		TeamMemberRuleSummary *bool `json:"teamMemberRuleSummary,omitempty"`
		Alias
	}{ 
		Entities: o.Entities,
		
		Conversation: o.Conversation,
		
		MetricType: o.MetricType,
		
		IsTeamMemberRuleSummary: o.IsTeamMemberRuleSummary,
		
		TeamMemberRuleSummary: o.TeamMemberRuleSummary,
		Alias:    (Alias)(o),
	})
}

func (o *V2mobiusalertstopicalertsummary) UnmarshalJSON(b []byte) error {
	var V2mobiusalertstopicalertsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &V2mobiusalertstopicalertsummaryMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := V2mobiusalertstopicalertsummaryMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Conversation, ok := V2mobiusalertstopicalertsummaryMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if MetricType, ok := V2mobiusalertstopicalertsummaryMap["metricType"].(string); ok {
		o.MetricType = &MetricType
	}
    
	if IsTeamMemberRuleSummary, ok := V2mobiusalertstopicalertsummaryMap["isTeamMemberRuleSummary"].(bool); ok {
		o.IsTeamMemberRuleSummary = &IsTeamMemberRuleSummary
	}
    
	if TeamMemberRuleSummary, ok := V2mobiusalertstopicalertsummaryMap["teamMemberRuleSummary"].(bool); ok {
		o.TeamMemberRuleSummary = &TeamMemberRuleSummary
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *V2mobiusalertstopicalertsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
