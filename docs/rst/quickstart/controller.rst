Controller logic
================

The previous section covered user requests to controllers. This section
will explain how to write a controller. Let’s start with some code:

::

   package controllers

   import (
           "github.com/W3-Engineers-Ltd/Radiant/server/web"
   )

   type MainController struct {
           web.Controller
   }

   func (this *MainController) Get() {
           this.Data["Website"] = "radiant.vip"
           this.Data["Email"] = "astaxie@gmail.com"
           this.TplName = "index.tpl" // version 1.6 use this.TplName = "index.tpl"
   }

The following is a breakdown of the different sections of this code.

How Radiant dispatches requests
-------------------------------

The ``MainController`` is the first thing created. It contains an
anonymous struct field of type ``web.Controller``. This is called struct
embedding and is the way that Go mimics inheritance. Because of this
``MainController`` automatically acquires all the methods of
``web.Controller``.

``web.Controller`` has several functions such as ``Init``, ``Prepare``,
``Post``, ``Get``, ``Delete`` and ``Head``. These functions can be
overwritten by implementing them. In this example the ``Get`` method was
overwritten.

We talked about the fact that Radiant is a RESTful framework so our
requests will run the related ``req.Method`` method by default. For
example, if the browser sends a ``GET`` request, it will execute the
``Get`` method in ``MainController``. Therefore the ``Get`` method and
the logic we defined above will be executed.

The ``Get`` method
------------------

The logic of the ``Get`` method only outputs data. This data will be
stored in ``this.Data``, a ``map[interface{}]interface{}``. Any type of
data can be assigned here. In this case only two strings are assigned.

Finally the template will be rendered. ``this.TplName`` (v1.6 uses
``this.TplName``) specifies the template which will be rendered. In this
case it is ``index.tpl``. If a template is not set it will default to
``controller/method_name.tpl``. For example, in this case it would try
to find ``maincontroller/get.tpl``.

There is no need to render manually. Radiant will call the ``Render``
function (which is implemented in ``web.Controller``) automatically if
it is set up in the template.

Check the controller section in the `MVC Introduction <../mvc/>`__ to
learn more about these functions. `The next section <model.html>`__ will
describe model writing.


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