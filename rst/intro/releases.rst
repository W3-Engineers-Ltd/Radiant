radiant 2.0.2
=============

`Change
Log <https://github.com/W3-Engineers-Ltd/Radiant/releases/tag/v2.0.2-beta.1>`__

radiant 2.0.1
=============

When we release v2.0.0, something wrong and then we re-release v2.0.0.

But the checksum in sum.golang.cn is immutable, so we release this
version

radiant 2.0.0
=============

Refactor
~~~~~~~~

1. Support the new project structure. We split the whole framework into
   four parts: 1.1 server: including web module. 1.2 client: including
   ORM, cache, httplib modules. 1.3 task: supporting timed tasks and
   other cyclic tasks. 1.4 core: including validation, config, logs and
   admin modules.
2. Add ``adapter`` module which is used to reduce the effort of
   upgrading Radiant from v1.x to v2.x
3. Add ``context.Context`` for ``cache``, ``httplib``, ``session``,
   ``task``, ``ORM`` modules’ API.
4. Add ``error`` as a return value for ``cache``, ``httplib``,
   ``session``, ``task``. Now users are able to know more info about
   what happen inside Radiant.
5. Decouple modules from each other. All modules only depend on ``core``
   package.
6. Support tracing, metrics in ORM, web, httplib modules.
7. Introduce ``filter-chain`` patter to support AOP.

Feature:
~~~~~~~~

1.  Allow Healthcheck endpoint return JSON for Kubernetes.
    `4055 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4055>`__
2.  Support ``ClientAuth`` for TLS.
    `4116 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4116>`__
3.  ``orm.RawSeter`` supports ``orm.Fielder``.
    `4191 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4191>`__
4.  Add a new MySQL operator for strict case-sensitive query.
    `4198 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4198>`__
5.  Using ``filter-chain`` pattern in ``orm`` module. Support
    opentracing and prometheus by using filter.
    `4141 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4141>`__
6.  Support ``prometheus`` filter for ``httplib`` module.
    `4145 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4145>`__
7.  Add additional options to redis session prov.
    `4137 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4137>`__
8.  Support default value filter for ``orm`` module.
    `4156 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4156>`__
9.  Add methods ``Unmarshaler``, ``Sub``, ``OnChange`` for ``Configer``
    module. `4175 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4175>`__
10. Custom Log Formatter.
    `4174 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4174>`__,
    `4179 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4179>`__,
    `4188 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4188>`__
11. Support the time precision for time.Time type.
    `4186 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4186>`__
12. Support ``etcd`` in Config module.
    `4195 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4195>`__
13. Optimize rawSet.QueryRows to avoid many unnecessary calls to
    parseStructTag.
    `4210 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4210>`__
14. Allow users to ignore some table when run ORM commands.
    `4211 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4211>`__
15. PostgresQueryBuilder
    `4205 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4205>`__
16. Provides a powerful ``LogFormatter`` implementation
    PatternLogFormatter.\ `4229 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4229>`__
17. Support custom ES index name.
    `4233 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4233>`__
18. Support multiple web servers.
    `4234 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4234>`__
19. Support toml config.
    `4262 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4262>`__
20. Using unmarshaler to parse config in web module.
    `4266 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4266>`__
21. Add MaxUploadFile to provide more safety uploading control.
    `4275 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4275>`__
22. Support using json string to init session.
    `4277 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4277>`__
23. Support global instance for config module.
    `4278 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4278>`__
24. Session: adds CookieSameSite to ManagerConfig.
    `4226 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4226>`__

Fix:
~~~~

1.  Fix reconnection bug in logs/conn.go.
    `4056 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4056>`__
2.  Return 413 when request payload too large.
    `4058 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4058>`__
3.  Fix ``index out of range`` in session module when ``len(sid) < 2``.
    `4068 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4068>`__
4.  Fix concurrent issue of context/input Query method.
    `4066 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4066>`__
5.  Allow using environment variable to specific the config file.
    `4111 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4111>`__
6.  XSRF add secure and http only flag.
    `4126 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4126>`__
7.  Fix temporary create failed on Windows
    `4244 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4244>`__
8.  Fix:return error after inserting data when primary key is string.
    `4150 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4150>`__
9.  Fix the bug that Fielder’s SetRaw is not called when calling
    orm.Raw() to query from database.
    `4160 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4160>`__
10. Fix: return error when calling ``InsertOrUpdate`` is successful with
    string primary key.
    `4158 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4158>`__
11. Fix the problem that the nested structure of queryRow() cannot
    assign values
    `4173 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4173>`__
12. Empty field in validator.Error when label struct tag is not
    declared. `4225 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4225>`__
13. Fix deadlock in task module.
    `4246 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4246>`__
14. Fix: form entity too large casue run out of memory.
    `4272 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4272>`__

Doc:
~~~~

1. Fix typo. `4251 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4251>`__,
   `4135 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4135>`__,
   `4107 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4107>`__

radiant 1.12.3
==============

.. _feature-1:

Feature:
~~~~~~~~

1. Allow Healthcheck endpoint return JSON for Kubernetes.
   `4055 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4055>`__
2. Support ``ClientAuth`` for TLS.
   `4116 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4116>`__
3. ``orm.RawSeter`` support ``orm.Fielder``.
   `4191 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4191>`__
4. Add a new MySQL operator for strict case sensitive query.
   `4198 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4198>`__

.. _fix-1:

Fix:
~~~~

1.  Fix reconnection bug in logs/conn.go.
    `4056 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4056>`__
2.  Return 403 when request payload too large.
    `4058 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4058>`__
3.  Fix race condition for Prepare Statement cache.
    `4061 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4061>`__
4.  Fix ``index out of range`` in session module when ``len(sid) < 2``.
    `4068 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4068>`__
5.  Fix concurrent issue of context/input Query method.
    `4066 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4066>`__
6.  Allow using environment variable to specific the config file.
    `4111 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4111>`__
7.  XSRF add secure and http only flag.
    `4126 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4126>`__
8.  Fix temporary create failed on Windows
    `4244 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4244>`__
9.  Session: adds CookieSameSite to ManagerConfig.
    `4226 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4226>`__
10. Make stmt cache smaller to avoid ``too many statement`` error.
    `4261 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4261>`__

radiant 1.12.2
==============

1.  Fix old process didn’t exist when graceful restart in radiant 1.12.0
    `#4005 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4005>`__
2.  Enhance: Print stack while orm abnormally exit
    `#3743 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3743>`__
3.  Enhance: Replacing lock with read lock in GetMapData
    `#3803 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3803>`__
4.  Fix: Get the real location of the log directory if the path is
    symbolic path
    `#3818 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3818>`__
5.  Fix: Cache, context, session: add lock to fix inconsistent field
    protection `#3922 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3922>`__
6.  Fix: Encoded url(with slash) router mismatch problem
    `#3943 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3943>`__
7.  Fix: genRouterCode method generate wrong codes
    `#3981 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3981>`__
8.  Enhance: Using LRU algorithm, ignoring big file and using max cache
    size to reduce the memory usage of file cache
    `#3984 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3984>`__
9.  Fix: Set max DB connections
    `#3985 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3985>`__
10. Fix: SQLite don’t support SELECT … FOR UPDATE
    `#3992 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3992>`__
11. Enhance: Add Transfer-Encoding header in httplib’s PostFile method
    `#3993 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3993>`__
12. Enhance: Support bit operation in ORM
    `#3994 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3994>`__
13. Fix: net/http Middleware set via RunWithMiddleware or
    App.Run(middleware) doesn’t work when “BConfig.Listen.Graceful” is
    set to true `#3995 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3995>`__
14. Fix: Empty field in validator.Error when label struct tag is not
    declared `#4001 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4001>`__
15. Fix: panic: send on closed channel after closing logger
    `#4004 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4004>`__
16. Enhance: Store RouterPattern before filter execute
    `#4007 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4007>`__
17. Fix: Using HTMLEscapeString in adminui.go to avoid XSS attack
    `#4018 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4018>`__
18. Fix: Process not closed when graceful set to true
    `#4005 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4005>`__
19. Enhance: Use scan instead of keys in redis
    `#4016 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4016>`__
20. Feature: Support prometheus
    `#4021 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4021>`__
21. Fix: Can’t create more than max_prepared_stmt_count statements
    `#4025 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4025>`__
22. Enhance: Support more mobile number pattern
    `#4027 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4027>`__
23. Fix: Can’t set section name
    `#4027 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4027>`__
24. Fix: strings.Repeat panic in orm/db.go
    `#4032 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4032>`__
25. Enhance: Make redis client idle timeout configurable
    `#4033 <https://github.com/W3-Engineers-Ltd/Radiant/pull/4033>`__

radiant 1.10.0
==============

1.  Update log.go add GetLevel Function to Log
    `#2970 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2970>`__
2.  Fix a typo “conflict”
    `#2971 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2971>`__
3.  Bug on private fields
    `#2978 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2978>`__
4.  Fix access log console unexpected ‘:raw-latex:`\n`’ at end of each
    log. `#2976 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2976>`__
5.  Fix Documentation for HTTP status codes descriptions.
    `#2992 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2992>`__
6.  Redis cache: make MaxIdle configurable
    `#3004 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3004>`__
7.  Update: Fix migration generate SQL
    `#3017 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3017>`__
8.  Handle pointer validation
    `#3046 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3046>`__
9.  Fix the issue TaseCase TestFormatHeader_0 is failed
    `#3066 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3066>`__
10. Fix BEEGO_RUNMODE
    `#3064 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3064>`__
11. Swagger: Allow example values with different types, allow example
    for enum. `#3085 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3085>`__
12. Fix the bug: unable to add column with ALTER TABLE
    `#2999 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2999>`__
13. Set default Radiant RunMode to production
    `#3076 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3076>`__
14. Fix typo `#3103 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3103>`__
15. In dev mode, template parse error cause program lock
    `#3126 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3126>`__
16. Amend a very minor typo in a variable name
    `#3115 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3115>`__
17. When log maxSize set big int，FileWrite Init fail
    `#3109 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3109>`__
18. Change github.com/garyburd/redigo to newest branch
    github.com/gomodul…
    `#3100 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3100>`__
19. ExecElem.FieldByName as local variable
    `#3039 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3039>`__
20. Allow log prefix
    `#3145 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3145>`__
21. Refactor yaml config for support multilevel
    `#3127 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3127>`__
22. Create redis_cluster.go
    `#3175 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3175>`__
23. Add field comment on create table
    `#3190 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3190>`__
24. Update: use PathEscape replace QueryEscape
    `#3200 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3200>`__
25. Update gofmt
    `#3206 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3206>`__
26. Update: Htmlquote Htmlunquote
    `#3202 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3202>`__
27. Add ‘FOR UPDATE’ support for querySet
    `#3208 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3208>`__
28. Debug stringsToJSON
    `#3171 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3171>`__
29. Fix defaut value bug, and add config for maxfiles
    `#3185 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3185>`__
30. Fix: correct MaxIdleConnsPerHost value to net/http default 100.
    `#3230 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3230>`__
31. Fix: When multiply comment routers on one func
    `#3217 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3217>`__
32. Send ErrNoRows if the query returns zero rows … in method orm_query…
    `#3247 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3247>`__
33. Fix typo `#3245 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3245>`__
34. Add session redis IdleTimeout config
    `#3239 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3239>`__
35. Fix the wrong status code in prod
    `#3226 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3226>`__
36. Add method to set the data depending on the accepted
    `#3182 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3182>`__
37. Fix Unexpected EOF bug in staticfile
    `#3152 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3152>`__
38. Add code style for logs README
    `#3146 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3146>`__
39. Fix response http code
    `#3142 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3142>`__
40. Improve access log
    `#3141 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3141>`__
41. Auto create log dir
    `#3105 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3105>`__
42. Html escape before display path, avoid xss
    `#3022 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3022>`__
43. Acquire lock when access config data
    `#3250 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3250>`__
44. Fix orm fields SetRaw function error judge problem
    `#2985 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2985>`__
45. Fix template rendering with automatic mapped parameters (see #2979)
    `#2981 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2981>`__
46. Fix the model can not be registered correctly on Ubuntu 32bit
    `#2997 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2997>`__
47. Feature/yaml
    `#3181 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3181>`__
48. Feature/autocert
    `#3249 <https://github.com/W3-Engineers-Ltd/Radiant/pull/3249>`__

radiant 1.9.0
=============

1.  Fix the new repo address for casbin
    `#2654 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2654>`__
2.  Fix cache/memory fatal error: concurrent map iteration and map write
    `#2726 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2726>`__
3.  AddAPPStartHook func modify
    `#2724 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2724>`__
4.  Fix panic: sync: negative WaitGroup counter
    `#2717 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2717>`__
5.  incorrect error rendering (wrong status)
    `#2712 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2712>`__
6.  validation: support int64 int32 int16 and int8 type
    `#2728 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2728>`__
7.  validation: support required option for some struct tag valids
    `#2741 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2741>`__
8.  Fix big form parse issue
    `#2725 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2725>`__
9.  File log add RotatePerm
    `#2683 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2683>`__
10. Fix Oracle placehold
    `#2749 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2749>`__
11. Supported gzip for req.Header has Content-Encoding: gzip
    `#2754 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2754>`__
12. Add new Database Migrations
    `#2744 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2744>`__
13. Radiant auto generate sort ControllerComments
    `#2766 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2766>`__
14. added statusCode and pattern to FilterMonitorFunc
    `#2692 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2692>`__
15. fix the bugs in the “ParseBool” function in the file of config.go
    `#2740 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2740>`__

radical 1.9.0
-------------

1. Added MySQL year data type
   `#443 <https://github.com/W3-Engineers-Ltd/Radiant/pull/443>`__
2. support multiple http methods
   `#445 <https://github.com/W3-Engineers-Ltd/Radiant/pull/445>`__
3. The DDL migration can now be generated by adding a -ddl and a proper
   “alter” or “create” as argument value.
   `#455 <https://github.com/W3-Engineers-Ltd/Radiant/pull/455>`__
4. Fix: docs generator skips everything containing ‘vendor’
   `#454 <https://github.com/W3-Engineers-Ltd/Radiant/pull/454>`__
5. get these tables information in custom the option
   `#441 <https://github.com/W3-Engineers-Ltd/Radiant/pull/441>`__
6. read ref(pk) `#444 <https://github.com/W3-Engineers-Ltd/Radiant/pull/444>`__
7. Add command radical server to server static folder.

radiant 1.7.1
=============

New features: 1. Added IP for access log
`#2156 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2156>`__ 2.
ReadForUpdate or ORM
`#2158 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2158>`__ 3. Parameters
binding supports
form，columns[0].Data=foo&columns[1].Data=bar&columns[2].Data=baz
`#2111 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2111>`__ 4. Added
``radiant.BConfig.RecoverFunc`` for custom recover method.
`#2004 <https://github.com/W3-Engineers-Ltd/Radiant/issues/2004>`__ 5. memcache
cache supports byte and string. So as to cache struct by
gob\ `#1521 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1521>`__ 6. ORM
delete by condition.
`#1802 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1802>`__ 7. swagger
doc supports yaml
`#2162 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2162>`__ 8. Added
RunController and RunMethod for customized router rules
`#2017 <https://github.com/W3-Engineers-Ltd/Radiant/issues/2017>`__

Bug fixes: 1. No / added while visiting static folder who contains
index.html. For example visit /swagger won’t redirect to /swagger/ so
that relitive css and js won’t be access.
`#2142 <https://github.com/W3-Engineers-Ltd/Radiant/issues/2142>`__ 2. Time in
Radiant admin UI displayed alphabetically other than ordered by us or
ms. `#1877 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1877>`__ 3. Crash
while captcha generates by custom height and width.
`#2161 <https://github.com/W3-Engineers-Ltd/Radiant/issues/2161>`__ 4. Panic
while empty body requested with DELETE method when CopyBody enabled.
`#1656 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1656>`__

radiant 1.7.0
=============

New features: 1. Improved Filter speed by 7.5+ times
`#1799 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1799>`__ 2. Multiple
level for Gzip compression
`#1808 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1808>`__ 3. Negative
numbers for ORM PK
`#1810 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1810>`__ 4. Custom
auto-increasing ID for ORM
`#1826 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1826>`__ 5. Improved
Context file downloading: check file existence before
download\ `#1827 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1827>`__ 6.
``GetLogger`` method for log module
`#1832 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1832>`__

::

   package main

   import "github.com/W3-Engineers-Ltd/Radiant/logs"

   func main() {
       logs.Warn("this is a warn message")

       l := logs.GetLogger("HTTP")
       l.Println("this is a message of http")

       logs.GetLogger("orm").Println("this is a message of orm")

       logs.Debug("my book is bought in the year of ", 2016)
       logs.Info("this %s cat is %v years old", "yellow", 3)
       logs.Error(1024, "is a very", "good", 2.5, map[string]int{"id": 1})
       logs.Critical("oh my god")
   }

|image0| 7. Log for session if error occurred.
`#1833 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1833>`__ 8. Public
methods for logs: ``EnableFuncCallDepth`` and ``SetLogFuncCallDepth``
for setting function call level.
`#1837 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1837>`__ 9. Use
``go run`` to run radiant project
`#1840 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1840>`__ 10. Added
``ExecuteTemplate`` method which is used to access template other than
use map since map is not safe for concurrent reading and writing.
`#1848 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1848>`__ 11. ``time``
type for ORM field
`#1856 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1856>`__ 12. ORM One
only fetch one record
`#1874 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1874>`__ 13. ORM suports
json jsonb type `#1875 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1875>`__
14. ORM uses text type by default.
`#1879 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1879>`__ 15. session
configurations: ``EnableSidInHttpHeader`` ``EnableSidInUrlQuery``
``SessionNameInHttpHeader`` let user pass sid in http header or in URL.
`#1897 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1897>`__ 16. Shorten
fileanme of auto-generated router file name.
`#1924 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1924>`__ 17. Complex
template engine. ace jade
`#1940 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1940>`__

::

   radiant.AddTemplateEngine("ace", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
           aceOptions := &ace.Options{DynamicReload: true, FuncMap: funcs}
           aceBasePath := filepath.Join(root, "base/base")
           aceInnerPath := filepath.Join(root, strings.TrimSuffix(path, ".ace"))

           tpl, err := ace.Load(aceBasePath, aceInnerPath, aceOptions)
           if err != nil {
               return nil, fmt.Errorf("error loading ace template: %v", err)
           }

           return tpl, nil
       })

`#1940 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1940>`__ 18. session
suports ssdb `#1953 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1953>`__
19. RenderForm supports required
`#1993 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1993>`__ 20. Beautified
radiant logs `#1997 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1997>`__
|image1| 21. ORM suports ``time.Time`` pointer in struct
`#2006 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2006>`__ 22.
``TplPrefix`` in Controller for setting prefix folder in baseController
`#2030 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2030>`__ 23. js function
checking in jsonb to avoid non-exist methods.
`#2045 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2045>`__ 24.
``InsertOrUpdate`` method in ORM
`#2053 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2053>`__ 25. Filter
method added parameter for resetting parameters. Because when using
``radiant.InsertFilter("*", radiant.BeforeStatic, RedirectHTTP)``
parameter will be assigned to ``:splat`` which will affect other useful
routers. `#2085 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2085>`__ 26.
session initialized by object other than json. *It might have issue for
the projects use session module separately.*
`#2096 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2096>`__ 27. Upgraded
Swagger to 2.0. The code generated now doesn’t rely on API. radiant
generat swagger.json directly.

bugfix: 1. ``/m`` redirect to ``/m/`` automatically in static reouters.
`#1792 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1792>`__ 2. Parsing
config file error while testing
`#1794 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1794>`__ 3. Race
condition while rotate file.
`#1803 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1803>`__ 4. Fixed
multiple response.WriteHeader calls error.
`#1805 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1805>`__ 5. Fixed panic
if primary key is uint in ORM
`#1828 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1828>`__ 6. Fixed panic
if current time is less than 2000 while rotate logs. `# <>`__ 7. Fixed
XSRF reuse caused by context
reuse.\ `#1863 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1863>`__ 8.
Panic while InsertMulti \* type in ORM
`#1882 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1882>`__ 9. Multiple
execution of task in a very short time.
`#1909 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1909>`__ 10. Garbled
file name in IE `#1912 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1912>`__
11. ORM DISTINCT
`#1938 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1938>`__ 12. Can’t use
int while setting file permit in Logs module.
`#1948 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1948>`__
`#2003 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2003>`__ 13. Empty
foreign key for QueryRow and QueryRows.
`#1964 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1964>`__ 14. Retrieving
scheme from X-Forwarded-Proto when it isn’t none.
`#2050 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2050>`__ 15. Add query
parameters for redirecting static path to ``path/``
`#2064 <https://github.com/W3-Engineers-Ltd/Radiant/pull/2064>`__

radiant 1.6.1
=============

New features

1. Oracle driver for ORM
2. inline mode for ORM Model
3. ssdb engine for Cache
4. Color scheme configure for console out
5. travis integration
6. mulitfile engine for Log. Write logs from different levels to
   different files.

bugfix： 1. cookie time config 2. Router rule mapping
`#1580 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1580>`__ 3. No logs
before radiant.Run() 4. Returning nil while []string is empty in config
5. Wrong comment for ini engine 6. Log time delay while store log
asynchronously 7. Config file parsed twice. 8. Can’t handle ``()`` in
URL for regex router. 9. Chinese encoding issue in mail 10. No Distinct
in ORM 11. Compiling error in Layout 12. Wrong file name in logrotate
13. Invalid CORS if CORS plugin fail. 14. Conflicting between path
params and router params in filters 15. Return 404 other than 200 if
static files are not found. 16. Added GroupBy interface 17. Static file
crush caused by accessing map concurrently of Go 1.6 18. Extra newline
output by json.Encoder of JSONBody in httplib 19. Missing log when Close
if use flush in log under asynchronous mode.

radiant 1.6.0
=============

New features:

1.  ``log`` supports rotating files like ``xx.2013-01-01.2.log``
    `#1265 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1265>`__
2.  ``context.response`` supports Flush, Hijack, CloseNotify
3.  ORM supports Distinct
    `#1276 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1276>`__
4.  ``map_get`` template method
    `#1305 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1305>`__
5.  ORM supports `tidb <https://github.com/pingcap/tidb>`__ engine
    `#1366 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1366>`__
6.  httplib request supports []string
    `#1308 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1308>`__
7.  ORM ``querySeter`` added ``GroupBy`` method
    `#1345 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1345>`__
8.  Session’s MySQL engine supports custom table name
    `#1348 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1348>`__
9.  Performance of log’s file engine improved 30%; Supports set log
    file’s permission
    `#1560 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1560>`__
10. Get session by query
    `#1507 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1507>`__
11. Cache module supports multiple Cache objects.
12. validation supports custom validation functions

bugfix:

1.  ``bind`` method in ``context`` caused crash when parameter is empty.
    `#1245 <https://github.com/W3-Engineers-Ltd/Radiant/issues/1245>`__
2.  manytomany in ORM reverse error
    `#671 <https://github.com/W3-Engineers-Ltd/Radiant/issues/671>`__
3.  http: multiple response.WriteHeader calls
    `#1329 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1329>`__
4.  ParseForm uses local timezone while parsing date
    `#1343 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1343>`__
5.  Emails sent by log’s SMTP engine can’t be authorised
6.  Fixed some issues in router: ``/topic/:id/?:auth``,
    ``/topic/:id/?:auth:int``
    `#1349 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1349>`__
7.  Fixed the crash caused by nil while parsing comment documentation.
    `#1367 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1367>`__
8.  Can’t read ``index.html`` in static folder
9.  ``dbBase.Update`` doesn’t return err if failed
    `#1384 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1384>`__
10. ``Required`` in ``validation`` only works for int but not for int64
11. orm: Fix handling of rel(fk) to model with string pk
    `#1379 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1379>`__
12. graceful error while both http and https enabled
    `#1414 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1414>`__
13. If ListenTCP4 enabled and httpaddr is empty, it still listens TCP6
14. migration doesn’t support postgres
    `#1434 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1434>`__
15. Default values of ORM text, bool will cause error while creating
    tables.
16. graceful panic: negative WaitGroup counter

Improvement:

1.  Moved example to `samples <https://github.com/radiant/samples>`__
2.  Passed golint
3.  Rewrote router, improved performance by 3 times.
4.  Used ``sync.Pool`` for ``context`` to improve performance
5.  Improved template compiling speed.
    `#1298 <https://github.com/W3-Engineers-Ltd/Radiant/pull/1298>`__
6.  Improved config
7.  Refactored whole codebase for readability and maintainability
8.  Moved all init code into ``AddAPPStartHook``
9.  Removed ``middleware``. Will only use ``plugins``
10. Refactored ``Error`` handling.

Radiant 1.5.0
=============

New Features:

1. Graceful shutdown
2. Added ``JsonBody`` method to ``httplib`` which supporting sending raw
   body as JSON format
3. Added ``AcceptsHtml`` ``AcceptsXml`` ``AcceptsJson`` methods to
   ``context input``
4. Get config files from Runmode first
5. ``httplib`` supports ``gzip``
6. ``log`` module stop using asynchronous mode by default
7. ``validation`` supports recursion
8. Added ``apk mime``
9. ``ORM`` supports ``eq`` an ``ne``

Bugfixes:

1.  Wrong parameters for ledis driver.
2.  When user refresh the page after the captcha code expired from the
    cache, it returns 404. Generating new captcha code for reloading.
3.  Controller defines Error exception
4.  cookie doesn’t work in window IE
5.  GetIn returns nil error while getting non-exist variable
6.  More cellphone validation code
7.  Wrong router matching
8.  The ``panic`` returns http 200
9.  The database setting erros caused by redis session
10. The issue that https and http don’t share session
11. Memcache session driver returns error if it’s empty

Radiant 1.4.3
=============

New Features:

1.  ORM support default settting
2.  improve logs/file line count
3.  sesesion ledis support select db
4.  session redis support select db
5.  cache redis support select db
6.  ``UrlFor`` support all type of the parameters
7.  controller ``GetInt/GetString`` function support default value,
    like: ``GetInt("a",12)``
8.  add ``CompareNot/NotNil`` template function
9.  support Controller defeine error，\ `controller
    Error <http://radiant.vip/docs/mvc/controller/errors.html#controller%E5%AE%9A%E4%B9%89error>`__
10. ``ParseForm`` support slices of ints and strings
11. improve ORM interface

bugfix: 1. context get wrong subdomain 2. ``radiant.AppConfig.Strings``
when the strings is empty, always return ``[]string{}`` 3.
utils/pagination can’t modify the attributes 4. when the request url is
empty, route tree crashes 5. can’t click the link to run the task in
adminui 6. FASTCGI restart didn’t delete the unix Socket file

Radiant 1.4.2
=============

New Features:

1.  Added SQL Constructor inspired by ZEND ORM.
2.  Added ``GetInt()``, ``GetInt8()``, ``GetInt16()``, ``GetInt32()``,
    ``GetInt64()`` for Controller.
3.  Improved the logging. Added ``FilterHandler`` for filter logging
    output.
4.  Static folder supports ``index.html``. Automatically adding ``/``
    for static folders.
5.  ``flash`` supports ``success`` and ``set`` methods.
6.  Config for ignoring case for routers: ``RouterCaseSensitive``. Case
    sensitive by default.
7.  Configs load based on environment:
    ``radiant.AppConfig.String("myvar")`` return 456 on dev mode and
    return 123 on the other modes.

       runmode = dev myvar = 123 [dev] myvar = 456

8.  Added ``include`` for ``ini`` config files:

       appname = btest include b.conf

9.  Added ``paginator`` utils.
10. Added ``BEEGO_RUNMODE`` environment variable. You can change the
    application mode by changing this environment variable.
11. Added Json function for fetching ``statistics`` in ``toolbox``.
12. Attachements support for mail utils.
13. Turn on fastcgi by standard IO.
14. Using ``SETEX`` command to support the old version redis in redis
    Session engine.
15. RenderForm supports html id and class by using id and class tag.
16. ini config files support BOM head.
17. Added new Session engine ``ledis``.
18. Improved file uploading in ``httplib``. Supporting extremely large
    files by using ``io.Pipe``.
19. Binding to TCP4 address by default. It will bind to ipv6 in GO.
    Added config variable ``ListenTCP4``.
20. off/on/yes/no/1/0 will parse to ``bool`` in form rendering. Support
    time format.
21. Simplify the generating of SeesionID. Using golang buildin ``rand``
    function other than ``hmac_sha1``.

bugfix:

1.  XSRF verification failure while ``PUT`` and ``DELETE`` cased by
    lowercased ``_method``
2.  No error message returned while initialize the cache by
    ``StartAndGC``
3.  Can’t set ``User-Agent`` in ``httplib``
4.  Improved ``DelStaticPath``
5.  Only finding files in the first static folder when using multiple
    static folders
6.  ``Filter`` functions can’t execute after ``AfterExec`` and
    ``FinishRouter``
7.  Fixed uninitialized mime
8.  Wrong file name and line number in the log
9.  Can’t send the request while only uploading one file in ``httplib``
10. Improved the ``Abort`` output message. It couldn’t out undefined
    error message before.
11. Fixed the issue that can’t add inner Filter while no out Filter set
    in the nested ``namespaces``
12. Router mapping error while router has multiple level parameters.
    #824
13. The information lossing while having many ``namespaces`` for the
    commented router. #770
14. ``urlfor`` function calling useless {{placeholder}} #759

Radiant 1.4.1
=============

New features:

1. ``context.Input.Url`` get path info without domain scheme.
2. Added plugin ``apiauth`` to simulate the ``AWS`` encrypted requests.
3. Simplified the debug output for router info.
4. Supportting pointer type in ORM.
5. Added ``BasicAuth``, cache for multiple requests

bugfix: 1. Router *.* can’t be parsed

Radiant 1.3.0
=============

Hi guys! After the hard working for one month, we are so excited to
release Radiant 1.3.0. We brought many useful features. `Upgrade
notes <http://radiant.vip/docs/intro/upgrade.html>`__

The brand new router system
^^^^^^^^^^^^^^^^^^^^^^^^^^^

We rewrote the router system to tree router. It improved the performance
significantly and supported more formats.

For the routers below:

::

   /user/astaxie
   /user/:username

If the request is ``/user/astaxie``, it will match fixed router which is
the first one; If the request is ``/user/slene``, it will match the
second one. The register order doesn’t matter.

namespace is more elegant
^^^^^^^^^^^^^^^^^^^^^^^^^

``namespace`` is designed for modular applications. It was using chain
style similar to jQuery in previous version but ``gofmt`` can’t format
it very well. Now we are using multi parameters style: (The chain style
still works)

::

   ns :=
   radiant.NewNamespace("/v1",
       radiant.NSNamespace("/shop",
           radiant.NSGet("/:id", func(ctx *context.Context) {
               ctx.Output.Body([]byte("shopinfo"))
           }),
       ),
       radiant.NSNamespace("/order",
           radiant.NSGet("/:id", func(ctx *context.Context) {
               ctx.Output.Body([]byte("orderinfo"))
           }),
       ),
       radiant.NSNamespace("/crm",
           radiant.NSGet("/:id", func(ctx *context.Context) {
               ctx.Output.Body([]byte("crminfo"))
           }),
       ),
   )

For more information please check
`namespace <http://radiant.vip/docs/mvc/controller/router.html#namespace>`__

Annotation Router
^^^^^^^^^^^^^^^^^

::

   // CMS API
   type CMSController struct {
       radiant.Controller
   }

   func (c *CMSController) URLMapping() {
       c.Mapping("StaticBlock", c.StaticBlock)
       c.Mapping("AllBlock", c.AllBlock)
   }

   // @router /staticblock/:key [get]
   func (this *CMSController) StaticBlock() {

   }

   // @router /all/:key [get]
   func (this *CMSController) AllBlock() {
   }

`Annotation
Router <http://radiant.vip/docs/mvc/controller/router.html#annotations>`__

Automated API Document
^^^^^^^^^^^^^^^^^^^^^^

Automated document is a very cool feature that I wish to have. Now it
became real in Radiant. As I said Radiant will not only boost the
development of API but also make the API easy to use for the user.

The API document can be generated by annotations automatically and can
be tested online.

|image2|

|image3|

For more information please check `Automated
Document <http://radiant.vip/docs/advantage/docs.html>`__

config supports different Runmode
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

You can set configurations for different Runmode under their own
sections. Radiant will take the configurations of current Runmode by
default. For example:

::

   appname = radicalpkg
   httpaddr = "127.0.0.1"
   httpport = 9090
   runmode ="dev"
   autorender = false
   autorecover = false
   viewspath = "myview"

   [dev]
   httpport = 8080
   [prod]
   httpport = 8088
   [test]
   httpport = 8888

The configurations above set up httpport for dev, prod and test
environment. Radiant will take httpport = 8080 for current runmode
“dev”.

Support Two Way Authentication for SSL
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

::

   config := tls.Config{
       ClientAuth: tls.RequireAndVerifyClientCert,
       Certificates: []tls.Certificate{cert},
       ClientCAs: pool,
   }
   config.Rand = rand.Reader

   radiant.RadicalApp.Server.TLSConfig = &config

radiant.Run supports parameter
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

``radiant.Run()`` Run on ``HttpPort`` by default

``radiant.Run(":8089")``

``radiant.Run("127.0.0.1:8089")``

Increased XSRFKEY token from 15 characters to 32 characters.
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Removed hot reload
^^^^^^^^^^^^^^^^^^

Template function supports Config. Get Config value from Template easily.
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

::

   {{config returnType key defaultValue}}

   {{config "int" "httpport" 8080}}

httplib supports cookiejar. Thanks to curvesft
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

orm suports time format. If empty return nil other than 0000.00.00 Thanks to JessonChan
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

config module supports parsing a json array. Thanks to chrisport
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

bug fix
~~~~~~~

-  Fixed static folder infinite loop
-  Fixed typo

Radiant 1.2.0
=============

Hi guys! After one month of hard work, we released the new awesome
version 1.2.0. Radiant is the fastest Go framework in the latest `Web
Framework
Benchmarks <http://www.techempower.com/benchmarks/#section=data-r9&hw=i7&test=json>`__
already though our goal is to make Radiant the best and easiest
framework to use. In this new release, we improved even more in both
usability and performance which is closer to native Go.

New Features:
~~~~~~~~~~~~~

1. ``namespace`` Support
^^^^^^^^^^^^^^^^^^^^^^^^

::

       radiant.NewNamespace("/v1").
           Filter("before", auth).
           Get("/notallowed", func(ctx *context.Context) {
           ctx.Output.Body([]byte("notAllowed"))
       }).
           Router("/version", &AdminController{}, "get:ShowAPIVersion").
           Router("/changepassword", &UserController{}).
           Namespace(
           radiant.NewNamespace("/shop").
               Filter("before", sentry).
               Get("/:id", func(ctx *context.Context) {
               ctx.Output.Body([]byte("notAllowed"))
           }))

The code above supports the URL requests below:

::

   GET       /v1/notallowed
   GET       /v1/version
   GET       /v1/changepassword
   POST      /v1/changepassword
   GET       /v1/shop/123

``namespace`` also supports pre-filters, conditions checking and
unlimited nested ``namespace``

2. Supporting more flexible router modes
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Custom functions from RESTful router

::

   radiant.Get(router, radiant.FilterFunc)
   radiant.Post(router, radiant.FilterFunc)
   radiant.Put(router, radiant.FilterFunc)
   radiant.Head(router, radiant.FilterFunc)
   radiant.Options(router, radiant.FilterFunc)
   radiant.Delete(router, radiant.FilterFunc)

   radiant.Get("/user", func(ctx *context.Context) {
       ctx.Output.Body([]byte("Get userlist"))
   })

More flexible Handler

``radiant.Handler(router, http.Handler)``

Integrating other services easily

::

   import (
       "http"
       "github.com/gorilla/rpc"
       "github.com/gorilla/rpc/json"
   )

   func init() {
       s := rpc.NewServer()
       s.RegisterCodec(json.NewCodec(), "application/json")
       s.RegisterService(new(HelloService), "")
       radiant.Handler("/rpc", s)
   }

3. Binding request parameters to object directly
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

For example: this request parameters

::

   ?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie

::

   var id int
   ctx.Input.Bind(&id, "id")  //id ==123

   var isok bool
   ctx.Input.Bind(&isok, "isok")  //isok ==true

   var ft float64
   ctx.Input.Bind(&ft, "ft")  //ft ==1.2

   ol := make([]int, 0, 2)
   ctx.Input.Bind(&ol, "ol")  //ol ==[1 2]

   ul := make([]string, 0, 2)
   ctx.Input.Bind(&ul, "ul")  //ul ==[str array]

   user struct{Name}
   ctx.Input.Bind(&user, "user")  //user =={Name:"astaxie"}

4. Optimized the form parsing flow and improved the performance
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

5. Added more testcases
^^^^^^^^^^^^^^^^^^^^^^^

6. Added links for admin monitoring module
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

7. supporting saving struct into session
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

8.httplib supports file upload interface
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

::

   b:=httplib.Post("http://radiant.vip/")
   b.Param("username","astaxie")
   b.Param("password","123456")
   b.PostFile("uploadfile1", "httplib.pdf")
   b.PostFile("uploadfile2", "httplib.txt")
   str, err := b.String()
   if err != nil {
       t.Fatal(err)
   }

``httplib`` also supports custom protocol version

9. ORM supports all the unexport fields of struct
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

10. Enable XSRF in controller level. XSRF can only be controlled in the whole project level. However, you may want to have more control for XSRF, so we let you control it in Prepare function in controller level. Default is true which means using the global setting.
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

::

   func (a *AdminController) Prepare(){
          a.EnableXSRF = false
   }

11. controller supports ServeFormatted function which supports calling ServeJson or ServeXML based on the request’s Accept
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

12. session supports memcache engine
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

13. The Download function of Context supports custom download file name
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Bug Fixes
---------

1. Fixed the bug that session’s Cookie engine can’t set expiring time
2. Fixed the bug of saving and parsing flash data
3. Fixed all the problems of ``go vet``
4. Fixed the bug of ParseFormOrMulitForm
5. Fixed the bug that only POST can parse raw body. Now all the requests
   except GET and HEAD support raw body.
6. Fixed the bug that config module can’t parse ``xml`` and ``yaml``

Radiant 1.1.4
=============

This is an emergency release for solving a serious security problem.
Please update to the latest version! By the way released all changes
together.

1. fixed a security problem. I will show the details in
   radiant/security.html later.

2. ``statifile`` move to new file.

3. move dependence of the third libs,if you use this module in your
   application: session/cache/config, please import the submodule of the
   third libs:

   ::

      import (
           "github.com/W3-Engineers-Ltd/Radiant"
         _ "github.com/W3-Engineers-Ltd/Radiant/session/mysql"
      )

4. modify some functions to private.

5. improve the FormParse.

released date: 2014-04-08

Radiant 1.1.3
=============

this is a hot fixed:

1. console engine for logs.It will not run if there’s no config.

2. Radiant 1.1.2 support ``go run main.go``, but if ``main.go`` bot
   abute the Radiant’s project rule,use own AppConfigPath or not exist
   app.conf will panic.

3. Radiant 1.1.2 supports ``go test`` parse config,but actually when
   call TestRadiantInit still can’t parseconfig

released date: 2014-04-04

Radiant 1.1.2
=============

The improvements:

1.  Added ExceptMethodAppend fuction which supports filter out some
    functions while run autorouter
2.  Supporting user-defined FlashName, FlashSeperator
3.  ORM supports user-defined types such as type MyInt int
4.  Fixed validation module return user-defined validating messages
5.  Improved logs module, added Init processing errors. Changed some
    unnecessory public function to private
6.  Added PostgreSQL engine for session module
7.  logs module supports output caller filename and line number. Added
    EnableFuncCallDepth function, closed by default.
8.  Fixed bugs of Cookie engine in session module
9.  Improved the error message for templates parsing error
10. Allowing modifing Context by Filter to skip Radiant’s routering
    rules and using uder-defined routering rules. Added parameters
    RunController and RunMethod
11. Supporting to run Radiant APP by using ``go run main.go``
12. Supporting to run test cases by using ``go test``. Added
    TestRadiantInit function.

released date: 2014-04-03

Radiant 1.1.1
=============

Added some new features and fixed some bugs in this release.

1.  File engine can’t delete file in session module which will raise
    reading failure.
2.  File cache can’t read struct. Improved god automating register
3.  New couchbase engine for session module
4.  httplib supports transport and proxy
5.  Improved the Cookie function in context which support httponly by
    default as well as some other default parameters.
6.  Improved validation module to support different cellphone No.
7.  Made getstrings function to as same as getstring which doesn’t need
    parseform
8.  Redis engine in session module will return error while connection
    failure
9.  Fixed the bug of unable to add GroupRouters
10. Fixed the bugs for multiple static files, routes matching bug and
    display the static folder automatically
11. Added GetDB to get connected \*sql.DB in ORM
12. Added ResetModelCache for ORM to reset the struct which has already
    registered the cache in order to write tests easily
13. Supporting between in ORM
14. Supporting sql.Null\* type in ORM
15. Modified auto_now_add which will skip time setting if there is
    default value.

released date: 2014-03-12

Radiant 1.1.0
=============

Added some new features and fixed some bugs in this release.

New features

1.  Supporting AddAPPStartHook function
2.  Supporting plugin mode; Supporting AddGroupRouter for configuring
    plugin routes.
3.  Response supporting HiJacker interface
4.  AddFilter supports batch matching
5.  Refactored session module, supporting Cookie engine
6.  Performance benchmark for ORM
7.  Added strings interface for config which allows configuration
8.  Supporting template render control in controller level
9.  Added basicauth plugin which can implement authentication easily
10. #436 insert multiple objects
11. #384 query map to struct

bugfix

1. Fixed the bug of FileCache
2. Fixed the import lib of websocket
3. Changed http status from 200 to 500 when there are internal error.
4. gmfim map in memzipfile.go file should use some synchronization
   mechanism (for example sync.RWMutex) otherwise it errors sometimes.
5. Fixed #440 on_delete bug that not getting delted automatically
6. Fixed #441 timezone bug

released date: 2014-02-10

Radiant 1.0 release
===================

After four months code refactoring, we released the first stable version
of Radiant. We did a lot of refactoring and improved a lot in detail.
Here is the list of the main improvements:

1. Modular design. Right now Radiant is a light weight assembling
   framework with eight powerful stand alone modules including cache,
   config, logs, sessions, httplibs, toolbox, orm and context. It might
   have more in the future. You can use all of these stand alone modules
   in your other applications directly no matter it’s web applications
   or any other applications such as web games and mobile games.

2. Supervisor module. In the real world engineering, after the
   deployment of the application, we need to do many kinds of statistics
   and analytics for the application such as QPS statistics, GC
   analytics, memory and CPU monitoring and so on. When the live issue
   happends we also want to debug and profile our application on live.
   All of these real world engineering features are included in Radiant.
   You can enable the supervisor module in Radiant and visit it from
   default port 8088.

3. Detailed document. We rewritten all the document. We improved the
   document based on many advices from the users. To make it communicate
   easier for different language speakers, now the comments of the
   document in each language are separated.

4. Demos. We provided three examples, chat room, url shortener and todo
   list. You can understand and use Radiant easier and faster by
   learning the demos.

5. Redesigned Radiant website. Nice people from Radiant community helped
   Radiant for logo design and website design.

6. More and more users. We listed our typical users in our homepage.
   They are all big companies and they are using Radiant for their
   products already. Radiant already tested by those live applications.

7. Growing active communities. There are more than 390 issues on github,
   more than 36 contributors and more than 700 commits. Google groups is
   also growing.

8. More and more applications in Radiant. There are some open source
   applications as well. E.g.: CMS system:
   https://github.com/insionng/toropress and admin system:
   https://github.com/radiant/admin

9. Powerful assistance tools. radical is used to assist the development
   of Radiant applications. It can create, compile, package the Radiant
   application easily.

released date: 2013-12-19

.. |image0| image:: https://cloud.githubusercontent.com/assets/707691/14017109/f608b658-f1ff-11e5-8d57-72030cfe4f5d.png
.. |image1| image:: https://cloud.githubusercontent.com/assets/1248967/16153054/f654b08e-34a4-11e6-894d-24f16ab847a7.png
.. |image2| image:: ../images/docs.png
.. |image3| image:: ../images/doc_test.png

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