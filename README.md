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

Just call `prismic-infos` with the URL of the API.

Example:

```bash
$ prismic-infos https://lesbonneschoses.prismic.io/api
Releases:
- release (id=master ref=UlfoxUnM08QWYXdl) {master} “Master”
Forms:
- form (id=selections) “Product selections”
- form (id=macarons) “Macarons”
- form (id=cupcakes) “Cupcakes”
- form (id=pies) “Little Pies”
- form (id=blog) “The Blog”
- form (id=stores) “Stores”
- form (id=jobs) “Jobs”
- form (id=featured) “Featured”
- form (id=products) “All Products”
- form (id=everything)
Bookmarks:
- bookmark (ref=UlfoxUnM0wkXYXbf) “about”
- bookmark (ref=UlfoxUnM0wkXYXbd) “jobs”
- bookmark (ref=UlfoxUnM0wkXYXbp) “stores”
Types:
- type (id=blog-post) “Blog post”
- type (id=store) “Store”
- type (id=article) “Site-level article”
- type (id=selection) “Products selection”
- type (id=job-offer) “Job offer”
- type (id=product) “Product”
Tags: Macaron, Pie, Cupcake, Featured
```

### Options

To get some help

```bash
$ prismic-infos -h
Usage of prismic-infos:
  -only="all": The infos to display (releases,forms,bookmarks,types,tags)
  -token="": The access token
  -utc=false: Display dates in UTC
```

To display only some informations:

```bash
$ prismic-infos -only=releases,bookmarks https://lesbonneschoses.prismic.io/api
Releases:
- release (id=master ref=UlfoxUnM08QWYXdl) {master} “Master”
Bookmarks:
- bookmark (ref=UlfoxUnM0wkXYXbf) “about”
- bookmark (ref=UlfoxUnM0wkXYXbd) “jobs”
- bookmark (ref=UlfoxUnM0wkXYXbp) “stores”
```

To display UTC dates:

```bash
prismic-infos -utc https://dohzya-test.prismic.io/api
```

You can also use a (permanant) access token:

```bash
prismic-infos -token="$YOURTOKEN" https://yourrepo.prismic.io/api
```
