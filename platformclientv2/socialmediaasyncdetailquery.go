package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Socialmediaasyncdetailquery
type Socialmediaasyncdetailquery struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval *string `json:"interval,omitempty"`

	// TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	TimeZone *string `json:"timeZone,omitempty"`

	// Filter - Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
	Filter *Socialmediaqueryfilter `json:"filter,omitempty"`

	// PageSize - The number of results per page
	PageSize *int `json:"pageSize,omitempty"`

	// Order - Sorting of results based on time
	Order *string `json:"order,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Socialmediaasyncdetailquery) SetField(field string, fieldValue interface{}) {
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

func (o Socialmediaasyncdetailquery) MarshalJSON() ([]byte, error) {
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
	type Alias Socialmediaasyncdetailquery
	
	return json.Marshal(&struct { 
		Interval *string `json:"interval,omitempty"`
		
		TimeZone *string `json:"timeZone,omitempty"`
		
		Filter *Socialmediaqueryfilter `json:"filter,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		Order *string `json:"order,omitempty"`
		Alias
	}{ 
		Interval: o.Interval,
		
		TimeZone: o.TimeZone,
		
		Filter: o.Filter,
		
		PageSize: o.PageSize,
		
		Order: o.Order,
		Alias:    (Alias)(o),
	})
}

func (o *Socialmediaasyncdetailquery) UnmarshalJSON(b []byte) error {
	var SocialmediaasyncdetailqueryMap map[string]interface{}
	err := json.Unmarshal(b, &SocialmediaasyncdetailqueryMap)
	if err != nil {
		return err
	}
	
	if Interval, ok := SocialmediaasyncdetailqueryMap["interval"].(string); ok {
		o.Interval = &Interval
	}
    
	if TimeZone, ok := SocialmediaasyncdetailqueryMap["timeZone"].(string); ok {
		o.TimeZone = &TimeZone
	}
    
	if Filter, ok := SocialmediaasyncdetailqueryMap["filter"].(map[string]interface{}); ok {
		FilterString, _ := json.Marshal(Filter)
		json.Unmarshal(FilterString, &o.Filter)
	}
	
	if PageSize, ok := SocialmediaasyncdetailqueryMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if Order, ok := SocialmediaasyncdetailqueryMap["order"].(string); ok {
		o.Order = &Order
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Socialmediaasyncdetailquery) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
