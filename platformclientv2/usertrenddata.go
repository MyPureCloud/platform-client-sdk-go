package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Usertrenddata
type Usertrenddata struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateStartWorkday - Start workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`

	// DateEndWorkday - End workday used as the date range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`

	// PercentOfGoal - Percent of goal
	PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`

	// AverageValue - Average metric value
	AverageValue *float64 `json:"averageValue,omitempty"`

	// RankTotalPoints - Rank, ordered by total points
	RankTotalPoints *int `json:"rankTotalPoints,omitempty"`

	// RankPercentagePoints - Rank, ordered by percentage of points
	RankPercentagePoints *int `json:"rankPercentagePoints,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Usertrenddata) SetField(field string, fieldValue interface{}) {
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

func (o Usertrenddata) MarshalJSON() ([]byte, error) {
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
	type Alias Usertrenddata
	
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
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		PercentOfGoal *float64 `json:"percentOfGoal,omitempty"`
		
		AverageValue *float64 `json:"averageValue,omitempty"`
		
		RankTotalPoints *int `json:"rankTotalPoints,omitempty"`
		
		RankPercentagePoints *int `json:"rankPercentagePoints,omitempty"`
		Alias
	}{ 
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		PercentOfGoal: o.PercentOfGoal,
		
		AverageValue: o.AverageValue,
		
		RankTotalPoints: o.RankTotalPoints,
		
		RankPercentagePoints: o.RankPercentagePoints,
		Alias:    (Alias)(o),
	})
}

func (o *Usertrenddata) UnmarshalJSON(b []byte) error {
	var UsertrenddataMap map[string]interface{}
	err := json.Unmarshal(b, &UsertrenddataMap)
	if err != nil {
		return err
	}
	
	if dateStartWorkdayString, ok := UsertrenddataMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := UsertrenddataMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if PercentOfGoal, ok := UsertrenddataMap["percentOfGoal"].(float64); ok {
		o.PercentOfGoal = &PercentOfGoal
	}
    
	if AverageValue, ok := UsertrenddataMap["averageValue"].(float64); ok {
		o.AverageValue = &AverageValue
	}
    
	if RankTotalPoints, ok := UsertrenddataMap["rankTotalPoints"].(float64); ok {
		RankTotalPointsInt := int(RankTotalPoints)
		o.RankTotalPoints = &RankTotalPointsInt
	}
	
	if RankPercentagePoints, ok := UsertrenddataMap["rankPercentagePoints"].(float64); ok {
		RankPercentagePointsInt := int(RankPercentagePoints)
		o.RankPercentagePoints = &RankPercentagePointsInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Usertrenddata) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
