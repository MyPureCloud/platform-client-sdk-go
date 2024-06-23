package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Flowexecutiondataqueryresult - This is the metadata of an executionData entry for a flow.
type Flowexecutiondataqueryresult struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// StartDateTime - The start time for the execution of this flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// EndDateTime - The end time for the execution of this flow. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	EndDateTime *time.Time `json:"endDateTime,omitempty"`

	// FlowId - The id of the flow that was executed.
	FlowId *string `json:"flowId,omitempty"`

	// FlowVersion - The version of the flow that was executed.
	FlowVersion *string `json:"flowVersion,omitempty"`

	// ConversationId - The id of the conversation that executed this flow.
	ConversationId *string `json:"conversationId,omitempty"`

	// FlowType - The type of flow.
	FlowType *string `json:"flowType,omitempty"`

	// FlowErrorReason - If the flow errored out this is the reason.
	FlowErrorReason *string `json:"flowErrorReason,omitempty"`

	// FlowWarningReason - If the flow had a warning, this is the reason.
	FlowWarningReason *string `json:"flowWarningReason,omitempty"`

	// FlowName - The name of the flow.
	FlowName *string `json:"flowName,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Flowexecutiondataqueryresult) SetField(field string, fieldValue interface{}) {
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

func (o Flowexecutiondataqueryresult) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "StartDateTime","EndDateTime", }
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
	type Alias Flowexecutiondataqueryresult
	
	StartDateTime := new(string)
	if o.StartDateTime != nil {
		
		*StartDateTime = timeutil.Strftime(o.StartDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		StartDateTime = nil
	}
	
	EndDateTime := new(string)
	if o.EndDateTime != nil {
		
		*EndDateTime = timeutil.Strftime(o.EndDateTime, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		EndDateTime = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		StartDateTime *string `json:"startDateTime,omitempty"`
		
		EndDateTime *string `json:"endDateTime,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		FlowErrorReason *string `json:"flowErrorReason,omitempty"`
		
		FlowWarningReason *string `json:"flowWarningReason,omitempty"`
		
		FlowName *string `json:"flowName,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		StartDateTime: StartDateTime,
		
		EndDateTime: EndDateTime,
		
		FlowId: o.FlowId,
		
		FlowVersion: o.FlowVersion,
		
		ConversationId: o.ConversationId,
		
		FlowType: o.FlowType,
		
		FlowErrorReason: o.FlowErrorReason,
		
		FlowWarningReason: o.FlowWarningReason,
		
		FlowName: o.FlowName,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Flowexecutiondataqueryresult) UnmarshalJSON(b []byte) error {
	var FlowexecutiondataqueryresultMap map[string]interface{}
	err := json.Unmarshal(b, &FlowexecutiondataqueryresultMap)
	if err != nil {
		return err
	}
	
	if Id, ok := FlowexecutiondataqueryresultMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := FlowexecutiondataqueryresultMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if startDateTimeString, ok := FlowexecutiondataqueryresultMap["startDateTime"].(string); ok {
		StartDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", startDateTimeString)
		o.StartDateTime = &StartDateTime
	}
	
	if endDateTimeString, ok := FlowexecutiondataqueryresultMap["endDateTime"].(string); ok {
		EndDateTime, _ := time.Parse("2006-01-02T15:04:05.999999Z", endDateTimeString)
		o.EndDateTime = &EndDateTime
	}
	
	if FlowId, ok := FlowexecutiondataqueryresultMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if FlowVersion, ok := FlowexecutiondataqueryresultMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if ConversationId, ok := FlowexecutiondataqueryresultMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if FlowType, ok := FlowexecutiondataqueryresultMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowErrorReason, ok := FlowexecutiondataqueryresultMap["flowErrorReason"].(string); ok {
		o.FlowErrorReason = &FlowErrorReason
	}
    
	if FlowWarningReason, ok := FlowexecutiondataqueryresultMap["flowWarningReason"].(string); ok {
		o.FlowWarningReason = &FlowWarningReason
	}
    
	if FlowName, ok := FlowexecutiondataqueryresultMap["flowName"].(string); ok {
		o.FlowName = &FlowName
	}
    
	if SelfUri, ok := FlowexecutiondataqueryresultMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Flowexecutiondataqueryresult) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
