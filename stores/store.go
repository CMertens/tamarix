package stores


type store interface {
	Open(opts string) (err error)
	Close() error

	MarketplaceCreate()(err error)
	MarketplaceGet()(err error)
	MarketplaceView()(err error)
	MarketplaceUpdate()(err error)
	MarketplaceDelete()(err error)

	StorefrontCreate()(err error)
	StorefrontGet()(err error)
	StorefrontView()(err error)
	StorefrontUpdate()(err error)
	StorefrontDelete()(err error)

	CatalogCreate()(err error)
	CatalogGet()(err error)
	CatalogView()(err error)
	CatalogUpdate()(err error)
	CatalogDelete()(err error)

	ProductCreate()(err error)
	ProductGet()(err error)
	ProductView()(err error)
	ProductUpdate()(err error)
	ProductDelete()(err error)

	SkuCreate()(err error)
	SkuGet()(err error)
	SkuView()(err error)
	SkuUpdate()(err error)
	SkuDelete()(err error)

	
}