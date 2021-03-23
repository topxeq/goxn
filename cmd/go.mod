module github.com/topxeq/goxn/goxn

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/topxeq/goxn v0.0.0
	github.com/topxeq/tk v0.0.0-20210318011618-0e03ab785f68
)

replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/goxn v0.0.0 => ../

