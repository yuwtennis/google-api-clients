""" """

import sys
from pathlib import Path

PACKAGE_ROOT = str(Path(__file__).resolve().parent.parent.parent)
sys.path.append(PACKAGE_ROOT)

from google_api_tutorials.factories.credential import CredentialFactory
from google_api_tutorials.factories.service import ServiceFactory

if __name__ == "__main__":

    print(PACKAGE_ROOT)
    print(sys.path)
