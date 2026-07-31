package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Reportingturntoolcall
type Reportingturntoolcall struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// ToolId - Represents the identifier of the tool called.
	ToolId *string `json:"toolId,omitempty"`

	// ToolName - Represents the name of the tool used in the event.
	ToolName *string `json:"toolName,omitempty"`

	// ToolType - Represents the type of tool used in the event.
	ToolType *string `json:"toolType,omitempty"`

	// TargetId - Represents the identifier of the target that the tool is using.
	TargetId *string `json:"targetId,omitempty"`

	// Status - Represents whether the tool call was successful or not.
	Status *string `json:"status,omitempty"`

	// ErrorText - Represents the error returned by the tool in the event of a failure.
	ErrorText *string `json:"errorText,omitempty"`

	// DateInvoked - Represents the starting time of the tool call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateInvoked *time.Time `json:"dateInvoked,omitempty"`

	// LatencyMs - Represents the time it took the tool call to execute.
	LatencyMs *int `json:"latencyMs,omitempty"`

	// Origin - Represents the origin of the tool call.
	Origin *string `json:"origin,omitempty"`

	// KnowledgeMetadata - Represents various metadata of knowledge calls used by the tool if the tool is configured to use knowledge.
	KnowledgeMetadata *Reportingturnknowledgemetadata `json:"knowledgeMetadata,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Reportingturntoolcall) SetField(field string, fieldValue interface{}) {
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

func (o Reportingturntoolcall) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateInvoked", }
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
	type Alias Reportingturntoolcall
	
	DateInvoked := new(string)
	if o.DateInvoked != nil {
		
		*DateInvoked = timeutil.Strftime(o.DateInvoked, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateInvoked = nil
	}
	
	return json.Marshal(&struct { 
		ToolId *string `json:"toolId,omitempty"`
		
		ToolName *string `json:"toolName,omitempty"`
		
		ToolType *string `json:"toolType,omitempty"`
		
		TargetId *string `json:"targetId,omitempty"`
		
		Status *string `json:"status,omitempty"`
		
		ErrorText *string `json:"errorText,omitempty"`
		
		DateInvoked *string `json:"dateInvoked,omitempty"`
		
		LatencyMs *int `json:"latencyMs,omitempty"`
		
		Origin *string `json:"origin,omitempty"`
		
		KnowledgeMetadata *Reportingturnknowledgemetadata `json:"knowledgeMetadata,omitempty"`
		Alias
	}{ 
		ToolId: o.ToolId,
		
		ToolName: o.ToolName,
		
		ToolType: o.ToolType,
		
		TargetId: o.TargetId,
		
		Status: o.Status,
		
		ErrorText: o.ErrorText,
		
		DateInvoked: DateInvoked,
		
		LatencyMs: o.LatencyMs,
		
		Origin: o.Origin,
		
		KnowledgeMetadata: o.KnowledgeMetadata,
		Alias:    (Alias)(o),
	})
}

func (o *Reportingturntoolcall) UnmarshalJSON(b []byte) error {
	var ReportingturntoolcallMap map[string]interface{}
	err := json.Unmarshal(b, &ReportingturntoolcallMap)
	if err != nil {
		return err
	}
	
	if ToolId, ok := ReportingturntoolcallMap["toolId"].(string); ok {
		o.ToolId = &ToolId
	}
    
	if ToolName, ok := ReportingturntoolcallMap["toolName"].(string); ok {
		o.ToolName = &ToolName
	}
    
	if ToolType, ok := ReportingturntoolcallMap["toolType"].(string); ok {
		o.ToolType = &ToolType
	}
    
	if TargetId, ok := ReportingturntoolcallMap["targetId"].(string); ok {
		o.TargetId = &TargetId
	}
    
	if Status, ok := ReportingturntoolcallMap["status"].(string); ok {
		o.Status = &Status
	}
    
	if ErrorText, ok := ReportingturntoolcallMap["errorText"].(string); ok {
		o.ErrorText = &ErrorText
	}
    
	if dateInvokedString, ok := ReportingturntoolcallMap["dateInvoked"].(string); ok {
		DateInvoked, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateInvokedString)
		o.DateInvoked = &DateInvoked
	}
	
	if LatencyMs, ok := ReportingturntoolcallMap["latencyMs"].(float64); ok {
		LatencyMsInt := int(LatencyMs)
		o.LatencyMs = &LatencyMsInt
	}
	
	if Origin, ok := ReportingturntoolcallMap["origin"].(string); ok {
		o.Origin = &Origin
	}
    
	if KnowledgeMetadata, ok := ReportingturntoolcallMap["knowledgeMetadata"].(map[string]interface{}); ok {
		KnowledgeMetadataString, _ := json.Marshal(KnowledgeMetadata)
		json.Unmarshal(KnowledgeMetadataString, &o.KnowledgeMetadata)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Reportingturntoolcall) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
