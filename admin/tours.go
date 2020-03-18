package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetToursTable(ctx *context.Context) table.Table {

	toursTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := toursTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)

	info.SetTable("tours").SetTitle("Tours").SetDescription("Tours")

	formList := toursTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)

	formList.SetTable("tours").SetTitle("Tours").SetDescription("Tours")

	return toursTable
}
