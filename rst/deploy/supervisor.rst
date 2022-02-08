Supervisord
===========

Supervisord is a very useful process manager implemented in Python.
Supervisord can change your non-daemon application into a daemon
application. The application needs to be a non-daemon app. So if you
want to use Supervisord to manage nginx, you need to set daemon off to
run nginx in non-daemon mode.

Install Supervisord
-------------------

1. install setuptools

   ::

       wget https://pypi.python.org/packages/2.7/s/setuptools/setuptools-0.6c11-py2.7.egg

       sh setuptools-0.6c11-py2.7.egg 

       easy_install supervisor

       echo_supervisord_conf >/etc/supervisord.conf

       mkdir /etc/supervisord.conf.d

2. config ``/etc/supervisord.conf``

   ::

       [include]
       files = /etc/supervisord.conf.d/*.conf

3. Create new application to be managed

   ::

       cd /etc/supervisord.conf.d
       vim radicalpkg.conf

   Configurations：

   ::

       [program:radicalpkg]
       directory = /opt/app/radicalpkg
       command = /opt/app/radicalpkg/radicalpkg
       autostart = true
       startsecs = 5
       user = root
       redirect_stderr = true
       stdout_logfile = /var/log/supervisord/radicalpkg.log

Supervisord Manage
------------------

Supervisord provides two commands, supervisord and supervisorctl:

-  supervisord: Initialize Supervisord, run configed processes
-  supervisorctl stop programxxx: Stop process programxxx. programxxx is
   configed name in [program:radicalpkg]. Here is radicalpkg.
-  supervisorctl start programxxx: Run the process.
-  supervisorctl restart programxxx: Restart the process.
-  supervisorctl stop groupworker: Restart all processes in group
   groupworker
-  supervisorctl stop all: Stop all processes. Notes: start, restart and
   stop won’t reload the latest configs.
-  supervisorctl reload: Reload the latest configs.
-  supervisorctl update: Reload all the processes who’s config has
   changed.

..

         Notes: The processes stopped by ``stop`` manually won’t restart
         after reload or update.

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