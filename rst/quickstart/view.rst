Creating views
==============

In the previous example, when creating the Controller the line
``this.TplName = "index.tpl"`` was used to declare the template to be
rendered. By default ``radiant.Controller`` supports ``tpl`` and
``html`` extensions. Other extensions can be added by calling
``radiant.AddTemplateExt``.

Radiant uses the default ``html/template`` engine built into Go, so view
displays show data using standard Go templates. You can find more
information about using Go templates at `Building Web Applications with
Golang <https://github.com/astaxie/build-web-application-with-golang/blob/master/en/07.4.html>`__.

Let’s look at an example:

::

   <!DOCTYPE html>

   <html>
       <head>
           <title>Radiant</title>
           <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
       </head>
       <body>
           <header class="hero-unit" style="background-color:#A9F16C">
               <div class="container">
                   <div class="row">
                       <div class="hero-text">
                           <h1>Welcome to Radiant!</h1>
                           <p class="description">
                               Radiant is a simple & powerful Go web framework which is inspired by tornado and sinatra.
                               <br />
                               Official website: <a href="http://{{.Website}}">{{.Website}}</a>
                               <br />
                               Contact me: {{.Email}}
                           </p>
                       </div>
                   </div>
               </div>
           </header>
       </body>
   </html>

The data was assigned into a map type variable ``Data`` in the
controller, which is used as the rendering context. The data can now be
accessed and output by using the keys ``.Website`` and ``.Email``.

`The next section <static.html>`__. will describe how to serve static
files.

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