package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Activityplanresponse
type Activityplanresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name - The name of the activity plan
	Name *string `json:"name,omitempty"`

	// ManagementUnits - The management units to which this activity plan applies. Empty list or null means this activity plan applies to the entire business unit
	ManagementUnits *[]Managementunitreference `json:"managementUnits,omitempty"`

	// Description - The description of this activity plan
	Description *string `json:"description,omitempty"`

	// ActivityCode - The activity code associated with this activity plan. It is recommended to load and cache the entire list of activity codes rather than look up individual codes
	ActivityCode *Activitycodereference `json:"activityCode,omitempty"`

	// VarType - The type of the activity plan
	VarType *string `json:"type,omitempty"`

	// InitialSchedulePeriod - The initial schedule period of the activity plan
	InitialSchedulePeriod *Schedulingperiod `json:"initialSchedulePeriod,omitempty"`

	// LengthMinutes - The length of the activity in minutes
	LengthMinutes *int `json:"lengthMinutes,omitempty"`

	// GroupSettings - Group settings for this activity plan
	GroupSettings *Groupsettings `json:"groupSettings,omitempty"`

	// RecurrenceSettings - Recurrence settings for this activity plan
	RecurrenceSettings *Recurrencesettings `json:"recurrenceSettings,omitempty"`

	// AttendeesSearchRule - Attendee search rule for this activity plan
	AttendeesSearchRule *Usersearchrule `json:"attendeesSearchRule,omitempty"`

	// Facilitated - Whether the sessions created by this activity plan should be facilitated
	Facilitated *bool `json:"facilitated,omitempty"`

	// FacilitatorsSearchRule - Facilitator search rule for this activity plan
	FacilitatorsSearchRule *Usersearchrule `json:"facilitatorsSearchRule,omitempty"`

	// TransitionTimeMinutes - Transition time in minutes between facilitated sessions
	TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`

	// ServiceGoalImpactOverrides - Allowable service goal impact override settings for this activity plan
	ServiceGoalImpactOverrides *Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`

	// OptimizationObjective - The optimization objective of this activity plan
	OptimizationObjective *string `json:"optimizationObjective,omitempty"`

	// FixedAvailability - Fixed availability configuration for this activity plan
	FixedAvailability *[]Fixedavailability `json:"fixedAvailability,omitempty"`

	// State - The state of this activity plan
	State *string `json:"state,omitempty"`

	// CountsAsPaidTime - Whether the activity should count as paid time
	CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`

	// CreatedDate - The date the activity plan was created, in ISO-8601 format
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// CreatedBy - The user who created this activity plan
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// ModifiedDate - The date the activity plan was modified, in ISO-8601 format
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// ModifiedBy - The last user to modify this activity plan
	ModifiedBy *Userreference `json:"modifiedBy,omitempty"`

	// LastRunDate - The date the activity plan was last run, in ISO-8601 format
	LastRunDate *time.Time `json:"lastRunDate,omitempty"`

	// LastRunBy - The last user to run this activity plan
	LastRunBy *Userreference `json:"lastRunBy,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Activityplanresponse) SetField(field string, fieldValue interface{}) {
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

func (o Activityplanresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate","LastRunDate", }
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
	type Alias Activityplanresponse
	
	CreatedDate := new(string)
	if o.CreatedDate != nil {
		
		*CreatedDate = timeutil.Strftime(o.CreatedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreatedDate = nil
	}
	
	ModifiedDate := new(string)
	if o.ModifiedDate != nil {
		
		*ModifiedDate = timeutil.Strftime(o.ModifiedDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifiedDate = nil
	}
	
	LastRunDate := new(string)
	if o.LastRunDate != nil {
		
		*LastRunDate = timeutil.Strftime(o.LastRunDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		LastRunDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ManagementUnits *[]Managementunitreference `json:"managementUnits,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		ActivityCode *Activitycodereference `json:"activityCode,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		InitialSchedulePeriod *Schedulingperiod `json:"initialSchedulePeriod,omitempty"`
		
		LengthMinutes *int `json:"lengthMinutes,omitempty"`
		
		GroupSettings *Groupsettings `json:"groupSettings,omitempty"`
		
		RecurrenceSettings *Recurrencesettings `json:"recurrenceSettings,omitempty"`
		
		AttendeesSearchRule *Usersearchrule `json:"attendeesSearchRule,omitempty"`
		
		Facilitated *bool `json:"facilitated,omitempty"`
		
		FacilitatorsSearchRule *Usersearchrule `json:"facilitatorsSearchRule,omitempty"`
		
		TransitionTimeMinutes *int `json:"transitionTimeMinutes,omitempty"`
		
		ServiceGoalImpactOverrides *Activityplanservicegoalimpactoverrides `json:"serviceGoalImpactOverrides,omitempty"`
		
		OptimizationObjective *string `json:"optimizationObjective,omitempty"`
		
		FixedAvailability *[]Fixedavailability `json:"fixedAvailability,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		CountsAsPaidTime *bool `json:"countsAsPaidTime,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		ModifiedBy *Userreference `json:"modifiedBy,omitempty"`
		
		LastRunDate *string `json:"lastRunDate,omitempty"`
		
		LastRunBy *Userreference `json:"lastRunBy,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ManagementUnits: o.ManagementUnits,
		
		Description: o.Description,
		
		ActivityCode: o.ActivityCode,
		
		VarType: o.VarType,
		
		InitialSchedulePeriod: o.InitialSchedulePeriod,
		
		LengthMinutes: o.LengthMinutes,
		
		GroupSettings: o.GroupSettings,
		
		RecurrenceSettings: o.RecurrenceSettings,
		
		AttendeesSearchRule: o.AttendeesSearchRule,
		
		Facilitated: o.Facilitated,
		
		FacilitatorsSearchRule: o.FacilitatorsSearchRule,
		
		TransitionTimeMinutes: o.TransitionTimeMinutes,
		
		ServiceGoalImpactOverrides: o.ServiceGoalImpactOverrides,
		
		OptimizationObjective: o.OptimizationObjective,
		
		FixedAvailability: o.FixedAvailability,
		
		State: o.State,
		
		CountsAsPaidTime: o.CountsAsPaidTime,
		
		CreatedDate: CreatedDate,
		
		CreatedBy: o.CreatedBy,
		
		ModifiedDate: ModifiedDate,
		
		ModifiedBy: o.ModifiedBy,
		
		LastRunDate: LastRunDate,
		
		LastRunBy: o.LastRunBy,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Activityplanresponse) UnmarshalJSON(b []byte) error {
	var ActivityplanresponseMap map[string]interface{}
	err := json.Unmarshal(b, &ActivityplanresponseMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActivityplanresponseMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := ActivityplanresponseMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if ManagementUnits, ok := ActivityplanresponseMap["managementUnits"].([]interface{}); ok {
		ManagementUnitsString, _ := json.Marshal(ManagementUnits)
		json.Unmarshal(ManagementUnitsString, &o.ManagementUnits)
	}
	
	if Description, ok := ActivityplanresponseMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if ActivityCode, ok := ActivityplanresponseMap["activityCode"].(map[string]interface{}); ok {
		ActivityCodeString, _ := json.Marshal(ActivityCode)
		json.Unmarshal(ActivityCodeString, &o.ActivityCode)
	}
	
	if VarType, ok := ActivityplanresponseMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if InitialSchedulePeriod, ok := ActivityplanresponseMap["initialSchedulePeriod"].(map[string]interface{}); ok {
		InitialSchedulePeriodString, _ := json.Marshal(InitialSchedulePeriod)
		json.Unmarshal(InitialSchedulePeriodString, &o.InitialSchedulePeriod)
	}
	
	if LengthMinutes, ok := ActivityplanresponseMap["lengthMinutes"].(float64); ok {
		LengthMinutesInt := int(LengthMinutes)
		o.LengthMinutes = &LengthMinutesInt
	}
	
	if GroupSettings, ok := ActivityplanresponseMap["groupSettings"].(map[string]interface{}); ok {
		GroupSettingsString, _ := json.Marshal(GroupSettings)
		json.Unmarshal(GroupSettingsString, &o.GroupSettings)
	}
	
	if RecurrenceSettings, ok := ActivityplanresponseMap["recurrenceSettings"].(map[string]interface{}); ok {
		RecurrenceSettingsString, _ := json.Marshal(RecurrenceSettings)
		json.Unmarshal(RecurrenceSettingsString, &o.RecurrenceSettings)
	}
	
	if AttendeesSearchRule, ok := ActivityplanresponseMap["attendeesSearchRule"].(map[string]interface{}); ok {
		AttendeesSearchRuleString, _ := json.Marshal(AttendeesSearchRule)
		json.Unmarshal(AttendeesSearchRuleString, &o.AttendeesSearchRule)
	}
	
	if Facilitated, ok := ActivityplanresponseMap["facilitated"].(bool); ok {
		o.Facilitated = &Facilitated
	}
    
	if FacilitatorsSearchRule, ok := ActivityplanresponseMap["facilitatorsSearchRule"].(map[string]interface{}); ok {
		FacilitatorsSearchRuleString, _ := json.Marshal(FacilitatorsSearchRule)
		json.Unmarshal(FacilitatorsSearchRuleString, &o.FacilitatorsSearchRule)
	}
	
	if TransitionTimeMinutes, ok := ActivityplanresponseMap["transitionTimeMinutes"].(float64); ok {
		TransitionTimeMinutesInt := int(TransitionTimeMinutes)
		o.TransitionTimeMinutes = &TransitionTimeMinutesInt
	}
	
	if ServiceGoalImpactOverrides, ok := ActivityplanresponseMap["serviceGoalImpactOverrides"].(map[string]interface{}); ok {
		ServiceGoalImpactOverridesString, _ := json.Marshal(ServiceGoalImpactOverrides)
		json.Unmarshal(ServiceGoalImpactOverridesString, &o.ServiceGoalImpactOverrides)
	}
	
	if OptimizationObjective, ok := ActivityplanresponseMap["optimizationObjective"].(string); ok {
		o.OptimizationObjective = &OptimizationObjective
	}
    
	if FixedAvailability, ok := ActivityplanresponseMap["fixedAvailability"].([]interface{}); ok {
		FixedAvailabilityString, _ := json.Marshal(FixedAvailability)
		json.Unmarshal(FixedAvailabilityString, &o.FixedAvailability)
	}
	
	if State, ok := ActivityplanresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if CountsAsPaidTime, ok := ActivityplanresponseMap["countsAsPaidTime"].(bool); ok {
		o.CountsAsPaidTime = &CountsAsPaidTime
	}
    
	if createdDateString, ok := ActivityplanresponseMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if CreatedBy, ok := ActivityplanresponseMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if modifiedDateString, ok := ActivityplanresponseMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if ModifiedBy, ok := ActivityplanresponseMap["modifiedBy"].(map[string]interface{}); ok {
		ModifiedByString, _ := json.Marshal(ModifiedBy)
		json.Unmarshal(ModifiedByString, &o.ModifiedBy)
	}
	
	if lastRunDateString, ok := ActivityplanresponseMap["lastRunDate"].(string); ok {
		LastRunDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", lastRunDateString)
		o.LastRunDate = &LastRunDate
	}
	
	if LastRunBy, ok := ActivityplanresponseMap["lastRunBy"].(map[string]interface{}); ok {
		LastRunByString, _ := json.Marshal(LastRunBy)
		json.Unmarshal(LastRunByString, &o.LastRunBy)
	}
	
	if SelfUri, ok := ActivityplanresponseMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Activityplanresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
