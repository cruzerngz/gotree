# gotree

## [What is this?](https://www.youtube.com/watch?v=1sFbLppuhhs&t=28)

<img src="https://cdn.vox-cdn.com/thumbor/4Fflu59FB8lQ1XLfvPXTAN7MdJY=/0x0:2880x1615/1525x1525/filters:focal(1396x447:1856x907):format(webp)/cdn.vox-cdn.com/uploads/chorus_image/image/66531693/image___2020_03_19T170444.621.0.png" width=35%>

I'm learning some golang. So here's (a part of) the linux `tree` command implemented in go.

```sh
git clone https://github.com/cruzerngz/gotree.git
cd gotree

## if you have make installed
make
./bin/gotree --path <your-path>

## if you don't have make installed
go build ./cmd/gotree.go
./gotree --path <your-path>
```

## How small/fast is it?
The unstripped binary with debug information is 2.1MB large, 1.4MB stripped.
In comparison, the `tree` command takes up a measly 87KB, though it links to external libraries.

Surprisingly, `gotree` is faster than `tree` at doing... things.
Keep in mind that `tree` accepts something like 50 command line flags while `gotree` only accepts 1.

I ran these on a directory with 4123 subdirectories, 46505 files; and...
```
Benchmark 1: ./bin/gotree --path ..
  Time (mean ± σ):      82.3 ms ±   3.0 ms    [User: 37.8 ms, System: 46.7 ms]
  Range (min … max):    78.7 ms …  96.3 ms    36 runs

Benchmark 2: tree ..
  Time (mean ± σ):     166.1 ms ±   3.2 ms    [User: 70.2 ms, System: 94.9 ms]
  Range (min … max):   160.3 ms … 170.4 ms    17 runs

Summary
  './bin/gotree --path ..' ran
    2.02 ± 0.08 times faster than 'tree ..'

```
---

Thank you for coming to my TED talk
