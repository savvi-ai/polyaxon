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
 * OpenAPI spec version: 1.0.73
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1Container', 'model/V1Environment', 'model/V1Init', 'model/V1Volume'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./V1Container'), require('./V1Environment'), require('./V1Init'), require('./V1Volume'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.V1Service = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1Container, root.PolyaxonSdk.V1Environment, root.PolyaxonSdk.V1Init, root.PolyaxonSdk.V1Volume);
  }
}(this, function(ApiClient, V1Container, V1Environment, V1Init, V1Volume) {
  'use strict';

  /**
   * The V1Service model module.
   * @module model/V1Service
   * @version 1.0.73
   */

  /**
   * Constructs a new <code>V1Service</code>.
   * @alias module:model/V1Service
   * @class
   */
  var exports = function() {
  };

  /**
   * Constructs a <code>V1Service</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/V1Service} obj Optional instance to populate.
   * @return {module:model/V1Service} The populated <code>V1Service</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();
      if (data.hasOwnProperty('kind'))
        obj.kind = ApiClient.convertToType(data['kind'], 'String');
      if (data.hasOwnProperty('environment'))
        obj.environment = V1Environment.constructFromObject(data['environment']);
      if (data.hasOwnProperty('connections'))
        obj.connections = ApiClient.convertToType(data['connections'], ['String']);
      if (data.hasOwnProperty('volumes'))
        obj.volumes = ApiClient.convertToType(data['volumes'], [V1Volume]);
      if (data.hasOwnProperty('init'))
        obj.init = ApiClient.convertToType(data['init'], [V1Init]);
      if (data.hasOwnProperty('sidecars'))
        obj.sidecars = ApiClient.convertToType(data['sidecars'], [V1Container]);
      if (data.hasOwnProperty('container'))
        obj.container = V1Container.constructFromObject(data['container']);
      if (data.hasOwnProperty('ports'))
        obj.ports = ApiClient.convertToType(data['ports'], ['Number']);
      if (data.hasOwnProperty('rewritePath'))
        obj.rewritePath = ApiClient.convertToType(data['rewritePath'], 'Boolean');
    }
    return obj;
  }

  /**
   * @member {String} kind
   * @default 'service'
   */
  exports.prototype.kind = 'service';

  /**
   * @member {module:model/V1Environment} environment
   */
  exports.prototype.environment = undefined;

  /**
   * @member {Array.<String>} connections
   */
  exports.prototype.connections = undefined;

  /**
   * Volumes is a list of volumes that can be mounted.
   * @member {Array.<module:model/V1Volume>} volumes
   */
  exports.prototype.volumes = undefined;

  /**
   * @member {Array.<module:model/V1Init>} init
   */
  exports.prototype.init = undefined;

  /**
   * @member {Array.<module:model/V1Container>} sidecars
   */
  exports.prototype.sidecars = undefined;

  /**
   * @member {module:model/V1Container} container
   */
  exports.prototype.container = undefined;

  /**
   * @member {Array.<Number>} ports
   */
  exports.prototype.ports = undefined;

  /**
   * Rewrite path to remove polyaxon base url(i.e. /v1/services/namespace/owner/project/). Default is false, the service shoud handle a base url.
   * @member {Boolean} rewritePath
   */
  exports.prototype.rewritePath = undefined;

  return exports;

}));
