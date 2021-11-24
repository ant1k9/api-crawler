![CI](https://github.com/ant1k9/api-crawler/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/ant1k9/api-crawler/branch/main/graph/badge.svg)](https://codecov.io/gh/ant1k9/api-crawler)

### Api-Crawler

Crawl your sites with the given configuration

```bash
$ goose -dir migrations postgres "host=127.0.0.1 port=5432 user=<user> password=<password> dbname=<dbname> sllmode=disable"
$ make build
$ ./bin/api-crawler add-plugin <some-type>
$ ./bin/api-crawler crawl <some-type>
```
