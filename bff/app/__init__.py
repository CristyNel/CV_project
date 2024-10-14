from flask import Flask
from flask_bcrypt import Bcrypt
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from decouple import config

# initialize flask app
app = Flask(__name__, template_folder="../templates", static_folder="../static")

# load configuration from environment variable
app.config.from_object(config("APP_SETTINGS"))

# initialize extensions
bcrypt = Bcrypt(app)
db = SQLAlchemy(app)
migrate = Migrate(app, db)

# import routes or additional app setup here
from . import app
