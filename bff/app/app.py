# * CV_project/bff/app/app.py

"""
Frontend with flask to interact with the API
Routes
"""

import logging
import json
import os

import requests
from flask import Flask, request, jsonify, render_template, send_from_directory, abort
from dotenv import load_dotenv
from flask_cors import CORS

load_dotenv()

IP = os.getenv("API_IP", "cv_api-service")
PORT = os.getenv("API_PORT", "8080")

project_root = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
template_folder = os.path.join(project_root, "templates")
static_folder = os.path.join(project_root, "static")
app = Flask(__name__, template_folder=template_folder, static_folder=static_folder)
CORS(app)

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger("flask.app")
logger.propagate = False  # Disable log propagation

app.logger.info("\033[1;96;1m * * * 🔗 project_root    = %s\033[0m", project_root)
app.logger.info("\033[1;96;1m * * * 🔗 template_folder = %s\033[0m", template_folder)
app.logger.info("\033[1;96;1m * * * 🔗 static_folder   = %s\033[0m", static_folder)


@app.route("/", methods=["GET"])
def home():
    """Home route to fetch and display users."""
    url = f"http://{IP}:{PORT}/users"
    app.logger.info("\033[1;96;1m * * * 👤 Fetching users from: %s\033[0m", url)
    try:
        response = requests.get(url=url, timeout=12)
        response.raise_for_status()
        if response.status_code == 200 and response.text:
            data = response.json()
            if data is None:
                app.logger.error("\033[1;91;1m * * * 🆘 Received None data \033[0m")
                return abort(500, description="Internal Server Error")
        else:
            app.logger.error(
                "\033[1;91;1m * * * 🆘 Received empty response or non-200 status \033[0m"
            )
            return abort(500, description="Internal Server Error")
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request failed: %s\033[0m", e, exc_info=True
        )
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html", users=data)


@app.route("/users", methods=["GET"])
def get_users():
    """Route to fetch and display all users."""
    url = f"http://{IP}:{PORT}/users"
    app.logger.info("\033[1;96;1m * * * 👤 Fetching users from: %s\033[0m", url)
    try:
        response = requests.get(url=url, timeout=12)
        response.raise_for_status()
        if response.status_code == 200 and response.text:
            data = response.json()
        else:
            app.logger.error("Received empty response or non-200 status")
            return abort(500, description="Internal Server Error")
    except requests.exceptions.RequestException as e:
        app.logger.error("Request failed: %s", e, exc_info=True)
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html", users=data)


@app.route("/user", methods=["POST"])
def add_user():
    """route to add a new user"""
    url = f"http://{IP}:{PORT}/user"
    app.logger.info("\033[1;96;1m * * * 👤 Add new user, POST to %s\033[0m", url)
    try:
        requests.post(url=url, timeout=10)
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Add user request failed: %s\033[0m", e, exc_info=True
        )
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html")


@app.route("/user/<user_id>", methods=["PUT"])
def edit_user(user_id):
    """route to edit an existing user"""
    edited_data = request.json
    url = f"http://{IP}:{PORT}/user/{user_id}"
    app.logger.info("\033[1;96;1m * * * 📝 Edit existing user, PUT to %s\033[0m", url)
    headers = {"Content-Type": "application/json"}
    try:
        requests.put(url, data=json.dumps(edited_data), headers=headers, timeout=10)
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Edit user request failed: %s\033[0m",
            e,
            exc_info=True,
        )
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html")


@app.route("/postform", methods=["GET"])
def get_postform():
    """route to display the post form"""
    return render_template("forms/post_form.html")


@app.route("/editform", methods=["GET"])
def get_user():
    """route to display the edit form"""
    url = f"http://{IP}:{PORT}/user"
    app.logger.info("\033[1;96;1m * * * 📝 Edit form, GET to %s\033[0m", url)
    try:
        u = requests.get(url=url, timeout=10)
        u.raise_for_status()
        data = u.json()
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Get user request failed: %s\033[0m", e, exc_info=True
        )
        abort(500, description="Internal Server Error")
    return render_template("forms/edit_form.html", data=data)


@app.route("/template1/<user_id>", methods=["GET"])
def generate_template1(user_id):
    """route to generate template 1"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=1&user={user_id}"
    app.logger.info("\033[1;96;1m * * * 🎨 Generate template 1, GET to %s\033[0m", url)
    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for user data failed: %s\033[0m",
            e,
            exc_info=True,
        )
        abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for template 1 PDF failed: %s\033[0m",
            e,
            exc_info=True,
        )
        abort(500, description="Internal Server Error")

    return render_template("view/template1.html", data=data, pdf_data=pdf_data)


@app.route("/template2", methods=["GET"])
def generate_template2():
    """route to generate template 2"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=2"
    app.logger.info("\033[1;96;1m * * * 🎨 Generate template 2, GET to %s\033[0m", url)
    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for user data failed: %s\033[0m",
            e,
            exc_info=True,
        )
        abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for template 2 PDF failed: %s\033[0m",
            e,
            exc_info=True,
        )
        abort(500, description="Internal Server Error")

    return render_template("view/template2.html", data=data, pdf_data=pdf_data)


@app.route("/template3", methods=["GET"])
def generate_template3():
    """route to generate template 3"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    app.logger.info("\033[1;96;1m * * * 🎨 Generate template 3, GET to %s\033[0m", url)
    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for user data failed: %s\033[0m",
            e,
            exc_info=True,
        )
        return abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request for template 3 PDF failed: %s\033[0m",
            e,
            exc_info=True,
        )
        return abort(500, description="Internal Server Error")

    return render_template("view/template3.html", data=data, pdf_data=pdf_data)


@app.route("/login", methods=["GET", "POST"])
def loginuser():
    """route to login a user"""
    if request.method == "GET":
        app.logger.info("\033[1;96;1m * * * 🔓 Login, GET to %s\033[0m", url)
        return render_template("forms/loginform.html")
    elif request.method == "POST":
        url = f"http://{IP}:{PORT}/login"
        app.logger.info("\033[1;96;1m * * * 🔓 Login, POST to %s\033[0m", url)
        try:
            r = requests.post(
                url, data=request.form, headers=request.headers, timeout=10
            )
            r.raise_for_status()
            if r.status_code == 200:
                return render_template("view/greet.html")
            else:
                return render_template("forms/loginform.html")
        except requests.exceptions.RequestException as e:
            app.logger.error(
                "\033[1;91;1m * * * 🆘 Login request failed: %s\033[0m",
                e,
                exc_info=True,
            )
            return render_template("forms/loginform.html")
    return abort(405)


@app.route("/logout", methods=["GET"])
def logoutuser():
    """route to logout a user"""
    url = f"http://{IP}:{PORT}/logout"
    app.logger.info("\033[1;96;1m * * * 🔓 Logout, GET to %s\033[0m", url)
    try:
        r = requests.post(url, timeout=10)
        r.raise_for_status()
    except requests.exceptions.RequestException as e:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Logout request failed: %s\033[0m", e, exc_info=True
        )
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html")


@app.route("/signup", methods=["GET", "POST"])
def signupuser():
    """route to sign up a user"""
    if request.method == "GET":
        app.logger.info("\033[1;96;1m * * * 🔐 Signup, GET to %s\033[0m", url)
        return render_template("forms/signupform.html")
    if request.method == "POST":
        url = f"http://{IP}:{PORT}/signup"
        app.logger.info("\033[1;96;1m * * * 🔐 Signup, POST to %s\033[0m", url)
        try:
            r = requests.post(
                url, data=request.form, headers=request.headers, timeout=10
            )
            r.raise_for_status()
            if r.status_code == 200:
                return render_template("forms/signupform.html")
            else:
                return render_template("forms/signupform.html")
        except requests.exceptions.RequestException as e:
            app.logger.error(
                "\033[1;91;1m * * * 🆘 Signup request failed: %s\033[0m",
                e,
                exc_info=True,
            )
            return render_template("forms/signupform.html")
    return abort(405)


@app.route("/user", methods=["DELETE"])
def delete_user():
    """Route to delete a user"""
    user_id = request.args.get("id")

    if not user_id:
        app.logger.error("\033[1;91;1m * * * 🆘 User ID is missing: %s\033[0m", user_id)
        return jsonify({"success": False, "error": "User ID required"}), 400

    url = f"http://{IP}:{PORT}/user/{user_id}"
    app.logger.info("\033[1;96;1m * * * 🧹 Delete user, DELETE to %s\033[0m", url)
    try:
        response = requests.delete(url, timeout=10)
        response.raise_for_status()
        app.logger.info(
            "\033[1;96;1m * * * ✅ Successfully deleted user with ID %s\033[0m", user_id
        )
        return (
            jsonify(
                {"success": True, "message": f"User {user_id} deleted successfully"}
            ),
            200,
        )
    except requests.exceptions.HTTPError as http_err:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 HTTP error occurred: %s\033[0m",
            http_err,
            exc_info=True,
        )
        return (
            jsonify(
                {
                    "success": False,
                    "error": "User deletion failed",
                    "details": str(http_err),
                }
            ),
            500,
        )
    except requests.exceptions.RequestException as req_err:
        app.logger.error(
            "\033[1;91;1m * * * 🆘 Request error occurred: %s\033[0m",
            req_err,
            exc_info=True,
        )
        return (
            jsonify(
                {
                    "success": False,
                    "error": "User deletion failed",
                    "details": str(req_err),
                }
            ),
            500,
        )


@app.route("/favicon.ico")
def favicon():
    """Serve the favicon.ico file."""
    favicon_path = os.path.join(app.root_path, "..", "templates/view/favicon.ico")
    app.logger.info(
        "\033[1;96;1m * * * 🔍 Fetching favicon from: %s\033[0m", favicon_path
    )
    return send_from_directory(
        os.path.join(app.root_path, "..", "templates/view"), "favicon.ico"
    )


@app.route("/styles/<path:filename>")
def serve_css(filename):
    """Serve CSS files from the static/styles directory."""
    css_path = os.path.join(app.root_path, "..", "static/styles", filename)
    app.logger.info("\033[1;96;1m * * * 🔍 Fetching CSS from: %s\033[0m", css_path)
    return send_from_directory(
        os.path.join(app.root_path, "..", "static/styles"), filename
    )


@app.route("/js/<path:filename>")
def serve_js(filename):
    """Serve JavaScript files from the static/js directory."""
    js_path = os.path.join(app.root_path, "..", "static/js", filename)
    app.logger.info("\033[1;96;1m * * * 🔍 Fetching JS from: %s\033[0m", js_path)
    return send_from_directory(os.path.join(app.root_path, "..", "static/js"), filename)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)
