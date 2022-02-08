Cache Module
============

Radiant’s cache module is used for caching data, inspired by
``database/sql``. It supports four cache providers: file, memcache,
memory and redis. You can install it by:

::

   github.com/W3-Engineers-Ltd/Radiant/client/cache

If you use the ``memcache`` or ``redis`` provider, you should first
install:

::

   go get -u github.com/W3-Engineers-Ltd/Radiant/client/cache/memcache

and then import:

::

   import _ "github.com/W3-Engineers-Ltd/Radiant/client/cache/memcache"

Basic Usage
-----------

First step is importing the package:

::

   import (
       "github.com/W3-Engineers-Ltd/Radiant/client/cache"
   )

Then initialize a global variable object:

::

   bm, err := cache.NewCache("memory", `{"interval":60}`)

Then we can use ``bm`` to modify the cache:

::

   bm.Put("astaxie", 1, 10*time.Second)
   bm.Get("astaxie")
   bm.IsExist("astaxie")
   bm.Delete("astaxie")

Provider Settings
-----------------

Here is how to configure the four providers:

-  memory

   ``interval`` stands for GC time, which means the cache will be
   cleared every 60s:

   ::

        {"interval":60}

-  file

   ::

        {"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":2,"EmbedExpiry":120}

-  redis

   redis uses
   `redigo <https://github.com/garyburd/redigo/tree/master/redis>`__

   ::

        {"key":"collectionName","conn":":6039","dbNum":"0","password":"thePassWord"}

   -  key: the Redis collection name
   -  conn: Redis connection info
   -  dbNum: Select the DB having the specified zero-based numeric
      index.
   -  password: the password for connecting password-protected Redis
      server

-  memcache

   memcache uses
   `vitess <http://code.google.com/p/vitess/go/memcache>`__

   ::

        {"conn":"127.0.0.1:11211"}

Creating your own provider
--------------------------

The cache module uses the Cache interface, so you can create your own
cache provider by implementing this interface and registering it.

.. code:: go

   type Cache interface {
       // Get a cached value by key.
       Get(ctx context.Context, key string) (interface{}, error)
       // GetMulti is a batch version of Get.
       GetMulti(ctx context.Context, keys []string) ([]interface{}, error)
       // Set a cached value with key and expire time.
       Put(ctx context.Context, key string, val interface{}, timeout time.Duration) error
       // Delete cached value by key.
       Delete(ctx context.Context, key string) error
       // Increment a cached int value by key, as a counter.
       Incr(ctx context.Context, key string) error
       // Decrement a cached int value by key, as a counter.
       Decr(ctx context.Context, key string) error
       // Check if a cached value exists or not.
       IsExist(ctx context.Context, key string) (bool, error)
       // Clear all cache.
       ClearAll(ctx context.Context) error
       // Start gc routine based on config string settings.
       StartAndGC(config string) error
   }

Register your provider:

::

   func init() {
       cache.Register("myowncache", NewOwnCache())
   }

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