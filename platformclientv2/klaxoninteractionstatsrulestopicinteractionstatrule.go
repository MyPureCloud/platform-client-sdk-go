package platformclientv2
import (
	"github.com/leekchan/timeutil"
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

func (u *Klaxoninteractionstatsrulestopicinteractionstatrule) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxoninteractionstatsrulestopicinteractionstatrule

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		DimensionValue *string `json:"dimensionValue,omitempty"`
		
		DimensionValueName *string `json:"dimensionValueName,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		NumericRange *string `json:"numericRange,omitempty"`
		
		Statistic *string `json:"statistic,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		InAlarm *bool `json:"inAlarm,omitempty"`
		
		Enabled *bool `json:"enabled,omitempty"`
		
		NotificationUsers *[]Klaxoninteractionstatsrulestopicnotificationuser `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		Dimension: u.Dimension,
		
		DimensionValue: u.DimensionValue,
		
		DimensionValueName: u.DimensionValueName,
		
		Metric: u.Metric,
		
		MediaType: u.MediaType,
		
		NumericRange: u.NumericRange,
		
		Statistic: u.Statistic,
		
		Value: u.Value,
		
		InAlarm: u.InAlarm,
		
		Enabled: u.Enabled,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
