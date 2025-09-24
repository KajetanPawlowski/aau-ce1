import random
import time
import iot_connection as iot

print("Generating random numbers from 1-1000...")
print("Press Ctrl+C to stop...")

try:
    while True:
        number = random.randint(1, 1000)
        payload = f"Random number: {number}"
        iot.send_message(payload)
        time.sleep(10)  # Wait 0.5 seconds between numbers
except KeyboardInterrupt:
    print("\nStopped generating numbers.")