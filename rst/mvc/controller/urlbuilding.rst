URL Building
------------

If it can match URLs, can Radiant also generate them? Of course it can.
To build a URL to a specific function you can use the URLFor() function.
It accepts the name of the function of Controller as first argument and
a number of keyword arguments, each corresponding to the variable part
of the URL rule. Unknown variable parts are appended to the URL as query
parameters. Here are some examples:

Here is the controller definition:

::

   type TestController struct {
       web.Controller
   }

   func (this *TestController) Get() {
       this.Data["Username"] = "astaxie"
       this.Ctx.Output.Body([]byte("ok"))
   }

   func (this *TestController) List() {
       this.Ctx.Output.Body([]byte("i am list"))
   }

   func (this *TestController) Params() {
       this.Ctx.Output.Body([]byte(this.Ctx.Input.Params["0"] + this.Ctx.Input.Params["1"] + this.Ctx.Input.Params["2"]))
   }

   func (this *TestController) Myext() {
       this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
   }

   func (this *TestController) GetUrl() {
       this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
   }

This is how you register the router:

::

   web.Router("/api/list", &TestController{}, "*:List")
   web.Router("/person/:last/:first", &TestController{})
   web.AutoRouter(&TestController{})

This is how you generate the url:

::

   web.URLFor("TestController.List")
   // Output /api/list

   web.URLFor("TestController.Get", ":last", "xie", ":first", "asta")
   // Output /person/xie/asta

   web.URLFor("TestController.Myext")
   // Output /Test/Myext

   web.URLFor("TestController.GetUrl")
   // Output /Test/GetUrl

This is how you use it in a template
------------------------------------

radiant has already registered the template function ``urlfor``. You can
use it like this:

::

   {{urlfor "TestController.List"}}
   // Output /api/list

   {{urlfor "TestController.Get" ":last" "xie" ":first" "asta"}}
   // Output /person/xie/asta

Why would you want to build URLs instead of hard-coding them into your
templates? There are three good reasons for this:

1. Reversing is often more descriptive than hard-coding the URLs. More
   importantly, it allows you to change URLs in one go, without having
   to remember to change URLs all over the place.
2. URL building will handle escaping of special characters and Unicode
   data transparently for you, so you don’t have to deal with them.

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