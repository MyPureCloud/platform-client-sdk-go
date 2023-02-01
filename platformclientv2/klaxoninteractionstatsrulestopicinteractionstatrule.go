package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Klaxoninteractionstatsrulestopicinteractionstatrule
type Klaxoninteractionstatsrulestopicinteractionstatrule struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Klaxoninteractionstatsrulestopicinteractionstatrule) SetField(field string, fieldValue interface{}) {
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

func (o Klaxoninteractionstatsrulestopicinteractionstatrule) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{  }

		// Construct object
		newObj := make(map[string]interface{})
		for fieldName := range o.SetFieldNames {
			// Get initial field value
			fieldValue := val.FieldByName(fieldName).Interface()

			// Apply value formatting overrides
			if fieldValue == nil || reflect.ValueOf(fieldValue).IsNil()  {
				// Do nothing. Just catching this case to avoid trying to custom serialize a nil value
			} else if contains(dateTimeFields, fieldName) {
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
		Alias
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
		Alias:    (Alias)(o),
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
