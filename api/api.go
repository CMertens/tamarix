package api

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/guregu/kami"
	_ "golang.org/x/net/context"
)

var (
	defaultVersion = "v1.0"
)

type API struct {
	handler http.Handler
	client  http.Client
	logger  logrus.Entry
	version string
	// config
}

func NewAPI() *API {
	api := &API{}

	mux := kami.New()

	mux.Get("/", api.Index)

	/* MARKETPLACE ADMIN ENDPOINTS */

	/* STOREFRONT ADMIN ENDPOINTS */

	/* USER API ENDPOINTS */
	mux.Get("/marketplaces", api.MarketplacesList)
	mux.Get("/marketplace/:id", api.MarketplaceView)

	mux.Get("/storefronts", api.StorefrontsList)
	mux.Get("/storefront/:id", api.StorefrontView)

	mux.Get("/storefront/:storeid/catalogs", api.CatalogsList)
	mux.Get("/catalog/:id", api.CatalogView)

	mux.Get("/catalog/:catid/categories", api.CategoriesList)
	mux.Get("/category/:id", api.CategoryView)

	mux.Get("/product/:id", api.ProductView)
	mux.Post("/product/:id/configure", api.GetSkuFromProductConfiguration)
	mux.Get("/sku/:id", api.SkuView)
	mux.Get("/sku/:id/price", api.GetSkuPrice)

	/* User Account Management */
	mux.Get("/account/:id", api.AccountView)

	mux.Get("/account/:id/addresses", api.AddressesListFromAccount)
	mux.Post("/account/:id/address", api.AddressCreate)
	mux.Get("/address/:id", api.AddressView)

	mux.Get("/account/:id/paymethods", api.PaymethodsListFromAccount)
	mux.Post("/account/:id/paymethod", api.PaymethodCreate)

	mux.Get("/account/:id/credits", api.AccountCreditView)

	mux.Get("/account/:id/entitlements", api.AccountEntitlementsList)
	mux.Get("/entitlement/:id", api.EntitlementView)
	mux.Put("/entitlement/:id", api.EntitlementUpdate)

	mux.Post("/register", api.AccountCreate)
	mux.Post("/login", api.Login)
	mux.Post("/logout", api.Logout)
	mux.Post("/reset/request", api.ResetPassword)
	mux.Post("/reset/validate", api.ValidateReset)
	mux.Post("/reset/changepwd", api.ChangePassword)

	/* User Purchase Orders */
	mux.Get("/account/:id/orders", api.OrdersListFromAccount)
	mux.Post("/account/:id/order", api.OrderCreateFromAccount)
	mux.post("/order", api.OrderCreate)
	mux.Put("/order/:id", api.OrderUpdate)
	mux.Get("/order/:id", api.OrderView)

	mux.Get("/order/:id/orderitems", api.OrderItemsListFromOrder)
	mux.Post("/order/:id/orderitem", api.OrderItemCreateFromOrder)
	mux.Get("/orderitem/:id", api.OrderItemView)
	mux.Put("/orderitem/:id", api.OrderItemUpdate)

	/* User Invoices */
	mux.Get("/account/:id/invoices", api.InvoicesListFromAccount)
	mux.Get("/order/:id/invoices", api.InvoicesListFromOrder)
	mux.Post("/order/:id/invoice", api.InvoiceCreateFromOrder)
	mux.Put("/invoice/:id", api.InvoiceUpdate)

	mux.Get("/invoice/:id/invoiceitems", api.InvoiceItemListFromInvoice)
	mux.Get("/invoiceitem/:id", api.InvoiceItemView)

	/* Coupon */
	mux.Get("/account/:id/coupons", api.CouponsListFromAccount)
	mux.Post("/account/:id/coupon", api.CouponCreateFromAccount)
	mux.Get("/coupon/:id", api.CouponView)
	mux.Put("/coupon/:id", api.CouponUpdate)

	/* User RMAs */
	mux.Get("/account/:id/rmas", api.RMAsListFromAccount)
	mux.Get("/invoice/:id/rmas", api.RMAsListFromInvoice)
	mux.Post("/invoice/:id/rma", api.RMACreateFromInvoice)
	mux.Get("/rma/:id", api.RMAView)
	mux.Put("/rma/:id", api.RMAUpdate)

	/* User Payments */
	mux.Get("/account/:id/payments", api.PaymentsListFromAccount)
	mux.Get("/invoice/:id/payments", api.PaymentsListFromInvoice)
	mux.Get("/payment/:id", api.PaymentView)
	mux.Post("/invoice/:id/payment", api.PaymentCreateFromInvoice)

	/* User Payment Cancellation */
	mux.Post("/payment/:id/paymentcancellation", api.PaymentCancellationCreateFromPayment)
	mux.Get("/paymentcancellation/:id", api.PaymentCancellationView)
	mux.Put("/paymentcancellation/:id", api.PaymentCancellationUpdate)

	/* User Shipments */
	mux.Get("/account/:id/shipments", api.ShipmentsListFromAccount)
	mux.Get("/invoice/:id/shipments", api.ShipmentsListFromInvoice)
	mux.Get("/shipment/:id", api.ShipmentView)
	mux.Put("/shipment/:id", api.ShipmentUpdate)

	/* Subscriptions */
	mux.Get("/account/:id/subscriptions", api.SubscriptionsListFromAccount)
	mux.Get("/subscription/:id", api.SubscriptionView)
	mux.Put("/subscription/:id", api.SubscriptionUpdate)
	mux.Post("/account/:id/subscription", api.SubscriptionCreateFromAccount)

	return api
}
