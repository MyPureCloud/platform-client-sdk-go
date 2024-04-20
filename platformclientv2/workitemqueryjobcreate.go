package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Workitemqueryjobcreate
type Workitemqueryjobcreate struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageSize - The total page size requested. Default 25
	PageSize *int `json:"pageSize,omitempty"`

	// PageNumber - The page number requested
	PageNumber *int `json:"pageNumber,omitempty"`

	// Filters - List of filter objects to be used in the search.
	Filters *[]Workitemqueryjobfilter `json:"filters,omitempty"`

	// Expands - List of entity attributes to be expanded in the result.
	Expands *[]string `json:"expands,omitempty"`

	// Attributes - List of entity attributes to be retrieved in the result.
	Attributes *[]string `json:"attributes,omitempty"`

	// Sort - Sort
	Sort *Workitemqueryjobsort `json:"sort,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Workitemqueryjobcreate) SetField(field string, fieldValue interface{}) {
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

func (o Workitemqueryjobcreate) MarshalJSON() ([]byte, error) {
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
	type Alias Workitemqueryjobcreate
	
	return json.Marshal(&struct { 
		PageSize *int `json:"pageSize,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		Filters *[]Workitemqueryjobfilter `json:"filters,omitempty"`
		
		Expands *[]string `json:"expands,omitempty"`
		
		Attributes *[]string `json:"attributes,omitempty"`
		
		Sort *Workitemqueryjobsort `json:"sort,omitempty"`
		Alias
	}{ 
		PageSize: o.PageSize,
		
		PageNumber: o.PageNumber,
		
		Filters: o.Filters,
		
		Expands: o.Expands,
		
		Attributes: o.Attributes,
		
		Sort: o.Sort,
		Alias:    (Alias)(o),
	})
}

func (o *Workitemqueryjobcreate) UnmarshalJSON(b []byte) error {
	var WorkitemqueryjobcreateMap map[string]interface{}
	err := json.Unmarshal(b, &WorkitemqueryjobcreateMap)
	if err != nil {
		return err
	}
	
	if PageSize, ok := WorkitemqueryjobcreateMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if PageNumber, ok := WorkitemqueryjobcreateMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if Filters, ok := WorkitemqueryjobcreateMap["filters"].([]interface{}); ok {
		FiltersString, _ := json.Marshal(Filters)
		json.Unmarshal(FiltersString, &o.Filters)
	}
	
	if Expands, ok := WorkitemqueryjobcreateMap["expands"].([]interface{}); ok {
		ExpandsString, _ := json.Marshal(Expands)
		json.Unmarshal(ExpandsString, &o.Expands)
	}
	
	if Attributes, ok := WorkitemqueryjobcreateMap["attributes"].([]interface{}); ok {
		AttributesString, _ := json.Marshal(Attributes)
		json.Unmarshal(AttributesString, &o.Attributes)
	}
	
	if Sort, ok := WorkitemqueryjobcreateMap["sort"].(map[string]interface{}); ok {
		SortString, _ := json.Marshal(Sort)
		json.Unmarshal(SortString, &o.Sort)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Workitemqueryjobcreate) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
