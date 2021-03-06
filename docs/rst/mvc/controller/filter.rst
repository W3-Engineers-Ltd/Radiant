Filters
=======

Radiant supports custom filter middlewares. E.g.: user authentication
and force redirection.

Activating Filters
==================

Before filters can be used, filters must be activated.

Filters can be activated at the code level:

``web.BConfig.WebConfig.Session.SessionOn = true``

Filters can also be activated in the configuration file:

``SessionOn = true``

Attempting to use a filter without activation will cause a
``Handler crashed with error runtime error: invalid memory address or nil pointer dereference``
error

Inserting Filters
=================

A filter function can be inserted as follows:

.. code:: go


   web.InsertFilter(pattern string, pos int, filter FilterFunc, opts ...FilterOpt)

This is the FilterFunc signature:

.. code:: go

   type FilterFunc func(*context.Context)

The *context* must be imported if this has not already been done:

.. code:: go

   import "github.com/W3-Engineers-Ltd/Radiant/server/web/context"

InsertFilter’s four parameters:

-  ``pattern``: string or regex to match against router rules. Use
   ``/*`` to match all.
-  ``pos``: the place to execute the Filter. There are five fixed
   parameters representing different execution processes.

   -  web.BeforeStatic: Before finding the static file.
   -  web.BeforeRouter: Before finding router.
   -  web.BeforeExec: After finding router and before executing the
      matched Controller.
   -  web.AfterExec: After executing Controller.
   -  web.FinishRouter: After finishing router.

-  ``filter``: filter function type FilterFunc func(*context.Context)
-  ``opts``:

   1. web.WithReturnOnOutput: whether to continue running if has output.
      default is false.
   2. web.WithResetParams: whether to reset parameters to their previous
      values after the filter has completed.
   3. web.WithCaseSensitive: whether case sensitive

..

   from radiant version 1.3 AddFilter has been removed

Here is an example to authenticate if the user is logged in for all
requests:

.. code:: go

   var FilterUser = func(ctx *context.Context) {
       if strings.HasPrefix(ctx.Input.URL(), "/login") {
           return
       }
       
       _, ok := ctx.Input.Session("uid").(int)
       if !ok {
           ctx.Redirect(302, "/login")
       }
   }

   web.InsertFilter("/*", web.BeforeRouter, FilterUser)

..

   Filters which use session must be executed after ``BeforeRouter``
   because session is not initialized before that. web session module
   must be enabled first. (see `Session
   control <../controller/session.html>`__)

Filters can be run against requests which use a regex router rule for
matching:

.. code:: go

   var FilterUser = func(ctx *context.Context) {
       _, ok := ctx.Input.Session("uid").(int)
       if !ok {
           ctx.Redirect(302, "/login")
       }
   }
   web.InsertFilter("/user/:id([0-9]+)", web.BeforeRouter, FilterUser)

Filter Implementation UrlManager
--------------------------------

Context.Input has new features ``RunController`` and ``RunMethod`` from
radiant version 1.1.2. These can control the router in the filter and
skip the Radiant router rule.

For example:

.. code:: go

   var UrlManager = func(ctx *context.Context) {
       // read urlMapping data from database
       urlMapping := model.GetUrlMapping()
       for baseurl,rule := range urlMapping {
           if baseurl == ctx.Request.RequestURI {
               ctx.Input.RunController = rule.controller
               ctx.Input.RunMethod = rule.method
               break
           }
       }
   }

   web.InsertFilter("/*", web.BeforeRouter, UrlManager)

Filter和FilterChain
-------------------

In v1.x, we can’t invoke next ``Filter`` inside a ``Filter``. So we got
a problem: we could not do something “surrounding” request execution.

For example, if we want to do:

::

   func filter() {
       // do something before serving request
       handleRequest()
       // do something after serving request
   }

The typical cases are tracing and metrics.

So we enhance ``Filter`` by designing a new interface:

.. code:: go

   type FilterChain func(next FilterFunc) FilterFunc

Here is a simple example:

.. code:: go

   package main

   import (
       "github.com/W3-Engineers-Ltd/Radiant/core/logs"
       "github.com/W3-Engineers-Ltd/Radiant/server/web"
       "github.com/W3-Engineers-Ltd/Radiant/server/web/context"
   )

   func main() {
       web.InsertFilterChain("/*", func(next web.FilterFunc) web.FilterFunc {
           return func(ctx *context.Context) {
               // do something
               logs.Info("hello")
               // don't forget this
               next(ctx)

               // do something
           }
       })
   }

In this example, we only output “hello” and then we invoke next filter.

Prometheus例子
~~~~~~~~~~~~~~

.. code:: go

   package main

   import (
       "time"

       "github.com/W3-Engineers-Ltd/Radiant/server/web"
       "github.com/W3-Engineers-Ltd/Radiant/server/web/filter/prometheus"
   )

   func main() {
       // we start admin service
       // Prometheus will fetch metrics data from admin service's port
       web.BConfig.Listen.EnableAdmin = true

       web.BConfig.AppName = "my app"

       ctrl := &MainController{}
       web.Router("/hello", ctrl, "get:Hello")
       fb := &prometheus.FilterChainBuilder{}
       web.InsertFilterChain("/*", fb.FilterChain)
       web.Run(":8080")
       // after you start the server
       // and GET http://localhost:8080/hello
       // access http://localhost:8088/metrics
       // you can see something looks like:
       // http_request_web_sum{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1002
       // http_request_web_count{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
       // http_request_web_sum{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1004
       // http_request_web_count{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
   }

   type MainController struct {
       web.Controller
   }

   func (ctrl *MainController) Hello() {
       time.Sleep(time.Second)
       ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
   }

If you don’t use Radiant’s admin service, don’t forget to expose
``prometheus``\ ’s port.

Opentracing例子
~~~~~~~~~~~~~~~

.. code:: go

   package main

   import (
       "time"

       "github.com/W3-Engineers-Ltd/Radiant/server/web"
       "github.com/W3-Engineers-Ltd/Radiant/server/web/filter/opentracing"
   )

   func main() {
       // don't forget this to inject the opentracing API's implementation
       // opentracing2.SetGlobalTracer()

       web.BConfig.AppName = "my app"

       ctrl := &MainController{}
       web.Router("/hello", ctrl, "get:Hello")
       fb := &opentracing.FilterChainBuilder{}
       web.InsertFilterChain("/*", fb.FilterChain)
       web.Run(":8080")
       // after you start the server
   }

   type MainController struct {
       web.Controller
   }

   func (ctrl *MainController) Hello() {
       time.Sleep(time.Second)
       ctrl.Ctx.ResponseWriter.Write([]byte("Hello, world"))
   }

Don’t forget to using ``SetGlobalTracer`` to initialize opentracing.

.. toctree::
   :maxdepth: 4
   :caption: Contents:

   rst/quickstart
   rst/quickstart/README

.. toctree::
   :maxdepth: 4
   :caption: Quickstart Introduction:


   rst/quickstart/new
   rst/quickstart/router
   rst/quickstart/controller
   rst/quickstart/model
   rst/quickstart/view

.. toctree::
   :maxdepth: 4
   :caption: Introduction:

   rst/intro/Introduction
   rst/advantage/README
   rst/install/install
   rst/install/radical

.. toctree::
   :maxdepth: 4
   :caption: MVC Introduction:

   rst/mvc/README
.. toctree::
   :maxdepth: 4
   :caption: Controller:

   rst/mvc/controller/config
   rst/mvc/controller/controller
   rst/mvc/controller/error
   rst/mvc/controller/filter
   rst/mvc/controller/flash
   rst/mvc/controller/jsonxml
   rst/mvc/controller/params
   rst/mvc/controller/router
   rst/mvc/controller/session
   rst/mvc/controller/urlbuilding
   rst/mvc/controller/validation
   rst/mvc/controller/xsrf
.. toctree::
   :maxdepth: 4
   :caption: Models:

   rst/mvc/model/overview
   rst/mvc/model/orm
   rst/mvc/model/object
   rst/mvc/model/query
   rst/mvc/model/rawsql
   rst/mvc/model/querybuilder
   rst/mvc/model/transaction
   rst/mvc/model/models
   rst/mvc/model/commandline
   rst/mvc/model/test
   rst/mvc/model/custome_fields
   rst/mvc/model/faq

.. toctree::
   :maxdepth: 4
   :caption: Views:

   rst/mvc/view/view
   rst/mvc/view/template
   rst/mvc/view/static
   rst/mvc/view/page
   rst/mvc/view/global_variables

.. toctree::
   :maxdepth: 4
   :caption: Modules:

   rst/module/README
   rst/module/session
   rst/module/cache
   rst/module/logs
   rst/module/httplib
   rst/module/context
   rst/module/task
   rst/module/config

.. toctree::
   :maxdepth: 4
   :caption: Advanced Radiant:

   rst/advantage/README
   rst/advantage/monitor
   rst/advantage/docs

.. toctree::
   :maxdepth: 4
   :caption: Deployment:

   rst/deploy/README
   rst/deploy/radiant
   rst/deploy/supervisor
   rst/deploy/systemctl
   rst/deploy/nginx
   rst/deploy/apache

.. toctree::
   :maxdepth: 4
   :caption: FAQ:

   rst/faq/FAQ