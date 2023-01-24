package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Intradayplanninggrouprequest
type Intradayplanninggrouprequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// BusinessUnitDate - Requested date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	BusinessUnitDate *time.Time `json:"businessUnitDate,omitempty"`

	// Categories - The metric categories
	Categories *[]string `json:"categories,omitempty"`

	// PlanningGroupIds - The IDs of the planning groups for which to fetch data.  Omitting or passing an empty list will return all available planning groups
	PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`

	// IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data. Required, defaults to 15
	IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Intradayplanninggrouprequest) SetField(field string, fieldValue interface{}) {
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

func (o Intradayplanninggrouprequest) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{  }
		localDateTimeFields := []string{  }
		dateFields := []string{ "BusinessUnitDate", }

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
	type Alias Intradayplanninggrouprequest
	
	BusinessUnitDate := new(string)
	if o.BusinessUnitDate != nil {
		*BusinessUnitDate = timeutil.Strftime(o.BusinessUnitDate, "%Y-%m-%d")
	} else {
		BusinessUnitDate = nil
	}
	
	return json.Marshal(&struct { 
		BusinessUnitDate *string `json:"businessUnitDate,omitempty"`
		
		Categories *[]string `json:"categories,omitempty"`
		
		PlanningGroupIds *[]string `json:"planningGroupIds,omitempty"`
		
		IntervalLengthMinutes *int `json:"intervalLengthMinutes,omitempty"`
		Alias
	}{ 
		BusinessUnitDate: BusinessUnitDate,
		
		Categories: o.Categories,
		
		PlanningGroupIds: o.PlanningGroupIds,
		
		IntervalLengthMinutes: o.IntervalLengthMinutes,
		Alias:    (Alias)(o),
	})
}

func (o *Intradayplanninggrouprequest) UnmarshalJSON(b []byte) error {
	var IntradayplanninggrouprequestMap map[string]interface{}
	err := json.Unmarshal(b, &IntradayplanninggrouprequestMap)
	if err != nil {
		return err
	}
	
	if businessUnitDateString, ok := IntradayplanninggrouprequestMap["businessUnitDate"].(string); ok {
		BusinessUnitDate, _ := time.Parse("2006-01-02", businessUnitDateString)
		o.BusinessUnitDate = &BusinessUnitDate
	}
	
	if Categories, ok := IntradayplanninggrouprequestMap["categories"].([]interface{}); ok {
		CategoriesString, _ := json.Marshal(Categories)
		json.Unmarshal(CategoriesString, &o.Categories)
	}
	
	if PlanningGroupIds, ok := IntradayplanninggrouprequestMap["planningGroupIds"].([]interface{}); ok {
		PlanningGroupIdsString, _ := json.Marshal(PlanningGroupIds)
		json.Unmarshal(PlanningGroupIdsString, &o.PlanningGroupIds)
	}
	
	if IntervalLengthMinutes, ok := IntradayplanninggrouprequestMap["intervalLengthMinutes"].(float64); ok {
		IntervalLengthMinutesInt := int(IntervalLengthMinutes)
		o.IntervalLengthMinutes = &IntervalLengthMinutesInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Intradayplanninggrouprequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
