package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Commandstatus
type Commandstatus struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Expiration - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	Expiration *time.Time `json:"expiration,omitempty"`

	// UserId
	UserId *string `json:"userId,omitempty"`

	// StatusCode
	StatusCode *string `json:"statusCode,omitempty"`

	// CommandType
	CommandType *string `json:"commandType,omitempty"`

	// Document
	Document *Document `json:"document,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Commandstatus) SetField(field string, fieldValue interface{}) {
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

func (o Commandstatus) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "Expiration", }
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
	type Alias Commandstatus
	
	Expiration := new(string)
	if o.Expiration != nil {
		
		*Expiration = timeutil.Strftime(o.Expiration, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		Expiration = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Expiration *string `json:"expiration,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		StatusCode *string `json:"statusCode,omitempty"`
		
		CommandType *string `json:"commandType,omitempty"`
		
		Document *Document `json:"document,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Expiration: Expiration,
		
		UserId: o.UserId,
		
		StatusCode: o.StatusCode,
		
		CommandType: o.CommandType,
		
		Document: o.Document,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Commandstatus) UnmarshalJSON(b []byte) error {
	var CommandstatusMap map[string]interface{}
	err := json.Unmarshal(b, &CommandstatusMap)
	if err != nil {
		return err
	}
	
	if Id, ok := CommandstatusMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := CommandstatusMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if expirationString, ok := CommandstatusMap["expiration"].(string); ok {
		Expiration, _ := time.Parse("2006-01-02T15:04:05.999999Z", expirationString)
		o.Expiration = &Expiration
	}
	
	if UserId, ok := CommandstatusMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if StatusCode, ok := CommandstatusMap["statusCode"].(string); ok {
		o.StatusCode = &StatusCode
	}
    
	if CommandType, ok := CommandstatusMap["commandType"].(string); ok {
		o.CommandType = &CommandType
	}
    
	if Document, ok := CommandstatusMap["document"].(map[string]interface{}); ok {
		DocumentString, _ := json.Marshal(Document)
		json.Unmarshal(DocumentString, &o.Document)
	}
	
	if SelfUri, ok := CommandstatusMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Commandstatus) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
