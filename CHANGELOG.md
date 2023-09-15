# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to the following versioning pattern:

Given a version number MAJOR.MINOR.PATCH, increment:

- MAJOR version when the **API** version is incremented. This may include backwards incompatible changes;
- MINOR version when **breaking changes** are introduced OR **new functionalities** are added in a backwards compatible manner;
- PATCH version when backwards compatible bug **fixes** are implemented.


## [Unreleased]
### Removed 
- AccountCreated, Created and Owned attributes to DictKey resource
- AccountNumber and BranchCode attributes to PaymentPreview resource

### Changed
- AccountNumber and BranchCode docstring attributes to DictKey resource

## [0.3.1] - 2023-09-14
### Changed
- core version

## [0.3.0] - 2023-05-31
### Added
- CorporateBalance resource
- CorporateCard resource
- CorporateHolder resource
- CorporateInvoice resource
- CorporatePurchase resource
- CorporateRule resource
- CorporateTransaction resource
- CorporateWithdrawal resource
- CardMethod sub-resource
- MerchantCategory sub-resource
- MerchantCountry sub-resource
- rules attribute to Invoice resource
- Invoice.Rule sub-resource

## [0.2.0] - 2023-03-22
### Changed
- metadata Transfer attribute from struct to map

## [0.1.0] - 2023-03-17
### Added
- metadata attribute to Transfer resource

## [0.0.1] - 2023-01-26
### Added
- Full Stark Bank API v2 compatibility
