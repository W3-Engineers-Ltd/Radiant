Command Line
============

You can call ``orm.RunCommand()`` after you registered models and
database(s) as follows:

.. code:: go

   func main() {
       // orm.RegisterModel...
       // orm.RegisterDataBase...
       ...
       orm.RunCommand()
   }

.. code:: bash

   go build main.go
   ./main orm
   # Get help by just run it.
   # If possible, go run main.go orm has the same result.

Database Schema Generation
--------------------------

.. code:: bash

   ./main orm syncdb -h
   Usage of orm command: syncdb:
     -db="default": DataBase alias
     -force=false: drop tables before create
     -v=false: verbose info

Use the ``-force=1`` flag to force drop tables and re-create.

Use the ``-v`` flag to print SQL statements.

--------------

Use program to create tables:

.. code:: go

   // Database alias.
   name := "default"

   // Drop table and re-create.
   force := true

   // Print log.
   verbose := true

   // Error.
   err := orm.RunSyncdb(name, force, verbose)
   if err != nil {
       fmt.Println(err)
   }

Even if you do not enable ``force`` mode, ORM also will auto-add new
fields and indexes, but you have to deal with delete operations
yourself.

Print SQL Statements
--------------------

.. code:: bash

   ./main orm sqlall -h
   Usage of orm command: syncdb:
     -db="default": DataBase alias name

Use database with alias ``default`` as default.

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