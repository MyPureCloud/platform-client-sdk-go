package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrorentity
type Bulkerrorentity struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// Retryable
	Retryable *bool `json:"retryable,omitempty"`


	// Entity
	Entity *Entity `json:"entity,omitempty"`


	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`

}

func (o *Bulkerrorentity) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkerrorentity
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		Retryable *bool `json:"retryable,omitempty"`
		
		Entity *Entity `json:"entity,omitempty"`
		
		Details *[]Bulkerrordetail `json:"details,omitempty"`
		*Alias
	}{ 
		Code: o.Code,
		
		Message: o.Message,
		
		Status: o.Status,
		
		Retryable: o.Retryable,
		
		Entity: o.Entity,
		
		Details: o.Details,
		Alias:    (*Alias)(o),
	})
}

func (o *Bulkerrorentity) UnmarshalJSON(b []byte) error {
	var BulkerrorentityMap map[string]interface{}
	err := json.Unmarshal(b, &BulkerrorentityMap)
	if err != nil {
		return err
	}
	
	if Code, ok := BulkerrorentityMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := BulkerrorentityMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Status, ok := BulkerrorentityMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Retryable, ok := BulkerrorentityMap["retryable"].(bool); ok {
		o.Retryable = &Retryable
	}
	
	if Entity, ok := BulkerrorentityMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Details, ok := BulkerrorentityMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkerrorentity) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
