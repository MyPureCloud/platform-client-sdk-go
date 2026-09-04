package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactsearchrequest
type Contactsearchrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// PageNumber - Page number (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - Page size (limited to fetching first 1,000 records; pageNumber * pageSize must be <= 1,000)
	PageSize *int `json:"pageSize,omitempty"`

	// DivisionIds - Which divisions to search, up to 50
	DivisionIds *[]string `json:"divisionIds,omitempty"`

	// Expand - Which fields, if any, to expand
	Expand *[]string `json:"expand,omitempty"`

	// Operation - Search operation to execute, currently supports {@code simpleSearch} only.
	Operation *Contactsearchoperation `json:"operation,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactsearchrequest) SetField(field string, fieldValue interface{}) {
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

func (o Contactsearchrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Contactsearchrequest
	
	return json.Marshal(&struct { 
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		DivisionIds *[]string `json:"divisionIds,omitempty"`
		
		Expand *[]string `json:"expand,omitempty"`
		
		Operation *Contactsearchoperation `json:"operation,omitempty"`
		Alias
	}{ 
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		DivisionIds: o.DivisionIds,
		
		Expand: o.Expand,
		
		Operation: o.Operation,
		Alias:    (Alias)(o),
	})
}

func (o *Contactsearchrequest) UnmarshalJSON(b []byte) error {
	var ContactsearchrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactsearchrequestMap)
	if err != nil {
		return err
	}
	
	if PageNumber, ok := ContactsearchrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := ContactsearchrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if DivisionIds, ok := ContactsearchrequestMap["divisionIds"].([]interface{}); ok {
		DivisionIdsString, _ := json.Marshal(DivisionIds)
		json.Unmarshal(DivisionIdsString, &o.DivisionIds)
	}
	
	if Expand, ok := ContactsearchrequestMap["expand"].([]interface{}); ok {
		ExpandString, _ := json.Marshal(Expand)
		json.Unmarshal(ExpandString, &o.Expand)
	}
	
	if Operation, ok := ContactsearchrequestMap["operation"].(map[string]interface{}); ok {
		OperationString, _ := json.Marshal(Operation)
		json.Unmarshal(OperationString, &o.Operation)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactsearchrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
