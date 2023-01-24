package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbuintradaydataupdatetopicbuintradayresult
type Wfmbuintradaydataupdatetopicbuintradayresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// StartDate
	StartDate *time.Time `json:"startDate,omitempty"`

	// EndDate
	EndDate *time.Time `json:"endDate,omitempty"`

	// IntervalLengthMinutes
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

	// IntradayDataGroupings
	IntradayDataGroupings *[]Wfmbuintradaydataupdatetopicbuintradaydatagroup `json:"intradayDataGroupings,omitempty"`

	// Categories
	Categories *[]string `json:"categories,omitempty"`

	// NoDataReason
	NoDataReason *string `json:"noDataReason,omitempty"`

	// Schedule
	Schedule *Wfmbuintradaydataupdatetopicbuschedulereference `json:"schedule,omitempty"`

	// ShortTermForecast
	ShortTermForecast *Wfmbuintradaydataupdatetopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbuintradaydataupdatetopicbuintradayresult) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbuintradaydataupdatetopicbuintradayresult) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbuintradaydataupdatetopicbuintradayresult
	
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
		StartDate *string `json:"startDate,omitempty"`
		
		EndDate *string `json:"endDate,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		
		IntradayDataGroupings *[]Wfmbuintradaydataupdatetopicbuintradaydatagroup `json:"intradayDataGroupings,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		NoDataReason *string `json:"noDataReason,omitempty"`
		
		Schedule *Wfmbuintradaydataupdatetopicbuschedulereference `json:"schedule,omitempty"`
		
		ShortTermForecast *Wfmbuintradaydataupdatetopicbushorttermforecastreference `json:"shortTermForecast,omitempty"`
		Alias
	}{ 
		StartDate: StartDate,
		
		EndDate: EndDate,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		IntradayDataGroupings: o.IntradayDataGroupings,
		
		Categories: o.Categories,
		
		NoDataReason: o.NoDataReason,
		
		Schedule: o.Schedule,
		
		ShortTermForecast: o.ShortTermForecast,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbuintradaydataupdatetopicbuintradayresult) UnmarshalJSON(b []byte) error {
	var WfmbuintradaydataupdatetopicbuintradayresultMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbuintradaydataupdatetopicbuintradayresultMap)
	if err != nil {
		return err
	}
	
	if startDateString, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["startDate"].(string); ok {
		StartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateString)
		o.StartDate = &StartDate
	}
	
	if endDateString, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["endDate"].(string); ok {
		EndDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateString)
		o.EndDate = &EndDate
	}
	
	if IntervalLengthMinutes, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if IntradayDataGroupings, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["intradayDataGroupings"].([]interface{}); ok {
		IntradayDataGroupingsString, _ := json.Marshal(IntradayDataGroupings)
		json.Unmarshal(IntradayDataGroupingsString, &o.IntradayDataGroupings)
	}
	
	if Categories, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if NoDataReason, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["noDataReason"].(string); ok {
		o.NoDataReason = &NoDataReason
	}
    
	if Schedule, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["schedule"].(map[string]interface{}); ok {
		ScheduleString, _ := json.Marshal(Schedule)
		json.Unmarshal(ScheduleString, &o.Schedule)
	}
	
	if ShortTermForecast, ok := WfmbuintradaydataupdatetopicbuintradayresultMap["shortTermForecast"].(map[string]interface{}); ok {
		ShortTermForecastString, _ := json.Marshal(ShortTermForecast)
		json.Unmarshal(ShortTermForecastString, &o.ShortTermForecast)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbuintradaydataupdatetopicbuintradayresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
