#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import os

from polyaxon import settings
from polyaxon.env_vars.getters.project import get_project_or_local
from polyaxon.env_vars.keys import (
    POLYAXON_KEYS_COLLECT_ARTIFACTS,
    POLYAXON_KEYS_COLLECT_RESOURCES,
    POLYAXON_KEYS_RUN_INSTANCE,
)
from polyaxon.exceptions import PolyaxonClientException
from polyaxon.managers.run import RunManager
from polyaxon.utils.bool_utils import to_bool


def get_run_or_local(run_uuid=None, is_cli: bool = False):
    if run_uuid:
        return run_uuid
    if is_cli:
        return RunManager.get_config_or_raise().uuid

    run = RunManager.get_config()
    if run:
        return run.uuid
    return None


def get_project_run_or_local(project=None, run_uuid=None, is_cli: bool = True):
    user, project_name = get_project_or_local(project, is_cli=is_cli)
    run_uuid = get_run_or_local(run_uuid, is_cli=is_cli)
    return user, project_name, run_uuid


def get_collect_artifact():
    """If set, Polyaxon will collect artifacts"""
    return to_bool(os.getenv(POLYAXON_KEYS_COLLECT_ARTIFACTS, None), handle_none=True)


def get_collect_resources():
    """If set, Polyaxon will collect resources"""
    return to_bool(os.getenv(POLYAXON_KEYS_COLLECT_RESOURCES, None), handle_none=True)


def get_log_level():
    """If set on the polyaxonfile it will return the log level."""
    return settings.CLIENT_CONFIG.log_level


def get_run_info(run_instance: str = None):
    run_instance = run_instance or os.getenv(POLYAXON_KEYS_RUN_INSTANCE, None)
    if not run_instance:
        raise PolyaxonClientException(
            "Could not get run info, "
            "please make sure this is run is correctly started by Polyaxon."
        )

    parts = run_instance.split(".")
    if not len(parts) == 4:
        raise PolyaxonClientException(
            "run instance is invalid `{}`, "
            "please make sure this is run is correctly started by Polyaxon.".format(
                run_instance
            )
        )
    return parts[0], parts[1], parts[-1]
