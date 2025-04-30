# Changelog

## 1.13.1 (2025-04-30)

Full Changelog: [v1.13.0...v1.13.1](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.13.0...v1.13.1)

### Bug Fixes

* handle empty bodies in WithJSONSet ([7a3b436](https://github.com/terminaldotshop/terminal-sdk-go/commit/7a3b4362b51dc35a4ef17279358a767b9bdd78ad))

## 1.13.0 (2025-04-24)

Full Changelog: [v1.12.0...v1.13.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.12.0...v1.13.0)

### Features

* **api:** product variant descriptions ([3082bf9](https://github.com/terminaldotshop/terminal-sdk-go/commit/3082bf9a5f8bd4241d03b92f9a39f292783f79d7))

## 1.12.0 (2025-04-24)

Full Changelog: [v1.11.0...v1.12.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.11.0...v1.12.0)

### Features

* **api:** global region, typed tracking status, variant tags ([a86f157](https://github.com/terminaldotshop/terminal-sdk-go/commit/a86f1574e07b61153855155280764048e0807edf))


### Chores

* **ci:** only use depot for staging repos ([6486fa8](https://github.com/terminaldotshop/terminal-sdk-go/commit/6486fa80efe15f6fbb15f82f05c98f99b56f06be))
* **internal:** codegen related update ([9ce0cb9](https://github.com/terminaldotshop/terminal-sdk-go/commit/9ce0cb966ee1acac5bbce9051ec9d485ac13baf1))

## 1.11.0 (2025-04-23)

Full Changelog: [v1.10.0...v1.11.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.10.0...v1.11.0)

### Features

* **api:** include price on subscriptions ([0ae77c5](https://github.com/terminaldotshop/terminal-sdk-go/commit/0ae77c58807f9fc59a0583ffacfc987d030b5cc2))


### Chores

* **ci:** add timeout thresholds for CI jobs ([b973633](https://github.com/terminaldotshop/terminal-sdk-go/commit/b9736336388490ac7cc264b1613bfd6ebec2ba25))

## 1.10.0 (2025-04-18)

Full Changelog: [v1.9.0...v1.10.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.9.0...v1.10.0)

### Features

* **api:** update subscription route ([7d0c254](https://github.com/terminaldotshop/terminal-sdk-go/commit/7d0c2543a0572b399ccd2877c90105cb680454a1))
* **client:** add support for reading base URL from environment variable ([431d9c8](https://github.com/terminaldotshop/terminal-sdk-go/commit/431d9c850b57d89557359986c200dcd5f7b9c16c))


### Chores

* **docs:** document pre-request options ([675cfb4](https://github.com/terminaldotshop/terminal-sdk-go/commit/675cfb46a4dfb847b58ab396cb082b52c9c4688e))


### Documentation

* update documentation links to be more uniform ([31d4a93](https://github.com/terminaldotshop/terminal-sdk-go/commit/31d4a93bf11aa66162b40bc9b9d476fba4bf09b7))

## 1.9.0 (2025-04-14)

Full Changelog: [v1.8.0...v1.9.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.8.0...v1.9.0)

### Features

* **api:** include created timestamps ([cc2cabd](https://github.com/terminaldotshop/terminal-sdk-go/commit/cc2cabdab8c3353a9403a5dcc992d9e78946ea48))


### Chores

* **internal:** expand CI branch coverage ([89374f1](https://github.com/terminaldotshop/terminal-sdk-go/commit/89374f18d74e349fc3f9308a6c559ddf1f3e4981))
* **internal:** reduce CI branch coverage ([ff6f722](https://github.com/terminaldotshop/terminal-sdk-go/commit/ff6f722ca5de9c66a721b70cc9c7997fa32a1db4))

## 1.8.0 (2025-04-09)

Full Changelog: [v1.7.3...v1.8.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.7.3...v1.8.0)

### Features

* **api:** include shipment tracking info on order ([#282](https://github.com/terminaldotshop/terminal-sdk-go/issues/282)) ([f52506f](https://github.com/terminaldotshop/terminal-sdk-go/commit/f52506f8f7ef73eeae8256b9242d70ef366503ae))
* **client:** support custom http clients ([#280](https://github.com/terminaldotshop/terminal-sdk-go/issues/280)) ([3a60175](https://github.com/terminaldotshop/terminal-sdk-go/commit/3a60175bd2c0ed835f59ab55f0ca15c25b54f35d))

## 1.7.3 (2025-04-03)

Full Changelog: [v1.7.2...v1.7.3](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.7.2...v1.7.3)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#277](https://github.com/terminaldotshop/terminal-sdk-go/issues/277)) ([01da641](https://github.com/terminaldotshop/terminal-sdk-go/commit/01da641d51a572a99569469faff208025ac2f619))

## 1.7.2 (2025-04-02)

Full Changelog: [v1.7.1...v1.7.2](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.7.1...v1.7.2)

### Bug Fixes

* pluralize `list` response variables ([#274](https://github.com/terminaldotshop/terminal-sdk-go/issues/274)) ([2fcc263](https://github.com/terminaldotshop/terminal-sdk-go/commit/2fcc263bec413f2313c605104a6627598cc61250))

## 1.7.1 (2025-04-01)

Full Changelog: [v1.7.0...v1.7.1](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.7.0...v1.7.1)

### Bug Fixes

* remove trailing / for environments ([#272](https://github.com/terminaldotshop/terminal-sdk-go/issues/272)) ([45b4237](https://github.com/terminaldotshop/terminal-sdk-go/commit/45b4237bf93642e39bbb7fc5677d62a2c28c615b))


### Chores

* fix typos ([#269](https://github.com/terminaldotshop/terminal-sdk-go/issues/269)) ([68bcda1](https://github.com/terminaldotshop/terminal-sdk-go/commit/68bcda1b16f8d6080df071eedb0d62e370a2995e))

## 1.7.0 (2025-03-26)

Full Changelog: [v1.6.0...v1.7.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.6.0...v1.7.0)

### Features

* **client:** improve default client options support ([#262](https://github.com/terminaldotshop/terminal-sdk-go/issues/262)) ([8363df7](https://github.com/terminaldotshop/terminal-sdk-go/commit/8363df71a37e9cea79811f11310f38b264f3feb1))


### Bug Fixes

* **test:** return early after test failure ([#267](https://github.com/terminaldotshop/terminal-sdk-go/issues/267)) ([023997e](https://github.com/terminaldotshop/terminal-sdk-go/commit/023997e49a770b3924f6844ab8a53302a21ff3a6))


### Chores

* add request options to client tests ([#266](https://github.com/terminaldotshop/terminal-sdk-go/issues/266)) ([7baf25a](https://github.com/terminaldotshop/terminal-sdk-go/commit/7baf25a31fc2fb1130cedac39f30a7623755caef))
* **docs:** improve security documentation ([#265](https://github.com/terminaldotshop/terminal-sdk-go/issues/265)) ([f0b5e0e](https://github.com/terminaldotshop/terminal-sdk-go/commit/f0b5e0edff094ad3a6e14d3d3344c6a163d3e059))
* **internal:** remove extra empty newlines ([#264](https://github.com/terminaldotshop/terminal-sdk-go/issues/264)) ([497456f](https://github.com/terminaldotshop/terminal-sdk-go/commit/497456f72051b08f2666ced1937655ffa38c191e))

## 1.6.0 (2025-03-13)

Full Changelog: [v1.5.0...v1.6.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.5.0...v1.6.0)

### Features

* **api:** region model ([#259](https://github.com/terminaldotshop/terminal-sdk-go/issues/259)) ([e95fc8b](https://github.com/terminaldotshop/terminal-sdk-go/commit/e95fc8b2b9f019e259fbe5c6a10db99207999b10))

## 1.5.0 (2025-03-13)

Full Changelog: [v1.4.0...v1.5.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.4.0...v1.5.0)

### Features

* **api:** region type ([#256](https://github.com/terminaldotshop/terminal-sdk-go/issues/256)) ([20bf8a4](https://github.com/terminaldotshop/terminal-sdk-go/commit/20bf8a41c632dae0d8abb0dbb62cf36d164a0d48))

## 1.4.0 (2025-03-13)

Full Changelog: [v1.3.0...v1.4.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.3.0...v1.4.0)

### Features

* **api:** add region to `GET /view/init` ([#253](https://github.com/terminaldotshop/terminal-sdk-go/issues/253)) ([0e625b8](https://github.com/terminaldotshop/terminal-sdk-go/commit/0e625b8a779055f2d66722ab7ae1d37111a317e0))

## 1.3.0 (2025-03-13)

Full Changelog: [v1.2.0...v1.3.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.2.0...v1.3.0)

### Features

* **api:** remove gift cards ([#250](https://github.com/terminaldotshop/terminal-sdk-go/issues/250)) ([4c567d0](https://github.com/terminaldotshop/terminal-sdk-go/commit/4c567d079166b3b1316e0537851d391a2f4b5158))

## 1.2.0 (2025-03-13)

Full Changelog: [v1.1.0...v1.2.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.1.0...v1.2.0)

### Features

* **api:** clear cart api ([#247](https://github.com/terminaldotshop/terminal-sdk-go/issues/247)) ([1626076](https://github.com/terminaldotshop/terminal-sdk-go/commit/162607648438a9d41c3442318fecfc34ea68dbf4))

## 1.1.0 (2025-03-11)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v1.0.0...v1.1.0)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#242](https://github.com/terminaldotshop/terminal-sdk-go/issues/242)) ([aea5e57](https://github.com/terminaldotshop/terminal-sdk-go/commit/aea5e571f95dbffb0327bc4159af389e05be2082))
* **api:** gift cards ([#245](https://github.com/terminaldotshop/terminal-sdk-go/issues/245)) ([cfd2f1b](https://github.com/terminaldotshop/terminal-sdk-go/commit/cfd2f1b8ae5e253a99be0f9b7206a61dc272cb61))
* **client:** accept RFC6838 JSON content types ([#243](https://github.com/terminaldotshop/terminal-sdk-go/issues/243)) ([70a0209](https://github.com/terminaldotshop/terminal-sdk-go/commit/70a0209aacdba5b3c60a59f3c2a3af6941f10f25))
* **client:** allow custom baseurls without trailing slash ([#240](https://github.com/terminaldotshop/terminal-sdk-go/issues/240)) ([e220ba3](https://github.com/terminaldotshop/terminal-sdk-go/commit/e220ba3a51f63f79ec4c2ce0a2f68fae857ad00c))


### Refactors

* tidy up dependencies ([#244](https://github.com/terminaldotshop/terminal-sdk-go/issues/244)) ([030e080](https://github.com/terminaldotshop/terminal-sdk-go/commit/030e0802286970dfac9af324d9e42f0e2ef91ec6))

## 1.0.0 (2025-03-07)

Full Changelog: [v0.1.0-alpha.61...v1.0.0](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.61...v1.0.0)

### Features

* **api:** manual updates ([#237](https://github.com/terminaldotshop/terminal-sdk-go/issues/237)) ([1dae804](https://github.com/terminaldotshop/terminal-sdk-go/commit/1dae80452c795546dae0ffae9b85ec65e40019a1))

## 0.1.0-alpha.61 (2025-02-28)

Full Changelog: [v0.1.0-alpha.60...v0.1.0-alpha.61](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.60...v0.1.0-alpha.61)

### Features

* **api:** manual updates ([#234](https://github.com/terminaldotshop/terminal-sdk-go/issues/234)) ([2f5afb6](https://github.com/terminaldotshop/terminal-sdk-go/commit/2f5afb69264d64e8520abaff8545b669306a4b34))

## 0.1.0-alpha.60 (2025-02-28)

Full Changelog: [v0.1.0-alpha.59...v0.1.0-alpha.60](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.59...v0.1.0-alpha.60)

### Features

* **api:** manual updates ([#232](https://github.com/terminaldotshop/terminal-sdk-go/issues/232)) ([b175e18](https://github.com/terminaldotshop/terminal-sdk-go/commit/b175e184b2e68402ffb2e50b4d064a7c1bc15458))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#230](https://github.com/terminaldotshop/terminal-sdk-go/issues/230)) ([066a70a](https://github.com/terminaldotshop/terminal-sdk-go/commit/066a70a65f790f862eb03fc49d0335d24326f31e))

## 0.1.0-alpha.59 (2025-02-26)

Full Changelog: [v0.1.0-alpha.58...v0.1.0-alpha.59](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.58...v0.1.0-alpha.59)

### Features

* **api:** manual updates ([#227](https://github.com/terminaldotshop/terminal-sdk-go/issues/227)) ([423ce20](https://github.com/terminaldotshop/terminal-sdk-go/commit/423ce20c96308bb363e897d86fbac87a086fafcb))

## 0.1.0-alpha.58 (2025-02-26)

Full Changelog: [v0.1.0-alpha.57...v0.1.0-alpha.58](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.57...v0.1.0-alpha.58)

### Features

* **api:** manual updates ([#225](https://github.com/terminaldotshop/terminal-sdk-go/issues/225)) ([761d739](https://github.com/terminaldotshop/terminal-sdk-go/commit/761d7398948215dfe1b54d60dffe08c9f4df6497))


### Chores

* **internal:** codegen related update ([#223](https://github.com/terminaldotshop/terminal-sdk-go/issues/223)) ([3e85cef](https://github.com/terminaldotshop/terminal-sdk-go/commit/3e85cefee52fa4c6e12d050727ac6558ebc844a5))

## 0.1.0-alpha.57 (2025-02-25)

Full Changelog: [v0.1.0-alpha.56...v0.1.0-alpha.57](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.56...v0.1.0-alpha.57)

### Features

* **api:** manual updates ([#220](https://github.com/terminaldotshop/terminal-sdk-go/issues/220)) ([2cc0d42](https://github.com/terminaldotshop/terminal-sdk-go/commit/2cc0d4218295f0d7067c00ef879df246d211cb0a))

## 0.1.0-alpha.56 (2025-02-17)

Full Changelog: [v0.1.0-alpha.55...v0.1.0-alpha.56](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.55...v0.1.0-alpha.56)

### Features

* **api:** manual updates ([#217](https://github.com/terminaldotshop/terminal-sdk-go/issues/217)) ([225d656](https://github.com/terminaldotshop/terminal-sdk-go/commit/225d656544b5a136e266c2d91d31f0b3c79dee92))

## 0.1.0-alpha.55 (2025-02-16)

Full Changelog: [v0.1.0-alpha.54...v0.1.0-alpha.55](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.54...v0.1.0-alpha.55)

### Features

* **api:** manual updates ([#214](https://github.com/terminaldotshop/terminal-sdk-go/issues/214)) ([a5e3cb5](https://github.com/terminaldotshop/terminal-sdk-go/commit/a5e3cb522495e237818ec4b3658890e06fd3c75e))

## 0.1.0-alpha.54 (2025-02-16)

Full Changelog: [v0.1.0-alpha.53...v0.1.0-alpha.54](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.53...v0.1.0-alpha.54)

### Features

* **api:** manual updates ([#211](https://github.com/terminaldotshop/terminal-sdk-go/issues/211)) ([f6e3e07](https://github.com/terminaldotshop/terminal-sdk-go/commit/f6e3e0788854b5d26818237d3f88956381c7ca67))

## 0.1.0-alpha.53 (2025-02-16)

Full Changelog: [v0.1.0-alpha.52...v0.1.0-alpha.53](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.52...v0.1.0-alpha.53)

### Features

* **api:** manual updates ([#208](https://github.com/terminaldotshop/terminal-sdk-go/issues/208)) ([8fd6ee6](https://github.com/terminaldotshop/terminal-sdk-go/commit/8fd6ee6b4cbc183e4f7a5a111469e4851b7ead30))

## 0.1.0-alpha.52 (2025-02-15)

Full Changelog: [v0.1.0-alpha.51...v0.1.0-alpha.52](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.51...v0.1.0-alpha.52)

### Bug Fixes

* **client:** don't truncate manually specified filenames ([#205](https://github.com/terminaldotshop/terminal-sdk-go/issues/205)) ([dc87fdb](https://github.com/terminaldotshop/terminal-sdk-go/commit/dc87fdb9b3031515f5155df355ca413690996c68))

## 0.1.0-alpha.51 (2025-02-11)

Full Changelog: [v0.1.0-alpha.50...v0.1.0-alpha.51](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.50...v0.1.0-alpha.51)

### Bug Fixes

* do not call path.Base on ContentType ([#202](https://github.com/terminaldotshop/terminal-sdk-go/issues/202)) ([ead9d41](https://github.com/terminaldotshop/terminal-sdk-go/commit/ead9d4189493102383f087f47fb6c0962e799d4a))

## 0.1.0-alpha.50 (2025-02-07)

Full Changelog: [v0.1.0-alpha.49...v0.1.0-alpha.50](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.49...v0.1.0-alpha.50)

### Bug Fixes

* fix early cancel when RequestTimeout is provided for streaming requests ([#199](https://github.com/terminaldotshop/terminal-sdk-go/issues/199)) ([60ace27](https://github.com/terminaldotshop/terminal-sdk-go/commit/60ace278088ebffabde0bfe7989b01c9fab30d81))

## 0.1.0-alpha.49 (2025-02-06)

Full Changelog: [v0.1.0-alpha.48...v0.1.0-alpha.49](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.48...v0.1.0-alpha.49)

### Features

* **api:** manual updates ([#197](https://github.com/terminaldotshop/terminal-sdk-go/issues/197)) ([83b2530](https://github.com/terminaldotshop/terminal-sdk-go/commit/83b2530720abd2624dc64272a769f62489a21455))
* **client:** send `X-Stainless-Timeout` header ([#195](https://github.com/terminaldotshop/terminal-sdk-go/issues/195)) ([3c87644](https://github.com/terminaldotshop/terminal-sdk-go/commit/3c876448a43571f45505b8e7f20dbdf4d019ea48))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#196](https://github.com/terminaldotshop/terminal-sdk-go/issues/196)) ([aca6c1d](https://github.com/terminaldotshop/terminal-sdk-go/commit/aca6c1d440403cdeab878aa9868fbaac62e322f5))


### Documentation

* document raw responses ([#193](https://github.com/terminaldotshop/terminal-sdk-go/issues/193)) ([81f2f8c](https://github.com/terminaldotshop/terminal-sdk-go/commit/81f2f8c59c5f762a0e279459d7d0dd30b68a285d))

## 0.1.0-alpha.48 (2025-02-01)

Full Changelog: [v0.1.0-alpha.47...v0.1.0-alpha.48](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.47...v0.1.0-alpha.48)

### Bug Fixes

* fix unicode encoding for json ([#191](https://github.com/terminaldotshop/terminal-sdk-go/issues/191)) ([210268d](https://github.com/terminaldotshop/terminal-sdk-go/commit/210268dd0c0027f4b22f3b01ad592c88fb39c239))


### Chores

* refactor client tests ([#189](https://github.com/terminaldotshop/terminal-sdk-go/issues/189)) ([85d43c0](https://github.com/terminaldotshop/terminal-sdk-go/commit/85d43c054441df93ecc40c480d0b3f40ee67a30f))

## 0.1.0-alpha.47 (2025-01-21)

Full Changelog: [v0.1.0-alpha.46...v0.1.0-alpha.47](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.46...v0.1.0-alpha.47)

### Bug Fixes

* fix apijson.Port for embedded structs ([#186](https://github.com/terminaldotshop/terminal-sdk-go/issues/186)) ([a695598](https://github.com/terminaldotshop/terminal-sdk-go/commit/a695598d4461c70013b445ce2e786a922f7cbc47))
* fix apijson.Port for embedded structs ([#187](https://github.com/terminaldotshop/terminal-sdk-go/issues/187)) ([5163450](https://github.com/terminaldotshop/terminal-sdk-go/commit/516345068fab4c658797d41c87cae3a4825defd7))


### Chores

* **internal:** codegen related update ([#181](https://github.com/terminaldotshop/terminal-sdk-go/issues/181)) ([0159817](https://github.com/terminaldotshop/terminal-sdk-go/commit/0159817807f010f3a35dcd775b646e187a6ae292))
* **internal:** codegen related update ([#183](https://github.com/terminaldotshop/terminal-sdk-go/issues/183)) ([a001347](https://github.com/terminaldotshop/terminal-sdk-go/commit/a001347f7517921e581d0deaa1ffab650266e390))
* **internal:** codegen related update ([#184](https://github.com/terminaldotshop/terminal-sdk-go/issues/184)) ([6280ce0](https://github.com/terminaldotshop/terminal-sdk-go/commit/6280ce06138426e1a618c395168e5d12cd34fe9d))
* **internal:** codegen related update ([#185](https://github.com/terminaldotshop/terminal-sdk-go/issues/185)) ([300e52d](https://github.com/terminaldotshop/terminal-sdk-go/commit/300e52d20da3b22cd8d47994bf509ecde62b18e3))

## 0.1.0-alpha.46 (2024-12-17)

Full Changelog: [v0.1.0-alpha.45...v0.1.0-alpha.46](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.45...v0.1.0-alpha.46)

### Features

* **api:** manual updates ([#176](https://github.com/terminaldotshop/terminal-sdk-go/issues/176)) ([d2a7d27](https://github.com/terminaldotshop/terminal-sdk-go/commit/d2a7d27fc8bd7446f1511c290b0406cf4d78005a))

## 0.1.0-alpha.45 (2024-12-17)

Full Changelog: [v0.1.0-alpha.44...v0.1.0-alpha.45](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.44...v0.1.0-alpha.45)

### Features

* **api:** manual updates ([#172](https://github.com/terminaldotshop/terminal-sdk-go/issues/172)) ([930f6af](https://github.com/terminaldotshop/terminal-sdk-go/commit/930f6afe1e5ecb6ff85b17ce6f5e1d7e80c7450c))
* **api:** manual updates ([#174](https://github.com/terminaldotshop/terminal-sdk-go/issues/174)) ([6374181](https://github.com/terminaldotshop/terminal-sdk-go/commit/6374181679c1d6cf85db126711f4d5d0a90224ce))

## 0.1.0-alpha.44 (2024-12-16)

Full Changelog: [v0.1.0-alpha.43...v0.1.0-alpha.44](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.43...v0.1.0-alpha.44)

### Features

* **api:** manual updates ([#169](https://github.com/terminaldotshop/terminal-sdk-go/issues/169)) ([4bc452a](https://github.com/terminaldotshop/terminal-sdk-go/commit/4bc452a445fd9b916d60101c782c811c759cbd99))

## 0.1.0-alpha.43 (2024-12-16)

Full Changelog: [v0.1.0-alpha.42...v0.1.0-alpha.43](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.42...v0.1.0-alpha.43)

### Features

* **api:** manual updates ([#166](https://github.com/terminaldotshop/terminal-sdk-go/issues/166)) ([3d1dc54](https://github.com/terminaldotshop/terminal-sdk-go/commit/3d1dc54ef1b9e574987991449ebc0832c03ae241))

## 0.1.0-alpha.42 (2024-12-16)

Full Changelog: [v0.1.0-alpha.41...v0.1.0-alpha.42](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.41...v0.1.0-alpha.42)

### Features

* **api:** manual updates ([#163](https://github.com/terminaldotshop/terminal-sdk-go/issues/163)) ([c48f870](https://github.com/terminaldotshop/terminal-sdk-go/commit/c48f87036c27f02e95447d432feaf616ca2514dc))

## 0.1.0-alpha.41 (2024-12-13)

Full Changelog: [v0.1.0-alpha.40...v0.1.0-alpha.41](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.40...v0.1.0-alpha.41)

### Features

* **api:** update via SDK Studio ([#161](https://github.com/terminaldotshop/terminal-sdk-go/issues/161)) ([5197fa6](https://github.com/terminaldotshop/terminal-sdk-go/commit/5197fa602cdfb8bec55b3d188f9efc7dd50e3711))


### Chores

* **internal:** version bump ([#159](https://github.com/terminaldotshop/terminal-sdk-go/issues/159)) ([1893630](https://github.com/terminaldotshop/terminal-sdk-go/commit/1893630ab9d2293cc7949edb3b0e889fa32fe1aa))

## 0.1.0-alpha.40 (2024-12-13)

Full Changelog: [v0.1.0-alpha.39...v0.1.0-alpha.40](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.39...v0.1.0-alpha.40)

### Features

* **api:** update via SDK Studio ([#155](https://github.com/terminaldotshop/terminal-sdk-go/issues/155)) ([d71d0b9](https://github.com/terminaldotshop/terminal-sdk-go/commit/d71d0b9d1380546efd39db04eda1db709a743353))
* **api:** update via SDK Studio ([#156](https://github.com/terminaldotshop/terminal-sdk-go/issues/156)) ([6bed197](https://github.com/terminaldotshop/terminal-sdk-go/commit/6bed197010ddbd06b71b61b66704222f6ce68f2b))


### Chores

* **internal:** version bump ([#158](https://github.com/terminaldotshop/terminal-sdk-go/issues/158)) ([ca09e12](https://github.com/terminaldotshop/terminal-sdk-go/commit/ca09e12df9db3ea013f88e8647454ed0d8c40ab6))

## 0.1.0-alpha.39 (2024-12-13)

Full Changelog: [v0.1.0-alpha.38...v0.1.0-alpha.39](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.38...v0.1.0-alpha.39)

### Features

* **api:** update via SDK Studio ([#152](https://github.com/terminaldotshop/terminal-sdk-go/issues/152)) ([327d6f5](https://github.com/terminaldotshop/terminal-sdk-go/commit/327d6f57757b183bc23eb2cbb07d85cf07fd41c7))

## 0.1.0-alpha.38 (2024-12-13)

Full Changelog: [v0.1.0-alpha.37...v0.1.0-alpha.38](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.37...v0.1.0-alpha.38)

### Features

* **api:** update via SDK Studio ([#149](https://github.com/terminaldotshop/terminal-sdk-go/issues/149)) ([aed8143](https://github.com/terminaldotshop/terminal-sdk-go/commit/aed814385eb95e059bc9d00a4cbba48a1dccf9f5))

## 0.1.0-alpha.37 (2024-12-13)

Full Changelog: [v0.1.0-alpha.36...v0.1.0-alpha.37](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.36...v0.1.0-alpha.37)

### Features

* **api:** update via SDK Studio ([#146](https://github.com/terminaldotshop/terminal-sdk-go/issues/146)) ([da75ca3](https://github.com/terminaldotshop/terminal-sdk-go/commit/da75ca35fa7949708092f085abe1add2060e0253))

## 0.1.0-alpha.36 (2024-12-13)

Full Changelog: [v0.1.0-alpha.35...v0.1.0-alpha.36](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.35...v0.1.0-alpha.36)

### Features

* **api:** api update ([#138](https://github.com/terminaldotshop/terminal-sdk-go/issues/138)) ([cafdba4](https://github.com/terminaldotshop/terminal-sdk-go/commit/cafdba42b7fb7c0fbe79030758fd29533e7367af))
* **api:** api update ([#140](https://github.com/terminaldotshop/terminal-sdk-go/issues/140)) ([d8695aa](https://github.com/terminaldotshop/terminal-sdk-go/commit/d8695aa6b50dd8e1ea73daafb308f43eee2d39c5))
* **api:** update via SDK Studio ([#141](https://github.com/terminaldotshop/terminal-sdk-go/issues/141)) ([b9c2ac9](https://github.com/terminaldotshop/terminal-sdk-go/commit/b9c2ac9f0e7caa38c6880c26094571b9777d51ba))
* **api:** update via SDK Studio ([#142](https://github.com/terminaldotshop/terminal-sdk-go/issues/142)) ([3c4c9cc](https://github.com/terminaldotshop/terminal-sdk-go/commit/3c4c9cc59bdacabbfd664db7a01a5d144cf63324))
* **api:** update via SDK Studio ([#143](https://github.com/terminaldotshop/terminal-sdk-go/issues/143)) ([51d4eed](https://github.com/terminaldotshop/terminal-sdk-go/commit/51d4eedb02b0bf5586c9ce77dfaea3774f315526))


### Chores

* configure new SDK language ([2a22ad9](https://github.com/terminaldotshop/terminal-sdk-go/commit/2a22ad900edceeb68025c2fd0324d500960f8f05))

## 0.1.0-alpha.35 (2024-12-04)

Full Changelog: [v0.1.0-alpha.34...v0.1.0-alpha.35](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.34...v0.1.0-alpha.35)

### Features

* **api:** update via SDK Studio ([#135](https://github.com/terminaldotshop/terminal-sdk-go/issues/135)) ([ee83fe2](https://github.com/terminaldotshop/terminal-sdk-go/commit/ee83fe285738375e02e4c906b2af5cec1a7c951e))

## 0.1.0-alpha.34 (2024-11-17)

Full Changelog: [v0.1.0-alpha.33...v0.1.0-alpha.34](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.33...v0.1.0-alpha.34)

### Features

* **api:** update via SDK Studio ([#132](https://github.com/terminaldotshop/terminal-sdk-go/issues/132)) ([555fb76](https://github.com/terminaldotshop/terminal-sdk-go/commit/555fb76c3fed3ff79b5450635fd0b11c71b5ab54))

## 0.1.0-alpha.33 (2024-11-17)

Full Changelog: [v0.1.0-alpha.32...v0.1.0-alpha.33](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.32...v0.1.0-alpha.33)

### Features

* **api:** api update ([#129](https://github.com/terminaldotshop/terminal-sdk-go/issues/129)) ([434dd55](https://github.com/terminaldotshop/terminal-sdk-go/commit/434dd554b5cac5fe4d3e5fced7a32c226262441b))
* **api:** api update ([#131](https://github.com/terminaldotshop/terminal-sdk-go/issues/131)) ([3bcbd62](https://github.com/terminaldotshop/terminal-sdk-go/commit/3bcbd627083da9e8c387e8b22a807b9275051cf7))

## 0.1.0-alpha.32 (2024-11-17)

Full Changelog: [v0.1.0-alpha.31...v0.1.0-alpha.32](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.31...v0.1.0-alpha.32)

### Features

* **api:** update via SDK Studio ([#127](https://github.com/terminaldotshop/terminal-sdk-go/issues/127)) ([3626993](https://github.com/terminaldotshop/terminal-sdk-go/commit/36269938621cf514a7fdea516c82d801fcb95468))

## 0.1.0-alpha.31 (2024-11-17)

Full Changelog: [v0.1.0-alpha.30...v0.1.0-alpha.31](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.30...v0.1.0-alpha.31)

### Features

* **api:** update via SDK Studio ([#125](https://github.com/terminaldotshop/terminal-sdk-go/issues/125)) ([9a5c3b8](https://github.com/terminaldotshop/terminal-sdk-go/commit/9a5c3b8e57123cea3b6d169eabb2858bd320a3ca))

## 0.1.0-alpha.30 (2024-11-17)

Full Changelog: [v0.1.0-alpha.29...v0.1.0-alpha.30](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.29...v0.1.0-alpha.30)

### Features

* **api:** api update ([#120](https://github.com/terminaldotshop/terminal-sdk-go/issues/120)) ([43ec053](https://github.com/terminaldotshop/terminal-sdk-go/commit/43ec053a620942fc9889b83e08585cbee0daffd6))

## 0.1.0-alpha.29 (2024-11-17)

Full Changelog: [v0.1.0-alpha.28...v0.1.0-alpha.29](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.28...v0.1.0-alpha.29)

### Features

* **api:** api update ([#117](https://github.com/terminaldotshop/terminal-sdk-go/issues/117)) ([995ef57](https://github.com/terminaldotshop/terminal-sdk-go/commit/995ef577923c454d5f060cd831843e8a66b576d4))
* **api:** update via SDK Studio ([#119](https://github.com/terminaldotshop/terminal-sdk-go/issues/119)) ([74a8d8b](https://github.com/terminaldotshop/terminal-sdk-go/commit/74a8d8b3235bcf7f8d18f356173c1e80d29aa56a))

## 0.1.0-alpha.28 (2024-11-17)

Full Changelog: [v0.1.0-alpha.27...v0.1.0-alpha.28](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.27...v0.1.0-alpha.28)

### Features

* **api:** update via SDK Studio ([#115](https://github.com/terminaldotshop/terminal-sdk-go/issues/115)) ([a0dc5e4](https://github.com/terminaldotshop/terminal-sdk-go/commit/a0dc5e45313ce1e0e6bc6447b0c1b61d00f35a8e))

## 0.1.0-alpha.27 (2024-11-13)

Full Changelog: [v0.1.0-alpha.26...v0.1.0-alpha.27](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.26...v0.1.0-alpha.27)

### Features

* **api:** api update ([#107](https://github.com/terminaldotshop/terminal-sdk-go/issues/107)) ([b880524](https://github.com/terminaldotshop/terminal-sdk-go/commit/b880524f5a58af739cc187c881438ce4ac6bd583))
* **api:** api update ([#109](https://github.com/terminaldotshop/terminal-sdk-go/issues/109)) ([19e1bd2](https://github.com/terminaldotshop/terminal-sdk-go/commit/19e1bd2fcb00949944c0d1ba94b578eec0bba4f9))
* **api:** update via SDK Studio ([#113](https://github.com/terminaldotshop/terminal-sdk-go/issues/113)) ([de6de7b](https://github.com/terminaldotshop/terminal-sdk-go/commit/de6de7b7c677836a6cee2d8c568e37635e0a889b))


### Chores

* rebuild project due to codegen change ([#110](https://github.com/terminaldotshop/terminal-sdk-go/issues/110)) ([185979c](https://github.com/terminaldotshop/terminal-sdk-go/commit/185979c9ce80c896cc404b762c168e8996c65fee))
* rebuild project due to codegen change ([#111](https://github.com/terminaldotshop/terminal-sdk-go/issues/111)) ([8ae9c3d](https://github.com/terminaldotshop/terminal-sdk-go/commit/8ae9c3d7bef59ee8989067b7859f77a6a5098223))
* rebuild project due to codegen change ([#112](https://github.com/terminaldotshop/terminal-sdk-go/issues/112)) ([4fb82b3](https://github.com/terminaldotshop/terminal-sdk-go/commit/4fb82b3477cfc5c3f4ffb4942d3361e93b9a99a6))

## 0.1.0-alpha.26 (2024-10-21)

Full Changelog: [v0.1.0-alpha.25...v0.1.0-alpha.26](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.25...v0.1.0-alpha.26)

### Features

* **api:** api update ([#104](https://github.com/terminaldotshop/terminal-sdk-go/issues/104)) ([a4253d8](https://github.com/terminaldotshop/terminal-sdk-go/commit/a4253d8939339bec1ea7497b269e81abee112c4c))
* **api:** api update ([#105](https://github.com/terminaldotshop/terminal-sdk-go/issues/105)) ([a55c65a](https://github.com/terminaldotshop/terminal-sdk-go/commit/a55c65a5612d0fbedafae9c74bcca10df0abf858))
* **api:** OpenAPI spec update via Stainless API ([#100](https://github.com/terminaldotshop/terminal-sdk-go/issues/100)) ([e57c88a](https://github.com/terminaldotshop/terminal-sdk-go/commit/e57c88afd119dcd7c13981429d50abb5a855c1cd))
* **api:** OpenAPI spec update via Stainless API ([#101](https://github.com/terminaldotshop/terminal-sdk-go/issues/101)) ([660a7b0](https://github.com/terminaldotshop/terminal-sdk-go/commit/660a7b024094505dfec95fd13e1e053526d7900c))
* **api:** OpenAPI spec update via Stainless API ([#102](https://github.com/terminaldotshop/terminal-sdk-go/issues/102)) ([675c86a](https://github.com/terminaldotshop/terminal-sdk-go/commit/675c86a5ea86af7819e4e72a07c75b423a8d01ff))
* **api:** OpenAPI spec update via Stainless API ([#103](https://github.com/terminaldotshop/terminal-sdk-go/issues/103)) ([8c2e078](https://github.com/terminaldotshop/terminal-sdk-go/commit/8c2e0789594cfa52b1acb39b22d46d35e4c04541))
* **api:** OpenAPI spec update via Stainless API ([#97](https://github.com/terminaldotshop/terminal-sdk-go/issues/97)) ([e83c8d2](https://github.com/terminaldotshop/terminal-sdk-go/commit/e83c8d29558090a0bed9717c6002344ced11a0e6))
* **api:** OpenAPI spec update via Stainless API ([#99](https://github.com/terminaldotshop/terminal-sdk-go/issues/99)) ([ed00d94](https://github.com/terminaldotshop/terminal-sdk-go/commit/ed00d94d382652352b7e962fe716d9112d9f2a0d))

## 0.1.0-alpha.25 (2024-07-03)

Full Changelog: [v0.1.0-alpha.24...v0.1.0-alpha.25](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.24...v0.1.0-alpha.25)

### Features

* **api:** OpenAPI spec update via Stainless API ([#94](https://github.com/terminaldotshop/terminal-sdk-go/issues/94)) ([cf914c4](https://github.com/terminaldotshop/terminal-sdk-go/commit/cf914c4649d23eb66ae0b7ab10a0370c247c247c))
* **api:** update via SDK Studio ([#96](https://github.com/terminaldotshop/terminal-sdk-go/issues/96)) ([b1c200b](https://github.com/terminaldotshop/terminal-sdk-go/commit/b1c200b788cfca79c8884c7954ace63c0c6ad563))

## 0.1.0-alpha.24 (2024-07-03)

Full Changelog: [v0.1.0-alpha.23...v0.1.0-alpha.24](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.23...v0.1.0-alpha.24)

### Features

* **api:** OpenAPI spec update via Stainless API ([#88](https://github.com/terminaldotshop/terminal-sdk-go/issues/88)) ([42347a2](https://github.com/terminaldotshop/terminal-sdk-go/commit/42347a20f3c43551d444adda0fce4af66e34c8b6))
* **api:** OpenAPI spec update via Stainless API ([#90](https://github.com/terminaldotshop/terminal-sdk-go/issues/90)) ([07369c5](https://github.com/terminaldotshop/terminal-sdk-go/commit/07369c518d67df456cc78344023d04865ccc748c))
* **api:** OpenAPI spec update via Stainless API ([#92](https://github.com/terminaldotshop/terminal-sdk-go/issues/92)) ([a79f62d](https://github.com/terminaldotshop/terminal-sdk-go/commit/a79f62d9e987804501c20339a4fde9362622a142))
* **api:** update via SDK Studio ([#91](https://github.com/terminaldotshop/terminal-sdk-go/issues/91)) ([2dcae76](https://github.com/terminaldotshop/terminal-sdk-go/commit/2dcae765dbbea666615f72fb4e609e7af917f206))
* **api:** update via SDK Studio ([#93](https://github.com/terminaldotshop/terminal-sdk-go/issues/93)) ([8038917](https://github.com/terminaldotshop/terminal-sdk-go/commit/803891784c0df43a80947157e1962b20433d0f36))

## 0.1.0-alpha.23 (2024-06-29)

Full Changelog: [v0.1.0-alpha.22...v0.1.0-alpha.23](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.22...v0.1.0-alpha.23)

### Features

* **api:** OpenAPI spec update via Stainless API ([#83](https://github.com/terminaldotshop/terminal-sdk-go/issues/83)) ([5543d8f](https://github.com/terminaldotshop/terminal-sdk-go/commit/5543d8f6655df201fbb321ee119282457936b754))
* **api:** OpenAPI spec update via Stainless API ([#85](https://github.com/terminaldotshop/terminal-sdk-go/issues/85)) ([eb616d2](https://github.com/terminaldotshop/terminal-sdk-go/commit/eb616d298084f3571dee5a084089890ecbb64961))
* **api:** update via SDK Studio ([#86](https://github.com/terminaldotshop/terminal-sdk-go/issues/86)) ([c918a31](https://github.com/terminaldotshop/terminal-sdk-go/commit/c918a317ffa7e2e00c6095d8082ec30a830218cd))

## 0.1.0-alpha.22 (2024-06-28)

Full Changelog: [v0.1.0-alpha.21...v0.1.0-alpha.22](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.21...v0.1.0-alpha.22)

### Features

* **api:** OpenAPI spec update via Stainless API ([#79](https://github.com/terminaldotshop/terminal-sdk-go/issues/79)) ([6310a43](https://github.com/terminaldotshop/terminal-sdk-go/commit/6310a430a37c56fab9368e1734e6538fa53cb02d))
* **api:** update via SDK Studio ([#81](https://github.com/terminaldotshop/terminal-sdk-go/issues/81)) ([40f1948](https://github.com/terminaldotshop/terminal-sdk-go/commit/40f19487b6739ba939f6a71460812109f24d8316))

## 0.1.0-alpha.21 (2024-06-28)

Full Changelog: [v0.1.0-alpha.20...v0.1.0-alpha.21](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.20...v0.1.0-alpha.21)

### Features

* **api:** OpenAPI spec update via Stainless API ([#75](https://github.com/terminaldotshop/terminal-sdk-go/issues/75)) ([9435c0d](https://github.com/terminaldotshop/terminal-sdk-go/commit/9435c0da5114c0dc9e627f9b54d6297d7d1fbef1))
* **api:** update via SDK Studio ([#77](https://github.com/terminaldotshop/terminal-sdk-go/issues/77)) ([bd212a3](https://github.com/terminaldotshop/terminal-sdk-go/commit/bd212a378b30bf596e01693515fba3ad33e2f87c))

## 0.1.0-alpha.20 (2024-06-28)

Full Changelog: [v0.1.0-alpha.19...v0.1.0-alpha.20](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.19...v0.1.0-alpha.20)

### Features

* **api:** OpenAPI spec update via Stainless API ([#71](https://github.com/terminaldotshop/terminal-sdk-go/issues/71)) ([696f8e3](https://github.com/terminaldotshop/terminal-sdk-go/commit/696f8e37e55712a7825513b861555d56dae1c17e))
* **api:** update via SDK Studio ([#73](https://github.com/terminaldotshop/terminal-sdk-go/issues/73)) ([8d435bc](https://github.com/terminaldotshop/terminal-sdk-go/commit/8d435bcc3bb97f7a8b559445a58c99ac077dda10))

## 0.1.0-alpha.19 (2024-06-27)

Full Changelog: [v0.1.0-alpha.18...v0.1.0-alpha.19](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.18...v0.1.0-alpha.19)

### Features

* **api:** update via SDK Studio ([#68](https://github.com/terminaldotshop/terminal-sdk-go/issues/68)) ([2ed2bb9](https://github.com/terminaldotshop/terminal-sdk-go/commit/2ed2bb9345883a1a12d2bc1f24a985fee7bf6781))

## 0.1.0-alpha.18 (2024-06-26)

Full Changelog: [v0.1.0-alpha.17...v0.1.0-alpha.18](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.17...v0.1.0-alpha.18)

### Features

* **api:** OpenAPI spec update via Stainless API ([#63](https://github.com/terminaldotshop/terminal-sdk-go/issues/63)) ([d97ec7d](https://github.com/terminaldotshop/terminal-sdk-go/commit/d97ec7d9b5e48d1a4f4830f2f443811ac444a1fe))
* **api:** OpenAPI spec update via Stainless API ([#65](https://github.com/terminaldotshop/terminal-sdk-go/issues/65)) ([8baba16](https://github.com/terminaldotshop/terminal-sdk-go/commit/8baba162d6a6f597b7f3d66822b95733072da372))
* **api:** OpenAPI spec update via Stainless API ([#66](https://github.com/terminaldotshop/terminal-sdk-go/issues/66)) ([98d9f5b](https://github.com/terminaldotshop/terminal-sdk-go/commit/98d9f5bf3870aeda22c451fac425dc053f2cd421))

## 0.1.0-alpha.17 (2024-06-25)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* **api:** update via SDK Studio ([#60](https://github.com/terminaldotshop/terminal-sdk-go/issues/60)) ([141d088](https://github.com/terminaldotshop/terminal-sdk-go/commit/141d088a67a11cbab50c7dd42d02d5f4401663b0))

## 0.1.0-alpha.16 (2024-06-25)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Features

* **api:** OpenAPI spec update via Stainless API ([#56](https://github.com/terminaldotshop/terminal-sdk-go/issues/56)) ([682ff23](https://github.com/terminaldotshop/terminal-sdk-go/commit/682ff234fe9937efae27acb25d215d9644b5eaeb))
* **api:** update via SDK Studio ([#58](https://github.com/terminaldotshop/terminal-sdk-go/issues/58)) ([de5fc02](https://github.com/terminaldotshop/terminal-sdk-go/commit/de5fc0218a4f3081302dad8c1e7b83af6854f223))

## 0.1.0-alpha.15 (2024-06-25)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **api:** OpenAPI spec update via Stainless API ([#52](https://github.com/terminaldotshop/terminal-sdk-go/issues/52)) ([4e5d9d6](https://github.com/terminaldotshop/terminal-sdk-go/commit/4e5d9d6914fdfe7e6ba4bcd7d258f9eee0d729f8))
* **api:** update via SDK Studio ([#54](https://github.com/terminaldotshop/terminal-sdk-go/issues/54)) ([443dd0e](https://github.com/terminaldotshop/terminal-sdk-go/commit/443dd0e1f5bfe15be9b8417b9dab73e8811c2a59))

## 0.1.0-alpha.14 (2024-06-24)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* **api:** update via SDK Studio ([#49](https://github.com/terminaldotshop/terminal-sdk-go/issues/49)) ([0b6f8ba](https://github.com/terminaldotshop/terminal-sdk-go/commit/0b6f8bafb063a493dbed4dc73ddc17560293e580))

## 0.1.0-alpha.13 (2024-06-24)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Features

* **api:** OpenAPI spec update via Stainless API ([#44](https://github.com/terminaldotshop/terminal-sdk-go/issues/44)) ([3c7b912](https://github.com/terminaldotshop/terminal-sdk-go/commit/3c7b9128da078780bd59dffc89987082deef7c5a))
* **api:** update via SDK Studio ([#46](https://github.com/terminaldotshop/terminal-sdk-go/issues/46)) ([8a34095](https://github.com/terminaldotshop/terminal-sdk-go/commit/8a34095a0e7d34c9fa0e3a39743c8f999ce30573))

## 0.1.0-alpha.12 (2024-06-20)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** update via SDK Studio ([#41](https://github.com/terminaldotshop/terminal-sdk-go/issues/41)) ([bd8127e](https://github.com/terminaldotshop/terminal-sdk-go/commit/bd8127e1af062b98c71aa9709e336e49e8b8b677))

## 0.1.0-alpha.11 (2024-06-20)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Features

* **api:** update via SDK Studio ([#38](https://github.com/terminaldotshop/terminal-sdk-go/issues/38)) ([7ad3a72](https://github.com/terminaldotshop/terminal-sdk-go/commit/7ad3a7223234fdc1d1fa1f481e8679e784901723))

## 0.1.0-alpha.10 (2024-06-20)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Features

* **api:** OpenAPI spec update via Stainless API ([#35](https://github.com/terminaldotshop/terminal-sdk-go/issues/35)) ([b0df3b7](https://github.com/terminaldotshop/terminal-sdk-go/commit/b0df3b7afaa9cd58274c4a0d64fbdf9d91b6349d))
* **api:** OpenAPI spec update via Stainless API ([#37](https://github.com/terminaldotshop/terminal-sdk-go/issues/37)) ([027c7be](https://github.com/terminaldotshop/terminal-sdk-go/commit/027c7be77b5140d660f09e14c29e5c0f251ad5f3))

## 0.1.0-alpha.9 (2024-06-19)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Features

* **api:** update via SDK Studio ([#32](https://github.com/terminaldotshop/terminal-sdk-go/issues/32)) ([a860dd7](https://github.com/terminaldotshop/terminal-sdk-go/commit/a860dd78eca0b56fa356258e58bb2a580a882961))

## 0.1.0-alpha.8 (2024-06-19)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** OpenAPI spec update via Stainless API ([#24](https://github.com/terminaldotshop/terminal-sdk-go/issues/24)) ([34989f4](https://github.com/terminaldotshop/terminal-sdk-go/commit/34989f42e35c7eccbe4d5206aebb3192e0188873))
* **api:** OpenAPI spec update via Stainless API ([#27](https://github.com/terminaldotshop/terminal-sdk-go/issues/27)) ([571b4c3](https://github.com/terminaldotshop/terminal-sdk-go/commit/571b4c37d3022558d9b34276f0ae54973e4676ee))
* **api:** OpenAPI spec update via Stainless API ([#30](https://github.com/terminaldotshop/terminal-sdk-go/issues/30)) ([699af11](https://github.com/terminaldotshop/terminal-sdk-go/commit/699af11583c87861c907c2c1ae42f6ffe1152a71))
* **api:** update via SDK Studio ([#26](https://github.com/terminaldotshop/terminal-sdk-go/issues/26)) ([82e1ced](https://github.com/terminaldotshop/terminal-sdk-go/commit/82e1cedeadbdd486b6649544e202a39863e9aa7f))
* **api:** update via SDK Studio ([#28](https://github.com/terminaldotshop/terminal-sdk-go/issues/28)) ([ee42232](https://github.com/terminaldotshop/terminal-sdk-go/commit/ee42232e1e3a0dc09ffba2649abfe0e1c295d579))

## 0.1.0-alpha.7 (2024-05-17)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** OpenAPI spec update via Stainless API ([#19](https://github.com/terminaldotshop/terminal-sdk-go/issues/19)) ([9e024d8](https://github.com/terminaldotshop/terminal-sdk-go/commit/9e024d8d485186dac83a0cfad26f3f90aca826d3))
* **api:** OpenAPI spec update via Stainless API ([#21](https://github.com/terminaldotshop/terminal-sdk-go/issues/21)) ([468fbb4](https://github.com/terminaldotshop/terminal-sdk-go/commit/468fbb4859d1e96570154800c4090c5d98f05cd3))
* **api:** OpenAPI spec update via Stainless API ([#22](https://github.com/terminaldotshop/terminal-sdk-go/issues/22)) ([82ea930](https://github.com/terminaldotshop/terminal-sdk-go/commit/82ea930255f72fdf15592730f1e0ea80a5201e1c))

## 0.1.0-alpha.6 (2024-05-14)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** OpenAPI spec update via Stainless API ([#15](https://github.com/terminaldotshop/terminal-sdk-go/issues/15)) ([9486960](https://github.com/terminaldotshop/terminal-sdk-go/commit/94869603fd510de3737f9ebd4d5554346f7ec5ca))
* **api:** OpenAPI spec update via Stainless API ([#17](https://github.com/terminaldotshop/terminal-sdk-go/issues/17)) ([bf39bee](https://github.com/terminaldotshop/terminal-sdk-go/commit/bf39beecdeb8030465e9778ce19755f2713b7b60))

## 0.1.0-alpha.5 (2024-05-14)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** update via SDK Studio ([#12](https://github.com/terminaldotshop/terminal-sdk-go/issues/12)) ([d579714](https://github.com/terminaldotshop/terminal-sdk-go/commit/d579714aa731435f37705fa733849ce44f1495d2))

## 0.1.0-alpha.4 (2024-05-14)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** OpenAPI spec update via Stainless API ([#9](https://github.com/terminaldotshop/terminal-sdk-go/issues/9)) ([415091a](https://github.com/terminaldotshop/terminal-sdk-go/commit/415091a348a82b66cea976a7b55f745f398efd55))

## 0.1.0-alpha.3 (2024-05-14)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** update via SDK Studio ([#7](https://github.com/terminaldotshop/terminal-sdk-go/issues/7)) ([ab172f4](https://github.com/terminaldotshop/terminal-sdk-go/commit/ab172f4711d98d02ea12313f7c3ec695ed4acfa6))

## 0.1.0-alpha.2 (2024-05-14)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** OpenAPI spec update via Stainless API ([#4](https://github.com/terminaldotshop/terminal-sdk-go/issues/4)) ([678e09a](https://github.com/terminaldotshop/terminal-sdk-go/commit/678e09ae72cdab6465f4bcc6ca38598c8f0938e2))

## 0.1.0-alpha.1 (2024-05-14)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/terminaldotshop/terminal-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** OpenAPI spec update via Stainless API ([7b5fdbd](https://github.com/terminaldotshop/terminal-sdk-go/commit/7b5fdbd35e1d5aaab4ce706cfdfd1b483331f101))
* **api:** OpenAPI spec update via Stainless API ([8fd196f](https://github.com/terminaldotshop/terminal-sdk-go/commit/8fd196f1aadd962aff4e88dda762309c46cc004d))
* **api:** update via SDK Studio ([f452dbd](https://github.com/terminaldotshop/terminal-sdk-go/commit/f452dbdcbf852fbe98213129527f66b7d42228ec))
* **api:** update via SDK Studio ([53be291](https://github.com/terminaldotshop/terminal-sdk-go/commit/53be291b06735b37bcf03e68f257440a8077bc62))
* **api:** update via SDK Studio ([d489aa8](https://github.com/terminaldotshop/terminal-sdk-go/commit/d489aa8f73bc794c2189d50d7d25aa921ccfd257))


### Chores

* configure new SDK language ([30f6e4e](https://github.com/terminaldotshop/terminal-sdk-go/commit/30f6e4e1e3862939c3d89ae456b39874b20eb603))
* go live ([#1](https://github.com/terminaldotshop/terminal-sdk-go/issues/1)) ([c11d6b2](https://github.com/terminaldotshop/terminal-sdk-go/commit/c11d6b2d3799c5db544e69b9dda8a376ec2947b1))
