# Simple tool for reading from loki and writing back

I needed to mutate some log lines in Loki, which is something it doesn't do.

So to accomplish this, I wrote this tool to read some lines, modify them via the query, and write them back (filtering only some labels, note this is done pretty crudely in loki.go and ony really applies to the query I used. In newer Loki versions there is a `drop` in logql syntax which could maybe replace this all together)

This is a really rough implementation and notably doesn't handle one VERY important case of having multiple log lines with the exact same nanosecond timestamp and different values.

Loki allows storing multiple entries at the same timestamp with different values. The data I was working with here does not have this condition so this code does not properly handle this if it were to exist.