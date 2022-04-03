# Autodoc helper

## Description
Small https://www.autodoc.ru/ price checker with telegram notifications

## Build

- `make build`
- Run or add to crontab list

## Usage:

- app version `./bin/price_checker --config ./config.yml --version`
- add to checklist `./bin/price_checker --config ./config.yml --add --partnumber <partnumber> --manufacter <manufacter_id>`
- check by id `./bin/price_checker --config ./config.yml --check --partnumber <partnumber> --manufacter <manufacter_id>`
- check all in list `./bin/price_checker --config ./config.yml --check-all`
- cleanup db `./bin/price_checker --config ./config.yml --cleanup`


## Changelog

- 1.0: first release
- 1.1: add cleanup, change flags