# hello-actions
A simple repository to test the behaviours of Github Actions.

## Github Workflows  and atributes 

Github Workflows support YAML based files to setup and configurare workfllowsz to work with Github Actions.

```
name: Hello
on: [push]
jobs:
  build:
    rans-on: ubuntu-latest
    steps:
      - name: checkout
        uses: action/checkout@v4
      - name: Executing code
        run: echo Hello worldd!!
```

### Glosary:
- on -> When is going to be executed the workflow (OnPush action or workflow_dispatch).
- jobs:  -> Every workflow must have at least 1 job.
  -  build -> Every job has multiple steps. In this case the "build steps"
    -  rans-on: -> On which plataform will run (Windows-latest, ubuntu-latest or MacOs-lat es).
    -  steps: -> Actions that will done the job. 
      - name:  -> --Optional to describe the step. 
      - uses: -> Define which github action to use (usefulkk to declare, upload, download or set up go environment).
       - run  -> execute Terminal code (base or PowerShell).

## Examples 

Based on the linkedin's  course **"Github Actions: Event-Driven Automatation"**, i write 3 sample configuration files. 
- [Github workflow explaining basic multi-job configuration.](.github/workflows/multi-job.yml)
- [Github Workflow explaining basic on Creating and deleting artifact.](.github/workflows/artifacts.yml)
- [Github workflow Sample configuration pipeline for TEST, BUILD AND BUILT-TEST. ](.github/workflows/pipeline.yml)
