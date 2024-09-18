# CV Management System API

## Overview : This project provides a platform for managing user data and generating custom CV templates

![Cv_examples](https://venngage-wordpress.s3.amazonaws.com/uploads/2021/11/section-3-resume-banner-1-1.png)

## Screenshots: development

![SCreenshot: 18/09](./bff/static/img/screenshots/Screenshot_18-09-24.png)

```sh
CV_project
├── api                                 # Go API         - backend
│   ├── CV_project                      # Executable file ...  (cd $PathCvProject/api && go build -o CV_project main.go)
│   ├── example.pdf                     # Sample PDF
│   ├── go.mod                          # Go modules     - dependency management
│   ├── go.sum                          # Checksum       - dependencies
│   └── main.go                         # Main entry point
├── bff                                 # Flask-based    - frontend
│   ├── app.py                          # Main Flask application file
│   ├── css                             # CSS files      - styling
│   │   └── users.css                   # Custom CSS     - user-related pages
│   ├── templates                       # HTML templates - frontend
│   │   ├── edit_form.html              # HTML     form  - editing user data
│   │   ├── greet.html                  # Greeting page
│   │   ├── home.html                   # Home     page
│   │   ├── loginform.html              # Login    form
│   │   ├── populate_template.html      # Template       - populating CVs
│   │   ├── post_form.html              # Form           - posting new content
│   │   ├── signupform.html             # Signup form
│   │   ├── template1.html              # CV template 1
│   │   ├── template2.html              # CV template 2
│   │   └── template3.html              # CV template 3
│   └── users.js                        # JavaScript     - user-related functionality
├── .gitignore                          # Git ignore     - version control
├── sql                                 # SQL files      - database schema
│   ├── schemadump.sql                  # Schema creation and sample data
│   └── schema.sql                      # Schema creation only
└── src                                 # Source code
    └── __init__.py                     # Initialization and configuration
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
