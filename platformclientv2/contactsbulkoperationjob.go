package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Contactsbulkoperationjob
type Contactsbulkoperationjob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique job identifier.
	Id *string `json:"id,omitempty"`

	// State - The job state.
	State *string `json:"state,omitempty"`

	// VarType - The job type.
	VarType *string `json:"type,omitempty"`

	// TotalRecords - Total records that will be impacted by the bulk operation.
	TotalRecords *int `json:"totalRecords,omitempty"`

	// CompletedRecords - Amount of records that have been impacted by the bulk operation.
	CompletedRecords *int `json:"completedRecords,omitempty"`

	// PercentComplete - Percentage of records that have been impacted by the bulk operation.
	PercentComplete *int `json:"percentComplete,omitempty"`

	// FailureReason - Information on failure reason.
	FailureReason *Errorinfo `json:"failureReason,omitempty"`

	// DownloadURI - URI to download the original backup contacts.
	DownloadURI *string `json:"downloadURI,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Contactsbulkoperationjob) SetField(field string, fieldValue interface{}) {
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

func (o Contactsbulkoperationjob) MarshalJSON() ([]byte, error) {
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
	type Alias Contactsbulkoperationjob
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		TotalRecords *int `json:"totalRecords,omitempty"`
		
		CompletedRecords *int `json:"completedRecords,omitempty"`
		
		PercentComplete *int `json:"percentComplete,omitempty"`
		
		FailureReason *Errorinfo `json:"failureReason,omitempty"`
		
		DownloadURI *string `json:"downloadURI,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		State: o.State,
		
		VarType: o.VarType,
		
		TotalRecords: o.TotalRecords,
		
		CompletedRecords: o.CompletedRecords,
		
		PercentComplete: o.PercentComplete,
		
		FailureReason: o.FailureReason,
		
		DownloadURI: o.DownloadURI,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Contactsbulkoperationjob) UnmarshalJSON(b []byte) error {
	var ContactsbulkoperationjobMap map[string]interface{}
	err := json.Unmarshal(b, &ContactsbulkoperationjobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := ContactsbulkoperationjobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if State, ok := ContactsbulkoperationjobMap["state"].(string); ok {
		o.State = &State
	}
    
	if VarType, ok := ContactsbulkoperationjobMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if TotalRecords, ok := ContactsbulkoperationjobMap["totalRecords"].(float64); ok {
		TotalRecordsInt := int(TotalRecords)
		o.TotalRecords = &TotalRecordsInt
	}
	
	if CompletedRecords, ok := ContactsbulkoperationjobMap["completedRecords"].(float64); ok {
		CompletedRecordsInt := int(CompletedRecords)
		o.CompletedRecords = &CompletedRecordsInt
	}
	
	if PercentComplete, ok := ContactsbulkoperationjobMap["percentComplete"].(float64); ok {
		PercentCompleteInt := int(PercentComplete)
		o.PercentComplete = &PercentCompleteInt
	}
	
	if FailureReason, ok := ContactsbulkoperationjobMap["failureReason"].(map[string]interface{}); ok {
		FailureReasonString, _ := json.Marshal(FailureReason)
		json.Unmarshal(FailureReasonString, &o.FailureReason)
	}
	
	if DownloadURI, ok := ContactsbulkoperationjobMap["downloadURI"].(string); ok {
		o.DownloadURI = &DownloadURI
	}
    
	if SelfUri, ok := ContactsbulkoperationjobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Contactsbulkoperationjob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
