package stores


type store interface {
	Open(opts string) (err error)
	Close() error

	/* TK: We'll start with a minimal interface surface that uses reflection
	 		to dispatch CRUD methods. */
	Create(interface{})(result interface{}, err error)
	View(interface{})(result interface{}, err error)
	Update(interface{})(result interface{}, err error)
	Delete(interface{})(result interface{}, err error)
	List(interface{})(result interface{}, err error)

	/*
	MarketplaceCreate()(err error)
	MarketplaceList()(err error)
	MarketplaceView()(err error)
	MarketplaceUpdate()(err error)
	MarketplaceDelete()(err error)

	StorefrontCreate()(err error)
	StorefrontList()(err error)
	StorefrontView()(err error)
	StorefrontUpdate()(err error)
	StorefrontDelete()(err error)

	CatalogCreate()(err error)
	CatalogList()(err error)
	CatalogView()(err error)
	CatalogUpdate()(err error)
	CatalogDelete()(err error)

	ProductCreate()(err error)
	ProductList()(err error)
	ProductView()(err error)
	ProductUpdate()(err error)
	ProductDelete()(err error)

	SkuCreate()(err error)
	SkuList()(err error)
	SkuView()(err error)
	SkuUpdate()(err error)
	SkuDelete()(err error)
	*/
	
}