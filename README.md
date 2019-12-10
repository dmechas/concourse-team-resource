# Concourse Team Resource

Get and set concourse teams from concourse inpired by [concourse-pipeline-resource](https://github.com/concourse/concourse-pipeline-resource)

## Installing

Use this resource by adding the following to the `resource_types` section of a pipeline config:

```yaml
---
resource_types:
- name: concourse-team
  type: docker-image
  source:
    repository: dmechas/concourse-team-resource
```

See [concourse docs](https://concourse-ci.org/resource-types.html) for more details on adding `resource_types` to a pipeline config.
