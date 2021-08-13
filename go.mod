module github.com/topxeq/goxn

go 1.16

require (
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc
	github.com/go-sql-driver/mysql v1.5.0
	github.com/godror/godror v0.15.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0
	github.com/topxeq/tk v0.0.0
)

// replace github.com/topxeq/tk v0.0.0 => ../tk

replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v0.0.0 => ../tk

replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

replace github.com/topxeq/text v0.0.0 => ../text
