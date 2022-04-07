# outline

10am
- start with a few intros and explainations 15m


| Time  | Content                         | Duration | Materials |
|-------|---------------------------------|----------|-----------|
| 10    | Intro about me, lecture outline | 20m      | Slides    |
| 10.20 | Explain generic approach        | 20m      | Slides    |
|       |                                 |          |           |
|       |                                 |          |           |


# Lecture outline

points to cover
- about me
- points to cover
  - intro to gitops
  - traditional ways of maintaining infra
    - ways to deploy
    - ssh
    - pull based (redis)
    - pros and cons
    - security setup - (restrict ssh access to servers mostly, or to cloud accounts)
    - click based uploads
    
  - gitops
    - push or pull based
    - git repo management
    - access controls
    - security setup (restrict access to gitops repos)
    - infrastructure management as code (IaC) - terraform 
  - Live demo
    - infra
      - with and without gitops (deploy the pets store application, in DO Apps)
    - updates to the petstore application
      - how to deploy to multiple environments
      - introduce 12factor.net (esp build once, dpeloy many times rule )
      - envvars structuring



Demo

code
2 apis 
-> get my pets - returns (random values like sadlkjf,akjsdf in stage) (proper animals in prod) - based on env variable - animals
-> healthcheck - shows the mode, commit sha, time

we add a new update 
-> 3rd api, says Hello

- traditional way
  clone the demo app
  Create a DO app by hand 
  deploy to 2 env - staging and prod, by hand
  - push and build the docker push to the registry
  - goto the GUI deploy staging (keep one ENV var: MODE - 2 values prod and stage) by selecting the version - env MODE = stage 
  - goto the GUI deploy prod 
  
- GitOps way
  - clone demo app
  - create the app via terraform, set the env vars
  - build the app and push to a container registry
  - 
  - deploy to stage via push to this repo
  - deploy to prod via push to this repo


DO App creation list 
container registry
App
env vars 
update the Build files, with container registry name
add github actions file


Flowchart


points to segue

outline is made so that aud understands the current system on which the new approach is based.



# Db commands

to connect 
psql <connection string>
run a script 
\i <path to script>
show dbs
\l

show tables
\dt

## connect to DO dbs
https://docs.digitalocean.com/products/databases/postgresql/how-to/connect/




