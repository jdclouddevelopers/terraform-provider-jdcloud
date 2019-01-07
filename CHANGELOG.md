## 0.0.1 (Unreleased)

FEATURES:

* Importer for various resources is introduced [GH-17]
* Website files prepared [GH-18]
* CHANGELOG.md created [GH-18]

BUG FIXES:

* Unable to set ACL when OSS bucket is created [GH-17]
* No validation on `subnet_id` and `vpc_id` when creating RDS instance[GH-15]

IMPROVEMENTS:

* Code format modified. Errors will be returned when trying to invoke `d.Set("List/Set",List/Set)` [GH-17]
* `*schema.TypeList` -> `*schema.TypeSet` [GH-17]
* GNUMakefile introduced [GH-17]
* Scripts for various purpose introduced [GH-18] 
* .travis.yml modified for a more detailed testing process [GH-18]