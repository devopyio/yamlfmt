# yamlfmt

Formats yaml files.

## Install

go get -u github.com/devopyio/yamlfmt

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
