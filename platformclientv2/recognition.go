package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Recognition
type Recognition struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Recipient - The recipient of the recognition
	Recipient *Userreference `json:"recipient,omitempty"`

	// CreatedBy - The creator of the recognition
	CreatedBy *Userreference `json:"createdBy,omitempty"`

	// DateCreated - The creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// VarType - The type of recognition
	VarType *string `json:"type,omitempty"`

	// Title - The recognition title
	Title *string `json:"title,omitempty"`

	// Note - The recognition note
	Note *string `json:"note,omitempty"`

	// ContextType - The context type (optional)
	ContextType *string `json:"contextType,omitempty"`

	// ContextId - The context id (optional)
	ContextId *string `json:"contextId,omitempty"`

	// DateDisplayed - The displayed date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateDisplayed *time.Time `json:"dateDisplayed,omitempty"`

	// DateAcknowledged - The acknowledged date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	DateAcknowledged *time.Time `json:"dateAcknowledged,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Recognition) SetField(field string, fieldValue interface{}) {
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

func (o Recognition) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateDisplayed","DateAcknowledged", }
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
	type Alias Recognition
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateDisplayed := new(string)
	if o.DateDisplayed != nil {
		
		*DateDisplayed = timeutil.Strftime(o.DateDisplayed, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateDisplayed = nil
	}
	
	DateAcknowledged := new(string)
	if o.DateAcknowledged != nil {
		
		*DateAcknowledged = timeutil.Strftime(o.DateAcknowledged, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateAcknowledged = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Recipient *Userreference `json:"recipient,omitempty"`
		
		CreatedBy *Userreference `json:"createdBy,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		VarType *string `json:"type,omitempty"`
		
		Title *string `json:"title,omitempty"`
		
		Note *string `json:"note,omitempty"`
		
		ContextType *string `json:"contextType,omitempty"`
		
		ContextId *string `json:"contextId,omitempty"`
		
		DateDisplayed *string `json:"dateDisplayed,omitempty"`
		
		DateAcknowledged *string `json:"dateAcknowledged,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Recipient: o.Recipient,
		
		CreatedBy: o.CreatedBy,
		
		DateCreated: DateCreated,
		
		VarType: o.VarType,
		
		Title: o.Title,
		
		Note: o.Note,
		
		ContextType: o.ContextType,
		
		ContextId: o.ContextId,
		
		DateDisplayed: DateDisplayed,
		
		DateAcknowledged: DateAcknowledged,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Recognition) UnmarshalJSON(b []byte) error {
	var RecognitionMap map[string]interface{}
	err := json.Unmarshal(b, &RecognitionMap)
	if err != nil {
		return err
	}
	
	if Id, ok := RecognitionMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Recipient, ok := RecognitionMap["recipient"].(map[string]interface{}); ok {
		RecipientString, _ := json.Marshal(Recipient)
		json.Unmarshal(RecipientString, &o.Recipient)
	}
	
	if CreatedBy, ok := RecognitionMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if dateCreatedString, ok := RecognitionMap["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if VarType, ok := RecognitionMap["type"].(string); ok {
		o.VarType = &VarType
	}
    
	if Title, ok := RecognitionMap["title"].(string); ok {
		o.Title = &Title
	}
    
	if Note, ok := RecognitionMap["note"].(string); ok {
		o.Note = &Note
	}
    
	if ContextType, ok := RecognitionMap["contextType"].(string); ok {
		o.ContextType = &ContextType
	}
    
	if ContextId, ok := RecognitionMap["contextId"].(string); ok {
		o.ContextId = &ContextId
	}
    
	if dateDisplayedString, ok := RecognitionMap["dateDisplayed"].(string); ok {
		DateDisplayed, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateDisplayedString)
		o.DateDisplayed = &DateDisplayed
	}
	
	if dateAcknowledgedString, ok := RecognitionMap["dateAcknowledged"].(string); ok {
		DateAcknowledged, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateAcknowledgedString)
		o.DateAcknowledged = &DateAcknowledged
	}
	
	if SelfUri, ok := RecognitionMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Recognition) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
