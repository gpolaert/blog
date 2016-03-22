# Multi-line support

## What will try to do?
In this post, I'm going to show how handle multi lines with the common log shippers:
* rsyslog,
* syslog-ng
* fluentd
* logstash
* nxlo (soon)


I will illustrated each shippers with two common examples: Java error stacks and PostgreSQL log events.
Java and Postgres are quite good produces multi-line logs, for SQL queries, or for THE NullPointerException.
Parsing multi-line to collect them into a single event can be rapidly a nightmare.

So, don't panic and take a seat. Debugging and monitoring should be more pleasant soon.

Before start, let's talk about some prerequisites and context if you want to reproduce examples.
1. All logs are written into a file. I mainly use a `tail-like` command.
2. All logs are stream to a single pipe: in others terms, I assume that all lines for an event
are contiguous. In this post, I won't show you how to do when multiple events are cross.
3. We don't focus to the event output.


### Input events (PostgreSQL/Java)
In order to reproduce examples, you have to configure the log section in PostgreSQL. Edit
`$PG_DATA/postgresql.conf` and add/replace the following section:
```properties

# In order to convert multi-line to a single message, you need to bypass syslog handler
log_destination = 'stderr'

# Collect the stderr in write logs into a file
logging_collector = on

# For this tutorial, all lines are written to a single file
# Rolling files are still recommended in production environment
log_filename = 'postgresql.log'


# Tell Postgres to be more verbose. Adapt it for your needs
client_min_messages = debug3
log_min_messages = debug3
log_min_error_statement = debug3
log_min_duration_statement = 0

# Customize log prefix
log_line_prefix = '%t postgres[%p]: '
```

Then restart your postgres server.

## Parse PostgreSQL multi-lines
Before start, let's see how logs look like, `tail $PG_DATA/pg_log/postgresql.log`.

```
2016-03-21 11:22:09 UTC postgres[423]: DEBUG:  name: unnamed; blockState:       STARTED; state: INPROGR, xid/subid/cid: 0/1/0, nestlvl: 1, children:
2016-03-21 11:22:10 UTC postgres[423]: DEBUG:  StartTransactionCommand
2016-03-21 11:22:10 UTC postgres[423]: STATEMENT:  select *
        from
        foo;
2016-03-21 11:22:10 UTC postgres[423]: DEBUG:  StartTransaction
...
```


### RSyslog (since v8.10.0)
Multi-line support is recent and required at minium the 8.10.0 rsyslog version.
Before start, check the version with `rsyslogd -v`.

Multu-line support is only available on `imfile` module. So, your logs need to be write
in a specific file before be parsed and sent to syslog daemon. Send logs to a facility (local0 for instance),
won't work.

```properties
imfile(

)
```

### Syslog-ng


### Fluentd

The `in_tail` plugin allows to parse multi-lines. You have to set up `multiline` format in order to
enable the multi-line parsing.

The syntax is quite simple. You have to define a patter for the split lines. Regex are ruby-compatible.
In our example, you still break lines on "SEVERITY:". Next lines start with spaces or tabs.


Fluentd allows to add format1 trough format20 in order to parse different line format. Here, the plugin
use the grok format. See the [documentation](http://) for more details.

On the above input, this configuration

```xml
<source>
  @type tail
  path  /var/lib/postgresql/data/pg_log/postgresql.log
  pos_file /var/log/td-agent/postgresql.log.pos
  tag postgres
  format multiline
  format_firstline /^(STATEMENT|LOG|DEBUG|FATAL|ERROR|WARNING|[A-Z]+):/
  format1 /^(?<severity>(STATEMENT|LOG|DEBUG|FATAL|ERROR|WARNING|[A-Z]+)): (?<message>.*)$/
  time_format %b %d %H:%M:%S
</source>
```

Will produce these events

```json


```
### Logstash