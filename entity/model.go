package entity

type BaseStreamMessage struct {
	EventType string `mapstructure:"event_type"`
	SentAt    string `mapstructure:"sent_at"`
}

type BaseItemType struct {
	NftId     string               `mapstructure:"nft_id"`
	Permalink string               `mapstructure:"permalink"`
	Metadata  BaseItemMetadataType `mapstructure:"metadata"`
	Chain     Chain                `mapstructure:"chain"`
}
type PayloadItemAndColl struct {
	Item       BaseItemType `mapstructure:"item"`
	Collection Collection   `mapstructure:"collection"`
	OrderHash  string       `mapstructure:"order_hash"`
}

type Collection struct {
	Slug string `mapstructure:"slug"`
}
type Chain struct {
	Name string `mapstructure:"name"`
}
type ItemListedEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemListedEventPayload `mapstructure:"payload"`
}

type ItemListedEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	Quantity           int          `mapstructure:"quantity"`
	ListingType        string       `mapstructure:"listing_type"`
	ListingDate        string       `mapstructure:"listing_date"`
	ExpirationDate     string       `mapstructure:"expiration_date"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	BasePrice          string       `mapstructure:"base_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	IsPrivate          bool         `mapstructure:"is_private"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}

type ItemReceivedOfferEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemReceivedOfferEventPayload `mapstructure:"payload"`
}
type ItemReceivedOfferEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	Quantity           int          `mapstructure:"quantity"`
	CreatedDate        string       `mapstructure:"created_date"`
	ExpirationDate     string       `mapstructure:"expiration_date"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	BasePrice          string       `mapstructure:"base_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}
type ItemReceivedBidEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemReceivedBidEventPayload `mapstructure:"payload"`
}
type ItemReceivedBidEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	Quantity           int          `mapstructure:"quantity"`
	CreatedDate        string       `mapstructure:"created_date"`
	ExpirationDate     string       `mapstructure:"expiration_date"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	BasePrice          string       `mapstructure:"base_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}
type ItemSoldEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemSoldEventPayload `mapstructure:"payload"`
}
type ItemSoldEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	ListingType        string       `mapstructure:"listing_type"`
	ClosingDate        string       `mapstructure:"closing_date"`
	Transaction        Transaction  `mapstructure:"transaction"`
	Maker              Account      `mapstructure:"maker"`
	Taker              Account      `mapstructure:"taker"`
	SalePrice          string       `mapstructure:"sale_price"`
	PaymentToken       PaymentToken `mapstructure:"payment_token"`
	IsPrivate          bool         `mapstructure:"is_private"`
	EventTimestamp     string       `mapstructure:"event_timestamp"`
}
type ItemTransferredEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemTransferredEventPayload `mapstructure:"payload"`
}
type ItemTransferredEventPayload struct {
	PayloadItemAndColl `mapstructure:",squash"`
	FromAccount        Account     `mapstructure:"from_account"`
	Quantity           int         `mapstructure:"quantity"`
	ToAccount          Account     `mapstructure:"to_account"`
	Transaction        Transaction `mapstructure:"transaction"`
	EventTimestamp     string      `mapstructure:"event_timestamp"`
}
type ItemCancelledEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemCancelledEventPayload `mapstructure:"payload"`
}
type ItemCancelledEventPayload struct {
	ItemListedEventPayload `mapstructure:",squash"`
	Transaction            Transaction `mapstructure:"transaction"`

	//PayloadItemAndColl `mapstructure:",squash"`
	//Quantity           int          `mapstructure:"quantity"`
	//ListingType        string       `mapstructure:"listing_type"`
	//PaymentToken       PaymentToken `mapstructure:"payment_token"`
	//EventTimestamp     string       `mapstructure:"event_timestamp"`
}
type ItemMetadataUpdateEvent struct {
	BaseStreamMessage `mapstructure:",squash"`
	Payload           ItemMetadataUpdatePayload `mapstructure:"payload"`
}
type ItemMetadataUpdatePayload struct {
	PayloadItemAndColl   `mapstructure:",squash"`
	BaseItemMetadataType `mapstructure:",squash"`
	Description          string  `mapstructure:"description"`
	BackgroundColor      string  `mapstructure:"background_color"`
	Traits               []Trait `mapstructure:"traits"`
}

type BaseItemMetadataType struct {
	Name         string `mapstructure:"name"`
	ImageUrl     string `mapstructure:"image_url"`
	AnimationUrl string `mapstructure:"animation_url"`
	MetadataUrl  string `mapstructure:"metadata_url"`
}

type Transaction struct {
	Hash      string `mapstructure:"hash"`
	Timestamp string `mapstructure:"timestamp"`
}

type Account struct {
	Address string `mapstructure:"address"`
}

type Trait struct {
	TraitType   string `mapstructure:"trait_type"`
	Value       string `mapstructure:"value"`
	DisplayType string `mapstructure:"display_type"`
	MaxValue    int    `mapstructure:"max_value"`
	TraitCount  string `mapstructure:"trait_count"`
	Order       int    `mapstructure:"order"`
}

type PaymentToken struct {
	Address  string `mapstructure:"address"`
	Decimals int    `mapstructure:"decimals"`
	EthPrice string `mapstructure:"eth_price"`
	Name     string `mapstructure:"name"`
	Symbol   string `mapstructure:"symbol"`
	UsdPrice string `mapstructure:"usd_price"`
}
