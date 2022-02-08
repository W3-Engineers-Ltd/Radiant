Automated API Documentation
===========================

Automated documentation is a very cool feature that I found to be
desirable. Now it became a reality in Radiant. As I said Radiant will
not only boost the development of your API but also make the API easy to
use for the user.

Radiant implemented the `swagger specification <http://swagger.io/>`__
for API documentation. It’s very easy to create powerful interactive API
documentation.

Ok, let’s try it out now. First let’s create a new API application by
``Radical api radiantapi``

API global settings
===================

Add the following comments at the top of ``routers/router.go``:

::

   // @APIVersion 1.0.0
   // @Title mobile API
   // @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
   // @Contact astaxie@gmail.com
   package routers

The comments above set the global information. The available settings:

-  @APIVersion
-  @Title
-  @Description
-  @Contact
-  @TermsOfServiceUrl
-  @License
-  @LicenseUrl
-  @Name
-  @URL
-  @LicenseUrl
-  @License
-  @Schemes
-  @Host

Router Parsing
--------------

Right now automated API documentation only supports ``NSNamespace`` and
``NSInclude`` and it only supports two levels of parsing. The first
level is the API version and the second level is the modules.

This only works for ``dev`` environment. We think that, all API must be
tested, and if users are able to generate API in non-dev environment,
some users may use it in production environment.

In v2.x, a big change is that we scan the directory which is configured
by `CommentRouterPath </en-US/mvc/controller/config.html>`__.

But we only generate router files, you must call ``Include`` method to
use it.

::

   func init() {
       ns :=
           web.NewNamespace("/v1",
               web.NSNamespace("/customer",
                   web.NSInclude(
                       &controllers.CustomerController{},
                       &controllers.CustomerCookieCheckerController{},
                   ),
               ),
               web.NSNamespace("/catalog",
                   web.NSInclude(
                       &controllers.CatalogController{},
                   ),
               ),
               web.NSNamespace("/newsletter",
                   web.NSInclude(
                       &controllers.NewsLetterController{},
                   ),
               ),
               web.NSNamespace("/cms",
                   web.NSInclude(
                       &controllers.CMSController{},
                   ),
               ),
               web.NSNamespace("/suggest",
                   web.NSInclude(
                       &controllers.SearchController{},
                   ),
               ),
           )
       web.AddNamespace(ns)
   }

Application Comment
-------------------

This is the most important part of comment. For example:

::

   package controllers

   import "github.com/W3-Engineers-Ltd/Radiant/server/web"

   // CMS API
   type CMSController struct {
       web.Controller
   }

   func (c *CMSController) URLMapping() {
       c.Mapping("StaticBlock", c.StaticBlock)
       c.Mapping("Product", c.Product)
   }

   // @Title getStaticBlock
   // @Summary getStaticBlock
   // @Deprecated Deprecated
   // @Description get all the staticblock by key
   // @Param   key path    string  true    "The static block key." default_value
   // @Success 200 {object} ZDT.ZDTMisc.CmsResponse
   // @Failure 400 Bad request
   // @Failure 404 Not found
   // @Accept json
   // @router /staticblock/:key [get]
   func (c *CMSController) StaticBlock() {

   }

   // @Title Get Product list
   // @Description Get Product list by some info
   // @Success 200 {object} models.ZDTProduct.ProductList
   // @Param   category_id     query   int false       "category id"
   // @Param   brand_id    query   int false       "brand id"
   // @Param   query   query   string  false       "query of search"
   // @Param   segment query   string  false       "segment"
   // @Param   sort    query   string  false       "sort option"
   // @Param   dir     query   string  false       "direction asc or desc"
   // @Param   offset  query   int     false       "offset"
   // @Param   limit   query   int     false       "count limit"
   // @Param   price           query   float       false       "price"
   // @Param   special_price   query   bool        false       "whether this is special price"
   // @Param   size            query   string      false       "size filter"
   // @Param   color           query   string      false       "color filter"
   // @Param   format          query   bool        false       "choose return format"
   // @Failure 400 no enough input
   // @Failure 500 get products common error
   // @router /products [get]
   func (c *CMSController) Product() {

   }

In the code above, we defined the comment on top of ``CMSController`` is
the information for this module. Then we defined the comment for every
controller’s methods.

Below is a list of supported comments for generating swagger APIs:

-  @Accept Aceept type json/xml/html/plain
-  @Deprecated Deprecated flag.
-  @Title

   The title for this API. It’s a string, and all the content after the
   first space will be parsed as the title.

-  @Description

   The description for this API. It’s a string, and all the content
   after the first space will be parsed as the description.

-  @Param

   ``@Param`` defines the parameters sent to the server. There are five
   columns for each ``@Param``:

   1. parameter key;
   2. parameter sending type; It can be ``formData``, ``query``,
      ``path``, ``body`` or ``header``. ``formData`` means the parameter
      sends by POST ( set Content-Type to
      application/x-www-form-urlencoded ) . ``query`` means the
      parameter sends by GET in url. ``path`` means the parameter in the
      url path, such as key in the former example. ``body`` means the
      raw data send from request body. ``header`` means the parameter is
      in request header.
   3. parameter data type
   4. required
   5. comment
   6. default value

-  @Success

   The success message returned to client. Three parameters.

   1. status code.
   2. return type; Must wrap with {}.
   3. returned object or string. For {object}, use path and the object
      name of your project here and ``radical`` tool will look up the
      object while generating the docs. For example
      ``models.ZDTProduct.ProductList`` represents ``ProductList``
      object under ``/models/ZDTProduct``

   ..

            Use space to separate these three parameters

-  @Failure

   The failure message returned to client. Two parameters separated by
   space.

   1. Status code.
   2. Error message.

-  @router

   Router information. Two parameters separated by space.

   1. The request’s router address.
   2. Supported request methods. Wrap in ``[]``. Use ``,`` to separate
      multiple methods.

Generate documentation automatically
------------------------------------

Make it work by following the steps: 1. Enable docs by setting
``EnableDocs = true`` in ``conf/app.conf``. 2. Use
``radical run -downdoc=true -gendoc=true`` to run your API application
and rebuild documentation automatically. 3. Visit ``/swagger`` in your
project. (see image #1 below)

Your API documentation is available now. Open your browser and check it
out.

|image0|

|image1|

Problems You May Have
---------------------

1. CORS Two solutions:

   1. Integrate ``swagger`` into the application. Download
      `swagger <https://github.com/web/swagger/releases>`__ and put it
      into project folder. (``radical run -downdoc=true`` will also
      download it and put it into project folder) And before
      ``web.Run()`` in ``func main()`` of ``main.go``
      ``go  if web.BConfig.RunMode == "dev" {      web.BConfig.WebConfig.DirectoryIndex = true      web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"  }``
      And then visit ``/swagger`` in your project.

   2. Make API support CORS
      ``go  ctx.Output.Header("Access-Control-Allow-Origin", "*")``

2. Other problems. This is a feature used in my own project. If you have
   some other problems please fire issues to us.

.. |image0| image:: ../images/docs.png
.. |image1| image:: ../images/doc_test.png


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