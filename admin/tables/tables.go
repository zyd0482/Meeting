package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "meeting_meets" => http://localhost:9033/admin/info/meeting_meets
//
var Generators = map[string]table.Generator{
	"meeting_meets": GetMeetingMeetsTable,
}
