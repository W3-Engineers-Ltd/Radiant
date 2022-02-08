Parsing Configuration Files
===========================

The config module is used for parsing configuration files, inspired by
``database/sql``. It supports ini, json, xml and yaml files. You can
install it by:

::

   go get github.com/W3-Engineers-Ltd/Radiant/core/config

If you want to parse xml or yaml, you should first install:

::

   go get -u github.com/W3-Engineers-Ltd/Radiant/core/config/xml

and then import:

::

   import _ "github.com/W3-Engineers-Ltd/Radiant/core/config/xml"

Remote configure middleware
===========================

Now we support ``etcd`` as the implementation.

Usage
=====

Using package
-------------

In v2.x, Radiant create a ``globalInstance``, so that users could use
``config`` module directly.

.. code:: go

   val, err := config.String("mykey")

Radiant use ``ini`` implementation and loads config from
``config/app.conf``.

If the file not found or got some error, Radiant outputs some warning
log.

Or you can initialize the ``globalInstance`` by:

.. code:: go

   _ import "github.com/W3-Engineers-Ltd/Radiant/core/config/toml"
   err := InitGlobalInstance("toml", "some config")
   // ...
   val, err := config.String("mykey")
   // ...

Create instance manually
------------------------

Initialize a parser object:

::

   iniconf, err := NewConfig("ini", "testini.conf")
   if err != nil {
       t.Fatal(err)
   }

Get data from parser:

::

   iniconf.String("appname")

Parser methods
~~~~~~~~~~~~~~

Here are the parser’s methods:

::

   // Configer defines how to get and set value from configuration raw data.
   type Configer interface {
       // support section::key type in given key when using ini type.
       Set(key, val string) error

       // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
       String(key string) (string, error)
       // get string slice
       Strings(key string) ([]string, error)
       Int(key string) (int, error)
       Int64(key string) (int64, error)
       Bool(key string) (bool, error)
       Float(key string) (float64, error)
       // support section::key type in key string when using ini and json type; Int,Int64,Bool,Float,DIY are same.
       DefaultString(key string, defaultVal string) string
       // get string slice
       DefaultStrings(key string, defaultVal []string) []string
       DefaultInt(key string, defaultVal int) int
       DefaultInt64(key string, defaultVal int64) int64
       DefaultBool(key string, defaultVal bool) bool
       DefaultFloat(key string, defaultVal float64) float64

       // DIY return the original value
       DIY(key string) (interface{}, error)

       GetSection(section string) (map[string]string, error)

       Unmarshaler(prefix string, obj interface{}, opt ...DecodeOption) error
       Sub(key string) (Configer, error)
       OnChange(key string, fn func(value string))
       SaveConfigFile(filename string) error
   }

Notice: 1. All ``Default*`` methods, default value will be returned if
key not found or go some error; 2. ``DIY`` returns original value. When
you want to use this method, you should be care of value’s type. 3.
``GetSection`` returns all configure items under the ``section``.
``section`` has different meaning in different implementation. 4.
``Unmarshaler`` try to decode the configs to ``obj``. And the parameter
``prefix`` is similar with ``section``. 5. ``Sub`` is similar to
``GetSection``, but ``Sub`` will wrap result as an ``Configer``
instance. 6. ``Onchange`` is used to listen change event. But most of
file-based implementations don’t support this method. 7.
``SaveConfigFile`` output configs to a file. 8. Some implementation
support key like ``a.b.c.d``, but not all implementations support it and
not all of them use ``.`` as separator.

Configuration sections
~~~~~~~~~~~~~~~~~~~~~~

The ini file supports configuration sections. You can get values inside
a section by using ``section::key``.

For example:

::

   [demo]
   key1 = "asta"
   key2 = "xie"

You can use ``iniconf.String("demo::key2")`` to get the value.

How to Obtain Environment Variables
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

After Pull Request “Support get environment variables in config #1636”
was merged into the code, radiant supports using environment variables
in the configuration file.

The format for this is ``${ENVIRONMENTVARIABLE}`` within the
configuration file which is equivalent to
``value = os.Getenv('ENVIRONMENTVARIABLE')``. Radiant will only check
for environment variables if the value begins with ``${`` and ends with
``}``.

Additionally, a default value can be configured for the case that there
is no environment variable set or the environment variable is empty.
This is accomplished by using the format ``${ENVVAR||defaultvalue}``,
for example ``${GOPATH||/home/asataxie/workspace/go}``. This ``||`` is
used to split environment values and default values. See
``/config/config_test.go`` in the `radiant
repo <https://github.com/W3-Engineers-Ltd/Radiant>`__ for more examples and edge
cases about how these environment variables and default values are
parsed.

For example:

::

   password = ${MyPWD}
   token = ${TOKEN||astaxie}
   user = ${MyUser||radiant}

If the environment variable ``$TOKEN`` is set, its value will be used
for the ``token`` configuration value and
``radiant.AppConfig.String("token")`` would return its value. If
``$TOKEN`` is not set, the value would then be the string ``astaxie``.

**Please note**: The environment variables are only read when the
configuration file is parsed, not when configuration item is obtained by
a function like ``radiant.AppConfig.String(string)``.

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