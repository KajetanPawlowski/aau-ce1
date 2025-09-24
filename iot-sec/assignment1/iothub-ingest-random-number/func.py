import json
import datetime
import os
import azure.functions as func
from azure.cosmos import CosmosClient

# env vars
COSMOS_URL = os.environ["COSMOS_URL"]
COSMOS_KEY = os.environ["COSMOS_KEY"]
DATABASE_NAME = "IoTData"
CONTAINER_NAME = "Telemetry"

client = CosmosClient(COSMOS_URL, credential=COSMOS_KEY)
container = client.get_database_client(DATABASE_NAME).get_container_client(CONTAINER_NAME)

def main(event: func.EventHubEvent):
    body = json.loads(event.get_body())
    # add server timestamp if needed
    body['ingest_time'] = datetime.datetime.utcnow().isoformat()
    container.create_item(body)