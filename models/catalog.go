package models

type Catalog struct {
	ID          string   `json:"id"`
	StoreID     string   `json:"storeid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageurls"`
}

type Category struct {
	ID          string   `json:"id"`
	CatalogID   string   `json:"catalogid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageurls"`
}

type ProductProperty struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	OptionType  string   `json:"type"`
	Values      []string `json:"values"`
	Min         float32  `json:"min"`
	Max         float32  `json:"max"`
	Increment   float32  `json:"incr"`
}

type ProductPropertyValuesList struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Min    float32  `json:"min"`
	Max    float32  `json:"max"`
	Values []string `json:"values"`
}

type ProductPropertyConfigurationRule struct {
	ID       string                      `json:"id"`
	SrcName  string                      `json:"srcname"`
	SrcValue string                      `json:"srcvalue"`
	Targets  []ProductPropertyValuesList `json:"tgts"`
}

type ProductPropertyValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProductOptions struct {
	AllowDynamicSkus bool `json:"allowDynamicSkus"`
}

type Product struct {
	ID          string                             `json:"id"`
	CatalogID   string                             `json:"catalogid"`
	Name        string                             `json:"name"`
	Description string                             `json:"description"`
	Properties  []ProductProperty                  `json:"properties"`
	Rules       []ProductPropertyConfigurationRule `json:"rules"`
	Options     ProductOptions                     `json:"options"`
}

type Sku struct {
	ID          string                 `json:"id"`
	ProductID   string                 `json:"productid"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Config      []ProductPropertyValue `json:"config"`
}
