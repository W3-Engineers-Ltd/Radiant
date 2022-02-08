Flash Messages
==============

Flash messages are not related to Adobe/Macromedia Flash. They are
temporary messages between two logic blocks. All flash messages will be
cleared after the very next logic block. They are normally used to send
notes and error messages. Their use is suited for the
`Post/Redirect/Get <http://en.wikipedia.org/wiki/Post/Redirect/Get>`__
model. For example:

.. code:: go

   // Display settings message
   func (c *MainController) Get() {
       flash := web.ReadFromRequest(&c.Controller)
       if n, ok := flash.Data["notice"]; ok {
           // Display settings successful
           c.TplName = "set_success.html"
       } else if n, ok = flash.Data["error"]; ok {
           // Display error messages
           c.TplName = "set_error.html"
       } else {
           // Display default settings page
           this.Data["list"] = GetInfo()
           c.TplName = "setting_list.html"
       }
   }

   // Process settings messages
   func (c *MainController) Post() {
       flash := web.NewFlash()
       setting := Settings{}
       valid := Validation{}
       c.ParseForm(&setting)
       if b, err := valid.Valid(setting); err != nil {
           flash.Error("Settings invalid!")
           flash.Store(&c.Controller)
           c.Redirect("/setting", 302)
           return
       } else if b != nil {
           flash.Error("validation err!")
           flash.Store(&c.Controller)
           c.Redirect("/setting", 302)
           return
       }
       saveSetting(setting)
       flash.Notice("Settings saved!")
       flash.Store(&c.Controller)
       c.Redirect("/setting", 302)
   }

The logic of the code above is as follows:

1. Execute GET method. There’s no flash data, so display settings page.
2. After submission, execute POST and initialize a flash. If checking
   failed, set error flash message. If checking passed, save settings
   and set flash message to successful.
3. Redirect to GET request.
4. GET request receives flash message and executes the related logic.
   Show error page or success page based on the type of message.

``ReadFromRequest`` assigns messages to flash, so you can use it in your
template:

::

   {{.flash.error}}
   {{.flash.warning}}
   {{.flash.success}}
   {{.flash.notice}}

There are 4 different levels of flash messages:

-  Notice: Notice message
-  Success: Success message
-  Warning: Warning message
-  Error: Error message

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