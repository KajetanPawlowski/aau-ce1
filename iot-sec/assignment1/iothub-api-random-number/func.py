import datetime
import os
import azure.functions as func
from azure.cosmos import CosmosClient

COSMOS_URL = os.environ["COSMOS_URL"]
COSMOS_KEY = os.environ["COSMOS_KEY"]
DATABASE_NAME = "IoTData"
CONTAINER_NAME = "Telemetry"

client = CosmosClient(COSMOS_URL, credential=COSMOS_KEY)
container = client.get_database_client(DATABASE_NAME).get_container_client(CONTAINER_NAME)

def main(req: func.HttpRequest) -> func.HttpResponse:
    period = req.route_params.get('period')  # 'today', 'last-hour', 'last-5min'
    now = datetime.datetime.utcnow()

    if period == 'today':
        start = now.replace(hour=0, minute=0, second=0, microsecond=0)
    elif period == 'last-hour':
        start = now - datetime.timedelta(hours=1)
    elif period == 'last-5min':
        start = now - datetime.timedelta(minutes=5)
    else:
        return func.HttpResponse("Invalid period", status_code=400)

    query = "SELECT VALUE AVG(c.value) FROM c WHERE c.timestamp >= @start"
    items = list(container.query_items(
        query=query,
        parameters=[{"name":"@start", "value": start.isoformat()}],
        enable_cross_partition_query=True))

    avg = items[0] if items else None
    return func.HttpResponse(str(avg), mimetype="application/json")