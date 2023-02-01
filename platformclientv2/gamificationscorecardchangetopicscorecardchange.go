package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Gamificationscorecardchangetopicscorecardchange
type Gamificationscorecardchangetopicscorecardchange struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Workday
	Workday *string `json:"workday,omitempty"`

	// DivisionId
	DivisionId *string `json:"divisionId,omitempty"`

	// TeamId
	TeamId *string `json:"teamId,omitempty"`

	// PerformanceProfileId
	PerformanceProfileId *string `json:"performanceProfileId,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// PerformanceMetrics
	PerformanceMetrics *[]Gamificationscorecardchangetopicperformancemetric `json:"performanceMetrics,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Gamificationscorecardchangetopicscorecardchange) SetField(field string, fieldValue interface{}) {
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

func (o Gamificationscorecardchangetopicscorecardchange) MarshalJSON() ([]byte, error) {
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
	type Alias Gamificationscorecardchangetopicscorecardchange
	
	return json.Marshal(&struct { 
		Workday *string `json:"workday,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		TeamId *string `json:"teamId,omitempty"`
		
		PerformanceProfileId *string `json:"performanceProfileId,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		PerformanceMetrics *[]Gamificationscorecardchangetopicperformancemetric `json:"performanceMetrics,omitempty"`
		Alias
	}{ 
		Workday: o.Workday,
		
		DivisionId: o.DivisionId,
		
		TeamId: o.TeamId,
		
		PerformanceProfileId: o.PerformanceProfileId,
		
		UserId: o.UserId,
		
		PerformanceMetrics: o.PerformanceMetrics,
		Alias:    (Alias)(o),
	})
}

func (o *Gamificationscorecardchangetopicscorecardchange) UnmarshalJSON(b []byte) error {
	var GamificationscorecardchangetopicscorecardchangeMap map[string]interface{}
	err := json.Unmarshal(b, &GamificationscorecardchangetopicscorecardchangeMap)
	if err != nil {
		return err
	}
	
	if Workday, ok := GamificationscorecardchangetopicscorecardchangeMap["workday"].(string); ok {
		o.Workday = &Workday
	}
    
	if DivisionId, ok := GamificationscorecardchangetopicscorecardchangeMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if TeamId, ok := GamificationscorecardchangetopicscorecardchangeMap["teamId"].(string); ok {
		o.TeamId = &TeamId
	}
    
	if PerformanceProfileId, ok := GamificationscorecardchangetopicscorecardchangeMap["performanceProfileId"].(string); ok {
		o.PerformanceProfileId = &PerformanceProfileId
	}
    
	if UserId, ok := GamificationscorecardchangetopicscorecardchangeMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if PerformanceMetrics, ok := GamificationscorecardchangetopicscorecardchangeMap["performanceMetrics"].([]interface{}); ok {
		PerformanceMetricsString, _ := json.Marshal(PerformanceMetrics)
		json.Unmarshal(PerformanceMetricsString, &o.PerformanceMetrics)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Gamificationscorecardchangetopicscorecardchange) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
