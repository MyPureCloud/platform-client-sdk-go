package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Annotation
type Annotation struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
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

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Annotation) SetField(field string, fieldValue interface{}) {
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

func (o Annotation) MarshalJSON() ([]byte, error) {
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
		Alias
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
		Alias:    (Alias)(o),
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
