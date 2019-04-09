The below compares using `babel-cli` (via `os/exec`) vs doing the same via HTTP POST to a running Node.js server.
 

```bash
BenchmarkBabelHTTP-4         300           4084183 ns/op            4512 B/op         60 allocs/op
BenchmarkBabelExec-4           2         514577991 ns/op           11248 B/op         86 allocs/op
```

4084183 is 4,08 milliseconds; still not very fast, but it's still 125x faster than the  `os/exec` variant.