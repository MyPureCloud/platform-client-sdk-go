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

func (u *Klaxoninteractionstatsalertstopicinteractionstatalert) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Klaxoninteractionstatsalertstopicinteractionstatalert

	
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
		Id: u.Id,
		
		Name: u.Name,
		
		RuleId: u.RuleId,
		
		Dimension: u.Dimension,
		
		DimensionValue: u.DimensionValue,
		
		DimensionValueName: u.DimensionValueName,
		
		Metric: u.Metric,
		
		MediaType: u.MediaType,
		
		NumericRange: u.NumericRange,
		
		Statistic: u.Statistic,
		
		Value: u.Value,
		
		Unread: u.Unread,
		
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		NotificationUsers: u.NotificationUsers,
		
		AlertTypes: u.AlertTypes,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
