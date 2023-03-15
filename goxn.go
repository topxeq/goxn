package goxn

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/topxeq/charlang"
	"github.com/topxeq/qlang"
	_ "github.com/topxeq/qlang/lib/builtin" // 导入 builtin 包
	_ "github.com/topxeq/qlang/lib/chan"
	"github.com/topxeq/sqltk"
	"github.com/topxeq/xie"

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

	qlmimemultipart "github.com/topxeq/qlang/lib/mime/multipart"

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

	// qlgithub_360EntSecGroupSkylar_excelize "github.com/topxeq/qlang/lib/github.com/360EntSecGroup-Skylar/excelize"
	qlgithub_xuri_excelize "github.com/topxeq/qlang/lib/github.com/xuri/excelize/v2"

	qlgithub_stretchr_objx "github.com/topxeq/qlang/lib/github.com/stretchr/objx"

	qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi "github.com/topxeq/qlang/lib/github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	qlgithub_topxeq_afero "github.com/topxeq/qlang/lib/github.com/topxeq/afero"

	qlgithub_domodwyer_mailyak "github.com/topxeq/qlang/lib/github.com/domodwyer/mailyak"
	qlgithub_topxeq_docxrepl "github.com/topxeq/qlang/lib/github.com/topxeq/docxrepl"

	qlgithub_topxeq_socks "github.com/topxeq/qlang/lib/github.com/topxeq/socks"

	qlgithub_topxeq_regexpx "github.com/topxeq/qlang/lib/github.com/topxeq/regexpx"

	qlgithub_topxeq_xmlx "github.com/topxeq/qlang/lib/github.com/topxeq/xmlx"

	qlgithub_topxeq_awsapi "github.com/topxeq/qlang/lib/github.com/topxeq/awsapi"

	qlgithub_topxeq_charlang "github.com/topxeq/qlang/lib/github.com/topxeq/charlang"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"github.com/topxeq/tk"
)

var versionG = "v3.8.3"
var VersionG = versionG

var notFoundG = interface{}(errors.New("not found"))

var initFlag bool = false

var scriptPathG = ""

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

func isErrX(vA interface{}) bool {
	if vA == nil {
		return false
	}

	_, ok := vA.(error)

	if ok {
		return true
	}

	nv2, ok := vA.(string)

	if ok {
		return tk.IsErrStr(nv2)
	}

	return false
}

func getErrStrX(vA interface{}) string {
	if vA == nil {
		return ""
	}

	nv1, ok := vA.(error)

	if ok {
		return nv1.Error()
	}

	nv2, ok := vA.(string)

	if ok {
		if tk.IsErrStr(nv2) {
			return tk.GetErrStr(nv2)
		}
	}

	return ""
}

func fnASRSE(fn func(string) (string, error)) func(args ...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		if len(args) != 1 {
			return tk.Errf("not enough parameters")
		}

		s := tk.ToStr(args[0])

		strT, errT := fn(s)

		if errT != nil {
			return errT
		}

		return strT
	}
}

func fnASRSe(fn func(string) string) func(args ...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		if len(args) != 1 {
			return tk.Errf("not enough parameters")
		}

		s := tk.ToStr(args[0])

		strT := fn(s)

		if tk.IsErrStr(strT) {
			return tk.Errf("%v", tk.GetErrStr(strT))
		}

		return strT
	}
}

func fnASSRSe(fn func(string, string) string) func(args ...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		if len(args) < 2 {
			return tk.Errf("not enough parameters")
		}

		s1 := tk.ToStr(args[0])
		s2 := tk.ToStr(args[1])

		strT := fn(s1, s2)

		if tk.IsErrStr(strT) {
			return tk.Errf("%v", tk.GetErrStr(strT))
		}

		return strT
	}
}

func fnASSVRSe(fn func(string, ...string) string) func(args ...interface{}) interface{} {
	return func(args ...interface{}) interface{} {
		lenT := len(args)
		if lenT < 1 {
			return tk.Errf("not enough parameters")
		}

		s1 := tk.ToStr(args[0])

		var strT string

		if lenT < 2 {
			strT = fn(s1)
		} else {
			strT = fn(s1, tk.ObjectsToS(args[1:])...)
		}

		if tk.IsErrStr(strT) {
			return tk.Errf("%v", tk.GetErrStr(strT))
		}

		return strT
	}
}

func newStringBuilder() *strings.Builder {
	return new(strings.Builder)
}

func fromJSONX(vA interface{}) interface{} {
	jsonA, ok := vA.(string)

	if !ok {
		return tk.Errf("string type required")
	}

	rsT, errT := tk.FromJSON(jsonA)

	if errT != nil {
		return errT
	}

	return rsT
}

func getNowDateStrCompact() string {
	return tk.GetNowTimeString()[0:8]
}

func isNil(vA interface{}) bool {
	if vA == nil {
		return true
	}

	if vA == spec.Undefined {
		return true
	}

	return tk.IsNil(vA)
}

func trim(vA interface{}, argsA ...string) string {
	if vA == nil {
		return ""
	}

	if vA == spec.Undefined {
		return ""
	}

	if nv, ok := vA.(string); ok {
		return tk.Trim(nv, argsA...)
	}

	return tk.Trim(fmt.Sprintf("%v", vA), argsA...)
}

func newWebView2() {

}

func initGUI() {

}

func getInputGUI() {

}

func getPasswordGUI() {

}

func getListItemGUI() {

}

func getListItemsGUI() {

}

func getColorGUI() {

}

func getDateGUI() {

}

func getConfirmGUI() {

}

func showInfoGUI() {

}

func showErrorGUI() {

}

func selectFileToSaveGUI() {

}

func selectFileGUI() {

}

func selectDirectoryGUI() {

}

func printValue(nameA string) {
	return
}

func defined(nameA string) bool {
	return false

}

func getStack() string {
	return tk.ErrStrf("not available")
}

func getVars() string {
	return tk.ErrStrf("not available")
}

func typeOfVar(nameA string) string {
	return tk.ErrStrf("not available")
}

var leBufG []string
var leLineEndG string = "\n"
var leSilentG bool = false

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

	return tk.JoinLines(leBufG, leLineEndG)
}

func leLoadFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	strT, errT := tk.LoadStringFromFileE(fileNameA)

	if errT != nil {
		if !leSilentG {
			tk.Pl("failed load file to leBuf: %v", errT)
		}

		return errT
	}

	leBufG = tk.SplitLines(strT)
	// leBufG, errT = tk.LoadStringListBuffered(fileNameA, false, false)

	return nil
}

func leAppendFile(fileNameA string) error {
	if leBufG == nil {
		leClear()
	}

	strT, errT := tk.LoadStringFromFileE(fileNameA)

	if errT != nil {
		if !leSilentG {
			tk.Pl("failed load file to leBuf: %v", errT)
		}

		return errT
	}

	leBufG = append(leBufG, tk.SplitLines(strT)...)
	// leBufG, errT = tk.LoadStringListBuffered(fileNameA, false, false)

	return nil
}

func leLoadUrl(urlA string) error {
	if leBufG == nil {
		leClear()
	}

	strT := tk.DownloadWebPageX(urlA)

	if tk.IsErrStr(strT) {
		if !leSilentG {
			tk.Pl("failed load URL to leBuf: %v", tk.GetErrStr(strT))
		}

		return tk.ErrStrToErr(strT)
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

	textT := tk.JoinLines(leBufG, leLineEndG)

	if tk.IsErrStr(textT) {
		if !leSilentG {
			tk.Pl("failed save leBuf to File: %v", tk.GetErrStr(textT))
		}
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
		if !leSilentG {
			tk.Pl("failed load clipboard to leBuf: %v", tk.GetErrStr(textT))
		}

		return tk.Errf(tk.GetErrStr(textT))
	}

	leBufG = tk.SplitLines(textT)

	return nil
}

func leSaveClip() error {
	if leBufG == nil {
		leClear()
	}

	textT := tk.JoinLines(leBufG, leLineEndG)

	if tk.IsErrStr(textT) {
		if !leSilentG {
			tk.Pl("failed save leBuf to clipboard: %v", tk.GetErrStr(textT))
		}

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
		textT := tk.JoinLines(leBufG, leLineEndG)

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
		if !leSilentG {
			tk.Pl("line index out of range: %v", idxA)
		}

		return tk.Errf("line index out of range")
	}

	tk.Pln(leBufG[idxA])

	return nil
}

func leSort(descentA bool) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	if descentA {
		sort.Sort(sort.Reverse(sort.StringSlice(leBufG)))
	} else {
		sort.Sort(sort.StringSlice(leBufG))
	}

	return nil
}

func leConvertToUTF8(srcEncA ...string) error {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.Errf("buffer not initalized")
	}

	encT := ""

	if len(srcEncA) > 0 {
		encT = srcEncA[0]
	}

	leBufG = tk.SplitLines(tk.ConvertStringToUTF8(tk.JoinLines(leBufG, leLineEndG), encT))

	return nil
}

func leLineEnd(lineEndA ...string) string {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.ErrStrf("buffer not initalized")
	}

	if len(lineEndA) > 0 {
		leLineEndG = lineEndA[0]
	} else {
		return leLineEndG
	}

	return ""
}

func leSilent(silentA ...bool) bool {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return false
	}

	if len(silentA) > 0 {
		leSilentG = silentA[0]
		return leSilentG
	}

	return leSilentG
}

func leGetLine(idxA int) string {
	if leBufG == nil {
		leClear()
	}

	if leBufG == nil {
		return tk.ErrStrf("buffer not initalized")
	}

	if idxA < 0 || idxA >= len(leBufG) {
		if !leSilentG {
			tk.Pl("line index out of range: %v", idxA)
		}

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
		if !leSilentG {
			tk.Pl("line index out of range: %v", idxA)
		}

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
		if !leSilentG {
			tk.Pl("line index out of range: %v", idxA)
		}

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
		if !leSilentG {
			tk.Pl("line index out of range: %v", startA)
		}

		return tk.Errf("start line index out of range")
	}

	if endA < 0 || endA >= len(leBufG) {
		if !leSilentG {
			tk.Pl("line index out of range: %v", endA)
		}

		return tk.Errf("end line index out of range")
	}

	if startA > endA {
		if !leSilentG {
			tk.Pl("start line index greater than end line index: %v", startA)
		}

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

	_, ok := vA.(error)
	if ok {
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

	rsT := fmt.Sprintf("%v", vA)

	if tk.IfSwitchExistsWhole(argsA, "-trim") {
		rsT = tk.Trim(rsT)
	}

	return rsT
}

func nilToEmptyOk(vA interface{}, argsA ...string) (string, bool) {

	if vA == nil {
		return "", true
	}

	if vA == spec.Undefined {
		return "", false
	}

	if tk.IsNil(vA) {
		return "", true
	}

	_, ok := vA.(error)
	if ok {
		return "", true
	}

	if (argsA != nil) && (len(argsA) > 0) {
		vf, ok := vA.(float64)
		if ok {
			if tk.IfSwitchExistsWhole(argsA, "-nofloat") {
				return tk.ToStr(int(vf)), true
			} else {
				return tk.Float64ToStr(vA.(float64)), true
			}
		}

	}

	rsT := fmt.Sprintf("%v", vA)

	if tk.IfSwitchExistsWhole(argsA, "-trim") {
		rsT = tk.Trim(rsT)
	}

	return rsT, true
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

func isValidNotEmpty(vA interface{}, argsA ...string) bool {
	rsT := isValid(vA, argsA...)

	if rsT {
		nv, ok := vA.(string)
		if ok {
			if nv == "" {
				return false
			}
		}
	}

	return rsT
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

func NewFuncInterfaceInterface(funcA interface{}) func(interface{}) interface{} {
	funcT := (funcA).(*execq.Function)
	f := func(s interface{}) interface{} {
		r := funcT.Call(execq.NewStack(), s).([]interface{})

		if len(r) < 1 {
			return nil
		}

		return r[0].(interface{})
	}

	return f
}

func NewFuncInterfacesInterface(funcA interface{}) func(...interface{}) interface{} {
	funcT := (funcA).(*execq.Function)
	f := func(s ...interface{}) interface{} {
		// tk.Pl("x2: %v", s)
		r0 := funcT.Call(execq.NewStack(), s...)

		r, ok := r0.([]interface{})

		if ok {
			if r == nil {
				return nil
			}

			if len(r) < 1 {
				return nil
			}

			rs, ok := r[0].(interface{})
			if ok {
				return rs
			}

			return nil

		} else {
			return r0
		}

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

func isUndefined(vA interface{}) bool {
	if vA == spec.Undefined {
		return true
	}

	return false
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

func getCfgString(fileNameA string) string {
	basePathT, errT := tk.EnsureBasePath("gox")

	if errT == nil {
		cfgPathT := tk.JoinPath(basePathT, fileNameA)

		cfgStrT := tk.Trim(tk.LoadStringFromFile(cfgPathT))

		if !tk.IsErrorString(cfgStrT) {
			return cfgStrT
		}

		return tk.ErrStrF("failed to get config string: %v", tk.GetErrorString(cfgStrT))

	}

	return tk.ErrStrF("failed to get config string: %v", errT)
}

func setCfgString(fileNameA string, strA string) string {
	basePathT, errT := tk.EnsureBasePath("gox")

	if errT == nil {
		cfgPathT := tk.JoinPath(basePathT, fileNameA)

		rsT := tk.SaveStringToFile(strA, cfgPathT)

		if tk.IsErrorString(rsT) {
			return tk.ErrStrF("failed to save config string: %v", tk.GetErrorString(rsT))
		}

		return ""

	}

	return tk.ErrStrF("failed to save config string: %v", errT)
}

func runCode(codeA string, argsA ...string) interface{} {
	if !initFlag {
		initFlag = true
		importQLNonGUIPackages()
	}

	vmT := qlang.New("-noexit")

	// if argsA != nil && len(argsA) > 0 {
	vmT.SetVar("argsG", argsA)
	// } else {
	// 	vmT.SetVar("argsG", os.Args)
	// }

	retG := notFoundG

	errT := vmT.SafeEval(codeA)

	if errT != nil {
		return errT
	}

	rs, ok := vmT.GetVar("outG")

	if ok {
		if rs != nil {
			return rs
		}
	}

	if retG != notFoundG {
		return retG
	}

	return retG
}

func runScript(codeA string, modeA string, argsA ...string) interface{} {

	if true { //modeA == "" || modeA == "0" || modeA == "ql" {
		vmT := qlang.New()

		// if argsA != nil && len(argsA) > 0 {
		vmT.SetVar("argsG", argsA)
		// }

		retG := notFoundG

		errT := vmT.SafeEval(codeA)

		if errT != nil {
			return errT
		}

		rs, ok := vmT.GetVar("outG")

		if ok {
			if rs != nil {
				return rs
			}
		}

		return retG
	} else {
		return tk.SystemCmd("gox", append([]string{codeA}, argsA...)...)
	}

}

func runFile(argsA ...string) interface{} {
	lenT := len(argsA)

	if lenT < 1 {
		return nil
	}

	fcT := tk.LoadStringFromFile(argsA[0])

	if tk.IsErrorString(fcT) {
		return tk.Errf("Invalid file content: %v", tk.GetErrorString(fcT))
	}

	return runScript(fcT, "", argsA[1:]...)
}

func getMagic(numberA int) string {
	if numberA < 0 {
		return tk.GenerateErrorString("invalid magic number")
	}

	typeT := numberA % 10

	var fcT string

	if typeT == 8 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https://gitee.com/topxeq/gox/raw/master/magic/%v.gox", numberA), nil, "", 30)

	} else if typeT == 7 {
		fcT = tk.DownloadPageUTF8(tk.Spr("https: //raw.githubusercontent.com/topxeq/gox/master/magic/%v.gox", numberA), nil, "", 30)
	} else {
		return tk.GenerateErrorString("invalid magic number")
	}

	return fcT

}

func magic(numberA int, argsA ...string) interface{} {
	fcT := getMagic(numberA)

	if tk.IsErrorString(fcT) {
		return tk.ErrorStringToError(fcT)
	}

	return runCode(fcT, argsA...)

}

func newCharFunc(funcA interface{}) *charlang.Function {
	funcT := (funcA).(*execq.Function)
	// f := func(s interface{}) (interface{}, error) {
	// 	r := funcT.Call(execq.NewStack(), s).([]interface{})

	// 	if r[1] == nil {
	// 		return r[0].(interface{}), nil
	// 	}

	// 	return r[0].(interface{}), r[1].(error)
	// }

	// return f

	return &charlang.Function{
		Value: func(argsA ...charlang.Object) (charlang.Object, error) {
			s := make([]interface{}, 0, len(argsA))

			for _, v := range argsA {
				s = append(s, charlang.ConvertFromObject(v))
			}

			r := funcT.Call(execq.NewStack(), s...).([]interface{})

			if r[1] == nil {
				return charlang.ConvertToObject(r[0].(interface{})), nil
			}

			return charlang.ConvertToObject(r[0].(interface{})), charlang.NewCommonError(r[1].(error).Error())
		},
	}
}

func importQLNonGUIPackages() {
	// import native functions and global variables
	var defaultExports = map[string]interface{}{
		// 其中 tk.开头的函数都是github.com/topxeq/tk包中的，可以去pkg.go.dev/github.com/topxeq/tk查看函数定义

		// common related 一般函数
		"defined":         defined,               // 查看某变量是否已经定义，注意参数是字符串类型的变量名，例： if defined("a") {...}
		"pass":            tk.Pass,               // 没有任何操作的函数，一般用于脚本结尾避免脚本返回一个结果导致输出乱了
		"isDefined":       isDefined,             // 判断某变量是否已经定义，与defined的区别是传递的是变量名而不是字符串方式的变量，例： if isDefined(a) {...}
		"isDef":           isDefined,             // 等同于isDef
		"isUndefined":     isUndefined,           // 判断某变量是否未定义
		"isUndef":         isUndefined,           // 等同于isUndefined
		"isNil":           isNil,                 // 判断一个变量或表达式是否为nil
		"isValid":         isValid,               // 判断某变量是否已经定义，并且不是nil，如果传入了第二个参数，还可以判断该变量是否类型是该类型，例： if isValid(a, "string") {...}
		"isValidNotEmpty": isValidNotEmpty,       // 判断某变量是否已经定义，并且不是nil或空字符串，如果传入了第二个参数，还可以判断该变量是否类型是该类型，例： if isValid(a, "string") {...}
		"isValidX":        isValidNotEmpty,       // 等同于isValidNotEmpty
		"eval":            qlEval,                // 运行一段Gox语言代码并获得其返回值，返回值可以放于名为outG的全局变量中，也可以作为最后一个表达式的返回值返回
		"typeOf":          tk.TypeOfValue,        // 给出某变量的类型名
		"typeOfReflect":   tk.TypeOfValueReflect, // 给出某变量的类型名（使用了反射方式）
		"typeOfVar":       typeOfVar,             // 给出某变量的内部类型名，注意参数是字符串类型的变量名
		"exit":            tk.Exit,               // 立即退出脚本的执行，可以带一个整数作为参数，也可以没有
		"setValue":        tk.SetValue,           // 用反射的方式设定一个变量的值
		"getValue":        tk.GetValue,           // 用反射的方式获取一个变量的值
		"getPointer":      tk.GetPointer,         // 用反射的方式获取一个变量的指针
		"getAddr":         tk.GetAddr,            // 用反射的方式获取一个变量的地址
		"setVar":          tk.SetVar,             // 设置一个全局变量，例： setVar("a", "value of a")
		"getVar":          tk.GetVar,             // 获取一个全局变量的值，例： v = getVar("a")
		"ifThenElse":      tk.IfThenElse,         // 相当于三元操作符a?b:c，但注意a、b、c三个表达式仍需语法正确
		"ifElse":          tk.IfThenElse,         // 相当于ifThenElse
		"ifThen":          tk.IfThenElse,         // 相当于ifThenElse
		"deepClone":       tk.DeepClone,
		"deepCopy":        tk.DeepCopyFromTo,
		"run":             runFile,
		"runCode":         runCode, // 运行一段Gox代码（新开一个虚拟机），传入参数除第一个是代码字符串外，后面可以跟多个参数，如果是字符串参数会加入新虚拟机的argsG变量中，如果是字符串数组也会都加入argsG中，如果是映射，则会按照键值加入全局变量中，例如：runCode("pln", {"arg1": 1.2, "arg2": "true"})
		"runScript":       runScript,
		"magic":           magic,

		// debug relate 调试相关
		"dump":   tk.Dump, // 输出一个或多个对象信息供参考
		"dumpf":  tk.Dumpf,
		"sdump":  tk.Sdump, // 生成一个或多个对象信息供参考
		"sdumpf": tk.Sdumpf,

		// output related 输出相关
		"pv":        printValue,   // 输出一个变量的值，注意参数是字符串类型的变量名，例： pv("a")
		"pr":        tk.Pr,        // 等同于其他语言中的print
		"prf":       tk.Printf,    // 等同于其他语言中的printf
		"pln":       tk.Pln,       // 等同于其他语言中的println
		"printfln":  tk.Pl,        // 等同于其他语言中的printf，但多输出一个回车换行
		"pl":        tk.Pl,        // 等同于printfln
		"sprintf":   fmt.Sprintf,  // 等同于其他语言中的sprintf
		"spr":       fmt.Sprintf,  // 等同于sprintf
		"fprintf":   fmt.Fprintf,  // 等同于其他语言中的frintf
		"plv":       tk.Plv,       // 输出某变量或表达式的内容/值，以Go语言内部的表达方式，例如字符串将加上双引号
		"plvx":      tk.Plvx,      // 输出某变量或表达式的内容/值和类型等信息
		"plo":       tk.Plo,       // 输出某变量或表达式的内容/值和类型等信息
		"plos":      tk.Plos,      // 输出多个变量或表达式的内容/值和类型等信息
		"plosr":     tk.Plosr,     // 输出多个变量或表达式的内容/值和类型等信息，每个换行输出
		"plNow":     tk.PlNow,     // 相当于pl，但前面多加了一个时间标记
		"plVerbose": tk.PlVerbose, // 相当于pl，但前面多了一个布尔类型的参数，可以传入一个verbose变量，指定是否输出该信息，例：
		// v = false
		// plVerbose(v, "a: %v", 3) // 由于v的值为false，因此本条语句将不输出
		"vpl":    tk.PlVerbose, // 等同于plVerbose
		"plvsr":  tk.Plvsr,     // 输出多个变量或表达式的值，每行一个
		"plerr":  tk.PlErr,     // 快捷输出一个error类型的值
		"plErr":  tk.PlErr,     // 快捷输出一个error类型的值
		"plErrX": tk.PlErrX,    // 快捷输出一个error类型或TXERROR:开始的字符串的值
		"plExit": tk.PlAndExit, // 相当于pl然后exit退出脚本的执行

		// input related 输入相关
		"getChar":      tk.GetChar,           // 从命令行获取用户的输入，成功返回一个表示字符字符串(控制字符代码+字符代码)，否则返回error对象
		"getChar2":     tk.GetChar2,          // 从命令行获取用户的输入，成功返回一个表示字符ASCII码的字符串，否则返回error对象
		"getInput":     tk.GetUserInput,      // 从命令行获取用户的输入
		"getInputf":    tk.GetInputf,         // 从命令行获取用户的输入，同时可以用printf先输出一个提示信息
		"getPasswordf": tk.GetInputPasswordf, // 从命令行获取密码输入，输入信息将不显示

		// math related数学相关
		"bitXor":       tk.BitXor,               // 异或运算
		"adjustFloat":  tk.AdjustFloat,          // 去除浮点数的计算误差，用法：adjustFloat(4.000000002, 2)，第二个参数表示保留几位小数点后数字
		"getRandomInt": tk.GetRandomIntLessThan, // 获取[0-maxA)之间的随机整数
		"getRandom":    tk.GetRandomFloat,       // 获取[0.0-1.0)之间的随机浮点数

		// string related 字符串相关
		"trim":                 trim,               // 取出字符串前后的空白字符，可选的第二个参数可以是待去掉的字符列表，等同于tk.Trim, 但支持Undefind（转空字符串）和nil
		"strTrim":              tk.Trim,            // 等同于tk.Trim
		"trimSafely":           tk.TrimSafely,      // 取出字符串前后的空白字符，非字符串则返回默认值空，可以通过第二个（可选）参数设置默认值
		"trimx":                tk.TrimSafely,      // 等同于trimSafely
		"trimX":                tk.TrimSafely,      // 等同于trimSafely
		"trimStart":            strings.TrimPrefix, // 去除前导子字符串
		"trimEnd":              strings.TrimSuffix, // 去除末尾子字符串
		"toLower":              strings.ToLower,    // 字符串转小写
		"toUpper":              strings.ToUpper,    // 字符串转大写
		"padStr":               tk.PadString,       // 字符串补零等填充操作，例如 s1 = padStr(s0, 5, "-fill=0", "-right=true")，第二个参数是要补齐到几位，默认填充字符串fill为字符串0，right（表示是否在右侧填充）为false（也可以直接写成-right），因此上例等同于padStr(s0, 5)，如果fill字符串不止一个字符，最终补齐数量不会多于第二个参数指定的值，但有可能少
		"strPad":               tk.PadString,
		"limitStr":             tk.LimitString,            // 超长字符串截短，用法 s2 = limitStr("abcdefg", 3, "-suffix=...")，将得到abc...，suffix默认为...
		"strContains":          strings.Contains,          // 判断字符串中是否包含某个字串
		"strContainsIn":        tk.ContainsIn,             // 判断字符串中是否包含某几个字串
		"strReplace":           tk.Replace,                // 替换字符串中的字串
		"strReplaceIn":         tk.StringReplace,          // strReplaceIn("2020-02-02 08:09:15", "-", "", ":", "", " ", "")
		"strJoin":              strJoin,                   // 连接一个字符串数组，以指定的分隔符，例： s = strJoin(listT, "\n")
		"strSplit":             strings.Split,             // 拆分一个字符串为数组，例： listT = strSplit(strT, "\n")
		"strSplitByLen":        tk.SplitByLen,             // 按长度拆分一个字符串为数组，注意由于是rune，可能不是按字节长度，例： listT = strSplitByLen(strT, 10)，可以加第三个参数表示字节数不能超过多少，加第四个参数表示分隔符（遇上分隔符从分隔符后重新计算长度，也就是说分割长度可以超过指定的个数，一般用于有回车的情况）
		"splitLines":           tk.SplitLines,             // 相当于strSplit(strT, "\n")
		"strSplitLines":        tk.SplitLines,             // 相当于splitLines
		"startsWith":           tk.StartsWith,             // 判断字符串是否以某子串开头
		"strStartsWith":        tk.StartsWith,             // 等同于startsWith
		"endsWith":             tk.EndsWith,               // 判断字符串是否以某子串结尾
		"strEndsWith":          tk.EndsWith,               // 等同于endsWith
		"strIn":                tk.InStrings,              // 判断字符串是否在一个字符串列表中出现，函数定义： strIn(strA string, argsA ...string) bool，第一个可变参数如果以“-”开头，将表示参数开关，-it表示忽略大小写，并且trim再比较（strA并不trim）
		"strFindAll":           tk.FindSubStringAll,       // 寻找字符串中某个子串出现的所有位置，函数定义： func strFindAll(strA string, subStrA string) [][]int，每个匹配是两个整数，分别表示开头和结尾（不包含）
		"newStringBuilder":     newStringBuilder,          // 新建一个strings.Builder对象
		"newStringBuffer":      newStringBuilder,          // 同newStringBuilder
		"getNowStr":            tk.GetNowTimeStringFormal, // 获取一个表示当前时间的字符串，格式：2020-02-02 08:09:15
		"getNowString":         tk.GetNowTimeStringFormal, // 等同于getNowStr
		"getNowStrCompact":     tk.GetNowTimeString,       // 获取一个简化的表示当前时间的字符串，格式：20200202080915
		"getNowStringCompact":  tk.GetNowTimeStringFormal, // 等同于getNowStringCompact
		"getNowDateStrCompact": getNowDateStrCompact,      // 获取一个简化的表示当前日期的字符串，格式：20210215
		"genTimeStamp":         tk.GetTimeStampMid,        // 生成时间戳，格式为13位的Unix时间戳1640133706954，例：timeStampT = genTimeStamp(time.Now())
		"genRandomStr":         tk.GenerateRandomStringX,  // 生成随机字符串，函数定义： genRandomStr("-min=6", "-max=8", "-noUpper", "-noLower", "-noDigit", "-special", "-space", "-invalid")
		"generateRandomString": tk.GenerateRandomString,   // 生成随机字符串，函数定义： (minCharA, maxCharA int, hasUpperA, hasLowerA, hasDigitA, hasSpecialCharA, hasSpaceA bool, hasInvalidChars bool) string

		// regex related 正则表达式相关
		"regMatch":        tk.RegMatchX,          // 判断某字符串是否完整符合某表达式，例： if regMatch(mailT, `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,8})$`) {...}
		"regContains":     tk.RegContainsX,       // 判断某字符串是否包含符合正则表达式的子串，例： if regContains("abccd", "b.c") {...}
		"regContainsIn":   tk.RegContainsIn,      // 判断字符串中是否包含符合正则表达式的某几个字串
		"regCount":        tk.RegCount,           // 判断某字符串包含几个符合正则表达式的子串
		"regFind":         tk.RegFindFirstX,      // 根据正则表达式在字符串中寻找第一个匹配，函数定义： func regFind(strA, patternA string, groupA int) string
		"regFindAll":      tk.RegFindAllX,        // 根据正则表达式在字符串中寻找所有匹配，函数定义： func regFindAll(strA, patternA string, groupA int) []string
		"regFindIndex":    tk.RegFindFirstIndexX, // 根据正则表达式在字符串中第一个匹配的为止，函数定义： func regFindIndex(strA, patternA string) (int, int)
		"regFindAllIndex": tk.RegFindAllIndexX,   // 根据正则表达式搜索在字符串中所有匹配，函数定义： func regFindAllIndex(strA, patternA string) [][]int
		"regReplace":      tk.RegReplaceX,        // 根据正则表达式在字符串中进行替换，函数定义： regReplace(strA, patternA, replaceA string) string, 例：regReplace("abcdefgabcdfg", "(b.*)f(ga.*?)g", "${1}_${2}")，结果是abcde_gabcdf
		"regSplit":        tk.RegSplitX,          // 根据正则表达式分割字符串（以符合条件的匹配来分割），函数定义： regSplit(strA, patternA string, nA ...int) []string

		// conversion related 转换相关
		"nilToEmpty":      nilToEmpty,                     // 将nil、error等值都转换为空字符串，其他的转换为字符串, 加-nofloat参数将浮点数转换为整数，-trim参数将结果trim
		"nilToEmptyOk":    nilToEmptyOk,                   // 将nil、error等值都转换为空字符串，其他的转换为字符串, 加-nofloat参数将浮点数转换为整数，-trim参数将结果trim，第二个返回值是bool类型，如果值是undefined，则返回false，其他情况为true
		"intToStr":        tk.IntToStrX,                   // 整数转字符串
		"strToInt":        tk.StrToIntWithDefaultValue,    // 字符串转整数
		"floatToStr":      tk.Float64ToStr,                // 浮点数转字符串
		"strToFloat":      tk.StrToFloat64,                // 字符串转浮点数，如果第二个参数（可选）存在，则默认错误时返回该值，否则错误时返回-1
		"timeToStr":       tk.FormatTime,                  // 时间转字符串，函数定义: timeToStr(timeA time.Time, formatA ...string) string，formatA可为"2006-01-02 15:04:05"（默认值）等字符串，为compact代表“20060102150405”
		"timeStampToTime": tk.GetTimeFromUnixTimeStampMid, // Unix时间戳转时间（time.Time），支持10位和13位的时间戳，用法: timeT = timeToStr(timeStampToTime("1641139200"), "compact") ，得到20220103000000
		"formatTime":      tk.FormatTime,                  // 等同于timeToStr
		"strToTime":       strToTime,                      // 字符串转时间
		"toTime":          tk.ToTime,                      // 字符串或时间转时间
		"bytesToData":     tk.BytesToData,                 // 字节数组转任意类型变量，可选参数-endian=B或L指定使用BigEndian字节顺序还是LittleEndian，函数定义func(bytesA []byte, dataA interface{}, optsA ...string) error，其中dataA为接收变量
		"dataToBytes":     tk.DataToBytes,                 // 任意类型值转字节数组，可选参数-endian=B或L指定使用BigEndian字节顺序还是LittleEndian
		"toStr":           tk.ToStr,                       // 任意值转字符串
		"toInt":           tk.ToInt,                       // 任意值转整数
		"toFloat":         tk.ToFloat,                     // 任意值转浮点数
		"toByte":          tk.ToByte,                      // 任意值转字节
		"toSimpleMap":     tk.SimpleMapToString,           // 将一个map（map[string]string或map[string]interface{}）转换为Simple Map字符串
		"fromSimpleMap":   tk.LoadSimpleMapFromString,     // 将一个Simple Map字符串转换为map[string]string

		"hexToBytes": tk.HexToBytes, // 将16进制字符串转换为字节数组([]byte)
		"bytesToHex": tk.BytesToHex, // 将字节数组([]byte)转换为16进制字符串
		"hexEncode":  tk.StrToHex,   // 16进制编码
		"hex":        tk.StrToHex,   // 等同于hexEncode
		"strToHex":   tk.StrToHex,   // 等同于hexEncode
		"toHex":      tk.ToHex,      // 将任意值转换为16进制形式，注意是小写格式
		"hexDecode":  tk.HexToStr,   // 16进制解码
		"hexToStr":   tk.HexToStr,   // 等同于hexDecode

		"mssToMsi": tk.MSS2MSI, // 转换map[string]string到map[string]interface{}
		"msiToMss": tk.MSI2MSS, // 转换map[string]interface{}到map[string]string

		"mssToCharMap": charlang.MssToMap, // 转换map[string]string到charlang中的map
		"msiToCharMap": charlang.MsiToMap, // 转换map[string]interface{}到charlang中的map

		"toInterface": tk.ToInterface, // 任意值转interface{}
		"toPointer":   tk.ToPointer,   // 任意值转相应的指针
		"toVar":       tk.ToVar,       // 任意值（*interface{}）转相应的值

		// array/map related 数组（切片）/映射（字典）相关
		"removeItems":   tk.RemoveItemsInArray,               // 从切片中删除指定的项，例： removeItems(aryT, 3, 5)，注意这是表示删除序号为3到5的项目（序号从0开始），共三项
		"removeItem":    tk.RemoveItemsInArray,               // 等同于removeItems
		"remove":        tk.RemoveItemsInArray,               // 等同于removeItems
		"getMapString":  tk.SafelyGetStringForKeyWithDefault, // 从映射中获得指定的键值，避免返回nil，函数定义：func getMapString(mapA map[string]string, keyA string, defaultA ...string) string， 不指定defaultA将返回空字符串
		"getMapItem":    getMapItem,                          // 类似于getMapString，但可以取任意类型的值
		"getArrayItem":  getArrayItem,                        // 类似于getMapItem，但是是取一个切片中指定序号的值
		"joinList":      tk.JoinList,                         // 类似于strJoin，但可以连接任意类型的值
		"arrayContains": tk.ArrayContains,                    // 判断数组中是否包含某值

		// object related 对象有关

		"newObject": tk.NewObject, // 新建一个对象，目前支持stack, set(hashset), treeset, list(arraylist), linklist(linkedlist), tree(btree), stringBuffer(stringBuilder), bytesBuffer, error(err), errorString(errStr), string(TXString), StringRing等，用法：objT = newObject("stack")或objT = newObject("tree", 5)创建五层的btree树等
		"newObj":    tk.NewObject, // 等同于newObject

		// error related 错误处理相关
		"isError":          tk.IsError,            // 判断表达式的值是否为error类型
		"isErr":            tk.IsError,            // 等同于isError
		"isErrX":           tk.IsErrX,             // 判断表达式的值是否为error类型，同时也判断是否是TXERROR:开始的字符串
		"isErrStr":         tk.IsErrStr,           // 判断字符串是否是TXERROR:开始的字符串
		"checkError":       tk.CheckError,         // 检查变量，如果是error则立即停止脚本的执行
		"checkErr":         tk.CheckError,         // 等同于checkError
		"checkErrX":        tk.CheckErrX,          // 检查变量，如果是error或TXERROR:开始的字符串则输出相应信息后立即停止脚本的执行
		"checkErrf":        tk.CheckErrf,          // 检查变量，如果是error则立即停止脚本的执行，之前可以printfln输出信息
		"checkErrorString": tk.CheckErrorString,   // 检查变量，如果是TXERROR:开始的字符串则立即停止脚本的执行
		"checkErrStr":      tk.CheckErrStr,        // 等同于checkErrorString
		"checkErrStrf":     tk.CheckErrStrf,       // 检查变量，如果是TXERROR:开始的字符串则立即停止脚本的执行，之前可以printfln输出信息
		"fatalf":           tk.Fatalf,             // printfln输出信息后终止脚本的执行
		"fatalfc":          tk.FatalfByCondition,  // printfln输出信息后如果第一个参数为false，才终止脚本的执行
		"fatalfi":          tk.FatalfByCondition,  // 同fatalfc
		"errStr":           tk.ErrStr,             // 生成TXERROR:开始的字符串
		"errStrf":          tk.ErrStrF,            // 生成TXERROR:开始的字符串，类似sprintf的用法
		"getErrStr":        tk.GetErrStr,          // 从TXERROR:开始的字符串获取其后的错误信息
		"getErrStrX":       tk.GetErrStrX,         // 从error对象或TXERROR:开始的字符串获取其中的错误信息，返回为空字符串一般表示没有错误
		"errf":             tk.Errf,               // 生成error类型的变量，其中提示信息类似sprintf的用法
		"errToEmptyStr":    tk.ErrorToEmptyString, // 将任意值转为string，如果是error类型的变量则转为空字符串

		// encode/decode related 编码/解码相关
		"xmlEncode":          tk.EncodeToXMLString,    // 编码为XML
		"xmlDecode":          tk.FromXMLWithDefault,   // 解码XML为对象，函数定义：(xmlA string, defaultA interface{}) interface{}
		"fromXML":            tk.FromXMLX,             // 解码XML为etree.Element对象，函数定义：fromXML(xmlA string, pathA ...interface{}) interface{}，出错时返回error，否则返回*etree.Element对象
		"fromXml":            tk.FromXMLX,             // 等同于fromXML
		"toXML":              tk.ToXML,                // 编码数据为XML格式，可选参数-indent, -cdata, -root=ABC, -rootAttr={"f1", "v1"}, -default="<xml>ab c</xml>"
		"toXml":              tk.ToXML,                // 等同于toXML
		"htmlEncode":         tk.EncodeHTML,           // HTML编码（&nbsp;等）
		"htmlDecode":         tk.DecodeHTML,           // HTML解码
		"urlEncode":          tk.UrlEncode2,           // URL编码（http://www.aaa.com -> http%3A%2F%2Fwww.aaa.com）
		"urlEncodeX":         tk.UrlEncode,            // 增强URL编码（会将+和\n等也编码）
		"urlDecode":          tk.UrlDecode,            // URL解码
		"base64Encode":       tk.EncodeToBase64,       // Base64编码，输入参数是[]byte字节数组
		"toBase64":           tk.ToBase64,             // Base64编码，输入参数是[]byte字节数组或字符串
		"base64Decode":       tk.DecodeFromBase64,     // base64解码，返回两个参数（第二个是error类型）的结果
		"fromBase64":         tk.FromBase64,           // base64解码，返回字节数组
		"md5Encode":          tk.MD5Encrypt,           // MD5编码
		"md5":                tk.MD5Encrypt,           // 等同于md5Encode
		"jsonEncode":         tk.ObjectToJSON,         // JSON编码
		"jsonDecode":         tk.JSONToObject,         // JSON解码
		"toJSON":             tk.ToJSONX,              // 增强的JSON编码，建议使用，函数定义： toJSON(objA interface{}, optsA ...string) string，参数optsA可选。例：s = toJSON(textA, "-indent", "-sort")
		"toJson":             tk.ToJSONX,              // 等同于toJSON
		"toJSONX":            tk.ToJSONX,              // 等同于toJSON
		"toJsonX":            tk.ToJSONX,              // 等同于toJSON
		"fromJSON":           tk.FromJSONWithDefault,  // 增强的JSON解码，函数定义： fromJSON(jsonA string, defaultA ...interface{}) interface{}
		"fromJson":           tk.FromJSONWithDefault,  // 等同于fromJSON
		"fromJSONX":          fromJSONX,               // 增强的JSON解码，建议使用，函数定义： fromJSON(jsonA string) interface{}，如果解码失败，返回error对象
		"fromJsonX":          fromJSONX,               // 等同于fromJSONX
		"getJSONNode":        tk.GetJSONNode,          // 获取JSON中的某个节点，未取到则返回nil，示例： getJSONNode("{\"ID\":1,\"Name\":\"Reds\",\"Colors\":[\"Crimson\",\"Red\",\"Ruby\",\"Maroon\"]}", "Colors", 0)
		"getJsonNode":        tk.GetJSONNode,          // 等同于getJSONNode
		"simpleEncode":       tk.EncodeStringCustomEx, // 简单编码，主要为了文件名和网址名不含非法字符
		"simpleDecode":       tk.DecodeStringCustom,   // 简单编码的解码，主要为了文件名和网址名不含非法字符
		"tableToMSSArray":    tk.TableToMSSArray,      // 参见dbRecsToMapArray，主要用于处理数据库查询结果
		"tableToMSSMap":      tk.TableToMSSMap,        // 类似tableToMSSArray，但还加上一个ID作为主键成为字典/映射类型
		"tableToMSSMapArray": tk.TableToMSSMapArray,   // 类似tableToMSSMap，但主键下的键值是一个数组，其中每一项是一个map[string]string

		// encrypt/decrypt related 加密/解密相关
		"encryptStr":   tk.EncryptStringByTXDEF,            // 加密字符串，第二个参数（可选）是密钥字串
		"encryptText":  tk.EncryptStringByTXDEF,            // 等同于encryptStr
		"encryptTextX": fnASSVRSe(tk.EncryptStringByTXDEF), // 等同于encryptStr，但错误时返回error对象
		"decryptStr":   tk.DecryptStringByTXDEF,            // 解密字符串，第二个参数（可选）是密钥字串
		"decryptText":  tk.DecryptStringByTXDEF,            // 等同于decryptStr
		"decryptTextX": fnASSVRSe(tk.DecryptStringByTXDEF), // 等同于decryptStr，但错误时返回error对象
		"encryptData":  tk.EncryptDataByTXDEF,              // 加密二进制数据（[]byte类型），第二个参数（可选）是密钥字串
		"decryptData":  tk.DecryptDataByTXDEF,              // 解密二进制数据（[]byte类型），第二个参数（可选）是密钥字串

		// log related 日志相关
		"setLogFile": tk.SetLogFile,         // 设置日志文件路径，下面有关日志的函数将用到
		"logf":       tk.LogWithTimeCompact, // 输出到日志文件，函数定义： func logf(formatA string, argsA ...interface{})
		"logPrint":   logPrint,              // 同时输出到标准输出和日志文件

		// number related 数字相关
		"abs": tk.Abs,

		// system related 系统相关
		"getClipText":       tk.GetClipText,                 // 从系统剪贴板获取文本，例： textT = getClipText()
		"setClipText":       tk.SetClipText,                 // 设定系统剪贴板中的文本，例： setClipText("测试")
		"getEnv":            tk.GetEnv,                      // 获取系统环境变量
		"setEnv":            tk.SetEnv,                      // 设定系统环境变量
		"systemCmd":         tk.SystemCmd,                   // 执行一条系统命令，例如： systemCmd("cmd", "/k", "copy a.txt b.txt")
		"openFile":          tk.RunWinFileWithSystemDefault, // 用系统默认的方式打开一个文件，例如： openFile("a.jpg")
		"ifFileExists":      tk.IfFileExists,                // 判断文件是否存在
		"fileExists":        tk.IfFileExists,                // 等同于ifFileExists
		"joinPath":          filepath.Join,                  // 连接文件路径，等同于Go语言标准库中的path/filepath.Join
		"ensureDir":         tk.EnsureMakeDirs,
		"getFileSize":       tk.GetFileSizeCompact,           // 获取文件大小
		"getFileInfo":       tk.GetFileInfo,                  // 获取文件信息，返回map[string]string
		"getFileList":       tk.GetFileList,                  // 获取指定目录下的符合条件的所有文件，例：listT = getFileList(pathT, "-recursive", "-pattern=*", "-exclusive=*.txt", "-withDir", "-verbose"), -compact 参数将只给出Abs、Size、IsDir三项, -dirOnly参数将只列出目录（不包含文件）
		"getFiles":          tk.GetFileList,                  // 等同于getFileList
		"createFile":        tk.CreateFile,                   // 等同于tk.CreateFile
		"createTempFile":    tk.CreateTempFile,               // 等同于tk.CreateTempFile
		"copyFile":          tk.CopyFile,                     // 等同于tk.CopyFile，可带参数-force和-bufferSize=100000
		"removeFile":        tk.RemoveFile,                   // 等同于tk.RemoveFile
		"renameFile":        tk.RenameFile,                   // 等同于tk.RenameFile
		"loadText":          tk.LoadStringFromFile,           // 从文件中读取文本字符串，函数定义：func loadText(fileNameA string) string，出错时返回TXERROR:开头的字符串指明原因
		"loadTextX":         fnASRSE(tk.LoadStringFromFileE), // 从文件中读取文本字符串，函数定义：func loadText(fileNameA string) string，出错时返回error对象
		"saveText":          tk.SaveStringToFile,             // 将字符串保存到文件，函数定义： func saveText(strA string, fileA string) string
		"saveTextX":         fnASSRSe(tk.SaveStringToFile),   // 将字符串保存到文件，如果失败返回error对象
		"appendText":        tk.AppendStringToFile,           // 将字符串增加到文件末尾，函数定义： func appendText(strA string, fileA string) string
		"appendTextX":       fnASSRSe(tk.AppendStringToFile), // 将字符串增加到文件末尾，如果失败返回error对象
		"loadBytes":         tk.LoadBytesFromFile,            // 从文件中读取二进制数据，函数定义：func loadBytes(fileNameA string, numA ...int) interface{}，返回[]byte或error，第二个参数没有或者小于零的话表示读取所有
		"loadBytesX":        tk.LoadBytesFromFile,            // 等同于loadBytes
		"saveBytes":         tk.SaveBytesToFileE,             // 将二进制数据保存到文件，函数定义： func saveBytes(bytesA []byte, fileA string) error
		"saveBytesX":        tk.SaveBytesToFileE,             // 等同于saveBytes
		"sleep":             tk.Sleep,                        // 休眠指定的秒数，例：sleep(30)，可以是小数
		"sleepSeconds":      tk.SleepSeconds,                 // 基本等同于sleep，但只能是整数秒
		"sleepMilliSeconds": tk.SleepMilliSeconds,            // 类似于sleep，但单位是毫秒
		"sleepMS":           tk.SleepMilliSeconds,            // 等同于sleepMilliSeconds

		"getAppDir":    tk.GetApplicationPath,
		"getCurDir":    tk.GetCurrentDir,
		"getConfigDir": fnASRSE(tk.EnsureBasePath),

		// time related 时间相关

		"now": time.Now, // 获取当前时间

		// command-line 命令行处理相关
		"getParameter":   tk.GetParameterByIndexWithDefaultValue, // 按顺序序号获取命令行参数，其中0代表第一个参数，也就是软件名称或者命令名称，1开始才是第一个参数，注意参数不包括开关，即类似-verbose=true这样的，函数定义：func getParameter(argsA []string, idxA int, defaultA string) string
		"getParam":       tk.GetParam,                            // 类似于getParameter，只是后两个参数都是可选，默认是1和""（空字符串），且顺序随意
		"getSwitch":      tk.GetSwitchWithDefaultValue,           // 获取命令行参数中的开关，用法：tmps = getSwitch(args, "-verbose=", "false")，第三个参数是默认值（如果在命令行中没取到的话返回该值）
		"getIntSwitch":   tk.GetSwitchWithDefaultIntValue,        // 与getSwitch类似，但获取到的是整型（int）的值
		"getFloatSwitch": tk.GetSwitchWithDefaultFloatValue,      // 与getSwitch类似，但获取到的是浮点数（float64）的值
		"switchExists":   tk.IfSwitchExistsWhole,                 // 判断命令行参数中是否存在开关（完整的，），用法：flag = switchExists(args, "-restart")
		"ifSwitchExists": tk.IfSwitchExistsWhole,                 // 等同于switchExists
		"parseCommand":   tk.ParseCommandLine,                    // 等同于tk.ParseCommandLine

		// network related 网络相关
		"newSSHClient": tk.NewSSHClient, // 新建一个SSH连接，以便执行各种SSH操作，例：
		// clientT, errT = newSSHClient(hostName, port, userName, password)
		// defer clientT.Close() // 别忘了用完关闭网络连接

		// outT, errT = clientT.Run(`ls -p; cat abc.txt`) // 执行一条或多条命令
		// errT = clientT.Upload(`./abc.txt`, strReplace(joinPath(pathT, `abc.txt`), `\`, "/")) // 上传文件
		// errT = clientT.Download(`down.txt`, `./down.txt`) // 下载文件
		// bytesT, errT = clientT.GetFileContent(`/root/test/down.txt`) // 获取某个文件的二进制内容[]byte
		"mapToPostData": tk.MapToPostData,    // 从一个映射（map）对象生成进行POST请求的参数对象，函数定义func mapToPostData(postDataA map[string]string) url.Values
		"getWebPage":    tk.DownloadPageUTF8, // 进行一个网络HTTP请求并获得服务器返回结果，或者下载一个网页，函数定义func getWebPage(urlA string, postDataA url.Values, customHeaders string, timeoutSecsA time.Duration, optsA ...string) string
		// customHeadersA 是自定义请求头，内容是多行文本形如 charset: utf-8。如果冒号后还有冒号，要替换成`
		// 返回结果是TXERROR字符串，即如果是以TXERROR:开头，则表示错误信息，否则是网页或请求响应
		"getWeb": tk.DownloadWebPageX, // 进行一个网络HTTP请求并获得服务器返回结果，或者下载一个网页，函数定义func getWebPage(urlA string, postDataA map[string]string, optsA ...string) string
		// 除了urlA，所有参数都是可选；
		// optsA支持-verbose， -detail， -timeout=30（秒），-encoding=utf-8/gb2312/gbk/gb18030等,
		// 如果要添加FORM形式的POST的数据，则直接传入一个url.Values类型的数据，或者map[string]string或者map[string]interface{}的参数即可，也可以用开关参数-post={"Key1": "Value1", "Key2": "Value2"}这样传入JSON，此时请求将自动转为POST方式（默认是GET方式）
		// 如果要直接POST数据，则直接传入-postBody=ABCDEFG这样的信息即可，其中ABCDEFG是所需POST的字符串，例如getWeb("http://abc.com:8001/sap/bc/srt/rfc/sap/getSvc", "-postBody=<XML><data1>Test</data1></XML>", `-headers={"Content-Type":"text/xml; charset=utf-8", "SOAPAction":""}`, "-timeout=15")，此时请求将自动转为POST方式（默认是GET方式），另外也可以直接传入一个[]byte类型的参数
		// 如需添加自定义请求头，则添加开关参数类似：-headers={"content-type": "text/plain; charset=utf-8;"}
		// 返回结果是TXERROR字符串，即如果是以TXERROR:开头，则表示错误信息，否则是网页或请求响应
		"downloadFile": tk.DownloadFile, // 从网络下载一个文件，函数定义func downloadFile(urlA, dirA, fileNameA string, argsA ...string) string
		"httpRequest":  tk.RequestX,     // 进行一个网络HTTP请求并获得服务器返回结果，函数定义func httpRequest(urlA, methodA, reqBodyA string, customHeadersA string, timeoutSecsA time.Duration, optsA ...string) (string, error)
		// 其中methodA可以是"GET"，"POST"等
		// customHeadersA 是自定义请求头，内容是多行文本形如 charset: utf-8。如果冒号后还有冒号，要替换成`
		"postRequest": tk.PostRequestX, // 进行一个POST网络请求并获得服务器返回结果，函数定义func postRequest(urlA, reqBodyA string, customHeadersA string, timeoutSecsA time.Duration, optsA ...string) (string, error)
		// 其中reqBodyA是POST的body
		// customHeadersA 是自定义请求头，内容是多行文本形如 charset: utf-8。如果冒号后还有冒号，要替换成`
		// timeoutSecsA是请求超时的秒数
		// optsA是一组字符串，可以是-verbose和-detail，均表示是否输出某些信息
		"getFormValue":         tk.GetFormValueWithDefaultValue,  // 从HTTP请求中获取字段参数，可以是Query参数，也可以是POST参数，函数定义func getFormValue(reqA *http.Request, keyA string, defaultA string) string
		"formValueExist":       tk.IfFormValueExists,             // 判断HTTP请求中的是否有某个字段参数，函数定义func formValueExist(reqA *http.Request, keyA string) bool
		"ifFormValueExist":     tk.IfFormValueExists,             // 等同于formValueExist
		"formToMap":            tk.FormToMap,                     // 将HTTP请求中的form内容转换为map（字典/映射类型），例：mapT = formToMap(req.Form)
		"generateJSONResponse": tk.GenerateJSONPResponseWithMore, // 生成Web API服务器的JSON响应，支持JSONP，例：return generateJSONResponse("fail", sprintf("数据库操作失败：%v", errT), req)
		"genResp":              tk.GenerateJSONPResponseWithMore, // 等同于generateJSONResponse
		"writeResp":            tk.WriteResponse,                 // 写http输出，函数原型writeResp(resA http.ResponseWriter, strA string) error
		"writeRespHeader":      tk.WriteResponseHeader,           // 写http响应头的状态（200、404等），函数原型writeRespHeader(resA http.ResponseWriter, argsA ...interface{}) error，例：writeRespHeader(http.StatusOK)
		"setRespHeader":        tk.SetResponseHeader,             // 设置http响应头中的内容，函数原型setRespHeader(resA http.ResponseWriter, keyA string, valueA string) error，例：setRespHeader(responseG, "Content-Type", "text/json; charset=utf-8")
		"jsonRespToHtml":       tk.JSONResponseToHTML,            // 类似{"Status":"fail", "Value":"failed to connect DB"}的JSON响应转换为通用的简单的错误网页
		"getMimeType":          tk.GetMimeTypeByExt,              // 根据文件扩展名获取MIME类型

		"replaceHtmlByMap":      tk.ReplaceHtmlByMap,
		"cleanHtmlPlaceholders": tk.CleanHtmlPlaceholders,

		// database related
		"dbConnect": sqltk.ConnectDBX, // 连接数据库以便后续读写操作，例：
		// dbT = dbConnect("sqlserver", "server=127.0.0.1;port=1443;portNumber=1443;user id=user;password=userpass;database=db1")
		// 	if isError(dbT) {
		// 		fatalf("打开数据库%v错误：%v", dbT)
		// 	}
		// }
		// defer dbT.Close()

		"dbClose": sqltk.CloseDBX, // 关闭数据库连接，例：
		// errT := dbClose(dbT)
		// 	if isError(rs) {
		// 		fatalf("关闭数据库时发生错误：%v", rs)
		// 	}
		// }

		"dbExec": sqltk.ExecDBX, // 进行数据库操作，例：
		// rs := dbExec(dbT, `insert into table1 (field1,id,field2) values('value1',1,'value2')`
		// 	if isError(rs) {
		// 		fatalf("新增数据库记录时发生错误：%v", rs)
		// 	}
		// }
		// insertID, affectedRows = rs[0], rs[1]

		"dbQuery": sqltk.QueryDBX, // 进行数据库查询，所有字段结果都将转换为字符串，返回结果为[]map[string]string，用JSON格式表达类似：[{"Field1": "Value1", "Field2": "Value2"},{"Field1": "Value1a", "Field2": "Value2a"}]，例：
		// sqlRsT = dbQuery(dbT, `SELECT * FROM TABLE1 WHERE ID=3`)
		// if isError(sqlRsT) {
		//		fatalf("查询数据库错误：%v", dbT)
		//	}
		// pl("在数据库中找到%v条记录", len(sqlRsT))

		"dbQueryRecs": sqltk.QueryDBRecsX, // 进行数据库查询，所有字段结果都将转换为字符串，返回结果为[][]string，即二维数组，其中第一行为表头字段名：[["Field1", "Field2"],["Value1","Value2"]]，例：
		// sqlRsT = dbQueryRecs(dbT, `SELECT * FROM TABLE1 WHERE ID=3`)
		// if isErr(sqlRsT) {
		//		fatalf("查询数据库错误：%v", sqlRsT)
		//	}
		// pl("在数据库中找到%v条记录", len(sqlRsT))

		"dbQueryMap": sqltk.QueryDBMapX, // 进行数据库查询，所有字段结果都将转换为字符串，返回结果为map[string]map[string]string，即将dbQuery的结果再加上一个索引，例：{"Value1": {"Field1": "Value1"}, "Value2": {"Field2": "Value2"}}
		// sqlRsT = dbQueryMap(dbT, `SELECT * FROM TABLE1 WHERE ID=3`, "ID")
		// if isErr(sqlRsT) {
		//		fatalf("查询数据库错误：%v", sqlRsT)
		//	}
		// pl("在数据库中找到结果：%v", sqlRsT)

		"dbQueryMapArray": sqltk.QueryDBMapArrayX, // 进行数据库查询，所有字段结果都将转换为字符串，返回结果为map[string][]map[string]string，即将dbQueryMap的结果中，每一个键值中可以是一个数组（[]map[string]string类型），例：{"Value1": [{"Field1": "Value1"}, {"Field1": "Value1a"}], "Value2": [{"Field1": "Value2"}, {"Field1": "Value2a"}, {"Field1": "Value2b"}]}
		// sqlRsT = dbQueryMapArray(dbT, `SELECT * FROM TABLE1 WHERE ID=3`, "ID")
		// if isErr(sqlRsT) {
		//		fatalf("查询数据库错误：%v", sqlRsT)
		//	}
		// pl("在数据库中找到结果：%v", sqlRsT)

		"dbQueryCount": sqltk.QueryCountX, // 与dbQuery类似，但主要进行数量查询，也支持结果只有一个整数的查询，例：
		// sqlRsT = dbQueryCount(dbT, `SELECT COUNT(*) FROM TABLE1 WHERE ID>3`)
		// if isError(sqlRsT) {
		//		fatalf("查询数据库错误：%v", dbT)
		//	}
		// pl("在数据库中共有符合条件的%v条记录", sqlRsT)

		"dbQueryFloat": sqltk.QueryFloatX, // 与dbQueryCount类似，但主要进行返回一个浮点数结果的查询，例：
		// sqlRsT = dbQueryFloat(dbT, `SELECT PRICE FROM TABLE1 WHERE ID=3`)
		// if isError(sqlRsT) {
		//		fatalf("查询数据库错误：%v", dbT)
		//	}
		// pl("查询结果为%v", sqlRsT)

		"dbQueryString": sqltk.QueryStringX, // 与dbQueryCount类似，但主要支持结果只有一个字符串的查询

		"dbFormat":       sqltk.FormatSQLValue, // 将字符串转换为可用在SQL语句中的字符串（将单引号变成双单引号）
		"formatSQLValue": sqltk.FormatSQLValue, // 将字符串转换为可用在SQL语句中的字符串（将单引号变成双单引号）

		"dbOneLineRecordToMap": sqltk.OneLineRecordToMap, // 将只有一行（加标题行两行）的SQL语句查询结果（[][]string格式）变为类似{"Field1": "Value1", "Field2": "Value2"}的map[string]string格式

		"dbOneColumnRecordsToArray": sqltk.OneColumnRecordsToArray, // 将只有一列的SQL语句查询结果（[][]string格式）变为类似["Value1", "Value2"]的[]string格式

		"dbRecsToMapArray": sqltk.RecordsToMapArray, // 将多行行（第一行为标头字段行）的SQL语句查询结果（[][]string格式）变为类似[{"Field1": "Value1", "Field2": "Value2"},{"Field1": "Value1a", "Field2": "Value2a"}]的[]map[string]string格式

		"dbRecsToMapArrayMap": sqltk.RecordsToMapArrayMap, // 将多行行（第一行为标头字段行）的SQL语句查询结果（[][]string格式）变为类似dbQueryMapArray函数返回的结果

		// line editor related 内置行文本编辑器有关
		"leClear":       leClear,         // 清空行文本编辑器缓冲区，例：leClear()
		"leLoadStr":     leLoadString,    // 行文本编辑器缓冲区载入指定字符串内容，例：leLoadStr("abc\nbbb\n结束")
		"leSetAll":      leLoadString,    // 等同于leLoadString
		"leSaveStr":     leSaveString,    // 取出行文本编辑器缓冲区中内容，例：s = leSaveStr()
		"leGetAll":      leSaveString,    // 等同于leSaveStr
		"leLoad":        leLoadFile,      // 从文件中载入文本到行文本编辑器缓冲区中，例：err = leLoad(`c:\test.txt`)
		"leLoadFile":    leLoadFile,      // 等同于leLoad
		"leAppendFile":  leAppendFile,    // 从文件中载入文本追加到行文本编辑器缓冲区中，例：err = leAppendFile(`c:\test.txt`)
		"leSave":        leSaveFile,      // 将行文本编辑器缓冲区中内容保存到文件中，例：err = leSave(`c:\test.txt`)
		"leSaveFile":    leSaveFile,      // 等同于leSave
		"leLoadClip":    leLoadClip,      // 从剪贴板中载入文本到行文本编辑器缓冲区中，例：err = leLoadClip()
		"leSaveClip":    leSaveClip,      // 将行文本编辑器缓冲区中内容保存到剪贴板中，例：err = leSaveClip()
		"leLoadUrl":     leLoadUrl,       // 从网址URL载入文本到行文本编辑器缓冲区中，例：err = leLoadUrl(`http://example.com/abc.txt`)
		"leInsert":      leInsertLine,    // 行文本编辑器缓冲区中的指定位置前插入指定内容，例：err = leInsert(3， "abc")
		"leInsertLine":  leInsertLine,    // 等同于leInsert
		"leAppend":      leAppendLine,    // 行文本编辑器缓冲区中的最后追加指定内容，例：err = leAppendLine("abc")
		"leAppendLine":  leAppendLine,    // 等同于leAppend
		"leSet":         leSetLine,       // 设定行文本编辑器缓冲区中的指定行为指定内容，例：err = leSet(3， "abc")
		"leSetLine":     leSetLine,       // 等同于leSet
		"leSetLines":    leSetLines,      // 设定行文本编辑器缓冲区中指定范围的多行为指定内容，例：err = leSetLines(3, 5， "abc\nbbb")
		"leRemove":      leRemoveLine,    // 删除行文本编辑器缓冲区中的指定行，例：err = leRemove(3)
		"leRemoveLine":  leRemoveLine,    // 等同于leRemove
		"leRemoveLines": leRemoveLines,   // 删除行文本编辑器缓冲区中指定范围的多行，例：err = leRemoveLines(1, 3)
		"leViewAll":     leViewAll,       // 查看行文本编辑器缓冲区中的所有内容，例：allText = leViewAll()
		"leView":        leViewLine,      // 查看行文本编辑器缓冲区中的指定行，例：lineText = leView(18)
		"leSort":        leSort,          // 将行文本编辑器缓冲区中的行进行排序，唯一参数（可省略，默认为false）表示是否降序排序，例：errT = leSort(true)
		"leEnc":         leConvertToUTF8, // 将行文本编辑器缓冲区中的文本转换为UTF-8编码，如果不指定原始编码则默认为GB18030编码
		"leLineEnd":     leLineEnd,       // 读取或设置行文本编辑器缓冲区中行末字符（一般是\n或\r\n），不带参数是获取，带参数是设置
		"leSilent":      leSilent,        // 读取或设置行文本编辑器的静默模式（布尔值），不带参数是获取，带参数是设置

		// GUI related start
		// gui related 图形界面相关
		"initGUI":             initGUI,             // GUI操作，一般均需调用initGUI来进行初始化，例：initGUI()
		"getConfirmGUI":       getConfirmGUI,       // 显示一个提示信息并让用户确认的对话框，例：getConfirmGUI("对话框标题", "信息内容")，注意，从第二个参数开始可以类似于printf那样带格式化字符串和任意长度参数值，例如getConfirmGUI("对话框标题", "信息内容=%v", abc)
		"getInputGUI":         getInputGUI,         // 显示一个提示信息并让用户输入信息的对话框，例：getInputGUI("请输入……", "姓名")，注意，从第3个参数开始为可选参数，可以有-ok=确认按钮标题，-cancel=取消按钮标题，分别表示确认按钮与取消按钮的标题（默认分别为OK和Cancel），例如getInputGUI("对话框标题", "信息内容", "-ok=确定", "-cancel=关闭")，返回输入字符串，如果按了取消按钮，将返回TXERROR:开始的空字符串
		"getPasswordGUI":      getPasswordGUI,      // 显示一个提示信息并让用户输入密码/口令的对话框，例：getPasswordGUI("请输入……", "密码")，注意，从第3个参数开始为可选参数，可以有-ok=确认按钮标题，-cancel=取消按钮标题，分别表示确认按钮与取消按钮的标题（默认分别为OK和Cancel），例如getPasswordGUI("对话框标题", "信息内容", "-ok=确定", "-cancel=关闭")
		"getListItemGUI":      getListItemGUI,      // 提供单选列表供用户选择，结果格式是选中的字符串或者TXERROR字符串；示例：getListItemGUI("请选择", "所需的颜色", ["红色","黄色"]...)
		"getListItemsGUI":     getListItemsGUI,     // 提供多选列表供用户选择，结果格式是选中的字符串数组或者TXERROR字符串；示例：getListItemGUI("请选择", "所需的颜色", ["红色","黄色","蓝色"]...)
		"getColorGUI":         getColorGUI,         // 获取用户选择的颜色，结果格式是FFEEDD或者TXERROR字符串；示例：getColorGUI("请选择颜色", "CCCCCC")
		"getDateGUI":          getDateGUI,          // 获取用户选择的日期，结果格式是20210218或者TXERROR字符串；示例：getDateGUI("请选择……", "开始日期")，注意，从第二个参数开始可以类似于printf那样带格式化字符串和任意长度参数值，例如getPasswordGUI("对话框标题", "信息内容=%v", abc)
		"showInfoGUI":         showInfoGUI,         // 显示一个提示信息的对话框，例：showInfoGUI("对话框标题", "信息内容")，注意，从第二个参数开始可以类似于printf那样带格式化字符串和任意长度参数值，例如showInfoGUI("对话框标题", "信息内容=%v", abc)
		"showErrorGUI":        showErrorGUI,        // 显示一个错误或警告信息的对话框，例：showErrorGUI("对话框标题", "错误或警告内容")，注意，从第二个参数开始可以类似于printf那样带格式化字符串和任意长度参数值，例如showErrorGUI("对话框标题", "信息内容=%v", abc)
		"selectFileToSaveGUI": selectFileToSaveGUI, // 图形化选取用于保存数据的文件，例：fileName = selectFileToSaveGUI("-title=请选择文件……", "-filterName=所有文件", "-filter=*", "-start=.")，参数均为可选，start是默认起始目录
		"selectFileGUI":       selectFileGUI,       // 图形化选取文件，例：fileName = selectFileGUI("-title=请选择文件……", "-filterName=所有文件", "-filter=*", "-start=.")，参数均为可选，start是默认起始目录
		"selectDirectoryGUI":  selectDirectoryGUI,  // 图形化选取目录，例：dirName = selectDirectoryGUI("-title=请选择目录……", "-start=.")，参数均为可选，start是默认起始目录

		"newWebView2": newWebView2, // 新建一个WebView2的窗口

		// GUI related end

		// compress/uncompress related 压缩解压缩相关函数

		"compress":   tk.Compress,
		"uncompress": tk.Uncompress,

		"compressText":   tk.CompressText,
		"uncompressText": tk.UncompressText,

		// misc related 杂项相关函数
		"dealRef": tk.DealRef,

		"lockN":    tk.LockN, // lock a global lock, 0 <= N < 10
		"unlockN":  tk.UnlockN,
		"tryLockN": tk.TryLockN,

		"readLockN":    tk.RLockN, // read lock a global lock, 0 <= N < 10
		"readUnlockN":  tk.RUnlockN,
		"tryReadLockN": tk.TryRLockN,

		"sortX":            tk.SortX,                        // 排序各种数据，用法：sort([{"f1": 1}, {"f1": 2}], "-key=f1", "-desc")
		"newFunc":          NewFuncB,                        // 将Gox语言中的定义的函数转换为Go语言中类似 func f() 的形式
		"newFuncII":        NewFuncInterfaceInterface,       // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a interface{}) interface{} 的形式
		"newFuncIIE":       NewFuncInterfaceInterfaceErrorB, // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a interface{}) (interface{}, error) 的形式
		"newFuncIsI":       NewFuncInterfacesInterface,      // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a ...interface{}) interface{} 的形式
		"newFuncSSE":       NewFuncStringStringErrorB,       // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a string) (string, error) 的形式
		"newFuncSS":        NewFuncStringStringB,            // 将Gox语言中的定义的函数转换为Go语言中类似 func f(a string) string 的形式
		"newCharFunc":      newCharFunc,                     // 将Gox语言中的定义的函数转换为Charlang语言中类似 func f() 的形式
		"newStringRing":    tk.NewStringRing,                // 创建一个字符串环，大小固定，后进的会将先进的最后一个顶出来
		"getCfgStr":        getCfgString,                    // 从根目录（Windows下为C:\，*nix下为/）的gox子目录中获取文件名为参数1的配置项字符串
		"setCfgStr":        setCfgString,                    // 向根目录（Windows下为C:\，*nix下为/）的gox子目录中写入文件名为参数1，内容为参数2的配置项字符串，例：saveCfgStr("timeout", "30")
		"genQR":            tk.GenerateQR,                   // 生成二维码，例：genQR("http://www.example.com", "-level=2"), level 0..3，越高容错性越好，但越大
		"newChar":          charlang.NewChar,                // new a charlang script VM
		"runChar":          charlang.RunChar,                // run a charlang script VM
		"runCharCode":      charlang.RunCharCode,            // run a charlang script
		"runXie":           xie.RunCode,                     // run a xielang script
		"quickCompileChar": charlang.QuickCompile,           // compile a charlang script VM
		"quickRunChar":     charlang.QuickRun,               // run a charlang script VM
		"newCharAny":       charlang.NewAny,                 // create a interface{} pointer in charlang
		"newCharAnyValue":  charlang.NewAnyValue,            // create a interface{} value in charlang
		"toCharValue":      charlang.ConvertToObject,        // convert to a interface{} value in charlang
		"wrapError":        tk.WrapError,                    //
		"renderMarkdown":   tk.RenderMarkdown,               // 将Markdown格式字符串渲染为HTML

		"genToken":   tk.GenerateToken, // 生成令牌，用法：genToken("appCode", "userID", "userRole", "-secret=abc")
		"checkToken": tk.CheckToken,    // 检查令牌，如果成功，返回类似“appCode|userID|userRole|”的字符串；失败返回TXERROR字符串

		// global variables 全局变量
		"timeFormatG":        tk.TimeFormat,        // 用于时间处理时的时间格式，值为"2006-01-02 15:04:05"
		"timeFormatCompactG": tk.TimeFormatCompact, // 用于时间处理时的简化时间格式，值为"20060102150405"

		"getSystemEndian": tk.GetSystemEndian, // 获取系统的字节顺序，返回binary.BigEndian或binary.LittleEndian
		"getStack":        getStack,           // 获取堆栈
		"getVars":         getVars,            // 获取当前变量表

		"scriptPathG": scriptPathG, // 所执行脚本的路径
		"versionG":    versionG,    // Gox/Goxc的版本号
		"leBufG":      leBufG,      // 内置行文本编辑器所用的编辑缓冲区

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

	qlang.Import("mime_multipart", qlmimemultipart.Exports)

	qlang.Import("net", qlnet.Exports)
	qlang.Import("net_http", qlnethttp.Exports)
	qlang.Import("http", qlnethttp.Exports)
	qlang.Import("net_http_cookiejar", qlnet_http_cookiejar.Exports)
	qlang.Import("net_http_httputil", qlnet_http_httputil.Exports)
	qlang.Import("net_mail", qlnet_mail.Exports)
	qlang.Import("net_rpc", qlnet_rpc.Exports)
	qlang.Import("net_rpc_jsonrpc", qlnet_rpc_jsonrpc.Exports)
	qlang.Import("net_smtp", qlnet_smtp.Exports)
	qlang.Import("net_url", qlneturl.Exports)
	qlang.Import("url", qlneturl.Exports)

	qlang.Import("os", qlos.Exports)
	qlang.Import("os_exec", qlos_exec.Exports)
	qlang.Import("os_signal", qlos_signal.Exports)
	qlang.Import("os_user", qlos_user.Exports)
	qlang.Import("path", qlpath.Exports)
	qlang.Import("path_filepath", qlpathfilepath.Exports)
	qlang.Import("filepath", qlpathfilepath.Exports)

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
	qlang.Import("etree", qlgithubbeeviketree.Exports)
	qlang.Import("github_topxeq_sqltk", qlgithubtopxeqsqltk.Exports)
	qlang.Import("sqltk", qlgithubtopxeqsqltk.Exports)

	qlang.Import("github_topxeq_xmlx", qlgithub_topxeq_xmlx.Exports)

	qlang.Import("github_topxeq_awsapi", qlgithub_topxeq_awsapi.Exports)
	qlang.Import("awsapi", qlgithub_topxeq_awsapi.Exports)

	qlang.Import("github_topxeq_charlang", qlgithub_topxeq_charlang.Exports)
	qlang.Import("charlang", qlgithub_topxeq_charlang.Exports)

	qlang.Import("github_fogleman_gg", qlgithub_fogleman_gg.Exports)
	qlang.Import("gg", qlgithub_fogleman_gg.Exports)

	qlang.Import("github_360EntSecGroupSkylar_excelize", qlgithub_xuri_excelize.Exports)
	qlang.Import("github_xuri_excelize", qlgithub_xuri_excelize.Exports)
	qlang.Import("excelize", qlgithub_xuri_excelize.Exports)

	qlang.Import("github_stretchr_objx", qlgithub_stretchr_objx.Exports)

	qlang.Import("github_aliyun_alibabacloudsdkgo_services_dysmsapi", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)
	qlang.Import("aliyunsms", qlgithub_aliyun_alibabacloudsdkgo_services_dysmsapi.Exports)

	qlang.Import("github_topxeq_afero", qlgithub_topxeq_afero.Exports)
	qlang.Import("memfs", qlgithub_topxeq_afero.Exports)

	qlang.Import("github_topxeq_socks", qlgithub_topxeq_socks.Exports)

	qlang.Import("github_topxeq_regexpx", qlgithub_topxeq_regexpx.Exports)

	qlang.Import("github_domodwyer_mailyak", qlgithub_domodwyer_mailyak.Exports)
	qlang.Import("mailyak", qlgithub_domodwyer_mailyak.Exports)

	qlang.Import("github_topxeq_docxrepl", qlgithub_topxeq_docxrepl.Exports)

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

	if tk.StartsWith(codeA, "//TXDEF#") {
		tmps := tk.DecryptStringByTXDEF(codeA, "topxeq")

		if !tk.IsErrStr(tmps) {
			codeA = tmps
		}
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

func RunScriptOnHttp(codeA string, resA http.ResponseWriter, reqA *http.Request, inputA string, argsA []string, parametersA map[string]string, optionsA ...string) (string, error) {
	if tk.IfSwitchExists(optionsA, "-verbose") {
		tk.Pl("Starting...")
	}

	if !initFlag {
		initFlag = true
		importQLNonGUIPackages()
	}

	if tk.StartsWith(codeA, "//TXDEF#") {
		tmps := tk.DecryptStringByTXDEF(codeA, "topxeq")

		if !tk.IsErrStr(tmps) {
			codeA = tmps
		}
	}

	vmT := qlang.New("-noexit")

	vmT.SetVar("inputG", inputA)

	vmT.SetVar("argsG", argsA)

	vmT.SetVar("basePathG", tk.GetSwitch(optionsA, "-base=", ""))

	vmT.SetVar("paraMapG", parametersA)

	vmT.SetVar("requestG", reqA)

	vmT.SetVar("responseG", resA)

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
