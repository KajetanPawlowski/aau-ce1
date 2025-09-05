from flask import Flask

# Initialize the Flask application
app = Flask(__name__)

# Define a route for the root URL
@app.route('/')
def home():
    return "Hello, World!"

# Run the app when this script is executed directly
if __name__ == '__main__':
    # Start the development server on host 0.0.0.0 to allow external access
    app.run(host='0.0.0.0', port=5000, debug=True)
