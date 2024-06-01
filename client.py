import requests

def send_message(message):
    server_url = "http://localhost:8080/message"
    response = requests.post(server_url, json={'message': message})
    return response.text

if __name__ == '__main__':
    while True:
        user_input = input("Enter a message: ")
        response = send_message(user_input)
        print(f"Server response: {response}")
