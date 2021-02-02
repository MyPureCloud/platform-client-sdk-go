package platformclientv2
import (
	"encoding/json"
)

// Adherencesettings - Schedule Adherence Configuration
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

// String returns a JSON representation of the model
func (o *Adherencesettings) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
