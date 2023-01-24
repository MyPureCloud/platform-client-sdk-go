package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Dialerauditrequest
type Dialerauditrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// QueryPhrase - The word or words to search for.
	QueryPhrase *string `json:"queryPhrase,omitempty"`

	// QueryFields - The fields in which to search for the queryPhrase.
	QueryFields *[]string `json:"queryFields,omitempty"`

	// Facets - The fields to facet on.
	Facets *[]Auditfacet `json:"facets,omitempty"`

	// Filters - The fields to filter on.
	Filters *[]Auditfilter `json:"filters,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Dialerauditrequest) SetField(field string, fieldValue interface{}) {
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

func (o Dialerauditrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Dialerauditrequest
	
	return json.Marshal(&struct { 
		QueryPhrase *string `json:"queryPhrase,omitempty"`
		
		QueryFields *[]string `json:"queryFields,omitempty"`
		
		Facets *[]Auditfacet `json:"facets,omitempty"`
		
		Filters *[]Auditfilter `json:"filters,omitempty"`
		Alias
	}{ 
		QueryPhrase: o.QueryPhrase,
		
		QueryFields: o.QueryFields,
		
		Facets: o.Facets,
		
		Filters: o.Filters,
		Alias:    (Alias)(o),
	})
}

func (o *Dialerauditrequest) UnmarshalJSON(b []byte) error {
	var DialerauditrequestMap map[string]interface{}
	err := json.Unmarshal(b, &DialerauditrequestMap)
	if err != nil {
		return err
	}
	
	if QueryPhrase, ok := DialerauditrequestMap["queryPhrase"].(string); ok {
		o.QueryPhrase = &QueryPhrase
	}
    
	if QueryFields, ok := DialerauditrequestMap["queryFields"].([]interface{}); ok {
		QueryFieldsString, _ := json.Marshal(QueryFields)
		json.Unmarshal(QueryFieldsString, &o.QueryFields)
	}
	
	if Facets, ok := DialerauditrequestMap["facets"].([]interface{}); ok {
		FacetsString, _ := json.Marshal(Facets)
		json.Unmarshal(FacetsString, &o.Facets)
	}
	
	if Filters, ok := DialerauditrequestMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Dialerauditrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
