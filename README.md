# Autodoc helper

## Description
Small https://www.autodoc.ru/ price checker with telegram notifications

## Build

- `make build`
- Run or add to crontab list

## Usage:

- Add to checklist `./bin/price_checker --config ./config.yml --add <partnumber> --manufacter <manufacter_id>`

- check all `./bin/price_checker --config ./config.yml --check-all`

- check `./bin/price_checker --config ./config.yml --check <partnumber>`

- app version `./bin/price_checker --config ./config.yml --version`
- cleanup db `./bin/price_checker --config ./config.yml --cleanup`


## Changelog

- 1.0: first release