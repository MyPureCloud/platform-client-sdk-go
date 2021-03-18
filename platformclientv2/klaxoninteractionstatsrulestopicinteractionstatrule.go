package platformclientv2
import (
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxoninteractionstatsrulestopicinteractionstatrule
type Klaxoninteractionstatsrulestopicinteractionstatrule struct { 
	// Id
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


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


	// InAlarm
	InAlarm *bool `json:"inAlarm,omitempty"`


	// Enabled
	Enabled *bool `json:"enabled,omitempty"`


	// NotificationUsers
	NotificationUsers *[]Klaxoninteractionstatsrulestopicnotificationuser `json:"notificationUsers,omitempty"`


	// AlertTypes
	AlertTypes *[]string `json:"alertTypes,omitempty"`

}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
