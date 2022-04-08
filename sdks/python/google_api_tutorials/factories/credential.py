""" """
import os

from google.oauth2 import service_account


class CredentialFactory:
    """ """

    @staticmethod
    def create_with_subject(scopes, subject_email):
        """

        Parameters
        ----------
        scopes: list
            Scopes to authorize request
        subject_email: str
            Email to impersonate

        Returns
        -------
        Credentials:
            https://google-auth.readthedocs.io/en/master/reference/google.auth.credentials.html#google.auth.credentials.Credentials

        """
        creds = service_account.Credentials.from_service_account_file(os.getenv("GOOGLE_APPLICATION_CREDENTIAL"))
        scoped_cred = creds.with_scopes(scopes)
        delegated_cred = scoped_cred.with_subject(subject_email)

        return delegated_cred
