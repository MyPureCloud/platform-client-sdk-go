package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Createactivityplanrequest
type Createactivityplanrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Name - The name of the activity plan
	Name *string `json:"name,omitempty"`

	// ManagementUnitIds - The management units to which this activity plan applies. Empty list or null means this activity plan applies to the entire business unit
	ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`

	// Description - The description of the activity plan
	Description *string `json:"description,omitempty"`

	// ActivityCodeId - The activity code associated with the activity plan
	ActivityCodeId *string `json:"activityCodeId,omitempty"`

	// VarType - The type of the activity plan
	VarType *string `json:"type,omitempty"`

	// LengthMinutes - The length in minutes of the activity plan
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// InitialSchedulePeriod - The initial scheduling period for the activity plan
	InitialSchedulePeriod *Schedulingperiod `json:"initialSchedulePeriod,omitempty"`

	// GroupSettings - Group settings for the activity plan
	GroupSettings *Groupsettings `json:"groupSettings,omitempty"`

	// RecurrenceSettings - Settings controlling recurrence for the activity plan. If not set the activity plan will only occur once
	RecurrenceSettings *Recurrencesettings `json:"recurrenceSettings,omitempty"`

	// AttendeesSearchRule - Attendee search rule for this activity plan
	AttendeesSearchRule *Usersearchrule `json:"attendeesSearchRule,omitempty"`

	// Facilitated - Whether the sessions created by this activity plan should be facilitated
	Facilitated *bool `json:"facilitated,omitempty"`

	// FacilitatorsSearchRule - Facilitator search rule for this activity plan
	FacilitatorsSearchRule *Usersearchrule `json:"facilitatorsSearchRule,omitempty"`

	// TransitionTimeMinutes - Transition time in minutes between facilitated sessions
	TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`

	// ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan. If not set the business unit setting will be used
	ServiceGoalImpactOverrides *Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`

	// OptimizationObjective - The optimization objective of this activity plan
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// State - The state of this activity plan
	State *string `json:"state,omitempty"`

	// CountsAsPaidTime - Whether the activity should count as paid time
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// FixedAvailability - Fixed availability configuration for the activity plan
	FixedAvailability *[]Fixedavailability `json:"fixedAvailability,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Createactivityplanrequest) SetField(field string, fieldValue interface{}) {
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

func (o Createactivityplanrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Createactivityplanrequest
	
	return json.Marshal(&struct { 
		Name *string `json:"name,omitempty"`
		
		ManagementUnitIds *[]string `json:"managementUnitIds,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCodeId *string `json:"activityCodeId,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		InitialSchedulePeriod *Schedulingperiod `json:"initialSchedulePeriod,omitempty"`
		
		GroupSettings *Groupsettings `json:"groupSettings,omitempty"`
		
		RecurrenceSettings *Recurrencesettings `json:"recurrenceSettings,omitempty"`
		
		AttendeesSearchRule *Usersearchrule `json:"attendeesSearchRule,omitempty"`
		
		Facilitated *bool `json:"facilitated,omitempty"`
		
		FacilitatorsSearchRule *Usersearchrule `json:"facilitatorsSearchRule,omitempty"`
		
		TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`
		
		ServiceGoalImpactOverrides *Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`
		
		OptimizationObjective *string `json:"optimizationObjective,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		FixedAvailability *[]Fixedavailability `json:"fixedAvailability,omitempty"`
		Alias
	}{ 
		Name: o.Name,
		
		ManagementUnitIds: o.ManagementUnitIds,
		
		Description: o.Description,
		
		ActivityCodeId: o.ActivityCodeId,
		
		VarType: o.VarType,
		
		LengthMinutes: o.LengthMinutes,
		
		InitialSchedulePeriod: o.InitialSchedulePeriod,
		
		GroupSettings: o.GroupSettings,
		
		RecurrenceSettings: o.RecurrenceSettings,
		
		AttendeesSearchRule: o.AttendeesSearchRule,
		
		Facilitated: o.Facilitated,
		
		FacilitatorsSearchRule: o.FacilitatorsSearchRule,
		
		TransitionTimeMinutes: o.TransitionTimeMinutes,
		
		ServiceGoalImpactOverrides: o.ServiceGoalImpactOverrides,
		
		OptimizationObjective: o.OptimizationObjective,
		
		State: o.State,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		FixedAvailability: o.FixedAvailability,
		Alias:    (Alias)(o),
	})
}

func (o *Createactivityplanrequest) UnmarshalJSON(b []byte) error {
	var CreateactivityplanrequestMap map[string]interface{}
	err := json.Unmarshal(b, &CreateactivityplanrequestMap)
	if err != nil {
		return err
	}
	
	if Name, ok := CreateactivityplanrequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnitIds, ok := CreateactivityplanrequestMap["managementUnitIds"].([]interface{}); ok {
		ManagementUnitIdsString, _ := json.Marshal(ManagementUnitIds)
		json.Unmarshal(ManagementUnitIdsString, &o.ManagementUnitIds)
	}
	
	if Description, ok := CreateactivityplanrequestMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCodeId, ok := CreateactivityplanrequestMap["activityCodeId"].(string); ok {
		o.ActivityCodeId = &ActivityCodeId
	}
    
	if VarType, ok := CreateactivityplanrequestMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if LengthMinutes, ok := CreateactivityplanrequestMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if InitialSchedulePeriod, ok := CreateactivityplanrequestMap["initialSchedulePeriod"].(map[string]interface{}); ok {
		InitialSchedulePeriodString, _ := json.Marshal(InitialSchedulePeriod)
		json.Unmarshal(InitialSchedulePeriodString, &o.InitialSchedulePeriod)
	}
	
	if GroupSettings, ok := CreateactivityplanrequestMap["groupSettings"].(map[string]interface{}); ok {
		GroupSettingsString, _ := json.Marshal(GroupSettings)
		json.Unmarshal(GroupSettingsString, &o.GroupSettings)
	}
	
	if RecurrenceSettings, ok := CreateactivityplanrequestMap["recurrenceSettings"].(map[string]interface{}); ok {
		RecurrenceSettingsString, _ := json.Marshal(RecurrenceSettings)
		json.Unmarshal(RecurrenceSettingsString, &o.RecurrenceSettings)
	}
	
	if AttendeesSearchRule, ok := CreateactivityplanrequestMap["attendeesSearchRule"].(map[string]interface{}); ok {
		AttendeesSearchRuleString, _ := json.Marshal(AttendeesSearchRule)
		json.Unmarshal(AttendeesSearchRuleString, &o.AttendeesSearchRule)
	}
	
	if Facilitated, ok := CreateactivityplanrequestMap["facilitated"].(bool); ok {
		o.Facilitated = &Facilitated
	}
    
	if FacilitatorsSearchRule, ok := CreateactivityplanrequestMap["facilitatorsSearchRule"].(map[string]interface{}); ok {
		FacilitatorsSearchRuleString, _ := json.Marshal(FacilitatorsSearchRule)
		json.Unmarshal(FacilitatorsSearchRuleString, &o.FacilitatorsSearchRule)
	}
	
	if TransitionTimeMinutes, ok := CreateactivityplanrequestMap["transitionTimeMinutes"].(float64); ok {
		TransitionTimeMinutesInt := int(TransitionTimeMinutes)
		o.TransitionTimeMinutes = &TransitionTimeMinutesInt
	}
	
	if ServiceGoalImpactOverrides, ok := CreateactivityplanrequestMap["serviceGoalImpactOverrides"].(map[string]interface{}); ok {
		ServiceGoalImpactOverridesString, _ := json.Marshal(ServiceGoalImpactOverrides)
		json.Unmarshal(ServiceGoalImpactOverridesString, &o.ServiceGoalImpactOverrides)
	}
	
	if OptimizationObjective, ok := CreateactivityplanrequestMap["optimizationObjective"].(string); ok {
		o.OptimizationObjective = &OptimizationObjective
	}
    
	if State, ok := CreateactivityplanrequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if CountsAsPaidTime, ok := CreateactivityplanrequestMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if FixedAvailability, ok := CreateactivityplanrequestMap["fixedAvailability"].([]interface{}); ok {
		FixedAvailabilityString, _ := json.Marshal(FixedAvailability)
		json.Unmarshal(FixedAvailabilityString, &o.FixedAvailability)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Createactivityplanrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
