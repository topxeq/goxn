module github.com/topxeq/goxn

go 1.16

require (
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c // indirect
	github.com/denisenkom/go-mssqldb v0.12.3
	github.com/go-sql-driver/mysql v1.7.0
	github.com/godror/godror v0.36.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/sijms/go-ora/v2 v2.5.25
	github.com/topxeq/charlang v0.0.0-20220722003130-49bba2664ad6
	github.com/topxeq/qlang v0.0.0
	github.com/topxeq/sqltk v0.0.0-20230223005953-f9932d23950c
	github.com/topxeq/tk v1.0.1
	github.com/topxeq/xie v0.0.0
)

// replace github.com/topxeq/tk v0.0.0 => ../tk

// replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

// replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/tk v1.0.1 => ../tk

// replace github.com/topxeq/xmlx v0.2.0 => ../xmlx

// replace github.com/topxeq/sqltk v0.0.0 => ../sqltk

replace github.com/topxeq/qlang v0.0.0 => ../qlang

replace github.com/topxeq/xie v0.0.0 => ../xie

// replace github.com/topxeq/text v0.0.0 => ../text

// replace github.com/topxeq/charlang v0.0.0 => ../charlang

// replace github.com/topxeq/goph v0.0.0 => ../goph

// replace github.com/topxeq/go-sciter v0.0.0 => ../go-sciter

// replace github.com/topxeq/gods v0.0.0 => ../gods
