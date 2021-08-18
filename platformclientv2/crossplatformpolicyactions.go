package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Crossplatformpolicyactions
type Crossplatformpolicyactions struct { 
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

func (u *Crossplatformpolicyactions) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		RetainRecording: u.RetainRecording,
		
		DeleteRecording: u.DeleteRecording,
		
		AlwaysDelete: u.AlwaysDelete,
		
		AssignEvaluations: u.AssignEvaluations,
		
		AssignMeteredEvaluations: u.AssignMeteredEvaluations,
		
		AssignMeteredAssignmentByAgent: u.AssignMeteredAssignmentByAgent,
		
		AssignCalibrations: u.AssignCalibrations,
		
		RetentionDuration: u.RetentionDuration,
		
		MediaTranscriptions: u.MediaTranscriptions,
		
		IntegrationExport: u.IntegrationExport,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Crossplatformpolicyactions) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
