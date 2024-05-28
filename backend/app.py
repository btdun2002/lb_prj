from flask import Flask, request

app = Flask(__name__)

@app.route('/')
def index():
    return "Hello from Backend Server! This is the main endpoint.\n"

@app.route('/health')
def health():
    return "OK",200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
