package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Interactionstatsalert
type Interactionstatsalert struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name - Name of the rule that generated the alert
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


	// RuleId - The id of the rule.
	RuleId *string `json:"ruleId,omitempty"`


	// Unread - Indicates if the alert has been read.
	Unread *bool `json:"unread,omitempty"`


	// StartDate - The date/time the alert was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDate *time.Time `json:"startDate,omitempty"`


	// EndDate - The date/time the owning rule exiting in alarm status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDate *time.Time `json:"endDate,omitempty"`


	// NotificationUsers - The ids of users who were notified of alarm state change.
	NotificationUsers *[]User `json:"notificationUsers,omitempty"`


	// AlertTypes - A collection of notification methods.
	AlertTypes *[]string `json:"alertTypes,omitempty"`


	// RuleUri
	RuleUri *string `json:"ruleUri,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Interactionstatsalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Interactionstatsalert
	
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
		
		Dimension *string `json:"dimension,omitempty"`
		
		DimensionValue *string `json:"dimensionValue,omitempty"`
		
		Metric *string `json:"metric,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		NumericRange *string `json:"numericRange,omitempty"`
		
		Statistic *string `json:"statistic,omitempty"`
		
		Value *float64 `json:"value,omitempty"`
		
		RuleId *string `json:"ruleId,omitempty"`
		
		Unread *bool `json:"unread,omitempty"`
		
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		NotificationUsers *[]User `json:"notificationUsers,omitempty"`
		
		AlertTypes *[]string `json:"alertTypes,omitempty"`
		
		RuleUri *string `json:"ruleUri,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Dimension: o.Dimension,
		
		DimensionValue: o.DimensionValue,
		
		Metric: o.Metric,
		
		MediaType: o.MediaType,
		
		NumericRange: o.NumericRange,
		
		Statistic: o.Statistic,
		
		Value: o.Value,
		
		RuleId: o.RuleId,
		
		Unread: o.Unread,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: o.NotificationUsers,
		
		AlertTypes: o.AlertTypes,
		
		RuleUri: o.RuleUri,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Interactionstatsalert) UnmarshalJSON(b []byte) error {
	var InteractionstatsalertMap map[string]interface{}
	err := json.Unmarshal(b, &InteractionstatsalertMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InteractionstatsalertMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := InteractionstatsalertMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Dimension, ok := InteractionstatsalertMap["dimension"].(string); ok {
		o.Dimension = &Dimension
	}
    
	if DimensionValue, ok := InteractionstatsalertMap["dimensionValue"].(string); ok {
		o.DimensionValue = &DimensionValue
	}
    
	if Metric, ok := InteractionstatsalertMap["metric"].(string); ok {
		o.Metric = &Metric
	}
    
	if MediaType, ok := InteractionstatsalertMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if NumericRange, ok := InteractionstatsalertMap["numericRange"].(string); ok {
		o.NumericRange = &NumericRange
	}
    
	if Statistic, ok := InteractionstatsalertMap["statistic"].(string); ok {
		o.Statistic = &Statistic
	}
    
	if Value, ok := InteractionstatsalertMap["value"].(float64); ok {
		o.Value = &Value
	}
    
	if RuleId, ok := InteractionstatsalertMap["ruleId"].(string); ok {
		o.RuleId = &RuleId
	}
    
	if Unread, ok := InteractionstatsalertMap["unread"].(bool); ok {
		o.Unread = &Unread
	}
    
	if startDateString, ok := InteractionstatsalertMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := InteractionstatsalertMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if NotificationUsers, ok := InteractionstatsalertMap["notificationUsers"].([]interface{}); ok {
		NotificationUsersString, _ := json.Marshal(NotificationUsers)
		json.Unmarshal(NotificationUsersString, &o.NotificationUsers)
	}
	
	if AlertTypes, ok := InteractionstatsalertMap["alertTypes"].([]interface{}); ok {
		AlertTypesString, _ := json.Marshal(AlertTypes)
		json.Unmarshal(AlertTypesString, &o.AlertTypes)
	}
	
	if RuleUri, ok := InteractionstatsalertMap["ruleUri"].(string); ok {
		o.RuleUri = &RuleUri
	}
    
	if SelfUri, ok := InteractionstatsalertMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Interactionstatsalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
