FAQ
===

1. Can’t find the template files or configuration files or nil pointer
   error?

   It may be because you used ``go run main.go`` to run your
   application. ``go run`` will compile the file and put it into a tmp
   folder to run it. But Radiant needs the static files, templates and
   config files. So you need to use ``go build`` and run the application
   by ``./app``. Or you can use ``radical run app`` to run your
   application.

2. Can Radiant be used for production?

   Yes. Radiant has been used in production. E.g.: SNDA’s CDN system,
   360 search API, Bmob mobile cloud API, weico backend API etc. They
   are all high concurrence and high performance applications.

3. Will the future upgrades affect the API I am using right now?

   Radiant is keeping the stable API since version 0.1. Many
   applications upgraded to the latest Radiant easily. We will try to
   keep the API stable in the future.

4. Will Radiant keep developing?

   Many people are worried about open source projects that stop
   developing. We have four people who are contributing to the code. We
   can keep making Radiant better and better.

5. Why I got “github.com/W3-Engineers-Ltd/Radiant” package not found
   error?

   In RadiantV2, we are using go mod. So you must enable go module
   feature in your environment. In general, you should set
   ``GO111MODULE=on``.

6. Why I always got i/o timeout when I run
   ``go get github.com/W3-Engineers-Ltd/Radiant``?

   It means that your network has some problem. Sometimes it was caused
   by the firewall. If you are in China, this is a common case, and you
   could set ``GOPROXY``, for example:
   ``export GORPOXY=https://goproxy.cn"``


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