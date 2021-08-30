# Introducing Watch Log
A Go command line tool for watching `json` logs.

## Screenshot
// TODO

# Getting Started (WIP)

## Install

```bash
go -u get github.com/lexcao/watch-log
```

## Usage

For single file:
```bash
$ wlog -f sample.log
```

For streaming: `|` is needed
```bash
# live tail file
$ tail -f sample.log | wlog

# live tail console
$ docker logs -f my-app | wlog
```

## Command while Watching

Type `:` will show the command line like `vim`


```
help
- will show the help of each command

omit [field1, field2...] 
- `o` for short
- will omit json field for showing after formatting a json

pick [field1, field2...]
- `p` for short
- will pick json field for showing after formatting a json (reverse of `omit`)

match [FIELD] = [REGEX]
- `m` for short
- will match the [REGEX] on the given [FIELD]

hightlight  
```

All these commands are available for entry command
e.g.
```bash
# for file
$ wlog -f sample.log -o 
# for stream
$ tail -f sample.log | wlog -o [level,test]
```

## More Info

```bash
$ wlog --help
```

# Structure

```
main.go // Entry Point
cmd
pkg 
| app
| loader
| parser
| pipeline
  | omit
  | match
  | pick
  | highlight
| render
```

## Flow

1. Load file / stream
2. Parse json / other format line by line
3. Process into pipeline with [omit, match, pick, highlight]
4. Render to console
