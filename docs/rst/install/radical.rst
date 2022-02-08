Introduction to radical tool
============================

Radical tool is a project for rapid Radiant development. With radical
tool developers can create, auto compile and reload, develop, test, and
deploy Radiant applications quickly and easily.

Installing radical tool
-----------------------

Install radical tool with the following command:

::

   go get github.com/W3-Engineers-Ltd/Radical

Update the radical tool with the following command:

::

   go get -u github.com/W3-Engineers-Ltd/Radical

``radical`` is installed into ``GOPATH/bin`` by default. You need to add
``GOPATH/bin`` to your PATH, otherwise the ``radical`` command won’t
work.

radical tool commands
---------------------

Type ``radical`` in command line and the following messages with be
displayed:

::

   radical is a tool for managing Radiant framework.

   Usage:

       radical command [arguments]

   The commands are:

       new         Create a Radiant application
       run         run the app and start a Web server for development
       pack        Compress a Radiant project into a single file
       api         create an API Radiant application
       bale        packs non-Go files to Go source files
       version     show the radical, Radiant and Go version
       generate    source code generator
       migrate     run database migrations

Command ``new``
~~~~~~~~~~~~~~~

The ``new`` command can create a new web project. You can create a new
Radiant project by typing ``radical new <project name>`` under
``$GOPATH/src``. This will generate all the default project folders and
files:

::

   radical new myproject
   [INFO] Creating application...
   /gopath/src/myproject/
   /gopath/src/myproject/conf/
   /gopath/src/myproject/controllers/
   /gopath/src/myproject/models/
   /gopath/src/myproject/static/
   /gopath/src/myproject/static/js/
   /gopath/src/myproject/static/css/
   /gopath/src/myproject/static/img/
   /gopath/src/myproject/views/
   /gopath/src/myproject/conf/app.conf
   /gopath/src/myproject/controllers/default.go
   /gopath/src/myproject/views/index.tpl
   /gopath/src/myproject/main.go
   13-11-25 09:50:39 [SUCC] New application successfully created!

::

   myproject
   ├── conf
   │   └── app.conf
   ├── controllers
   │   └── default.go
   ├── main.go
   ├── models
   ├── routers
   │   └── router.go
   ├── static
   │   ├── css
   │   ├── img
   │   └── js
   ├── tests
   │   └── default_test.go
   └── views
       └── index.tpl

   8 directories, 4 files

Command ``api``
~~~~~~~~~~~~~~~

The ``new`` command is used for crafting new web applications. The
``api`` command is used to create new API applications. Here is the
result of running ``radical api project_name``:

::

   radical api apiproject
   create app folder: /gopath/src/apiproject
   create conf: /gopath/src/apiproject/conf
   create controllers: /gopath/src/apiproject/controllers
   create models: /gopath/src/apiproject/models
   create tests: /gopath/src/apiproject/tests
   create conf app.conf: /gopath/src/apiproject/conf/app.conf
   create controllers default.go: /gopath/src/apiproject/controllers/default.go
   create tests default.go: /gopath/src/apiproject/tests/default_test.go
   create models object.go: /gopath/src/apiproject/models/object.go
   create main.go: /gopath/src/apiproject/main.go

Below is the generated project structure of a new API application:

::

   apiproject
   ├── conf
   │   └── app.conf
   ├── controllers
   │   └── object.go
   │   └── user.go
   ├── docs
   │   └── doc.go
   ├── main.go
   ├── models
   │   └── object.go
   │   └── user.go
   ├── routers
   │   └── router.go
   └── tests
       └── default_test.go

Compare this to the ``radical new myproject`` command seen earlier. Note
that the new API application doesn’t have a ``static`` and ``views``
folder.

You can also create a model and controller based on the database schema
by providing database conn:

``radical api [appname] [-tables=""] [-driver=mysql] [-conn=root:@tcp(127.0.0.1:3306)/test]``

Command ``run``
~~~~~~~~~~~~~~~

The ``radical run`` command will supervise the file system of any
Radiant project using
`inotify <http://en.wikipedia.org/wiki/Inotify>`__. The results will
autocompile and display immediately after any modification in the
Radiant project folders.

::

   13-11-25 09:53:04 [INFO] Uses 'myproject' as 'appname'
   13-11-25 09:53:04 [INFO] Initializing watcher...
   13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/controllers)
   13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject/models)
   13-11-25 09:53:04 [TRAC] Directory(/gopath/src/myproject)
   13-11-25 09:53:04 [INFO] Start building...
   13-11-25 09:53:16 [SUCC] Build was successful
   13-11-25 09:53:16 [INFO] Restarting myproject ...
   13-11-25 09:53:16 [INFO] ./myproject is running...

Visting ``http://localhost:8080/`` with a web browser will display your
app running:

|image0|

After modifying the ``default.go`` file in the ``controllers`` folder,
the following output will be displayed in the command line:

::

   13-11-25 10:11:20 [EVEN] "/gopath/src/myproject/controllers/default.go": DELETE|MODIFY
   13-11-25 10:11:20 [INFO] Start building...
   13-11-25 10:11:20 [SKIP] "/gopath/src/myproject/controllers/default.go": CREATE
   13-11-25 10:11:23 [SKIP] "/gopath/src/myproject/controllers/default.go": MODIFY
   13-11-25 10:11:23 [SUCC] Build was successful
   13-11-25 10:11:23 [INFO] Restarting myproject ...
   13-11-25 10:11:23 [INFO] ./myproject is running...

Refresh the browser to show the results of the new modifications.

Command ``pack``
~~~~~~~~~~~~~~~~

The ``pack`` command is used to compress the project into a single file.
The compressed file can be deployed by uploading and extracting the zip
file to the server.

::

   radical pack
   app path: /gopath/src/apiproject
   GOOS darwin GOARCH amd64
   build apiproject
   build success
   exclude prefix:
   exclude suffix: .go:.DS_Store:.tmp
   file write to `/gopath/src/apiproject/apiproject.tar.gz`

The compressed file will be in the project folder:

::

   rwxr-xr-x  1 astaxie  staff  8995376 11 25 22:46 apiproject
   -rw-r--r--  1 astaxie  staff  2240288 11 25 22:58 apiproject.tar.gz
   drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 conf
   drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 controllers
   -rw-r--r--  1 astaxie  staff      509 11 25 22:31 main.go
   drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 models
   drwxr-xr-x  3 astaxie  staff      102 11 25 22:31 tests

Command ``bale``
~~~~~~~~~~~~~~~~

This command is currently only available to the developer team. It is
used to compress all static files in to a single binary file so that
they do not need to carry static files including js, css, images and
views when publishing the project. Those files will be self-extracting
with non-overwrite when the program starts.

Command ``version``
~~~~~~~~~~~~~~~~~~~

This command displays the version of ``radical``, ``radiant``, and
``go``.

.. code:: shell

   $ radical version
   radical   :1.2.2
   Radiant :1.4.2
   Go    :go version go1.3.3 darwin/amd64

This command try to output radiant’s version. It works well for GOPATH
mode. Radical finds radiant’s version from $GOPATH/src/astaxie/radiant
directory.

So when we use GOMOD mode, and we don’t download radiant’s source code,
Radical could not find the version’s information.

Command ``generate``
~~~~~~~~~~~~~~~~~~~~

This command will generate the routers by analyzing the functions in
controllers.

.. code:: shell

   radical generate scaffold [scaffoldname] [-fields=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
       The generate scaffold command will do a number of things for you.
       -fields: a list of table fields. Format: field:type, ...
       -driver: [mysql | postgres | sqlite], the default is mysql
       -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test
       example: radical generate scaffold post -fields="title:string,body:text"

   radical generate model [modelname] [-fields=""]
       generate RESTful model based on fields
       -fields: a list of table fields. Format: field:type, ...

   radical generate controller [controllerfile]
       generate RESTful controllers

   radical generate view [viewpath]
       generate CRUD view in viewpath

   radical generate migration [migrationfile] [-fields=""]
       generate migration file for making database schema update
       -fields: a list of table fields. Format: field:type, ...

   radical generate docs
       generate swagger doc file

   radical generate test [routerfile]
       generate testcase

   radical generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]
       generate appcode based on an existing database
       -tables: a list of table names separated by ',', default is empty, indicating all tables
       -driver: [mysql | postgres | sqlite], the default is mysql
       -conn:   the connection string used by the driver.
                default for mysql:    root:@tcp(127.0.0.1:3306)/test
                default for postgres: postgres://postgres:postgres@127.0.0.1:5432/postgres
       -level:  [1 | 2 | 3], 1 = models; 2 = models,controllers; 3 = models,controllers,router

Command ``migrate``
~~~~~~~~~~~~~~~~~~~

This command will run database migration scripts.

::

   radical migrate [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
       run all outstanding migrations
       -driver: [mysql | postgresql | sqlite], the default is mysql
       -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

   radical migrate rollback [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
       rollback the last migration operation
       -driver: [mysql | postgresql | sqlite], the default is mysql
       -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

   radical migrate reset [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
       rollback all migrations
       -driver: [mysql | postgresql | sqlite], the default is mysql
       -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

   radical migrate refresh [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]
       rollback all migrations and run them all again
       -driver: [mysql | postgresql | sqlite], the default is mysql
       -conn:   the connection string used by the driver, the default is root:@tcp(127.0.0.1:3306)/test

radical tool configuration
--------------------------

The file ``radical.json`` in the radical tool source code folder is the
Radiant configuration file. This file is still under development, but
some options are already available to use:

-  ``"version": 0``: version of file, for checking incompatible format
   version.
-  ``"go_install": false``: if you use a full import path like
   ``github.com/user/repo/subpkg`` you can enable this option to run
   ``go install`` and speed up you build processes.
-  ``"watch_ext": []``: add other file extensions to watch (only watch
   ``.go`` files by default). For example, ``.ini``, ``.conf``, etc.
-  ``"dir_structure":{}``: if your folder names are not the same as MVC
   classic names you can use this option to change them.
-  ``"cmd_args": []``: add command arguments for every start.
-  ``"envs": []``: set environment variables for every start.

.. |image0| image:: ../images/radicalrun.png


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