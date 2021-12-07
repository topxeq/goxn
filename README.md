# goxn
Embeddable version of Gox language. A network-server/micro-service mode is available as well.

## Usage

goxn -port=:80 -dir=. -webDir=.

use -dir= to specify the main micro-service path, -webDir= to specify the static web/file server path.

## Example

```
goxn -port=:8080 -dir=c:\scripts -webDir=d:\web
```

A web/application server will start and listen on port 8080. Any scripts file(.gox) written by Gox language will be executed(and is expected to return proper http response). For example, if there is a file named "test.gox" in c:\scripts, the URL to access the service is http://127.0.0.1:8080/wms/test .

And all the files in d:\web, will be served as static web file, i.e. browse to http://127.0.0.1:8080/index.html to access index.html in d:\web. 