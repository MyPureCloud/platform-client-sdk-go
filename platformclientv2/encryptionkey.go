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


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (u *Encryptionkey) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Encryptionkey

	
	CreateDate := new(string)
	if u.CreateDate != nil {
		
		*CreateDate = timeutil.Strftime(u.CreateDate, "%Y-%m-%dT%H:%M:%S.%fZ")
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
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: u.Id,
		
		Name: u.Name,
		
		CreateDate: CreateDate,
		
		KeydataSummary: u.KeydataSummary,
		
		User: u.User,
		
		LocalEncryptionConfiguration: u.LocalEncryptionConfiguration,
		
		SelfUri: u.SelfUri,
		Alias:    (*Alias)(u),
	})
}

// String returns a JSON representation of the model
func (o *Encryptionkey) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
