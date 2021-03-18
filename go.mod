module github.com/topxeq/goxn

go 1.16

require (
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.985 // indirect
	github.com/beevik/etree v1.1.1-0.20200718192613-4a2f8b9d084c
	github.com/denisenkom/go-mssqldb v0.0.0-20200428022330-06a60b6afbbc
	github.com/go-sql-driver/mysql v1.5.0
	github.com/godror/godror v0.15.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/stretchr/objx v0.3.0
	github.com/topxeq/afero v0.0.0-20200914073911-38c8390e9ef4 // indirect
	github.com/topxeq/awsapi v0.0.0-20191115074250-1192cb0fdb97 // indirect
	github.com/topxeq/qlang v0.0.0-20210316075839-fafaa56f2136
	github.com/topxeq/sqltk v0.0.0-20210112052931-55ad87cc9be3
	github.com/topxeq/tk v0.0.0-20210318011618-0e03ab785f68
)

// replace github.com/topxeq/tk v0.0.0 => ../tk

replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1
