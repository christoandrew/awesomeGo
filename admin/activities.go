package admin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetActivitiesTable(ctx *context.Context) table.Table {

	activitiesTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := activitiesTable.GetInfo()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Created_at", "created_at", db.Datetime)
	info.AddField("Updated_at", "updated_at", db.Datetime)
	info.AddField("Deleted_at", "deleted_at", db.Datetime)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Description", "description", db.Varchar)

	info.SetTable("activities").SetTitle("Activities").SetDescription("Activities")

	formList := activitiesTable.GetForm()

	formList.AddField("Id", "id", db.Int, form.Default).FieldNotAllowAdd()
	formList.AddField("Created_at", "created_at", db.Datetime, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Datetime, form.Datetime)
	formList.AddField("Deleted_at", "deleted_at", db.Datetime, form.Datetime)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)

	formList.SetTable("activities").SetTitle("Activities").SetDescription("Activities")

	return activitiesTable
}