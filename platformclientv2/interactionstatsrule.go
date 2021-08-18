package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Interactionstatsrule
type Interactionstatsrule struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name of the rule
	Name *string `json:"name,omitempty"`


	// Dimension - The dimension of concern.
	Dimension *string `json:"dimension,omitempty"`


	// DimensionValue - The value of the dimension.
	DimensionValue *string `json:"dimensionValue,omitempty"`


	// Metric - The metric to be assessed.
	Metric *string `json:"metric,omitempty"`


	// MediaType - The media type.
	MediaType *string `json:"mediaType,omitempty"`


	// NumericRange - The comparison descriptor used against the metric's value.
	NumericRange *string `json:"numericRange,omitempty"`


	// Statistic - The statistic of concern for the metric.
	Statistic *string `json:"statistic,omitempty"`


	// Value - The threshold value.
	Value *float64 `json:"value,omitempty"`


	// Enabled - Indicates if the rule is enabled.
	Enabled *bool `json:"enabled,omitempty"`


	// InAlarm - Indicates if the rule is in alarm state.
	InAlarm *bool `json:"inAlarm,omitempty"`


	// NotificationUsers - The ids of users who will be notified of alarm state change.
	NotificationUsers *[]User `json:"notificationUsers,omitempty"`


	// AlertTypes - A collection of notification methods.
	AlertTypes *[]string `json:"alertTypes,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Interactionstatsrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Interactionstatsrule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		DimensionValue *string `json:"dimensionValue,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		NumericRange *string `json:"numericRange,omitempty"`
		
		Statistic *string `json:"statistic,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		NotificationUsers *[]User `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Dimension: u.Dimension,
		
		DimensionValue: u.DimensionValue,
		
		Metric: u.Metric,
		
		MediaType: u.MediaType,
		
		NumericRange: u.NumericRange,
		
		Statistic: u.Statistic,
		
		Value: u.Value,
		
		Enabled: u.Enabled,
		
		InAlarm: u.InAlarm,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Interactionstatsrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
