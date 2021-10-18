# Go File Walk Test

Quick comparison between different Go walk implementations.

https://engineering.kablamo.com.au/posts/2021/quick-comparison-between-go-file-walk-implementations

Install them and use hyperfine to test.

```
hyperfine 'filepathwalk' 'iafan' 'karrick' 'michealtjones' 'readdir'
```

Or run like so to confirm the output is the same for all.

```
filepathwalk && iafan && karrick && michealtjones && readdir
```
