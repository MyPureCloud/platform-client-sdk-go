package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactbulkeditrequest
type Contactbulkeditrequest struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ContactListFilterId - Contact List Filter ID.
	ContactListFilterId *string `json:"contactListFilterId,omitempty"`

	// Criteria - Criteria to filter the contacts by.
	Criteria *Contactbulksearchcriteria `json:"criteria,omitempty"`

	// ContactIds - Contact IDs to be bulk edited.
	ContactIds *[]string `json:"contactIds,omitempty"`

	// Contact - Contact object with details of fields used for patching.
	Contact *Dialercontact `json:"contact,omitempty"`

	// GenerateDownloadURI - Whether to do backup export as part of Bulk Operation or not. Default: true.
	GenerateDownloadURI *bool `json:"generateDownloadURI,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactbulkeditrequest) SetField(field string, fieldValue interface{}) {
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

func (o Contactbulkeditrequest) MarshalJSON() ([]byte, error) {
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
	type Alias Contactbulkeditrequest
	
	return json.Marshal(&struct { 
		ContactListFilterId *string `json:"contactListFilterId,omitempty"`
		
		Criteria *Contactbulksearchcriteria `json:"criteria,omitempty"`
		
		ContactIds *[]string `json:"contactIds,omitempty"`
		
		Contact *Dialercontact `json:"contact,omitempty"`
		
		GenerateDownloadURI *bool `json:"generateDownloadURI,omitempty"`
		Alias
	}{ 
		ContactListFilterId: o.ContactListFilterId,
		
		Criteria: o.Criteria,
		
		ContactIds: o.ContactIds,
		
		Contact: o.Contact,
		
		GenerateDownloadURI: o.GenerateDownloadURI,
		Alias:    (Alias)(o),
	})
}

func (o *Contactbulkeditrequest) UnmarshalJSON(b []byte) error {
	var ContactbulkeditrequestMap map[string]interface{}
	err := json.Unmarshal(b, &ContactbulkeditrequestMap)
	if err != nil {
		return err
	}
	
	if ContactListFilterId, ok := ContactbulkeditrequestMap["contactListFilterId"].(string); ok {
		o.ContactListFilterId = &ContactListFilterId
	}
    
	if Criteria, ok := ContactbulkeditrequestMap["criteria"].(map[string]interface{}); ok {
		CriteriaString, _ := json.Marshal(Criteria)
		json.Unmarshal(CriteriaString, &o.Criteria)
	}
	
	if ContactIds, ok := ContactbulkeditrequestMap["contactIds"].([]interface{}); ok {
		ContactIdsString, _ := json.Marshal(ContactIds)
		json.Unmarshal(ContactIdsString, &o.ContactIds)
	}
	
	if Contact, ok := ContactbulkeditrequestMap["contact"].(map[string]interface{}); ok {
		ContactString, _ := json.Marshal(Contact)
		json.Unmarshal(ContactString, &o.Contact)
	}
	
	if GenerateDownloadURI, ok := ContactbulkeditrequestMap["generateDownloadURI"].(bool); ok {
		o.GenerateDownloadURI = &GenerateDownloadURI
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactbulkeditrequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
