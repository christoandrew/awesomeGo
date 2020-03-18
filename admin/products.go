package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetProductsTable(ctx *context.Context) table.Table {

	productsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := productsTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Code", "code", db.Varchar)
	info.AddField("Price", "price", db.Int)

	info.SetTable("products").SetTitle("Products").SetDescription("Products")

	formList := productsTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Code", "code", db.Varchar, form.Text)
	formList.AddField("Price", "price", db.Int, form.Number)

	formList.SetTable("products").SetTitle("Products").SetDescription("Products")

	return productsTable
}
