package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Note
type Note struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Division - The division to which this entity belongs.
	Division *Writablestarrabledivision `json:"division,omitempty"`

	// EntityId - The id of the contact or organization to which this note refers. This only needs to be set for input when using the Bulk APIs.
	EntityId *string `json:"entityId,omitempty"`

	// EntityType - This is only need to be set when using Bulk API. Using any other value than contact or organization will result in null being used.
	EntityType *string `json:"entityType,omitempty"`

	// NoteText
	NoteText *string `json:"noteText,omitempty"`

	// ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`

	// CreatedBy - When creating or updating a note, only User.id is required. User object is fully populated when expanding a note.
	CreatedBy *User `json:"createdBy,omitempty"`

	// ExternalDataSources - Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Note) SetField(field string, fieldValue interface{}) {
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

func (o Note) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "ModifyDate","CreateDate", }
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
	type Alias Note
	
	ModifyDate := new(string)
	if o.ModifyDate != nil {
		
		*ModifyDate = timeutil.Strftime(o.ModifyDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		ModifyDate = nil
	}
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Division *Writablestarrabledivision `json:"division,omitempty"`
		
		EntityId *string `json:"entityId,omitempty"`
		
		EntityType *string `json:"entityType,omitempty"`
		
		NoteText *string `json:"noteText,omitempty"`
		
		ModifyDate *string `json:"modifyDate,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		CreatedBy *User `json:"createdBy,omitempty"`
		
		ExternalDataSources *[]Externaldatasource `json:"externalDataSources,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Division: o.Division,
		
		EntityId: o.EntityId,
		
		EntityType: o.EntityType,
		
		NoteText: o.NoteText,
		
		ModifyDate: ModifyDate,
		
		CreateDate: CreateDate,
		
		CreatedBy: o.CreatedBy,
		
		ExternalDataSources: o.ExternalDataSources,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Note) UnmarshalJSON(b []byte) error {
	var NoteMap map[string]interface{}
	err := json.Unmarshal(b, &NoteMap)
	if err != nil {
		return err
	}
	
	if Id, ok := NoteMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Division, ok := NoteMap["division"].(map[string]interface{}); ok {
		DivisionString, _ := json.Marshal(Division)
		json.Unmarshal(DivisionString, &o.Division)
	}
	
	if EntityId, ok := NoteMap["entityId"].(string); ok {
		o.EntityId = &EntityId
	}
    
	if EntityType, ok := NoteMap["entityType"].(string); ok {
		o.EntityType = &EntityType
	}
    
	if NoteText, ok := NoteMap["noteText"].(string); ok {
		o.NoteText = &NoteText
	}
    
	if modifyDateString, ok := NoteMap["modifyDate"].(string); ok {
		ModifyDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", modifyDateString)
		o.ModifyDate = &ModifyDate
	}
	
	if createDateString, ok := NoteMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if CreatedBy, ok := NoteMap["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ExternalDataSources, ok := NoteMap["externalDataSources"].([]interface{}); ok {
		ExternalDataSourcesString, _ := json.Marshal(ExternalDataSources)
		json.Unmarshal(ExternalDataSourcesString, &o.ExternalDataSources)
	}
	
	if SelfUri, ok := NoteMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Note) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
