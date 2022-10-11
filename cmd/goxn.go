package main

import (
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/topxeq/goxn"
	"github.com/topxeq/tk"

	"gitee.com/topxeq/xie"
)

var muxG *http.ServeMux
var portG = ":80"
var sslPortG = ":443"
var basePathG = "."
var webPathG = "."
var certPathG = "."
var verboseG = false

func doWms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
	}

	reqT := tk.GetFormValueWithDefaultValue(req, "wms", "")
	if verboseG {
		tk.Pl("RequestURI: %v", req.RequestURI)
	}

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/wms") {
			reqT = req.RequestURI[4:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(tk.ErrStrf("操作失败：%v", "invalid vo format")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if !tk.EndsWith(fileNameT, ".gox") {
		fileNameT += ".gox"
	}

	fcT := tk.LoadStringFromFile(filepath.Join(basePathG, fileNameT))
	if tk.IsErrStr(fcT) {
		res.Write([]byte(tk.ErrStrf("操作失败：%v", tk.GetErrStr(fcT))))
		return
	}

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

	toWriteT, errT = goxn.RunScriptOnHttp(fcT, res, req, paraMapT["input"], nil, paraMapT, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")

		errStrT := tk.ErrStrf("操作失败：%v", errT)
		tk.Pln(errStrT)
		res.Write([]byte(errStrT))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))
}

func replaceHtml(strA string, mapA map[string]string) string {
	if mapA == nil {
		return strA
	}

	for k, v := range mapA {
		strA = tk.Replace(strA, "TX_"+k+"_XT", v)
	}

	return strA
}

func genFailCompact(titleA, msgA string, optsA ...string) string {
	// mapT := map[string]string{
	// 	"msgTitle":    titleA,
	// 	"msg":         msgA,
	// 	"subMsg":      "",
	// 	"actionTitle": "返回",
	// 	"actionHref":  "javascript:history.back();",
	// }

	// var fileNameT = "fail.html"

	// if tk.IfSwitchExists(optsA, "-compact") {
	// 	fileNameT = "failcompact.html"
	// }

	// tmplT := tk.LoadStringFromFile(filepath.Join(basePathG, "tmpl", fileNameT))

	// tmplT = replaceHtml(tmplT, mapT)

	tmplT := tk.ErrStrf("%v: %v", titleA, msgA)

	return tmplT
}

// Xielang
func doXms(res http.ResponseWriter, req *http.Request) {
	if res != nil {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", "*")
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	if req != nil {
		req.ParseForm()
	}

	// tk.Pl("xms: %v", req)

	reqT := tk.GetFormValueWithDefaultValue(req, "xms", "")

	if reqT == "" {
		if tk.StartsWith(req.RequestURI, "/xms") {
			reqT = req.RequestURI[4:]
		}
	}

	tmps := tk.Split(reqT, "?")
	if len(tmps) > 1 {
		reqT = tmps[0]
	}

	if tk.StartsWith(reqT, "/") {
		reqT = reqT[1:]
	}

	// tk.Pl("charms: %v", reqT)

	var paraMapT map[string]string
	var errT error

	vo := tk.GetFormValueWithDefaultValue(req, "vo", "")

	if vo == "" {
		paraMapT = tk.FormToMap(req.Form)
	} else {
		paraMapT, errT = tk.MSSFromJSON(vo)

		if errT != nil {
			res.Write([]byte(genFailCompact("action failed", "invalid vo format", "-compact")))
			return
		}
	}

	if verboseG {
		tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)
	}

	toWriteT := ""

	fileNameT := reqT

	if !tk.EndsWith(fileNameT, ".xie") {
		fileNameT += ".xie"
	}

	fcT := tk.LoadStringFromFile(filepath.Join(basePathG, fileNameT))
	if tk.IsErrStr(fcT) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStr(fcT), "-compact")))
		return
	}

	// envT := make(map[string]interface{})

	// envT["argsG"] = paraMapT
	// envT["requestG"] = req
	// envT["responseG"] = res
	// envT["reqNameG"] = reqT

	vmT := xie.NewXie(nil)

	vmT.SetVar("argsG", paraMapT)
	vmT.SetVar("requestG", req)
	vmT.SetVar("responseG", res)
	vmT.SetVar("reqNameG", reqT)
	vmT.SetVar("basePathG", basePathG)

	// vmT.SetVar("inputG", objA)

	lrs := vmT.Load(fcT)

	if tk.IsErrStr(lrs) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStr(lrs), "-compact")))
		return
	}

	// var argsT []string = tk.JSONToStringArray(tk.GetSwitch(optsA, "-args=", "[]"))

	// if argsT != nil {
	// 	vmT.VarsM["argsG"] = argsT
	// } else {
	// 	vmT.VarsM["argsG"] = []string{}
	// }

	rs := vmT.Run()

	if errT != nil {
		res.Write([]byte(genFailCompact("action failed", errT.Error(), "-compact")))
		return
	}

	if tk.IsErrStr(rs) {
		res.Write([]byte(genFailCompact("action failed", tk.GetErrStr(rs), "-compact")))
		return
	}

	toWriteT = rs

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

}

var staticFS http.Handler = nil

func serveStaticDirHandler(w http.ResponseWriter, r *http.Request) {
	if staticFS == nil {
		// tk.Pl("staticFS: %#v", staticFS)
		// staticFS = http.StripPrefix("/w/", http.FileServer(http.Dir(filepath.Join(basePathG, "w"))))
		hdl := http.FileServer(http.Dir(webPathG))
		// tk.Pl("hdl: %#v", hdl)
		staticFS = hdl
	}

	old := r.URL.Path

	// tk.Pl("urlPath: %v", r.URL.Path)

	name := filepath.Join(webPathG, path.Clean(old))

	// tk.Pl("name: %v", name)

	info, err := os.Lstat(name)
	if err == nil {
		if !info.IsDir() {
			staticFS.ServeHTTP(w, r)
			// http.ServeFile(w, r, name)
		} else {
			if tk.IfFileExists(filepath.Join(name, "index.html")) {
				staticFS.ServeHTTP(w, r)
			} else {
				http.NotFound(w, r)
			}
		}
	} else {
		http.NotFound(w, r)
	}

}

func startHttpsServer(portA string) {
	if !tk.StartsWith(portA, ":") {
		portA = ":" + portA
	}

	err := http.ListenAndServeTLS(portA, filepath.Join(certPathG, "server.crt"), filepath.Join(certPathG, "server.key"), muxG)
	if err != nil {
		tk.PlNow("failed to start https: %v", err)
	}

}

func main() {

	portG = tk.GetSwitch(os.Args, "-port=", portG)
	sslPortG = tk.GetSwitch(os.Args, "-sslPort=", sslPortG)

	verboseG = tk.IfSwitchExistsWhole(os.Args, "-verbose")

	if !tk.StartsWith(portG, ":") {
		portG = ":" + portG
	}

	if !tk.StartsWith(sslPortG, ":") {
		sslPortG = ":" + sslPortG
	}

	basePathG = tk.GetSwitch(os.Args, "-dir=", basePathG)
	webPathG = tk.GetSwitch(os.Args, "-webDir=", basePathG)
	certPathG = tk.GetSwitch(os.Args, "-certDir=", certPathG)

	muxG = http.NewServeMux()

	muxG.HandleFunc("/wms/", doWms)
	muxG.HandleFunc("/wms", doWms)

	muxG.HandleFunc("/xms/", doXms)
	muxG.HandleFunc("/xms", doXms)

	muxG.HandleFunc("/", serveStaticDirHandler)

	tk.PlNow("goxn %v -port=%v -sslPort=%v -dir=%v -webDir=%v -certDir=%v", goxn.VersionG, portG, sslPortG, basePathG, webPathG, certPathG)

	if sslPortG != "" {
		tk.PlNow("try starting ssl server on %v...", sslPortG)
		go startHttpsServer(sslPortG)
	}

	tk.Pl("try starting server on %v ...", portG)
	err := http.ListenAndServe(portG, muxG)

	if err != nil {
		tk.PlNow("failed to start: %v", err)
	}

	// resultT, errT := goxn.RunScript(tk.LoadStringFromFile(os.Args[1]), tk.GetSwitch(os.Args, "-input="), os.Args, nil)

	// tk.CheckErrf("error: %v", errT)

	// tk.Pl("%v", resultT)

}
