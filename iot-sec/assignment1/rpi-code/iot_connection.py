# test_connection.py
from azure.iot.device import IoTHubDeviceClient
from iot_secrets import SHARED_ACCESS as accesskey
import json
import time

# Replace with your actual connection string
HOSTNAME = "iot-sec-hub.azure-devices.net"
DEVICE_ID = "rpi-group4"
SHARED_ACCESS = accesskey

CONNECTION_STRING = f"HostName={HOSTNAME};DeviceId={DEVICE_ID};SharedAccessKey={SHARED_ACCESS}"

def test_connection():
    try:
        print("Testing connection...")
        print(f"Hub: {HOSTNAME}")
        print(f"Device: {DEVICE_ID}")
        
        # Create client with more detailed error handling
        client = IoTHubDeviceClient.create_from_connection_string(CONNECTION_STRING)
        
        # Try to connect
        client.connect()
        print("✅ Successfully connected to IoT Hub!")
        
        # Send test message
        client.send_message('{"test": "connection successful"}')
        print("✅ Test message sent!")
        
        client.disconnect()
        
    except Exception as e:
        print(f"❌ Connection failed: {e}")
        print(f"Error type: {type(e).__name__}")


def send_message(payload):
    try:
        client = IoTHubDeviceClient.create_from_connection_string(CONNECTION_STRING)
        client.connect()
        
        message_json = json.dumps(payload)
        client.send_message(message_json)
        print(f"✅ Message sent: {message_json}")
        
        client.disconnect()
    except Exception as e:
        print(f"❌ Failed to send message: {e}")

if __name__ == "__main__":
    test_connection()