Installing Radiant
==================

You can use the classic Go way to install Radiant:

::

   go get github.com/W3-Engineers-Ltd/Radiant

Frequently asked questions:

-  git is not installed. Please install git for your system.
-  git https is not accessible. Please config local git and close https
   validation:

   ::

        git config --global http.sslVerify false

-  How can I install Radiant offline? There is no good solution for now.
   We will create packages for downloading and installing for future
   releases.

Upgrading Radiant
=================

You can upgrade Radiant through Go command or download and upgrade from
source code.

-  Through Go command (Recommended):

   ::

        go get -u github.com/W3-Engineers-Ltd/Radiant

-  Through source code: visit
   ``https://github.com/W3-Engineers-Ltd/Radiant`` and download the
   source code. Copy and overwrite to path
   ``$GOPATH/src/github.com/W3-Engineers-Ltd/Radiant``. Then run
   ``go install`` to upgrade Radiant:

   ::

        go install  github.com/W3-Engineers-Ltd/Radiant

**Upgrading Prior to 1.0:** The API of Radiant is stable after 1.0 and
compatible with every upgrade. If you are still using a version lower
than 1.0 you may need to configure your parameters based on the latest
API.


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