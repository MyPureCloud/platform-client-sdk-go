package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Longtermrequirementsservicegoal
type Longtermrequirementsservicegoal struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// UseAsaTarget - Toggle for target average speed of answer from service goals
	UseAsaTarget *bool `json:"useAsaTarget,omitempty"`

	// TargetAsaSec - Value of target average speed of answer from service goals
	TargetAsaSec *float64 `json:"targetAsaSec,omitempty"`

	// UseServiceLevelTarget - Toggle for service level objective from service goals
	UseServiceLevelTarget *bool `json:"useServiceLevelTarget,omitempty"`

	// ServiceLevelObjectiveSeconds - Value of service level objective seconds from service goals
	ServiceLevelObjectiveSeconds *float64 `json:"serviceLevelObjectiveSeconds,omitempty"`

	// ServiceLevelGoalPercent - Value of service level objective percent from service goals
	ServiceLevelGoalPercent *float64 `json:"serviceLevelGoalPercent,omitempty"`

	// UseAbandonRateGoal - Toggle for abandon rate from service goals
	UseAbandonRateGoal *bool `json:"useAbandonRateGoal,omitempty"`

	// AbandonRateGoalPercent - Value of abandon rate objective
	AbandonRateGoalPercent *float64 `json:"abandonRateGoalPercent,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Longtermrequirementsservicegoal) SetField(field string, fieldValue interface{}) {
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

func (o Longtermrequirementsservicegoal) MarshalJSON() ([]byte, error) {
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
	type Alias Longtermrequirementsservicegoal
	
	return json.Marshal(&struct { 
		UseAsaTarget *bool `json:"useAsaTarget,omitempty"`
		
		TargetAsaSec *float64 `json:"targetAsaSec,omitempty"`
		
		UseServiceLevelTarget *bool `json:"useServiceLevelTarget,omitempty"`
		
		ServiceLevelObjectiveSeconds *float64 `json:"serviceLevelObjectiveSeconds,omitempty"`
		
		ServiceLevelGoalPercent *float64 `json:"serviceLevelGoalPercent,omitempty"`
		
		UseAbandonRateGoal *bool `json:"useAbandonRateGoal,omitempty"`
		
		AbandonRateGoalPercent *float64 `json:"abandonRateGoalPercent,omitempty"`
		Alias
	}{ 
		UseAsaTarget: o.UseAsaTarget,
		
		TargetAsaSec: o.TargetAsaSec,
		
		UseServiceLevelTarget: o.UseServiceLevelTarget,
		
		ServiceLevelObjectiveSeconds: o.ServiceLevelObjectiveSeconds,
		
		ServiceLevelGoalPercent: o.ServiceLevelGoalPercent,
		
		UseAbandonRateGoal: o.UseAbandonRateGoal,
		
		AbandonRateGoalPercent: o.AbandonRateGoalPercent,
		Alias:    (Alias)(o),
	})
}

func (o *Longtermrequirementsservicegoal) UnmarshalJSON(b []byte) error {
	var LongtermrequirementsservicegoalMap map[string]interface{}
	err := json.Unmarshal(b, &LongtermrequirementsservicegoalMap)
	if err != nil {
		return err
	}
	
	if UseAsaTarget, ok := LongtermrequirementsservicegoalMap["useAsaTarget"].(bool); ok {
		o.UseAsaTarget = &UseAsaTarget
	}
    
	if TargetAsaSec, ok := LongtermrequirementsservicegoalMap["targetAsaSec"].(float64); ok {
		o.TargetAsaSec = &TargetAsaSec
	}
    
	if UseServiceLevelTarget, ok := LongtermrequirementsservicegoalMap["useServiceLevelTarget"].(bool); ok {
		o.UseServiceLevelTarget = &UseServiceLevelTarget
	}
    
	if ServiceLevelObjectiveSeconds, ok := LongtermrequirementsservicegoalMap["serviceLevelObjectiveSeconds"].(float64); ok {
		o.ServiceLevelObjectiveSeconds = &ServiceLevelObjectiveSeconds
	}
    
	if ServiceLevelGoalPercent, ok := LongtermrequirementsservicegoalMap["serviceLevelGoalPercent"].(float64); ok {
		o.ServiceLevelGoalPercent = &ServiceLevelGoalPercent
	}
    
	if UseAbandonRateGoal, ok := LongtermrequirementsservicegoalMap["useAbandonRateGoal"].(bool); ok {
		o.UseAbandonRateGoal = &UseAbandonRateGoal
	}
    
	if AbandonRateGoalPercent, ok := LongtermrequirementsservicegoalMap["abandonRateGoalPercent"].(float64); ok {
		o.AbandonRateGoalPercent = &AbandonRateGoalPercent
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Longtermrequirementsservicegoal) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
