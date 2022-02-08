Client Request
==============

Similar to Curl, httplib is used to simulate http requests sent by
clients. Similar to jQuery, it supports method chaining. It’s easy to
use and it can be installed by:

::

   go get github.com/W3-Engineers-Ltd/Radiant/client/httplib

Basic Usage
-----------

Import package:

::

   import (
       "github.com/W3-Engineers-Ltd/Radiant/client/httplib"
   )   

Initialize request method and url:

::

   req := httplib.Get("http://radiant.vip/")

Send the request and retrieve the data in the response:

::

   str, err := req.String()
   if err != nil {
       t.Fatal(err)
   }
   fmt.Println(str)

Method Functions
----------------

httplib supports these methods:

-  ``Get(url string)``
-  ``Post(url string)``
-  ``Put(url string)``
-  ``Delete(url string)``
-  ``Head(url string)``

Debug Output
------------

Enable debug information output:

::

   req.Debug(true)

Then it will output debug information:

::

   httplib.Get("http://radiant.vip/").Debug(true).Response()

   // Output
   GET / HTTP/0.0
   Host: radiant.vip
   User-Agent: radiantServer

HTTPS Request
-------------

If the requested scheme is https, we need to set the TLS of client:

::

   req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

`Learn more about TLS
settings <http://gowalker.org/crypto/tls#Config>`__

Set Timeout
-----------

Can set request timeout and data reading timeout by:

::

   req.SetTimeout(connectTimeout, readWriteTimeout)

It is a function of request object. So it can be done like this:

::

   httplib.Get("http://radiant.vip/").SetTimeout(100 * time.Second, 30 * time.Second).Response()

Set Request Params
------------------

For Put or Post requests, we may need to send parameters. Parameters can
be set in the following manner:

::

   req := httplib.Post("http://radiant.vip/")
   req.Param("username","astaxie")
   req.Param("password","123456")

Send big data
-------------

To simulate file uploading or to send big data, one can use the ``Body``
function:

::

   req := httplib.Post("http://radiant.vip/")
   bt,err:=ioutil.ReadFile("hello.txt")
   if err!=nil{
       log.Fatal("read file err:",err)
   }
   req.Body(bt)

Set header
----------

To simulate header values, e.g.:

::

   Accept-Encoding:gzip,deflate,sdch
   Host:radiant.vip
   User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36

Can use ``Header`` function:

::

   req := httplib.Post("http://radiant.vip/")
   req.Header("Accept-Encoding","gzip,deflate,sdch")
   req.Header("Host","radiant.vip")
   req.Header("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

Upload file
-----------

PostFile function requires the first parameter to be the name of form
and the second parameter is the filename or filepath you want to send.

::

   b:=httplib.Post("http://radiant.vip/")
   b.Param("username","astaxie")
   b.Param("password","123456")
   b.PostFile("uploadfile1", "httplib.pdf")
   b.PostFile("uploadfile2", "httplib.txt")
   str, err := b.String()
   if err != nil {
       t.Fatal(err)
   }

Get Response
------------

The settings above are before sending request, how can we get response
after request? Here are the ways:

+------------------+--------------+-----------------------------------+
| Method           | Type         | Description                       |
+==================+==============+===================================+
| ``req.Response() | ``(*http.Res | This is a ``http.Response``       |
| ``               | ponse, error | object. You can get data from it. |
|                  | )``          |                                   |
+------------------+--------------+-----------------------------------+
| ``req.Bytes()``  | ``([]byte, e | Return raw response body.         |
|                  | rror)``      |                                   |
+------------------+--------------+-----------------------------------+
| ``req.String()`` | ``(string, e | Return raw response body.         |
|                  | rror)``      |                                   |
+------------------+--------------+-----------------------------------+
| ``req.ToFile(fil | ``error``    | Save response body into a file.   |
| ename string)``  |              |                                   |
+------------------+--------------+-----------------------------------+
| ``req.ToJSON(res | ``error``    | Parse JSON response into the      |
| ult interface{}) |              | result object.                    |
| ``               |              |                                   |
+------------------+--------------+-----------------------------------+
| ``req.ToXml(resu | ``error``    | Parse XML response into the       |
| lt interface{})` |              | result object.                    |
| `                |              |                                   |
+------------------+--------------+-----------------------------------+

Filter
======

In order to support some AOP feature, e.g. logs, tracing, we designed
``filter-chain`` for httplib.

There are two key interfaces:

.. code:: go

   type FilterChain func(next Filter) Filter

   type Filter func(ctx context.Context, req *RadiantHTTPRequest) (*http.Response, error)

This is a typical usage of ``Filter-Chain`` pattern. So you must invoke
``next(...)`` when you want to implement your own logic.

Here is an example：

.. code:: go

   func myFilter(next httplib.Filter) httplib.Filter {
       return func(ctx context.Context, req *httplib.RadiantHTTPRequest) (*http.Response, error) {
           r := req.GetRequest()
           logs.Info("hello, here is the filter: ", r.URL)
           // Never forget invoke this. Or the request will not be sent
           return next(ctx, req)
       }
   }

And we could register this filter as global filter:

.. code:: go

       httplib.SetDefaultSetting(httplib.RadiantHTTPSettings{

           FilterChains: []httplib.FilterChain{
               myFilter,
           },

           UserAgent:        "radiantServer",
           ConnectTimeout:   60 * time.Second,
           ReadWriteTimeout: 60 * time.Second,
           Gzip:             true,
           DumpBody:         true,
       })

Sometimes you only want to use the filter for specific requests:

.. code:: go

   req.AddFilters(myFilter)

We provide some filters.

Prometheus Filter
-----------------

It’s used to support ``Prometheus`` framework to collect metric data.

.. code:: go

       builder := prometheus.FilterChainBuilder{
           AppName: "My-test",
           ServerName: "User-server-1",
           RunMode: "dev",
       }
       req := httplib.Get("http://radiant.vip/")
       // only work for this request, or using SetDefaultSetting to support all requests
       req.AddFilters(builder.FilterChain)

       resp, err := req.Response()
       if err != nil {
           logs.Error("could not get response: ", err)
       } else {
           logs.Info(resp)
       }

If you don’t use Radiant’s admin service, you must expose ``prometheus``
port manually.

Opentracing Filter
------------------

.. code:: go

       builder := opentracing.FilterChainBuilder{}
       req := httplib.Get("http://radiant.vip/")
       // only work for this request, or using SetDefaultSetting to support all requests
       req.AddFilters(builder.FilterChain)

       resp, err := req.Response()
       if err != nil {
           logs.Error("could not get response: ", err)
       } else {
           logs.Info(resp)
       }

Don’t forget to register ``Opentracing`` real implementation.

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