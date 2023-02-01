package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Adhocrecordingtopicrecordingdatav2
type Adhocrecordingtopicrecordingdatav2 struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// DateCreated
	DateCreated *time.Time `json:"dateCreated,omitempty"`

	// DateModified
	DateModified *time.Time `json:"dateModified,omitempty"`

	// Workspace
	Workspace *Adhocrecordingtopicworkspacedata `json:"workspace,omitempty"`

	// CreatedBy
	CreatedBy *Adhocrecordingtopicuserdata `json:"createdBy,omitempty"`

	// ContentType
	ContentType *string `json:"contentType,omitempty"`

	// ContentLength
	ContentLength *int `json:"contentLength,omitempty"`

	// Filename
	Filename *string `json:"filename,omitempty"`

	// ChangeNumber
	ChangeNumber *int `json:"changeNumber,omitempty"`

	// DateUploaded
	DateUploaded *time.Time `json:"dateUploaded,omitempty"`

	// UploadedBy
	UploadedBy *Adhocrecordingtopicuserdata `json:"uploadedBy,omitempty"`

	// LockInfo
	LockInfo *Adhocrecordingtopiclockdata `json:"lockInfo,omitempty"`

	// SelfUri
	SelfUri *string `json:"selfUri,omitempty"`

	// DurationMillieconds
	DurationMillieconds *int `json:"durationMillieconds,omitempty"`

	// Conversation
	Conversation *Adhocrecordingtopicconversationdata `json:"conversation,omitempty"`

	// Read
	Read *bool `json:"read,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Adhocrecordingtopicrecordingdatav2) SetField(field string, fieldValue interface{}) {
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

func (o Adhocrecordingtopicrecordingdatav2) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "DateCreated","DateModified","DateUploaded", }
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
	type Alias Adhocrecordingtopicrecordingdatav2
	
	DateCreated := new(string)
	if o.DateCreated != nil {
		
		*DateCreated = timeutil.Strftime(o.DateCreated, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateCreated = nil
	}
	
	DateModified := new(string)
	if o.DateModified != nil {
		
		*DateModified = timeutil.Strftime(o.DateModified, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateModified = nil
	}
	
	DateUploaded := new(string)
	if o.DateUploaded != nil {
		
		*DateUploaded = timeutil.Strftime(o.DateUploaded, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		DateUploaded = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		DateCreated *string `json:"dateCreated,omitempty"`
		
		DateModified *string `json:"dateModified,omitempty"`
		
		Workspace *Adhocrecordingtopicworkspacedata `json:"workspace,omitempty"`
		
		CreatedBy *Adhocrecordingtopicuserdata `json:"createdBy,omitempty"`
		
		ContentType *string `json:"contentType,omitempty"`
		
		ContentLength *int `json:"contentLength,omitempty"`
		
		Filename *string `json:"filename,omitempty"`
		
		ChangeNumber *int `json:"changeNumber,omitempty"`
		
		DateUploaded *string `json:"dateUploaded,omitempty"`
		
		UploadedBy *Adhocrecordingtopicuserdata `json:"uploadedBy,omitempty"`
		
		LockInfo *Adhocrecordingtopiclockdata `json:"lockInfo,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		
		DurationMillieconds *int `json:"durationMillieconds,omitempty"`
		
		Conversation *Adhocrecordingtopicconversationdata `json:"conversation,omitempty"`
		
		Read *bool `json:"read,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		DateCreated: DateCreated,
		
		DateModified: DateModified,
		
		Workspace: o.Workspace,
		
		CreatedBy: o.CreatedBy,
		
		ContentType: o.ContentType,
		
		ContentLength: o.ContentLength,
		
		Filename: o.Filename,
		
		ChangeNumber: o.ChangeNumber,
		
		DateUploaded: DateUploaded,
		
		UploadedBy: o.UploadedBy,
		
		LockInfo: o.LockInfo,
		
		SelfUri: o.SelfUri,
		
		DurationMillieconds: o.DurationMillieconds,
		
		Conversation: o.Conversation,
		
		Read: o.Read,
		Alias:    (Alias)(o),
	})
}

func (o *Adhocrecordingtopicrecordingdatav2) UnmarshalJSON(b []byte) error {
	var Adhocrecordingtopicrecordingdatav2Map map[string]interface{}
	err := json.Unmarshal(b, &Adhocrecordingtopicrecordingdatav2Map)
	if err != nil {
		return err
	}
	
	if Id, ok := Adhocrecordingtopicrecordingdatav2Map["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := Adhocrecordingtopicrecordingdatav2Map["name"].(string); ok {
		o.Name = &Name
	}
    
	if dateCreatedString, ok := Adhocrecordingtopicrecordingdatav2Map["dateCreated"].(string); ok {
		DateCreated, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateCreatedString)
		o.DateCreated = &DateCreated
	}
	
	if dateModifiedString, ok := Adhocrecordingtopicrecordingdatav2Map["dateModified"].(string); ok {
		DateModified, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateModifiedString)
		o.DateModified = &DateModified
	}
	
	if Workspace, ok := Adhocrecordingtopicrecordingdatav2Map["workspace"].(map[string]interface{}); ok {
		WorkspaceString, _ := json.Marshal(Workspace)
		json.Unmarshal(WorkspaceString, &o.Workspace)
	}
	
	if CreatedBy, ok := Adhocrecordingtopicrecordingdatav2Map["createdBy"].(map[string]interface{}); ok {
		CreatedByString, _ := json.Marshal(CreatedBy)
		json.Unmarshal(CreatedByString, &o.CreatedBy)
	}
	
	if ContentType, ok := Adhocrecordingtopicrecordingdatav2Map["contentType"].(string); ok {
		o.ContentType = &ContentType
	}
    
	if ContentLength, ok := Adhocrecordingtopicrecordingdatav2Map["contentLength"].(float64); ok {
		ContentLengthInt := int(ContentLength)
		o.ContentLength = &ContentLengthInt
	}
	
	if Filename, ok := Adhocrecordingtopicrecordingdatav2Map["filename"].(string); ok {
		o.Filename = &Filename
	}
    
	if ChangeNumber, ok := Adhocrecordingtopicrecordingdatav2Map["changeNumber"].(float64); ok {
		ChangeNumberInt := int(ChangeNumber)
		o.ChangeNumber = &ChangeNumberInt
	}
	
	if dateUploadedString, ok := Adhocrecordingtopicrecordingdatav2Map["dateUploaded"].(string); ok {
		DateUploaded, _ := time.Parse("2006-01-02T15:04:05.999999Z", dateUploadedString)
		o.DateUploaded = &DateUploaded
	}
	
	if UploadedBy, ok := Adhocrecordingtopicrecordingdatav2Map["uploadedBy"].(map[string]interface{}); ok {
		UploadedByString, _ := json.Marshal(UploadedBy)
		json.Unmarshal(UploadedByString, &o.UploadedBy)
	}
	
	if LockInfo, ok := Adhocrecordingtopicrecordingdatav2Map["lockInfo"].(map[string]interface{}); ok {
		LockInfoString, _ := json.Marshal(LockInfo)
		json.Unmarshal(LockInfoString, &o.LockInfo)
	}
	
	if SelfUri, ok := Adhocrecordingtopicrecordingdatav2Map["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    
	if DurationMillieconds, ok := Adhocrecordingtopicrecordingdatav2Map["durationMillieconds"].(float64); ok {
		DurationMilliecondsInt := int(DurationMillieconds)
		o.DurationMillieconds = &DurationMilliecondsInt
	}
	
	if Conversation, ok := Adhocrecordingtopicrecordingdatav2Map["conversation"].(map[string]interface{}); ok {
		ConversationString, _ := json.Marshal(Conversation)
		json.Unmarshal(ConversationString, &o.Conversation)
	}
	
	if Read, ok := Adhocrecordingtopicrecordingdatav2Map["read"].(bool); ok {
		o.Read = &Read
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Adhocrecordingtopicrecordingdatav2) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
