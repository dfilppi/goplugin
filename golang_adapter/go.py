########
# Copyright (c) 2019 Cloudify Platform Ltd. All rights reserved
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#        http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
############


from cloudify.exceptions import NonRecoverableError
from cloudify import ctx
from cloudify.proxy.server import HTTPCtxProxy
import os
import json
import subprocess
import tempfile


def install():
    """ Perform a lazy install of golang source
    """
    # if plugin_path exists, this was already done
    if 'plugin_path' in ctx.instance.runtime_properties:
        return

    # verify go available
    res = subprocess.call(["go", "version"])
    if res != 0:
        raise NonRecoverableError("go not available")

    # get go source: construct source tree
    dirpath = tempfile.mkdtemp()

    gosrc_dir = os.path.dirname(os.path.abspath(__file__))+"/src"
    subprocess.call(["cp", "-r", gosrc_dir, dirpath])

    # set gopath and build
    os.environ['GOPATH'] = dirpath
    res = subprocess.call(["go", "install", "plugin"])
    if res != 0:
        raise NonRecoverableError("go install failed on plugin")

    ctx.instance.runtime_properties['plugin_path'] = dirpath + "/bin/plugin"


##
# Plugin operations
##
def callgo(func, args, **kwargs):
    """ Should do whatever "create" is defined as in go code.  It is assumed
        that the go implementation main function accepts a funcion name and
        a varargs list
    """
    install()
    proxy_server = HTTPCtxProxy(ctx._get_current_object())

    try:
        exepath = ctx.instance.runtime_properties['plugin_path']
        # below will fail on windows
        res = subprocess.call(
            [exepath, str(proxy_server.port), func, json.dumps(args)])
        if res != 0:
            raise NonRecoverableError("func {} execution faild".format(func))
    finally:
        proxy_server.close()
