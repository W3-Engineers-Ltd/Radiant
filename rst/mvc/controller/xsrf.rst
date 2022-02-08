Cross-Site Request Forgery
==========================

XSRF, `Cross-Site Request
Forgery <http://en.wikipedia.org/wiki/Cross-site_request_forgery>`__, is
an important security concern for web development. Radiant has built in
XSRF protection which assigns each user a randomized cookie that is used
to verify requests. XSRF protection can be activated by setting
``EnableXSRF = true`` in the configuration file:

::

   EnableXSRF = true
   XSRFKey = 61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o
   XSRFExpire = 3600 // set cookie expire in 3600 seconds, default to 60 seconds if not specified

XSRF protection can also be enabled in the main application entry
function:

::

   web.BConfig.WebConfig.EnableXSRF = true
   web.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
   web.BConfig.WebConfig.XSRFExpire = 3600

When XSRF is enabled Radiant will set a cookie ``_xsrf`` for every user.
Radiant will refuse any ``POST``, ``PUT``, or ``DELETE`` request that
does not include this cookie. If XSRF protection is enabled a field must
be added to provide an ``_xsrf`` value to every form. This can be added
directly in the template with ``XSRFFormHTML()``.

A global expiration time should be set using ``web.XSRFExpire``. This
value can be also be set for individual logic functions:

.. code:: go

   func (this *HomeController) Get(){
       this.XSRFExpire = 7200
       this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
   }

**XSRF** works with HTTPS protocol. In Radiant 2.x, the cookie storing
XSRF token has two flag:
`secure <https://en.wikipedia.org/wiki/Secure_cookie>`__ and
`http-only <https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies>`__.

In Radiant 1.x (<=1.12.2), we don’t have this two flags, so it’s not
safe because attackers is able to steal the XSRF token.


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