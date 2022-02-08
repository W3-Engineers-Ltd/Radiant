Live Monitor
============

We discussed the toolbox module before. It will listen on
``127.0.0.1:8088`` by default when the application is running. It can’t
be accessed from the internet but you can browse to it by other means
such as by nginx proxy.

         For security reason it is recommended that you block port 8088
         in firewall.

Monitor is disabled by default. You can enable it by adding the
following line in ``conf/app.conf`` file:

::

   EnableAdmin = true

Also you can change the port it listens on:

::

   AdminAddr = "localhost"
   AdminPort = 8088

Open browser and visit ``http://localhost:8088/`` you will see
``Welcome to Admin Dashboard``.

Requests statistics
-------------------

Browse to ``http://localhost:8088/qps`` and you will see the following:

|image0|

Performance profiling
---------------------

You can also see the information for ``goroutine``, ``heap``,
``threadcreate``, ``block``, ``cpuprof``, ``memoryprof``, ``gc summary``
and do profiling.

Healthcheck
-----------

You need to manually register the healthcheck logic to see the status of
the healthcheck from ``http://localhost:8088/healthcheck``

Tasks
-----

You can add task in your application and check the task status or
trigger it manually.

-  Check task status: ``http://localhost:8088/task``
-  Run task manually:
   ``http://localhost:8088/runtask?taskname=task_name``

Config Status
-------------

After the development of the application, we may also want to know the
config when the application is running. Radiant’s Monitor also provided
this feature.

-  Show all configurations:
   ``http://localhost:8088/listconf?command=conf``
-  Show all routers: ``http://localhost:8088/listconf?command=router``
-  Show all filters: ``http://localhost:8088/listconf?command=filter``

.. |image0| image:: ../images/monitoring.png

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