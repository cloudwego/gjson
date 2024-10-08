# GJSON
This is a wrapper lib of [gjson](https://github.com/tidwall/gjson), which is accelerated by [sonic](https://github.com/bytedance/sonic)'s algorithm.

## Performance
It's usually faster than original one, especially for large JSON. (see codes [here](testdata/gjson_timing_test.go))
```
goversion: 1.22.0
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) Platinum 8260 CPU @ 2.40GHz
                         │  origin.out   │             target.out              │
                         │    sec/op     │   sec/op     vs base                │
GetComplexPath/small-32     3.000µ ±  6%   2.591µ ± 6%  -13.65% (p=0.000 n=10)
GetComplexPath/medium-32   24.855µ ±  5%   6.464µ ± 8%  -74.00% (p=0.000 n=10)
GetComplexPath/large-32    1309.1µ ±  3%   270.0µ ± 8%  -79.37% (p=0.000 n=10)
GetSimplePath/small-32      702.1n ± 11%   634.3n ± 5%   -9.66% (p=0.000 n=10)
GetSimplePath/medium-32     9.558µ ±  3%   1.744µ ± 4%  -81.76% (p=0.000 n=10)
GetSimplePath/Large-32     342.03µ ± 14%   37.83µ ± 5%  -88.94% (p=0.000 n=10)
geomean                     24.64µ         7.576µ       -69.26%

                         │  origin.out  │             target.out              │
                         │     B/op     │    B/op     vs base                 │
GetComplexPath/small-32    104.0 ± 0%     104.0 ± 0%       ~ (p=1.000 n=10) ¹
GetComplexPath/medium-32   16.00 ± 0%     16.00 ± 0%       ~ (p=1.000 n=10) ¹
GetComplexPath/large-32    16.00 ± 0%     16.00 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/small-32     0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/medium-32    0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/Large-32     0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
geomean                               ²               +0.00%                ²
¹ all samples are equal
² summaries must be >0 to compute geomean

                         │  origin.out  │             target.out              │
                         │  allocs/op   │ allocs/op   vs base                 │
GetComplexPath/small-32    6.000 ± 0%     6.000 ± 0%       ~ (p=1.000 n=10) ¹
GetComplexPath/medium-32   2.000 ± 0%     2.000 ± 0%       ~ (p=1.000 n=10) ¹
GetComplexPath/large-32    2.000 ± 0%     2.000 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/small-32     0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/medium-32    0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
GetSimplePath/Large-32     0.000 ± 0%     0.000 ± 0%       ~ (p=1.000 n=10) ¹
geomean                               ²               +0.00%       
```


## Options

### FastPath

This option is used to cache parsed simple paths and search JSON path all-in-once in C funtions, to reduce the overhead of c-go interaction. By default, this option is disabled, because it will and consumes a little more memory than default. You can enable it by set environment variable `GJSON_FAST_PATH=1`.

### FastString

By default, Gjson doesn't use the SIMD algorithm when decoding a string and does not validate UTF8 either, thus its string-value APIs' behaviors are slow and different from `encoding/json`. You can change the behaviors by setting the environment variable below:

- `GJSON_FAST_STRING=1`: `SIMD-implemented` string parsing. The string-value APIs' behaviors will be the same as sonic/decoder's default behavior, with 2~3X times the speed of default string-parsing.
- `GJSON_FAST_STRING=2`, `SIMD-implemented` string parsing, and UTF-8 validating. String-value APIs' behaviors will be totally same with `encoding/json.Decode` and also faster than the default.

### Benchmark

This is a benchmark on the above options (see [codes](testdata/gjson_timing_test.go)):

```
goversion: 1.22.0
goos: linux
goarch: amd64
pkg: github.com/tidwall/gjson/testdata
cpu: Intel(R) Xeon(R) Platinum 8260 CPU @ 2.40GHz
BenchmarkFastPath/normal/small-32                  741.1 ns/op        0 B/op        0 allocs/op
BenchmarkFastPath/normal/medium-32                  1846 ns/op        0 B/op        0 allocs/op
BenchmarkFastPath/normal/Large-32                  38675 ns/op        0 B/op        0 allocs/op
BenchmarkFastPath/fast-path/small-32               364.4 ns/op        0 B/op        0 allocs/op
BenchmarkFastPath/fast-path/medium-32               1662 ns/op        0 B/op        0 allocs/op
BenchmarkFastPath/fast-path/Large-32               38976 ns/op        0 B/op        0 allocs/op
BenchmarkParseString/normal/small-32               401.8 ns/op      192 B/op        2 allocs/op
BenchmarkParseString/normal/medium-32               4600 ns/op     2560 B/op        2 allocs/op
BenchmarkParseString/normal/large-32               63202 ns/op    27904 B/op        2 allocs/op
BenchmarkParseString/fast-string/small-32          219.7 ns/op       96 B/op        1 allocs/op
BenchmarkParseString/fast-string/medium-32          2076 ns/op     1408 B/op        1 allocs/op
BenchmarkParseString/fast-string/large-32          15814 ns/op    14336 B/op        1 allocs/op
BenchmarkParseString/validate-string/small-32      231.6 ns/op       96 B/op        1 allocs/op
BenchmarkParseString/validate-string/medium-32      2116 ns/op     1408 B/op        1 allocs/op
BenchmarkParseString/validate-string/large-32      16779 ns/op    14336 B/op        1 allocs/op
```
