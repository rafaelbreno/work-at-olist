## My Local Specifications
- OS: Arch Linux 
- Shell: zsh
- Text Editor: Neovim([dotfiles](https://github.com/rafaelbreno/vimfiles))

------

## Project Specificiations

### Technologies
- WebServer: NGINX
- Language: Go 1.16
- Database: Postgres 13.1

### Structure
#### _.docker/_
- Here will be configured all Docker local deployment files
- _go/_ - Project Language
- _postgres/_ - Project Database
- _nginx/_ - Project Webserver

------

## Todo:
- [ ] Finish API
- [ ] Write Tests
- [ ] Deploy on Heroku
- [ ] Apply
- [ ] Add How to Deploy locally
    - [ ] Without Docker
    - [ ] With Docker
- [ ] Implement CI
- [ ] Write project's documentation
    - Deployment
    - API
- [ ] Implement CLI commands
    - Add CSV files
    - `go run add_csg.go --file=foo_bar.csv`

------

### Twelve Factors
1. Codebase
    - Git as VSC and a PR system(not using Issues because this is a Fork not the real __repo__ itself)
    - Using CircleCI to run the Pipeline
2. Dependencies
    - Go itself has the _Go Modules_ to manage their dependencies
3. Config
    - The project config will be stored in _'.env'_ file
    - Because I'll be deploying it, it'll have: _'.env.example'_ _'.env.local.docker'_ _'.env.local'_ _'.env.prod'_ 
    - Just to ease the work
4. Backing Services
    - Database
        - Local: Postgres Container
        - Prod: Heroku Postgres
5. Build, release, run
    - Define branchs with _release/_ prefix and it's version
6. Processes
    - Standalone script    
7. Port Binding
    - Each service has it's own port, Go app, Postgres
8. Concurrency
    - Go has native Concurrency, so we can define workers on Nginx
9. Disposability
    - ???
10. Dev/prod parity
    - The only change between dev and prod is the environment file
    - To run locally it's used the Docker Compose, to the similiarity between the prod environment and the developer environment is similar
11. Logs
    - ???
12. Admin processes
    - ???
