package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Userbestpointsitem
type Userbestpointsitem struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// GranularityType - Best points aggregation interval granularity
	GranularityType *string `json:"granularityType,omitempty"`

	// Points - Gamification points
	Points *int `json:"points,omitempty"`

	// DateStartWorkday - Start workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`

	// DateEndWorkday - End workday of the best points aggregation interval. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`

	// Rank - The rank of this user
	Rank *int `json:"rank,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Userbestpointsitem) SetField(field string, fieldValue interface{}) {
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

func (o Userbestpointsitem) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStartWorkday","DateEndWorkday", }

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
	type Alias Userbestpointsitem
	
	DateStartWorkday := new(string)
	if o.DateStartWorkday != nil {
		*DateStartWorkday = timeutil.Strftime(o.DateStartWorkday, "%Y-%m-%d")
	} else {
		DateStartWorkday = nil
	}
	
	DateEndWorkday := new(string)
	if o.DateEndWorkday != nil {
		*DateEndWorkday = timeutil.Strftime(o.DateEndWorkday, "%Y-%m-%d")
	} else {
		DateEndWorkday = nil
	}
	
	return json.Marshal(&struct { 
		GranularityType *string `json:"granularityType,omitempty"`
		
		Points *int `json:"points,omitempty"`
		
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		Rank *int `json:"rank,omitempty"`
		Alias
	}{ 
		GranularityType: o.GranularityType,
		
		Points: o.Points,
		
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		Rank: o.Rank,
		Alias:    (Alias)(o),
	})
}

func (o *Userbestpointsitem) UnmarshalJSON(b []byte) error {
	var UserbestpointsitemMap map[string]interface{}
	err := json.Unmarshal(b, &UserbestpointsitemMap)
	if err != nil {
		return err
	}
	
	if GranularityType, ok := UserbestpointsitemMap["granularityType"].(string); ok {
		o.GranularityType = &GranularityType
	}
    
	if Points, ok := UserbestpointsitemMap["points"].(float64); ok {
		PointsInt := int(Points)
		o.Points = &PointsInt
	}
	
	if dateStartWorkdayString, ok := UserbestpointsitemMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := UserbestpointsitemMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if Rank, ok := UserbestpointsitemMap["rank"].(float64); ok {
		RankInt := int(Rank)
		o.Rank = &RankInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Userbestpointsitem) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
