# CV_project/bff/app/app.py

"""
Frontend with flask to interact with the API
Routes
"""

import logging
import argparse
import json
import os

import requests
from flask import Flask, abort, jsonify, render_template, send_from_directory, request
from dotenv import load_dotenv
from flask_cors import CORS

print("Start")
load_dotenv()
parser = argparse.ArgumentParser()
parser.add_argument("-i", "--ip", help="API IP", default="cv_api_container")
parser.add_argument("-p", "--port", help="API PORT", default="8080")
args = vars(parser.parse_args())

IP = args["ip"]
PORT = args["port"]

app = Flask(__name__, template_folder="../templates")
CORS(app)

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger("flask.app")


@app.route("/", methods=["GET"])
def home():
    """Home route to fetch and display users."""
    url = f"http://{IP}:{PORT}/users"
    app.logger.info("Fetching users from: %s", url)
    try:
        print(f"Fetching users from: {url}")
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
    app.logger.info("Sending POST request to %s", url)
    try:
        requests.post(url=url, timeout=10)
    except requests.exceptions.RequestException as e:
        app.logger.error("Add user request failed: %s", e, exc_info=True)
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html")


@app.route("/user/<user_id>", methods=["PUT"])
def edit_user(user_id):
    """route to edit an existing user"""
    edited_data = request.json
    url = f"http://{IP}:{PORT}/user/{user_id}"
    headers = {"Content-Type": "application/json"}
    try:
        requests.put(url, data=json.dumps(edited_data), headers=headers, timeout=10)
    except requests.exceptions.RequestException as e:
        app.logger.error("Edit user request failed: %s", e, exc_info=True)
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
    try:
        u = requests.get(url=url, timeout=10)
        u.raise_for_status()
        data = u.json()
    except requests.exceptions.RequestException as e:
        app.logger.error("Get user request failed: %s", e, exc_info=True)
        abort(500, description="Internal Server Error")
    return render_template("forms/edit_form.html", data=data)


@app.route("/template1/<user_id>", methods=["GET"])
def generate_template1(user_id):
    """route to generate template 1"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=1&user={user_id}"

    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for user data failed: %s", e, exc_info=True)
        abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for template 1 PDF failed: %s", e, exc_info=True)
        abort(500, description="Internal Server Error")

    return render_template("view/template1.html", data=data, pdf_data=pdf_data)


@app.route("/template2", methods=["GET"])
def generate_template2():
    """route to generate template 2"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf?template=2"

    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for user data failed: %s", e, exc_info=True)
        abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for template 2 PDF failed: %s", e, exc_info=True)
        abort(500, description="Internal Server Error")

    return render_template("view/template2.html", data=data, pdf_data=pdf_data)


@app.route("/template3", methods=["GET"])
def generate_template3():
    """route to generate template 3"""
    url = f"http://{IP}:{PORT}/user"
    url2 = f"http://{IP}:{PORT}/pdf"
    try:
        r = requests.get(url=url, timeout=10)
        r.raise_for_status()
        data = r.json()
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for user data failed: %s", e, exc_info=True)
        return abort(500, description="Internal Server Error")

    try:
        r2 = requests.get(url=url2, timeout=10)
        r2.raise_for_status()
        pdf_data = r2.content
    except requests.exceptions.RequestException as e:
        app.logger.error("Request for template 3 PDF failed: %s", e, exc_info=True)
        return abort(500, description="Internal Server Error")

    return render_template("view/template3.html", data=data, pdf_data=pdf_data)


@app.route("/login", methods=["GET", "POST"])
def loginuser():
    """route to login a user"""
    if request.method == "GET":
        return render_template("forms/loginform.html")
    elif request.method == "POST":
        url = f"http://{IP}:{PORT}/login"
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
            app.logger.error("Login request failed: %s", e, exc_info=True)
            return render_template("forms/loginform.html")
    return abort(405)


@app.route("/logout", methods=["GET"])
def logoutuser():
    """route to logout a user"""
    url = f"http://{IP}:{PORT}/logout"
    try:
        r = requests.post(url, timeout=10)
        r.raise_for_status()
    except requests.exceptions.RequestException as e:
        app.logger.error("Logout request failed: %s", e, exc_info=True)
        return abort(500, description="Internal Server Error")
    return render_template("view/home.html")


@app.route("/signup", methods=["GET", "POST"])
def signupuser():
    """route to sign up a user"""
    if request.method == "GET":
        return render_template("forms/signupform.html")
    if request.method == "POST":
        url = f"http://{IP}:{PORT}/signup"
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
            app.logger.error("Signup request failed: %s", e, exc_info=True)
            return render_template("forms/signupform.html")
    return abort(405)


@app.route("/user", methods=["DELETE"])
def delete_user():
    """route to delete a user"""
    user_id = request.args.get("id")

    if not user_id:
        app.logger.error("User ID is missing")
        return jsonify({"error": "User ID required"}), 400

    url = f"http://{IP}:{PORT}/user/{user_id}"
    app.logger.info("Sending DELETE request to %s", url)
    try:
        response = requests.delete(url, timeout=10)
        response.raise_for_status()
        app.logger.info("Successfully deleted user with ID %s", user_id)
        return jsonify({"message": f"User {user_id} deleted successfully"}), 200
    except requests.exceptions.HTTPError as http_err:
        app.logger.error("HTTP error occurred: %s", http_err, exc_info=True)
        return jsonify({"error": "User deletion failed", "details": str(http_err)}), 500
    except requests.exceptions.RequestException as req_err:
        app.logger.error("Request error occurred: %s", req_err, exc_info=True)
        return jsonify({"error": "User deletion failed", "details": str(req_err)}), 500


@app.route("/favicon.ico")
def favicon():
    """Serve the favicon.ico file."""
    print(
        f"Fetching favicon from: {os.path.join(app.root_path, 'static/images/favicon.ico')}"
    )
    return app.send_static_file("favicon.ico")


@app.route("/<path:filename>")
def serve_css(filename):
    """Serve CSS files from the static/styles directory."""
    print(
        f"Fetching CSS from: {os.path.join(app.root_path, 'static/styles', filename)}"
    )
    return send_from_directory(os.path.join(app.root_path, "static/styles"), filename)


# CV_project/bff/static/js
@app.route("/js/<path:filename>")
def serve_js(filename):
    """Serve JavaScript files from the static/js directory."""
    print(f"Fetching JS from: {os.path.join(app.root_path, 'static/js', filename)}")
    return send_from_directory(os.path.join(app.root_path, "static/js"), filename)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000, debug=True)
