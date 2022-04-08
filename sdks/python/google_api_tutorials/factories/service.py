""" """
from googleapiclient.discovery import build


class ServiceFactory:
    """ """

    @staticmethod
    def create(service_name, version, credentials, **kwargs):
        """

        Parameters
        ----------
        service_name: str
            Service to use
        version: str
            API version
        credentials: google.auth.credentials.Credentials
        kwargs:
            See
            https://googleapis.github.io/google-api-python-client/docs/epy/googleapiclient.discovery-module.html#build
        """

        return build(service_name, version, credentials=credentials, **kwargs)
