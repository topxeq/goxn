package goxn

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/topxeq/qlang"
	_ "github.com/topxeq/qlang/lib/builtin" // 导入 builtin 包
	_ "github.com/topxeq/qlang/lib/chan"
	"github.com/topxeq/sqltk"

	execq "github.com/topxeq/qlang/exec"
	"github.com/topxeq/qlang/spec"

	// import standard packages
	qlarchivezip "github.com/topxeq/qlang/lib/archive/zip"
	qlbufio "github.com/topxeq/qlang/lib/bufio"
	qlbytes "github.com/topxeq/qlang/lib/bytes"

	qlcrypto "github.com/topxeq/qlang/lib/crypto"
	qlcryptoaes "github.com/topxeq/qlang/lib/crypto/aes"
	qlcryptocipher "github.com/topxeq/qlang/lib/crypto/cipher"
	qlcryptohmac "github.com/topxeq/qlang/lib/crypto/hmac"
	qlcryptomd5 "github.com/topxeq/qlang/lib/crypto/md5"
	qlcryptorand "github.com/topxeq/qlang/lib/crypto/rand"
	qlcryptorsa "github.com/topxeq/qlang/lib/crypto/rsa"
	qlcryptosha1 "github.com/topxeq/qlang/lib/crypto/sha1"
	qlcryptosha256 "github.com/topxeq/qlang/lib/crypto/sha256"
	qlcryptox509 "github.com/topxeq/qlang/lib/crypto/x509"

	qldatabasesql "github.com/topxeq/qlang/lib/database/sql"

	qlencodingbase64 "github.com/topxeq/qlang/lib/encoding/base64"
	qlencodingbinary "github.com/topxeq/qlang/lib/encoding/binary"
	qlencodingcsv "github.com/topxeq/qlang/lib/encoding/csv"
	qlencodinggob "github.com/topxeq/qlang/lib/encoding/gob"
	qlencodinghex "github.com/topxeq/qlang/lib/encoding/hex"
	qlencodingjson "github.com/topxeq/qlang/lib/encoding/json"
	qlencodingpem "github.com/topxeq/qlang/lib/encoding/pem"
	qlencodingxml "github.com/topxeq/qlang/lib/encoding/xml"

	qlerrors "github.com/topxeq/qlang/lib/errors"
	qlflag "github.com/topxeq/qlang/lib/flag"
	qlfmt "github.com/topxeq/qlang/lib/fmt"

	qlhashfnv "github.com/topxeq/qlang/lib/hash/fnv"

	qlhtml "github.com/topxeq/qlang/lib/html"
	qlhtmltemplate "github.com/topxeq/qlang/lib/html/template"

	qlimage "github.com/topxeq/qlang/lib/image"
	qlimage_color "github.com/topxeq/qlang/lib/image/color"
	qlimage_color_palette "github.com/topxeq/qlang/lib/image/color/palette"
	qlimage_draw "github.com/topxeq/qlang/lib/image/draw"
	qlimage_gif "github.com/topxeq/qlang/lib/image/gif"
	qlimage_jpeg "github.com/topxeq/qlang/lib/image/jpeg"
	qlimage_png "github.com/topxeq/qlang/lib/image/png"

	qlio "github.com/topxeq/qlang/lib/io"
	qlio_fs "github.com/topxeq/qlang/lib/io/fs"
	qlioioutil "github.com/topxeq/qlang/lib/io/ioutil"

	qllog "github.com/topxeq/qlang/lib/log"

	qlmath "github.com/topxeq/qlang/lib/math"
	qlmathbig "github.com/topxeq/qlang/lib/math/big"
	qlmathbits "github.com/topxeq/qlang/lib/math/bits"
	qlmathrand "github.com/topxeq/qlang/lib/math/rand"

	qlnet "github.com/topxeq/qlang/lib/net"
	qlnethttp "github.com/topxeq/qlang/lib/net/http"
	qlnet_http_cookiejar "github.com/topxeq/qlang/lib/net/http/cookiejar"
	qlnet_http_httputil "github.com/topxeq/qlang/lib/net/http/httputil"
	qlnet_mail "github.com/topxeq/qlang/lib/net/mail"
	qlnet_rpc "github.com/topxeq/qlang/lib/net/rpc"
	qlnet_rpc_jsonrpc "github.com/topxeq/qlang/lib/net/rpc/jsonrpc"
	qlnet_smtp "github.com/topxeq/qlang/lib/net/smtp"
	qlneturl "github.com/topxeq/qlang/lib/net/url"

	qlos "github.com/topxeq/qlang/lib/os"
	qlos_exec "github.com/topxeq/qlang/lib/os/exec"
	qlos_signal "github.com/topxeq/qlang/lib/os/signal"
	qlos_user "github.com/topxeq/qlang/lib/os/user"

	qlpath "github.com/topxeq/qlang/lib/path"
	qlpathfilepath "github.com/topxeq/qlang/lib/path/filepath"

	qlreflect "github.com/topxeq/qlang/lib/reflect"
	qlregexp "github.com/topxeq/qlang/lib/regexp"
	qlruntime "github.com/topxeq/qlang/lib/runtime"
	qlruntimedebug "github.com/topxeq/qlang/lib/runtime/debug"

	qlsort "github.com/topxeq/qlang/lib/sort"
	qlstrconv "github.com/topxeq/qlang/lib/strconv"
	qlstrings "github.com/topxeq/qlang/lib/strings"
	qlsync "github.com/topxeq/qlang/lib/sync"

	qltext_template "github.com/topxeq/qlang/lib/text/template"
	qltime "github.com/topxeq/qlang/lib/time"

	qlunicode "github.com/topxeq/qlang/lib/unicode"
	qlunicode_utf8 "github.com/topxeq/qlang/lib/unicode/utf8"

	// import 3rd party packages
	qlgithubbeeviketree "github.com/topxeq/qlang/lib/github.com/beevik/etree"
	qlgithubtopxeqsqltk "github.com/topxeq/qlang/lib/github.com/topxeq/sqltk"
	qlgithubtopxeqtk "github.com/topxeq/qlang/lib/github.com/topxeq/tk"

	qlgithub_fogleman_gg "github.com/topxeq/qlang/lib/github.com/fogleman/gg"

	qlgithub_360EntSecGroupSkylar_excelize "github.com/topxeq/qlang/lib/github.com/360EntSecGroup-Skylar/excelize"

	qlgithub_stretchr_objx "github.com/topxeq/qlang/lib/github.com/stretchr/objx"

	qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi "github.com/topxeq/qlang/lib/github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	qlgithub_topxeq_afero "github.com/topxeq/qlang/lib/github.com/topxeq/afero"

	qlgithub_domodwyer_mailyak "github.com/topxeq/qlang/lib/github.com/domodwyer/mailyak"

	qlgithub_topxeq_socks "github.com/topxeq/qlang/lib/github.com/topxeq/socks"

	qlgithub_topxeq_regexpx "github.com/topxeq/qlang/lib/github.com/topxeq/regexpx"

	qlgithub_topxeq_xmlx "github.com/topxeq/qlang/lib/github.com/topxeq/xmlx"

	qlgithub_topxeq_awsapi "github.com/topxeq/qlang/lib/github.com/topxeq/awsapi"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/topxeq/tk"
)

var versionG = "0.96a"

var notFoundG = interface{}(errors.New("not found"))

var initFlag bool = false

func qlEval(strA string) string {
	vmT := qlang.New("-noexit")

	errT := vmT.SafeEval(strA)

	if errT != nil {
		return tk.ErrStr(errT.Error())
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		return tk.Spr("%v", rs)
	}

	return ""
}

// native functions

// func printValue(nameA string) {

// 	v, ok := qlVMG.GetVar(nameA)

// 	if !ok {
// 		tk.Pl("no variable by the name found: %v", nameA)
// 		return
// 	}

// 	tk.Pl("%v(%T): %v", nameA, v, v)

// }

// func defined(nameA string) bool {

// 	_, ok := qlVMG.GetVar(nameA)

// 	return ok

// }

var leBufG []string

func leClear() {
	leBufG = make([]string, 0, 100)
}

func leLoadString(strA string) {
	if leBufG == nil {
		leClear()
	}

	leBufG = tk.SplitLines(strA)
}

func leSaveString() string {
	if leBufG == nil {
		leClear()
	}

	return tk.JoinLines(leBufG)
}

func leLoadFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	strT, errT := tk.LoadStringFromFileE(fileNameA)

	if errT != nil {
		return errT
	}

	leBufG = tk.SplitLines(strT)
	// leBufG, errT = tk.LoadStringListBuffered(fileNameA, false, false)

	return nil
}

func leSaveFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	var errT error

	textT := tk.JoinLines(leBufG)

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	errT = tk.SaveStringToFileE(textT, fileNameA)

	return errT
}

func leLoadClip() error {
	if leBufG == nil {
		leClear()
	}

	textT := tk.GetClipText()

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	leBufG = tk.SplitLines(textT)

	return nil
}

func leSaveClip() error {
	if leBufG == nil {
		leClear()
	}

	textT := tk.JoinLines(leBufG)

	if tk.IsErrStr(textT) {
		return tk.Errf(tk.GetErrStr(textT))
	}

	return tk.SetClipText(textT)
}

func leViewAll(argsA ...string) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if tk.IfSwitchExistsWhole(argsA, "-nl") {
		textT := tk.JoinLines(leBufG)

		tk.Pln(textT)

	} else {
		for i, v := range leBufG {
			tk.Pl("%v: %v", i, v)
		}
	}

	return nil
}

func leViewLine(idxA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	tk.Pln(leBufG[idxA])

	return nil
}

func leGetLine(idxA int) string {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.ErrStrf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.ErrStrf("line index out of range")
	}

	return leBufG[idxA]
}

func leSetLine(idxA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	leBufG[idxA] = strA

	return nil
}

func leSetLines(startA int, endA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	if startA > endA {
		return tk.Errf("start index greater than end index")
	}

	listT := tk.SplitLines(strA)

	if endA < 0 {
		rs := make([]string, 0, len(leBufG)+len(listT))

		rs = append(rs, listT...)
		rs = append(rs, leBufG...)

		leBufG = rs

		return nil
	}

	if startA >= len(leBufG) {
		leBufG = append(leBufG, listT...)

		return nil
	}

	if startA < 0 {
		startA = 0
	}

	if endA >= len(leBufG) {
		endA = len(leBufG) - 1
	}

	rs := make([]string, 0, len(leBufG)+len(listT)-1)

	rs = append(rs, leBufG[:startA]...)
	rs = append(rs, listT...)
	rs = append(rs, leBufG[endA+1:]...)

	leBufG = rs

	return nil
}

func leInsertLine(idxA int, strA string) error {
	if leBufG == nil {
		leClear()
	}

	// if leBufG == nil {
	// 	return tk.Errf("buffer not initalized")
	// }

	// if idxA < 0 || idxA >= len(leBufG) {
	// 	return tk.Errf("line index out of range")
	// }

	if idxA < 0 {
		idxA = 0
	}

	listT := tk.SplitLines(strA)

	if idxA >= len(leBufG) {
		leBufG = append(leBufG, listT...)
	} else {
		rs := make([]string, 0, len(leBufG)+1)

		rs = append(rs, leBufG[:idxA]...)
		rs = append(rs, listT...)
		rs = append(rs, leBufG[idxA:]...)

		leBufG = rs

	}

	return nil
}

func leAppendLine(strA string) error {
	if leBufG == nil {
		leClear()
	}

	// if leBufG == nil {
	// 	return tk.Errf("buffer not initalized")
	// }

	// if idxA < 0 || idxA >= len(leBufG) {
	// 	return tk.Errf("line index out of range")
	// }

	listT := tk.SplitLines(strA)

	leBufG = append(leBufG, listT...)

	return nil
}

func leRemoveLine(idxA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		return tk.Errf("line index out of range")
	}

	rs := make([]string, 0, len(leBufG)+1)

	rs = append(rs, leBufG[:idxA]...)
	rs = append(rs, leBufG[idxA+1:]...)

	leBufG = rs

	return nil
}

func leRemoveLines(startA int, endA int) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if startA < 0 || startA >= len(leBufG) {
		return tk.Errf("start line index out of range")
	}

	if endA < 0 || endA >= len(leBufG) {
		return tk.Errf("end line index out of range")
	}

	if startA > endA {
		return tk.Errf("start line index greater than end line index")
	}

	rs := make([]string, 0, len(leBufG)+1)

	rs = append(rs, leBufG[:startA]...)
	rs = append(rs, leBufG[endA+1:]...)

	leBufG = rs

	return nil
}

func nilToEmpty(vA interface{}, argsA ...string) string {

	if vA == nil {
		return ""
	}

	if vA == spec.Undefined {
		return ""
	}

	if tk.IsNil(vA) {
		return ""
	}

	if (argsA != nil) && (len(argsA) > 0) {
		vf, ok := vA.(float64)
		if ok {
			if tk.IfSwitchExistsWhole(argsA, "-nofloat") {
				return tk.ToStr(int(vf))
			} else {
				return tk.Float64ToStr(vA.(float64))
			}
		}
	}

	return fmt.Sprintf("%v", vA)

}

func isValid(vA interface{}, argsA ...string) bool {

	if vA == nil {
		return false
	}

	if vA == spec.Undefined {
		return false
	}

	if tk.IsNil(vA) {
		return false
	}

	if (argsA != nil) && (len(argsA) > 0) {
		typeT := fmt.Sprintf("%T", vA)

		if typeT == argsA[0] {
			return true
		} else {
			return false
		}
	}

	return true
}

func logPrint(formatA string, argsA ...interface{}) {
	tk.Pl(formatA, argsA...)
	tk.LogWithTimeCompact(formatA, argsA...)
}

// -1 return random item
func getArrayItem(aryA interface{}, idxA int, defaultA ...interface{}) interface{} {
	var hasDefaultT = false
	if len(defaultA) > 0 {
		hasDefaultT = true
	}

	if aryA == nil {
		if hasDefaultT {
			return defaultA[0]
		}

		return ""
	}

	switch aryT := aryA.(type) {
	case []interface{}:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []string:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []int:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []float64:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	case []bool:
		lenT := len(aryT)

		if lenT < 0 || (idxA < -1 || idxA >= lenT) {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		if idxA == -1 {
			return aryT[tk.GetRandomIntLessThan(lenT)]
		}

		return aryT[idxA]
	}

	return ""

}

func getMapItem(mapA interface{}, keyA string, defaultA ...interface{}) interface{} {
	var hasDefaultT = false
	if len(defaultA) > 0 {
		hasDefaultT = true
	}

	if mapA == nil {
		if hasDefaultT {
			return defaultA[0]
		}

		return ""
	}

	switch mapT := mapA.(type) {
	case map[string]interface{}:
		itemT, ok := mapT[keyA]
		if !ok {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		return itemT
	case map[string]string:
		itemT, ok := mapT[keyA]
		if !ok {
			if hasDefaultT {
				return defaultA[0]
			}

			return ""
		}

		return itemT
	}

	return ""
}

func NewFuncB(funcA interface{}) func() {
	funcT := (funcA).(*execq.Function)
	f := func() {
		funcT.Call(execq.NewStack())

		return
	}

	return f
}

func NewFuncInterfaceInterfaceErrorB(funcA interface{}) func(interface{}) (interface{}, error) {
	funcT := (funcA).(*execq.Function)
	f := func(s interface{}) (interface{}, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r[1] == nil {
			return r[0].(interface{}), nil
		}

		return r[0].(interface{}), r[1].(error)
	}

	return f
}

func NewFuncStringStringErrorB(funcA interface{}) func(string) (string, error) {
	funcT := (funcA).(*execq.Function)
	f := func(s string) (string, error) {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if r == nil {
			return "", tk.Errf("nil result")
		}

		if len(r) < 2 {
			return "", tk.Errf("incorrect return argument count")
		}

		if r[1] == nil {
			return r[0].(string), nil
		}

		return r[0].(string), r[1].(error)
	}

	return f
}

func NewFuncStringStringB(funcA interface{}) func(string) string {
	funcT := (funcA).(*execq.Function)
	f := func(s string) string {
		return funcT.Call(execq.NewStack(), s).(string)
	}

	return f
}

func intToStr(n interface{}, defaultA ...string) string {
	var defaultT string = ""
	if (defaultA != nil) && (len(defaultA) > 0) {
		defaultT = defaultA[0]
	}

	switch nv := n.(type) {
	case int:
		return fmt.Sprintf("%v", nv)
	case int8:
		return fmt.Sprintf("%v", nv)
	case int16:
		return fmt.Sprintf("%v", nv)
	case int32:
		return fmt.Sprintf("%v", nv)
	case int64:
		return fmt.Sprintf("%v", nv)
	case float64:
		return tk.Float64ToStr(nv)
	case float32:
		tmps := fmt.Sprintf("%f", nv)
		if tk.Contains(tmps, ".") {
			tmps = strings.TrimRight(tmps, "0")
			tmps = strings.TrimRight(tmps, ".")
		}

		return tmps
	case string:
		nT, errT := strconv.ParseInt(nv, 10, 0)
		if errT != nil {
			return defaultT
		}

		return fmt.Sprintf("%v", nT)
	default:
		nT, errT := strconv.ParseInt(fmt.Sprintf("%v", nv), 10, 0)
		if errT != nil {
			return defaultT
		}

		return fmt.Sprintf("%v", nT)
	}

}

func strJoin(aryA interface{}, sepA string, defaultA ...string) string {
	var defaultT string = ""
	if (defaultA != nil) && (len(defaultA) > 0) {
		defaultT = defaultA[0]
	}

	if aryA == nil {
		return defaultT
	}

	switch v := aryA.(type) {
	case []string:
		return strings.Join(v, sepA)
	case []interface{}:
		var bufT strings.Builder
		for j, jv := range v {
			if j > 0 {
				bufT.WriteString(sepA)
			}

			bufT.WriteString(fmt.Sprintf("%v", jv))
		}

		return bufT.String()
	}

	return defaultT
}

func isDefined(vA interface{}) bool {
	if vA == spec.Undefined {
		return false
	}

	return true
}

func strToTime(strA string, formatA ...string) interface{} {
	formatT := tk.TimeFormat

	if (formatA != nil) && (len(formatA) > 0) {
		formatT = formatA[0]
	}

	timeT, errT := tk.StrToTimeByFormat(strA, formatT)

	if errT != nil {
		return spec.Undefined
	}

	return timeT
}

func importQLNonGUIPackages() {
	// import native functions and global variables
	var defaultExports = map[string]interface{}{
		// common related
		"pass": tk.Pass,
		// "defined":       defined,
		"isDefined":     isDefined,
		"isValid":       isValid,
		"eval":          qlEval,
		"typeOf":        tk.TypeOfValue,
		"typeOfReflect": tk.TypeOfValueReflect,
		"exit":          tk.Exit,
		"setValue":      tk.SetValue,
		"getValue":      tk.GetValue,
		"setVar":        tk.SetVar,
		"getVar":        tk.GetVar,
		"isNil":         tk.IsNil,
		"deepClone":     tk.DeepClone,
		"deepCopy":      tk.DeepCopyFromTo,
		// "run":           runFile,
		// "runCode":       runCode,
		// "runScript":     runScript,
		// "magic":         magic,

		// output related
		"pr":        tk.Pr,
		"pln":       tk.Pln,
		"prf":       tk.Printf,
		"printfln":  tk.Pl,
		"pl":        tk.Pl,
		"sprintf":   fmt.Sprintf,
		"spr":       fmt.Sprintf,
		"fprintf":   fmt.Fprintf,
		"plv":       tk.Plv,
		"plvx":      tk.Plvx,
		"plNow":     tk.PlNow,
		"plVerbose": tk.PlVerbose,
		// "pv":        printValue,
		"plvsr":  tk.Plvsr,
		"plerr":  tk.PlErr,
		"plExit": tk.PlAndExit,

		// math related
		"bitXor": tk.BitXor,

		// string related
		"trim":             tk.Trim,
		"strTrim":          tk.Trim,
		"strContains":      strings.Contains,
		"strReplace":       tk.Replace,
		"strJoin":          strJoin,
		"strSplit":         strings.Split,
		"getNowStr":        tk.GetNowTimeStringFormal,
		"getNowStrCompact": tk.GetNowTimeString,
		"splitLines":       tk.SplitLines,
		"startsWith":       tk.StartsWith,
		"endsWith":         tk.EndsWith,
		"strStartsWith":    tk.StartsWith,
		"strEndsWith":      tk.EndsWith,

		// regex related
		"regMatch":     tk.RegMatchX,
		"regContains":  tk.RegContainsX,
		"regFind":      tk.RegFindFirstX,
		"regFindAll":   tk.RegFindAllX,
		"regFindIndex": tk.RegFindFirstIndexX,
		"regReplace":   tk.RegReplaceX,
		"regSplit":     tk.RegSplitX,

		// conversion related
		"nilToEmpty": nilToEmpty,
		"strToInt":   tk.StrToIntWithDefaultValue,
		"strToTime":  strToTime,
		"timeToStr":  tk.FormatTime,
		"intToStr":   tk.IntToStrX,
		"floatToStr": tk.Float64ToStr,
		"toStr":      tk.ToStr,
		"toInt":      tk.ToInt,
		"toFloat":    tk.ToFloat,
		"toLower":    strings.ToLower,
		"toUpper":    strings.ToUpper,

		// array/map related
		"remove":       tk.RemoveItemsInArray,
		"getMapString": tk.SafelyGetStringForKeyWithDefault,
		"getMapItem":   getMapItem,
		"getArrayItem": getArrayItem,
		"joinList":     tk.JoinList,

		// error related
		"isError":          tk.IsError,
		"isErrStr":         tk.IsErrStr,
		"checkError":       tk.CheckError,
		"checkErrorString": tk.CheckErrorString,
		"checkErrf":        tk.CheckErrf,
		"checkErrStr":      tk.CheckErrStr,
		"checkErrStrf":     tk.CheckErrStrf,
		"fatalf":           tk.Fatalf,
		"errStr":           tk.ErrStr,
		"errStrf":          tk.ErrStrF,
		"getErrStr":        tk.GetErrStr,
		"errf":             tk.Errf,

		// encode/decode related
		"xmlEncode":    tk.EncodeToXMLString,
		"htmlEncode":   tk.EncodeHTML,
		"htmlDecode":   tk.DecodeHTML,
		"base64Encode": tk.EncodeToBase64,
		"base64Decode": tk.DecodeFromBase64,
		"md5Encode":    tk.MD5Encrypt,
		"md5":          tk.MD5Encrypt,
		"hexEncode":    tk.StrToHex,
		"hexDecode":    tk.HexToStr,
		"jsonEncode":   tk.ObjectToJSON,
		"jsonDecode":   tk.JSONToObject,
		"toJSON":       tk.ToJSONX,
		"fromJSON":     tk.FromJSONWithDefault,
		"simpleEncode": tk.EncodeStringCustomEx,
		"simpleDecode": tk.DecodeStringCustom,

		// input related
		"getInput":     tk.GetUserInput,
		"getInputf":    tk.GetInputf,
		"getPasswordf": tk.GetInputPasswordf,

		// log related
		"setLogFile": tk.SetLogFile,
		"logf":       tk.LogWithTimeCompact,
		"logPrint":   logPrint,

		// system related
		"getClipText":  tk.GetClipText,
		"setClipText":  tk.SetClipText,
		"systemCmd":    tk.SystemCmd,
		"ifFileExists": tk.IfFileExists,
		"fileExists":   tk.IfFileExists,
		"joinPath":     filepath.Join,
		"getFileSize":  tk.GetFileSizeCompact,
		"getFileList":  tk.GetFileList,
		"loadText":     tk.LoadStringFromFile,
		"saveText":     tk.SaveStringToFile,
		"loadBytes":    tk.LoadBytesFromFileE,
		"saveBytes":    tk.SaveBytesToFile,
		"sleep":        tk.SleepSeconds,
		"sleepSeconds": tk.SleepSeconds,

		// command-line
		"getParameter":   tk.GetParameterByIndexWithDefaultValue,
		"getSwitch":      tk.GetSwitchWithDefaultValue,
		"getIntSwitch":   tk.GetSwitchWithDefaultIntValue,
		"switchExists":   tk.IfSwitchExistsWhole,
		"ifSwitchExists": tk.IfSwitchExistsWhole,

		// network related
		"newSSHClient":         tk.NewSSHClient,
		"mapToPostData":        tk.MapToPostData,
		"getWebPage":           tk.DownloadPageUTF8,
		"downloadFile":         tk.DownloadFile,
		"httpRequest":          tk.RequestX,
		"getFormValue":         tk.GetFormValueWithDefaultValue,
		"formValueExist":       tk.IfFormValueExists,
		"ifFormValueExist":     tk.IfFormValueExists,
		"formToMap":            tk.FormToMap,
		"generateJSONResponse": tk.GenerateJSONPResponseWithMore,

		// database related
		"dbConnect":     sqltk.ConnectDBX,
		"dbExec":        sqltk.ExecDBX,
		"dbQuery":       sqltk.QueryDBX,
		"dbQueryCount":  sqltk.QueryCountX,
		"dbQueryString": sqltk.QueryStringX,

		// line editor related
		"leClear":       leClear,
		"leLoadStr":     leLoadString,
		"leSetAll":      leLoadString,
		"leSaveStr":     leSaveString,
		"leGetAll":      leSaveString,
		"leLoad":        leLoadFile,
		"leLoadFile":    leLoadFile,
		"leSave":        leSaveFile,
		"leSaveFile":    leSaveFile,
		"leLoadClip":    leLoadClip,
		"leSaveClip":    leSaveClip,
		"leInsert":      leInsertLine,
		"leInsertLine":  leInsertLine,
		"leAppend":      leAppendLine,
		"leAppendLine":  leAppendLine,
		"leSet":         leSetLine,
		"leSetLine":     leSetLine,
		"leSetLines":    leSetLines,
		"leRemove":      leRemoveLine,
		"leRemoveLine":  leRemoveLine,
		"leRemoveLines": leRemoveLines,
		"leViewAll":     leViewAll,
		"leView":        leViewLine,

		// GUI related start
		// gui related
		// "initGUI":             initGUI,
		// "getConfirmGUI":       getConfirmGUI,
		// "showInfoGUI":         showInfoGUI,
		// "showErrorGUI":        showErrorGUI,
		// "selectFileToSaveGUI": selectFileToSaveGUI,
		// "selectFileGUI":       selectFileGUI,
		// "selectDirectoryGUI":  selectDirectoryGUI,

		// GUI related end

		// misc
		"newFunc":    NewFuncB,
		"newFuncIIE": NewFuncInterfaceInterfaceErrorB,
		"newFuncSSE": NewFuncStringStringErrorB,
		"newFuncSS":  NewFuncStringStringB,

		// global variables
		"timeFormatG":        tk.TimeFormat,
		"timeFormatCompactG": tk.TimeFormatCompact,

		// "scriptPathG": scriptPathG,
		"versionG": versionG,
		"leBufG":   leBufG,
	}

	qlang.Import("", defaultExports)

	qlang.Import("archive_zip", qlarchivezip.Exports)
	qlang.Import("bufio", qlbufio.Exports)
	qlang.Import("bytes", qlbytes.Exports)

	qlang.Import("crypto", qlcrypto.Exports)
	qlang.Import("crypto_aes", qlcryptoaes.Exports)
	qlang.Import("crypto_cipher", qlcryptocipher.Exports)
	qlang.Import("crypto_hmac", qlcryptohmac.Exports)
	qlang.Import("crypto_md5", qlcryptomd5.Exports)
	qlang.Import("crypto_rand", qlcryptorand.Exports)
	qlang.Import("crypto_rsa", qlcryptorsa.Exports)
	qlang.Import("crypto_sha256", qlcryptosha256.Exports)
	qlang.Import("crypto_sha1", qlcryptosha1.Exports)
	qlang.Import("crypto_x509", qlcryptox509.Exports)

	qlang.Import("database_sql", qldatabasesql.Exports)

	qlang.Import("encoding_pem", qlencodingpem.Exports)
	qlang.Import("encoding_base64", qlencodingbase64.Exports)
	qlang.Import("encoding_binary", qlencodingbinary.Exports)
	qlang.Import("encoding_csv", qlencodingcsv.Exports)
	qlang.Import("encoding_gob", qlencodinggob.Exports)
	qlang.Import("encoding_hex", qlencodinghex.Exports)
	qlang.Import("encoding_json", qlencodingjson.Exports)
	qlang.Import("encoding_xml", qlencodingxml.Exports)

	qlang.Import("errors", qlerrors.Exports)

	qlang.Import("flag", qlflag.Exports)
	qlang.Import("fmt", qlfmt.Exports)

	qlang.Import("hash_fnv", qlhashfnv.Exports)

	qlang.Import("html", qlhtml.Exports)
	qlang.Import("html_template", qlhtmltemplate.Exports)

	qlang.Import("image", qlimage.Exports)
	qlang.Import("image_color", qlimage_color.Exports)
	qlang.Import("image_color_palette", qlimage_color_palette.Exports)
	qlang.Import("image_draw", qlimage_draw.Exports)
	qlang.Import("image_gif", qlimage_gif.Exports)
	qlang.Import("image_jpeg", qlimage_jpeg.Exports)
	qlang.Import("image_png", qlimage_png.Exports)

	qlang.Import("io", qlio.Exports)
	qlang.Import("io_ioutil", qlioioutil.Exports)
	qlang.Import("io_fs", qlio_fs.Exports)

	qlang.Import("log", qllog.Exports)

	qlang.Import("math", qlmath.Exports)
	qlang.Import("math_big", qlmathbig.Exports)
	qlang.Import("math_bits", qlmathbits.Exports)
	qlang.Import("math_rand", qlmathrand.Exports)

	qlang.Import("net", qlnet.Exports)
	qlang.Import("net_http", qlnethttp.Exports)
	qlang.Import("net_http_cookiejar", qlnet_http_cookiejar.Exports)
	qlang.Import("net_http_httputil", qlnet_http_httputil.Exports)
	qlang.Import("net_mail", qlnet_mail.Exports)
	qlang.Import("net_rpc", qlnet_rpc.Exports)
	qlang.Import("net_rpc_jsonrpc", qlnet_rpc_jsonrpc.Exports)
	qlang.Import("net_smtp", qlnet_smtp.Exports)
	qlang.Import("net_url", qlneturl.Exports)

	qlang.Import("os", qlos.Exports)
	qlang.Import("os_exec", qlos_exec.Exports)
	qlang.Import("os_signal", qlos_signal.Exports)
	qlang.Import("os_user", qlos_user.Exports)
	qlang.Import("path", qlpath.Exports)
	qlang.Import("path_filepath", qlpathfilepath.Exports)

	qlang.Import("reflect", qlreflect.Exports)
	qlang.Import("regexp", qlregexp.Exports)

	qlang.Import("runtime", qlruntime.Exports)
	qlang.Import("runtime_debug", qlruntimedebug.Exports)

	qlang.Import("sort", qlsort.Exports)
	qlang.Import("strconv", qlstrconv.Exports)
	qlang.Import("strings", qlstrings.Exports)
	qlang.Import("sync", qlsync.Exports)

	qlang.Import("text_template", qltext_template.Exports)
	qlang.Import("time", qltime.Exports)

	qlang.Import("unicode", qlunicode.Exports)
	qlang.Import("unicode_utf8", qlunicode_utf8.Exports)

	// 3rd party

	qlang.Import("github_topxeq_tk", qlgithubtopxeqtk.Exports)
	qlang.Import("tk", qlgithubtopxeqtk.Exports)

	qlang.Import("github_beevik_etree", qlgithubbeeviketree.Exports)
	qlang.Import("github_topxeq_sqltk", qlgithubtopxeqsqltk.Exports)

	qlang.Import("github_topxeq_xmlx", qlgithub_topxeq_xmlx.Exports)

	qlang.Import("github_topxeq_awsapi", qlgithub_topxeq_awsapi.Exports)

	qlang.Import("github_fogleman_gg", qlgithub_fogleman_gg.Exports)

	qlang.Import("github_360EntSecGroupSkylar_excelize", qlgithub_360EntSecGroupSkylar_excelize.Exports)

	qlang.Import("github_stretchr_objx", qlgithub_stretchr_objx.Exports)

	qlang.Import("github_aliyun_alibabacloudsdkgo_services_dysmsapi", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)

	qlang.Import("github_topxeq_afero", qlgithub_topxeq_afero.Exports)

	qlang.Import("github_topxeq_socks", qlgithub_topxeq_socks.Exports)

	qlang.Import("github_topxeq_regexpx", qlgithub_topxeq_regexpx.Exports)

	qlang.Import("github_domodwyer_mailyak", qlgithub_domodwyer_mailyak.Exports)

}

func InitVM() {
	// qlang.SetOnPop(func(v interface{}) {
	// 	retG = v
	// })

	// qlang.SetDumpCode("1")

	if !initFlag {
		initFlag = true
		importQLNonGUIPackages()
	}

}

func RunScript(codeA, inputA string, argsA []string, parametersA map[string]string, optionsA ...string) (string, error) {
	if tk.IfSwitchExists(optionsA, "-verbose") {
		tk.Pl("Starting...")
	}

	if !initFlag {
		initFlag = true
		importQLNonGUIPackages()
	}

	vmT := qlang.New("-noexit")

	vmT.SetVar("inputG", inputA)

	vmT.SetVar("argsG", argsA)

	vmT.SetVar("basePathG", tk.GetSwitch(optionsA, "-base=", ""))

	vmT.SetVar("paraMapG", parametersA)

	retT := ""

	errT := vmT.SafeEval(codeA)

	if errT != nil {
		return retT, errT
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		if rs != nil {
			strT, ok := rs.(string)
			if ok {
				return strT, nil
			}

			return fmt.Sprintf("%v", rs), nil
		}

		return retT, nil
	}

	return retT, nil
}

func doJapi(resA http.ResponseWriter, reqA *http.Request) string {
	if reqA != nil {
		reqA.ParseForm()
	}

	reqT := tk.GetFormValueWithDefaultValue(reqA, "req", "")

	if resA != nil {
		resA.Header().Set("Access-Control-Allow-Origin", "*")
		resA.Header().Set("Access-Control-Allow-Headers", "*")
		resA.Header().Set("Content-Type", "text/json;charset=utf-8")
	}

	resA.WriteHeader(http.StatusOK)

	vo := tk.GetFormValueWithDefaultValue(reqA, "vo", "")

	var paraMapT map[string]string
	var errT error

	if vo == "" {
		paraMapT = tk.FormToMap(reqA.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			return tk.GenerateJSONPResponse("success", "invalid vo format", reqA)
		}
	}

	switch reqT {
	case "debug":
		return tk.GenerateJSONPResponse("success", fmt.Sprintf("%v", reqA), reqA)

	case "requestinfo":
		rs := tk.Spr("%#v", reqA)

		return tk.GenerateJSONPResponse("success", rs, reqA)

	case "test":

		return tk.GenerateJSONPResponse("success", "test respone", reqA)

	case "runScript":
		scriptT := paraMapT["script"]
		if scriptT == "" {
			return tk.GenerateJSONPResponse("fail", fmt.Sprintf("empty script"), reqA)
		}

		retT, errT := RunScript(scriptT, paraMapT["input"], nil, nil)

		var errStrT string = ""

		if errT != nil {
			errStrT = fmt.Sprintf("%v", errT)
		}

		return tk.GenerateJSONPResponseWithMore("success", retT, reqA, "Error", errStrT)

	case "runFileScript":
		scriptT := paraMapT["script"]
		if scriptT == "" {
			return tk.GenerateJSONPResponse("fail", tk.Spr("empty script"), reqA)
		}

		baseDirT := paraMapT["base"]
		if baseDirT == "" {
			baseDirT = "."
		}

		fcT := tk.LoadStringFromFile(filepath.Join(baseDirT, scriptT))
		if tk.IsErrStr(fcT) {
			return tk.GenerateJSONPResponseWithMore("fail", "", reqA, "Error", tk.GetErrStr(fcT))
		}

		retT, errT := RunScript(fcT, paraMapT["input"], nil, nil)

		var errStrT string = ""

		if errT != nil {
			errStrT = fmt.Sprintf("%v", errT)
		}

		return tk.GenerateJSONPResponseWithMore("success", retT, reqA, "Error", errStrT)
	}

	return tk.GenerateJSONPResponse("fail", "unknown request", reqA)

}

func japiHandler(w http.ResponseWriter, req *http.Request) {
	rs := doJapi(w, req)

	w.Write([]byte(rs))
}

func StartServer(portA string, codeA string) error {
	muxT := http.NewServeMux()

	if strings.ContainsAny(codeA, " /") {
		return tk.Errf("failed to start server: %v", "invalid password")
	}

	if codeA == "" {
		muxT.HandleFunc("/japi", japiHandler)
	} else {
		muxT.HandleFunc("/japi/"+codeA, japiHandler)
	}

	errT := http.ListenAndServe(portA, muxT)

	if errT != nil {
		return tk.Errf("failed to start server: %v", errT)
	}

	return nil
}
