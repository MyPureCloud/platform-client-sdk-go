package platformclientv2
import (
	"time"
	"encoding/json"
)

// Klaxoninteractionstatsalertstopicinteractionstatalert
type Klaxoninteractionstatsalertstopicinteractionstatalert struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// RuleId
	RuleId *string `json:"ruleId,omitempty"`


	// Dimension
	Dimension *string `json:"dimension,omitempty"`


	// DimensionValue
	DimensionValue *string `json:"dimensionValue,omitempty"`


	// DimensionValueName
	DimensionValueName *string `json:"dimensionValueName,omitempty"`


	// Metric
	Metric *string `json:"metric,omitempty"`


	// MediaType
	MediaType *string `json:"mediaType,omitempty"`


	// NumericRange
	NumericRange *string `json:"numericRange,omitempty"`


	// Statistic
	Statistic *string `json:"statistic,omitempty"`


	// Value
	Value *float32 `json:"value,omitempty"`


	// Unread
	Unread *bool `json:"unread,omitempty"`


	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`


	// NotificationUsers
	NotificationUsers *[]Klaxoninteractionstatsalertstopicnotificationuser `json:"notificationUsers,omitempty"`


	// AlertTypes
	AlertTypes *[]string `json:"alertTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
