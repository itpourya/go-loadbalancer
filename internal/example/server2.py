from flask import Flask
import sys

app = Flask(__name__)

serverName = "server-2"

@app.route("/")
def hello():
    return serverName