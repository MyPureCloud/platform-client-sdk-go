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


	// AbsoluteLocation - Offset of annotation (milliseconds) from start of recording (after removing the cumulative duration of all pauses).
	AbsoluteLocation *int `json:"absoluteLocation,omitempty"`


	// AbsoluteDurationMs - Duration of annotation (milliseconds).
	AbsoluteDurationMs *int `json:"absoluteDurationMs,omitempty"`


	// RecordingLocation - Offset of annotation (milliseconds) from start of recording, adjusted for any recording cuts
	RecordingLocation *int `json:"recordingLocation,omitempty"`


	// RecordingDurationMs - Duration of annotation (milliseconds), adjusted for any recording cuts.
	RecordingDurationMs *int `json:"recordingDurationMs,omitempty"`


	// User - User that created this annotation (if any).
	User *User `json:"user,omitempty"`


	// Description - Text of annotation. Maximum character limit is 500.
	Description *string `json:"description,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Annotation) MarshalJSON() ([]byte, error) {
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
		Id: o.Id,
		
		Name: o.Name,
		
		VarType: o.VarType,
		
		Location: o.Location,
		
		DurationMs: o.DurationMs,
		
		AbsoluteLocation: o.AbsoluteLocation,
		
		AbsoluteDurationMs: o.AbsoluteDurationMs,
		
		RecordingLocation: o.RecordingLocation,
		
		RecordingDurationMs: o.RecordingDurationMs,
		
		User: o.User,
		
		Description: o.Description,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Annotation) UnmarshalJSON(b []byte) error {
	var AnnotationMap map[string]interface{}
	err := json.Unmarshal(b, &AnnotationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := AnnotationMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := AnnotationMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if VarType, ok := AnnotationMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Location, ok := AnnotationMap["location"].(float64); ok {
		LocationInt := int(Location)
		o.Location = &LocationInt
	}
	
	if DurationMs, ok := AnnotationMap["durationMs"].(float64); ok {
		DurationMsInt := int(DurationMs)
		o.DurationMs = &DurationMsInt
	}
	
	if AbsoluteLocation, ok := AnnotationMap["absoluteLocation"].(float64); ok {
		AbsoluteLocationInt := int(AbsoluteLocation)
		o.AbsoluteLocation = &AbsoluteLocationInt
	}
	
	if AbsoluteDurationMs, ok := AnnotationMap["absoluteDurationMs"].(float64); ok {
		AbsoluteDurationMsInt := int(AbsoluteDurationMs)
		o.AbsoluteDurationMs = &AbsoluteDurationMsInt
	}
	
	if RecordingLocation, ok := AnnotationMap["recordingLocation"].(float64); ok {
		RecordingLocationInt := int(RecordingLocation)
		o.RecordingLocation = &RecordingLocationInt
	}
	
	if RecordingDurationMs, ok := AnnotationMap["recordingDurationMs"].(float64); ok {
		RecordingDurationMsInt := int(RecordingDurationMs)
		o.RecordingDurationMs = &RecordingDurationMsInt
	}
	
	if User, ok := AnnotationMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if Description, ok := AnnotationMap["description"].(string); ok {
		o.Description = &Description
	}
    
	if SelfUri, ok := AnnotationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Annotation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
