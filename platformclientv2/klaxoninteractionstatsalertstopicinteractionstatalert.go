package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
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

func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxoninteractionstatsalertstopicinteractionstatalert
	
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
		
		Name *string `json:"name,omitempty"`
		
		RuleId *string `json:"ruleId,omitempty"`
		
		Dimension *string `json:"dimension,omitempty"`
		
		DimensionValue *string `json:"dimensionValue,omitempty"`
		
		DimensionValueName *string `json:"dimensionValueName,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		NumericRange *string `json:"numericRange,omitempty"`
		
		Statistic *string `json:"statistic,omitempty"`
		
		Value *float32 `json:"value,omitempty"`
		
		Unread *bool `json:"unread,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		NotificationUsers *[]Klaxoninteractionstatsalertstopicnotificationuser `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		RuleId: o.RuleId,
		
		Dimension: o.Dimension,
		
		DimensionValue: o.DimensionValue,
		
		DimensionValueName: o.DimensionValueName,
		
		Metric: o.Metric,
		
		MediaType: o.MediaType,
		
		NumericRange: o.NumericRange,
		
		Statistic: o.Statistic,
		
		Value: o.Value,
		
		Unread: o.Unread,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		Alias:    (*Alias)(o),
	})
}

func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) UnmarshalJSON(b []byte) error {
	var KlaxoninteractionstatsalertstopicinteractionstatalertMap map[string]interface{}
	err := json.Unmarshal(b, &KlaxoninteractionstatsalertstopicinteractionstatalertMap)
	if err != nil {
		return err
	}
	
	if Id, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if RuleId, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["ruleId"].(string); ok {
		o.RuleId = &RuleId
	}
    
	if Dimension, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if DimensionValue, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["dimensionValue"].(string); ok {
		o.DimensionValue = &DimensionValue
	}
    
	if DimensionValueName, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["dimensionValueName"].(string); ok {
		o.DimensionValueName = &DimensionValueName
	}
    
	if Metric, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MediaType, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if NumericRange, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["numericRange"].(string); ok {
		o.NumericRange = &NumericRange
	}
    
	if Statistic, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["statistic"].(string); ok {
		o.Statistic = &Statistic
	}
    
	if Value, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["value"].(float64); ok {
		ValueFloat32 := float32(Value)
		o.Value = &ValueFloat32
	}
    
	if Unread, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["unread"].(bool); ok {
		o.Unread = &Unread
	}
    
	if startDateString, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if NotificationUsers, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := KlaxoninteractionstatsalertstopicinteractionstatalertMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
