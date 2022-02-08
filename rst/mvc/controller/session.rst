Session control
===============

Radiant has a built-in session module that supports memory, file, mysql,
redis, couchbase, memcache and postgres as the save provider. Other
providers can be implemented according to the interface.

To use session in Radiant switch it on in the main function:

::

   web.BConfig.WebConfig.Session.SessionOn = true

Or it can be activated in the configuration file:

::

   SessionOn = true

After being switched on, session can used be used like this:

.. code:: go

   func (this *MainController) Get() {
       v := this.GetSession("asta")
       if v == nil {
           this.SetSession("asta", int(1))
           this.Data["num"] = 0
       } else {
           this.SetSession("asta", v.(int)+1)
           this.Data["num"] = v.(int)
       }
       this.TplName = "index.tpl"
   }

There are several useful methods to handle session:

-  SetSession(name string, value interface{})
-  GetSession(name string) interface{}
-  DelSession(name string)
-  SessionRegenerateID()
-  DestroySession()

The most commonly used methods are ``SetSession``, ``GetSession``, and
``DelSession``.

Custom logic can also be used:

::

   sess := this.StartSession()
   defer sess.SessionRelease()

sess object has following methods:

-  Set
-  Get
-  Delete
-  SessionID
-  SessionRelease
-  Flush

SetSession, GetSession and DelSession methods are recommended for
session operation as it will release resource automatically.

Here are some parameters used in the Session module:

-  SessionOn

   Enables Session. Default value is ``false``. Parameter name in
   configuration file: ``SessionOn``

-  SessionProvider Sets Session provider. Set to ``memory`` by default.
   ``File``, ``mysql`` and ``redis`` are also supported. Parameter name
   in configuration file: ``sessionprovider``.

-  SessionName Sets the cookie name. Session is stored in browser’s
   cookies by default. The default name is radiantsessionID. Parameter
   name in configuration file: ``sessionname``.

-  SessionGCMaxLifetime Sets the Session expire time. Default value is
   ``3600s``. Parameter name in configuration file:
   ``sessiongcmaxlifetime``.

-  SessionProviderConfig Sets the save path or connection string for
   file, mysql or redis. Default value is empty. Parameter name in
   configuration file: ``sessionproviderconfig``.

-  SessionHashFunc Sets the function used to generate sessionid. The
   default value is ``sha1``.

-  SessionCookieLifeTime Sets the cookie expire time. The cookie is used
   to store data in client.

Package Installation
--------------------

If you are not using Go modules, manual installation may be required.

*Note: Radiant >= 1.1.3 removed all dependencies*

.. code:: bash

   # Couchbase
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/couchbase

   # Ledis
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/ledis

   # Memcache
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/memcache

   # MySQL
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/mysql

   # Postgres
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/postgres

   # Redis
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis

   # Redis (cluster mode)
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis_cluster

   # Redis (sentinel)
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis_sentinel

   # SSDB
   go get -u github.com/W3-Engineers-Ltd/Radiant/server/web/session/ssdb

Example Usage
-------------

Couchbase
~~~~~~~~~

SessionProviderConfig is connection address using
`couchbase <https://github.com/couchbaselabs/go-couchbase>`__.

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
     _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/couchbase"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "couchbase"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "http://bucketname:bucketpass@myserver:8091/"
   }

File
~~~~

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "file"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "/tmp"
   }

Memcache
~~~~~~~~

SessionProviderConfig is the connection address using
`memcache <https://github.com/radiant/memcache>`__.

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
     _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/memcache"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "memcache"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:7080"
   }

MySQL
~~~~~

SessionProviderConfig is the connection address using
`go-sql-driver <https://github.com/go-sql-driver/mysql>`__.

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
     _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/mysql"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "mysql"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "username:password@protocol(address)/dbname?param=value"
   }

Postgres
~~~~~~~~

SessionProviderConfig is the connection address using
`postgres <https://github.com/lib/pq>`__.

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
     _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/postgres"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "postgresql"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
   }

Redis
~~~~~

SessionProviderConfig is the connection address using
`redigo <https://github.com/garyburd/redigo>`__.

.. code:: go

   // main.go
   package main

   import (
     "github.com/W3-Engineers-Ltd/Radiant/server/web"
     _ "github.com/W3-Engineers-Ltd/Radiant/server/web/session/redis"
   )

   func init() {
     web.BConfig.WebConfig.Session.SessionOn = true
     web.BConfig.WebConfig.Session.SessionProvider = "redis"
     web.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"
   }

Note:
-----

Session uses ``gob`` to register objects. When using a session engine
other than ``memory``, objects must be registered in session before they
can be used. Use ``gob.Register()`` to register them in ``init()``
function.

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