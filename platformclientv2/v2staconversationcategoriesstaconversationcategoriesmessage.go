package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// V2staconversationcategoriesstaconversationcategoriesmessage
type V2staconversationcategoriesstaconversationcategoriesmessage struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// OrganizationId
	OrganizationId *string `json:"organizationId,omitempty"`

	// ConversationId
	ConversationId *string `json:"conversationId,omitempty"`

	// MediaType
	MediaType *string `json:"mediaType,omitempty"`

	// TranscriptIds
	TranscriptIds *[]string `json:"transcriptIds,omitempty"`

	// CategoryIds
	CategoryIds *[]string `json:"categoryIds,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *V2staconversationcategoriesstaconversationcategoriesmessage) SetField(field string, fieldValue interface{}) {
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

func (o V2staconversationcategoriesstaconversationcategoriesmessage) MarshalJSON() ([]byte, error) {
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
	type Alias V2staconversationcategoriesstaconversationcategoriesmessage
	
	return json.Marshal(&struct { 
		OrganizationId *string `json:"organizationId,omitempty"`
		
		ConversationId *string `json:"conversationId,omitempty"`
		
		MediaType *string `json:"mediaType,omitempty"`
		
		TranscriptIds *[]string `json:"transcriptIds,omitempty"`
		
		CategoryIds *[]string `json:"categoryIds,omitempty"`
		Alias
	}{ 
		OrganizationId: o.OrganizationId,
		
		ConversationId: o.ConversationId,
		
		MediaType: o.MediaType,
		
		TranscriptIds: o.TranscriptIds,
		
		CategoryIds: o.CategoryIds,
		Alias:    (Alias)(o),
	})
}

func (o *V2staconversationcategoriesstaconversationcategoriesmessage) UnmarshalJSON(b []byte) error {
	var V2staconversationcategoriesstaconversationcategoriesmessageMap map[string]interface{}
	err := json.Unmarshal(b, &V2staconversationcategoriesstaconversationcategoriesmessageMap)
	if err != nil {
		return err
	}
	
	if OrganizationId, ok := V2staconversationcategoriesstaconversationcategoriesmessageMap["organizationId"].(string); ok {
		o.OrganizationId = &OrganizationId
	}
    
	if ConversationId, ok := V2staconversationcategoriesstaconversationcategoriesmessageMap["conversationId"].(string); ok {
		o.ConversationId = &ConversationId
	}
    
	if MediaType, ok := V2staconversationcategoriesstaconversationcategoriesmessageMap["mediaType"].(string); ok {
		o.MediaType = &MediaType
	}
    
	if TranscriptIds, ok := V2staconversationcategoriesstaconversationcategoriesmessageMap["transcriptIds"].([]interface{}); ok {
		TranscriptIdsString, _ := json.Marshal(TranscriptIds)
		json.Unmarshal(TranscriptIdsString, &o.TranscriptIds)
	}
	
	if CategoryIds, ok := V2staconversationcategoriesstaconversationcategoriesmessageMap["categoryIds"].([]interface{}); ok {
		CategoryIdsString, _ := json.Marshal(CategoryIds)
		json.Unmarshal(CategoryIdsString, &o.CategoryIds)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *V2staconversationcategoriesstaconversationcategoriesmessage) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
