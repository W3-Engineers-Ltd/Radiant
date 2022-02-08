ORM Test
========

Testing code:

-  Model definition
   `models_test.go <https://github.com/W3-Engineers-Ltd/Radiant/client/orm/models_test.go>`__
-  Test cases
   `orm_test.go <https://github.com/W3-Engineers-Ltd/Radiant/client/orm/orm_test.go>`__

MySQL
^^^^^

.. code:: bash

   mysql -u root -e 'create database orm_test;'
   export ORM_DRIVER=mysql
   export ORM_SOURCE="root:@/orm_test?charset=utf8"
   go test -v github.com/W3-Engineers-Ltd/Radiant/core/client/orm

Sqlite3
^^^^^^^

.. code:: bash

   touch /path/to/orm_test.db
   export ORM_DRIVER=sqlite3
   export ORM_SOURCE=/path/to/orm_test.db
   go test -v github.com/W3-Engineers-Ltd/Radiant/core/client/orm

PostgreSQL
^^^^^^^^^^

.. code:: bash

   psql -c 'create database orm_test;' -U postgres
   export ORM_DRIVER=postgres
   export ORM_SOURCE="user=postgres dbname=orm_test sslmode=disable"
   go test -v github.com/W3-Engineers-Ltd/Radiant/core/client/orm

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