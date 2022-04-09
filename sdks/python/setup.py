from distutils.core import setup

setup(name='GoogleApiTutorials',
      version='1.0',
      python_version='>=3.7',
      description='Libraries required for examples',
      author='Yu Watanabe',
      author_email='yu.w.tennis@gmail.com',
      url='',
      packages=['google_api_tutorials'],
      install_requires=[
            'google-auth ~= 2.6.2',
            'google-api-python-client ~= 2.43.0'
      ]
      )
