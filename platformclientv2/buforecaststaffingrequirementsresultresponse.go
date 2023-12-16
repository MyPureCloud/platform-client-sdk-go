package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Buforecaststaffingrequirementsresultresponse
type Buforecaststaffingrequirementsresultresponse struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BusinessUnitId - The ID of the business unit to which the forecast staffing requirements belongs
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// Forecast - The forecast reference
	Forecast *Bushorttermforecastreference `json:"forecast,omitempty"`

	// ReferenceStartDate - The reference start date for interval-based data for this forecast as an ISO-8601 string
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

	// WeekCount - The number of weeks in this forecast
	WeekCount *int `json:"weekCount,omitempty"`

	// IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

	// State - The state of the staffing requirements generation
	State *string `json:"state,omitempty"`

	// Results - The forecast staffing requirement results, Will be populated when state == 'Complete'
	Results *[]Buforecaststaffingrequirementsresult `json:"results,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Buforecaststaffingrequirementsresultresponse) SetField(field string, fieldValue interface{}) {
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

func (o Buforecaststaffingrequirementsresultresponse) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ReferenceStartDate", }
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
	type Alias Buforecaststaffingrequirementsresultresponse
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		Forecast *Bushorttermforecastreference `json:"forecast,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Results *[]Buforecaststaffingrequirementsresult `json:"results,omitempty"`
		Alias
	}{ 
		BusinessUnitId: o.BusinessUnitId,
		
		Forecast: o.Forecast,
		
		ReferenceStartDate: ReferenceStartDate,
		
		WeekCount: o.WeekCount,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		State: o.State,
		
		Results: o.Results,
		Alias:    (Alias)(o),
	})
}

func (o *Buforecaststaffingrequirementsresultresponse) UnmarshalJSON(b []byte) error {
	var BuforecaststaffingrequirementsresultresponseMap map[string]interface{}
	err := json.Unmarshal(b, &BuforecaststaffingrequirementsresultresponseMap)
	if err != nil {
		return err
	}
	
	if BusinessUnitId, ok := BuforecaststaffingrequirementsresultresponseMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if Forecast, ok := BuforecaststaffingrequirementsresultresponseMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if referenceStartDateString, ok := BuforecaststaffingrequirementsresultresponseMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if WeekCount, ok := BuforecaststaffingrequirementsresultresponseMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if IntervalLengthMinutes, ok := BuforecaststaffingrequirementsresultresponseMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if State, ok := BuforecaststaffingrequirementsresultresponseMap["state"].(string); ok {
		o.State = &State
	}
    
	if Results, ok := BuforecaststaffingrequirementsresultresponseMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Buforecaststaffingrequirementsresultresponse) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
