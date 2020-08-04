// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.threeten.bp.OffsetDateTime;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1ListOrganizationMembersResponse;
import org.openapitools.client.model.V1ListOrganizationsResponse;
import org.openapitools.client.model.V1Organization;
import org.openapitools.client.model.V1OrganizationMember;
import org.junit.Test;
import org.junit.Ignore;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for OrganizationsV1Api
 */
@Ignore
public class OrganizationsV1ApiTest {

    private final OrganizationsV1Api api = new OrganizationsV1Api();

    
    /**
     * Create organization
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createOrganizationTest() throws ApiException {
        V1Organization body = null;
        V1Organization response = api.createOrganization(body);

        // TODO: test validations
    }
    
    /**
     * Create organization member
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void createOrganizationMemberTest() throws ApiException {
        String owner = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.createOrganizationMember(owner, body);

        // TODO: test validations
    }
    
    /**
     * Delete organization
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOrganizationTest() throws ApiException {
        String owner = null;
        api.deleteOrganization(owner);

        // TODO: test validations
    }
    
    /**
     * Delete organization member details
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void deleteOrganizationMemberTest() throws ApiException {
        String owner = null;
        String user = null;
        api.deleteOrganizationMember(owner, user);

        // TODO: test validations
    }
    
    /**
     * Get organization
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization response = api.getOrganization(owner);

        // TODO: test validations
    }
    
    /**
     * Get organization member details
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOrganizationMemberTest() throws ApiException {
        String owner = null;
        String user = null;
        V1OrganizationMember response = api.getOrganizationMember(owner, user);

        // TODO: test validations
    }
    
    /**
     * Get organization settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void getOrganizationSettingsTest() throws ApiException {
        String owner = null;
        String organizationUser = null;
        String organizationUserEmail = null;
        String organizationName = null;
        Boolean organizationIsPublic = null;
        OffsetDateTime organizationCreatedAt = null;
        OffsetDateTime organizationUpdatedAt = null;
        String organizationRole = null;
        V1Organization response = api.getOrganizationSettings(owner, organizationUser, organizationUserEmail, organizationName, organizationIsPublic, organizationCreatedAt, organizationUpdatedAt, organizationRole);

        // TODO: test validations
    }
    
    /**
     * Get organization members
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationMembersTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        V1ListOrganizationMembersResponse response = api.listOrganizationMembers(owner, offset, limit, sort, query);

        // TODO: test validations
    }
    
    /**
     * List organizations names
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationNamesTest() throws ApiException {
        V1ListOrganizationsResponse response = api.listOrganizationNames();

        // TODO: test validations
    }
    
    /**
     * List organizations
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void listOrganizationsTest() throws ApiException {
        V1ListOrganizationsResponse response = api.listOrganizations();

        // TODO: test validations
    }
    
    /**
     * Patch organization
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.patchOrganization(owner, body);

        // TODO: test validations
    }
    
    /**
     * Patch organization member
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.patchOrganizationMember(owner, memberUser, body);

        // TODO: test validations
    }
    
    /**
     * Patch oranization settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void patchOrganizationSettingsTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.patchOrganizationSettings(owner, body);

        // TODO: test validations
    }
    
    /**
     * Update organization
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOrganizationTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.updateOrganization(owner, body);

        // TODO: test validations
    }
    
    /**
     * Update organization member
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOrganizationMemberTest() throws ApiException {
        String owner = null;
        String memberUser = null;
        V1OrganizationMember body = null;
        V1OrganizationMember response = api.updateOrganizationMember(owner, memberUser, body);

        // TODO: test validations
    }
    
    /**
     * Update organization settings
     *
     * 
     *
     * @throws ApiException
     *          if the Api call fails
     */
    @Test
    public void updateOrganizationSettingsTest() throws ApiException {
        String owner = null;
        V1Organization body = null;
        V1Organization response = api.updateOrganizationSettings(owner, body);

        // TODO: test validations
    }
    
}
