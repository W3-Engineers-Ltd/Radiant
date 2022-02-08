Getting started
===============

Installation
------------

Radiant contains sample applications to help you learn and use the
Radiant app framework.

You will need a `Go <https://golang.org>`__ 1.1+ installation for this
to work.


::

   go get -u github.com/W3-Engineers-Ltd/Radiant
   go get -u github.com/W3-Engineers-Ltd/Radical

For convenience, you should add ``$GOPATH/bin`` to your ``$PATH``
environment variable. Please make sure you have already set the
``$GOPATH`` environment variable.

If you haven’t set ``$GOPATH`` add it to the shell you’re using
(~/.profile, ~/.zshrc, ~/.cshrc or any other).

For example ``~/.zsh``

::

   echo 'export GOPATH="$HOME/go"' >> ~/.zsh

If you have already set ``$GOPATH``

::

   echo 'export PATH="$GOPATH/bin:$PATH"' >> ~/.profile # or ~/.zshrc, ~/.cshrc, whatever shell you use
   exec $SHELL

Want to quickly see how it works? Then just set things up like this:

::

   cd $GOPATH/src
   radical new hello
   cd hello
   radical run

Windows users：

::

   cd %GOPATH%/src
   radical new hello
   cd hello
   radical run

These commands help you:

1. Install Radiant into your ``$GOPATH``.
2. Install the Radical tool in your computer.
3. Create a new application called ``hello``.
4. Start hot compile.

Once it’s running, open a browser to http://localhost:8080/.

Simple example
--------------

The following example prints ``Hello world`` to your browser, it shows
how easy it is to build a web application with radiant.

::

   package main

   import (
       "github.com/W3-Engineers-Ltd/Radiant/server/web"
   )

   type MainController struct {
       web.Controller
   }

   func (this *MainController) Get() {
       this.Ctx.WriteString("hello world")
   }

   func main() {
       web.Router("/", &MainController{})
       web.Run()
   }

Save file as ``hello.go``, build and run it:

::

   $ go build -o hello hello.go
   $ ./hello

Open http://127.0.0.1:8080 in your browser and you will see
``hello world``.

What is happening in the scenes of the above example?

1. We import package ``github.com/W3-Engineers-Ltd/Radiant/server/web``.
   As we know, Go initializes packages and runs init() in every package
   (`more
   details <https://github.com/Unknwon/build-web-application-with-golang_EN/blob/master/eBook/02.3.html#main-function-and-init-function>`__),
   so Radiant initializes the ``RadicalApp`` application at this time.
2. Define the controller. We define a struct called ``MainController``
   with an anonymous field ``web.Controller``, so the ``MainController``
   has all methods that ``web.Controller`` has.
3. Define some RESTful methods. Due to the anonymous field above,
   ``MainController`` already has ``Get``, ``Post``, ``Delete``, ``Put``
   and other methods, these methods will be called when user sends a
   corresponding request (e.g. the ``Post`` method is called to handle
   requests using POST. Therefore, after we overloaded the ``Get``
   method in ``MainController``, all GET requests will use that method
   in ``MainController`` instead of in ``web.Controller``.
4. Define the main function. All applications in Go use ``main`` as
   their entry point like C does.
5. Register routers. This tells Radiant which controller is responsible
   for specific requests. Here we register ``MainController`` for ``/``,
   so all requests to ``/`` will be handed by ``MainController``. Be
   aware that the first argument is the path and the second one is
   pointer to the controller you want to register.
6. Run the application on port 8080 as default, press ``Ctrl+c`` to
   exit.

Following are shortcut ``.bat`` files for Windows users:

Create files ``step1.install-radical.bat`` and
``step2.new-radiant-app.bat`` under ``%GOPATH%/src``.

``step1.install-radical.bat``:

::

   set GOPATH=%~dp0..
   go build github.com\radiant\radical
   copy radical.exe %GOPATH%\bin\radical.exe
   del radical.exe
   pause

``step2.new-radiant-app.bat``:

::

   @echo Set value of APP same as your app folder
   set APP=coscms.com
   set GOPATH=%~dp0..
   set BEE=%GOPATH%\bin\radical
   %BEE% new %APP%
   cd %APP%
   echo %BEE% run %APP%.exe > run.bat
   echo pause >> run.bat
   start run.bat
   pause
   start http://127.0.0.1:8080

Click those two file in order will quick start your Radiant tour. And
just run ``run.bat`` in the future.

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