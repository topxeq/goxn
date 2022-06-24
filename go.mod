module github.com/topxeq/goxn

go 1.16

require (
	gitee.com/topxeq/xie v0.0.0
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.10.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/godror/godror v0.31.1-0.20220304150828-9df83c110931
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/topxeq/charlang v0.0.0-20220308001517-79d2b54a8942
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0
	github.com/topxeq/tk v1.0.1
)

// replace github.com/topxeq/tk v0.0.0 => ../tk

// replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

// replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v1.0.1 => ../tk

// replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

replace gitee.com/topxeq/xie v0.0.0 => ../../../gitee.com/topxeq/xie

// replace github.com/topxeq/text v0.0.0 => ../text

// replace github.com/topxeq/charlang v0.0.0 => ../charlang

// replace github.com/topxeq/goph v0.0.0 => ../goph

// replace github.com/topxeq/go-sciter v0.0.0 => ../go-sciter

// replace github.com/topxeq/gods v0.0.0 => ../gods
