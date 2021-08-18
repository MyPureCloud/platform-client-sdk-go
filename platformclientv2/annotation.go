package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Annotation
type Annotation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// VarType
	VarType *string `json:"type,omitempty"`


	// Location - Offset of annotation in milliseconds.
	Location *int `json:"location,omitempty"`


	// DurationMs - Duration of annotation in milliseconds.
	DurationMs *int `json:"durationMs,omitempty"`


	// AbsoluteLocation - Offset of annotation (milliseconds) from start of recording.
	AbsoluteLocation *int `json:"absoluteLocation,omitempty"`


	// AbsoluteDurationMs - Duration of annotation (milliseconds).
	AbsoluteDurationMs *int `json:"absoluteDurationMs,omitempty"`


	// RecordingLocation - Offset of annotation (milliseconds) from start of recording, adjusted for any recording cuts
	RecordingLocation *int `json:"recordingLocation,omitempty"`


	// RecordingDurationMs - Duration of annotation (milliseconds), adjusted for any recording cuts.
	RecordingDurationMs *int `json:"recordingDurationMs,omitempty"`


	// User - User that created this annotation (if any).
	User *User `json:"user,omitempty"`


	// Description - Text of annotation.
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Annotation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Annotation

	

	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Location *int `json:"location,omitempty"`
		
		DurationMs *int `json:"durationMs,omitempty"`
		
		AbsoluteLocation *int `json:"absoluteLocation,omitempty"`
		
		AbsoluteDurationMs *int `json:"absoluteDurationMs,omitempty"`
		
		RecordingLocation *int `json:"recordingLocation,omitempty"`
		
		RecordingDurationMs *int `json:"recordingDurationMs,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		Description *string `json:"description,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		VarType: u.VarType,
		
		Location: u.Location,
		
		DurationMs: u.DurationMs,
		
		AbsoluteLocation: u.AbsoluteLocation,
		
		AbsoluteDurationMs: u.AbsoluteDurationMs,
		
		RecordingLocation: u.RecordingLocation,
		
		RecordingDurationMs: u.RecordingDurationMs,
		
		User: u.User,
		
		Description: u.Description,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Annotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
