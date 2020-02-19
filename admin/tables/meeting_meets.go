package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMeetingMeetsTable(ctx *context.Context) table.Table {

    meetingMeetsTable := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := meetingMeetsTable.GetInfo()
	
	info.AddField("Id","id", db.Int).FieldFilterable()
	info.AddField("Type","type", db.Tinyint)
	info.AddField("Banner","banner", db.Varchar)
	info.AddField("Title","title", db.Varchar)
	info.AddField("Start_at","start_at", db.Timestamp)
	info.AddField("Place","place", db.Varchar)
	info.AddField("Fee","fee", db.Int)
	info.AddField("Person","person", db.Int)
	info.AddField("Content","content", db.Text)
	info.AddField("State","state", db.Tinyint)
	info.AddField("Created_at","created_at", db.Timestamp)
	info.AddField("Updated_at","updated_at", db.Timestamp)
	info.AddField("Deleted_at","deleted_at", db.Timestamp)
	
	info.SetTable("meeting_meets").SetTitle("Meeting_meets").SetDescription("Meeting_meets")

	formList := meetingMeetsTable.GetForm()
	
	formList.AddField("Id","id",db.Int,form.Default).FieldNotAllowAdd()
	formList.AddField("Type","type",db.Tinyint,form.Number)
	formList.AddField("Banner","banner",db.Varchar,form.Text)
	formList.AddField("Title","title",db.Varchar,form.Text)
	formList.AddField("Start_at","start_at",db.Timestamp,form.Datetime)
	formList.AddField("Place","place",db.Varchar,form.Text)
	formList.AddField("Fee","fee",db.Int,form.Number)
	formList.AddField("Person","person",db.Int,form.Number)
	formList.AddField("Content","content",db.Text,form.RichText)
	formList.AddField("State","state",db.Tinyint,form.Number)
	formList.AddField("Created_at","created_at",db.Timestamp,form.Datetime)
	formList.AddField("Updated_at","updated_at",db.Timestamp,form.Datetime)
	formList.AddField("Deleted_at","deleted_at",db.Timestamp,form.Datetime)
	
	formList.SetTable("meeting_meets").SetTitle("Meeting_meets").SetDescription("Meeting_meets")

	return meetingMeetsTable
}