""" """

import sys
from pathlib import Path

PACKAGE_ROOT = str(Path(__file__).parent.parent.parent.resolve())
sys.path.append(PACKAGE_ROOT)

if __name__ == "__main__":

    print(PACKAGE_ROOT)
    print(sys.path)