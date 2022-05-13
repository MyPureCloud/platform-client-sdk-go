package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Webchatguestmediarequest - Object representing the guest model of a media request of a chat conversation.
type Webchatguestmediarequest struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// Types - The types of media being requested.
	Types *[]string `json:"types,omitempty"`


	// State - The state of the media request, one of PENDING|ACCEPTED|DECLINED|TIMEDOUT|CANCELLED|ERRORED.
	State *string `json:"state,omitempty"`


	// CommunicationId - The ID of the new media communication, if applicable.
	CommunicationId *string `json:"communicationId,omitempty"`


	// SecurityKey - The security information related to a media request.
	SecurityKey *string `json:"securityKey,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Webchatguestmediarequest) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Webchatguestmediarequest
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		Types *[]string `json:"types,omitempty"`
		
		State *string `json:"state,omitempty"`
		
		CommunicationId *string `json:"communicationId,omitempty"`
		
		SecurityKey *string `json:"securityKey,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		Types: o.Types,
		
		State: o.State,
		
		CommunicationId: o.CommunicationId,
		
		SecurityKey: o.SecurityKey,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Webchatguestmediarequest) UnmarshalJSON(b []byte) error {
	var WebchatguestmediarequestMap map[string]interface{}
	err := json.Unmarshal(b, &WebchatguestmediarequestMap)
	if err != nil {
		return err
	}
	
	if Id, ok := WebchatguestmediarequestMap["id"].(string); ok {
		o.Id = &Id
	}
    
	if Name, ok := WebchatguestmediarequestMap["name"].(string); ok {
		o.Name = &Name
	}
    
	if Types, ok := WebchatguestmediarequestMap["types"].([]interface{}); ok {
		TypesString, _ := json.Marshal(Types)
		json.Unmarshal(TypesString, &o.Types)
	}
	
	if State, ok := WebchatguestmediarequestMap["state"].(string); ok {
		o.State = &State
	}
    
	if CommunicationId, ok := WebchatguestmediarequestMap["communicationId"].(string); ok {
		o.CommunicationId = &CommunicationId
	}
    
	if SecurityKey, ok := WebchatguestmediarequestMap["securityKey"].(string); ok {
		o.SecurityKey = &SecurityKey
	}
    
	if SelfUri, ok := WebchatguestmediarequestMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
    

	return nil
}

// String returns a JSON representation of the model
func (o *Webchatguestmediarequest) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
