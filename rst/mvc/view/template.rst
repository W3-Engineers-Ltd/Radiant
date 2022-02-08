Template Functions
==================

Radiant supports custom template functions but it must be set as below
before ``web.Run()``:

::

   func hello(in string)(out string){
       out = in + "world"
       return
   }

   web.AddFuncMap("hi",hello)

Then you can use these functions in template:

::

   {{.Content | hi}}

Here are Radiant’s built-in template functions:

-  dateformat

   Format time, return string. {{dateformat .Time
   “2006-01-02T15:04:05Z07:00”}}

-  date

   This is similar to date function in PHP. It can easily return time by
   string {{date .T “Y-m-d H:i:s”}}

-  compare

   Compare two objects. If they are the same return true otherwise
   return false. {{compare .A .B}}

-  substr

   Return sub string. supports UTF-8 string. {{substr .Str 0 30}}

-  html2str

   Parse html to string by removing tags like script and css and return
   text. {{html2str .Htmlinfo}}

-  str2html Parse string to HTML, no escaping. {{str2html .Strhtml}}

-  htmlquote

   Escaping html. {{htmlquote .quote}}

-  htmlunquote

   Unescaping to html. {{htmlunquote .unquote}}

-  renderform

   Generate form from StructTag. {{&struct \| renderform}}

-  assets_js

   Create a ``<script>`` tag from js src. {{assets_js src}}

-  assets_css

   Create a ``<link>`` tag from css src. {{assets_css src}}

-  config

   Get the value of AppConfig. {{config configType configKey
   defaultValue}}. configType must be String, Bool, Int, Int64, Float,
   or DIY

-  map_get

   Get value of ``map`` by key

   ::

        // In controller
        Data["m"] = map[string]interface{} {
            "a": 1,
            "1": map[string]float64{
                "c": 4,
            },
        }

        // In view
        {{ map_get m "a" }} // return 1
        {{ map_get m 1 "c" }} // return 4

-  urlfor

   Get the URL of a controller method

   ::

        {{urlfor "TestController.List"}}

   `more details </en-US/mvc/controller/urlbuilding.html>`__

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