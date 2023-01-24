package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxoninteractionstatsalertstopicinteractionstatalert
type Klaxoninteractionstatsalertstopicinteractionstatalert struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Klaxoninteractionstatsalertstopicinteractionstatalert) SetField(field string, fieldValue interface{}) {
	// Get Value object for field
	target := reflect.ValueOf(o)
	targetField := reflect.Indirect(target).FieldByName(field)

	// Set value
	if fieldValue != nil {
		targetField.Set(reflect.ValueOf(fieldValue))
	} else {
		// Must create a new Value (creates **type) then get its element (*type), which will be nil pointer of the appropriate type
		x := reflect.Indirect(reflect.New(targetField.Type()))
		targetField.Set(x)
	}

	// Add field to set field names list
	if o.SetFieldNames == nil {
		o.SetFieldNames = make(map[string]bool)
	}
	o.SetFieldNames[field] = true
}

func (o Klaxoninteractionstatsalertstopicinteractionstatalert) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDate","EndDate", }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if contains(dateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%fZ")
			} else if contains(localDateTimeFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%dT%H:%M:%S.%f")
			} else if contains(dateFields, fieldName) {
				fieldValue = timeutil.Strftime(toTime(fieldValue), "%Y-%m-%d")
			}

			// Assign value to field using JSON tag name
			newObj[getFieldName(reflect.TypeOf(&o), fieldName)] = fieldValue
		}

		// Marshal and return dynamically constructed interface
		return json.Marshal(newObj)
	}

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
		Alias
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
		Alias:    (Alias)(o),
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
