package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherencesettings
type Adherencesettings struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// SevereAlertThresholdMinutes - The threshold in minutes where an alert will be triggered when an agent is considered severely out of adherence
	SevereAlertThresholdMinutes *int `json:"severeAlertThresholdMinutes,omitempty"`

	// AdherenceTargetPercent - Target adherence percentage
	AdherenceTargetPercent *int `json:"adherenceTargetPercent,omitempty"`

	// AdherenceExceptionThresholdSeconds - The threshold in seconds for which agents should not be penalized for being momentarily out of adherence
	AdherenceExceptionThresholdSeconds *int `json:"adherenceExceptionThresholdSeconds,omitempty"`

	// NonOnQueueActivitiesEquivalent - Whether to treat all non-on-queue activities as equivalent for adherence purposes
	NonOnQueueActivitiesEquivalent *bool `json:"nonOnQueueActivitiesEquivalent,omitempty"`

	// TrackOnQueueActivity - Whether to track on-queue activities
	TrackOnQueueActivity *bool `json:"trackOnQueueActivity,omitempty"`

	// IgnoredActivityCategories - Activity categories that should be ignored for adherence purposes
	IgnoredActivityCategories *Ignoredactivitycategories `json:"ignoredActivityCategories,omitempty"`

	// IgnoredActivityCodeIds - Activity code IDs that should be ignored for adherence purposes
	IgnoredActivityCodeIds *Ignoredactivitycodeids `json:"ignoredActivityCodeIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adherencesettings) SetField(field string, fieldValue interface{}) {
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

func (o Adherencesettings) MarshalJSON() ([]byte, error) {
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
	type Alias Adherencesettings
	
	return json.Marshal(&struct { 
		SevereAlertThresholdMinutes *int `json:"severeAlertThresholdMinutes,omitempty"`
		
		AdherenceTargetPercent *int `json:"adherenceTargetPercent,omitempty"`
		
		AdherenceExceptionThresholdSeconds *int `json:"adherenceExceptionThresholdSeconds,omitempty"`
		
		NonOnQueueActivitiesEquivalent *bool `json:"nonOnQueueActivitiesEquivalent,omitempty"`
		
		TrackOnQueueActivity *bool `json:"trackOnQueueActivity,omitempty"`
		
		IgnoredActivityCategories *Ignoredactivitycategories `json:"ignoredActivityCategories,omitempty"`
		
		IgnoredActivityCodeIds *Ignoredactivitycodeids `json:"ignoredActivityCodeIds,omitempty"`
		Alias
	}{ 
		SevereAlertThresholdMinutes: o.SevereAlertThresholdMinutes,
		
		AdherenceTargetPercent: o.AdherenceTargetPercent,
		
		AdherenceExceptionThresholdSeconds: o.AdherenceExceptionThresholdSeconds,
		
		NonOnQueueActivitiesEquivalent: o.NonOnQueueActivitiesEquivalent,
		
		TrackOnQueueActivity: o.TrackOnQueueActivity,
		
		IgnoredActivityCategories: o.IgnoredActivityCategories,
		
		IgnoredActivityCodeIds: o.IgnoredActivityCodeIds,
		Alias:    (Alias)(o),
	})
}

func (o *Adherencesettings) UnmarshalJSON(b []byte) error {
	var AdherencesettingsMap map[string]interface{}
	err := json.Unmarshal(b, &AdherencesettingsMap)
	if err != nil {
		return err
	}
	
	if SevereAlertThresholdMinutes, ok := AdherencesettingsMap["severeAlertThresholdMinutes"].(float64); ok {
		SevereAlertThresholdMinutesInt := int(SevereAlertThresholdMinutes)
		o.SevereAlertThresholdMinutes = &SevereAlertThresholdMinutesInt
	}
	
	if AdherenceTargetPercent, ok := AdherencesettingsMap["adherenceTargetPercent"].(float64); ok {
		AdherenceTargetPercentInt := int(AdherenceTargetPercent)
		o.AdherenceTargetPercent = &AdherenceTargetPercentInt
	}
	
	if AdherenceExceptionThresholdSeconds, ok := AdherencesettingsMap["adherenceExceptionThresholdSeconds"].(float64); ok {
		AdherenceExceptionThresholdSecondsInt := int(AdherenceExceptionThresholdSeconds)
		o.AdherenceExceptionThresholdSeconds = &AdherenceExceptionThresholdSecondsInt
	}
	
	if NonOnQueueActivitiesEquivalent, ok := AdherencesettingsMap["nonOnQueueActivitiesEquivalent"].(bool); ok {
		o.NonOnQueueActivitiesEquivalent = &NonOnQueueActivitiesEquivalent
	}
    
	if TrackOnQueueActivity, ok := AdherencesettingsMap["trackOnQueueActivity"].(bool); ok {
		o.TrackOnQueueActivity = &TrackOnQueueActivity
	}
    
	if IgnoredActivityCategories, ok := AdherencesettingsMap["ignoredActivityCategories"].(map[string]interface{}); ok {
		IgnoredActivityCategoriesString, _ := json.Marshal(IgnoredActivityCategories)
		json.Unmarshal(IgnoredActivityCategoriesString, &o.IgnoredActivityCategories)
	}
	
	if IgnoredActivityCodeIds, ok := AdherencesettingsMap["ignoredActivityCodeIds"].(map[string]interface{}); ok {
		IgnoredActivityCodeIdsString, _ := json.Marshal(IgnoredActivityCodeIds)
		json.Unmarshal(IgnoredActivityCodeIdsString, &o.IgnoredActivityCodeIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherencesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
