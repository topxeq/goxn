package goxn

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/topxeq/qlang"
	_ "github.com/topxeq/qlang/lib/builtin" // 导入 builtin 包
	_ "github.com/topxeq/qlang/lib/chan"

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

var versionG = "0.93a"

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

func importQLNonGUIPackages() {
	var defaultExports = map[string]interface{}{
		"pass":             tk.Pass,
		"eval":             qlEval,
		"typeOf":           tk.TypeOfValue,
		"typeOfReflect":    tk.TypeOfValueReflect,
		"remove":           tk.RemoveItemsInArray,
		"pr":               tk.Pr,
		"pln":              tk.Pln,
		"prf":              tk.Printf,
		"printfln":         tk.Pl,
		"pl":               tk.Pl,
		"sprintf":          fmt.Sprintf,
		"fprintf":          fmt.Fprintf,
		"plv":              tk.Plv,
		"plvx":             tk.Plvx,
		"plvsr":            tk.Plvsr,
		"plerr":            tk.PlErr,
		"plExit":           tk.PlAndExit,
		"exit":             tk.Exit,
		"setValue":         tk.SetValue,
		"getValue":         tk.GetValue,
		"setVar":           tk.SetVar,
		"getVar":           tk.GetVar,
		"bitXor":           tk.BitXor,
		"isNil":            tk.IsNil,
		"isError":          tk.IsError,
		"strToInt":         tk.StrToIntWithDefaultValue,
		"intToStr":         tk.IntToStr,
		"floatToStr":       tk.Float64ToStr,
		"toStr":            tk.ToStr,
		"toInt":            tk.ToInt,
		"toFloat":          tk.ToFloat,
		"toLower":          strings.ToLower,
		"toUpper":          strings.ToUpper,
		"checkError":       tk.CheckError,
		"checkErrorString": tk.CheckErrorString,
		"checkErrf":        tk.CheckErrf,
		"checkErrStrf":     tk.CheckErrStrf,
		"fatalf":           tk.Fatalf,
		"isErrStr":         tk.IsErrStr,
		"errStr":           tk.ErrStr,
		"errStrf":          tk.ErrStrF,
		"getErrStr":        tk.GetErrStr,
		"errf":             tk.Errf,
		"getInput":         tk.GetUserInput,
		"getInputf":        tk.GetInputf,
		"deepClone":        tk.DeepClone,
		"deepCopy":         tk.DeepCopyFromTo,
		"getClipText":      tk.GetClipText,
		"setClipText":      tk.SetClipText,
		"trim":             tk.Trim,
		"systemCmd":        tk.SystemCmd,
		"newSSHClient":     tk.NewSSHClient,
		"getParameter":     tk.GetParameterByIndexWithDefaultValue,
		"getSwitch":        tk.GetSwitchWithDefaultValue,
		"getIntSwitch":     tk.GetSwitchWithDefaultIntValue,
		"switchExists":     tk.IfSwitchExistsWhole,
		"ifSwitchExists":   tk.IfSwitchExistsWhole,
		"xmlEncode":        tk.EncodeToXMLString,
		"htmlEncode":       tk.EncodeHTML,
		"htmlDecode":       tk.DecodeHTML,
		"base64Encode":     tk.EncodeToBase64,
		"base64Decode":     tk.DecodeFromBase64,
		"md5Encode":        tk.MD5Encrypt,
		"jsonEncode":       tk.ObjectToJSON,
		"jsonDecode":       tk.JSONToObject,
		"simpleEncode":     tk.EncodeStringCustomEx,
		"simpleDecode":     tk.DecodeStringCustom,

		"versionG": versionG,
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
