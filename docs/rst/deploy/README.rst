Releasing and Deploying
=======================

Development mode
~~~~~~~~~~~~~~~~

The application created by ``radical`` is in development mode by
default.

We can change the mode by:

::

   radiant.RunMode = "prod"

Or change it in conf/app.conf:

::

   runmode = prod

In development mode:

-  If you don’t have a views folder, it will show this kind of error:

   ::

        2013/04/13 19:36:17 [W] [stat views: no such file or directory]

-  Templates will load every time without cache.

-  If server throws error, the response will look like:

|image0|

.. _releasing-and-deploying-1:

Releasing and Deploying
~~~~~~~~~~~~~~~~~~~~~~~

The Go application is a bytecode file after compiling. You just need to
copy this file to the server and run it. But remember Radiant might also
include static files, configuration files and templates, so these three
folders also need to be copied to server while deploying.

::

   $ mkdir /opt/app/radicalpkg
   $ cp radicalpkg /opt/app/radicalpkg
   $ cp -fr views /opt/app/radicalpkg
   $ cp -fr static /opt/app/radicalpkg
   $ cp -fr conf /opt/app/radicalpkg

Here is the folder structure in ``/opt/app/radicalpkg``:

::

   .
   ├── conf
   │   ├── app.conf
   ├── static
   │   ├── css
   │   ├── img
   │   └── js
   └── views
       └── index.tpl
   ├── radicalpkg

Now we’ve copied our entire application to the server. Next step is
deploy it.

There are three ways to run it:

-  `Stand alone deploy <./radiant.html>`__
-  `Deploy with Supervisord <./supervisor.html>`__
-  `Deploy with Systemctl <./systemctl.html>`__

The application is exposed above, then usually we will have a nginx or
apache to serve pages and perform load balancing on our application.

-  `Deploy with Nginx <./nginx.html>`__
-  `Deploy with Apache <./apache.html>`__

.. |image0| image:: ./../images/dev.png


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