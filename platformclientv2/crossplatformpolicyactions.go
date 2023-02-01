package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformpolicyactions
type Crossplatformpolicyactions struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// RetainRecording - true to retain the recording associated with the conversation. Default = true
	RetainRecording *bool `json:"retainRecording,omitempty"`

	// DeleteRecording - true to delete the recording associated with the conversation. If retainRecording = true, this will be ignored. Default = false
	DeleteRecording *bool `json:"deleteRecording,omitempty"`

	// AlwaysDelete - true to delete the recording associated with the conversation regardless of the values of retainRecording or deleteRecording. Default = false
	AlwaysDelete *bool `json:"alwaysDelete,omitempty"`

	// AssignEvaluations
	AssignEvaluations *[]Evaluationassignment `json:"assignEvaluations,omitempty"`

	// AssignMeteredEvaluations
	AssignMeteredEvaluations *[]Meteredevaluationassignment `json:"assignMeteredEvaluations,omitempty"`

	// AssignMeteredAssignmentByAgent
	AssignMeteredAssignmentByAgent *[]Meteredassignmentbyagent `json:"assignMeteredAssignmentByAgent,omitempty"`

	// AssignCalibrations
	AssignCalibrations *[]Calibrationassignment `json:"assignCalibrations,omitempty"`

	// RetentionDuration
	RetentionDuration *Retentionduration `json:"retentionDuration,omitempty"`

	// MediaTranscriptions
	MediaTranscriptions *[]Mediatranscription `json:"mediaTranscriptions,omitempty"`

	// IntegrationExport - Policy action for exporting recordings using an integration to 3rd party s3.
	IntegrationExport *Integrationexport `json:"integrationExport,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Crossplatformpolicyactions) SetField(field string, fieldValue interface{}) {
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

func (o Crossplatformpolicyactions) MarshalJSON() ([]byte, error) {
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
	type Alias Crossplatformpolicyactions
	
	return json.Marshal(&struct { 
		RetainRecording *bool `json:"retainRecording,omitempty"`
		
		DeleteRecording *bool `json:"deleteRecording,omitempty"`
		
		AlwaysDelete *bool `json:"alwaysDelete,omitempty"`
		
		AssignEvaluations *[]Evaluationassignment `json:"assignEvaluations,omitempty"`
		
		AssignMeteredEvaluations *[]Meteredevaluationassignment `json:"assignMeteredEvaluations,omitempty"`
		
		AssignMeteredAssignmentByAgent *[]Meteredassignmentbyagent `json:"assignMeteredAssignmentByAgent,omitempty"`
		
		AssignCalibrations *[]Calibrationassignment `json:"assignCalibrations,omitempty"`
		
		RetentionDuration *Retentionduration `json:"retentionDuration,omitempty"`
		
		MediaTranscriptions *[]Mediatranscription `json:"mediaTranscriptions,omitempty"`
		
		IntegrationExport *Integrationexport `json:"integrationExport,omitempty"`
		Alias
	}{ 
		RetainRecording: o.RetainRecording,
		
		DeleteRecording: o.DeleteRecording,
		
		AlwaysDelete: o.AlwaysDelete,
		
		AssignEvaluations: o.AssignEvaluations,
		
		AssignMeteredEvaluations: o.AssignMeteredEvaluations,
		
		AssignMeteredAssignmentByAgent: o.AssignMeteredAssignmentByAgent,
		
		AssignCalibrations: o.AssignCalibrations,
		
		RetentionDuration: o.RetentionDuration,
		
		MediaTranscriptions: o.MediaTranscriptions,
		
		IntegrationExport: o.IntegrationExport,
		Alias:    (Alias)(o),
	})
}

func (o *Crossplatformpolicyactions) UnmarshalJSON(b []byte) error {
	var CrossplatformpolicyactionsMap map[string]interface{}
	err := json.Unmarshal(b, &CrossplatformpolicyactionsMap)
	if err != nil {
		return err
	}
	
	if RetainRecording, ok := CrossplatformpolicyactionsMap["retainRecording"].(bool); ok {
		o.RetainRecording = &RetainRecording
	}
    
	if DeleteRecording, ok := CrossplatformpolicyactionsMap["deleteRecording"].(bool); ok {
		o.DeleteRecording = &DeleteRecording
	}
    
	if AlwaysDelete, ok := CrossplatformpolicyactionsMap["alwaysDelete"].(bool); ok {
		o.AlwaysDelete = &AlwaysDelete
	}
    
	if AssignEvaluations, ok := CrossplatformpolicyactionsMap["assignEvaluations"].([]interface{}); ok {
		AssignEvaluationsString, _ := json.Marshal(AssignEvaluations)
		json.Unmarshal(AssignEvaluationsString, &o.AssignEvaluations)
	}
	
	if AssignMeteredEvaluations, ok := CrossplatformpolicyactionsMap["assignMeteredEvaluations"].([]interface{}); ok {
		AssignMeteredEvaluationsString, _ := json.Marshal(AssignMeteredEvaluations)
		json.Unmarshal(AssignMeteredEvaluationsString, &o.AssignMeteredEvaluations)
	}
	
	if AssignMeteredAssignmentByAgent, ok := CrossplatformpolicyactionsMap["assignMeteredAssignmentByAgent"].([]interface{}); ok {
		AssignMeteredAssignmentByAgentString, _ := json.Marshal(AssignMeteredAssignmentByAgent)
		json.Unmarshal(AssignMeteredAssignmentByAgentString, &o.AssignMeteredAssignmentByAgent)
	}
	
	if AssignCalibrations, ok := CrossplatformpolicyactionsMap["assignCalibrations"].([]interface{}); ok {
		AssignCalibrationsString, _ := json.Marshal(AssignCalibrations)
		json.Unmarshal(AssignCalibrationsString, &o.AssignCalibrations)
	}
	
	if RetentionDuration, ok := CrossplatformpolicyactionsMap["retentionDuration"].(map[string]interface{}); ok {
		RetentionDurationString, _ := json.Marshal(RetentionDuration)
		json.Unmarshal(RetentionDurationString, &o.RetentionDuration)
	}
	
	if MediaTranscriptions, ok := CrossplatformpolicyactionsMap["mediaTranscriptions"].([]interface{}); ok {
		MediaTranscriptionsString, _ := json.Marshal(MediaTranscriptions)
		json.Unmarshal(MediaTranscriptionsString, &o.MediaTranscriptions)
	}
	
	if IntegrationExport, ok := CrossplatformpolicyactionsMap["integrationExport"].(map[string]interface{}); ok {
		IntegrationExportString, _ := json.Marshal(IntegrationExport)
		json.Unmarshal(IntegrationExportString, &o.IntegrationExport)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Crossplatformpolicyactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
