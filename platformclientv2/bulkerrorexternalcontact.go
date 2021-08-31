package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Bulkerrorexternalcontact
type Bulkerrorexternalcontact struct { 
	// Code
	Code *string `json:"code,omitempty"`


	// Message
	Message *string `json:"message,omitempty"`


	// Status
	Status *int `json:"status,omitempty"`


	// Retryable
	Retryable *bool `json:"retryable,omitempty"`


	// Entity
	Entity *Externalcontact `json:"entity,omitempty"`


	// Details
	Details *[]Bulkerrordetail `json:"details,omitempty"`

}

func (o *Bulkerrorexternalcontact) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Bulkerrorexternalcontact
	
	return json.Marshal(&struct { 
		Code *string `json:"code,omitempty"`
		
		Message *string `json:"message,omitempty"`
		
		Status *int `json:"status,omitempty"`
		
		Retryable *bool `json:"retryable,omitempty"`
		
		Entity *Externalcontact `json:"entity,omitempty"`
		
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

func (o *Bulkerrorexternalcontact) UnmarshalJSON(b []byte) error {
	var BulkerrorexternalcontactMap map[string]interface{}
	err := json.Unmarshal(b, &BulkerrorexternalcontactMap)
	if err != nil {
		return err
	}
	
	if Code, ok := BulkerrorexternalcontactMap["code"].(string); ok {
		o.Code = &Code
	}
	
	if Message, ok := BulkerrorexternalcontactMap["message"].(string); ok {
		o.Message = &Message
	}
	
	if Status, ok := BulkerrorexternalcontactMap["status"].(float64); ok {
		StatusInt := int(Status)
		o.Status = &StatusInt
	}
	
	if Retryable, ok := BulkerrorexternalcontactMap["retryable"].(bool); ok {
		o.Retryable = &Retryable
	}
	
	if Entity, ok := BulkerrorexternalcontactMap["entity"].(map[string]interface{}); ok {
		EntityString, _ := json.Marshal(Entity)
		json.Unmarshal(EntityString, &o.Entity)
	}
	
	if Details, ok := BulkerrorexternalcontactMap["details"].([]interface{}); ok {
		DetailsString, _ := json.Marshal(Details)
		json.Unmarshal(DetailsString, &o.Details)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Bulkerrorexternalcontact) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
