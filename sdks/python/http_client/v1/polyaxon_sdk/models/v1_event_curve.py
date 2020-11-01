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

    The version of the OpenAPI document: 1.2.2-rc0
    Contact: contact@polyaxon.com
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from polyaxon_sdk.configuration import Configuration


class V1EventCurve(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'kind': 'V1EventCurveKind',
        'x': 'list[float]',
        'y': 'list[float]',
        'annotation': 'str'
    }

    attribute_map = {
        'kind': 'kind',
        'x': 'x',
        'y': 'y',
        'annotation': 'annotation'
    }

    def __init__(self, kind=None, x=None, y=None, annotation=None, local_vars_configuration=None):  # noqa: E501
        """V1EventCurve - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._kind = None
        self._x = None
        self._y = None
        self._annotation = None
        self.discriminator = None

        if kind is not None:
            self.kind = kind
        if x is not None:
            self.x = x
        if y is not None:
            self.y = y
        if annotation is not None:
            self.annotation = annotation

    @property
    def kind(self):
        """Gets the kind of this V1EventCurve.  # noqa: E501


        :return: The kind of this V1EventCurve.  # noqa: E501
        :rtype: V1EventCurveKind
        """
        return self._kind

    @kind.setter
    def kind(self, kind):
        """Sets the kind of this V1EventCurve.


        :param kind: The kind of this V1EventCurve.  # noqa: E501
        :type: V1EventCurveKind
        """

        self._kind = kind

    @property
    def x(self):
        """Gets the x of this V1EventCurve.  # noqa: E501


        :return: The x of this V1EventCurve.  # noqa: E501
        :rtype: list[float]
        """
        return self._x

    @x.setter
    def x(self, x):
        """Sets the x of this V1EventCurve.


        :param x: The x of this V1EventCurve.  # noqa: E501
        :type: list[float]
        """

        self._x = x

    @property
    def y(self):
        """Gets the y of this V1EventCurve.  # noqa: E501


        :return: The y of this V1EventCurve.  # noqa: E501
        :rtype: list[float]
        """
        return self._y

    @y.setter
    def y(self, y):
        """Sets the y of this V1EventCurve.


        :param y: The y of this V1EventCurve.  # noqa: E501
        :type: list[float]
        """

        self._y = y

    @property
    def annotation(self):
        """Gets the annotation of this V1EventCurve.  # noqa: E501


        :return: The annotation of this V1EventCurve.  # noqa: E501
        :rtype: str
        """
        return self._annotation

    @annotation.setter
    def annotation(self, annotation):
        """Sets the annotation of this V1EventCurve.


        :param annotation: The annotation of this V1EventCurve.  # noqa: E501
        :type: str
        """

        self._annotation = annotation

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1EventCurve):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1EventCurve):
            return True

        return self.to_dict() != other.to_dict()
