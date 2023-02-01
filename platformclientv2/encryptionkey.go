package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"reflect"
	"encoding/json"
	"strconv"
	"strings"
)

// Encryptionkey
type Encryptionkey struct { 
	// SetFieldNames defines the list of fields to use for controlled JSON serialization
	SetFieldNames map[string]bool `json:"-"`
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// CreateDate - create date of the key pair. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	CreateDate *time.Time `json:"createDate,omitempty"`

	// KeydataSummary - key data summary (base 64 encoded public key)
	KeydataSummary *string `json:"keydataSummary,omitempty"`

	// User - user that requested generation of public key
	User *User `json:"user,omitempty"`

	// LocalEncryptionConfiguration - Local configuration
	LocalEncryptionConfiguration *Localencryptionconfiguration `json:"localEncryptionConfiguration,omitempty"`

	// KeyConfigurationType - Key type used in this configuration
	KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`

	// KmsKeyArn - ARN of internal key to be wrapped by AWS KMS Symmetric key
	KmsKeyArn *string `json:"kmsKeyArn,omitempty"`

	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`
}

// SetField uses reflection to set a field on the model if the model has a property SetFieldNames, and triggers custom JSON serialization logic to only serialize properties that have been set using this function.
func (o *Encryptionkey) SetField(field string, fieldValue interface{}) {
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

func (o Encryptionkey) MarshalJSON() ([]byte, error) {
	// Special processing to dynamically construct object using only field names that have been set using SetField. This generates payloads suitable for use with PATCH API endpoints.
	if len(o.SetFieldNames) > 0 {
		// Get reflection Value
		val := reflect.ValueOf(o)

		// Known field names that require type overrides
		dateTimeFields := []string{ "CreateDate", }
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
	type Alias Encryptionkey
	
	CreateDate := new(string)
	if o.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(o.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
	} else {
		CreateDate = nil
	}
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		CreateDate *string `json:"createDate,omitempty"`
		
		KeydataSummary *string `json:"keydataSummary,omitempty"`
		
		User *User `json:"user,omitempty"`
		
		LocalEncryptionConfiguration *Localencryptionconfiguration `json:"localEncryptionConfiguration,omitempty"`
		
		KeyConfigurationType *string `json:"keyConfigurationType,omitempty"`
		
		KmsKeyArn *string `json:"kmsKeyArn,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CreateDate: CreateDate,
		
		KeydataSummary: o.KeydataSummary,
		
		User: o.User,
		
		LocalEncryptionConfiguration: o.LocalEncryptionConfiguration,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		KmsKeyArn: o.KmsKeyArn,
		
		SelfUri: o.SelfUri,
		Alias:    (Alias)(o),
	})
}

func (o *Encryptionkey) UnmarshalJSON(b []byte) error {
	var EncryptionkeyMap map[string]interface{}
	err := json.Unmarshal(b, &EncryptionkeyMap)
	if err != nil {
		return err
	}
	
	if Id, ok := EncryptionkeyMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := EncryptionkeyMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if createDateString, ok := EncryptionkeyMap["createDate"].(string); ok {
		CreateDate, _ := time.Parse("2006-01-02T15:04:05.999999Z", createDateString)
		o.CreateDate = &CreateDate
	}
	
	if KeydataSummary, ok := EncryptionkeyMap["keydataSummary"].(string); ok {
		o.KeydataSummary = &KeydataSummary
	}
    
	if User, ok := EncryptionkeyMap["user"].(map[string]interface{}); ok {
		UserString, _ := json.Marshal(User)
		json.Unmarshal(UserString, &o.User)
	}
	
	if LocalEncryptionConfiguration, ok := EncryptionkeyMap["localEncryptionConfiguration"].(map[string]interface{}); ok {
		LocalEncryptionConfigurationString, _ := json.Marshal(LocalEncryptionConfiguration)
		json.Unmarshal(LocalEncryptionConfigurationString, &o.LocalEncryptionConfiguration)
	}
	
	if KeyConfigurationType, ok := EncryptionkeyMap["keyConfigurationType"].(string); ok {
		o.KeyConfigurationType = &KeyConfigurationType
	}
    
	if KmsKeyArn, ok := EncryptionkeyMap["kmsKeyArn"].(string); ok {
		o.KmsKeyArn = &KmsKeyArn
	}
    
	if SelfUri, ok := EncryptionkeyMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Encryptionkey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
