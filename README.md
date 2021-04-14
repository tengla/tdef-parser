# TDEF parser

Made in golang

```
make all
```

Make sure your source file is utf-8, ie:
```
iconv -f ISO-8859-1 -t UTF-8 data/TDEF_R21_NO.txt > data/TDEF_R21_NO.utf8.tsv
```

Make json output
```
./dist/parser -file data/TDEF_R21_NO.utf8.tsv -print
```

You can pipe it to a file
```
./dist/parser -file data/TDEF_R21_NO.utf8.tsv -print > data.json
```

Or query with jq
```
./dist/parser -file data/TDEF_R21_NO.utf8.tsv -print | \
  jq '.Thds[].Thd|{TrainID,TrainOwner,InitialDateOfOperation}'
```
