package platformclientv2
import (
	"github.com/leekchan/timeutil"
	"encoding/json"
	"strconv"
	"strings"
)

// Messagingsticker
type Messagingsticker struct { 
	// Id - The globally unique identifier for the object.
	Id *string `json:"id,omitempty"`


	// Name
	Name *string `json:"name,omitempty"`


	// ProviderStickerId - The sticker Id of the sticker, assigned by the sticker provider.
	ProviderStickerId *int `json:"providerStickerId,omitempty"`


	// ProviderPackageId - The package Id of the sticker, assigned by the sticker provider.
	ProviderPackageId *int `json:"providerPackageId,omitempty"`


	// PackageName - The package name of the sticker, assigned by the sticker provider.
	PackageName *string `json:"packageName,omitempty"`


	// MessengerType - The type of the messenger provider.
	MessengerType *string `json:"messengerType,omitempty"`


	// StickerType - The type of the sticker.
	StickerType *string `json:"stickerType,omitempty"`


	// ProviderVersion - The version of the sticker, assigned by the provider.
	ProviderVersion *int `json:"providerVersion,omitempty"`


	// UriLocation
	UriLocation *string `json:"uriLocation,omitempty"`


	// SelfUri - The URI for this object
	SelfUri *string `json:"selfUri,omitempty"`

}

func (o *Messagingsticker) MarshalJSON() ([]byte, error) {
	// Redundant initialization to avoid unused import errors for models with no Time values
	_  = timeutil.Timedelta{}
	type Alias Messagingsticker
	
	return json.Marshal(&struct { 
		Id *string `json:"id,omitempty"`
		
		Name *string `json:"name,omitempty"`
		
		ProviderStickerId *int `json:"providerStickerId,omitempty"`
		
		ProviderPackageId *int `json:"providerPackageId,omitempty"`
		
		PackageName *string `json:"packageName,omitempty"`
		
		MessengerType *string `json:"messengerType,omitempty"`
		
		StickerType *string `json:"stickerType,omitempty"`
		
		ProviderVersion *int `json:"providerVersion,omitempty"`
		
		UriLocation *string `json:"uriLocation,omitempty"`
		
		SelfUri *string `json:"selfUri,omitempty"`
		*Alias
	}{ 
		Id: o.Id,
		
		Name: o.Name,
		
		ProviderStickerId: o.ProviderStickerId,
		
		ProviderPackageId: o.ProviderPackageId,
		
		PackageName: o.PackageName,
		
		MessengerType: o.MessengerType,
		
		StickerType: o.StickerType,
		
		ProviderVersion: o.ProviderVersion,
		
		UriLocation: o.UriLocation,
		
		SelfUri: o.SelfUri,
		Alias:    (*Alias)(o),
	})
}

func (o *Messagingsticker) UnmarshalJSON(b []byte) error {
	var MessagingstickerMap map[string]interface{}
	err := json.Unmarshal(b, &MessagingstickerMap)
	if err != nil {
		return err
	}
	
	if Id, ok := MessagingstickerMap["id"].(string); ok {
		o.Id = &Id
	}
	
	if Name, ok := MessagingstickerMap["name"].(string); ok {
		o.Name = &Name
	}
	
	if ProviderStickerId, ok := MessagingstickerMap["providerStickerId"].(float64); ok {
		ProviderStickerIdInt := int(ProviderStickerId)
		o.ProviderStickerId = &ProviderStickerIdInt
	}
	
	if ProviderPackageId, ok := MessagingstickerMap["providerPackageId"].(float64); ok {
		ProviderPackageIdInt := int(ProviderPackageId)
		o.ProviderPackageId = &ProviderPackageIdInt
	}
	
	if PackageName, ok := MessagingstickerMap["packageName"].(string); ok {
		o.PackageName = &PackageName
	}
	
	if MessengerType, ok := MessagingstickerMap["messengerType"].(string); ok {
		o.MessengerType = &MessengerType
	}
	
	if StickerType, ok := MessagingstickerMap["stickerType"].(string); ok {
		o.StickerType = &StickerType
	}
	
	if ProviderVersion, ok := MessagingstickerMap["providerVersion"].(float64); ok {
		ProviderVersionInt := int(ProviderVersion)
		o.ProviderVersion = &ProviderVersionInt
	}
	
	if UriLocation, ok := MessagingstickerMap["uriLocation"].(string); ok {
		o.UriLocation = &UriLocation
	}
	
	if SelfUri, ok := MessagingstickerMap["selfUri"].(string); ok {
		o.SelfUri = &SelfUri
	}
	

	return nil
}

// String returns a JSON representation of the model
func (o *Messagingsticker) String() string {
	j, _ := json.Marshal(o)
	str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

	return str
}
