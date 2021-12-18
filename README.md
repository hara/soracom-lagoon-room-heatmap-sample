# Sample Heat Map Renderer for SORACOM Lagoon Dynamic Image Visualization

This is a sample heatmap renderer for Soracom Dynamic Image Visualization of SORACOM Lagoon.
You send `GET /render/20-50`, renderer renders heat map with 20°C (temperature) and 50% (humidity).

## Prerequisite

* AWS Account
* AWS SAM CLI

## How to deployment

Configure AWS SAM CLI and

```
$ make build && make deploy
```

## How to use

Configure your Soracom Dynamic Image Visualization.

* Background:
  * prefix: `https://YOURAPI.execute-api.REGION.amazonaws.com/render/`
  * suffix: none
* Variable 0: temperature
* Variable 1: humidity
