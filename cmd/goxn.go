package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/topxeq/goxn"
	"github.com/topxeq/tk"
)

var muxG *http.ServeMux
var portG = ":80"
var basePathG = "."

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
	tk.Pl("RequestURI: %v", req.RequestURI)

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

	tk.Pl("[%v] REQ: %#v (%#v)", tk.GetNowTimeStringFormal(), reqT, paraMapT)

	toWriteT := ""

	fcT := tk.LoadStringFromFile(filepath.Join(basePathG, "wms", reqT+".gox"))
	if tk.IsErrStr(fcT) {
		res.Write([]byte(tk.ErrStrf("操作失败：%v", tk.GetErrStr(fcT))))
		return
	}

	// paraMapT["_reqHost"] = req.Host
	// paraMapT["_reqInfo"] = fmt.Sprintf("%#v", req)

	toWriteT, errT = goxn.RunScriptOnHttp(fcT, res, req, paraMapT["input"], nil, paraMapT, "-base="+basePathG)

	if errT != nil {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		res.Write([]byte(tk.ErrStrf("操作失败", tk.GetErrStr(errT.Error()))))
		return
	}

	if toWriteT == "TX_END_RESPONSE_XT" {
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	res.Write([]byte(toWriteT))
}

func main() {

	portG = tk.GetSwitch(os.Args, "-port=", portG)
	basePathG = tk.GetSwitch(os.Args, "-dir=", basePathG)

	muxG = http.NewServeMux()

	muxG.HandleFunc("/wms/", doWms)
	muxG.HandleFunc("/wms", doWms)

	tk.PlNow("try starting server on %v（basePath: %v）...", portG, basePathG)

	err := http.ListenAndServe(portG, muxG)

	if err != nil {
		tk.PlNow("failed to start: %v", err)
	}

	// resultT, errT := goxn.RunScript(tk.LoadStringFromFile(os.Args[1]), tk.GetSwitch(os.Args, "-input="), os.Args, nil)

	// tk.CheckErrf("error: %v", errT)

	// tk.Pl("%v", resultT)

}
