package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Adherencesettings
type Adherencesettings struct { 
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

}

func (o *Adherencesettings) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		SevereAlertThresholdMinutes: o.SevereAlertThresholdMinutes,
		
		AdherenceTargetPercent: o.AdherenceTargetPercent,
		
		AdherenceExceptionThresholdSeconds: o.AdherenceExceptionThresholdSeconds,
		
		NonOnQueueActivitiesEquivalent: o.NonOnQueueActivitiesEquivalent,
		
		TrackOnQueueActivity: o.TrackOnQueueActivity,
		
		IgnoredActivityCategories: o.IgnoredActivityCategories,
		Alias:    (*Alias)(o),
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
	

	return nil
}

// String returns a JSON representation of the model
func (o *Adherencesettings) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
