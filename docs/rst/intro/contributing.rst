Contributing to Radiant
=======================

Introduction
------------

Radiant is free and open source software, which means that anyone can
contribute to its development and progress under the Apache 2.0 License
(http://www.apache.org/licenses/LICENSE-2.0.html). Radiant’s source code
is hosted on github (https://github.com/W3-Engineers-Ltd/Radiant).

How can I become a contributor of Radiant?
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

You can fork, modify and then send a Pull Request to us. We will review
your code and give you feedback on your changes as soon as possible.

Pull Requests
-------------

The process for pull requests for new features and bug fixes are not the
same.

Bug fixes
~~~~~~~~~

Pull requests for bug fixes do not need to create an issue first. If you
have a solution to a bug, please describe your solution in detail in
your pull request.

Documentation improvements
~~~~~~~~~~~~~~~~~~~~~~~~~~

You can help improve the documentation by submitting a pull request to
the `radicaldoc <https://github.com/W3-Engineers-Ltd/Radiant-doc>`__
repository.

New features proposals
~~~~~~~~~~~~~~~~~~~~~~

Before you submit a pull request for a new feature, you should first
create an issue with ``[Proposal]`` in the title, describing the new
feature, as well as the implementation approach.

Proposals will be reviewed and discussed by the core contributors, and
can be adopted or potentially rejected.

Once a proposal is accepted, create an implementation of the new
features and submit it as a pull request. If the guidelines are not
followed the pull request will be rejected immediately.

Since Radiant follows the `Git
Flow <http://nvie.com/posts/a-successful-git-branching-model/>`__
branching model, ongoing development happens in the ``develop`` branch.
Therefore, please base your pull requests on the HEAD of the ``develop``
branch.

The git branches of Radiant
~~~~~~~~~~~~~~~~~~~~~~~~~~~

The master branch is relatively stable and the dev branch is for
developers. Here is a sample figure to show you how our branches work:

|image0|

For more information about the branching model:
http://nvie.com/posts/a-successful-git-branching-model/

A simple guideline for Git command
----------------------------------

You must have a github account, if not, please register one.

Fork 代码
~~~~~~~~~

1. Click
   `https://github.com/W3-Engineers-Ltd/Radiant <https://github.com/W3-Engineers-Ltd/Radiant>`__
2. Click “Fork” button which is on top right corner

Clone 代码
~~~~~~~~~~

We recommend using official repo as ``origin`` repo, and then add a
remote upstream to your repo.

If you already set SSH key, we recommend use SSH. The difference is
that, we don’t need to input the username and password to push changes.

Using SSH：

.. code:: bash

   git clone git@github.com:astaxie/radiant.git
   cd radiant
   git remote add upstream 'git@github.com:<your github username>/radiant.git'

Using HTTPS：

.. code:: bash

   git clone https://github.com/W3-Engineers-Ltd/Radiant.git
   cd radiant
   git remote add  'https://github.com/<you github username>/radiant.git'

The word ``upstream`` in command could be replaced with any word you
like.

fetch changes
~~~~~~~~~~~~~

Every time you want to something, you’d better fetch remote changes:

.. code:: bash

   git fetch

In this command, git only fetch ``origin`` repo。

If we want to fetch our remote repo changes:

.. code:: bash

   git fetch upstream

You can replace ``upstream`` with your repo name

create feature branch
~~~~~~~~~~~~~~~~~~~~~

我们在创建新的 feature 分支的时候，要先考虑清楚，从哪个分支切出来。
Before creating feature branch, we should think about choosing a branch
as base branch.

Assume that we want to merge the new feature to develop branch. In such
case:

.. code:: bash

   git checkout -b feature/my-feature origin/develop

Don’t forget to run ``git fetch`` before you create feature branch.

push commit
~~~~~~~~~~~

.. code:: bash

   git add .
   git commit
   git push upstream my-feature

make PR
~~~~~~~

Go to https://github.com/W3-Engineers-Ltd/Radiant, and make a Pull request

.. |image0| image:: ../images/git-branch-1.png

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