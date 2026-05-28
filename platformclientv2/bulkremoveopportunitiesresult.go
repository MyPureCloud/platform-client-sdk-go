package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkremoveopportunitiesresult
type Bulkremoveopportunitiesresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Status - The status indicating the result of the bulk operation for this item
	Status *string `json:"status,omitempty"`

	// VarError - The error result if the operation failed
	VarError *Bulkopportunitieserror `json:"error,omitempty"`

	// OpportunityId - The ID of the opportunity
	OpportunityId *string `json:"opportunityId,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Bulkremoveopportunitiesresult) SetField(field string, fieldValue interface{}) {
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

func (o Bulkremoveopportunitiesresult) MarshalJSON() ([]byte, error) {
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
	type Alias Bulkremoveopportunitiesresult
	
	return json.Marshal(&struct { 
		Status *string `json:"status,omitempty"`
		
		VarError *Bulkopportunitieserror `json:"error,omitempty"`
		
		OpportunityId *string `json:"opportunityId,omitempty"`
		Alias
	}{ 
		Status: o.Status,
		
		VarError: o.VarError,
		
		OpportunityId: o.OpportunityId,
		Alias:    (Alias)(o),
	})
}

func (o *Bulkremoveopportunitiesresult) UnmarshalJSON(b []byte) error {
	var BulkremoveopportunitiesresultMap map[string]interface{}
	err := json.Unmarshal(b, &BulkremoveopportunitiesresultMap)
	if err != nil {
		return err
	}
	
	if Status, ok := BulkremoveopportunitiesresultMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if VarError, ok := BulkremoveopportunitiesresultMap["error"].(map[string]interface{}); ok {
		VarErrorString, _ := json.Marshal(VarError)
		json.Unmarshal(VarErrorString, &o.VarError)
	}
	
	if OpportunityId, ok := BulkremoveopportunitiesresultMap["opportunityId"].(string); ok {
		o.OpportunityId = &OpportunityId
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkremoveopportunitiesresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
