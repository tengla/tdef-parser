# TDEF parser

Made in golang

```
make all
```

Make json output
```
./dist/parser -file <path to .tsv> -print
```

You can pipe it to a file
```
./dist/parser -file <path to .tsv> -print > data.json
```

Or query with jq
```
./dist/parser -file <path to .tsv> -print | \
  jq '.Thds[].Thd|{TrainID,TrainOwner,InitialDateOfOperation}'
```
