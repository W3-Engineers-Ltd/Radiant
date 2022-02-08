Project routing settings
========================

The previous section covered creating and running a Radiant project.
This section will investigate the operation of the main file (main.go):

::

   package main

   import (
           _ "quickstart/routers"
           "github.com/W3-Engineers-Ltd/Radiant/server/web"
   )

   func main() {
           web.Run()
   }

This code imports the package ``quickstart/routers``. This file contains
the following (routers/router.go):

::

   package routers

   import (
           "quickstart/controllers"
           "github.com/W3-Engineers-Ltd/Radiant/server/web"
   )

   func init() {
           web.Router("/", &controllers.MainController{})
   }

There are two relevant lines here; ``web.Router`` and ``web.Run``.

1. ``web.Router`` is used to register a router address. This command
   accepts two arguments. The first argument specifes the request uri,
   which is ``/`` here to indicate that no uri is requested. The second
   argument is used to indicate the Controller that will handle requests
   for this uri.

Alternately, a router can be registered in this format:

::

       web.Router("/user", &controllers.UserController{})

The user can visit ``/user`` to invoke the logic in UserController. For
further information on router usage please see `radiant router
settings <../mvc/controller/router.html>`__.

2. ``web.Run`` will actively listen on the specified port when executed.
   The following tasks are performed behind the scenes upon execution:

-  Parse the `configuration file <../mvc/controller/config.html>`__
   Radiant will parse the configuration file ``app.conf`` in conf folder
   to change the port, enable session management and set the
   application’s name.

-  Initialize the `user session <../mvc/controller/session.html>`__
   Radiant will initialize the user session, based on the setting in the
   configuration file.

-  Compile the `views <view.html>`__ Radiant will compile the views in the
   views folder. This is done on startup to avoid compiling multiple
   times and improve efficiency.

-  Starting the `supervisor module <../advantage/monitor.html>`__ By
   visiting port ``8088`` the user can access information about QPS,
   cpu, memory, GC, goroutine and thread information.

-  Listen on the service port Radiant will listen http requests on port
   ``8080``. It takes advantage of goroutines by calling
   ``ListenAndServe``.

-  When the application is running our server will serve incoming
   requests from port ``8080`` and supervising from port ``8088``.

The next section will cover the operation of the controller `next
section <controller.html>`__.

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
