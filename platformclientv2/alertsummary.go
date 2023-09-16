package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Alertsummary
type Alertsummary struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Entities - The entities who violated the rule condition over the duration of the alert.
	Entities *[]Alertsummaryentity `json:"entities,omitempty"`

	// Conversation - The id of the conversation that triggered the alert.  Only used for alerts based on instance-based conversation metrics.
	Conversation *Addressableentityref `json:"conversation,omitempty"`

	// MetricType - The metric type that is monitored.
	MetricType *string `json:"metricType,omitempty"`

	// EntitiesAreTeamMembers - Flag that indicated whether or not the alert is for a rule with a condition for all members of a team.
	EntitiesAreTeamMembers *bool `json:"entitiesAreTeamMembers,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Alertsummary) SetField(field string, fieldValue interface{}) {
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

func (o Alertsummary) MarshalJSON() ([]byte, error) {
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
	type Alias Alertsummary
	
	return json.Marshal(&struct { 
		Entities *[]Alertsummaryentity `json:"entities,omitempty"`
		
		Conversation *Addressableentityref `json:"conversation,omitempty"`
		
		MetricType *string `json:"metricType,omitempty"`
		
		EntitiesAreTeamMembers *bool `json:"entitiesAreTeamMembers,omitempty"`
		Alias
	}{ 
		Entities: o.Entities,
		
		Conversation: o.Conversation,
		
		MetricType: o.MetricType,
		
		EntitiesAreTeamMembers: o.EntitiesAreTeamMembers,
		Alias:    (Alias)(o),
	})
}

func (o *Alertsummary) UnmarshalJSON(b []byte) error {
	var AlertsummaryMap map[string]interface{}
	err := json.Unmarshal(b, &AlertsummaryMap)
	if err != nil {
		return err
	}
	
	if Entities, ok := AlertsummaryMap["entities"].([]interface{}); ok {
		EntitiesString, _ := json.Marshal(Entities)
		json.Unmarshal(EntitiesString, &o.Entities)
	}
	
	if Conversation, ok := AlertsummaryMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if MetricType, ok := AlertsummaryMap["metricType"].(string); ok {
		o.MetricType = &MetricType
	}
    
	if EntitiesAreTeamMembers, ok := AlertsummaryMap["entitiesAreTeamMembers"].(bool); ok {
		o.EntitiesAreTeamMembers = &EntitiesAreTeamMembers
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Alertsummary) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
