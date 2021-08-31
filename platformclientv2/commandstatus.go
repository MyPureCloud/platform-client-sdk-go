package platformclientv2
import (
	"time"
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Commandstatus
type Commandstatus struct { 
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

func (o *Commandstatus) MarshalJSON() ([]byte, error) {
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
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Expiration: Expiration,
		
		UserId: o.UserId,
		
		StatusCode: o.StatusCode,
		
		CommandType: o.CommandType,
		
		Document: o.Document,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
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
