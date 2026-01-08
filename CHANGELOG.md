# Changelog

## 0.3.0 (2026-01-08)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/openlayer-ai/openlayer-go/compare/v0.2.1...v0.3.0)

### Features

* **api:** add missing endpoints ([144444a](https://github.com/openlayer-ai/openlayer-go/commit/144444a432c251ac659830e9cba3235c32f9b8cf))

## 0.2.1 (2026-01-06)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/openlayer-ai/openlayer-go/compare/v0.2.0...v0.2.1)

### Chores

* **internal:** codegen related update ([4925276](https://github.com/openlayer-ai/openlayer-go/commit/49252767e49299b0b91c705cdb040e579a20b7bc))

## 0.2.0 (2025-12-17)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.1...v0.2.0)

### Features

* **api:** add test evaluation method ([cb22363](https://github.com/openlayer-ai/openlayer-go/commit/cb22363b161e9bb0a3b3d9a465c346457f6c1e9d))
* **api:** api update ([c4a4bd2](https://github.com/openlayer-ai/openlayer-go/commit/c4a4bd2460cad43f8a78de5c5c4e14b351746bcc))
* **api:** api update ([e8f5ac9](https://github.com/openlayer-ai/openlayer-go/commit/e8f5ac904eb179b5d5bfdc19ffe427b55a957efc))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([7ba0a42](https://github.com/openlayer-ai/openlayer-go/commit/7ba0a42cfbe23939bd1eff8526f63997cd22d5a0))

## 0.1.1 (2025-09-29)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0...v0.1.1)

### Bug Fixes

* bugfix for setting JSON keys with special characters ([93fe0f0](https://github.com/openlayer-ai/openlayer-go/commit/93fe0f0d1bfcca39d46f64a5c64de601495e1b52))

## 0.1.0 (2025-09-22)

Full Changelog: [v0.1.0-alpha.17...v0.1.0](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.17...v0.1.0)

### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([19cb06e](https://github.com/openlayer-ai/openlayer-go/commit/19cb06ee0bd56001bae8e04a0e21147b55da957c))


### Chores

* bump minimum go version to 1.22 ([4bbbeb0](https://github.com/openlayer-ai/openlayer-go/commit/4bbbeb01e4132c1775b72c88c95f888b8d7b0131))
* do not install brew dependencies in ./scripts/bootstrap by default ([039a9bd](https://github.com/openlayer-ai/openlayer-go/commit/039a9bdaef3563687eca21b7667fb968c670febc))
* update more docs for 1.22 ([8ef189a](https://github.com/openlayer-ai/openlayer-go/commit/8ef189a98c9ffa26c5a41634ae09643390672df7))

## 0.1.0-alpha.17 (2025-09-09)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **api:** add test creation endpoint ([50fa1ac](https://github.com/openlayer-ai/openlayer-go/commit/50fa1ac7d9d7bff9402b7be595e176cfab4d18e9))
* **api:** api update ([6a671cc](https://github.com/openlayer-ai/openlayer-go/commit/6a671cc3207209aabebdcad6ceab28512d25c2d2))
* **api:** api update ([9ca34f2](https://github.com/openlayer-ai/openlayer-go/commit/9ca34f2376b8e7736bc178b64fb8de6f14dc09f5))
* **api:** api update ([1464559](https://github.com/openlayer-ai/openlayer-go/commit/146455956798caf86c90a1e76a06663c862ab02a))
* **api:** api update ([945314e](https://github.com/openlayer-ai/openlayer-go/commit/945314e51ecff9e88c9fafe693a612d3d350c313))
* **api:** api update ([20c21a5](https://github.com/openlayer-ai/openlayer-go/commit/20c21a5bfc2a02ef41317797e52b966a29880738))
* **api:** api update ([915af77](https://github.com/openlayer-ai/openlayer-go/commit/915af773e51144f74c1517bb43eb6807867d4015))
* **api:** expose test retrieval endpoint ([cddc990](https://github.com/openlayer-ai/openlayer-go/commit/cddc990ebb0004ddee4ca0f8c538b543c700d254))
* **api:** expose test retrieval endpoint ([7ac3c6c](https://github.com/openlayer-ai/openlayer-go/commit/7ac3c6ccac9752d807519538a54b0afe223154fe))
* **api:** expose test update endpoint ([e8aba5e](https://github.com/openlayer-ai/openlayer-go/commit/e8aba5e6f17d645f5ccb543d259739d534a3da3a))
* **client:** add support for endpoint-specific base URLs in python ([820ffe8](https://github.com/openlayer-ai/openlayer-go/commit/820ffe88bf9a04be100ea8fb34794865b87d5f52))
* **client:** add support for reading base URL from environment variable ([9fd2355](https://github.com/openlayer-ai/openlayer-go/commit/9fd23557e5355ae4280fdea454c024eab8d9dc4d))
* **client:** support custom http clients ([#80](https://github.com/openlayer-ai/openlayer-go/issues/80)) ([773c56c](https://github.com/openlayer-ai/openlayer-go/commit/773c56c93858bfec9e73ccb22950919bd968d606))


### Bug Fixes

* **client:** clean up reader resources ([d608b57](https://github.com/openlayer-ai/openlayer-go/commit/d608b57fd158f1bdfeeeea89e9ebb911c3669b35))
* **client:** correctly update body in WithJSONSet ([e59a983](https://github.com/openlayer-ai/openlayer-go/commit/e59a9832b8dddd3612fdf9492b6c095ecd9eb589))
* handle empty bodies in WithJSONSet ([dfd870a](https://github.com/openlayer-ai/openlayer-go/commit/dfd870a9923922602af776bf62ae5da696c4f9f1))


### Chores

* **ci:** add timeout thresholds for CI jobs ([a41c7b6](https://github.com/openlayer-ai/openlayer-go/commit/a41c7b6366a1b1b380fa714ac527cd4c8a019112))
* **ci:** only use depot for staging repos ([abc6c92](https://github.com/openlayer-ai/openlayer-go/commit/abc6c92db4e9cbfd5e199571c521f9c3a7b02c73))
* **docs:** document pre-request options ([6cc0281](https://github.com/openlayer-ai/openlayer-go/commit/6cc028136dc67f77881c5c1f4d466254694a0679))
* **docs:** grammar improvements ([4dd364b](https://github.com/openlayer-ai/openlayer-go/commit/4dd364b2f9a3382a6b4880a3dc30521d14403a9b))
* improve devcontainer setup ([1322a10](https://github.com/openlayer-ai/openlayer-go/commit/1322a105c963120fab6553a83ef811cddc333e5d))
* **internal:** codegen related update ([8ad39a5](https://github.com/openlayer-ai/openlayer-go/commit/8ad39a515b9bc3c3a752682500f37193c6f44a70))
* **internal:** expand CI branch coverage ([14067c7](https://github.com/openlayer-ai/openlayer-go/commit/14067c7e9dfa7c1bcdfb9ded29ce3eaea12d1bd2))
* **internal:** reduce CI branch coverage ([08c00ca](https://github.com/openlayer-ai/openlayer-go/commit/08c00cadee41cd7f6b3bd2873a0a1f18ba997a99))
* make go mod tidy continue on error ([d402983](https://github.com/openlayer-ai/openlayer-go/commit/d402983d853f087862dd3676f7f7bd6fdddd9a96))
* **tests:** improve enum examples ([#82](https://github.com/openlayer-ai/openlayer-go/issues/82)) ([c4047ae](https://github.com/openlayer-ai/openlayer-go/commit/c4047aeebaf1c7ae8aa176572a1615746bee8bb5))


### Documentation

* update documentation links to be more uniform ([6245eac](https://github.com/openlayer-ai/openlayer-go/commit/6245eac124bdfa974d36963c293b4f6d1b791021))

## 0.1.0-alpha.16 (2025-04-03)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#78](https://github.com/openlayer-ai/openlayer-go/issues/78)) ([7c32753](https://github.com/openlayer-ai/openlayer-go/commit/7c327533032a42dfeb762a95138e940deb636337))
* **test:** return early after test failure ([#76](https://github.com/openlayer-ai/openlayer-go/issues/76)) ([e3db15b](https://github.com/openlayer-ai/openlayer-go/commit/e3db15bf6fc326d760dd2a3dec8c2a1557243bf4))


### Chores

* add request options to client tests ([#75](https://github.com/openlayer-ai/openlayer-go/issues/75)) ([47bbe8f](https://github.com/openlayer-ai/openlayer-go/commit/47bbe8f6fc702b16b17662158b89e447545834d5))
* **docs:** improve security documentation ([#74](https://github.com/openlayer-ai/openlayer-go/issues/74)) ([ae864fb](https://github.com/openlayer-ai/openlayer-go/commit/ae864fbc0ac02aac9944cfde9974c1966a799871))
* fix typos ([#77](https://github.com/openlayer-ai/openlayer-go/issues/77)) ([d5225a3](https://github.com/openlayer-ai/openlayer-go/commit/d5225a326c7781b7728234535403b76b2ef91a41))
* **internal:** remove workflow ([466df7b](https://github.com/openlayer-ai/openlayer-go/commit/466df7bd80c4a89431d6f8db27bbac59b6b30dff))

## 0.1.0-alpha.15 (2025-03-14)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **client:** accept RFC6838 JSON content types ([#68](https://github.com/openlayer-ai/openlayer-go/issues/68)) ([6596131](https://github.com/openlayer-ai/openlayer-go/commit/659613119e8c7b1574c5613daa610311213957ad))
* **client:** improve default client options support ([#70](https://github.com/openlayer-ai/openlayer-go/issues/70)) ([b22284c](https://github.com/openlayer-ai/openlayer-go/commit/b22284c7b950b30bd922aed075df97038f54d09c))


### Chores

* **internal:** codegen related update ([#66](https://github.com/openlayer-ai/openlayer-go/issues/66)) ([695347c](https://github.com/openlayer-ai/openlayer-go/commit/695347c6b725a93db08983610105f81e435bb913))
* **internal:** remove extra empty newlines ([#71](https://github.com/openlayer-ai/openlayer-go/issues/71)) ([44dc1c3](https://github.com/openlayer-ai/openlayer-go/commit/44dc1c324cd0c79bebfe46ee50da4b800eea0a2e))


### Refactors

* tidy up dependencies ([#69](https://github.com/openlayer-ai/openlayer-go/issues/69)) ([c676b56](https://github.com/openlayer-ai/openlayer-go/commit/c676b5617241008228c0a68138a09b52810506f6))

## 0.1.0-alpha.14 (2025-03-13)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* **api:** add endpoint to retrieve commit by id ([#64](https://github.com/openlayer-ai/openlayer-go/issues/64)) ([201cb4a](https://github.com/openlayer-ai/openlayer-go/commit/201cb4acd0e572cecb46ffdc29979c38a52846fb))
* **api:** api update ([1d5c8e7](https://github.com/openlayer-ai/openlayer-go/commit/1d5c8e7f9a98ab806517fbd7c2ac8c6ac456f04d))
* **client:** send `X-Stainless-Timeout` header ([d9948ef](https://github.com/openlayer-ai/openlayer-go/commit/d9948ef4b8b56ed71ddd2ea0f4519bf40dc5d52f))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([f7b96d9](https://github.com/openlayer-ai/openlayer-go/commit/f7b96d9ebe96937df548fe0797fbbaf8ba77109a))
* do not call path.Base on ContentType ([0b98d8e](https://github.com/openlayer-ai/openlayer-go/commit/0b98d8e38cb8716cc61ee62f9aa4b58250a4b581))
* fix early cancel when RequestTimeout is provided for streaming requests ([f0c9a39](https://github.com/openlayer-ai/openlayer-go/commit/f0c9a39e93d5aab9a2f78b95f1a6ee0aa27632ca))
* fix interface implementation stub names for unions ([0d8c74f](https://github.com/openlayer-ai/openlayer-go/commit/0d8c74f4d02b5d154357d98e892ac074470b474b))
* fix unicode encoding for json ([cdaee21](https://github.com/openlayer-ai/openlayer-go/commit/cdaee2125cb3c27f1a94cefc70d87ad5e12b79cc))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([3295c9b](https://github.com/openlayer-ai/openlayer-go/commit/3295c9bcc892bf7db4f1068120be3a86b86b1281))
* **internal:** codegen related update ([57b756f](https://github.com/openlayer-ai/openlayer-go/commit/57b756f1e31665115bc2cac4d7527f21965376ad))
* **internal:** codegen related update ([c2a18b7](https://github.com/openlayer-ai/openlayer-go/commit/c2a18b7d8ebe2df5f23428be410172a0dbd330cd))
* **internal:** codegen related update ([5582642](https://github.com/openlayer-ai/openlayer-go/commit/5582642f1d2b958566a842f4c9663baea39563cc))
* **internal:** fix devcontainers setup ([1d42cd7](https://github.com/openlayer-ai/openlayer-go/commit/1d42cd7adbc487450809a3da28cdb47cc11da3db))
* refactor client tests ([d2732fe](https://github.com/openlayer-ai/openlayer-go/commit/d2732fee321e2f829eafa4374c0c093d6e378f72))


### Documentation

* document raw responses ([1c66d2a](https://github.com/openlayer-ai/openlayer-go/commit/1c66d2a1e698e42799d142ff57ad1b2a51630b98))

## 0.1.0-alpha.13 (2025-01-02)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Chores

* **internal:** version bump ([#61](https://github.com/openlayer-ai/openlayer-go/issues/61)) ([f0934b2](https://github.com/openlayer-ai/openlayer-go/commit/f0934b2a2d0421134f7dee60dd1ebd655493ca70))

## 0.1.0-alpha.12 (2025-01-01)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Chores

* **internal:** codegen related update ([#59](https://github.com/openlayer-ai/openlayer-go/issues/59)) ([5df4e4d](https://github.com/openlayer-ai/openlayer-go/commit/5df4e4d53d5aa087e7e733a20dd52cd4783e5c2b))

## 0.1.0-alpha.11 (2024-12-19)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* **api:** api update ([#56](https://github.com/openlayer-ai/openlayer-go/issues/56)) ([c83eaec](https://github.com/openlayer-ai/openlayer-go/commit/c83eaec35bc52fbcb4110ad57e4ccc8be6daa224))

## 0.1.0-alpha.10 (2024-11-14)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Chores

* rebuild project due to codegen change ([#53](https://github.com/openlayer-ai/openlayer-go/issues/53)) ([5fbe9a0](https://github.com/openlayer-ai/openlayer-go/commit/5fbe9a0259a190c1cdefda2f4a9aba52ee77f915))

## 0.1.0-alpha.9 (2024-11-11)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** manual updates ([#49](https://github.com/openlayer-ai/openlayer-go/issues/49)) ([0c75483](https://github.com/openlayer-ai/openlayer-go/commit/0c754834bb8e3fbd493a4a656d785547c64bb532))
* **api:** update via SDK Studio ([#46](https://github.com/openlayer-ai/openlayer-go/issues/46)) ([0c269b0](https://github.com/openlayer-ai/openlayer-go/commit/0c269b042bba8934c5973214d2d665dc1cf362fc))
* **api:** update via SDK Studio ([#48](https://github.com/openlayer-ai/openlayer-go/issues/48)) ([27a5a99](https://github.com/openlayer-ai/openlayer-go/commit/27a5a9977c7741bbfcfe7a62d8e32a1a512cabac))


### Chores

* custom code changes ([#50](https://github.com/openlayer-ai/openlayer-go/issues/50)) ([3b63bd4](https://github.com/openlayer-ai/openlayer-go/commit/3b63bd474c0ac3d8404d8ed6196007f6eea1bfd0))
* rebuild project due to codegen change ([#51](https://github.com/openlayer-ai/openlayer-go/issues/51)) ([43b0dae](https://github.com/openlayer-ai/openlayer-go/commit/43b0dae82426763900cdbd37eccf2433c0d093b6))

## 0.1.0-alpha.8 (2024-09-24)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** OpenAPI spec update via Stainless API ([#45](https://github.com/openlayer-ai/openlayer-go/issues/45)) ([0f1932e](https://github.com/openlayer-ai/openlayer-go/commit/0f1932ea22a9f79f2f6552051a561ab21f049aa3))
* **api:** update via SDK Studio ([#42](https://github.com/openlayer-ai/openlayer-go/issues/42)) ([a0011e2](https://github.com/openlayer-ai/openlayer-go/commit/a0011e2a8fd344a97eff3f9c4dcc5a3a847677ee))
* **api:** update via SDK Studio ([#44](https://github.com/openlayer-ai/openlayer-go/issues/44)) ([258b832](https://github.com/openlayer-ai/openlayer-go/commit/258b8323d06d2c6ba62075553eeccb731c7895f7))

## 0.1.0-alpha.7 (2024-08-14)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Chores

* bump Go to v1.21 ([#38](https://github.com/openlayer-ai/openlayer-go/issues/38)) ([4552823](https://github.com/openlayer-ai/openlayer-go/commit/4552823edf865a06b312274e66135fcdcd36d758))
* **examples:** minor formatting changes ([#40](https://github.com/openlayer-ai/openlayer-go/issues/40)) ([aa3b623](https://github.com/openlayer-ai/openlayer-go/commit/aa3b6230b77e63ee0b644f81c5f683a28fcf76c0))

## 0.1.0-alpha.6 (2024-08-14)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Bug Fixes

* deserialization of struct unions that implement json.Unmarshaler ([#36](https://github.com/openlayer-ai/openlayer-go/issues/36)) ([bc92d1d](https://github.com/openlayer-ai/openlayer-go/commit/bc92d1dc5ac47f1f7fa17168112d100ac4b0bb65))


### Chores

* **ci:** bump prism mock server version ([#35](https://github.com/openlayer-ai/openlayer-go/issues/35)) ([a6ad713](https://github.com/openlayer-ai/openlayer-go/commit/a6ad713106b68733cc84391c9d6348a20a72eab1))
* **internal:** codegen related update ([#33](https://github.com/openlayer-ai/openlayer-go/issues/33)) ([4b4e1cd](https://github.com/openlayer-ai/openlayer-go/commit/4b4e1cde527d53cf6c7548bb27ccd6bb93c3c108))
* **internal:** version bump ([#32](https://github.com/openlayer-ai/openlayer-go/issues/32)) ([b0cfc04](https://github.com/openlayer-ai/openlayer-go/commit/b0cfc04e22de46846a7a0d39785b9ba0e99f23f3))

## 0.1.0-alpha.5 (2024-08-14)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** OpenAPI spec update via Stainless API ([#19](https://github.com/openlayer-ai/openlayer-go/issues/19)) ([b142e08](https://github.com/openlayer-ai/openlayer-go/commit/b142e0847a75b53b079925803bcc75c7c79e2bc0))
* **api:** OpenAPI spec update via Stainless API ([#23](https://github.com/openlayer-ai/openlayer-go/issues/23)) ([2529ae2](https://github.com/openlayer-ai/openlayer-go/commit/2529ae228b86f98a147cae383ffdbe74188d8b2c))
* **api:** OpenAPI spec update via Stainless API ([#24](https://github.com/openlayer-ai/openlayer-go/issues/24)) ([ca8e5fc](https://github.com/openlayer-ai/openlayer-go/commit/ca8e5fcf6aa42d026db5e174e94b98d25496af4d))
* **api:** update via SDK Studio ([#21](https://github.com/openlayer-ai/openlayer-go/issues/21)) ([f83d4cf](https://github.com/openlayer-ai/openlayer-go/commit/f83d4cf7708169f3ca03e22940a09aab95c3a15d))
* **api:** update via SDK Studio ([#22](https://github.com/openlayer-ai/openlayer-go/issues/22)) ([4af16b0](https://github.com/openlayer-ai/openlayer-go/commit/4af16b0774291fa1d0646ba33e9e88fe1084216d))
* **api:** update via SDK Studio ([#28](https://github.com/openlayer-ai/openlayer-go/issues/28)) ([6d5de0c](https://github.com/openlayer-ai/openlayer-go/commit/6d5de0c9a769b76e83cd308224198f1887e6e848))
* **api:** update via SDK Studio ([#29](https://github.com/openlayer-ai/openlayer-go/issues/29)) ([9bd7ba4](https://github.com/openlayer-ai/openlayer-go/commit/9bd7ba402687a7aab4214a386cd0d9f526c674a1))
* **api:** update via SDK Studio ([#30](https://github.com/openlayer-ai/openlayer-go/issues/30)) ([fe250fe](https://github.com/openlayer-ai/openlayer-go/commit/fe250fe2f2b6ba4baeb26bc50f06dc7fbe113649))


### Chores

* **ci:** limit release doctor target branches ([#25](https://github.com/openlayer-ai/openlayer-go/issues/25)) ([e2d8fc3](https://github.com/openlayer-ai/openlayer-go/commit/e2d8fc3b3427c5830f246ef91c0dd8b7020680a0))
* **ci:** remove unused release doctor ([#26](https://github.com/openlayer-ai/openlayer-go/issues/26)) ([6103065](https://github.com/openlayer-ai/openlayer-go/commit/61030659c93ba65ef0f15117ff6821452c737893))
* **tests:** update prism version ([#31](https://github.com/openlayer-ai/openlayer-go/issues/31)) ([b414437](https://github.com/openlayer-ai/openlayer-go/commit/b41443718f2fc5b163924cb50bf799ee78d52529))

## 0.1.0-alpha.4 (2024-07-08)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** OpenAPI spec update via Stainless API ([be7c0b4](https://github.com/openlayer-ai/openlayer-go/commit/be7c0b479f4bbdcf19cf6783f73dc2d4aae820e3))
* **api:** OpenAPI spec update via Stainless API ([#13](https://github.com/openlayer-ai/openlayer-go/issues/13)) ([bb83bc0](https://github.com/openlayer-ai/openlayer-go/commit/bb83bc0c22f8727d3d3352b4826b2c24bb34c8d4))
* **api:** OpenAPI spec update via Stainless API ([#14](https://github.com/openlayer-ai/openlayer-go/issues/14)) ([c23f741](https://github.com/openlayer-ai/openlayer-go/commit/c23f74110985d5f91ed0bfa08418d763ba47d471))
* **api:** OpenAPI spec update via Stainless API ([#16](https://github.com/openlayer-ai/openlayer-go/issues/16)) ([7ae861b](https://github.com/openlayer-ai/openlayer-go/commit/7ae861b1f28161696c347c208425d41a27c2da44))
* **api:** update via SDK Studio ([#10](https://github.com/openlayer-ai/openlayer-go/issues/10)) ([25f7816](https://github.com/openlayer-ai/openlayer-go/commit/25f781652e3e9645c0dc776ec0bc695c2fb68b2d))
* **api:** update via SDK Studio ([#15](https://github.com/openlayer-ai/openlayer-go/issues/15)) ([7b362a5](https://github.com/openlayer-ai/openlayer-go/commit/7b362a56e43a50f7c2229e39bc9fec649bd2aad0))

## 0.1.0-alpha.3 (2024-07-05)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([4cb6f92](https://github.com/openlayer-ai/openlayer-go/commit/4cb6f92efa8607ad5d1723e2b1bd8603a559cb34))
* **api:** update via SDK Studio ([50a2f3f](https://github.com/openlayer-ai/openlayer-go/commit/50a2f3fe50afa3227c55be1ba0a023711144aa34))
* **api:** update via SDK Studio ([#7](https://github.com/openlayer-ai/openlayer-go/issues/7)) ([9ea6352](https://github.com/openlayer-ai/openlayer-go/commit/9ea6352504266dcc33f3f083956130eabba809e8))


### Chores

* go live ([#9](https://github.com/openlayer-ai/openlayer-go/issues/9)) ([5ebc5e3](https://github.com/openlayer-ai/openlayer-go/commit/5ebc5e342ce1b239573aa227024094efb4af7aed))

## 0.1.0-alpha.2 (2024-06-07)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/openlayer-ai/openlayer-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** update via SDK Studio ([#4](https://github.com/openlayer-ai/openlayer-go/issues/4)) ([bf41bec](https://github.com/openlayer-ai/openlayer-go/commit/bf41bec163a95b3f83a7147bfed3c5027581ebda))

## 0.1.0-alpha.1 (2024-06-07)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/openlayer-ai/openlayer-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** OpenAPI spec update via Stainless API ([96fd32e](https://github.com/openlayer-ai/openlayer-go/commit/96fd32ed48543edc02daa2ab42e1108a171db9cd))
* **api:** OpenAPI spec update via Stainless API ([d57c0ae](https://github.com/openlayer-ai/openlayer-go/commit/d57c0ae7172c979daf3e42ca9b7104d53ec67f1a))
* **api:** OpenAPI spec update via Stainless API ([31c79dc](https://github.com/openlayer-ai/openlayer-go/commit/31c79dc6390877b054bf9dc626c74ac55e3133b6))
* **api:** OpenAPI spec update via Stainless API ([46e7470](https://github.com/openlayer-ai/openlayer-go/commit/46e7470e7c6c3d3ec0017e91f2e4b3d851af3259))
* **api:** update via SDK Studio ([04ebb3d](https://github.com/openlayer-ai/openlayer-go/commit/04ebb3d297a9dc7c932a0b383f5398f87b8564de))
* **api:** update via SDK Studio ([7c6a75e](https://github.com/openlayer-ai/openlayer-go/commit/7c6a75e841a731e8e9bae67d31af465f3d36095d))
* **api:** update via SDK Studio ([65994d2](https://github.com/openlayer-ai/openlayer-go/commit/65994d29a753cb503f60babbdd01bd94f7c4b656))
* **api:** update via SDK Studio ([c942498](https://github.com/openlayer-ai/openlayer-go/commit/c942498e7a09ea7cd52af0c8be6b5361de5eda8c))
* **api:** update via SDK Studio ([9c49f1e](https://github.com/openlayer-ai/openlayer-go/commit/9c49f1e52df5fababca4dae8084e3d93b0836452))
* **api:** update via SDK Studio ([398f60c](https://github.com/openlayer-ai/openlayer-go/commit/398f60c3161b2a5c29d900b6d628329443c4a537))
* **api:** update via SDK Studio ([fd04ba1](https://github.com/openlayer-ai/openlayer-go/commit/fd04ba1f7c462094e399614592d489a50f58fab4))
* **api:** update via SDK Studio ([00b9111](https://github.com/openlayer-ai/openlayer-go/commit/00b9111dfb0bff8d8a278ebddd5810a3daa54423))
* **api:** update via SDK Studio ([12e403c](https://github.com/openlayer-ai/openlayer-go/commit/12e403cc527a5ac1585861bb646710bf8f408dd9))
* **api:** update via SDK Studio ([a5c1abd](https://github.com/openlayer-ai/openlayer-go/commit/a5c1abda46d97efc1d538c4031eab559a3f09fd8))


### Chores

* configure new SDK language ([7751b7a](https://github.com/openlayer-ai/openlayer-go/commit/7751b7a940a74c8889106400750452b8d66d7b25))
* go live ([#1](https://github.com/openlayer-ai/openlayer-go/issues/1)) ([6bbd21a](https://github.com/openlayer-ai/openlayer-go/commit/6bbd21a6c51f074250e58afd7ec10c6668e97517))
* update SDK settings ([2366064](https://github.com/openlayer-ai/openlayer-go/commit/2366064a8af4450b0a57a0996b03a93da275d2ba))
* update SDK settings ([ccdd464](https://github.com/openlayer-ai/openlayer-go/commit/ccdd4641c80ab5199ef9d3fb2ad3a09ddbda3b4f))
* update SDK settings ([84ea1f7](https://github.com/openlayer-ai/openlayer-go/commit/84ea1f78705d4693c395d46a68b9c804b5b6e349))
* update SDK settings ([520a519](https://github.com/openlayer-ai/openlayer-go/commit/520a519b257d5de323a61bb0c4acdad7ee49ceb1))
