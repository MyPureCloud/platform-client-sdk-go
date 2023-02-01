package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workdayvaluestrend
type Workdayvaluestrend struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// DateStartWorkday - The start workday for the query range for the metric value trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateStartWorkday *time.Time `json:"dateStartWorkday,omitempty"`

	// DateEndWorkday - The end workday for the query range for the metric value trend. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateEndWorkday *time.Time `json:"dateEndWorkday,omitempty"`

	// DateReferenceWorkday - The reference workday used to determine the metric definition. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	DateReferenceWorkday *time.Time `json:"dateReferenceWorkday,omitempty"`

	// Division - The targeted division for the query
	Division *Division `json:"division,omitempty"`

	// User - The targeted user for the query
	User *Userreference `json:"user,omitempty"`

	// Timezone - The time zone used for aggregating metric values
	Timezone *string `json:"timezone,omitempty"`

	// Results - The metric value trends
	Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`

	// PerformanceProfile - The targeted performance profile for the average points
	PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`

	// Metric - The targeted metric for the average points
	Metric *Addressableentityref `json:"metric,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workdayvaluestrend) SetField(field string, fieldValue interface{}) {
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

func (o Workdayvaluestrend) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "DateStartWorkday","DateEndWorkday","DateReferenceWorkday", }

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
	type Alias Workdayvaluestrend
	
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
	
	DateReferenceWorkday := new(string)
	if o.DateReferenceWorkday != nil {
		*DateReferenceWorkday = timeutil.Strftime(o.DateReferenceWorkday, "%Y-%m-%d")
	} else {
		DateReferenceWorkday = nil
	}
	
	return json.Marshal(&struct { 
		DateStartWorkday *string `json:"dateStartWorkday,omitempty"`
		
		DateEndWorkday *string `json:"dateEndWorkday,omitempty"`
		
		DateReferenceWorkday *string `json:"dateReferenceWorkday,omitempty"`
		
		Division *Division `json:"division,omitempty"`
		
		User *Userreference `json:"user,omitempty"`
		
		Timezone *string `json:"timezone,omitempty"`
		
		Results *[]Workdayvaluesmetricitem `json:"results,omitempty"`
		
		PerformanceProfile *Addressableentityref `json:"performanceProfile,omitempty"`
		
		Metric *Addressableentityref `json:"metric,omitempty"`
		Alias
	}{ 
		DateStartWorkday: DateStartWorkday,
		
		DateEndWorkday: DateEndWorkday,
		
		DateReferenceWorkday: DateReferenceWorkday,
		
		Division: o.Division,
		
		User: o.User,
		
		Timezone: o.Timezone,
		
		Results: o.Results,
		
		PerformanceProfile: o.PerformanceProfile,
		
		Metric: o.Metric,
		Alias:    (Alias)(o),
	})
}

func (o *Workdayvaluestrend) UnmarshalJSON(b []byte) error {
	var WorkdayvaluestrendMap map[string]interface{}
	err := json.Unmarshal(b, &WorkdayvaluestrendMap)
	if err != nil {
		return err
	}
	
	if dateStartWorkdayString, ok := WorkdayvaluestrendMap["dateStartWorkday"].(string); ok {
		DateStartWorkday, _ := time.Parse("2006-01-02", dateStartWorkdayString)
		o.DateStartWorkday = &DateStartWorkday
	}
	
	if dateEndWorkdayString, ok := WorkdayvaluestrendMap["dateEndWorkday"].(string); ok {
		DateEndWorkday, _ := time.Parse("2006-01-02", dateEndWorkdayString)
		o.DateEndWorkday = &DateEndWorkday
	}
	
	if dateReferenceWorkdayString, ok := WorkdayvaluestrendMap["dateReferenceWorkday"].(string); ok {
		DateReferenceWorkday, _ := time.Parse("2006-01-02", dateReferenceWorkdayString)
		o.DateReferenceWorkday = &DateReferenceWorkday
	}
	
	if Division, ok := WorkdayvaluestrendMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if User, ok := WorkdayvaluestrendMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Timezone, ok := WorkdayvaluestrendMap["timezone"].(string); ok {
		o.Timezone = &Timezone
	}
    
	if Results, ok := WorkdayvaluestrendMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	
	if PerformanceProfile, ok := WorkdayvaluestrendMap["performanceProfile"].(map[string]interface{}); ok {
		PerformanceProfileString, _ := json.Marshal(PerformanceProfile)
		json.Unmarshal(PerformanceProfileString, &o.PerformanceProfile)
	}
	
	if Metric, ok := WorkdayvaluestrendMap["metric"].(map[string]interface{}); ok {
		MetricString, _ := json.Marshal(Metric)
		json.Unmarshal(MetricString, &o.Metric)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workdayvaluestrend) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
