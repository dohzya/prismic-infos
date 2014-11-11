Prismic-infos
=============

Display some informations on a prismic.io repository

Install
-------

To install, be sure to have a valid Go installation, then call

```bash
go get github.com/dohzya/prismic-infos
```

Use
---

Displaying some informations (the releases) on a prismic.io repository:

```bash
prismic-infos https://lesbonneschoses.prismic.io/api
```

In order to display UTC dates:

```bash
prismic-infos -utc https://dohzya-test.prismic.io/api
```

You can also use a (permanant) access token:

```bash
prismic-infos -token="$YOURTOKEN" https://yourrepo.prismic.io/api
```
