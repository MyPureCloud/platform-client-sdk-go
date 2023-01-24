package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Comparisonperiod
type Comparisonperiod struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Kpi - Key Performance Indicator optimised during the comparison period.
	Kpi *string `json:"kpi,omitempty"`

	// DateStarted - Start date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// DateEnded - End date of the comparison period. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateEnded *time.Time `json:"dateEnded,omitempty"`

	// KpiResults - KPI results for each metric
	KpiResults *[]Kpiresult `json:"kpiResults,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Comparisonperiod) SetField(field string, fieldValue interface{}) {
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

func (o Comparisonperiod) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateStarted","DateEnded", }
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
	type Alias Comparisonperiod
	
	DateStarted := new(string)
	if o.DateStarted != nil {
		
		*DateStarted = timeutil.Strftime(o.DateStarted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateStarted = nil
	}
	
	DateEnded := new(string)
	if o.DateEnded != nil {
		
		*DateEnded = timeutil.Strftime(o.DateEnded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateEnded = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Kpi *string `json:"kpi,omitempty"`
		
		DateStarted *string `json:"dateStarted,omitempty"`
		
		DateEnded *string `json:"dateEnded,omitempty"`
		
		KpiResults *[]Kpiresult `json:"kpiResults,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Kpi: o.Kpi,
		
		DateStarted: DateStarted,
		
		DateEnded: DateEnded,
		
		KpiResults: o.KpiResults,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Comparisonperiod) UnmarshalJSON(b []byte) error {
	var ComparisonperiodMap map[string]interface{}
	err := json.Unmarshal(b, &ComparisonperiodMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ComparisonperiodMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Kpi, ok := ComparisonperiodMap["kpi"].(string); ok {
		o.Kpi = &Kpi
	}
    
	if dateStartedString, ok := ComparisonperiodMap["dateStarted"].(string); ok {
		DateStarted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateStartedString)
		o.DateStarted = &DateStarted
	}
	
	if dateEndedString, ok := ComparisonperiodMap["dateEnded"].(string); ok {
		DateEnded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateEndedString)
		o.DateEnded = &DateEnded
	}
	
	if KpiResults, ok := ComparisonperiodMap["kpiResults"].([]interface{}); ok {
		KpiResultsString, _ := json.Marshal(KpiResults)
		json.Unmarshal(KpiResultsString, &o.KpiResults)
	}
	
	if SelfUri, ok := ComparisonperiodMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Comparisonperiod) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
