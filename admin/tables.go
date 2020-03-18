package admin

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "accommodations" => http://localhost:9033/admin/info/accommodations
// "activities" => http://localhost:9033/admin/info/activities
// "features" => http://localhost:9033/admin/info/features
// "lodges" => http://localhost:9033/admin/info/lodges
// "products" => http://localhost:9033/admin/info/products
// "tours" => http://localhost:9033/admin/info/tours
//
// example end
//
var Generators = map[string]table.Generator{
	"accommodations": GetAccommodationsTable,
	"activities":     GetActivitiesTable,
	"features":       GetFeaturesTable,
	"lodges":         GetLodgesTable,
	"products":       GetProductsTable,
	"tours":          GetToursTable,
	"users":          GetUserTable,

	// generators end
}
