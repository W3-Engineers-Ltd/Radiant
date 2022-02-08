JSON, XML, JSONP and YAML
=========================

Radiant is also designed for the creation of API applications. When we
build an API application, we often need to respond with JSON or XML.
Radiant provides a simple approach:

-  Respond with JSON data:

   .. code:: go

      type mystruct struct {
        FieldOne string `json:"field_one"`
      }

      func (this *AddController) Get() {
          mystruct := { ... }
          this.Data["json"] = &mystruct
          this.ServeJSON()
      }

   ServeJson will set ``content-type`` to ``application/json`` and
   JSONify the data.

-  Respond with XML data:

   .. code:: go

      func (this *AddController) Get() {
          mystruct := { ... }
          this.Data["xml"]=&mystruct
          this.ServeXML()
      }

   ServeXml will set ``content-type`` to ``application/xml`` and convert
   the data into XML.

-  Respond with jsonp

   .. code:: go

      func (this *AddController) Get() {
          mystruct := { ... }
          this.Data["jsonp"] = &mystruct
          this.ServeJSONP()
      }

   ServeJsonp will set ``content-type`` to ``application/javascript`` ,
   JSONify the data and respond to jsonp based on the request parameter
   ``callback``.

In version 1.6 names of methods were changed, it is ServeJSON(),
ServeXML(), ServeJSONP() from now on.

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