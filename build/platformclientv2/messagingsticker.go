package platformclientv2
import (
	"encoding/json"
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

// String returns a JSON representation of the model
func (o *Messagingsticker) String() string {
	j, _ := json.Marshal(o)
	return string(j)
}
