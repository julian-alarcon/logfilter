# logfilter

> The purpose of this app is just academic

This simple go app will take a text file and look for the strings `ERROR`,
`WARN`, `INFO` and `DEBUG`. If it finds any line with those strings will print
the output depending on the debugLevel defined.

## Requirements

You can run or compile it using golang version 1.16

### Input options

* `-pathLog` - Define a directory path of your log file to be filtered. E.g: `./logfilter -logPath logs/mylog.log`
* `-debugLevel` - Define the debug level desired. Valid numbers are 1=ERROR, 2=ERROR+WARN, 3=ERROR+WARN+INFO, 4=3=ERROR+WARN+INFO

## Run

```sh
$ go run . -logPath catalina.log -debugLevel 2
22:01:25.648 [http-nio-8080-exec-54] WARN  org.apache.solr.core.SolrCore - [ga] warning line
22:01:25.648 [http-nio-8080-exec-64] ERROR  org.apache.solr.core.SolrCore - [ga] error line
22:01:25.648 [http-nio-8080-exec-64] ERROR  org.apache.solr.core.SolrCore - [ga] error line
22:01:25.648 [http-nio-8080-exec-54] WARN  org.apache.solr.core.SolrCore - [ga] warning line
22:01:25.720 [http-nio-8080-exec-30] WARN  org.apache.solr.core.SolrCore - [ga] warning line
```

## Compile

```sh
go build .
```

## Execute

```sh
./logfilter -logPath catalina.log -debugLevel 4
22:01:25.648 [http-nio-8080-exec-54] WARN  org.apache.solr.core.SolrCore - [ga] warning line
22:01:25.648 [http-nio-8080-exec-64] ERROR  org.apache.solr.core.SolrCore - [ga] error line
22:01:25.648 [http-nio-8080-exec-64] ERROR  org.apache.solr.core.SolrCore - [ga] error line
22:01:25.648 [http-nio-8080-exec-54] WARN  org.apache.solr.core.SolrCore - [ga] warning line
22:01:25.720 [http-nio-8080-exec-30] DEBUG  org.apache.solr.core.SolrCore - [ga] debug line
22:01:25.720 [http-nio-8080-exec-30] DEBUG  org.apache.solr.core.SolrCore - [ga] debug line
22:01:25.720 [http-nio-8080-exec-30] INFO  org.apache.solr.core.SolrCore - [ga] info line
```

### Help

```sh
$ ./logfilter -help
Usage of ./logfilter:
  -debugLevel int
        Debug level. valid numbers are 1=ERROR, 2=ERROR+WARN, 3=ERROR+WARN+INFO, 4=3=ERROR+WARN+INFO (default 1)
  -logPath string
        Path of the log file to analyze (default "file.log")
```

## Golang Topics managed

* Setup a minimal go project (`go.mod`, `main.go`)
* Usage of slices instead of arrays. (`mySlice := []string{"ERROR", "WARN", "INFO", "DEBUG"}`)
* Variable naming in go. (small strings, camelCase, only uppercase if variable is going to be used outside)
* Variable declaration (`:=` instead of 2 lines, `var` and then `=`)
* Basic usage of pointers (`*`)
* How to import a library
* Known go debugger rules (Undeclared variables, format)
* Basic gitignore rules for Golang
* Basic procedural statements (`if` and `for`)
* Basic conditionals and string comparison
* `flag` usage for helpful command line apps
* Basic file management
* Printing in console options
* Basic error handling
