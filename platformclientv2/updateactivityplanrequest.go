package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Updateactivityplanrequest
type Updateactivityplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the activity plan
	Name *string `json:"name,omitempty"`

	// Description - The description of the activity plan
	Description *string `json:"description,omitempty"`

	// GroupSettings - Group settings for the activity plan
	GroupSettings *Valuewrappergroupsettings `json:"groupSettings,omitempty"`

	// AttendeesSearchRule - Attendee search rule for this activity plan
	AttendeesSearchRule *Valuewrapperusersearchrule `json:"attendeesSearchRule,omitempty"`

	// FacilitatorsSearchRule - Facilitator search rule for this activity plan
	FacilitatorsSearchRule *Valuewrapperusersearchrule `json:"facilitatorsSearchRule,omitempty"`

	// TransitionTimeMinutes - Transition time in minutes between facilitated sessions
	TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`

	// ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan
	ServiceGoalImpactOverrides *Valuewrapperactivityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`

	// OptimizationObjective - The optimization objective of this activity plan
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// State - The state of this activity plan
	State *string `json:"state,omitempty"`

	// FixedAvailability - Fixed availability configuration for the activity plan
	FixedAvailability *Listwrapperfixedavailability `json:"fixedAvailability,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Updateactivityplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Updateactivityplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Updateactivityplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		GroupSettings *Valuewrappergroupsettings `json:"groupSettings,omitempty"`
		
		AttendeesSearchRule *Valuewrapperusersearchrule `json:"attendeesSearchRule,omitempty"`
		
		FacilitatorsSearchRule *Valuewrapperusersearchrule `json:"facilitatorsSearchRule,omitempty"`
		
		TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`
		
		ServiceGoalImpactOverrides *Valuewrapperactivityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`
		
		OptimizationObjective *string `json:"optimizationObjective,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		FixedAvailability *Listwrapperfixedavailability `json:"fixedAvailability,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		Description: o.Description,
		
		GroupSettings: o.GroupSettings,
		
		AttendeesSearchRule: o.AttendeesSearchRule,
		
		FacilitatorsSearchRule: o.FacilitatorsSearchRule,
		
		TransitionTimeMinutes: o.TransitionTimeMinutes,
		
		ServiceGoalImpactOverrides: o.ServiceGoalImpactOverrides,
		
		OptimizationObjective: o.OptimizationObjective,
		
		State: o.State,
		
		FixedAvailability: o.FixedAvailability,
		Alias:    (Alias)(o),
	})
}

func (o *Updateactivityplanrequest) UnmarshalJSON(b []byte) error {
	var UpdateactivityplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &UpdateactivityplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := UpdateactivityplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Description, ok := UpdateactivityplanrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if GroupSettings, ok := UpdateactivityplanrequestMap["groupSettings"].(map[string]interface{}); ok {
		GroupSettingsString, _ := json.Marshal(GroupSettings)
		json.Unmarshal(GroupSettingsString, &o.GroupSettings)
	}
	
	if AttendeesSearchRule, ok := UpdateactivityplanrequestMap["attendeesSearchRule"].(map[string]interface{}); ok {
		AttendeesSearchRuleString, _ := json.Marshal(AttendeesSearchRule)
		json.Unmarshal(AttendeesSearchRuleString, &o.AttendeesSearchRule)
	}
	
	if FacilitatorsSearchRule, ok := UpdateactivityplanrequestMap["facilitatorsSearchRule"].(map[string]interface{}); ok {
		FacilitatorsSearchRuleString, _ := json.Marshal(FacilitatorsSearchRule)
		json.Unmarshal(FacilitatorsSearchRuleString, &o.FacilitatorsSearchRule)
	}
	
	if TransitionTimeMinutes, ok := UpdateactivityplanrequestMap["transitionTimeMinutes"].(float64); ok {
		TransitionTimeMinutesInt := int(TransitionTimeMinutes)
		o.TransitionTimeMinutes = &TransitionTimeMinutesInt
	}
	
	if ServiceGoalImpactOverrides, ok := UpdateactivityplanrequestMap["serviceGoalImpactOverrides"].(map[string]interface{}); ok {
		ServiceGoalImpactOverridesString, _ := json.Marshal(ServiceGoalImpactOverrides)
		json.Unmarshal(ServiceGoalImpactOverridesString, &o.ServiceGoalImpactOverrides)
	}
	
	if OptimizationObjective, ok := UpdateactivityplanrequestMap["optimizationObjective"].(string); ok {
		o.OptimizationObjective = &OptimizationObjective
	}
    
	if State, ok := UpdateactivityplanrequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if FixedAvailability, ok := UpdateactivityplanrequestMap["fixedAvailability"].(map[string]interface{}); ok {
		FixedAvailabilityString, _ := json.Marshal(FixedAvailability)
		json.Unmarshal(FixedAvailabilityString, &o.FixedAvailability)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Updateactivityplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
