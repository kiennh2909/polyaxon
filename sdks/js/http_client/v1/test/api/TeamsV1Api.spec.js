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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.1.6
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.TeamsV1Api();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('TeamsV1Api', function() {
    describe('createTeam', function() {
      it('should call createTeam successfully', function(done) {
        //uncomment below and update the code to test createTeam
        //instance.createTeam(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createTeamMember', function() {
      it('should call createTeamMember successfully', function(done) {
        //uncomment below and update the code to test createTeamMember
        //instance.createTeamMember(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteTeam', function() {
      it('should call deleteTeam successfully', function(done) {
        //uncomment below and update the code to test deleteTeam
        //instance.deleteTeam(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteTeamMember', function() {
      it('should call deleteTeamMember successfully', function(done) {
        //uncomment below and update the code to test deleteTeamMember
        //instance.deleteTeamMember(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getTeam', function() {
      it('should call getTeam successfully', function(done) {
        //uncomment below and update the code to test getTeam
        //instance.getTeam(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getTeamMember', function() {
      it('should call getTeamMember successfully', function(done) {
        //uncomment below and update the code to test getTeamMember
        //instance.getTeamMember(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listTeamMembers', function() {
      it('should call listTeamMembers successfully', function(done) {
        //uncomment below and update the code to test listTeamMembers
        //instance.listTeamMembers(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listTeamNames', function() {
      it('should call listTeamNames successfully', function(done) {
        //uncomment below and update the code to test listTeamNames
        //instance.listTeamNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listTeams', function() {
      it('should call listTeams successfully', function(done) {
        //uncomment below and update the code to test listTeams
        //instance.listTeams(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchTeam', function() {
      it('should call patchTeam successfully', function(done) {
        //uncomment below and update the code to test patchTeam
        //instance.patchTeam(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchTeamMember', function() {
      it('should call patchTeamMember successfully', function(done) {
        //uncomment below and update the code to test patchTeamMember
        //instance.patchTeamMember(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateTeam', function() {
      it('should call updateTeam successfully', function(done) {
        //uncomment below and update the code to test updateTeam
        //instance.updateTeam(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateTeamMember', function() {
      it('should call updateTeamMember successfully', function(done) {
        //uncomment below and update the code to test updateTeamMember
        //instance.updateTeamMember(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
