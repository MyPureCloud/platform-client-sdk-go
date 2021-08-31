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

func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		Dimension: o.Dimension,
		
		DimensionValue: o.DimensionValue,
		
		DimensionValueName: o.DimensionValueName,
		
		Metric: o.Metric,
		
		MediaType: o.MediaType,
		
		NumericRange: o.NumericRange,
		
		Statistic: o.Statistic,
		
		Value: o.Value,
		
		InAlarm: o.InAlarm,
		
		Enabled: o.Enabled,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) UnmarshalJSON(b []byte) error {
	var KlaxoninteractionstatsrulestopicinteractionstatruleMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxoninteractionstatsrulestopicinteractionstatruleMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Dimension, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
	
	if DimensionValue, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["dimensionValue"].(string); ok {
		o.DimensionValue = &DimensionValue
	}
	
	if DimensionValueName, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["dimensionValueName"].(string); ok {
		o.DimensionValueName = &DimensionValueName
	}
	
	if Metric, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["metric"].(string); ok {
		o.Metric = &Metric
	}
	
	if MediaType, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
	
	if NumericRange, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["numericRange"].(string); ok {
		o.NumericRange = &NumericRange
	}
	
	if Statistic, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["statistic"].(string); ok {
		o.Statistic = &Statistic
	}
	
	if Value, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
	
	if InAlarm, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["inAlarm"].(bool); ok {
		o.InAlarm = &InAlarm
	}
	
	if Enabled, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["enabled"].(bool); ok {
		o.Enabled = &Enabled
	}
	
	if NotificationUsers, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := KlaxoninteractionstatsrulestopicinteractionstatruleMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
