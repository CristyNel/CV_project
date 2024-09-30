# CV Management System API

## Overview : This project provides a platform for managing user data and generating custom CV templates

![Cv_examples](https://venngage-wordpress.s3.amazonaws.com/uploads/2021/11/section-3-resume-banner-1-1.png)

## Screenshots: development

![SCreenshot: 18/09](./bff/static/img/screenshots/Screenshot_18-09-24.png)

```sh
/CV_project
├── api                                 # go api         - backend
│   ├── cmd                             # command directory
│   │   ├── main.go                     # main entry point
│   │   └── main_test.go                # tests for main
│   ├── e2etests.yml                    # end-to-end tests configuration
│   ├── go.mod                          # go modules     - dependency management
│   ├── go.sum                          # checksum       - dependencies
│   ├── handlers                        # handlers for different routes
│   │   ├── home.go                     # home handler
│   │   ├── templates.go                # templates handler
│   │   └── users.go                    # users handler
│   ├── internal                        # internal packages
│   │   ├── app                         # application logic
│   │   │   └── app.go                  # application entry point
│   │   ├── database                    # database logic
│   │   │   └── db.go                   # database connection
│   │   └── utils                       # utility functions
│   │       └── utils.go                # utility functions implementation
│   ├── mock                            # mock handlers for testing
│   │   └── mock_handlers.go            # mock handlers
│   ├── models                          # data models
│   │   └── user.go                     # user model
│   ├── routes                          # route definitions
│   │   └── router.go                   # router setup
│   └── tests                           # tests
│       ├── coverage.html               # test coverage report
│       ├── coverage.out                # test coverage output
│       ├── home_test.go                # tests for home handler
│       ├── mock_handlers_test.go       # tests for mock handlers
│       ├── users_test.go               # tests for users handler
│       └── venom.log                   # venom test logs
├── bff                                 # flask-based    - frontend
│   ├── app                             # flask application
│   │   ├── app.py                      # main flask application file
│   │   └── __init__.py                 # initialization file
│   ├── requirements.txt                # python dependencies
│   ├── static                          # static files
│   │   ├── img                         # images
│   │   │   └── screenshots             # screenshots
│   │   ├── js                          # javascript files
│   │   │   └── users.js                # custom javascript - user-related functionality
│   │   └── styles                      # css files
│   │       └── users.css               # custom css - user-related pages
│   └── templates                       # html templates
│       ├── forms                       # form templates
│       │   ├── edit_form.html          # html form  - editing user data
│       │   ├── loginform.html          # login form
│       │   ├── post_form.html          # form       - posting new content
│       │   └── signupform.html         # signup form
│       └── view                        # view templates
│           ├── favicon.ico             # favicon
│           ├── greet.html              # greeting page
│           ├── home.html               # home page
│           ├── populate_template.html  # template    - populating cvs
│           ├── template1.html          # cv template 1
│           ├── template2.html          # cv template 2
│           └── template3.html          # cv template 3
├── docker-compose.yml                  # docker compose configuration
├── Dockerfile.api                      # dockerfile for the api
├── Dockerfile.bff                      # dockerfile for the bff
├── .env                                # environment variables
├── .github                             # github workflows
│   └── workflows                       # github actions workflows
│       ├── unit-tests.yml              # unit tests workflow
│       └── venom-tests.yml             # venom tests workflow
├── .gitignore                          # git ignore  - version control
├── Makefile                            # makefile for build automation
├── README.md                           # project documentation
└── sql                                 # sql files   - database schema
    ├── schemadump.sql                  # schema creation and sample data
    └── schema.sql                      # schema creation only
```

## Components

```plaintext
    Backend  (Go)           : Handles user data management, authentication, and PDF generation.
    Frontend (Python,Flask) : Provides the web interface for user interaction.
    Database (SQL)          : Stores user information.
```

## Prerequisites

- `Go`: _Backend development_
- `Flask`: _Frontend development_
- `MySQL database`: _Storing user data and templates_
- `wkhtmltopdf`: _PDF generation_
- `Git`: _Version control_
- `Docker & Docker Compose`: _Containerized deployment_

## Install Basic Tools

```sh
sudo apt update && sudo apt upgrade && sudo apt install -y git curl build-essential golang-go python3 python3-pip wkhtmltopdf docker.io docker-compose selinux-utils curl mysql-server
sudo mysql_secure_installation
pip install --break-system-packages Flask Flask-Bcrypt Flask-Migrate Flask-SQLAlchemy
```

## Replace Paths

```sh
PathCvProject="/bcn/github/CV_project"
grep -q "PathCvProject=" ~/.bashrc || echo "export PathCvProject=\"$PathCvProject\"                                         # Set path to CV project." >> ~/.bashrc && source ~/.bashrc
```

## Create DB, users table

```sh
sudo mysql -u root -p
CREATE DATABASE IF NOT EXISTS users;
USE users;
SOURCE /bcn/github/CV_project/sql/schemadump.sql;
```

## Verify successful import

```sh
mysql -u root -p users
SHOW DATABASES;
SHOW TABLES;
USE users;
DESCRIBE template;
DESCRIBE users;
SELECT * FROM template;
SELECT * FROM users;
```

## Change DB user password

```sh
ALTER USER 'CV_user'@'localhost' IDENTIFIED BY 'Y0ur_strong_password';
```

## Build the backend API

```sh
cd $PathCvProject/api
go mod tidy
go build -o CV_project main.go
export DB_USER="root"
export DB_PASSWORD="?????????????"
./CV_project
```

## BFF Flask app setup frontend

```sh
cd $PathCvProject/bff
python3 app.py -i 127.0.0.1 -p 8080
```

## Github

## SSH conection

```sh
GitSshKey="/PathTo/.ssh/github_rsa"
GitUsername="YourUsername"
GitEmail="YourEmail"
chmod 600 "$GitSshKey"
ssh-add "$GitSshKey"
git config --global user.name "$GitUsername"
git config --global user.email "$GitEmail"
git config --global http.sslBackend "openssl"
ssh -T git@github.com
```

### Commit & pull-push, avoid conflicts

```sh
echo "Enter commit message (Title case, infinitive verb, brief and clear summary of changes):"
read -p "CommitMssg: - " CommitMssg
cd "$PathCvProject" || exit
git add .
git commit -m "$CommitMssg"
git pull && git push origin main
```

## Start the project

```sh
cd $PathCvProject && make
```

## Docker

```sh
docker-compose build  # build
docker-compose up     # start
docker-compose up -d  # run background
docker-compose stop   # only stop
docker-compose down   # stops and removes containers
docker-compose ps     # view running containers
docker-compose rm     # removes stopped service containers
```

## Browser links

<https://miro.com/app/board/uXjVK6HA_1A=/>

<http://127.0.0.1:5000/template1>

<http://127.0.0.1:5000/template2>

<http://127.0.0.1:5000/template3>

---
