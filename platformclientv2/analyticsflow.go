package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Analyticsflow
type Analyticsflow struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// EndingLanguage - Flow ending language, e.g. en-us
	EndingLanguage *string `json:"endingLanguage,omitempty"`

	// EntryReason - The particular entry reason for this flow, e.g. an address, userId, or flowId
	EntryReason *string `json:"entryReason,omitempty"`

	// EntryType - The entry type for this flow, e.g. dnis, dialer, agent, flow, or direct
	EntryType *string `json:"entryType,omitempty"`

	// ExitReason - The exit reason for this flow, e.g. DISCONNECT
	ExitReason *string `json:"exitReason,omitempty"`

	// FlowId - The unique identifier of this flow
	FlowId *string `json:"flowId,omitempty"`

	// FlowName - The name of this flow at the time of flow execution
	FlowName *string `json:"flowName,omitempty"`

	// FlowType - The type of this flow
	FlowType *string `json:"flowType,omitempty"`

	// FlowVersion - The version of this flow
	FlowVersion *string `json:"flowVersion,omitempty"`

	// IssuedCallback - Flag indicating whether the flow issued a callback
	IssuedCallback *bool `json:"issuedCallback,omitempty"`

	// RecognitionFailureReason - The recognition failure reason causing to exit/disconnect
	RecognitionFailureReason *string `json:"recognitionFailureReason,omitempty"`

	// StartingLanguage - Flow starting language, e.g. en-us
	StartingLanguage *string `json:"startingLanguage,omitempty"`

	// TransferTargetAddress - The address of a flow transfer target, e.g. a phone number, an email address, or a queueId
	TransferTargetAddress *string `json:"transferTargetAddress,omitempty"`

	// TransferTargetName - The name of a flow transfer target
	TransferTargetName *string `json:"transferTargetName,omitempty"`

	// TransferType - The type of transfer for flows that ended with a transfer
	TransferType *string `json:"transferType,omitempty"`

	// Outcomes - Flow outcomes
	Outcomes *[]Analyticsflowoutcome `json:"outcomes,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Analyticsflow) SetField(field string, fieldValue interface{}) {
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

func (o Analyticsflow) MarshalJSON() ([]byte, error) {
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
	type Alias Analyticsflow
	
	return json.Marshal(&struct { 
		EndingLanguage *string `json:"endingLanguage,omitempty"`
		
		EntryReason *string `json:"entryReason,omitempty"`
		
		EntryType *string `json:"entryType,omitempty"`
		
		ExitReason *string `json:"exitReason,omitempty"`
		
		FlowId *string `json:"flowId,omitempty"`
		
		FlowName *string `json:"flowName,omitempty"`
		
		FlowType *string `json:"flowType,omitempty"`
		
		FlowVersion *string `json:"flowVersion,omitempty"`
		
		IssuedCallback *bool `json:"issuedCallback,omitempty"`
		
		RecognitionFailureReason *string `json:"recognitionFailureReason,omitempty"`
		
		StartingLanguage *string `json:"startingLanguage,omitempty"`
		
		TransferTargetAddress *string `json:"transferTargetAddress,omitempty"`
		
		TransferTargetName *string `json:"transferTargetName,omitempty"`
		
		TransferType *string `json:"transferType,omitempty"`
		
		Outcomes *[]Analyticsflowoutcome `json:"outcomes,omitempty"`
		Alias
	}{ 
		EndingLanguage: o.EndingLanguage,
		
		EntryReason: o.EntryReason,
		
		EntryType: o.EntryType,
		
		ExitReason: o.ExitReason,
		
		FlowId: o.FlowId,
		
		FlowName: o.FlowName,
		
		FlowType: o.FlowType,
		
		FlowVersion: o.FlowVersion,
		
		IssuedCallback: o.IssuedCallback,
		
		RecognitionFailureReason: o.RecognitionFailureReason,
		
		StartingLanguage: o.StartingLanguage,
		
		TransferTargetAddress: o.TransferTargetAddress,
		
		TransferTargetName: o.TransferTargetName,
		
		TransferType: o.TransferType,
		
		Outcomes: o.Outcomes,
		Alias:    (Alias)(o),
	})
}

func (o *Analyticsflow) UnmarshalJSON(b []byte) error {
	var AnalyticsflowMap map[string]interface{}
	err := json.Unmarshal(b, &AnalyticsflowMap)
	if err != nil {
		return err
	}
	
	if EndingLanguage, ok := AnalyticsflowMap["endingLanguage"].(string); ok {
		o.EndingLanguage = &EndingLanguage
	}
    
	if EntryReason, ok := AnalyticsflowMap["entryReason"].(string); ok {
		o.EntryReason = &EntryReason
	}
    
	if EntryType, ok := AnalyticsflowMap["entryType"].(string); ok {
		o.EntryType = &EntryType
	}
    
	if ExitReason, ok := AnalyticsflowMap["exitReason"].(string); ok {
		o.ExitReason = &ExitReason
	}
    
	if FlowId, ok := AnalyticsflowMap["flowId"].(string); ok {
		o.FlowId = &FlowId
	}
    
	if FlowName, ok := AnalyticsflowMap["flowName"].(string); ok {
		o.FlowName = &FlowName
	}
    
	if FlowType, ok := AnalyticsflowMap["flowType"].(string); ok {
		o.FlowType = &FlowType
	}
    
	if FlowVersion, ok := AnalyticsflowMap["flowVersion"].(string); ok {
		o.FlowVersion = &FlowVersion
	}
    
	if IssuedCallback, ok := AnalyticsflowMap["issuedCallback"].(bool); ok {
		o.IssuedCallback = &IssuedCallback
	}
    
	if RecognitionFailureReason, ok := AnalyticsflowMap["recognitionFailureReason"].(string); ok {
		o.RecognitionFailureReason = &RecognitionFailureReason
	}
    
	if StartingLanguage, ok := AnalyticsflowMap["startingLanguage"].(string); ok {
		o.StartingLanguage = &StartingLanguage
	}
    
	if TransferTargetAddress, ok := AnalyticsflowMap["transferTargetAddress"].(string); ok {
		o.TransferTargetAddress = &TransferTargetAddress
	}
    
	if TransferTargetName, ok := AnalyticsflowMap["transferTargetName"].(string); ok {
		o.TransferTargetName = &TransferTargetName
	}
    
	if TransferType, ok := AnalyticsflowMap["transferType"].(string); ok {
		o.TransferType = &TransferType
	}
    
	if Outcomes, ok := AnalyticsflowMap["outcomes"].([]interface{}); ok {
		OutcomesString, _ := json.Marshal(Outcomes)
		json.Unmarshal(OutcomesString, &o.Outcomes)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Analyticsflow) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
