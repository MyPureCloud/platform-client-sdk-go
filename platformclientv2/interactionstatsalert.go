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

func (u *Interactionstatsalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Interactionstatsalert

	
	StartDate := new(string)
	if u.StartDate != nil {
		
		*StartDate = timeutil.Strftime(u.StartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDate = nil
	}
	
	EndDate := new(string)
	if u.EndDate != nil {
		
		*EndDate = timeutil.Strftime(u.EndDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		Id: u.Id,
		
		Name: u.Name,
		
		Dimension: u.Dimension,
		
		DimensionValue: u.DimensionValue,
		
		Metric: u.Metric,
		
		MediaType: u.MediaType,
		
		NumericRange: u.NumericRange,
		
		Statistic: u.Statistic,
		
		Value: u.Value,
		
		RuleId: u.RuleId,
		
		Unread: u.Unread,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		
		RuleUri: u.RuleUri,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Interactionstatsalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
