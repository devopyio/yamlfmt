# yamlfmt

![Build Status](https://travis-ci.com/devopyio/yamlfmt.svg?branch=master)
[![yamlfmt](https://snapcraft.io/yamlfmt/badge.svg)](https://snapcraft.io/yamlfmt)

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/yamlfmt)

Formats yaml files.

## Install

### Linux:

```
sudo snap install yamlfmt
```

### Brew 

```
brew install devopyio/yamlfmt/yamlfmt
```

### Windows

```
scoop bucket add yamlfmt https://github.com/devopyio/scoop-yamlfmt
scoop install yamlfmt
```
### Using go get

```
go get -u github.com/devopyio/yamlfmt
```

## Usage

```
cat example.yaml | yamlfmt
```

Or you can simply execute:

```
yamlfmt -filename example.yaml
```

Example:

```
"groups": 
- "name": "etcd"
  "rules": 
  - "alert": "EtcdInsufficientMembers"
    "annotations": 
      "message": "Etcd cluster \"{{ $labels.job }}\": insufficient members ({{ $value }})."
    "expr": |
      count(up{job="etcd"} == 0) by (job) > (count(up{job="etcd"}) by (job) / 2 - 1)
    "for": "3m"
    "labels": 
      "severity": "critical"

```

Becomes:

```
groups:
- name: etcd
  rules:
  - alert: EtcdInsufficientMembers
    annotations:
      message: 'Etcd cluster "{{ $labels.job }}": insufficient members ({{ $value
        }}).'
    expr: |
      count(up{job="etcd"} == 0) by (job) > (count(up{job="etcd"}) by (job) / 2 - 1)
    for: 3m
    labels:
      severity: critical
```
