""" """
import datetime
import sys
import logging
import time
import os
from pathlib import Path

PACKAGE_ROOT = str(Path(__file__).resolve().parent.parent.parent)
sys.path.append(PACKAGE_ROOT)

from google_api_tutorials.factories.credential import CredentialFactory
from google_api_tutorials.factories.service import ServiceFactory

if __name__ == "__main__":

    logging.basicConfig(format="%(asctime)s - %(name)s - %(levelname)s - %(message)s")
    logger = logging.getLogger(__name__)
    logger.setLevel(logging.INFO)

    scopes = ["https://www.googleapis.com/auth/admin.reports.audit.readonly"]
    subject_email = os.getenv("GOOGLE_SUBJECT_EMAIL")
    next_page_token = ""
    start_time = datetime.datetime.utcnow() - datetime.timedelta(hours=96)
    page_cnt = 1

    creds = CredentialFactory.create_with_subject(scopes,subject_email)
    service = ServiceFactory.create("admin", "reports_v1", creds)

    while True:
        params = {
            "userKey" : "all",
            "applicationName": "drive",
            "startTime": f"{start_time.isoformat()}Z"
        }

        if next_page_token != "":
            params["pageToken"] = next_page_token

        resp = service.activities().list(**params).execute()

        items = resp.get('items', [])

        logger.info("Page: %s, Received %d items", page_cnt, len(items))
        for item in items:
            logger.info("Activity Record: Id.Time: %s , Id.ApplicationName: %s, "
                        "Id.UniqueQualifier: %s, Num of Events: %d",
                        item['id']['time'],
                        item['id']['applicationName'],
                        item['id']['uniqueQualifier'],
                        len(item['events']))

        if "nextPageToken" not in resp:
            break

        logger.info("NextPageToken: %s", resp["nextPageToken"])
        next_page_token = resp.get('nextPageToken')
        page_cnt += 1

        time.sleep(5)
