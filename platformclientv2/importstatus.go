package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Importstatus
type Importstatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// State - current status of the import
	State *string `json:"state,omitempty"`

	// TotalRecords - total number of records to be imported
	TotalRecords *int `json:"totalRecords,omitempty"`

	// CompletedRecords - number of records finished importing
	CompletedRecords *int `json:"completedRecords,omitempty"`

	// PercentComplete - percentage of records finished importing
	PercentComplete *int `json:"percentComplete,omitempty"`

	// FailureReason - if the import has failed, the reason for the failure
	FailureReason *string `json:"failureReason,omitempty"`

	// TargetContactListIds - The contact list Ids for target contact lists.
	TargetContactListIds *[]string `json:"targetContactListIds,omitempty"`

	// ListNamePrefix - The prefix for the contact list name
	ListNamePrefix *string `json:"listNamePrefix,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Importstatus) SetField(field string, fieldValue interface{}) {
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

func (o Importstatus) MarshalJSON() ([]byte, error) {
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
	type Alias Importstatus
	
	return json.Marshal(&struct { 
		State *string `json:"state,omitempty"`
		
		TotalRecords *int `json:"totalRecords,omitempty"`
		
		CompletedRecords *int `json:"completedRecords,omitempty"`
		
		PercentComplete *int `json:"percentComplete,omitempty"`
		
		FailureReason *string `json:"failureReason,omitempty"`
		
		TargetContactListIds *[]string `json:"targetContactListIds,omitempty"`
		
		ListNamePrefix *string `json:"listNamePrefix,omitempty"`
		Alias
	}{ 
		State: o.State,
		
		TotalRecords: o.TotalRecords,
		
		CompletedRecords: o.CompletedRecords,
		
		PercentComplete: o.PercentComplete,
		
		FailureReason: o.FailureReason,
		
		TargetContactListIds: o.TargetContactListIds,
		
		ListNamePrefix: o.ListNamePrefix,
		Alias:    (Alias)(o),
	})
}

func (o *Importstatus) UnmarshalJSON(b []byte) error {
	var ImportstatusMap map[string]interface{}
	err := json.Unmarshal(b, &ImportstatusMap)
	if err != nil {
		return err
	}
	
	if State, ok := ImportstatusMap["state"].(string); ok {
		o.State = &State
	}
    
	if TotalRecords, ok := ImportstatusMap["totalRecords"].(float64); ok {
		TotalRecordsInt := int(TotalRecords)
		o.TotalRecords = &TotalRecordsInt
	}
	
	if CompletedRecords, ok := ImportstatusMap["completedRecords"].(float64); ok {
		CompletedRecordsInt := int(CompletedRecords)
		o.CompletedRecords = &CompletedRecordsInt
	}
	
	if PercentComplete, ok := ImportstatusMap["percentComplete"].(float64); ok {
		PercentCompleteInt := int(PercentComplete)
		o.PercentComplete = &PercentCompleteInt
	}
	
	if FailureReason, ok := ImportstatusMap["failureReason"].(string); ok {
		o.FailureReason = &FailureReason
	}
    
	if TargetContactListIds, ok := ImportstatusMap["targetContactListIds"].([]interface{}); ok {
		TargetContactListIdsString, _ := json.Marshal(TargetContactListIds)
		json.Unmarshal(TargetContactListIdsString, &o.TargetContactListIds)
	}
	
	if ListNamePrefix, ok := ImportstatusMap["listNamePrefix"].(string); ok {
		o.ListNamePrefix = &ListNamePrefix
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Importstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
