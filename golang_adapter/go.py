from cloudify.decorators import operation
from cloudify.exceptions import NonRecoverableError, RecoverableError
from cloudify import ctx
from cloudify.proxy.server import HTTPCtxProxy
from functools import wraps
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
  res = subprocess.call(["go", "version" ])
  if res != 0:
    raise NonRecoverableError("go not available")

  # get go source: construct source tree
  dirpath = tempfile.mkdtemp()

  gosrc_dir = os.path.dirname(os.path.abspath(__file__))+"/src"
  subprocess.call(["cp","-r",gosrc_dir,dirpath])
  
  # set gopath and build
  os.environ['GOPATH'] = dirpath
  res = subprocess.call(["go", "install", "plugin"])
  if res != 0:
    raise NonRecoverableError("go install failed on plugin")

  ctx.instance.runtime_properties['plugin_path'] = dirpath + "/bin/plugin"


##
## Plugin operations
##
def callgo( func, args , **kwargs):
  """ Should do whatever "create" is defined as in go code.  It is assumed
      that the go implementation main function accepts a funcion name and
      a varargs list
  """
  install()
  proxy_server = HTTPCtxProxy(ctx._get_current_object())
  
  try:
    exepath = ctx.instance.runtime_properties['plugin_path']
    # below will fail on windows
    res = subprocess.call([ exepath , str(proxy_server.port), func, json.dumps( args )])
    if res == 1 :
      raise RecoverableError("func {} execution failed".format(func))
    if res == 2 :
      raise NonRecoverableError("func {} execution failed".format(func))
  finally:
    proxy_server.close()
