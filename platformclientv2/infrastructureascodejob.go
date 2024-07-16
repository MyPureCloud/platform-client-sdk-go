package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Infrastructureascodejob - Information about a CX infrastructure as code job
type Infrastructureascodejob struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// DryRun - Whether or not the job was a dry run
	DryRun *bool `json:"dryRun,omitempty"`

	// AcceleratorId - Accelerator associated with the job
	AcceleratorId *string `json:"acceleratorId,omitempty"`

	// DateSubmitted - Date and time on which job was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateSubmitted *time.Time `json:"dateSubmitted,omitempty"`

	// SubmittedBy - User who submitted the job
	SubmittedBy *Userreference `json:"submittedBy,omitempty"`

	// Status - Job status
	Status *string `json:"status,omitempty"`

	// ErrorInfo - Information about errors, if any
	ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`

	// Results - The output results of the terraform job
	Results *string `json:"results,omitempty"`

	// RollbackResults - The results of rolling back the job if there were errors.  Not returned if job was successful.
	RollbackResults *string `json:"rollbackResults,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Infrastructureascodejob) SetField(field string, fieldValue interface{}) {
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

func (o Infrastructureascodejob) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateSubmitted", }
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
	type Alias Infrastructureascodejob
	
	DateSubmitted := new(string)
	if o.DateSubmitted != nil {
		
		*DateSubmitted = timeutil.Strftime(o.DateSubmitted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateSubmitted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		DryRun *bool `json:"dryRun,omitempty"`
		
		AcceleratorId *string `json:"acceleratorId,omitempty"`
		
		DateSubmitted *string `json:"dateSubmitted,omitempty"`
		
		SubmittedBy *Userreference `json:"submittedBy,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorInfo *Errorinfo `json:"errorInfo,omitempty"`
		
		Results *string `json:"results,omitempty"`
		
		RollbackResults *string `json:"rollbackResults,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		DryRun: o.DryRun,
		
		AcceleratorId: o.AcceleratorId,
		
		DateSubmitted: DateSubmitted,
		
		SubmittedBy: o.SubmittedBy,
		
		Status: o.Status,
		
		ErrorInfo: o.ErrorInfo,
		
		Results: o.Results,
		
		RollbackResults: o.RollbackResults,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Infrastructureascodejob) UnmarshalJSON(b []byte) error {
	var InfrastructureascodejobMap map[string]interface{}
	err := json.Unmarshal(b, &InfrastructureascodejobMap)
	if err != nil {
		return err
	}
	
	if Id, ok := InfrastructureascodejobMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if DryRun, ok := InfrastructureascodejobMap["dryRun"].(bool); ok {
		o.DryRun = &DryRun
	}
    
	if AcceleratorId, ok := InfrastructureascodejobMap["acceleratorId"].(string); ok {
		o.AcceleratorId = &AcceleratorId
	}
    
	if dateSubmittedString, ok := InfrastructureascodejobMap["dateSubmitted"].(string); ok {
		DateSubmitted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateSubmittedString)
		o.DateSubmitted = &DateSubmitted
	}
	
	if SubmittedBy, ok := InfrastructureascodejobMap["submittedBy"].(map[string]interface{}); ok {
		SubmittedByString, _ := json.Marshal(SubmittedBy)
		json.Unmarshal(SubmittedByString, &o.SubmittedBy)
	}
	
	if Status, ok := InfrastructureascodejobMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorInfo, ok := InfrastructureascodejobMap["errorInfo"].(map[string]interface{}); ok {
		ErrorInfoString, _ := json.Marshal(ErrorInfo)
		json.Unmarshal(ErrorInfoString, &o.ErrorInfo)
	}
	
	if Results, ok := InfrastructureascodejobMap["results"].(string); ok {
		o.Results = &Results
	}
    
	if RollbackResults, ok := InfrastructureascodejobMap["rollbackResults"].(string); ok {
		o.RollbackResults = &RollbackResults
	}
    
	if SelfUri, ok := InfrastructureascodejobMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Infrastructureascodejob) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
