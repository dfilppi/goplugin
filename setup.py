import setuptools
import os

files = None
for _,_,files in os.walk('golang_adapter/src/plugin'):
  files = files

setuptools.setup(
    name='golang-test-plugin',
    version='0.0.1',
    author='dfilppi',
    author_email='dfilppi@gmail.com',
    description='example golang plugin for cloudify',
    packages=['golang_adapter'],
    package_data = {'golang_adapter': [ 'src/plugin/'+ f for f in files ] },
    install_requires = [
      'cloudify-plugins-common'
    ],
    license='LICENSE'
)
