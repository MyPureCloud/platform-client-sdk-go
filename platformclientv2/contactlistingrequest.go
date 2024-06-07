package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactlistingrequest
type Contactlistingrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactListFilterId - Contact List Filter ID.
	ContactListFilterId *string `json:"contactListFilterId,omitempty"`

	// Criteria - Criteria to filter the contacts by.
	Criteria *Contactbulksearchcriteria `json:"criteria,omitempty"`

	// PageNumber - Page number.
	PageNumber *int `json:"pageNumber,omitempty"`

	// PageSize - Page size. The max that will be returned is 50.
	PageSize *int `json:"pageSize,omitempty"`

	// ContactSorts - The order in which to sort contacts.
	ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactlistingrequest) SetField(field string, fieldValue interface{}) {
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

func (o Contactlistingrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Contactlistingrequest
	
	return json.Marshal(&struct { 
		ContactListFilterId *string `json:"contactListFilterId,omitempty"`
		
		Criteria *Contactbulksearchcriteria `json:"criteria,omitempty"`
		
		PageNumber *int `json:"pageNumber,omitempty"`
		
		PageSize *int `json:"pageSize,omitempty"`
		
		ContactSorts *[]Contactsort `json:"contactSorts,omitempty"`
		Alias
	}{ 
		ContactListFilterId: o.ContactListFilterId,
		
		Criteria: o.Criteria,
		
		PageNumber: o.PageNumber,
		
		PageSize: o.PageSize,
		
		ContactSorts: o.ContactSorts,
		Alias:    (Alias)(o),
	})
}

func (o *Contactlistingrequest) UnmarshalJSON(b []byte) error {
	var ContactlistingrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactlistingrequestMap)
	if err != nil {
		return err
	}
	
	if ContactListFilterId, ok := ContactlistingrequestMap["contactListFilterId"].(string); ok {
		o.ContactListFilterId = &ContactListFilterId
	}
    
	if Criteria, ok := ContactlistingrequestMap["criteria"].(map[string]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	
	if PageNumber, ok := ContactlistingrequestMap["pageNumber"].(float64); ok {
		PageNumberInt := int(PageNumber)
		o.PageNumber = &PageNumberInt
	}
	
	if PageSize, ok := ContactlistingrequestMap["pageSize"].(float64); ok {
		PageSizeInt := int(PageSize)
		o.PageSize = &PageSizeInt
	}
	
	if ContactSorts, ok := ContactlistingrequestMap["contactSorts"].([]interface{}); ok {
		ContactSortsString, _ := json.Marshal(ContactSorts)
		json.Unmarshal(ContactSortsString, &o.ContactSorts)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Contactlistingrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
