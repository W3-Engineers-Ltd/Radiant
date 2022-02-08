Deployment with Nginx
=====================

Go already has a standalone http server. But we still want to have nginx
to do more for us such as logging, CC attack and act as a static file
server because nginx performs well as a web server. So Go can just focus
on functionality and logic. We can also use the nginx proxy to deploy
multiple applications at the same time. Here is an example of two
applications that share port 80 but have different domains, and requests
are forwarding to different applications by nginx.

::

   server {
       listen       80;
       server_name  .a.com;

       charset utf-8;
       access_log  /home/a.com.access.log  main;

       location /(css|js|fonts|img)/ {
           access_log off;
           expires 1d;

           root "/path/to/app_a/static"
           try_files $uri @backend
       }

       location / {
           try_files /_not_exists_ @backend;
       }

       location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           proxy_set_header Host            $http_host;

           proxy_pass http://127.0.0.1:8080;
       }
   }

   server {
       listen       80;
       server_name  .b.com;

       charset utf-8;
       access_log  /home/b.com.access.log  main;

       location /(css|js|fonts|img)/ {
           access_log off;
           expires 1d;

           root "/path/to/app_b/static"
           try_files $uri @backend
       }

       location / {
           try_files /_not_exists_ @backend;
       }

       location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           proxy_set_header Host            $http_host;

           proxy_pass http://127.0.0.1:8081;
       }
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