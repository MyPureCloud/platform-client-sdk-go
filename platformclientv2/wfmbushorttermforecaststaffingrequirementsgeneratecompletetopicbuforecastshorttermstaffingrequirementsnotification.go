package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification
type Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BusinessUnitId
	BusinessUnitId *string `json:"businessUnitId,omitempty"`

	// State
	State *string `json:"state,omitempty"`

	// Forecast
	Forecast *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbushorttermforecastreference `json:"forecast,omitempty"`

	// WeekCount
	WeekCount *int `json:"weekCount,omitempty"`

	// IntervalLengthMinutes
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`

	// ReferenceStartDate
	ReferenceStartDate *time.Time `json:"referenceStartDate,omitempty"`

	// Results
	Results *[]Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults `json:"results,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification) SetField(field string, fieldValue interface{}) {
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

func (o Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification) MarshalJSON() ([]byte, error) {
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
	type Alias Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification
	
	ReferenceStartDate := new(string)
	if o.ReferenceStartDate != nil {
		
		*ReferenceStartDate = timeutil.Strftime(o.ReferenceStartDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ReferenceStartDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitId *string `json:"businessUnitId,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		Forecast *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbushorttermforecastreference `json:"forecast,omitempty"`
		
		WeekCount *int `json:"weekCount,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		
		ReferenceStartDate *string `json:"referenceStartDate,omitempty"`
		
		Results *[]Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsresults `json:"results,omitempty"`
		Alias
	}{ 
		BusinessUnitId: o.BusinessUnitId,
		
		State: o.State,
		
		Forecast: o.Forecast,
		
		WeekCount: o.WeekCount,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		
		ReferenceStartDate: ReferenceStartDate,
		
		Results: o.Results,
		Alias:    (Alias)(o),
	})
}

func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification) UnmarshalJSON(b []byte) error {
	var WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap map[string]interface{}
	err := json.Unmarshal(b, &WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap)
	if err != nil {
		return err
	}
	
	if BusinessUnitId, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["businessUnitId"].(string); ok {
		o.BusinessUnitId = &BusinessUnitId
	}
    
	if State, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["state"].(string); ok {
		o.State = &State
	}
    
	if Forecast, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["forecast"].(map[string]interface{}); ok {
		ForecastString, _ := json.Marshal(Forecast)
		json.Unmarshal(ForecastString, &o.Forecast)
	}
	
	if WeekCount, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["weekCount"].(float64); ok {
		WeekCountInt := int(WeekCount)
		o.WeekCount = &WeekCountInt
	}
	
	if IntervalLengthMinutes, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	
	if referenceStartDateString, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["referenceStartDate"].(string); ok {
		ReferenceStartDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", referenceStartDateString)
		o.ReferenceStartDate = &ReferenceStartDate
	}
	
	if Results, ok := WfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotificationMap["results"].([]interface{}); ok {
		ResultsString, _ := json.Marshal(Results)
		json.Unmarshal(ResultsString, &o.Results)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Wfmbushorttermforecaststaffingrequirementsgeneratecompletetopicbuforecastshorttermstaffingrequirementsnotification) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
