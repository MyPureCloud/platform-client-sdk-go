package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Actionmap
type Actionmap struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Version - The version of the action map.
	Version *int `json:"version,omitempty"`

	// IsActive - Whether the action map is active.
	IsActive *bool `json:"isActive,omitempty"`

	// DisplayName - Display name of the action map.
	DisplayName *string `json:"displayName,omitempty"`

	// TriggerWithSegments - Trigger action map if any segment in the list is assigned to a given customer.
	TriggerWithSegments *[]string `json:"triggerWithSegments,omitempty"`

	// TriggerWithEventConditions - List of event conditions that must be satisfied to trigger the action map.
	TriggerWithEventConditions *[]Eventcondition `json:"triggerWithEventConditions,omitempty"`

	// TriggerWithOutcomeProbabilityConditions - Probability conditions for outcomes that must be satisfied to trigger the action map.
	TriggerWithOutcomeProbabilityConditions *[]Outcomeprobabilitycondition `json:"triggerWithOutcomeProbabilityConditions,omitempty"`

	// PageUrlConditions - URL conditions that a page must match for web actions to be displayable.
	PageUrlConditions *[]Urlcondition `json:"pageUrlConditions,omitempty"`

	// Activation - Type of activation.
	Activation *Activation `json:"activation,omitempty"`

	// Weight - Weight of the action map with higher number denoting higher weight.
	Weight *int `json:"weight,omitempty"`

	// Action - The action that will be executed if this action map is triggered.
	Action *Actionmapaction `json:"action,omitempty"`

	// ActionMapScheduleGroups - The action map's associated schedule groups.
	ActionMapScheduleGroups *Actionmapschedulegroups `json:"actionMapScheduleGroups,omitempty"`

	// IgnoreFrequencyCap - Override organization-level frequency cap and always offer web engagements from this action map.
	IgnoreFrequencyCap *bool `json:"ignoreFrequencyCap,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

	// CreatedDate - Timestamp indicating when the action map was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreatedDate *time.Time `json:"createdDate,omitempty"`

	// ModifiedDate - Timestamp indicating when the action map was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`

	// StartDate - Timestamp at which the action map is scheduled to start firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate - Timestamp at which the action map is scheduled to stop firing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Actionmap) SetField(field string, fieldValue interface{}) {
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

func (o Actionmap) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreatedDate","ModifiedDate","StartDate","EndDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
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
	type Alias Actionmap
	
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
	
	StartDate := new(string)
	if o.StartDate != nil {
		
		*StartDate = timeutil.Strftime(o.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if o.EndDate != nil {
		
		*EndDate = timeutil.Strftime(o.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Version *int `json:"version,omitempty"`
		
		IsActive *bool `json:"isActive,omitempty"`
		
		DisplayName *string `json:"displayName,omitempty"`
		
		TriggerWithSegments *[]string `json:"triggerWithSegments,omitempty"`
		
		TriggerWithEventConditions *[]Eventcondition `json:"triggerWithEventConditions,omitempty"`
		
		TriggerWithOutcomeProbabilityConditions *[]Outcomeprobabilitycondition `json:"triggerWithOutcomeProbabilityConditions,omitempty"`
		
		PageUrlConditions *[]Urlcondition `json:"pageUrlConditions,omitempty"`
		
		Activation *Activation `json:"activation,omitempty"`
		
		Weight *int `json:"weight,omitempty"`
		
		Action *Actionmapaction `json:"action,omitempty"`
		
		ActionMapScheduleGroups *Actionmapschedulegroups `json:"actionMapScheduleGroups,omitempty"`
		
		IgnoreFrequencyCap *bool `json:"ignoreFrequencyCap,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		CreatedDate *string `json:"createdDate,omitempty"`
		
		ModifiedDate *string `json:"modifiedDate,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Version: o.Version,
		
		IsActive: o.IsActive,
		
		DisplayName: o.DisplayName,
		
		TriggerWithSegments: o.TriggerWithSegments,
		
		TriggerWithEventConditions: o.TriggerWithEventConditions,
		
		TriggerWithOutcomeProbabilityConditions: o.TriggerWithOutcomeProbabilityConditions,
		
		PageUrlConditions: o.PageUrlConditions,
		
		Activation: o.Activation,
		
		Weight: o.Weight,
		
		Action: o.Action,
		
		ActionMapScheduleGroups: o.ActionMapScheduleGroups,
		
		IgnoreFrequencyCap: o.IgnoreFrequencyCap,
		
		SelfUri: o.SelfUri,
		
		CreatedDate: CreatedDate,
		
		ModifiedDate: ModifiedDate,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		Alias:    (Alias)(o),
	})
}

func (o *Actionmap) UnmarshalJSON(b []byte) error {
	var ActionmapMap map[string]interface{}
	err := json.Unmarshal(b, &ActionmapMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ActionmapMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Version, ok := ActionmapMap["version"].(float64); ok {
		VersionInt := int(Version)
		o.Version = &VersionInt
	}
	
	if IsActive, ok := ActionmapMap["isActive"].(bool); ok {
		o.IsActive = &IsActive
	}
    
	if DisplayName, ok := ActionmapMap["displayName"].(string); ok {
		o.DisplayName = &DisplayName
	}
    
	if TriggerWithSegments, ok := ActionmapMap["triggerWithSegments"].([]interface{}); ok {
		TriggerWithSegmentsString, _ := json.Marshal(TriggerWithSegments)
		json.Unmarshal(TriggerWithSegmentsString, &o.TriggerWithSegments)
	}
	
	if TriggerWithEventConditions, ok := ActionmapMap["triggerWithEventConditions"].([]interface{}); ok {
		TriggerWithEventConditionsString, _ := json.Marshal(TriggerWithEventConditions)
		json.Unmarshal(TriggerWithEventConditionsString, &o.TriggerWithEventConditions)
	}
	
	if TriggerWithOutcomeProbabilityConditions, ok := ActionmapMap["triggerWithOutcomeProbabilityConditions"].([]interface{}); ok {
		TriggerWithOutcomeProbabilityConditionsString, _ := json.Marshal(TriggerWithOutcomeProbabilityConditions)
		json.Unmarshal(TriggerWithOutcomeProbabilityConditionsString, &o.TriggerWithOutcomeProbabilityConditions)
	}
	
	if PageUrlConditions, ok := ActionmapMap["pageUrlConditions"].([]interface{}); ok {
		PageUrlConditionsString, _ := json.Marshal(PageUrlConditions)
		json.Unmarshal(PageUrlConditionsString, &o.PageUrlConditions)
	}
	
	if Activation, ok := ActionmapMap["activation"].(map[string]interface{}); ok {
		ActivationString, _ := json.Marshal(Activation)
		json.Unmarshal(ActivationString, &o.Activation)
	}
	
	if Weight, ok := ActionmapMap["weight"].(float64); ok {
		WeightInt := int(Weight)
		o.Weight = &WeightInt
	}
	
	if Action, ok := ActionmapMap["action"].(map[string]interface{}); ok {
		ActionString, _ := json.Marshal(Action)
		json.Unmarshal(ActionString, &o.Action)
	}
	
	if ActionMapScheduleGroups, ok := ActionmapMap["actionMapScheduleGroups"].(map[string]interface{}); ok {
		ActionMapScheduleGroupsString, _ := json.Marshal(ActionMapScheduleGroups)
		json.Unmarshal(ActionMapScheduleGroupsString, &o.ActionMapScheduleGroups)
	}
	
	if IgnoreFrequencyCap, ok := ActionmapMap["ignoreFrequencyCap"].(bool); ok {
		o.IgnoreFrequencyCap = &IgnoreFrequencyCap
	}
    
	if SelfUri, ok := ActionmapMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if createdDateString, ok := ActionmapMap["createdDate"].(string); ok {
		CreatedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createdDateString)
		o.CreatedDate = &CreatedDate
	}
	
	if modifiedDateString, ok := ActionmapMap["modifiedDate"].(string); ok {
		ModifiedDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifiedDateString)
		o.ModifiedDate = &ModifiedDate
	}
	
	if startDateString, ok := ActionmapMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := ActionmapMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Actionmap) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
