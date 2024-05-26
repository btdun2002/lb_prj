from flask import Flask, request
import logging

app = Flask(__name__)

logging.basicConfig(level=logging.INFO)

@app.route('/')
def index():
    app.logger.info(f"Request from: {request.remote_addr}")
    return "Hello from Backend Server! This is the main endpoint."

@app.route('/health')
def health():
    return "OK",200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
