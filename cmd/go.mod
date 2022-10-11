module github.com/topxeq/goxn/goxn

go 1.16

require (
	gitee.com/topxeq/xie v0.0.0
	github.com/topxeq/goxn v0.0.0
	github.com/topxeq/tk v1.0.1
)

// replace github.com/360EntSecGroup-Skylar/excelize v1.4.1 => github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2

// replace github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2 => github.com/360EntSecGroup-Skylar/excelize v1.4.1

replace github.com/topxeq/goxn v0.0.0 => ../../goxn

replace github.com/topxeq/qlang v0.0.0 => ../../qlang

replace github.com/topxeq/tk v1.0.1 => ../../tk

replace github.com/topxeq/sqltk v0.0.0 => ../../sqltk

replace gitee.com/topxeq/xie v0.0.0 => ../../../../gitee.com/topxeq/xie

// // replace github.com/topxeq/text v0.0.0 => ../../text

// replace github.com/topxeq/charlang v0.0.0 => ../../charlang

// replace github.com/topxeq/goph v0.0.0 => ../../goph

// replace github.com/topxeq/go-sciter v0.0.0 => ../../go-sciter

// replace github.com/topxeq/gods v0.0.0 => ../../gods
