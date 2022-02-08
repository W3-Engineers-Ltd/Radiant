package sentry

import (
	"fmt"
	"github.com/W3-Engineers-Ltd/Radiant/core/logs"
	radiant "github.com/W3-Engineers-Ltd/Radiant/server/web"
	"github.com/W3-Engineers-Ltd/Radiant/server/web/context"
	"github.com/getsentry/sentry-go"
	"html/template"
	"runtime"
	"time"
)

var tpl = `
<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>UpNext Application Error</title>
    <style>
        html, body, body * {padding: 0; margin: 0;}
        #header {background:powderblue; border-bottom:solid 2px; padding: 20px 10px; text-align: center}
        #header h2{ }
        #footer {border-top:solid 1px #aaa; padding: 5px 10px; font-size: 12px; color:green;text-align: center}
        #content {padding: 5px; min-height:300px}
        #content .stack b{ font-size: 13px; color: red;}
        #content .stack pre{padding-left: 10px;}
        table {}
        td.t {text-align: right; padding-right: 5px; color: #888;}
    </style>
    <script type="text/javascript">
		var show = false;	
		function showHide(){
			var el = document.getElementById("error");
			if(show===false){
				el.style.display = "block";
				show = true;
			}else{
				el.style.display = "none";
				show = false;
			}
			
		}
    </script>
</head>
<body>
    <div id="header">
        <h2>UpNext Error</h2>
    </div>
    <div id="content">
		<h3>Request Method: </td><td>{{.RequestMethod}}</h3>
		<h3>Request URL: </td><td>{{.RequestURL}}</h3>
		<h3>RemoteAddr: </td><td>{{.RemoteAddr}}</h3> 
		<h3>UserAgent: </td><td>{{.UserAgent}}</h3> 
	
    </div>
	<div id="error" style="display: none">
		<pre>{{.AppError}}</pre>
	</div>
    <div id="footer">
        <p onclick="showHide()">Powered by {{ .RadiantVersion }}</p>
    </div>
	
</body>
</html>
`

// ShowErr render default application error page with error and stack string.
func showErr(err interface{}, ctx *context.Context, stack string) {
	t, _ := template.New("upnexterrortemplate").Parse(tpl)

	data := map[string]string{
		"AppError":       fmt.Sprintf("%s:%v\n%s", ctx.Input.Host(), err, stack),
		"RequestMethod":  ctx.Input.Method(),
		"RequestURL":     ctx.Input.URI(),
		"RemoteAddr":     ctx.Input.IP(),
		"UserAgent":      ctx.Input.UserAgent(),
		"RadiantVersion": "UpNext",
	}
	_ = t.Execute(ctx.ResponseWriter, data)
}

func SentryCaptureException(ctx *context.Context, err error) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetRequest(ctx.Request)
		scope.SetRequestBody(ctx.Input.RequestBody)
		scope.SetExtra("UserIP", ctx.Input.IP())
	})

	sentry.CaptureException(err)
}

func SentryCaptureMessage(ctx *context.Context, msg string) {
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetRequest(ctx.Request)
		scope.SetRequestBody(ctx.Input.RequestBody)
		scope.SetExtra("UserIP", ctx.Input.IP())
	})

	sentry.CaptureMessage(msg)
}

var ErrorTracker = func(ctx *context.Context, config *radiant.Config) {
	if err := recover(); err != nil {
		if err == radiant.ErrAbort {
			return
		}

		var stack string
		for i := 1; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			logs.Critical(fmt.Sprintf("%s:%d", file, line))
			stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line))
		}
		if ctx.Output.Status == 0 || ctx.Output.Status > 499 {
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetRequest(ctx.Request)
				scope.SetRequestBody(ctx.Input.RequestBody)
				scope.SetExtra("UserIP", ctx.Input.IP())
			})

			sentry.CurrentHub().Recover(fmt.Sprintf("%v\n%s", err, stack))
			sentry.Flush(time.Second * 5)
		}

		if !radiant.BConfig.RecoverPanic {
			panic(err)
		}
		if radiant.BConfig.EnableErrorsShow {
			if _, ok := radiant.ErrorMaps[fmt.Sprint(err)]; ok {
				radiant.Exception(uint64(ctx.Output.Status), ctx)
				return
			}
		}

		logs.Error(err)
		if ctx.Output.Status != 0 {
			ctx.ResponseWriter.WriteHeader(ctx.Output.Status)
		} else {
			ctx.ResponseWriter.WriteHeader(500)
		}
		if radiant.BConfig.RunMode == radiant.DEV {
			showErr(err, ctx, stack)
		} else {
			showErr(err, ctx, "No more info!")
		}

	}
}

func SentryInit() {
	sentryDNS, _ := radiant.AppConfig.String("SentryDSN")
	releaseVersion, _ := radiant.AppConfig.String("ASSET_HASH")
	serverName, _ := radiant.AppConfig.String("ServerName")
	runMode, _ := radiant.AppConfig.String("RunMode")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         sentryDNS,
		Environment: runMode,
		Release:     releaseVersion,
		ServerName:  serverName,
		Debug:       false,
	})
	if err != nil {
		logs.Error("sentry initialization error: %s", err)
	} else {
		fmt.Println("sentry running....")
	}
	defer sentry.Flush(2 * time.Second)
}

// How to use sentry

//Raise default panic error
/*v := []string{"hello world"}
fmt.Println(v[111])
*/

//for error log
//_, aErr := radiant.AppConfig.String("HelloWorld")
// utils.SentryCaptureException(ctx, aErr)
// or,
//sentry.CaptureException(aErr)

//for any message
// utils.SentryCaptureMessage(ctx, aErr)
// or,
//sentry.CaptureMessage("Your message as Hello world.")
