package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Textbotpromptsegment - Data for a single bot flow prompt segment.
type Textbotpromptsegment struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Text - The text of this prompt segment.
	Text *string `json:"text,omitempty"`

	// VarType - The segment type which describes any semantics about the 'text' and also indicates which other field might include additional relevant info.
	VarType *string `json:"type,omitempty"`

	// Format - Additional details describing the segment’s contents, which the client should honour where possible.
	Format *Format `json:"format,omitempty"`

	// Content - Details to display Rich Media content. This is only populated when the segment 'type' is 'Rich Media'.
	Content *[]Messagecontent `json:"content,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Textbotpromptsegment) SetField(field string, fieldValue interface{}) {
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

func (o Textbotpromptsegment) MarshalJSON() ([]byte, error) {
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
	type Alias Textbotpromptsegment
	
	return json.Marshal(&struct { 
		Text *string `json:"text,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Format *Format `json:"format,omitempty"`
		
		Content *[]Messagecontent `json:"content,omitempty"`
		Alias
	}{ 
		Text: o.Text,
		
		VarType: o.VarType,
		
		Format: o.Format,
		
		Content: o.Content,
		Alias:    (Alias)(o),
	})
}

func (o *Textbotpromptsegment) UnmarshalJSON(b []byte) error {
	var TextbotpromptsegmentMap map[string]interface{}
	err := json.Unmarshal(b, &TextbotpromptsegmentMap)
	if err != nil {
		return err
	}
	
	if Text, ok := TextbotpromptsegmentMap["text"].(string); ok {
		o.Text = &Text
	}
    
	if VarType, ok := TextbotpromptsegmentMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Format, ok := TextbotpromptsegmentMap["format"].(map[string]interface{}); ok {
		FormatString, _ := json.Marshal(Format)
		json.Unmarshal(FormatString, &o.Format)
	}
	
	if Content, ok := TextbotpromptsegmentMap["content"].([]interface{}); ok {
		ContentString, _ := json.Marshal(Content)
		json.Unmarshal(ContentString, &o.Content)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Textbotpromptsegment) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
