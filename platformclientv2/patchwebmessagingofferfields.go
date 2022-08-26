package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Patchwebmessagingofferfields
type Patchwebmessagingofferfields struct { 
	// OfferText - Text value to be used when inviting a visitor to engage with a web messaging offer.
	OfferText *string `json:"offerText,omitempty"`


	// ArchitectFlow - Flow to be invoked, overrides default flow when specified.
	ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`

}

func (o *Patchwebmessagingofferfields) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Patchwebmessagingofferfields
	
	return json.Marshal(&struct { 
		OfferText *string `json:"offerText,omitempty"`
		
		ArchitectFlow *Addressableentityref `json:"architectFlow,omitempty"`
		*Alias
	}{ 
		OfferText: o.OfferText,
		
		ArchitectFlow: o.ArchitectFlow,
		Alias:    (*Alias)(o),
	})
}

func (o *Patchwebmessagingofferfields) UnmarshalJSON(b []byte) error {
	var PatchwebmessagingofferfieldsMap map[string]interface{}
	err := json.Unmarshal(b, &PatchwebmessagingofferfieldsMap)
	if err != nil {
		return err
	}
	
	if OfferText, ok := PatchwebmessagingofferfieldsMap["offerText"].(string); ok {
		o.OfferText = &OfferText
	}
    
	if ArchitectFlow, ok := PatchwebmessagingofferfieldsMap["architectFlow"].(map[string]interface{}); ok {
		ArchitectFlowString, _ := json.Marshal(ArchitectFlow)
		json.Unmarshal(ArchitectFlowString, &o.ArchitectFlow)
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Patchwebmessagingofferfields) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
