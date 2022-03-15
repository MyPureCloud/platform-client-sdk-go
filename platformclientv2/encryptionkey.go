package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Encryptionkey
type Encryptionkey struct { 
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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Encryptionkey) MarshalJSON() ([]byte, error) {
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		CreateDate: CreateDate,
		
		KeydataSummary: o.KeydataSummary,
		
		User: o.User,
		
		LocalEncryptionConfiguration: o.LocalEncryptionConfiguration,
		
		KeyConfigurationType: o.KeyConfigurationType,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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
