package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Documentcontentblockwithhighlight
type Documentcontentblockwithhighlight struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// VarType - The type of the paragraph block.
	VarType *string `json:"type,omitempty"`

	// Text - Text. It must contain a value if the type of the block is Text.
	Text *Documenttext `json:"text,omitempty"`

	// Image - Image. It must contain a value if the type of the block is Image.
	Image *Documentbodyimage `json:"image,omitempty"`

	// Video - Video. It must contain a value if the type of the block is Video.
	Video *Documentbodyvideo `json:"video,omitempty"`

	// AnswerHighlight - The block highlight data.
	AnswerHighlight *Documentcontenthighlightindex `json:"answerHighlight,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Documentcontentblockwithhighlight) SetField(field string, fieldValue interface{}) {
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

func (o Documentcontentblockwithhighlight) MarshalJSON() ([]byte, error) {
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
	type Alias Documentcontentblockwithhighlight
	
	return json.Marshal(&struct { 
		VarType *string `json:"type,omitempty"`
		
		Text *Documenttext `json:"text,omitempty"`
		
		Image *Documentbodyimage `json:"image,omitempty"`
		
		Video *Documentbodyvideo `json:"video,omitempty"`
		
		AnswerHighlight *Documentcontenthighlightindex `json:"answerHighlight,omitempty"`
		Alias
	}{ 
		VarType: o.VarType,
		
		Text: o.Text,
		
		Image: o.Image,
		
		Video: o.Video,
		
		AnswerHighlight: o.AnswerHighlight,
		Alias:    (Alias)(o),
	})
}

func (o *Documentcontentblockwithhighlight) UnmarshalJSON(b []byte) error {
	var DocumentcontentblockwithhighlightMap map[string]interface{}
	err := json.Unmarshal(b, &DocumentcontentblockwithhighlightMap)
	if err != nil {
		return err
	}
	
	if VarType, ok := DocumentcontentblockwithhighlightMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Text, ok := DocumentcontentblockwithhighlightMap["text"].(map[string]interface{}); ok {
		TextString, _ := json.Marshal(Text)
		json.Unmarshal(TextString, &o.Text)
	}
	
	if Image, ok := DocumentcontentblockwithhighlightMap["image"].(map[string]interface{}); ok {
		ImageString, _ := json.Marshal(Image)
		json.Unmarshal(ImageString, &o.Image)
	}
	
	if Video, ok := DocumentcontentblockwithhighlightMap["video"].(map[string]interface{}); ok {
		VideoString, _ := json.Marshal(Video)
		json.Unmarshal(VideoString, &o.Video)
	}
	
	if AnswerHighlight, ok := DocumentcontentblockwithhighlightMap["answerHighlight"].(map[string]interface{}); ok {
		AnswerHighlightString, _ := json.Marshal(AnswerHighlight)
		json.Unmarshal(AnswerHighlightString, &o.AnswerHighlight)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Documentcontentblockwithhighlight) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
