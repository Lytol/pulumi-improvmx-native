# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .random import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "improvmx",
  "mod": "index",
  "fqn": "pulumi_improvmx",
  "classes": {
   "improvmx:index:Random": "Random"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "improvmx",
  "token": "pulumi:providers:improvmx",
  "fqn": "pulumi_improvmx",
  "class": "Provider"
 }
]
"""
)
