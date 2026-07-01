package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Capacityplanforecastinputstemplate
type Capacityplanforecastinputstemplate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ReferenceBusinessUnitDate - The reference date for interval-based data relative to the business unit time zone for the forecast inputs. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	ReferenceBusinessUnitDate *time.Time `json:"referenceBusinessUnitDate,omitempty"`

	// Granularity - Granularity of the intervals
	Granularity *string `json:"granularity,omitempty"`

	// Months - The list of months covered by this capacity plan, formatted as yyyy-MM, populated for monthly granularity
	Months *[]string `json:"months,omitempty"`

	// PlanningGroupsForecastData - The forecast data for the planning groups
	PlanningGroupsForecastData *[]Forecastinputplanninggroupdata `json:"planningGroupsForecastData,omitempty"`

	// CapacityPlanForecastSummary - The summary of forecast inputs for this capacity plan, for the selected granularity
	CapacityPlanForecastSummary *Capacityplanforecastmetrics `json:"capacityPlanForecastSummary,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Capacityplanforecastinputstemplate) SetField(field string, fieldValue interface{}) {
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

func (o Capacityplanforecastinputstemplate) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "ReferenceBusinessUnitDate", }

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
	type Alias Capacityplanforecastinputstemplate
	
	ReferenceBusinessUnitDate := new(string)
	if o.ReferenceBusinessUnitDate != nil {
		*ReferenceBusinessUnitDate = timeutil.Strftime(o.ReferenceBusinessUnitDate, "%Y-%m-%d")
	} else {
		ReferenceBusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		ReferenceBusinessUnitDate *string `json:"referenceBusinessUnitDate,omitempty"`
		
		Granularity *string `json:"granularity,omitempty"`
		
		Months *[]string `json:"months,omitempty"`
		
		PlanningGroupsForecastData *[]Forecastinputplanninggroupdata `json:"planningGroupsForecastData,omitempty"`
		
		CapacityPlanForecastSummary *Capacityplanforecastmetrics `json:"capacityPlanForecastSummary,omitempty"`
		Alias
	}{ 
		ReferenceBusinessUnitDate: ReferenceBusinessUnitDate,
		
		Granularity: o.Granularity,
		
		Months: o.Months,
		
		PlanningGroupsForecastData: o.PlanningGroupsForecastData,
		
		CapacityPlanForecastSummary: o.CapacityPlanForecastSummary,
		Alias:    (Alias)(o),
	})
}

func (o *Capacityplanforecastinputstemplate) UnmarshalJSON(b []byte) error {
	var CapacityplanforecastinputstemplateMap map[string]interface{}
	err := json.Unmarshal(b, &CapacityplanforecastinputstemplateMap)
	if err != nil {
		return err
	}
	
	if referenceBusinessUnitDateString, ok := CapacityplanforecastinputstemplateMap["referenceBusinessUnitDate"].(string); ok {
		ReferenceBusinessUnitDate, _ := time.Parse("2006-01-02", referenceBusinessUnitDateString)
		o.ReferenceBusinessUnitDate = &ReferenceBusinessUnitDate
	}
	
	if Granularity, ok := CapacityplanforecastinputstemplateMap["granularity"].(string); ok {
		o.Granularity = &Granularity
	}
    
	if Months, ok := CapacityplanforecastinputstemplateMap["months"].([]interface{}); ok {
		MonthsString, _ := json.Marshal(Months)
		json.Unmarshal(MonthsString, &o.Months)
	}
	
	if PlanningGroupsForecastData, ok := CapacityplanforecastinputstemplateMap["planningGroupsForecastData"].([]interface{}); ok {
		PlanningGroupsForecastDataString, _ := json.Marshal(PlanningGroupsForecastData)
		json.Unmarshal(PlanningGroupsForecastDataString, &o.PlanningGroupsForecastData)
	}
	
	if CapacityPlanForecastSummary, ok := CapacityplanforecastinputstemplateMap["capacityPlanForecastSummary"].(map[string]interface{}); ok {
		CapacityPlanForecastSummaryString, _ := json.Marshal(CapacityPlanForecastSummary)
		json.Unmarshal(CapacityPlanForecastSummaryString, &o.CapacityPlanForecastSummary)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Capacityplanforecastinputstemplate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
