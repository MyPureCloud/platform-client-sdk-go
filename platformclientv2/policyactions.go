package platformclientv2
import (
	"encoding/json"
)

// Policyactions
type Policyactions struct { 
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


	// AssignSurveys
	AssignSurveys *[]Surveyassignment `json:"assignSurveys,omitempty"`


	// RetentionDuration
	RetentionDuration *Retentionduration `json:"retentionDuration,omitempty"`


	// InitiateScreenRecording
	InitiateScreenRecording *Initiatescreenrecording `json:"initiateScreenRecording,omitempty"`


	// MediaTranscriptions
	MediaTranscriptions *[]Mediatranscription `json:"mediaTranscriptions,omitempty"`


	// IntegrationExport - Policy action for exporting recordings using an integration to 3rd party s3.
	IntegrationExport *Integrationexport `json:"integrationExport,omitempty"`

}

// String returns a JSON representation of the model
func (o *Policyactions) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
