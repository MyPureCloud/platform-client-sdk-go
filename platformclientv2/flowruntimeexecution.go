package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowruntimeexecution - Details about the current state of a Flow execution
type Flowruntimeexecution struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The flow execution ID
	Id *string `json:"id,omitempty"`

	// Name - The flow execution name.
	Name *string `json:"name,omitempty"`

	// FlowVersion - The Version of the flow definition of the flow execution.
	FlowVersion *Flowversion `json:"flowVersion,omitempty"`

	// DateLaunched - The time the flow was launched. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateLaunched *time.Time `json:"dateLaunched,omitempty"`

	// Status - The flow's running status, which indicates whether the flow is running normally or completed, etc.
	Status *string `json:"status,omitempty"`

	// DateCompleted - The time the flow completed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCompleted *time.Time `json:"dateCompleted,omitempty"`

	// CompletionReason - The completion reason set at the flow completion time, if applicable.
	CompletionReason *string `json:"completionReason,omitempty"`

	// FlowErrorInfo - Additional information if the flow is in error
	FlowErrorInfo *Errorbody `json:"flowErrorInfo,omitempty"`

	// OutputData - List of the flow's output variables, if any. Output variables are only supplied for Completed flows.
	OutputData *map[string]interface{} `json:"outputData,omitempty"`

	// Conversation - The conversation to which this Flow execution is related
	Conversation *Domainentityref `json:"conversation,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowruntimeexecution) SetField(field string, fieldValue interface{}) {
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

func (o Flowruntimeexecution) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateLaunched","DateCompleted", }
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
	type Alias Flowruntimeexecution
	
	DateLaunched := new(string)
	if o.DateLaunched != nil {
		
		*DateLaunched = timeutil.Strftime(o.DateLaunched, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateLaunched = nil
	}
	
	DateCompleted := new(string)
	if o.DateCompleted != nil {
		
		*DateCompleted = timeutil.Strftime(o.DateCompleted, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCompleted = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		FlowVersion *Flowversion `json:"flowVersion,omitempty"`
		
		DateLaunched *string `json:"dateLaunched,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		DateCompleted *string `json:"dateCompleted,omitempty"`
		
		CompletionReason *string `json:"completionReason,omitempty"`
		
		FlowErrorInfo *Errorbody `json:"flowErrorInfo,omitempty"`
		
		OutputData *map[string]interface{} `json:"outputData,omitempty"`
		
		Conversation *Domainentityref `json:"conversation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		FlowVersion: o.FlowVersion,
		
		DateLaunched: DateLaunched,
		
		Status: o.Status,
		
		DateCompleted: DateCompleted,
		
		CompletionReason: o.CompletionReason,
		
		FlowErrorInfo: o.FlowErrorInfo,
		
		OutputData: o.OutputData,
		
		Conversation: o.Conversation,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flowruntimeexecution) UnmarshalJSON(b []byte) error {
	var FlowruntimeexecutionMap map[string]interface{}
	err := json.Unmarshal(b, &FlowruntimeexecutionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowruntimeexecutionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FlowruntimeexecutionMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if FlowVersion, ok := FlowruntimeexecutionMap["flowVersion"].(map[string]interface{}); ok {
		FlowVersionString, _ := json.Marshal(FlowVersion)
		json.Unmarshal(FlowVersionString, &o.FlowVersion)
	}
	
	if dateLaunchedString, ok := FlowruntimeexecutionMap["dateLaunched"].(string); ok {
		DateLaunched, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateLaunchedString)
		o.DateLaunched = &DateLaunched
	}
	
	if Status, ok := FlowruntimeexecutionMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if dateCompletedString, ok := FlowruntimeexecutionMap["dateCompleted"].(string); ok {
		DateCompleted, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCompletedString)
		o.DateCompleted = &DateCompleted
	}
	
	if CompletionReason, ok := FlowruntimeexecutionMap["completionReason"].(string); ok {
		o.CompletionReason = &CompletionReason
	}
    
	if FlowErrorInfo, ok := FlowruntimeexecutionMap["flowErrorInfo"].(map[string]interface{}); ok {
		FlowErrorInfoString, _ := json.Marshal(FlowErrorInfo)
		json.Unmarshal(FlowErrorInfoString, &o.FlowErrorInfo)
	}
	
	if OutputData, ok := FlowruntimeexecutionMap["outputData"].(map[string]interface{}); ok {
		OutputDataString, _ := json.Marshal(OutputData)
		json.Unmarshal(OutputDataString, &o.OutputData)
	}
	
	if Conversation, ok := FlowruntimeexecutionMap["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if SelfUri, ok := FlowruntimeexecutionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowruntimeexecution) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
