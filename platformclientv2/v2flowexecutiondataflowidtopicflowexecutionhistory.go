package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2flowexecutiondataflowidtopicflowexecutionhistory
type V2flowexecutiondataflowidtopicflowexecutionhistory struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ExecutionId - The execution identifier which represents this unique instance of the flow that ran.
	ExecutionId *string `json:"executionId,omitempty"`

	// ConversationId - The Genesys Cloud conversation identifier associated with this flow instance execution data.
	ConversationId *string `json:"conversationId,omitempty"`

	// DivisionId - The division identifier for the division associated with the flow for this flow instance.
	DivisionId *string `json:"divisionId,omitempty"`

	// EndDateTime - The end date time for this flow instance execution data.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`

	// Endpoint - The public endpoint a user can use to retrieve the full historical log from the service.
	Endpoint *string `json:"endpoint,omitempty"`

	// Errors - If the flow invoked error handling, an array of errors.
	Errors *[]V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"errors,omitempty"`

	// Execution - An array of execution items that describe what happened when an Architect flow action container ran such as a flow, task, state or bot.
	Execution *[]V2flowexecutiondataflowidtopicexecution `json:"execution,omitempty"`

	// FlowExitReason - Provides information about why a flow ended.
	FlowExitReason *string `json:"flowExitReason,omitempty"`

	// FlowId - The flow identifier for this flow instance execution data.  This is the identifier of the flow configuration that users load up in Architect.
	FlowId *string `json:"flowId,omitempty"`

	// FlowIsDebug - Whether the flow that ran for this flow instance execution data was in debug mode.
	FlowIsDebug *bool `json:"flowIsDebug,omitempty"`

	// ExecutionItemsTruncated - If true, the execution items in this event have been truncated to be deliverable.
	ExecutionItemsTruncated *bool `json:"executionItemsTruncated,omitempty"`

	// FlowType - The flow type of the Architect flow that was run.
	FlowType *string `json:"flowType,omitempty"`

	// FlowVersion - The version of the flow for this flow instance execution data. Typically this is a numeric value like 1.0 represented as a string but can also be 'debug'
	FlowVersion *string `json:"flowVersion,omitempty"`

	// MessageType - If applicable, the type of message platform from which the message originated.
	MessageType *string `json:"messageType,omitempty"`

	// InvokingContext
	InvokingContext *V2flowexecutiondataflowidtopicinvokingcontext `json:"invokingContext,omitempty"`

	// StartDateTime - The start date time for this flow instance execution data.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// Warnings - If the flow encountered a warning during execution, this is an array of the warnings.
	Warnings *[]V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"warnings,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2flowexecutiondataflowidtopicflowexecutionhistory) SetField(field string, fieldValue interface{}) {
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

func (o V2flowexecutiondataflowidtopicflowexecutionhistory) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "EndDateTime","StartDateTime", }
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
	type Alias V2flowexecutiondataflowidtopicflowexecutionhistory
	
	EndDateTime := new(string)
	if o.EndDateTime != nil {
		
		*EndDateTime = timeutil.Strftime(o.EndDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDateTime = nil
	}
	
	StartDateTime := new(string)
	if o.StartDateTime != nil {
		
		*StartDateTime = timeutil.Strftime(o.StartDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDateTime = nil
	}
	
	return json.Marshal(&struct { 
		ExecutionId *string `json:"executionId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		DivisionId *string `json:"divisionId,omitempty"`
		
		EndDateTime *string `json:"endDateTime,omitempty"`
		
		Endpoint *string `json:"endpoint,omitempty"`
		
		Errors *[]V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"errors,omitempty"`
		
		Execution *[]V2flowexecutiondataflowidtopicexecution `json:"execution,omitempty"`
		
		FlowExitReason *string `json:"flowExitReason,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		FlowIsDebug *bool `json:"flowIsDebug,omitempty"`
		
		ExecutionItemsTruncated *bool `json:"executionItemsTruncated,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		MessageType *string `json:"messageType,omitempty"`
		
		InvokingContext *V2flowexecutiondataflowidtopicinvokingcontext `json:"invokingContext,omitempty"`
		
		StartDateTime *string `json:"startDateTime,omitempty"`
		
		Warnings *[]V2flowexecutiondataflowidtopicflowerrorwarninginfo `json:"warnings,omitempty"`
		Alias
	}{ 
		ExecutionId: o.ExecutionId,
		
		ConversationId: o.ConversationId,
		
		DivisionId: o.DivisionId,
		
		EndDateTime: EndDateTime,
		
		Endpoint: o.Endpoint,
		
		Errors: o.Errors,
		
		Execution: o.Execution,
		
		FlowExitReason: o.FlowExitReason,
		
		FlowId: o.FlowId,
		
		FlowIsDebug: o.FlowIsDebug,
		
		ExecutionItemsTruncated: o.ExecutionItemsTruncated,
		
		FlowType: o.FlowType,
		
		FlowVersion: o.FlowVersion,
		
		MessageType: o.MessageType,
		
		InvokingContext: o.InvokingContext,
		
		StartDateTime: StartDateTime,
		
		Warnings: o.Warnings,
		Alias:    (Alias)(o),
	})
}

func (o *V2flowexecutiondataflowidtopicflowexecutionhistory) UnmarshalJSON(b []byte) error {
	var V2flowexecutiondataflowidtopicflowexecutionhistoryMap map[string]interface{}
	err := json.Unmarshal(b, &V2flowexecutiondataflowidtopicflowexecutionhistoryMap)
	if err != nil {
		return err
	}
	
	if ExecutionId, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["executionId"].(string); ok {
		o.ExecutionId = &ExecutionId
	}
    
	if ConversationId, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if DivisionId, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["divisionId"].(string); ok {
		o.DivisionId = &DivisionId
	}
    
	if endDateTimeString, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["endDateTime"].(string); ok {
		EndDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateTimeString)
		o.EndDateTime = &EndDateTime
	}
	
	if Endpoint, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["endpoint"].(string); ok {
		o.Endpoint = &Endpoint
	}
    
	if Errors, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["errors"].([]interface{}); ok {
		ErrorsString, _ := json.Marshal(Errors)
		json.Unmarshal(ErrorsString, &o.Errors)
	}
	
	if Execution, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["execution"].([]interface{}); ok {
		ExecutionString, _ := json.Marshal(Execution)
		json.Unmarshal(ExecutionString, &o.Execution)
	}
	
	if FlowExitReason, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["flowExitReason"].(string); ok {
		o.FlowExitReason = &FlowExitReason
	}
    
	if FlowId, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if FlowIsDebug, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["flowIsDebug"].(bool); ok {
		o.FlowIsDebug = &FlowIsDebug
	}
    
	if ExecutionItemsTruncated, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["executionItemsTruncated"].(bool); ok {
		o.ExecutionItemsTruncated = &ExecutionItemsTruncated
	}
    
	if FlowType, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowVersion, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if MessageType, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["messageType"].(string); ok {
		o.MessageType = &MessageType
	}
    
	if InvokingContext, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["invokingContext"].(map[string]interface{}); ok {
		InvokingContextString, _ := json.Marshal(InvokingContext)
		json.Unmarshal(InvokingContextString, &o.InvokingContext)
	}
	
	if startDateTimeString, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["startDateTime"].(string); ok {
		StartDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateTimeString)
		o.StartDateTime = &StartDateTime
	}
	
	if Warnings, ok := V2flowexecutiondataflowidtopicflowexecutionhistoryMap["warnings"].([]interface{}); ok {
		WarningsString, _ := json.Marshal(Warnings)
		json.Unmarshal(WarningsString, &o.Warnings)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2flowexecutiondataflowidtopicflowexecutionhistory) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
