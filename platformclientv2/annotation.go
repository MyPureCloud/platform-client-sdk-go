package platformclientv2
import (
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

// String returns a JSON representation of the model
func (o *Annotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
