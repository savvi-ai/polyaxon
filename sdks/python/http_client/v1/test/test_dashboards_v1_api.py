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

# coding: utf-8

"""
    Polyaxon SDKs and REST API specification.

    Polyaxon SDKs and REST API specification.  # noqa: E501

    OpenAPI spec version: 1.0.73
    Contact: contact@polyaxon.com
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from __future__ import absolute_import

import unittest

import polyaxon_sdk
from polyaxon_sdk.api.dashboards_v1_api import DashboardsV1Api  # noqa: E501
from polyaxon_sdk.rest import ApiException


class TestDashboardsV1Api(unittest.TestCase):
    """DashboardsV1Api unit test stubs"""

    def setUp(self):
        self.api = polyaxon_sdk.api.dashboards_v1_api.DashboardsV1Api()  # noqa: E501

    def tearDown(self):
        pass

    def test_create_dashboard(self):
        """Test case for create_dashboard

        """
        pass

    def test_delete_dashboard(self):
        """Test case for delete_dashboard

        """
        pass

    def test_get_dashboard(self):
        """Test case for get_dashboard

        """
        pass

    def test_list_dashboard_names(self):
        """Test case for list_dashboard_names

        """
        pass

    def test_list_dashboards(self):
        """Test case for list_dashboards

        """
        pass

    def test_patch_dashboard(self):
        """Test case for patch_dashboard

        """
        pass

    def test_update_dashboard(self):
        """Test case for update_dashboard

        """
        pass


if __name__ == "__main__":
    unittest.main()
