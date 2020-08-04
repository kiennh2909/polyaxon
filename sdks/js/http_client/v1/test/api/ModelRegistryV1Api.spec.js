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
    instance = new PolyaxonSdk.ModelRegistryV1Api();
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

  describe('ModelRegistryV1Api', function() {
    describe('createModelRegistry', function() {
      it('should call createModelRegistry successfully', function(done) {
        //uncomment below and update the code to test createModelRegistry
        //instance.createModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteModelRegistry', function() {
      it('should call deleteModelRegistry successfully', function(done) {
        //uncomment below and update the code to test deleteModelRegistry
        //instance.deleteModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getModelRegistry', function() {
      it('should call getModelRegistry successfully', function(done) {
        //uncomment below and update the code to test getModelRegistry
        //instance.getModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listModelRegistry', function() {
      it('should call listModelRegistry successfully', function(done) {
        //uncomment below and update the code to test listModelRegistry
        //instance.listModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listModelRegistryNames', function() {
      it('should call listModelRegistryNames successfully', function(done) {
        //uncomment below and update the code to test listModelRegistryNames
        //instance.listModelRegistryNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchModelRegistry', function() {
      it('should call patchModelRegistry successfully', function(done) {
        //uncomment below and update the code to test patchModelRegistry
        //instance.patchModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateModelRegistry', function() {
      it('should call updateModelRegistry successfully', function(done) {
        //uncomment below and update the code to test updateModelRegistry
        //instance.updateModelRegistry(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
