package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Agentownedmappingpreview
type Agentownedmappingpreview struct { 
	// AgentOwnedColumn - The raw value of the agent-owned column
	AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`


	// Email - The email address of the user, if it exists
	Email *string `json:"email,omitempty"`


	// UserId - The id of the user, if it exists
	UserId *string `json:"userId,omitempty"`


	// Exists - Whether the user exists
	Exists *bool `json:"exists,omitempty"`


	// IsQueueMember - Whether the user is a member of the campaign's queue
	IsQueueMember *bool `json:"isQueueMember,omitempty"`


	// RecordCount - The number of contact records whose agent-owned column matches the raw value
	RecordCount *int `json:"recordCount,omitempty"`

}

func (o *Agentownedmappingpreview) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Agentownedmappingpreview
	
	return json.Marshal(&struct { 
		AgentOwnedColumn *string `json:"agentOwnedColumn,omitempty"`
		
		Email *string `json:"email,omitempty"`
		
		UserId *string `json:"userId,omitempty"`
		
		Exists *bool `json:"exists,omitempty"`
		
		IsQueueMember *bool `json:"isQueueMember,omitempty"`
		
		RecordCount *int `json:"recordCount,omitempty"`
		*Alias
	}{ 
		AgentOwnedColumn: o.AgentOwnedColumn,
		
		Email: o.Email,
		
		UserId: o.UserId,
		
		Exists: o.Exists,
		
		IsQueueMember: o.IsQueueMember,
		
		RecordCount: o.RecordCount,
		Alias:    (*Alias)(o),
	})
}

func (o *Agentownedmappingpreview) UnmarshalJSON(b []byte) error {
	var AgentownedmappingpreviewMap map[string]interface{}
	err := json.Unmarshal(b, &AgentownedmappingpreviewMap)
	if err != nil {
		return err
	}
	
	if AgentOwnedColumn, ok := AgentownedmappingpreviewMap["agentOwnedColumn"].(string); ok {
		o.AgentOwnedColumn = &AgentOwnedColumn
	}
    
	if Email, ok := AgentownedmappingpreviewMap["email"].(string); ok {
		o.Email = &Email
	}
    
	if UserId, ok := AgentownedmappingpreviewMap["userId"].(string); ok {
		o.UserId = &UserId
	}
    
	if Exists, ok := AgentownedmappingpreviewMap["exists"].(bool); ok {
		o.Exists = &Exists
	}
    
	if IsQueueMember, ok := AgentownedmappingpreviewMap["isQueueMember"].(bool); ok {
		o.IsQueueMember = &IsQueueMember
	}
    
	if RecordCount, ok := AgentownedmappingpreviewMap["recordCount"].(float64); ok {
		RecordCountInt := int(RecordCount)
		o.RecordCount = &RecordCountInt
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Agentownedmappingpreview) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
