from flask import Flask
import sys

app = Flask(__name__)

serverName = "server-1"

@app.route("/")
def hello():
    return serverName