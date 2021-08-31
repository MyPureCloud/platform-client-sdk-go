package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatconversation
type Webchatconversation struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Member - Chat Member
	Member *Webchatmemberinfo `json:"member,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Webchatconversation) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatconversation
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Member *Webchatmemberinfo `json:"member,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Member: o.Member,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatconversation) UnmarshalJSON(b []byte) error {
	var WebchatconversationMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatconversationMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatconversationMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := WebchatconversationMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if Member, ok := WebchatconversationMap["member"].(map[string]interface{}); ok {
		MemberString, _ := json.Marshal(Member)
		json.Unmarshal(MemberString, &o.Member)
	}
	
	if SelfUri, ok := WebchatconversationMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatconversation) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
