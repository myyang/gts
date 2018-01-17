# gts
Timestamp helper

## Install

```
go get github.com/myyang/gts
```

## Usage

```
Usage of gts:
  -dt string
        UTC datetime string to parse. Format: RFC3339 (default "1970-01-01T00:00:00Z")
  -p string
        Timestamp precision, default: 'ms'. Options: 's', 'ms', 'us', 'ns' (default "ms")
```

run in shell directly

```
$ gts
2017-12-27T09:29:37Z -> 1514366977392 (ms)

$ gts 1514366977
2017-12-27T09:29:37Z -> 1514366977000 (ms)

$ gts 1514366977392 -p 's'
2017-12-27T09:29:37Z -> 1514366977 (s)

$ gts -dt '1970-01-01T00:00:01Z'
1970-01-01T00:00:01Z -> 1000 (ms)

```
